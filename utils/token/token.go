package token

import(
	"github.com/dgrijalva/jwt-go"
	"github.com/yang-f/beauty/settings"
	"time"
	"fmt"
)

func Generate(user_id int) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"user_id": user_id,
	"exp": (time.Now().Add(time.Minute*60*24*2)).Unix(),
	})

	tokenString, err := token.SignedString(settings.HmacSampleSecret)
	return tokenString, err
}

func Valid(tokenString string) (int,error){
	token1, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return settings.HmacSampleSecret, nil
	})

	if claims, ok := token1.Claims.(jwt.MapClaims); ok && token1.Valid {
		return int(claims["user_id"].(float64)), nil 
	} else {
		return -1, err
	}
}

