package utils

import (
	"fmt"
	"os"
)

func OutputAnalyzerDataToFile(data map[string][][]string, fileName string) bool {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("An error occurred while creating the file:", err)
		return false
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("An error occurred while closing the file:", err)
		}
	}(file)

	for host, results := range data {
		_, err := file.WriteString(fmt.Sprintf("%s\n", host))
		if err != nil {
			fmt.Println("An error occurred while writing to the file:", err)
			return false
		}
		for _, result := range results {
			_, err := file.WriteString(fmt.Sprintf("%s\n", result))
			if err != nil {
				fmt.Println("An error occurred while writing to the file:", err)
				return false
			}
		}
		_, err = file.WriteString("\n")
		if err != nil {
			fmt.Println("An error occurred while writing to the file:", err)
			return false
		}
	}

	return true
}

func OutputAnalyzerDataToFileAppend(data map[string][][]string, filepath string) bool {
	var success = true
	var err error
	var f *os.File
	f, err = os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		success = false
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			success = false
		}
	}(f)
	for key, value := range data {
		_, err = f.WriteString(key + "\n")
		if err != nil {
			success = false
		}
		for _, v := range value {
			_, err = f.WriteString(v[0] + " " + v[1] + "\n")
			if err != nil {
				success = false
			}
		}
	}
	return success
}

func ReadAnalyzerDataFromFile(filepath string) (map[string][][]string, bool) {
	var success = true
	var err error
	var f *os.File
	f, err = os.Open(filepath)
	if err != nil {
		success = false
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			success = false
		}
	}(f)
	var data = make(map[string][][]string)
	var key string
	var value []string
	var line string
	for {
		_, err = f.Read([]byte(line))
		if err != nil {
			break
		}
		if line == "\n" {
			data[key] = append(data[key], value)
			value = nil
		} else {
			if key == "" {
				key = line
			} else {
				value = append(value, line)
			}
		}
	}
	return data, success
}
