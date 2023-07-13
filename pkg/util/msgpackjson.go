// Package util implements the conversion functions between JSON and Msgpack formats.
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
	"encoding/json"
	"log"

	"github.com/vmihailenco/msgpack/v5"
)

// JSONToMsgpack converts JSON data to MessagePack data
func JSONToMsgpack(jsonData []byte) []byte {
	var data interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal("Error unmarshalling JSON data:", err)
	}

	msgpackData, err := msgpack.Marshal(data)
	if err != nil {
		log.Fatal("Error marshalling MessagePack data:", err)
	}

	return msgpackData
}

// MsgpackToJSON converts MessagePack data to JSON data
func MsgpackToJSON(msgpackData []byte) []byte {
	var data interface{}
	err := msgpack.Unmarshal(msgpackData, &data)
	if err != nil {
		log.Fatal("Error unmarshalling MessagePack data:", err)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error marshalling JSON data:", err)
	}

	return jsonData
}
