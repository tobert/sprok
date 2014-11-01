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

import (
	"fmt"
	"strings"
	"syscall"
)

type Process struct {
	Env    map[string]string `"env"`
	Argv   []string          `"argv"`
	Stdin  string            `"stdin"`
	Stdout string            `"stdout"`
	Stderr string            `"stderr"`
}

// NewProcess returns a Process struct with the env map and argv allocated
// and all stdio pointed at /dev/null.
func NewProcess() Process {
	return Process{
		Env:    make(map[string]string),
		Argv:   make([]string, 1),
		Stdin:  "/dev/null",
		Stdout: "/dev/null",
		Stderr: "/dev/null",
	}
}

// Exec executes Argv with environment Env and file descriptors 1, 2, and 3 open on the
// files specified in Stdin, Stdout, and Stderr. They all default to /dev/null if left unspecified
// in the config or for empty string "".
func (p *Process) Exec() error {
	return syscall.Exec(p.Argv[0], p.Argv[1:], p.envPairs())
}

func (p *Process) String() string {
	env := strings.Join(p.envPairs(), " ")
	cmd := strings.Join(p.Argv, " ")
	// FOO=BAR cmd -arg1 -arg2 foo < /dev/null 1>/dev/null 2>/dev/null
	return fmt.Sprintf("%s %s <%s 1>%s 2>%s", env, cmd, p.Stdin, p.Stdout, p.Stderr)
}

func (p *Process) envPairs() []string {
	env := make([]string, len(p.Env))
	i := 0
	for key, value := range p.Env {
		env[i] = fmt.Sprintf("%s=%s", key, value)
		i++
	}
	return env
}
