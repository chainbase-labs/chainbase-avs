package prometheus

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"runtime"
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

func GetCPUModel() string {
	switch runtime.GOOS {
	case "darwin":
		return getMacOSCPUModel()
	case "linux":
		return getLinuxCPUModel()
	default:
		return fmt.Sprintf("Unsupported OS: %s", runtime.GOOS)
	}
}

func getMacOSCPUModel() string {
	cmd := exec.Command("sysctl", "-n", "machdep.cpu.brand_string")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}

func getLinuxCPUModel() string {
	cmd := exec.Command("cat", "/proc/cpuinfo")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	for _, line := range strings.Split(string(output), "\n") {
		if strings.HasPrefix(line, "model name") {
			return strings.TrimSpace(strings.Split(line, ":")[1])
		}
	}
	return "unknown"
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
