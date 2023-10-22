package temp

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ResponseStat struct {
	Name  string  `json:"name"`
	Value int `json:"value"`
}

type TempInfo struct {
	temp []int
}

type TempTracker struct {
	mu sync.Mutex

	currTemp TempInfo
}

func (t *TempTracker) GetAvgTemp() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	result := 0
	for _, v := range t.currTemp.temp {
		result += v
	}
	if len(t.currTemp.temp) == 0 {
		return 0
	}
	return result / len(t.currTemp.temp)
}

func (t *TempTracker) StartTracking() error {
	for {
		err := t.updateTemp()
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
}

func (t *TempTracker) GetTempInfo() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.GetAvgTemp()
}

func (t *TempTracker) updateTemp() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	tempInfo, err := getTempInfo()
	if err != nil {
		return err
	}
	t.currTemp = tempInfo
	return nil
}

func getTempInfo() (TempInfo, error) {
	array := []int{}
	//TODO: read all files not just two
	for i := 0; i < 2; i++ {
		path := filepath.Join("/sys/class/thermal/thermal_zone" + strconv.Itoa(i), "temp")
		log.Println(path)
		file, err := os.ReadFile(path)
		if err != nil {
			return TempInfo{}, err
		}
		dataFromFile := strings.TrimSuffix(string(file), "\n")
		temp, err := strconv.Atoi(dataFromFile)
		if err != nil {
			log.Println(err)
			return TempInfo{}, err
		}
		log.Println(file)
		array = append(array, temp)
	}
	return TempInfo{array}, nil
}