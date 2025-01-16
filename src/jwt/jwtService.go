package jwt

import (
	"awesomeProject/global"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	jwt2 "github.com/golang-jwt/jwt/v5"

	"time"
)


func CreateToken(exp int64)(string, error)  {
	claims := jwt2.RegisteredClaims{
		Subject: "nguyenvu",
		ID: "1",
		ExpiresAt: jwt2.NewNumericDate(time.Now().Add(time.Duration(exp)* time.Millisecond)),
		IssuedAt: jwt2.NewNumericDate(time.Now()),
		Issuer:	"AwesomeProject",
	}
	privateKey, err := getPrivateKey()
	token := jwt2.NewWithClaims(jwt2.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("error signing the token: %v", err)
	}
	return signedToken, err

}

func ExtractClaim(token string) (*jwt2.RegisteredClaims,error) {
	claims := &jwt2.RegisteredClaims{}

	t,err := jwt2.ParseWithClaims(token,claims, func(token *jwt2.Token) (interface{}, error) {
		return getPublicKey()
	})
	if err != nil {
		return nil, fmt.Errorf("could not parse token: %v", err)
	}
	if claims, ok := t.Claims.(*jwt2.RegisteredClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")


}

func getPrivateKey()(*rsa.PrivateKey,error)  {
	key := global.Config.JwtData
	keyBytes,err := base64.StdEncoding.DecodeString(key.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 key: %v", err)
	}
	privateKey,err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 key: %v", err)
	}
	if rsaKey, ok := privateKey.(*rsa.PrivateKey); ok {
		return rsaKey, nil
	}
	return nil, errors.New("private key is not RSA")

}

func getPublicKey()(*rsa.PublicKey,error)  {
	key := global.Config.JwtData
	keyBytes,err := base64.StdEncoding.DecodeString(key.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 key: %v", err)
	}
	publickey,err := x509.ParsePKIXPublicKey(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 key: %v", err)
	}
	if pubLey,ok := publickey.(*rsa.PublicKey); ok {
		return pubLey, nil
	}
	return nil, errors.New("private key is not RSA")
}
func VerifyTokenSubject(token string)(*jwt2.RegisteredClaims,error){
	
	claims, err := ExtractClaim(token)
	if err != nil {
		return nil, err
	}
	
	return claims, nil
}

