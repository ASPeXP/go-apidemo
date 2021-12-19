package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)
func Protect(tokenString string) error {
	_, err := jwt.Parse (tokenString, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			retur nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("===signature==="), nil 
	})

	return err
}