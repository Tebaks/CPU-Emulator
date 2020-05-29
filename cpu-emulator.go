package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type line struct {
	lineNumber int
	operation  string
	value1     int
}

func main() {
	accumulator := 0
	var memory [256]int
	programCounter := 0
	flag := 0
	var program []line
	working := true

	readFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		tokenizer := strings.Fields(fileScanner.Text())
		l := line{}
		l.lineNumber, err = strconv.Atoi(tokenizer[0])
		if err != nil {
			log.Fatalf("Line number must be integer.")
		}
		l.operation = tokenizer[1]
		if len(tokenizer) > 2 {
			l.value1, err = strconv.Atoi(tokenizer[2])
			if err != nil {
				log.Fatalf("value must be integer.")
			}

		program = append(program, l)
	}

	readFile.Close()

	for working {
		line := program[programCounter]
		switch line.operation {
		case "START":
			break
		case "LOAD":
			accumulator = line.value
			break
		case "LOADM":
			accumulator = memory[line.value]
			break
		case "STORE":
			memory[line.value] = accumulator
			break
		case "CMPM":
			if accumulator > memory[line.value] {
				flag = 1
			} else if accumulator < memory[line.value] {
				flag = -1
			} else {
				flag = 0
			}
			break
		case "CJMP":
			if flag > 0 {
				programCounter = line.value - 1
			}
			break
		case "JMP":
			programCounter = line.value - 1
			break
		case "ADD":
			accumulator += line.value
			break
		case "ADDM":
			accumulator += memory[line.value]
			break
		case "SUB":
			accumulator -= line.value
			break
		case "SUBM":
			accumulator -= memory[line.value]
			break
		case "MUL":
			accumulator *= line.value
			break
		case "MULM":
			accumulator *= memory[line.value]
			break
		case "DISP":
			fmt.Println(accumulator)
			break
		case "HALT":
			working = false
			break
		}
		programCounter++
	}
}
