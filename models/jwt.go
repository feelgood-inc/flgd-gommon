package models

import "github.com/golang-jwt/jwt"

type FeelgoodJWTPayload struct {
	UID       string   `json:"uid"`
	Email     string   `json:"email"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Type      string   `json:"type"`
	Roles     []string `json:"roles"`
}

type FeelgoodJWTCustomClaims struct {
	Email     string   `json:"email"`
	Type      string   `json:"type"`
	UID       string   `json:"uid"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Roles     []string `json:"roles"`
}

type FeelgoodJWTClaims struct {
	Claims FeelgoodJWTCustomClaims `json:"claims"`
	jwt.StandardClaims
}
