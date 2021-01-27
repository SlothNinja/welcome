package welcome

import (
	"net/http"

	"github.com/SlothNinja/sn"
	"github.com/gin-gonic/gin"
)

func (client Client) Index(c *gin.Context) {
	client.Log.Debugf("Entering")
	defer client.Log.Debugf("Exiting")

	cu, err := client.User.Current(c)
	if err != nil {
		client.Log.Debugf(err.Error())
	}
	c.HTML(http.StatusOK, "welcome/index", gin.H{
		"VersionID": sn.VersionID(),
		"CUser":     cu,
		"Context":   c})
}
