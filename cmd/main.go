package main

import (
	"github.com/golang/glog"
	"github.com/madhukirans/replayed/pkg/server"
	"github.com/madhukirans/replayed/pkg/types"
	"strconv"
	_ "net/http/pprof"
	"net/http"
)

var config *types.ReplayedConfig

func init() {
	//log.SetOutput(new(types.LogWriter))
	config = types.GetReplayedConfig()
	glog.Info("Config: v", config)
}

func main() {
	server.InitServer(config)
	http.HandleFunc("/", server.Handler)
	glog.Info("Starting server")
	http.ListenAndServe(":" + strconv.Itoa(config.Port), nil)
}
