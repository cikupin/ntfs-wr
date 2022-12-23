package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/oklog/ulid/v2"
)

// IsVolumeExist checks whether the given volume dir is exist or not
func IsVolumeExist(volumeDirName string) bool {
	_, err := os.Stat(fmt.Sprintf("/Volumes/%s", volumeDirName))
	return !os.IsNotExist(err)

}

// MakeVolumeDir creates a folder on /Volumes to mount external disk partition
// A returned folder name will assigned using ULID format
func MakeVolumeDir() (string, error) {
	dirName := strings.ToLower(ulid.Make().String())

	err := os.MkdirAll(fmt.Sprintf("/Volumes/%s", dirName), 0777)
	if err != nil {
		return "", err
	}
	return dirName, nil
}

// OpenVolumeDir will open finder on the selected volume path
func OpenVolumeDir(volumeDirName string) error {
	commands := fmt.Sprintf("open /Volumes/%s", volumeDirName)

	_, err := exec.Command("bash", "-c", commands).Output()
	if err != nil {
		return err
	}
	return nil
}

// RemoveVolumeDir removes the given directory on /Volumes path
func RemoveVolumeDir(volumeDirName string) error {
	return os.RemoveAll(fmt.Sprintf("/Volumes/%s", volumeDirName))
}
