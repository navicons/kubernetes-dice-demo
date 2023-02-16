package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	r := gin.Default()
	diceMax, err := strconv.Atoi(os.Getenv("APP_DICE_MAX"))
	if err != nil || diceMax > 6 || diceMax < 1 {
		diceMax = 6
	}

	btnTitle := os.Getenv("APP_BTN_ROLL_TITLE")
	if btnTitle == "" {
		btnTitle = "Roll"
	}

	r.GET("/backend/health", CORSMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	r.GET("/backend/settings", CORSMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"max":      diceMax,
			"btnTitle": btnTitle,
		})
	})

	r.GET("/backend/random", CORSMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"num":     rnd(diceMax),
			"victory": rand.Float32() < 0.5,
		})
	})
	r.Run()
}

func rnd(max int) int {
	r := rand.Intn(max)
	if r == 0 {
		r = 1
	}
	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
