package controller

import (
	"golang-startup/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary RedisGetSET
// @Tags Redis
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Security ApiKeyAuth
// @Param id path string true "ID"
// @Router /api/v1/redis/getset/{id} [post]
func RedisGetSET(c *gin.Context) {

	if data, err := services.RedisGetSET(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

// @Summary RedisHMGetSET
// @Tags Redis
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.User
// @Failure 400 {string} string
// @Security ApiKeyAuth
// @Router /api/v1/redis/hmgetset/{id} [post]
func RedisHMGetSET(c *gin.Context) {

	if data, err := services.RedisHMGetSET(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

// @Summary RedisListGetSet
// @Tags Redis
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Failure 400 {string} string
// @Security ApiKeyAuth
// @Router /api/v1/redis/listgetset [post]
func RedisListGetSet(c *gin.Context) {

	if data, err := services.RedisListGetSet(); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

// @Summary RedisSetGetSet
// @Tags Redis
// @Accept  json
// @Produce  json
// @Success 200 {array} integer
// @Failure 400 {string} string
// @Security ApiKeyAuth
// @Router /api/v1/redis/setgetset [post]
func RedisSetGetSet(c *gin.Context) {

	if data, err := services.RedisSetGetSet(); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

// @Summary RedisSortedsetGetSet
// @Tags Redis
// @Accept  json
// @Produce  json
// @Success 200 {array} models.BasketballTeam
// @Failure 400 {string} string
// @Security ApiKeyAuth
// @Router /api/v1/redis/sortedsetgetset [post]
func RedisSortedSetGetSet(c *gin.Context) {

	if data, err := services.RedisSortedSetGetSet(); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, data)
	}
}
