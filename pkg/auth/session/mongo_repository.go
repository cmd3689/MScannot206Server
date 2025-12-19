package session

import (
	"MScannot206/shared"
	"MScannot206/shared/entity"
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 7주일
const expireAfterSeconds = 3600 * 24 * 7

var ErrSessionRepositoryIsNil = errors.New("session repository is null")

func NewSessionRepository(
	ctx context.Context,
	client *mongo.Client,
	dbName string,
) (*SessionRepository, error) {

	if client == nil {
		return nil, errors.New("mongo client is null")
	}

	repo := &SessionRepository{
		client:  client,
		session: client.Database(dbName).Collection(shared.UserSession),
	}

	if err := repo.ensureIndexes(ctx); err != nil {
		return nil, err
	}

	return repo, nil
}

type SessionRepository struct {
	client  *mongo.Client
	session *mongo.Collection
}

func (r *SessionRepository) ensureIndexes(ctx context.Context) error {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "updated_at", Value: 1},
		},
		Options: options.Index().
			SetExpireAfterSeconds(expireAfterSeconds).
			SetName("session_ttl_idx"),
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err := r.session.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) SaveUserSessions(ctx context.Context, sessions []*entity.UserSession) error {
	if len(sessions) == 0 {
		return nil
	}

	models := make([]mongo.WriteModel, len(sessions))

	for i, session := range sessions {
		session.UpdatedAt = time.Now().UTC()

		filter := bson.D{
			{Key: "_id", Value: session.Uid},
		}

		update := bson.D{
			{Key: "$set", Value: session},
		}

		models[i] = mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)
	}

	_, err := r.session.BulkWrite(ctx, models, options.BulkWrite().SetOrdered(false))
	if err != nil {
		if bulkErr, ok := err.(mongo.BulkWriteException); ok {
			log.Warn().Msg("일부 세션 저장에 실패했습니다")
			for _, writeErr := range bulkErr.WriteErrors {
				log.Warn().Msgf("세션 저장 실패: %v", writeErr.Message)
			}
			return nil
		}
		return err
	}

	return nil
}

func (r *SessionRepository) ValidateUserSessions(ctx context.Context, sessions []*entity.UserSession) ([]string, []string, error) {
	if len(sessions) == 0 {
		return []string{}, []string{}, nil
	}

	sessionCount := len(sessions)
	validUids := make([]string, 0, sessionCount)
	invalidUids := make([]string, 0, (sessionCount / 2))

	targetUids := make([]string, 0, sessionCount)
	inputSessionMap := make(map[string]string, sessionCount)

	for _, s := range sessions {
		targetUids = append(targetUids, s.Uid)
		inputSessionMap[s.Uid] = s.Token
	}

	filter := bson.M{
		"_id": bson.M{"$in": targetUids},
	}

	cursor, err := r.session.Find(ctx, filter)
	if err != nil {
		return []string{}, []string{}, err
	}
	defer cursor.Close(ctx)

	var dbSessions []*entity.UserSession
	if err := cursor.All(ctx, &dbSessions); err != nil {
		return []string{}, []string{}, err
	}

	dbSessionMap := make(map[string]*entity.UserSession, len(dbSessions))
	for _, s := range dbSessions {
		dbSessionMap[s.Uid] = s
	}

	for _, uid := range targetUids {
		s, ok := dbSessionMap[uid]
		if ok && s.Token == inputSessionMap[uid] {
			validUids = append(validUids, uid)
		} else {
			invalidUids = append(invalidUids, uid)
		}
	}

	return validUids, invalidUids, nil
}
