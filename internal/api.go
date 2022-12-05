package internal

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"time"
)

type activeUser struct {
	IpAddr string `json:"ip_address"`
	Since  int64  `json:"since"`
}

type IUsersAPI interface {
	RegisterUrls()
	Run()
}

type GinUsersAPI struct {
	router       *gin.Engine
	activeUsers  map[string]time.Time
	mu           *sync.RWMutex
}

func NewGinUsersAPI() *GinUsersAPI {
	return &GinUsersAPI{
		router: gin.Default(),
		activeUsers: make(map[string]time.Time),
		mu: &sync.RWMutex{},
	}
}

func (a *GinUsersAPI) RegisterUrls()  {
	a.router.GET("user/ping", a.userPingHandler)
	a.router.GET("admin/users", a.getUsersHandler)
}

func (a *GinUsersAPI) Run() {
	if err := a.router.Run(); err != nil {
		log.Panic(err)
	}
}

func (a *GinUsersAPI) userPingHandler(c *gin.Context) {
	a.mu.Lock()
	a.activeUsers[c.Request.RemoteAddr] = time.Now()
	a.mu.Unlock()
	c.JSON(http.StatusOK, gin.H{})
}

func (a *GinUsersAPI) getUsersHandler(c *gin.Context) {
	now := time.Now()
	var activeUsers []activeUser
	a.mu.RLock()
	for ip, since := range a.activeUsers {
		diff := now.Unix() - since.Unix()
		if diff < 30*60*60 {
			activeUsers = append(activeUsers, activeUser{ip, since.Unix()})
		}
	}
	a.mu.RUnlock()
	c.JSON(http.StatusOK, gin.H{"active_users": activeUsers})
}