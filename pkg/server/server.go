package server

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/madhukirans/replayed/pkg/types"
	"io/ioutil"

	"github.com/golang/glog"
)

var config *types.ReplayedConfig
var serverBuf *bytes.Buffer

func StartServer(c *types.ReplayedConfig) *gin.Engine {
	config = c
	serverBuf = new(bytes.Buffer)

	r := gin.Default()
	//r.Use(s.Cors())

	v1 := r.Group("/")
	{
		v1.GET("/", GetData)
		v1.POST("/", PostData)
	}

	return r
}

func PostData(c *gin.Context) {
	//glog.Info("Reading client \n")
	x, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		glog.Error("Reading error %v", err)
	} else {
		serverBuf.Write(x)
	}
}

func GetData(c *gin.Context) {
	c.String(200, serverBuf.String())
}
