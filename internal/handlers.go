package internal

import "github.com/gin-gonic/gin"

func userPingHandler(c *gin.Context) {
	println(c.Request.RemoteAddr)
}

func getUsersHandler(c *gin.Context) {

}