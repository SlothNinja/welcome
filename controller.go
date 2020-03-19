package welcome

import (
	"net/http"

	"github.com/SlothNinja/log"
	"github.com/SlothNinja/sn"
	"github.com/SlothNinja/user"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	log.Debugf("Entering")
	defer log.Debugf("Exiting")

	cu := user.CurrentFrom(c)
	log.Debugf("cu: %#v", cu)
	c.HTML(http.StatusOK, "welcome/index", gin.H{
		"VersionID": sn.VersionID(),
		"CUser":     cu,
		"Context":   c})
}
