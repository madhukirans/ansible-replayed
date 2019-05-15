package server

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/madhukirans/replayed/pkg/types"
	"io/ioutil"

	"github.com/golang/glog"
	"sync"
	"net/http"
	"strconv"
	"fmt"
)

var config *types.ReplayedConfig
var serverBuf *bytes.Buffer
var clientRequestBufferSize int
var mutex sync.RWMutex
var serverBufferSize int

func InitServer(c *types.ReplayedConfig) {
	config = c
	serverBufferSize = c.BufferSizeInMB * 1024 * 1024
	clientRequestBufferSize = c.ClientRequestBufferSizeInKB * 1024
	serverBuf = new(bytes.Buffer)
}

var i int

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		mutex.RLock()
		w.Write([]byte(strconv.Itoa(i)))
		w.Write(serverBuf.Bytes())
		w.WriteHeader(http.StatusOK)
		mutex.RUnlock()
	case "POST":
		x, err := ioutil.ReadAll(r.Body)
		if err != nil {
			glog.Errorf("Reading error %v", err)
		} else {
			mutex.Lock()
			if serverBuf.Len()+len(x) < serverBufferSize {
				serverBuf.Write(x)
				w.WriteHeader(http.StatusAccepted)
			} else {
				w.WriteHeader(http.StatusInsufficientStorage)
			}
			i++
			mutex.Unlock()
		}
	case "DELETE":

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
		fmt.Println("method error")
	}
}

func PostData(c *gin.Context) {
	//glog.Info("Reading client \n")
	x, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		glog.Errorf("Reading error %v", err)
	} else {
		mutex.Lock()
		if serverBuf.Len()+len(x) < serverBufferSize {
			serverBuf.Write(x)
			c.JSON(http.StatusAccepted, gin.H{"status": "Ok"})
		} else {
			c.JSON(http.StatusInsufficientStorage, gin.H{"status": "StatusInsufficientStorage"})
		}
		mutex.Unlock()
	}
}

func GetData(c *gin.Context) {
	mutex.RLock()
	c.String(200, serverBuf.String())
	mutex.RUnlock()
}
