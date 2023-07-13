// Package cmd implements the msgpackfiend command tojson.
package cmd

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
	"log"
	"strings"

	"github.com/rudifa/msgpackfiend/pkg/util"
	"github.com/spf13/cobra"
)

// tojsonCmd represents the tojson command
var tojsonCmd = &cobra.Command{
	Use:   "tojson",
	Short: "Convert a msgpack file to json",
	Long:  `Convert a msgpack file to json...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tojson called")

		infile := cmd.Flag("infile").Value.String()
		outfile := cmd.Flag("outfile").Value.String()

		// log.Printf("infile: %s", infile)
		// log.Printf("outfile: %s", outfile)

		if len(args) > 0 {
			log.Fatal("unexpected arguments: ", strings.Join(args, " "))
		}

		bytes := util.Read(infile)
		msgpackbytes := util.MsgpackToJSON(bytes)
		util.Write(msgpackbytes, outfile)
	},
}

func init() {
	rootCmd.AddCommand(tojsonCmd)
	tojsonCmd.Flags().StringP("infile", "i", "", "Input file (default stdin)")
	tojsonCmd.Flags().StringP("outfile", "o", "", "Output file (default stdout)")
}
