package models

import (
	"github.com/dgrijalva/jwt-go"
)
type User struct {
	Name   string             `json:"name,omitempty" bson:"name, omitempty"`
	Password  string             `json:"album,omitempty" bson:"album, omitempty"`
	jwtKey   jwt.StandardClaims
}
