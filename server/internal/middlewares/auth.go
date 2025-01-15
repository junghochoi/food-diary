package auth

import (
	"fmt"
	"food-diary/internal/response"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
)

// Pass in the jwtSecret into auth middle ware
func AuthMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			jwtVerifyLogic(w, req, next, jwtSecret)
		})
	}
}

// Function to Verify the JWT token
func jwtVerifyLogic(w http.ResponseWriter, req *http.Request, next http.Handler, jwtSecret string) {
	log.Printf("Hitting the Auth Middleware\n")
	token := req.Header.Get("Authorization")
	if token == "" {
		response.Error(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	email, err := parseJWTToken(token, []byte(jwtSecret))
	if err != nil {
		log.Printf("Error parsing token: %s", err)
		response.Error(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	log.Printf("Received request from %s", email)
	next.ServeHTTP(w, req)
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func parseJWTToken(token string, hmacSecret []byte) (payload *Claims, err error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	// Check if the token is valid
	if err != nil {
		return nil, fmt.Errorf("Error Validating Token: %v", err)
	} else if claims, ok := t.Claims.(*Claims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("Error Parsing Token: %v", err)
}
