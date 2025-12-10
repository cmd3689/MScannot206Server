package mongodb

import (
	"MScannot206/pkg/shared"
	"MScannot206/shared/repository"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mongo_options "go.mongodb.org/mongo-driver/mongo/options"
)

func NewUserRepository(
	ctx context.Context,
	client *mongo.Client,
	dbName string,
) (*UserRepository, error) {

	if client == nil {
		return nil, errors.New("mongo client is null")
	}

	repo := &UserRepository{
		client: client,
	}

	repo.user = client.Database(dbName).Collection(shared.User)

	return repo, nil
}

type UserRepository struct {
	ctx    context.Context
	client *mongo.Client

	user *mongo.Collection
}

type userDoc struct {
	*repository.User `bson:",inline"`
	CreatedAt        time.Time `bson:"createdAt"`
}

func (r *UserRepository) Start() error {
	return nil
}

func (r *UserRepository) Stop() error {
	return nil
}

func (r *UserRepository) FindUserByUID(uid string) (*repository.User, error) {
	var doc userDoc

	filter := bson.D{
		{Key: "uid", Value: uid},
	}

	projection := bson.D{
		{Key: "uid", Value: 1},
	}

	opts := mongo_options.FindOne().
		SetProjection(projection)

	ret := r.user.FindOne(r.ctx, filter, opts)

	if ret.Err() == nil {
		if err := ret.Decode(&doc); err != nil {
			return nil, err
		}
	} else if ret.Err() == mongo.ErrNoDocuments {
		// 유저 데이터가 없는 경우 null 반환
		return nil, nil
	} else {
		return nil, ret.Err()
	}

	return doc.User, nil
}
