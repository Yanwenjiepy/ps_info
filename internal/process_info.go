package internal

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"ps_info/pkg/logger"
)

type PlatformContent struct {
	Platform string
	Family   string
	Version  string
}

type CPUContent struct {
	PhysicalNum int
	LogicalNum  int
}

type MemoryContent struct {
	TotalSize float64
	UsedSize  float64
}

// PlatformInfo return the name、family、version of the platform currently in use
func PlatformInfo() PlatformContent {

	platformInfoContent := PlatformContent{}

	platformName, platformFamily, platformVersion, err := host.PlatformInformation()
	if err != nil {
		logger.Log.Error("[PlatformInfo] Failed to get platform info, err: " + err.Error())
		return platformInfoContent
	}

	platformInfoContent.Platform = platformName
	platformInfoContent.Family = platformFamily
	platformInfoContent.Version = platformVersion

	return platformInfoContent
}

// CPUInfo return the cpu physical of logical core of the computer currently in use
func CPUInfo() CPUContent {

	cpuInfoContent := CPUContent{}

	cpuPhysicalNum, err := cpu.Counts(false)
	if err != nil {
		logger.Log.Error("[CPUInfo] Failed to get cpu physical number, err: " + err.Error())
		return cpuInfoContent
	}

	cpuLogicalNum, err := cpu.Counts(true)
	if err != nil {
		logger.Log.Error("[CPUInfo] Failed to get cpu logical number, err: " + err.Error())
		return cpuInfoContent
	}

	cpuInfoContent.PhysicalNum = cpuPhysicalNum
	cpuInfoContent.LogicalNum = cpuLogicalNum

	return cpuInfoContent
}

// CPUUseInfo return the CPU usage of the computer currently in use
func CPUUsage() []float64 {
	cpuUsePercentList, err := cpu.Percent(0, true)
	if err != nil {
		logger.Log.Error("[CPUUsage] Failed to get cpu percent, err: " + err.Error())
		return nil
	}
	return cpuUsePercentList
}

// MemoryInfo return the memory usage of the computer currently in use
func MemoryInfo() MemoryContent {

	memoryInfoContent := MemoryContent{}

	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		logger.Log.Error("[MemoryInfo] Failed to get virtual memory, err: " + err.Error())
		return memoryInfoContent
	}

	totalMemory := float64(virtualMemory.Total) / (1024 * 1024 * 1024)
	usedMemory := float64(virtualMemory.Used) / (1024 * 1024 * 1024)

	memoryInfoContent.TotalSize = totalMemory
	memoryInfoContent.UsedSize = usedMemory

	return memoryInfoContent
}
