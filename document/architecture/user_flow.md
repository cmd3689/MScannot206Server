# 유저/캐릭터 관리

본 문서에서는 유저의 캐릭터 관리와 관련된 흐름을 설명합니다.

## 목차

- [1. 캐릭터 생성 (Create Character Flow)](#1-캐릭터-생성-create-character-flow)
- [2. 캐릭터 제거 (Delete Character Flow)](#2-캐릭터-제거-delete-character-flow)

## 1. 캐릭터 생성 (Create Character Flow)

```mermaid
sequenceDiagram
    autonumber
    participant C as Client
    participant H as Handler
    participant S as Service
    participant R as Repository

    Note over C, H: Request
    C->>+H: POST /user/character/create
    Note over H, R: 유저 토큰 유효성 검사(중략)
    alt 토큰 유효
        H->>S: 캐릭터 조회 요청
        S->>+R: 캐릭터 조회
        R-->>-S: 캐릭터 리스트 반환
        S->>H: 캐릭터 리스트 반환
        H->>H: 캐릭터 생성 가능 여부 판단
        alt 캐릭터 생성 가능
            H->>+S: 캐릭터 생성 요청
            S->>+R: 캐릭터 생성
            Note over R: Database start
            R->>R: 생성 할 캐릭터 이름 저장
            R->>R: 유저-캐릭터 매핑 저장
            Note over R: Database end
            R-->>-S: 생성된 캐릭터 정보 반환
            S-->>-H: 생성된 캐릭터 정보 반환
        else 캐릭터 생성 불가
            H->>H: 오류 응답 추가
        end
    else 유효하지 않은 토큰
        H->>H: 오류 응답 추가
    end
    Note over H: 생성 된 캐릭터 정보 포함
    H->>H: 응답 생성
    
    Note over H, C: Response
    H-->>-C: 200 OK
```

## 2. 캐릭터 삭제 (Delete Character Flow)

```mermaid
sequenceDiagram
    autonumber
    participant C as Client
    participant H as Handler
    participant S as Service
    participant R as Repository

    Note over C, H: Request
    C->>+H: POST /user/character/delete
    Note over H, R: 유저 토큰 유효성 검사(중략)
        alt 토큰 유효
        H->>S: 캐릭터 조회 요청
        S->>+R: 캐릭터 조회
        R-->>-S: 캐릭터 리스트 반환
        S->>H: 캐릭터 리스트 반환
        H->>H: 캐릭터 삭제 가능 여부 판단
        alt 캐릭터 삭제 가능
            Note over H, S: 삭제 요청한 캐릭터의 Slot, Name 정보 포함
            H->>+S: 캐릭터 삭제 요청
            S->>+R: 캐릭터 삭제
            Note over R: Database start
            R->>R: 유저-캐릭터 매핑 삭제
            R->>R: 삭제 된 캐릭터 이름 삭제
            Note over R: Database end
            R-->>-S: 캐릭터 삭제 결과 반환
            S-->>-H: 캐릭터 삭제 결과 반환
        else 캐릭터 삭제 불가
            H->>H: 오류 응답 추가
        end
    else 유효하지 않은 토큰
        H->>H: 오류 응답 추가
    end
    H->>H: 응답 생성

    Note over H, C: Response
    H-->>-C: 200 OK

```
