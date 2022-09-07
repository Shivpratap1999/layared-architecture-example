package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//unexported type
type jwtclaims struct {
	userInformation string
	jwt.StandardClaims
}

//TokenProvider perform authorization process , it return string of *jwt.Token and error
func TokenProvider(userId string, jwtkey string) (jwtTokenstr string, err error) {
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = expirationTime
	log.Println("[jwt]access token generation ")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtTokenstr, err = token.SignedString([]byte(jwtkey))
	if err != nil {
		return "", err
	}
	return jwtTokenstr, nil
}

//TokenValidator perform Authentication process.
//It validate jwtToken string and return validation result (bool) and messege string.
func ValidateToken(jwtTokenStr string, jwtkey string) (status bool, err error) {
	token, err := extractToken(jwtTokenStr, jwtkey)
	if err != nil{
		return false, err
	}
	if !token.Valid{
		err = fmt.Errorf("invalid Jwt-token")
		return false, err
	}
	return true, err
}

// //ReNewToken reinitialise the Expiration time to the token actually it regenrate a fresh jwt.Token
// func ReNewToken(tokenStr string, jwtkey string) (jwtTokenstr string, err error) {
// 	claims := &jwtclaims{}
// 	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
// 		return
// 	}
// 	expirationTime := time.Now().Add(5 * time.Minute)
// 	claims.ExpiresAt = expirationTime.Unix()
// 	claims.userInformation, err = ValidateAndExtractUserID(tokenStr, jwtkey)
// 	if err != nil {
// 		return "", err
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	jwtTokenStr, err := token.SignedString(jwtkey)
// 	if err != nil {
// 		return "", err
// 	}
// 	return jwtTokenStr, nil
// }

//ExtractTokenUserInfo fetchs user information from jwt Token
func ValidateAndExtractUserID(tokenString string, jwtkey string) (userId string, err error) {
	token, err := extractToken(tokenString, jwtkey)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, ok := claims["user_id"].(string)
		if !ok {
			return "", fmt.Errorf("not access user-id")
		}
		return userId, nil
	}
	return "", fmt.Errorf("error when parsing claims")
}

func extractToken(tokenString, jwtAcckey string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method not matched")
		}
		return []byte(jwtAcckey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
