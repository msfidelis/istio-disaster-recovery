package system

import (
	"runtime"

	"github.com/zcalusic/sysinfo"
)

type Capabilities struct {
	Hostname   string `json:"hostname" binding:"required"`
	Cpus       int    `json:"cpus" binding:"required"`
	Os         string `json:"os" binding:"required"`
	Hypervisor string `json:"hypervisor" binding:"required"`
	Ram        uint   `json:"memory" binding:"required"`
}

func Info() Capabilities {

	var capabilities Capabilities
	var si sysinfo.SysInfo
	si.GetSysInfo()

	capabilities.Hostname = si.Node.Hostname
	capabilities.Cpus = runtime.NumCPU()
	capabilities.Os = si.OS.Name
	capabilities.Hypervisor = si.Node.Hypervisor
	capabilities.Ram = si.Memory.Size

	return capabilities

} 