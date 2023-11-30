package security

import (
	"fmt"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/config"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user model.User) (string, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return "", fmt.Errorf("Failed to create access token : %s", err.Error())
	}

	now := time.Now().UTC()
	end := now.Add(cfg.AccessTokenLifeTime)

	claims := &TokenMyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Id: user.Id,
	}

	token := jwt.NewWithClaims(cfg.JwtSigningMethod, claims)
	ss, err := token.SignedString(cfg.JwtSignatureKey)
	if err != nil {
		return "", fmt.Errorf("Failed to create access token : %s", err.Error())
	}
	return ss, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("Failed to verify access token : %s", err.Error())
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != cfg.JwtSigningMethod {
			return nil, fmt.Errorf("Invalid token signing method")
		}
		return cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Invalid parse token sdf : %s", err.Error())
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("Invalid token MapClaims")
	}
	return claims, nil
}

func GetIdFromToken(c *gin.Context) (string, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return "", fmt.Errorf("Failed to get id from token")
	}

	claimsMap := claims.(jwt.MapClaims)
	id := claimsMap["id"].(string)

	return id, nil
}
