package backend

import (
	"bytes"
	"os/exec"
	"strings"
	"syscall"
)

type WslService struct{}

func NewWslService() *WslService {
	return &WslService{}
}

func (s *WslService) GetDistros() ([]WslDistro, error) {
	allDistrosOutput := s.runWslCommand("--list", "--quiet")
	allList := s.parseWslList(allDistrosOutput)

	result := []WslDistro{}
	for _, name := range allList {
		result = append(result, WslDistro{
			Name: name,
		})
	}
	return result, nil
}

func (s *WslService) ScanPaths(distroName string, toolsToScan []string) ([]SelectableTool, error) {
	result := []SelectableTool{}

	for _, tool := range toolsToScan {
		cmd := exec.Command("wsl", "-d", distroName, "-e", "bash", "-c", "which "+tool+" 2>/dev/null")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, err := cmd.Output()
		if err == nil {
			binaryPath := strings.TrimSpace(strings.ReplaceAll(string(out), "\x00", ""))
			if binaryPath != "" {
				result = append(result, SelectableTool{
					ToolName:   tool,
					WslPath:    binaryPath,
					IsSelected: false,
				})
			}
		}
	}
	return result, nil
}

func (s *WslService) runWslCommand(args ...string) string {
	cmd := exec.Command("wsl", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	// WSL often outputs in UTF-16LE. Stripping null bytes makes it ASCII compatible.
	cleaned := bytes.ReplaceAll(out, []byte{0}, []byte{})
	return string(cleaned)
}

func (s *WslService) parseWslList(output string) []string {
	result := []string{}
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		clean := strings.TrimSpace(line)
		clean = strings.ReplaceAll(clean, "\r", "")
		if clean != "" {
			result = append(result, clean)
		}
	}
	return result
}
