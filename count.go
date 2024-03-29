package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func wordCount(cCtx *cli.Context, file string) (string, error) {
	content, err := os.ReadFile(file)
	words := strings.Fields(string(content))
	wordCount := len(words)
	countString := checkErrorAndFormatOutput(wordCount, file, err)
	return countString, nil
}

func byteCount(cCtx *cli.Context, file string) (string, error) {
	content, err := os.ReadFile(file)
	byteCount := len(content)
	countString := checkErrorAndFormatOutput(byteCount, file, err)
	return countString, nil
}

func lineCount(cCtx *cli.Context, file string) (string, error) {
	openFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer openFile.Close()
	scanner := bufio.NewScanner(openFile)

	lineNo := 0

	for scanner.Scan() {
		lineNo++
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	countString := checkErrorAndFormatOutput(lineNo, file, err)
	return countString, nil
}

func getAllCount(cCtx *cli.Context, file string) string {
	wordCountString, _ := wordCount(cCtx, file)
	byteCountString, _ := byteCount(cCtx, file)
	lineCountString, _ := lineCount(cCtx, file)

	wordCountValue := strings.Fields(wordCountString)[0]
	byteCountValue := strings.Fields(byteCountString)[0]
	lineCountValue := strings.Fields(lineCountString)[0]
	charCountValue := byteCountValue

	allCount := fmt.Sprintf("%v %v %v %v %v", lineCountValue, wordCountValue, charCountValue, byteCountValue, file)

	return allCount
}

func checkErrorAndFormatOutput(count int, file string, err error) string {
	if err != nil {
		log.Fatal(err)
	}
	countString := fmt.Sprintf("%v %-5v", count, file)
	return countString
}
