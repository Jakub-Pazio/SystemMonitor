package cpu

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ResponseStat struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type CpuInfo struct {
	userTime     int
	userNiceTime int
	systemTime   int
	idleTime     int
	ioWaitTime   int
	irqTime      int
	softIrqTime  int
}

type CpuTracker struct {
	mu sync.Mutex

	prev CpuInfo
	curr CpuInfo
}

func (c *CpuTracker) cpuPercentUsage() float64 {
	prevTotal := c.prev.userTime + c.prev.userNiceTime + c.prev.systemTime + c.prev.idleTime + c.prev.ioWaitTime + c.prev.irqTime + c.prev.softIrqTime
	currTotal := c.curr.userTime + c.curr.userNiceTime + c.curr.systemTime + c.curr.idleTime + c.curr.ioWaitTime + c.curr.irqTime + c.curr.softIrqTime
	totalDelta := currTotal - prevTotal
	idleDelta := c.curr.idleTime - c.prev.idleTime

	if totalDelta == 0 {
		return 0.0 // Avoid division by zero
	}
	return float64(totalDelta-idleDelta) / float64(totalDelta)
}

func (c *CpuTracker) updateUsage() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.prev = c.curr
	cpuInfo, err := getCpuInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.curr = cpuInfo
}

func (c *CpuTracker) StartTracking() error {
	for {
		c.updateUsage()
		time.Sleep(1 * time.Second)
	}
}

func (c *CpuTracker) GetCpuPercentUsage() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cpuPercentUsage()
}

func getCpuInfo() (CpuInfo, error) {
	data, err := os.ReadFile("/proc/stat")
	if err != nil {
		return CpuInfo{}, err
	}
	cpuLine := strings.Split(string(data), "\n")[0]
	cpuInfo := strings.Fields(cpuLine)[1:]
	userTime, err1 := strconv.Atoi(cpuInfo[0])
	userNiceTime, err2 := strconv.Atoi(cpuInfo[1])
	systemTime, err3 := strconv.Atoi(cpuInfo[2])
	idleTime, err4 := strconv.Atoi(cpuInfo[3])
	ioWaitTime, err5 := strconv.Atoi(cpuInfo[4])
	irqTime, err6 := strconv.Atoi(cpuInfo[5])
	softIrqTime, err7 := strconv.Atoi(cpuInfo[6])
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil {
		return CpuInfo{}, err
	}
	return CpuInfo{userTime, userNiceTime, systemTime, idleTime, ioWaitTime, irqTime, softIrqTime}, nil
}