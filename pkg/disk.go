package pkg

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

const fileSystem = "FAT_32"

func ListExternalDisk() (map[string]string, error) {
	commands := fmt.Sprintf("diskutil list | grep '%s' | awk '{print $(NF)}'", fileSystem)

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
