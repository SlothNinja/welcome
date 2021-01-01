package welcome

import (
	"net/http"

	"github.com/SlothNinja/log"
	"github.com/SlothNinja/sn"
	"github.com/gin-gonic/gin"
)

func (client Client) Index(c *gin.Context) {
	log.Debugf("Entering")
	defer log.Debugf("Exiting")

	cu, err := client.User.Current(c)
	if err != nil {
		log.Debugf(err.Error())
	}
	c.HTML(http.StatusOK, "welcome/index", gin.H{
		"VersionID": sn.VersionID(),
		"CUser":     cu,
		"Context":   c})
}
