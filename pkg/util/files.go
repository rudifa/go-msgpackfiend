// Package util implements the file operations.
package util

/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
import (
	"io"
	"log"
	"os"
)

// Read reads data from a file or stdin
func Read(infile string) []byte {
	var data []byte
	var err error

	if infile == "" {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal("Error reading from stdin:", err)
		}
	} else {
		f, err := os.Open(infile)
		if err != nil {
			log.Fatal("Error opening file:", err)
		}
		defer f.Close()

		data, err = io.ReadAll(f)
		if err != nil {
			log.Fatal("Error reading file:", err)
		}
	}

	return data
}

// Write writes data to a file or stdout
func Write(data []byte, outfile string) {
	var f io.Writer
	var err error

	if outfile == "" {
		f = os.Stdout
	} else {
		f, err = os.Create(outfile)
		if err != nil {
			log.Fatal("Error creating file:", err)
		}
		defer f.(*os.File).Close()
	}

	_, err = f.Write(data)
	if err != nil {
		log.Fatal("Error writing file:", err)
	}
}
