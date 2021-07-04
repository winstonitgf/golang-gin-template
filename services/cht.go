package services

import (
	"fmt"
	"golang-startup/global"
	"golang-startup/models"
	"math/rand"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

func RedisGetSET(id string) (string, error) {
	conn := global.Redis.Get()
	defer conn.Close()

	// 先從 redis 撈取資料
	if data, err := redis.String(conn.Do("GET", fmt.Sprintf("blog:%s", id))); err != nil {

		// 無法從 redis 撈取，改為 MySQL 撈取資料
		time.Sleep(2 * time.Second)
		blogLikeTimes := "4433"

		// 再次放入 redis
		if _, err := conn.Do("SET", fmt.Sprintf("blog:%s", id), blogLikeTimes); err != nil {
			fmt.Println(err.Error())
		}

		return blogLikeTimes, nil
	} else {
		return data, nil
	}
}

func RedisHMGetSET(id string) (*models.User, error) {
	conn := global.Redis.Get()
	defer conn.Close()

	userEntity := new(models.User)

	// 先從 redis 撈取資料
	if values, err := redis.Values(conn.Do("HGETALL", fmt.Sprintf("user:%s", id))); err != nil || len(values) == 0 {

		// 無法從 redis 撈取，改為 MySQL 撈取資料
		time.Sleep(2 * time.Second)
		if err := global.Mysql.First(userEntity, id).Error; err != nil {
			return nil, err
		}

		// 再次放入 redis
		if _, err := conn.Do("HMSET", fmt.Sprintf("user:%s", id),
			"id", userEntity.ID,
			"account", userEntity.Account,
			"nickname", userEntity.Nickname,
			"email", userEntity.Email,
			"phone", userEntity.Phone,
			"created_at", userEntity.CreatedAt,
			"updated_at", userEntity.UpdatedAt); err != nil {
			fmt.Println(err.Error())
		}

		return userEntity, nil
	} else {
		// fmt.Printf("%+v", userEntity)
		err = redis.ScanStruct(values, userEntity)
		return userEntity, nil
	}
}

func RedisListGetSet() ([]string, error) {

	conn := global.Redis.Get()
	defer conn.Close()

	usernameKey := "hr:jobs"
	var usernames []string

	// 先從 redis 撈取資料
	if values, err := redis.Values(conn.Do("LRANGE", usernameKey, "0", "-1")); err != nil || len(values) == 0 {

		// 無法從 redis 撈取，改為 MySQL 撈取資料
		time.Sleep(2 * time.Second)
		if err := global.Mysql.Model(models.User{}).Select("nickname").Find(&usernames).Error; err != nil {
			return nil, err
		}

		// 再次放入 redis
		if _, err := conn.Do("LPUSH", redis.Args{}.Add(usernameKey).AddFlat(usernames)...); err != nil {
			fmt.Println(err.Error())
		}

		return usernames, nil
	} else {

		redis.ScanSlice(values, &usernames)
		return usernames, nil
	}
}

func RedisSetGetSet() ([]int, error) {

	conn := global.Redis.Get()
	defer conn.Close()

	var sets = []string{"tag:1000:news", "tag:1012:news", "tag:1044:news"}

	// 先從 redis 撈取資料
	if values, err := redis.Values(conn.Do("SINTER", redis.Args{}.AddFlat(sets)...)); err != nil || len(values) == 0 {

		// 無法從 redis 撈取，改為 MySQL 撈取資料
		time.Sleep(2 * time.Second)

		// 假設把資料庫的資料，放進 redis 的 SET 中
		for _, set := range sets {

			// 隨機數字，模擬不同的 news ID
			var randomIDs []int
			for i := 0; i < 10; i++ {
				randomIDs = append(randomIDs, rand.Intn(10))
			}

			// 再次放入 redis
			if _, err := conn.Do("SADD", redis.Args{}.Add(set).AddFlat(randomIDs)...); err != nil {
				fmt.Println(err.Error())
			}
		}

		// 做一些亂數，模擬從資料庫回傳
		var randomIDs []int
		for i := 0; i < 10; i++ {
			randomIDs = append(randomIDs, rand.Intn(10))
		}

		return randomIDs, nil
	} else {

		var sinterResult []int
		redis.ScanSlice(values, &sinterResult)
		return sinterResult, nil
	}
}

func RedisSortedSetGetSet() ([]models.BasketballTeam, error) {

	conn := global.Redis.Get()
	defer conn.Close()

	sortedSetKey := "basketball:score"

	// 先從 redis 撈取資料
	if values, err := redis.Values(conn.Do("ZREVRANGE", sortedSetKey, "0", "-1", "withscores")); err != nil || len(values) == 0 {

		// 無法從 redis 撈取，改為 MySQL 撈取資料
		time.Sleep(2 * time.Second)

		// 假設把資料庫的資料，放進 redis 的 SET 中
		var basketballTeamEntities []models.BasketballTeam
		var basketballTeams = []string{"TeamA", "TeamB", "TeamC", "TeamD", "TeamE", "TeamF"}
		for _, basketballTeamName := range basketballTeams {

			basketballTeamEntity := new(models.BasketballTeam)
			basketballTeamEntity.Name = basketballTeamName
			basketballTeamEntity.Score = rand.Intn(100)
			basketballTeamEntities = append(basketballTeamEntities, *basketballTeamEntity)
			// 再次放入 redis
			if _, err := conn.Do("ZADD", sortedSetKey, basketballTeamEntity.Score, basketballTeamEntity.Name); err != nil {
				fmt.Println(err.Error())
			}
		}

		// 假設資料庫處理好相關邏輯，並回傳
		return basketballTeamEntities, nil
	} else {

		var basketballTeamEntities []models.BasketballTeam

		for i := 0; i+1 < len(values); i = i + 2 {
			name, ok := values[i].([]byte)
			if !ok {
				return nil, fmt.Errorf("expected []byte value but was %T", values[i])
			}

			score, ok := values[i+1].([]byte)
			if !ok {
				return nil, fmt.Errorf("expected []byte value but was %T", values[i])
			}

			basketballTeamEntity := new(models.BasketballTeam)
			basketballTeamEntity.Name = string(name)
			basketballTeamEntity.Score, err = strconv.Atoi(string(score))
			basketballTeamEntities = append(basketballTeamEntities, *basketballTeamEntity)
		}
		return basketballTeamEntities, nil
	}
}
