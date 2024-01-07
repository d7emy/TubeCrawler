package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func appendMultiCSV(path string, data [][]string) error {
	if !fileExists(path) {
		data = append([][]string{{"ID", "Title", "Channel Name", "Channel url", "Published Time", "Long", "View Count", "Browse ID"}}, data...)
	}
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	filePath := filepath.Join(dir, filepath.Base(path))
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	for _, line := range data {
	again:
		err = writer.Write(line)
		if err != nil {
			fmt.Println("Error writer.Write()", err)
			time.Sleep(time.Second)
			goto again
		}

		writer.Flush()
	}
	return nil
}

func appendCSV(path string, line []string) error {
	if !fileExists(path) {
		appendMultiCSV(path, [][]string{line})
		return nil
	}
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("Error os.MkdirAll()", err)
		return err
	}
	filePath := filepath.Join(dir, filepath.Base(path))
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error os.OpenFile()", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
again:
	err = writer.Write(line)
	if err != nil {
		fmt.Println("Error writer.Write()", err)
		time.Sleep(time.Second)
		goto again
	}

	writer.Flush()
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	panic(err)
}

func readCSV(filePath string) [][]string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Error os.Open():", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	//reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}
	return records
}
