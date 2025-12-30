# MScannot206Server&nbsp;![Go](https://img.shields.io/badge/Language-Go-00ADD8?style=flat&logo=go&logoColor=white) ![Go Version](https://img.shields.io/badge/Version-1.25.4-00ADD8?style=flat&logo=go&logoColor=white)

이 프로젝트는 [MScannot206](https://github.com/dek0058/MScannot206) 클라이언트를 보조하기 위한 콘솔 서버 입니다.

[메이플스토리 월드 크리에이터 이용약관](https://github.com/dek0058/MScannot206)을 준수하며, 해당 프로젝트는 비공식 프로젝트임을 알립니다.


## 목차

- [📋 요구사항](#-요구사항)
- [📚 API Documentation](#-api-documentation)
- [🏗️ 아키텍처](#️-아키텍처)
- [🖥️ 테스트 클라이언트](#️-테스트-클라이언트)

## 📋 요구사항

- [Go](https://go.dev/doc/install)
- [MongoDB](https://www.mongodb.com/try/download/community)


## 📚 API Documentation

상세한 API 명세는 아래 문서들을 참고해주세요.

- [🔐 로그인/인증 API (Login)](document/api/login.md)
- [👤 유저/캐릭터 API (User)](document/api/user.md)

## 🏗️ 아키텍처

### 메인 플로우

```mermaid
graph TD
    classDef user fill:#ffffff,stroke:#333,stroke-width:2px,color:#000000,font-weight:bold;
    classDef client fill:#E3F2FD,stroke:#1565C0,stroke-width:2px,color:#000000,font-weight:bold;
    classDef server fill:#E8F5E9,stroke:#2E7D32,stroke-width:2px,color:#000000,font-weight:bold;
    classDef db fill:#FFF3E0,stroke:#EF6C00,stroke-width:2px,color:#000000,font-weight:bold;

    User((User)):::user
    Client[Client]:::client

    subgraph Server_Area [Server Side]
        direction TB
        Handlers[Handler]:::server
        Services[Service]:::server
        Repositories[Repository]:::server
    end

    subgraph Data_Area [Persistence Layer]
        DB[("Database")]:::db
    end

    User--->|1.Connect|Client
    Client -->|2.API Request| Handlers
    Handlers -->|3.Call Method| Services
    Services -->|4.Request Data Access| Repositories
    Repositories -->|5.Query| DB
    DB -.->|6.Result| Repositories
    Repositories -.->|7.Return Entity/Model| Services
    Services -.->|8.Return DTO/Result| Handlers
    Handlers -.->|9.API Response| Client
```

### 상세 플로우
- [로그인/인증](document/architecture/auth_flow.md) - 서버의 로그인 및 인증 처리 흐름
- [유저/캐릭터 관리](document/architecture/user_flow.md) - 유저의 캐릭터 생성, 삭제 등의 처리 흐름

## 🖥️ 테스트 클라이언트

테스트 목적으로 제작된 간단한 콘솔 클라이언트가 포함되어 있습니다. 해당 클라이언트는 `pkg/testclient` 디렉토리에서 확인할 수 있습니다.

