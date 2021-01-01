package welcome

import (
	"github.com/SlothNinja/user"
	"github.com/gin-gonic/gin"
)

type Client struct {
	User user.Client
}

func NewClient(userClient user.Client) Client {
	return Client{
		User: userClient,
	}
}

func (client Client) AddRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/", client.Index)
	return r
}
