package welcome

import (
	"net/http"

	"github.com/SlothNinja/log"
	"github.com/SlothNinja/sn/v2"
	"github.com/SlothNinja/user/v2"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	log.Debugf("Entering")
	defer log.Debugf("Exiting")

	cu, err := user.FromSession(c)
	log.Debugf("cu: %#v\nerr: %v", cu, err)
	c.HTML(http.StatusOK, "welcome/index", gin.H{
		"VersionID": sn.VersionID(),
		"CUser":     cu,
		"Context":   c})
}
