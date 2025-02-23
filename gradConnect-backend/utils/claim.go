package utils

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID uint     `json:"user_id"`
	FullName  string   `json:"fullname"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}