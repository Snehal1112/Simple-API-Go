package models

import (
		"github.com/dgrijalva/jwt-go"
	)

type ClaimsData struct {
    JwtData `boil:"data" json:"data" toml:"data" yaml:"data"`
    jwt.StandardClaims
}
