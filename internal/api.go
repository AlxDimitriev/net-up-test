package internal

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type IUsersAPI interface {
	RegisterUrls()
	Run()
}

type GinUsersAPI struct {
	router *gin.Engine
	activeUsers map[string]time.Time
}

func NewGinUsersAPI() *GinUsersAPI {
	return &GinUsersAPI{router: gin.Default(), activeUsers: make(map[string]time.Time)}
}

func (a *GinUsersAPI) RegisterUrls()  {
	a.router.GET("user/ping", userPingHandler)
	a.router.GET("admin/users", getUsersHandler)
}

func (a *GinUsersAPI) Run() {
	if err := a.router.Run(); err != nil {
		log.Panic(err)
	}
}