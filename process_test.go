package sprok

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

import "testing"

func TestNewProcess(t *testing.T) {
	p := NewProcess()

	if p.Stdin != "/dev/null" {
		t.Error("stdin initialization failed")
	}
}

// This needs to be kept in sync with the format in
// process.go, which shouldn't change much ...
func TestProcessString(t *testing.T) {
	want := "cd /tmp && env TESTING=true ANOTHER=yesplease /bin/cat process.go <a 1>b 2>c"

	p := NewProcess()
	p.Chdir = "/tmp"
	p.Argv = []string{"/bin/cat", "process.go"}
	p.Env["TESTING"] = "true"
	p.Env["ANOTHER"] = "yesplease"
	p.Stdin = "a"
	p.Stdout = "b"
	p.Stderr = "c"

	if p.String() != want {
		t.Error("configured process did not return the expected string")
	}
}
