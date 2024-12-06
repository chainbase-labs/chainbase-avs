package metrics

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v3/mem"
)

func GetOutboundIP() string {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "unknown"
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(ip))
}

func GetCPUCore() uint32 {
	return uint32(runtime.NumCPU())
}

func GetTotalMemory() float64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0
	}
	return float64(v.Total) / (1024 * 1024)
}

func GetJobManagerStatus(jobManagerHost, jobManagerPort string) float64 {
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/config", jobManagerHost, jobManagerPort))

	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0

	}

	return 1
}
