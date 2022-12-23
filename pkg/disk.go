package pkg

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"

	zlog "github.com/rs/zerolog/log"
)

const FileSystem = "NTFS"

// ListExternalDisk return partition & volume from the external disk
func ListExternalDisk() (map[string]string, error) {
	commands := fmt.Sprintf("diskutil list | grep '%s' | awk '{print $(NF)}'", FileSystem)

	diskResultByte, err := exec.Command("bash", "-c", commands).Output()
	if err != nil {
		return nil, err
	}

	mappedDisk := getMappedDisk(string(diskResultByte))
	return mappedDisk, nil
}

func getMappedDisk(info string) map[string]string {
	regDisk := regexp.MustCompile(`(disk\d*)?s\d*`)

	mappedDisk := make(map[string]string)

	scanner := bufio.NewScanner(strings.NewReader(info))
	for scanner.Scan() {
		line := scanner.Text()
		match := regDisk.FindStringSubmatch(line)

		volume := fmt.Sprintf("/dev/%s", match[1])
		mappedDisk[match[0]] = volume
	}

	return mappedDisk
}

// MountDisk mounts partition to the given volume dir name
func MountDisk(disk string, volumeDirName string) error {
	zlog.Info().Msg("mounting volume is in progres...")

	commands := fmt.Sprintf("mount -t ntfs -o rw,auto,nobrowse /dev/%s /Volumes/%s", disk, volumeDirName)

	_, err := exec.Command("bash", "-c", commands).Output()
	if err != nil {
		return err
	}
	return nil
}

func UnmountDisk(volumeDirName string) error {
	zlog.Info().Msg("unmounting volume is in progres...")

	commands := fmt.Sprintf("umount /Volumes/%s", volumeDirName)

	_, err := exec.Command("bash", "-c", commands).Output()
	if err != nil {
		return err
	}

	// giving some time to make sure that the volume unmounted safely
	time.Sleep(2 * time.Second)

	return nil
}
