package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type FileId = int
type FreeMemoryId = int

type FilePart struct {
	from int
	to   int
}

type FreeMemoryPart struct {
	from int
	to   int
}
type Files = map[FileId][]FilePart
type FreeMemory = map[FreeMemoryId][]FreeMemoryPart
type Input struct {
	files      Files
	freeMemory FreeMemory
}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	isFile := true
	fileCounter := 0
	memoryCounter := 0
	memoryPos := 0
	files := make(map[FileId][]FilePart)
	freeMemory := make(map[FreeMemoryId][]FreeMemoryPart)
	for char := range input {
		if isFile {
			fileLength := char - '0'
			files[fileCounter] = []FilePart{}
			files[fileCounter] = append(files[fileCounter], FilePart{from: memoryPos, to: memoryPos + fileLength})
			memoryPos += fileLength
			fileCounter++
		} else {
			memoryLength := char - '0'
			freeMemory[memoryCounter] = []FreeMemoryPart{}
			freeMemory[memoryCounter] = append(freeMemory[memoryCounter], FreeMemoryPart{from: memoryPos, to: memoryPos + memoryLength})
			memoryPos += memoryLength
			memoryCounter++
		}
		isFile = !isFile
	}
	return Input{
		files:      files,
		freeMemory: freeMemory,
	}
}

func isContiguous(files Files) bool {
	// check that there is no space between files
	currentPos := files[0][0].from
	for _, file := range files {
		if file[0].from != currentPos {
			return false
		}
		currentPos = file[0].to
	}

}
func part1(input Input) int {
	return 0
}

func part2(input Input) int {
	return 0
}

func main() {
	startTime := time.Now()
	input := parseInput(input)
	fmt.Println("Parsed input in", time.Since(startTime))

	startTime = time.Now()
	res1 := part1(input)
	fmt.Println("Part 1:", res1, "in", time.Since(startTime))

	startTime = time.Now()
	res2 := part2(input)
	fmt.Println("Part 2:", res2, "in", time.Since(startTime))
}
