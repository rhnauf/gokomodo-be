package helper

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	Seller              = "seller"
	Buyer               = "Buyer"
	ErrNilClaims        = errors.New("claims not found")
	ErrCastingClaims    = errors.New("invalid claims")
	ErrMarshallClaims   = errors.New("invalid claims")
	ErrUnmarshallClaims = errors.New("invalid claims")
)

type CustomClaim struct {
	jwt.RegisteredClaims
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
}

func GenerateJWT(userId int, email string, jwtType string) (string, error) {
	// for simplicity it will be hardcoded
	jwtDuration := time.Duration(2) * time.Hour

	claim := &CustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtDuration)),
		},
		UserId: userId,
		Email:  email,
	}

	var privateKey string

	switch jwtType {
	case Seller:
		privateKey = os.Getenv("PRIVATE_KEY_SELLER")
	case Buyer:
		privateKey = os.Getenv("PRIVATE_KEY_BUYER")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	s, err := t.SignedString([]byte(privateKey))
	if err != nil {
		return "", err
	}

	return s, nil
}

func UnmarshallClaim(ctx context.Context) (*CustomClaim, error) {
	c := ctx.Value("claims")
	if c == nil {
		return nil, ErrNilClaims
	}

	cc, ok := c.(jwt.MapClaims)
	if !ok {
		return nil, ErrCastingClaims
	}

	jsonString, err := json.Marshal(cc)
	if err != nil {
		return nil, ErrMarshallClaims
	}

	var customClaim CustomClaim
	if err := json.Unmarshal([]byte(jsonString), &customClaim); err != nil {
		return nil, ErrUnmarshallClaims
	}

	return &customClaim, nil
}
