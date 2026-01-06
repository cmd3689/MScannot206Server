# 로그인/인증

본 문서에서는 서버의 인증 흐름에 대해 설명합니다. 주요 내용은 다음과 같습니다.

## 목차

- [1. 유저 토큰 인증 (User Token Authentication)](#1-유저-토큰-인증)
- [2. 로그인 (Login Flow)](#2-로그인-login-flow)

## 1. 유저 토큰 인증 (Authenticate Flow)

유저 세션의 토큰을 검증하는 흐름은 다음과 같습니다.

```mermaid
sequenceDiagram
    autonumber
    participant S as Service
    participant R as Repository
    participant DB as Database

    activate S
    S->>+R: 유저 세션에서 토큰 검증
    R->>+DB: 유저 uid로 유저 세션 조회
    DB-->>-R: 유저 세션 데이터 반환
    R->>R: 토큰 일치 여부 확인
    R->>-S: 유효한 토큰 여부 반환
    deactivate S

```

## 2. 로그인 (Login Flow)

클라이언트가 유저들의 로그인 요청을 보내면 서버는 다음 절차를 거쳐 인증을 수행합니다.

```mermaid
sequenceDiagram
    autonumber
    participant C as Client
    participant H as Handler
    participant S as Service
    participant R as Repository
    
    Note over C, H: Request
    C->>+H: POST /login
    H->>+S: 로그인 요청

    S->>+R: 유저 조회
    R-->>-S: User Entity (or nil)
    
    alt 유저 정보 없음
        S->>+R: 유저 생성
        R-->>-S: New User Entity
    end

    S->>-H: 로그인 될 유저 정보 반환
    H->>+S: 유저 세션 생성 요청

    S->>+R: 세션 저장
    R-->>-S: 세션 저장 결과

    S-->>-H: 유저 세션 반환
    Note over H: 유저 정보 및 세션 토큰 결합
    H->>H: 응답 생성
    
    Note over H, C: Response
    H-->>-C: 200 OK
```
