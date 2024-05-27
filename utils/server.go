package utils

import (
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type Server struct {
	Os   Os   `json:"os"`
	Cpu  Cpu  `json:"cpu"`
	Ram  Ram  `json:"ram"`
	Disk Disk `json:"disk"`
}

type Os struct {
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"numCpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"goVersion"`
	NumGoroutine int    `json:"numGoroutine"`
}

type Cpu struct {
	Cpus  []float64 `json:"cpus"`
	Cores int       `json:"cores"`
}

type Ram struct {
	UsedMB      int `json:"usedMb"`
	TotalMB     int `json:"totalMb"`
	UsedPercent int `json:"usedPercent"`
}

type Disk struct {
	UsedMB      int `json:"usedMb"`
	UsedGB      int `json:"usedGb"`
	TotalMB     int `json:"totalMb"`
	TotalGB     int `json:"totalGb"`
	UsedPercent int `json:"usedPercent"`
}

//@function: InitCPU
//@description: OS信息
//@return: o Os, err error

func InitOS() (o Os) {
	o.GOOS = runtime.GOOS
	o.NumCPU = runtime.NumCPU()
	o.Compiler = runtime.Compiler
	o.GoVersion = runtime.Version()
	o.NumGoroutine = runtime.NumGoroutine()
	return o
}

//@function: InitCPU
//@description: CPU信息
//@return: c Cpu, err error

func InitCPU() (c Cpu, err error) {
	if cores, err := cpu.Counts(false); err != nil {
		return c, err
	} else {
		c.Cores = cores
	}
	if cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return c, err
	} else {
		c.Cpus = cpus
	}
	return c, nil
}

//@function: InitRAM
//@description: RAM信息
//@return: r Ram, err error

func InitRAM() (r Ram, err error) {
	if u, err := mem.VirtualMemory(); err != nil {
		return r, err
	} else {
		r.UsedMB = int(u.Used) / MB
		r.TotalMB = int(u.Total) / MB
		r.UsedPercent = int(u.UsedPercent)
	}
	return r, nil
}

//@function: InitDisk
//@description: 硬盘信息
//@return: d Disk, err error

func InitDisk() (d Disk, err error) {
	percent := 0
	if u, err := disk.Usage("/"); err != nil {
		return d, err
	} else {
		d.UsedMB = int(u.Used) / MB
		d.UsedGB = int(u.Used) / GB
		d.TotalMB = int(u.Total) / MB
		d.TotalGB = int(u.Total) / GB
		percent = percent + int(u.UsedPercent)
	}
	// if u, err := disk.Usage("/home"); err != nil {
	// 	return d, err
	// } else {
	// 	d.UsedMB = d.UsedMB + int(u.Used)/MB
	// 	d.UsedGB = d.UsedGB + int(u.Used)/GB
	// 	d.TotalMB = d.TotalMB + int(u.Total)/MB
	// 	d.TotalGB = d.TotalGB + int(u.Total)/GB
	// 	percent = percent + int(u.UsedPercent)
	// }
	// percent = percent / 2
	d.UsedPercent = percent
	return d, nil
}

// cpu：CPU 相关；
// disk：磁盘相关；
// docker：docker 相关；
// host：主机相关；
// mem：内存相关；
// net：网络相关；
// process：进程相关；
// winservices：Windows 服务相关。

// 判断文件是否存在
func IsExistFile(pathFile string) (bool, error) {
	_, err := os.Stat(pathFile)
	return err == nil, err
}
