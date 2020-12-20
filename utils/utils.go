package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// GetExecutablePath gets the path of the current executable.
func GetExecutablePath() (string, error) {
	ex, err := os.Executable()

	if err != nil {
		return "", err
	}

	return filepath.Dir(ex), nil
}

// GetImageFromFilePath Reads the a file path and returns its content as bytes
func GetImageFromFilePath(filePath string) []byte {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes) // <--------------- here!

	return bytes
}

// SafeJoinPaths joins `path1` and `path2` and applies the correct separator.
func SafeJoinPaths(path1, path2 string) string {
	return filepath.FromSlash(filepath.Join(path1, path2))
}

// StringToInt converts a string to an int, and if the conversion fails, it uses the `defaultValue`.
func StringToInt(str string, defaultValue int) int {
	result, err := strconv.Atoi(str)

	if err != nil {
		return defaultValue
	}

	return result
}

// IntToString Converts Int to String (base 10)
func IntToString(value int) string {
	return strconv.FormatInt(int64(value), 10)
}

// ConvertJSONToStruct converts a json string to a struct
func ConvertJSONToStruct(text []byte, model interface{}) error {
	err := json.Unmarshal(text, &model)

	if err != nil {
		return err
	}

	return nil
}

// LoadJSONFromFile loads a JSON file from `path` and populated the given `model`.
func LoadJSONFromFile(path string, model interface{}) error {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &model)

	if err != nil {
		return err
	}

	return nil
}

// BytesToString casts bytes to string
func BytesToString(data []byte) string {
	return string(data[:])
}

// WriteJSONToFile writes the given `model` struct as JSON to the file at `path`
func WriteJSONToFile(path string, model interface{}) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(model)

	if err != nil {
		return err
	}

	return nil
}

// ParseConfig reads a `config.json` file on the same directory as the executable and parses its JSON to the given model.
func ParseConfig(model interface{}) error {
	path := SafeJoinPaths("./", "config.json")

	return LoadJSONFromFile(path, model)
}

// LogError logs a error inside a folder "logs" in the same location as the executable
func LogError(text string) error {
	path := "logs/" + time.Now().UTC().Format("20060102150405") + ".txt"

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer f.Close()

	if _, err := f.WriteString(text); err != nil {
		log.Println(err)
	}

	return nil
}
