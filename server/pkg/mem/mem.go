package mem

import (
	"bytes"
	"os"
	"strconv"
	"sync"
	"time"
)

type ResponseStat struct {
	Name  string  `json:"name"`
	Total int `json:"total"`
	Free int `json:"free"`
	Avail int `json:"avail"`
}

type MemInfo struct {
	Total int
	Free int
	Avail int
}

type MemTracker struct {
	mu sync.Mutex

	currMem MemInfo
}

func (t *MemTracker) GetMemInfo() MemInfo {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.currMem
}

func (t *MemTracker) StartTracking() error {
	for {
		err := t.updateMem()
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
}

func (t *MemTracker) updateMem() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	memInfo, err := getMemInfo()
	if err != nil {
		return err
	}
	t.currMem = memInfo
	return nil
}

func getMemInfo() (MemInfo, error) {
	file, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return MemInfo{}, err
	}
	var total, free, avail int
	for _, line := range bytes.Split(file, []byte("\n")) {
		if bytes.Contains(line, []byte("MemTotal")) {
			total, err = strconv.Atoi(string(bytes.Fields(line)[1]))
			if err != nil {
				return MemInfo{}, err
			}
		}
		if bytes.Contains(line, []byte("MemFree")) {
			free, err = strconv.Atoi(string(bytes.Fields(line)[1]))
			if err != nil {
				return MemInfo{}, err
			}
		}
		if bytes.Contains(line, []byte("MemAvailable")) {
			avail, err = strconv.Atoi(string(bytes.Fields(line)[1]))
			if err != nil {
				return MemInfo{}, err
			}
		}
	}
	return MemInfo{Total: total, Free: free, Avail: avail}, nil
}