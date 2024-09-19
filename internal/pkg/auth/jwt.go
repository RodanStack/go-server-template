package auth

import (
	"go-server-template/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	env *config.Env
}

func newJWT(env *config.Env) *JWT {
	return &JWT{
		env: env,
	}
}

// Generate generates a new JWT token.
func (j *JWT) Generate(userID int64) (string, error) {
	tokenExpirationHours := 24
	tokenExpiration := time.Now().Add(time.Hour * time.Duration(tokenExpirationHours)).Unix()

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     tokenExpiration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.env.JWTSecretKey))
}

// Validate validates a JWT token.
func (j *JWT) Validate(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKey
		}
		return j.env.JWTSecretKey, nil
	})

	return token, err
}
