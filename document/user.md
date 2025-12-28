# 👤 User API

유저 정보 조회, 캐릭터 생성 및 삭제 등 유저 데이터와 관련된 API 명세입니다.

## 목차
- [캐릭터 생성](#캐릭터-생성)
- [캐릭터 삭제](#캐릭터-삭제)

---

### 캐릭터 생성
빈 슬롯에 새로운 캐릭터를 생성합니다.

> **Endpoint**

| Method | URL |
| :---: | :--- |
| ![POST](https://img.shields.io/badge/POST-orange?style=for-the-badge) | `/api/v1/user/character/create` |

> **Request Body**

| Field | Type | Required | Description |
| :--- | :---: | :---: | :--- |
| `requests` | Array | ✅ | 캐릭터 생성 요청 리스트 |
| `requests[].uid` | String | ✅ | 사용자 고유 ID |
| `requests[].token` | String | ✅ | 사용자 인증 토큰 |
| `requests[].slot` | Integer | ✅ | 캐릭터 슬롯 번호 (1~3) |
| `requests[].name` | String | ✅ | 캐릭터 이름 (특수문자 불가) |

**Example:**
```json
{
  "requests": [
    {
      "uid": "12345678900000000",
      "token": "user_session_token",
      "slot": 1,
      "name": "토벤머리"
    },
    {
      "uid": "12345678900000001",
      "token": "user_session_token",
      "slot": 2,
      "name": "더벅머리"
    },
    {
      "uid": "12345678900000002",
      "token": "user_session_token",
      "slot": 3,
      "name": "토벤머리"
    }
  ]
}
```

> **Response Fields**
| Field | Type | Required | Description |
| :--- | :---: | :---: | :--- |
| `responses` | Array | ✅ | 캐릭터 생성 결과 리스트 |
| `responses[].uid` | String | ✅ | 사용자 고유 ID |
| `responses[].slot` | Integer | ❌ | 생성된 캐릭터의 슬롯 번호 |
| `responses[].error_code` | String | ❌ | 실패 사유 (에러 코드) |

**Example:**
**Success (200 OK)**
```json
{
  "data": {
    "responses": [
      {
        "uid": "12345678900000000"
        "slot": 1
      },
      {
        "uid": "12345678900000001"
        "slot": 2
      },
      {
        "uid": "12345678900000002",
        "error_code": "USER_CHARACTER_NAME_ALREADY_EXISTS_ERROR"
      }
    ]
  }
}
```

---

### 캐릭터 이름 중복 확인
특정 이름이 이미 사용 중인지 확인합니다.

> **Endpoint**

| Method | URL |
| :---: | :--- |
| ![POST](https://img.shields.io/badge/POST-orange?style=for-the-badge) | `/api/v1/user/character/create/check_name` |

> **Request Body**

| Field | Type | Required | Description |
| :--- | :---: | :---: | :--- |
| `requests` | Array | ✅ | 이름 중복 확인 요청 리스트 |
| `requests[].uid` | String | ✅ | 사용자 고유 ID |
| `requests[].token` | String | ✅ | 사용자 인증 토큰 |
| `requests[].name` | String | ✅ | 확인할 캐릭터 이름 |

**Example:**
```json
{
  "requests": [
    {
      "uid": "12345678900000000",
      "token": "user_session_token",
      "name": "토벤머리"
    },
    {
      "uid": "12345678900000001",
      "token": "user_session_token",
      "name": "더벅머리"
    },
    {
      "uid": "12345678900000002",
      "token": "user_session_token",
      "name": "토벤머리"
    }
  ]
}
```

> **Response Fields**

| Field | Type | Required | Description |
| :--- | :---: | :---: | :--- |
| `responses` | Array | ✅ | 이름 중복 확인 결과 리스트 |
| `responses[].uid` | String | ✅ | 사용자 고유 ID |
| `responses[].available` | Boolean | ✅ | 이름 사용 가능 여부 |
| `responses[].error_code` | String | ❌ | 실패 사유 (에러 코드) |

**Example:**
**Success (200 OK)**
```json
{
  "data": {
    "responses": [
      {
        "uid": "12345678900000000",
        "available": true
      },
      {
        "uid": "12345678900000001",
        "available": true
      },
      {
        "uid": "12345678900000002",
        "available": false
        "error_code": "USER_CHARACTER_NAME_ALREADY_EXISTS_ERROR"
      }
    ]
  }
}
```

---

### 캐릭터 삭제
특정 슬롯의 캐릭터를 삭제합니다.
*이 작업은 되돌릴 수 없습니다.*

> **Endpoint**

| Method | URL |
| :---: | :--- |
| ![POST](https://img.shields.io/badge/POST-orange?style=for-the-badge) | `/api/v1/user/character/delete` |

> **Request Body**

| Field | Type | Required | Description |
| :--- | :---: | :---: | :--- |
| `requests` | Array | ✅ | 캐릭터 삭제 요청 리스트 |
| `requests[].uid` | String | ✅ | 사용자 고유 ID |
| `requests[].token` | String | ✅ | 사용자 인증 토큰 |
| `requests[].slot` | Integer | ✅ | 삭제할 캐릭터의 슬롯 번호 (1~3) |

**Example:**
```json
{
  "requests": [
    {
      "uid": "12345678900000000",
      "token": "user_session_token",
      "slot": 1
    },
    {
      "uid": "12345678900000001",
      "token": "user_session_token",
      "slot": 2
    },
    {
      "uid": "12345678900000002",
      "token": "user_session_token",
      "slot": 1
    }
  ]
}
```

> **Response Fields**

| Field | Type | Required | Description |
| :--- | :---: | :---: | :--- |
| `responses` | Array | ✅ | 캐릭터 삭제 결과 리스트 |
| `responses[].uid` | String | ✅ | 사용자 고유 ID |
| `responses[].error_code` | String | ❌ | 실패 사유 (에러 코드) |

**Example:**
**Success (200 OK)**
```json
{
  "data": {
    "responses": [
      {
        "uid": "12345678900000000"
      },
      {
        "uid": "12345678900000001"
      },
      {
        "uid": "123456789000000002",
        "error_code": "USER_CHARACTER_SLOT_EMPTY_ERROR"
      }
    ]
  }
}
```
