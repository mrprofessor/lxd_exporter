package main

import (
	"log"

	lxd "github.com/lxc/lxd/client"
	lxd_api "github.com/lxc/lxd/shared/api"
	"github.com/prometheus/client_golang/prometheus"
)

//type collector struct {
//logger *log.Logger
//server lxd.InstanceServer
//}

//func NewCollector(
//logger *log.Logger, server lxd.InstanceServer) *collector {
//return &collector{logger: logger, server: server}
//}

var lxdContainerCount = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "lxd_container_count", Help: "Total number of lxd containers in a host",
})

func getLXDContainerCount(lxdServer lxd.InstanceServer) float64 {
	containersList, _ := lxdServer.GetInstances(lxd_api.InstanceTypeContainer)
	return float64(len(containersList))
}

func init() {

	server, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Fatalf("Unable to contact LXD server: %s", err)
		return
	}

	containersList, _ := server.GetInstances(lxd_api.InstanceTypeContainer)
	log.Println(containersList)

	prometheus.MustRegister(lxdContainerCount)
	lxdContainerCount.Set(getLXDContainerCount(server))
}
