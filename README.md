# Identity-Server на Golang
### Using SHA256 hashing for passwords and JWT-based authorization 
### Endpoints with parameters:
1. "auth/register":
⋅*Parameters:
  {
    "full_name":        "string",
    "e_mail":           "string",
    "nick_name":        "string",
    "age":              "integer",
    "password":         "string",
    "password_confirm": "string",
  }
 ⋅*Responses: 201 (Creeated) / 400 (Bad request)
2. "auth/login":
⋅⋅*Parameters:
  {
    "nick_name":        "string",
    "password":         "string",
  }
 ⋅⋅*Responses: 200 (Creeated) / 401 (Unauthorized) / 403 (Forbidden)
 ⋅⋅*Returns:
  {
     "jwt_token":     "string"
     "refresh_token": "string"
  }
3. "auth/welcome":
⋅⋅⋅*Parameters: Header-Authorization "Bearer {token}"
⋅⋅⋅*Responses: 200 (Creeated) / 400 (Bad request) / 500 (Internal)
⋅⋅⋅*Returns:
  {
   "id":                     "string"
   "full_name":              "string",
    "e_mail":                "string",
    "nick_name":             "string",
    "age":                   "integer",
    "password_hash":         "string",
}
4. "auth/reauth"
   ⋅⋅*Parameters:
  {
    "refresh_token": "string"
  }
 ⋅⋅*Responses: 200 (Creeated) / 400 (Bad request) / 500 (Internal)
 ⋅⋅*Returns:
  {
     "jwt_token":     "string"
     "refresh_token": "string"
  }
