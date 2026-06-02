package backend

import (
	"os"
	"strings"
	"syscall"
	"unsafe"
	"golang.org/x/sys/windows/registry"
)

type PathInjectorService struct{}

func NewPathInjectorService() *PathInjectorService {
	return &PathInjectorService{}
}

func (s *PathInjectorService) InjectPath(dirsToInject []string) error {
	return s.modifyPath(func(currentEntries []string) []string {
		toAdd := []string{}
		for _, dir := range dirsToInject {
			found := false
			for _, entry := range currentEntries {
				if strings.EqualFold(dir, entry) {
					found = true
					break
				}
			}
			if !found {
				toAdd = append(toAdd, dir)
			}
		}
		return append(toAdd, currentEntries...)
	})
}

func (s *PathInjectorService) RemovePath(dirsToRemove []string) error {
	return s.modifyPath(func(currentEntries []string) []string {
		cleaned := []string{}
		for _, entry := range currentEntries {
			shouldRemove := false
			for _, dir := range dirsToRemove {
				if strings.EqualFold(dir, entry) {
					shouldRemove = true
					break
				}
			}
			if !shouldRemove && entry != "" {
				cleaned = append(cleaned, entry)
			}
		}
		return cleaned
	})
}

func (s *PathInjectorService) modifyPath(modifier func([]string) []string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	currentPath, _, err := key.GetStringValue("PATH")
	if err != nil {
		currentPath = ""
	}

	currentEntries := strings.Split(currentPath, ";")
	
	newEntries := modifier(currentEntries)
	newPath := strings.Join(newEntries, ";")
	
	if currentPath != newPath {
		err = key.SetStringValue("PATH", newPath)
		if err != nil {
			return err
		}
		// Also update the environment variable for the current process, just in case
		os.Setenv("PATH", newPath)
		go s.broadcastSettingChange()
	}

	return nil
}

func (s *PathInjectorService) broadcastSettingChange() {
	envVar := syscall.StringToUTF16Ptr("Environment")
	// HWND_BROADCAST = 0xFFFF, WM_SETTINGCHANGE = 0x001A
	// Use SendMessageTimeout
	user32 := syscall.NewLazyDLL("user32.dll")
	sendMessageTimeout := user32.NewProc("SendMessageTimeoutW")
	sendMessageTimeout.Call(
		uintptr(0xFFFF),
		uintptr(0x001A),
		0,
		uintptr(unsafe.Pointer(envVar)),
		2, // SMTO_ABORTIFHUNG
		5000,
		0,
	)
}
