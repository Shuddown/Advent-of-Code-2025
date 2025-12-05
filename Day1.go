package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const url string = "https://adventofcode.com/2025/day/1/input"
const sessionCookie string = "53616c7465645f5f863f44e6e9b87ffb359229ccc8aa89b479fb6619203dc12408ed356a162e68f86d2cbb5e2a697283b26e7568de426d05682d2ee2086c21f0"
const inputFile string = "input.txt"

func getInputFile() {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	}
	req.AddCookie(cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	body, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	os.WriteFile(inputFile, body, 0644)
}

func abs(num int) int {
	if num >= 0 {
		return num
	}

	return -num
}

func main() {
	_, err := os.Stat(inputFile)
	if err != nil {
		getInputFile()
	}
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	partOneZeroCount := 0
	partTwoZeroCount := 0
	currPos := 50
	for scanner.Scan() {
		firstChar := scanner.Text()[0]
		turnAmount, err := strconv.Atoi(scanner.Text()[1:])
		if err != nil {
			panic(err)
		}
		if firstChar == 'L' {
			turnAmount *= -1
		}
		currWasZero := (currPos == 0)
		currPos += turnAmount
		partTwoZeroCount += abs(currPos / 100)
		if currPos <= 0 && !currWasZero {
			partTwoZeroCount += 1
		}
		currPos %= 100
		if currPos < 0 {
			currPos += 100
		}
		fmt.Println(currPos)
		if currPos == 0 {
			partOneZeroCount += 1
		}
	}
	fmt.Println(partOneZeroCount)
	fmt.Println(partTwoZeroCount)
}
