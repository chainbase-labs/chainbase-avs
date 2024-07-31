package prometheus

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetCPUModel(t *testing.T) {
	model := GetCPUModel()
	if model == "" {
		t.Error("GetCPUModel returned an empty string")
	}
	if model == "unknown" {
		t.Log("GetCPUModel returned 'unknown'. This might be expected in some environments.")
	}
	t.Logf("CPU Model: %s", model)

	if runtime.GOOS != "darwin" && runtime.GOOS != "linux" {
		if model != fmt.Sprintf("Unsupported OS: %s", runtime.GOOS) {
			t.Errorf("Expected 'Unsupported OS' message for non-darwin/linux OS, got: %s", model)
		}
	}
}

func TestGetOutboundIP(t *testing.T) {
	ip := GetOutboundIP()

	if ip == "unknown" {
		t.Log("getOutboundIP returned 'unknown'. This might be expected in some environments.")
	}
	t.Logf("IP: %s", ip)
}
