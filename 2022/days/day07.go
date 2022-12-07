package days

import (
	"aoc-2022/util"
	"fmt"
	"log"
	"sort"
	"strings"
)

func Day7() {
	fileContents := util.ReadFileLines("inputs/day07-1.txt")
	fmt.Printf("Part 1: %d\n", day7Part1(fileContents))
	fmt.Printf("Part 2: %d\n", day7Part2(fileContents))
}

type directory struct {
	name     string
	parent   *directory
	children []*directory
	// files    []file
	size int
}

type file struct {
	name string
	size int
}

func day7Part1(fileContents []string) int {
	rootDir := readDirectories(fileContents)
	// fmt.Printf("Dirs: %+v\n", rootDir)
	return getPart1Sum(*rootDir)
}

func readDirectories(lines []string) *directory {
	rootDir := directory{name: "/"}
	// fmt.Printf("Address of root: %p\n", &rootDir)
	curDir := &rootDir

	for _, l := range lines {
		// fmt.Println(l)
		if strings.HasPrefix(l, "$ cd") {
			splits := strings.Split(l, "cd ")
			dirName := splits[1]
			switch dirName {
			case "/":
				// fmt.Println("Moving to root")
				curDir = &rootDir
			case "..":
				// fmt.Printf("Moving up a dir to %s\n", curDir.parent.name)
				curDir = curDir.parent
			default:
				// fmt.Println("Moving to " + dirName)
				found := false
				for _, d := range curDir.children {
					// fmt.Printf("Checking %v\n", d)
					if d.name == dirName {
						// fmt.Printf("Moving to %v\n", *d)
						curDir = d
						found = true
						break
					}
				}
				if !found {
					log.Fatalf("Couldn't find directory %s\n", dirName)
				}
			}
		} else {
			if strings.HasPrefix(l, "$ ls") {
				// fmt.Println("Ignoring ls")
				continue
			} else if strings.HasPrefix(l, "dir") {
				splits := strings.Split(l, "dir ")
				// fmt.Printf("Adding parent %p\n", curDir)
				d := directory{name: splits[1], parent: curDir}
				// fmt.Printf("Adding directory %v\n", d)
				curDir.children = append(curDir.children, &d)
				// fmt.Printf("Directories in %s: %v\n", curDir.name, curDir.children)
			} else {
				splits := strings.Split(l, " ")
				size := util.GetIntFromString(splits[0])
				// f := file{name: splits[1], size: size}
				// curDir.files = append(curDir.files, f)
				curDir.size += size
				// fmt.Printf("File %s, size %d\n", splits[1], size)

				propDir := curDir.parent
				for propDir != nil {
					// fmt.Printf("Parent: %v\n", propDir.name)
					// fmt.Printf("Increasing %s by %d\n", propDir.name, size)
					propDir.size += size
					propDir = propDir.parent
				}
				// fmt.Printf("Files in %s: %v\n", curDir.name, curDir.files)
			}
		}
	}

	return &rootDir
}

func getPart1Sum(dir directory) int {
	vals := part1Recursive(dir, make([]int, 0))
	// fmt.Println(vals)
	var total int
	for _, v := range vals {
		total += v
	}
	return total
}

func part1Recursive(dir directory, vals []int) []int {
	var total int
	// fmt.Println(dir.name)
	if len(dir.children) == 0 {
		if dir.size <= 100000 {
			// fmt.Printf("Dir %s size %d\n", dir.name, dir.size)
			return append(vals, dir.size)
		}
		// fmt.Printf("Dir %s is too big\n", dir.name)
		return vals
	}
	newVals := make([]int, len(vals))
	for _, d := range dir.children {
		newVals = append(newVals, part1Recursive(*d, vals)...)
		// fmt.Println(newVals)
		if d.size <= 100000 {
			total += d.size
		}
	}
	// fmt.Printf("Dir %s has size %d\n", dir.name, dir.size)
	if dir.size <= 100000 {
		return append(newVals, dir.size)
	}
	return newVals
}

func day7Part2(fileContents []string) int {
	rootDir := readDirectories(fileContents)
	spaceUnused := 70000000 - rootDir.size
	spaceNeeded := 30000000 - spaceUnused
	// fmt.Println("Space needed ", spaceNeeded)

	vals := part2Recursive(*rootDir, make([]int, 0))
	sort.Ints(vals)

	for _, v := range vals {
		if v > spaceNeeded {
			return v
		}
	}
	log.Fatalln("No directory big enough")
	return 0
}

func part2Recursive(dir directory, vals []int) []int {
	var total int
	if len(dir.children) == 0 {
		return append(vals, dir.size)
	}
	newVals := make([]int, len(vals))
	for _, d := range dir.children {
		newVals = append(newVals, part2Recursive(*d, vals)...)
		total += d.size
	}
	return append(newVals, dir.size)
}
