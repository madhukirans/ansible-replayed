package server

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/madhukirans/replayed/pkg/types"
	"io/ioutil"

	"github.com/golang/glog"
	"sync"
)

var config *types.ReplayedConfig
var serverBuf *bytes.Buffer
var mutex sync.RWMutex

func StartServer(c *types.ReplayedConfig) *gin.Engine {
	config = c
	serverBuf = new(bytes.Buffer)
	r := gin.Default()
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
		glog.Errorf("Reading error %v", err)
	} else {
		mutex.Lock()
		serverBuf.Write(x)
		mutex.Unlock()
	}
}

func GetData(c *gin.Context) {
	mutex.RLock()
	c.String(200, serverBuf.String())
	mutex.RUnlock()
}
