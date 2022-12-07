package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type DirEntry struct {
	IsDir    bool
	Path     string
	Size     int
	Children []DirEntry
}

var calculatedDirSize = make(map[string]int)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	directories := parseInput(input)

	var answer int
	for path := range directories {
		size := getDirectorySize(path, directories)
		if size <= 100000 {
			answer += size
		}
	}
	return answer
}

func part2() int {
	directories := parseInput(input)

	var directorySizes []int
	for path := range directories {
		directorySizes = append(directorySizes, getDirectorySize(path, directories))
	}

	sort.Slice(directorySizes, func(i, j int) bool { return directorySizes[i] < directorySizes[j] })

	const totalDiskSpace = 70000000
	const requiredSpace = 30000000
	diskUsage := getDirectorySize("/", directories)
	amountToDelete := requiredSpace - (totalDiskSpace - diskUsage)
	for _, size := range directorySizes {
		if size >= amountToDelete {
			return size
		}
	}
	return -1
}

func parseInput(input string) map[string][]DirEntry {
	directories := make(map[string][]DirEntry)
	pwd := "/"
	cmds := strings.Split(input, "\n")
	for i := 0; i < len(cmds); i++ {

		var cmd, dir string
		fmt.Sscanf(cmds[i], "$ %s %s", &cmd, &dir)

		switch cmd {
		case "cd":
			pwd = cd(pwd, dir)
			break
		case "ls":
			j := i + 1
			for ; j < len(cmds); j++ {
				if cmds[j][0] == '$' {
					break
				}
				var d, name string
				fmt.Sscanf(cmds[j], "%s %s", &d, &name)

				var n int
				n, _ = strconv.Atoi(d)

				directories[pwd] = append(
					directories[pwd],
					DirEntry{
						IsDir:    d == "dir",
						Path:     cd(pwd, name),
						Size:     n,
						Children: make([]DirEntry, 0),
					},
				)
			}
			i = j - 1
			break
		}
	}
	return directories
}

func getDirectorySize(dir string, directories map[string][]DirEntry) int {
	if _, ok := calculatedDirSize[dir]; ok {
		return calculatedDirSize[dir]
	}
	var size int
	for _, dirEntry := range directories[dir] {
		if dirEntry.IsDir {
			size += getDirectorySize(cd(dir, dirEntry.Path), directories)
		} else {
			size += dirEntry.Size
		}
	}
	calculatedDirSize[dir] = size
	return calculatedDirSize[dir]
}

func cd(pwd, next string) string {
	switch {
	case next == "..":
		dest := pwd[:strings.LastIndex(pwd, "/")]
		if dest == "" {
			return "/"
		}
		return dest
	case next[0] == '/':
		return next
	case pwd[len(pwd)-1] != '/':
		return pwd + "/" + next
	default:
		return pwd + next
	}
}
