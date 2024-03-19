package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	LifeExpire  = 24 * time.Hour //会话过期为一天
	ShareSecret = "c064ecf48381d76e"
)

type JwtService struct {
	Uid      uint32
	NickName string
	UserName string
	Ip       string
	jwt.RegisteredClaims
}

func (c *JwtService) buildClaims() {
	now := time.Now()
	c.ExpiresAt = jwt.NewNumericDate(now.Add(LifeExpire)) //Expiration time
	c.IssuedAt = jwt.NewNumericDate(now)                  //Issuing time
	c.NotBefore = jwt.NewNumericDate(now)                 //Begin Effective time
}

func (c *JwtService) CreateToken() (string, error) {
	c.buildClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(ShareSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (c *JwtService) getKeyFun() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		return []byte(ShareSecret), nil
	}
}

func (c *JwtService) ParseToken(t string) error {
	token, err := jwt.ParseWithClaims(t, c, c.getKeyFun())
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*JwtService); ok && token.Valid {
		c = claims
		return nil
	}
	return errors.New("token not active yet")
}
