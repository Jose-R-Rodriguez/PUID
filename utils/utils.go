package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//Log checks and logs a error
func Log(err error, messages ...string) {
	if err != nil {
		var buffer strings.Builder
		for _, message := range messages {
			buffer.WriteString(message)
		}
		log.Printf("%v\n%v", err.Error(), buffer.String())
		os.Exit(1)
	}
}

/*
LoadJSONFromFile loads a Json's content into a byte array and returns it alongside it's filepointer
so the user may close it, if an error occurs we panic
*/
func LoadJSONFromFile(filePath string) ([]byte, *os.File) {
	jsonFile, err := os.Open(filePath)
	Log(err, "Couldn't load JSON")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue, jsonFile
}

/*
SafeLoadJSONFromFile loads a Json's content into a byte array and returns it alongside it's filepointer
so the user may close it
*/
func SafeLoadJSONFromFile(filePath string) ([]byte, *os.File) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Println("WARNING: ", err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue, jsonFile
}

/*
CompareStringSliceEquality compares 2 string slices for equality
*/
func CompareStringSliceEquality(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for index, eachA := range a {
		if eachA != b[index] {
			return false
		}
	}
	return true
}
