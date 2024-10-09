package metrics

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
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

func GetTotalMemory() float64 {
	cmd := exec.Command("grep", "MemTotal", "/proc/meminfo")
	output, err := cmd.Output()
	if err != nil {
		return 0
	}
	memStr := strings.TrimSpace(strings.Split(string(output), ":")[1])
	memKB, err := strconv.ParseFloat(strings.Split(memStr, " ")[0], 64)
	if err != nil {
		return 0
	}
	// convert KB to MB
	return memKB * 1024
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
