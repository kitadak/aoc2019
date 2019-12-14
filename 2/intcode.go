package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadData(datasetFile string) []int {
	fmt.Println("reading: " + datasetFile)
	file, err := os.Open(datasetFile)
	if err != nil {
		log.Fatal("Failed to load file", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	data, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read line", err)
	}
	stringData := strings.Split(data, ",")
	returnData := make([]int, len(stringData))
	for i := range returnData {
		returnData[i], _ = strconv.Atoi(stringData[i])
	}
	return returnData
}

func OpcodeOne(index int, data []int) {
	data[data[index+3]] = data[data[index+1]] + data[data[index+2]]
	return
}

func OpcodeTwo(index int, data []int) {
	data[data[index+3]] = data[data[index+1]] * data[data[index+2]]
	return
}

func RunProgram(data []int) {
	for i := 0; i < len(data); i += 4 {
		fmt.Println(i, len(data))
		fmt.Println(data)
		switch data[i] {
		case 1:
			OpcodeOne(i, data)
		case 2:
			OpcodeTwo(i, data)
		case 99:
			break
		}
	}
	return
}

func main() {
	fmt.Println("intcode.go")
	if len(os.Args) != 2 {
		log.Fatalf("Usage: ./main <input_file>, got %d", len(os.Args))
	}
	data := LoadData(os.Args[1])
	fmt.Println(data)
	RunProgram(data)
	fmt.Println(data)

}
