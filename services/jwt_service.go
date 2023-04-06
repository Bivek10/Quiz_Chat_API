package services

import (
	"strings"
	"time"

	"github.com/bivek/fmt_backend/errors"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	logger infrastructure.Logger
	env    infrastructure.Env
}

type JWTClamis struct {
	Email string
	jwt.StandardClaims
}

func NewJwtService(logger infrastructure.Logger, env infrastructure.Env) JWTService {
	return JWTService{
		logger: logger,
		env:    env,
	}
}

func (s JWTService) ParseToken(clientToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(clientToken, &JWTClamis{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTSecretKey), nil
	})
	if err != nil {
		if !strings.Contains(err.Error(), "expired") {
			s.logger.Zap.Error("Invalid token[ParseWithClaims] :", err.Error())
			err := errors.BadRequest.New("Invalid ID token")
			return nil, err
		}
		s.logger.Zap.Error("Invalid token[ParseWithClaims] :", err.Error())
		return nil, err
	}
	println(token.Valid)

	return token, nil
}

func (m JWTService) VerifyToken(c *gin.Context) (bool, error) {
	// Get the token from the request header
	header := c.GetHeader("Authorization")

	if header == "" {
		err := errors.BadRequest.New("Authorization token is required in header")
		err = errors.SetCustomMessage(err, "Authorization token is required in header")
		m.logger.Zap.Error("[GetHeader]: ", err.Error())
		return false, err
	}

	if !strings.Contains(header, "Bearer") {
		err := errors.BadRequest.New("Token type is required")
		m.logger.Zap.Error("Missing token type: ", err.Error())
		return false, err
	}

	tokenString := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
	token, err := m.ParseToken(tokenString)

	if err != nil {
		m.logger.Zap.Error("Error parsing token", err.Error())
		return false, err
	}

	claims, ok := token.Claims.(*JWTClamis)
	if !ok || !token.Valid {
		err := errors.BadRequest.New("Invalid token")
		err = errors.SetCustomMessage(err, "Invalid token")
		m.logger.Zap.Error("Invalid token [token.Valid]: ", err.Error())
		return false, err
	}
	// Can set anything in the request context and passes the request to the next handler.
	c.Set("email", claims.Email)
	return true, nil

}

func (m JWTService) VerifyRefreshToken(refreshToken string, c *gin.Context) (bool, error) {
	token, err := m.ParseToken(refreshToken)
	if err != nil {
		m.logger.Zap.Error("Error parsing token", err.Error())
		return false, err
	}

	claims, ok := token.Claims.(*JWTClamis)
	println(token.Valid)
	if !ok || !token.Valid {
		err := errors.BadRequest.New("Invalid token")
		err = errors.SetCustomMessage(err, "Invalid token")
		m.logger.Zap.Error("Invalid token [token.Valid]: ", err.Error())
		return false, err
	}

	c.Set("email", claims.Email)
	return true, nil
}

func (m JWTService) GenerateJWT(email string) (string, error, string, error) {

	clams := &JWTClamis{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	refreshClams :=
		&JWTClamis{
			Email: email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 48).Unix(),
			},
		}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, clams).SignedString([]byte(m.env.JWTSecretKey))
	refreshtoken, refresherr := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClams).SignedString([]byte(m.env.JWTSecretKey))
	return accessToken, err, refreshtoken, refresherr
}
