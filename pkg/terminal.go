package pkg

import (
	"fmt"
	"time"

	"ps_info/internal"
)

func Output() {

	platformInfoRes := internal.PlatformInfo()
	platformName, platformFamily, platformVersion := platformInfoRes.Platform, platformInfoRes.Family, platformInfoRes.Version

	cpuInfoRes := internal.CPUInfo()
	cpuPhysicalNum, cpuLogicalNum := cpuInfoRes.PhysicalNum, cpuInfoRes.LogicalNum

	for {
		fmt.Println("=====    Basic Information    =====")
		fmt.Printf("[OS]  Name: %s    Family: %s    Version: %s\n", platformName, platformFamily, platformVersion)
		fmt.Printf("[CPU]  Physcial number: %d    Logscial number:  %d\n", cpuPhysicalNum, cpuLogicalNum)
		fmt.Println()

		cpuPercentList := internal.CPUUsage()
		fmt.Println(cpuPercentList)
		fmt.Printf("[CPU]  cpu1: %.1f  cpu2: %.1f  cpu3: %.1f  cpu4: %.1f\n", cpuPercentList[0], cpuPercentList[1], cpuPercentList[2], cpuPercentList[3])
		fmt.Printf("[CPU]  cpu5: %.1f  cpu6: %.1f  cpu7: %.1f  cpu8: %.1f\n", cpuPercentList[4], cpuPercentList[5], cpuPercentList[6], cpuPercentList[7])

		memoryInfoRes := internal.MemoryInfo()
		totalMemory, usedMemory := memoryInfoRes.TotalSize, memoryInfoRes.UsedSize
		fmt.Printf("[RAM]  Total size: %.2fG    Used size: %.2fG\n", totalMemory, usedMemory)
		fmt.Println()

		fmt.Println("=====    Process Information    =====")

		time.Sleep(3 * time.Second)
	}

}

// OutputOriginal use the uilive library to update related information in the same place
// but cpu use percent 24%-30%, too higher
// func OutputOriginal() {
//	fmt.Println("=====    Basic Information    =====")
//
//	platformInfoRes := internal.PlatformInfo()
//	platformName, platformFamily, platformVersion := platformInfoRes.Platform, platformInfoRes.Family, platformInfoRes.Version
//	fmt.Printf("[OS]  Name: %s    Family: %s    Version: %s\n", platformName, platformFamily, platformVersion)
//
//	cpuInfoRes := internal.CPUInfo()
//	cpuPhysicalNum, cpuLogicalNum := cpuInfoRes.PhysicalNum, cpuInfoRes.LogicalNum
//	fmt.Printf("[CPU]  Physcial number: %d    Logscial number:  %d\n", cpuPhysicalNum, cpuLogicalNum)
//
//	terminalWriter := uilive.New()
//	terminalWriter.Start()
//
//	for {
//
//		cpuPercentList := internal.CPUUsage()
//		_, err := fmt.Fprintf(terminalWriter, "[CPU]  cpu1: %.1f  cpu2: %.1f  cpu3: %.1f  cpu4: %.1f\n", cpuPercentList[0], cpuPercentList[1], cpuPercentList[2], cpuPercentList[3])
//		if err != nil {
//			logger.Log.Error(fmt.Sprintf("Failed to output cpu info to stdout, and err is: %s", err.Error()))
//		}
//
//		_, err = fmt.Fprintf(terminalWriter.Newline(), "[CPU]  cpu5: %.1f  cpu6: %.1f  cpu7: %.1f  cpu8: %.1f\n", cpuPercentList[4], cpuPercentList[5], cpuPercentList[6], cpuPercentList[7])
//		if err != nil {
//			logger.Log.Error(fmt.Sprintf("Failed to output cpu info to stdout, and err is: %s", err.Error()))
//		}
//
//		memoryInfoRes := internal.MemoryInfo()
//		totalMemory, usedMemory := memoryInfoRes.TotalSize, memoryInfoRes.UsedSize
//		_, err = fmt.Fprintf(terminalWriter.Newline(), "[RAM]  Total size: %.2fG    Used size: %.2fG\n", totalMemory, usedMemory)
//		if err != nil {
//			logger.Log.Error(fmt.Sprintf("Failed to output cpu info to stdout, and err is: %s", err.Error()))
//		}
//
//		time.Sleep(1 * time.Second)
//	}
//}
