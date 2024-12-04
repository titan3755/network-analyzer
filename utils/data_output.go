package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
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

//func OutputAnalyzerDataToFileAppend(data map[string][][]string, filepath string) bool {
//	var success = true
//	var err error
//	var f *os.File
//	f, err = os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
//	if err != nil {
//		success = false
//	}
//	defer func(f *os.File) {
//		err := f.Close()
//		if err != nil {
//			success = false
//		}
//	}(f)
//	for key, value := range data {
//		_, err = f.WriteString(key + "\n")
//		if err != nil {
//			success = false
//		}
//		for _, v := range value {
//			_, err = f.WriteString(v[0] + " " + v[1] + "\n")
//			if err != nil {
//				success = false
//			}
//		}
//	}
//	return success
//}

func SplitLines(s string) []string {
	// Helper function to split the string into lines
	return strings.Split(s, "\n")
}

func SplitLine(s string) []string {
	// Split the line based on spaces (to capture individual components of the speedtest)
	return strings.Fields(s)
}

func ReadAnalyzerSpeedTestDataFromFile(filepath string) map[string][][]string {
	var data = make(map[string][][]string)
	var host string
	var results [][]string
	var result []string
	var file, err = os.Open(filepath)
	if err != nil {
		fmt.Println("An error occurred while opening the file:", err)
		return data
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("An error occurred while closing the file:", err)
		}
	}(file)

	var buf = make([]byte, 1024)
	var n int
	var err2 error
	for {
		n, err2 = file.Read(buf)
		if n == 0 || err2 != nil {
			break
		}
		var s = string(buf[:n])
		var lines = SplitLines(s)

		for _, line := range lines {
			// Skip empty lines
			if line == "" {
				continue
			}

			// Check for a new host (assuming the first non-empty line is the host)
			if !strings.HasPrefix(line, "[") && !strings.HasSuffix(line, "]") {
				// If we already have a host, store the results and reset
				if host != "" {
					data[host] = results
					results = make([][]string, 0)
				}
				host = line
				continue
			}

			// Now, we process the line inside the square brackets
			// Strip off the brackets
			re := regexp.MustCompile(`\[(.*)]`)
			matches := re.FindStringSubmatch(line)
			if len(matches) > 1 {
				// Now, we have the string inside the brackets
				result = SplitLine(matches[1])
				results = append(results, result)
			}
		}
	}

	// Don't forget to add the last host's results if there are any
	if host != "" {
		data[host] = results
	}

	return data
}

func ReadAnalyzerStabilityTestDataFromFile(filepath string) map[string][][]string {
	var data = make(map[string][][]string)
	var host string
	var results [][]string
	var result []string
	var file, err = os.Open(filepath)
	if err != nil {
		fmt.Println("An error occurred while opening the file:", err)
		return data
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("An error occurred while closing the file:", err)
		}
	}(file)

	var buf = make([]byte, 1024)
	var n int
	var err2 error
	for {
		n, err2 = file.Read(buf)
		if n == 0 || err2 != nil {
			break
		}
		var s = string(buf[:n])
		var lines = SplitLines(s)

		for _, line := range lines {
			// Skip empty lines
			if line == "" {
				continue
			}

			// If the line doesn't start with a bracket, it is a new host
			if !strings.HasPrefix(line, "[") && !strings.HasSuffix(line, "]") {
				// If we have a host, store the previous results before resetting
				if host != "" {
					data[host] = results
					results = make([][]string, 0)
				}
				// Set the new host
				host = line
				continue
			}

			// Now, we process the line inside the square brackets
			re := regexp.MustCompile(`\[(.*)]`)
			matches := re.FindStringSubmatch(line)
			if len(matches) > 1 {
				// Now, we have the string inside the brackets
				result = SplitLine(matches[1])
				results = append(results, result)
			}
		}
	}

	// Add the last host's results if any
	if host != "" {
		data[host] = results
	}

	return data
}
