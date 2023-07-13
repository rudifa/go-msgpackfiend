// Package util implements the hexdump function.
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
import "fmt"

// Hexdump returns a hexdump of the given data
func Hexdump(data []byte) string {
	var result string
	for i := 0; i < len(data); i += 16 {
		row := data[i:min(i+16, len(data))]
		hex := ""
		ascii := ""
		for j := 0; j < len(row); j++ {
			hex += fmt.Sprintf("%02x ", row[j])
			if row[j] >= 32 && row[j] <= 126 {
				ascii += string(row[j])
			} else {
				ascii += "."
			}
		}
		result += fmt.Sprintf("%08x: %-48s %s\n", i, hex, ascii)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
