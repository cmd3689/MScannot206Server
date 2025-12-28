# 🔐 Login API

로그인, 세션 갱신 등 인증과 관련된 API 명세입니다.

## 목차
- [로그인 (Login)](#로그인-login)

---

### 로그인 (Login)
사용자의 계정 정보를 검증하고 세션 토큰을 발급합니다.

> **Endpoint**

| Method | URL |
| :---: | :--- |
| ![POST](https://img.shields.io/badge/POST-orange?style=for-the-badge) | `/api/v1/login` |

> **Request Body**

| Field | Type | Required | Description |
| :--- | :---: | :---: | :--- |
| `uids` | Array | ✅ | 사용자 고유 ID |

**Example:**
```json
{
  "uids": ["12345678900000000", "12345678900000001", "12345678900000002", "12345678900000003"]
}
```

> **Response Fields**

| Field | Type | Required | Description |
| :--- | :---: | :---: | :--- |
| `success_uids` | Array | ✅ | 성공한 유저 리스트 |
| `success_uids[].user_entity` | Object | ✅ | 유저 상세 정보 객체 |
| `success_uids[].user_entity.uid` | String | ✅ | 유저 고유 식별자 |
| `success_uids[].user_entity.characters` | Array | ✅ | 보유 캐릭터 리스트 |
| `fail_uids` | Array | ✅ | 실패한 유저 리스트 |
| `fail_uids[].uid` | String | ✅ | 실패한 유저의 UID |
| `fail_uids[].error_code` | String | ❌ | 실패 사유 (에러 코드) |

**Example:**

**Success (200 OK)**
```json
{
  "data": {
    "success_uids": [
        {
            "user_entity": {
                "uid": "12345678900000000",
                "characters": [...],
            },
        },
        {
            "user_entity": {
                "uid": "12345678900000001",
                "characters": [...],
            },
        }
    ],
    "fail_uids": [
        {
            "uid": "12345678900000002",
            "error_code": "User Error Message..."
        },
        {
            "uid": "12345678900000003",
            "error_code": "User Error Message..."
        }
    ],
  }
}
```
---
