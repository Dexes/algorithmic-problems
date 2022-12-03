package main

import (
	"sort"
	"strings"
)

func removeSubfolders(folders []string) []string {
	sort.Strings(folders)

	rIndex, wIndex := 0, 0
	for rIndex < len(folders) {
		folders[wIndex] = folders[rIndex]
		for rIndex++; rIndex < len(folders) && isSubfolder(folders[wIndex], folders[rIndex]); rIndex++ {
		}

		wIndex++
	}

	return folders[:wIndex]
}

func isSubfolder(folder, subfolder string) bool {
	return len(subfolder) > len(folder) && subfolder[len(folder)] == '/' && strings.HasPrefix(subfolder, folder)
}
