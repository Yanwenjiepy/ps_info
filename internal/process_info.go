package internal

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"ps_info/logger"
)

type platformInfo struct {
	Platform string
	Family   string
	Version  string
}

type cpuInfo struct {
	PhysicalNum int
	LogicalNum  int
}

type memoryInfo struct {
	TotalSize float64
	UsedSize  float64
}

// PlatformInfo return the name、family、version of the platform currently in use
func PlatformInfo() *platformInfo {

	platformInfoContent := platformInfo{}

	platformName, platformFamily, platformVersion, err := host.PlatformInformation()
	if err != nil {
		logger.Log.Error("Failed to get platform info!")
		return &platformInfoContent
	}

	platformInfoContent.Platform = platformName
	platformInfoContent.Family = platformFamily
	platformInfoContent.Version = platformVersion

	return &platformInfoContent
}

// CPUInfo return the cpu physical of logical core of the computer currently in use
func CPUInfo() *cpuInfo {

	cpuInfoContent := cpuInfo{}

	cpuPhysicalNum, err := cpu.Counts(false)
	if err != nil {
		logger.Log.Error("Failed to get cpu physical number!")
		return &cpuInfoContent
	}

	cpuLogicalNum, err := cpu.Counts(true)
	if err != nil {
		logger.Log.Error("Failed to get cpu logical number!")
		return &cpuInfoContent
	}

	cpuInfoContent.PhysicalNum = cpuPhysicalNum
	cpuInfoContent.LogicalNum = cpuLogicalNum

	return &cpuInfoContent
}

// CPUUseInfo return the CPU usage of the computer currently in use
func CPUUsage() []float64 {
	cpuUsePercentList, err := cpu.Percent(0, true)
	if err != nil {
		logger.Log.Error("Failed to get cpu percent!")
		return nil
	}
	return cpuUsePercentList
}

// MemoryInfo return the memory usage of the computer currently in use
func MemoryInfo() *memoryInfo {

	memoryInfoContent := memoryInfo{}

	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		logger.Log.Error("Failed to get virtual memory!")
		return &memoryInfoContent
	}

	totalMemory := float64(virtualMemory.Total) / (1024 * 1024 * 1024)
	usedMemory := float64(virtualMemory.Used) / (1024 * 1024 * 1024)

	memoryInfoContent.TotalSize = totalMemory
	memoryInfoContent.UsedSize = usedMemory

	return &memoryInfoContent
}
