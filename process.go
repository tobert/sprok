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
	"log"
	"os"
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
		Env:    map[string]string{},
		Argv:   make([]string, 1),
		Stdin:  "/dev/null",
		Stdout: "/dev/null",
		Stderr: "/dev/null",
	}
}

// Exec executes Argv with environment Env and file descriptors
// 1, 2, and 3 open on the files specified in Stdin, Stdout,
// and Stderr. When output files are unspecified or an empty
// string, the file descriptors are left unmodified.
func (p *Process) Exec() error {
	var stdin, stdout, stderr *os.File

	if p.Stdin != "" {
		stdin = openFile(p.Stdin, os.O_RDONLY)

		err := syscall.Dup2(int(stdin.Fd()), int(os.Stdin.Fd()))
		if err != nil {
			log.Fatalf("Failed to redirect stdin: %s\n", err)
		}
	}

	if p.Stdout != "" {
		stdout = openFile(p.Stdout, os.O_WRONLY|os.O_CREATE|os.O_APPEND)

		err := syscall.Dup2(int(stdout.Fd()), int(os.Stdout.Fd()))
		if err != nil {
			log.Fatalf("Failed to redirect stdout: %s\n", err)
		}
	}

	dupStderr := false
	if p.Stderr != "" && p.Stderr == p.Stdout {
		stderr = stdout
		dupStderr = true
	} else if p.Stderr != "" {
		stderr = openFile(p.Stderr, os.O_WRONLY|os.O_CREATE|os.O_APPEND)
		dupStderr = true
	}

	if dupStderr {
		err := syscall.Dup2(int(stderr.Fd()), int(os.Stderr.Fd()))
		if err != nil {
			log.Fatalf("Failed to redirect stderr: %s\n", err)
		}
	}

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

func openFile(name string, flags int) *os.File {
	if name == "/dev/null" {
		name = os.DevNull
	}

	f, err := os.OpenFile(name, flags, 0644)
	if err != nil {
		log.Fatalf("Could not open file '%s': %s\n", name, err)
	}

	return f
}
