package token

import (
	"bao-bet365-api/model/env"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenService struct{}

var jwtSecret = []byte(env.Enviroment.Jwt.Secret)

func (this TokenService) GenerateToken(userId int) (token string, err error) {

	// 建立過期時間
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(env.Enviroment.Jwt.Expired) * time.Minute).Unix()

	// 找用戶
	// var user member.UcMember
	// if err = database.Eloquent.First(&user, userId).Error; err != nil {
	// 	return
	// }
	// token的宣告
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "codes.atavral.ninja/bao/sport/bao-cms-api",
		// "userId":  user.Id,
		// "uid":     user.Uid,
		// "account": user.Account,
		"exp": expireTime,
	})
	// 產生jwt token
	token, err = tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return
	}

	return
}

func (this TokenService) GenerateRefreshToken(userId int) (refreshToken string, err error) {

	// 建立過期時間
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(env.Enviroment.Jwt.Refresh) * time.Minute).Unix()

	// 找用戶
	// var user member.UcMember
	// if err = database.Eloquent.First(&user, userId).Error; err != nil {
	// 	return
	// }

	// token的宣告
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "codes.atavral.ninja/bao/sport/bao-cms-api",
		// "userId":  user.Id,
		// "uid":     user.Uid,
		// "account": user.Account,
		"exp": expireTime,
	})

	// 產生jwt token
	refreshToken, err = tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return
	}

	return
}

func (this TokenService) ValidateToken(uncheckToken string) (err error) {
	// 驗證token
	_, err = jwt.Parse(uncheckToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	// 錯誤回傳
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func (this TokenService) ParseToken(authToken string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {

		fmt.Println(err.Error())
		return nil, err
	}
}
