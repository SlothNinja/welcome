package welcome

import (
	"cloud.google.com/go/datastore"
	"github.com/SlothNinja/log"
	"github.com/SlothNinja/sn"
	"github.com/SlothNinja/user"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

const (
	msgEnter = "Entering"
	msgExit  = "Exiting"
)

type Client struct {
	*sn.Client
	User *user.Client
}

func NewClient(dClient *datastore.Client, uClient *user.Client, logger *log.Logger, cache *cache.Cache, router *gin.Engine) *Client {
	logger.Debugf(msgEnter)
	defer logger.Debugf(msgEnter)

	client := Client{
		Client: sn.NewClient(dClient, logger, cache, router),
		User:   uClient,
	}
	return client.addRoutes()
}

func (client *Client) addRoutes() *Client {
	log.Debugf(msgEnter)
	defer log.Debugf(msgExit)

	client.Router.GET("/", client.Index)
	return client
}
