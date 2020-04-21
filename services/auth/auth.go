package auth

import (
	"99live-cms-golang-api/env"
	"99live-cms-golang-api/models"
	"99live-cms-golang-api/structs"
	"fmt"

	"99live-cms-golang-api/packages"
	"99live-cms-golang-api/plugins/instances"
	"99live-cms-golang-api/services"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
)

type AuthService struct{}

var (
	errorList []error
)

func (AuthService) Login(userInfo models.User, ip string) (loginModel structs.LoginModel) {

	var err error

	// 取得用戶資料
	if loginModel.User, err = GetUserEntity(userInfo); err != nil {
		handleError(err)
		return
	}

	// 更新用戶狀態
	var updateUserModel models.User
	updateUserModel.Id = loginModel.User.Id
	updateUserModel.Status = 200
	updateUserModel.UpdatedAt = uint(time.Now().Unix())
	if err = instances.Mysql.Model(&models.User{}).UpdateColumn(&updateUserModel).Error; err != nil {
		handleError(err)
		return
	}

	// 取得 token
	if loginModel.Token.Token, err = getToken(loginModel.User, 1, "24h"); err != nil {
		handleError(err)
		return
	}

	// 取得 refresh token
	if loginModel.Token.RefreshToken, err = getToken(loginModel.User, 1, "240h"); err != nil {
		handleError(err)
		return
	}

	// 新增登入紀錄
	if err = addLoginHistory(loginModel.User.Id, ip); err != nil {
		handleError(err)
		return
	}

	return
}

func (AuthService) Verify(tokenString string) (userEntity models.User) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(env.Config.Jwt.Secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		packages.Print("用戶 ID： ", claims["user_id"], "驗證成功")

		var userParam models.User
		userParam.Id = uint(claims["user_id"].(float64))

		if userEntity, err = GetUserEntity(userParam); err != nil {
			handleError(err)
			return
		}
	} else {

		handleError(err)
		packages.Print("請求驗證失敗")
		return
	}

	return
}

func addLoginHistory(userId uint, ip string) error {

	var (
		err error
	)

	// 新增登入紀錄
	var loginHistoryEntity models.LoginHistory
	loginHistoryEntity.UserId = userId
	loginHistoryEntity.Ip = ip
	loginHistoryEntity.CreatedAt = uint(time.Now().Unix())
	if err = instances.Mysql.Create(&loginHistoryEntity).Error; err != nil {
		return err
	}

	return nil
}

// 取得用戶資料 (check redis)
func GetUserEntity(userInfo models.User) (userEntity models.User, err error) {

	// 取得 redis key
	var redisService services.RedisService
	redisKey := redisService.GetUserKey(userInfo.Account)

	// 取一個 redis 連線出來
	var conn redis.Conn
	if conn = instances.RedisPool.Get(); conn.Err() != nil {

		packages.Print("GetUserEntity: redis pool 出錯，從資料庫抓取")
		// 如果 redis pool 取得失敗，就直接撈資料庫
		if err = instances.Mysql.Where(&userInfo).First(&userEntity).Error; err != nil {
			return
		}

		// 然後把撈出來的用戶資料，再次放到 redis cache (這裡放 cache 失敗了也沒差，下次在呼叫的時後，也就會從資料庫抓)
		conn.Do("HMSET", redis.Args{redisKey}.AddFlat(userEntity)...)
	}
	defer conn.Close()

	// redis 中取出整個用戶資料
	redisValue, err := redis.Values(conn.Do("HGETALL", redisKey))
	if err != nil || len(redisValue) == 0 {

		packages.Print("GetUserEntity: REDIS 發生錯誤，從資料庫重新抓取")
		// 如果從 redis 取得失敗，就直接撈資料庫
		if err = instances.Mysql.Where(&userInfo).First(&userEntity).Error; err != nil {
			return
		}

		// 然後把撈出來的用戶資料，再次放到 redis cache (這裡放 cache 失敗了也沒差，下次在呼叫的時後，也就會從資料庫抓)
		conn.Do("HMSET", redis.Args{redisKey}.AddFlat(userEntity)...)

		return
	}

	// 把用戶資料放到 struct
	err = redis.ScanStruct(redisValue, &userEntity)
	if err != nil {
		return
	}
	packages.Print("GetUserEntity: redis 中取出整個用戶資料")

	return
}

// 取得 token
func getToken(userEntity models.User, roleId int, validTime string) (string, error) {

	// token 時效
	tokenValidTime, _ := time.ParseDuration(validTime)

	// jwt 密鑰
	jwtSecretKey := []byte(env.Config.Jwt.Secret)

	// 用時間的 UNIX 當成 JWT ID
	now := time.Now()
	jwtId := strconv.FormatInt(now.Unix(), 10)

	var jwtClaims structs.JwtClaims
	jwtClaims.UserId = int(userEntity.Id)
	jwtClaims.RoleId = roleId
	jwtClaims.StandardClaims = jwt.StandardClaims{
		Audience:  userEntity.Account,
		ExpiresAt: now.Add(tokenValidTime).Unix(),
		Id:        jwtId,
		IssuedAt:  now.Unix(),
		Issuer:    "99-live",
		Subject:   userEntity.Account,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	token, err := tokenClaims.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

// 錯誤處理
func handleError(errorPayload error) {
	errorList = append(errorList, errorPayload)
}

func (AuthService) Errors() []error {
	return errorList
}
