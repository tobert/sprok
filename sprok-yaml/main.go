package main

/*
 * Copyright 2014 Albert P. Tobey <atobey@datastax.com> @AlTobey
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/tobert/sprok"
	"gopkg.in/yaml.v2"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("this program requires exactly one argument")
	}

	js, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatalf("Error opening '%s' for read: %s\n", args[0], err)
	}

	proc := sprok.NewProcess()

	err = yaml.Unmarshal(js, &proc)
	if err != nil {
		log.Fatalf("Could not parse YAML data in file '%s': %s\n", args[0], err)
	}

	proc.Exec()
}
