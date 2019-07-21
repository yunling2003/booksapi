package auth

import (
	"fmt"
	"strconv"
	"time"
	"booksapi/src/config"
	"booksapi/src/util"
	"booksapi/src/models/orm"

	"github.com/dgrijalva/jwt-go"
)

type Payload struct {
	UserID              uint64 `json:"uid,omitempty"`
	WechatOpenID        string `json:"wid,omitempty"`
	EncryptedSessionKey string `json:"esk,omitempty"`
	Visitor             string `json:"visitor,omitempty"`
	jwt.StandardClaims
}

func CreateJwt(payload *Payload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	ss, err := token.SignedString([]byte(config.All["JwtSignKey"]))

	return ss, err
}

// ParseJwt will parse jwt token to payload
func ParseJwt(tokenString string) (bool, *Payload) {
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.All["JwtSignKey"]), nil
	})

	if err != nil {
		fmt.Println(err)
		return false, &Payload{}
	}

	payload, ok := token.Claims.(*Payload)
	if ok && token.Valid {
		return ok, payload
	}

	fmt.Printf("Result: %v, Valid: %v", ok, token.Valid)
	return false, payload
}

func (base *Payload) Gen() string {
	expireMinutes, _ := strconv.Atoi(config.All["token.expire.minutes"])
	base.ExpiresAt = time.Now().Add(time.Minute * time.Duration(expireMinutes)).Unix()
	tokenString, _ := CreateJwt(base)
	return tokenString
}

func Check(token string) (bool, *Payload) {
	return ParseJwt(token)
}

func GenVisitorJwt() string {	
	randomUser := util.MD5(fmt.Sprintf("%v", time.Now().UnixNano()))
	payload := Payload{
		Visitor: randomUser,
	}

	return payload.Gen()
}

func GenUserJwt(user *orm.User) string {
	payload := Payload{
		UserID:       user.ID,
		WechatOpenID: user.WechatOpenID,
		// EncryptedSessionKey: user.WechatSessionKey,
	}

	return payload.Gen()
}
