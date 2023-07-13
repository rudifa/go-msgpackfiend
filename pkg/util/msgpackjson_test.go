// Package util_test is a test package for the util package.
package util_test

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
	"fmt"
	"testing"

	"github.com/rudifa/msgpackfiend/pkg/util"
)

func TestJSONToMsgpack(t *testing.T) {

	// Read JSON data from file
	infile := "data/data.json"
	jsonData := util.Read(infile)

	// Convert JSON to MessagePack
	msgpackData := util.JSONToMsgpack(jsonData)

	// Write MessagePack data to file
	msgpackOut := "data/data.msgpack"
	util.Write([]byte(msgpackData), msgpackOut)

	fmt.Println("Conversion to MessagePack successful")

	// Read MessagePack data from file
	msgpackIn := "data/data.msgpack"
	msgpackData = util.Read(msgpackIn)

	// Convert MessagePack to JSON
	jsonData = util.MsgpackToJSON(msgpackData)

	// Write JSON data to file
	jsonOutfile := "data/data2.json"
	util.Write([]byte(jsonData), jsonOutfile)

	fmt.Println("Conversion to JSON successful")
}
