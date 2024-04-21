package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/rhnauf/gokomodo-be/internal/helper"
)

func AuthJWTSeller(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqHeader := r.Header.Get("Authorization")
		if !strings.Contains(reqHeader, "Bearer") {
			helper.HandleResponse(w, http.StatusUnauthorized, "missing authorization token", nil)
			return
		}

		tokenString := strings.Replace(reqHeader, "Bearer ", "", -1)

		claims, err := verifyToken(tokenString, helper.Seller)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				helper.HandleResponse(w, http.StatusUnauthorized, "token is expired", nil)
				return
			}
			log.Println("error verifying token =>", err)
			helper.HandleInternalServerError(w)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthJWTBuyer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqHeader := r.Header.Get("Authorization")
		if !strings.Contains(reqHeader, "Bearer") {
			helper.HandleResponse(w, http.StatusUnauthorized, "missing authorization token", nil)
			return
		}

		tokenString := strings.Replace(reqHeader, "Bearer ", "", -1)

		claims, err := verifyToken(tokenString, helper.Buyer)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				helper.HandleResponse(w, http.StatusUnauthorized, "token is expired", nil)
				return
			}
			log.Println("error verifying token =>", err)
			helper.HandleInternalServerError(w)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func verifyToken(tokenString string, jwtType string) (jwt.MapClaims, error) {
	var privateKey string

	switch jwtType {
	case helper.Seller:
		privateKey = os.Getenv("PRIVATE_KEY_SELLER")
	case helper.Buyer:
		privateKey = os.Getenv("PRIVATE_KEY_BUYER")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
