package main

import (
	"log"
	"os"
	"runtime"
	lxd "github.com/lxc/lxd/client"
	lxd_api "github.com/lxc/lxd/shared/api"
	"github.com/prometheus/client_golang/prometheus"
)

var hostConstantLabels = prometheus.Labels{
	"host_name": getHostName(),
}

var (
	hostLXDContainerCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "host_lxd_container_count",
		Help:        "Count of lxd containers in a host",
		ConstLabels: hostConstantLabels,
	})
	hostCPULogicalCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "host_cpu_logical_count",
		Help:        "Count of logical CPUs in a host",
		ConstLabels: hostConstantLabels,
	})
	hostCPUCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "host_cpu_count",
		Help:        "Count of physical CPUs in a host",
		ConstLabels: hostConstantLabels,
	})
	hostCPUCoreCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "host_cpu_core_count",
		Help:        "Count of CPU cores in a host",
		ConstLabels: hostConstantLabels,
	})
	HostDiskGb = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "host_disk_gb",
		Help:        "Host filesystem disk size in GB",
		ConstLabels: hostConstantLabels,
	})
	HostDiskFreeGb = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "host_disk_free_gb",
		Help:        "Free disk size in GB",
		ConstLabels: hostConstantLabels,
	})
	HostDiskAvailableGb = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "host_disk_available_gb",
		Help:        "Available disk size in GB",
		ConstLabels: hostConstantLabels,
	})
	hostMemoryMb = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "host_memory_mb",
		Help:        "Host total memory in MB",
		ConstLabels: hostConstantLabels,
	})
	hostMemoryFreeMb = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "host_memory_mb",
		Help:        "Host free memory size in MB",
		ConstLabels: hostConstantLabels,
	})
	hostMemoryAvailableMb = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "host_memory_mb",
		Help:        "Host available memory size in MB",
		ConstLabels: hostConstantLabels,
	})
)

func getHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Unable to find hostname: %s", err)
	}
	return hostname
}

func getLXDContainerCount() float64 {

	lxdServer, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Printf("Unable to contact LXD server: %s", err)
		return float64(-1)
	}
	containersList, _ := lxdServer.GetInstances(lxd_api.InstanceTypeContainer)
	return float64(len(containersList))
}

func getHostCPULogicalCount() float64 {
	return float64(runtime.NumCPU())
}

func Collector(logger *log.Logger) {
	prometheus.MustRegister(hostCPULogicalCount)
	logger.Println("Registerd host_cpu_logical_count metric")

	prometheus.MustRegister(hostLXDContainerCount)
	logger.Println("Registerd host_lxd_container_count metric")

	hostCPULogicalCount.Set(getHostCPULogicalCount())
	hostLXDContainerCount.Set(getLXDContainerCount())
}
