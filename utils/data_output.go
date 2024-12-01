package utils

import "os"

func OutputAnalyzerDataToFile(data map[string][][]string, filepath string) bool {
	var success = true
	var err error
	var f *os.File
	f, err = os.Create(filepath)
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
