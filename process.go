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
	"os/exec"
	"strings"
	"syscall"
)

// Process is the configuration for running a process
type Process struct {
	Chdir  string            `json:"chdir"  yaml:"chdir"`  // cd /tmp
	Env    map[string]string `json:"env"    yaml:"env"`    // env FOO=BAR
	Argv   []string          `json:"argv"   yaml:"argv"`   // "/bin/dd" "if=/dev/zero" "count=10"
	Stdin  string            `json:"stdin"  yaml:"stdin"`  // <"/dev/null"
	Stdout string            `json:"stdout" yaml:"stdout"` // >"zero.bin"
	Stderr string            `json:"stderr" yaml:"stderr"` // 2>"errors.log"
}

// NewProcess returns a Process struct with the env map and argv allocated
// and all stdio pointed at /dev/null.
// argv is allocated with a length of 1 to hold the command.
// Use append to add arguments.
func NewProcess() Process {
	return Process{
		Env:    map[string]string{},
		Argv:   []string{},
		Chdir:  "/",
		Stdin:  "",
		Stdout: "",
		Stderr: "",
	}
}

// Exec executes Argv with environment Env and file descriptors
// 1, 2, and 3 open on the files specified in Stdin, Stdout,
// and Stderr. When output files are unspecified or an empty
// string, the file descriptors are left unmodified.
// If argv[0] is stat-able (absolute or relative path), it is used as-is.
// When that fails the PATH searched using exec.LookPath().
func (p *Process) Exec() error {
	var stdin, stdout, stderr *os.File
	var err error

	// Always chdir before doing anything else, making relative paths
	// relative to the provided directory.
	err = os.Chdir(p.Chdir)
	if err != nil {
		log.Fatalf("Could not chdir to '%s': %s\n", p.Chdir, err)
	}

	// Check if it's relative to Chdir or an absolute path, either
	// way it will stat and return not-nil.
	fi, err := os.Stat(p.Argv[0])
	if err != nil {
		fpath, err := exec.LookPath(p.Argv[0])
		if err != nil {
			log.Fatalf("'%s' could not be found in PATH: %s\n", p.Argv[0], err)
		}
		p.Argv[0] = fpath
	}

	// Make sure argv[0] is an actual file before proceeding.
	fi, err = os.Stat(p.Argv[0])
	if err != nil {
		log.Fatalf("BUG: '%s' is not a valid path to a file: %s\n", p.Argv[0], err)
	}
	m := fi.Mode()
	if !m.IsRegular() {
		log.Fatalf("'%s' is not a file!\n", p.Argv[0])
	}

	// If stdin is set, remap it to the file specified on an fd opened read-only.
	if p.Stdin != "" {
		stdin, err = os.OpenFile(p.Stdin, os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("Could not open stdin target '%s': %s\n", p.Stdin, err)
		}

		err = syscall.Dup2(int(stdin.Fd()), int(os.Stdin.Fd()))
		if err != nil {
			log.Fatalf("Failed to redirect stdin: %s\n", err)
		}
	}

	// If stdout is set, remap it to the file specified on an fd opened
	// for append and write-only, it will be created if it does not exist.
	if p.Stdout != "" {
		stdout, err = os.OpenFile(p.Stdout, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("Could not open stdout target '%s': %s\n", p.Stdout, err)
		}

		err = syscall.Dup2(int(stdout.Fd()), int(os.Stdout.Fd()))
		if err != nil {
			log.Fatalf("Failed to redirect stdout: %s\n", err)
		}
	}

	// Same deal as stdout except if stdin and stdout have the same target, in
	// which case they will share the fd like 2>&1 does.
	if p.Stderr != "" {
		// there is no reason to open the file twice if they're the same file
		if p.Stderr == p.Stdout {
			stderr = stdout // will get dup2'd to the same fd
		} else {
			stderr, err = os.OpenFile(p.Stderr, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				log.Fatalf("Could not open stderr target '%s': %s\n", p.Stderr, err)
			}
		}

		err = syscall.Dup2(int(stderr.Fd()), int(os.Stderr.Fd()))
		if err != nil {
			log.Fatalf("Failed to redirect stderr: %s\n", err)
		}
	}

	return syscall.Exec(p.Argv[0], p.Argv, p.envPairs())
}

// String returns the process settings as a Bourne shell command.
func (p *Process) String() string {
	env := strings.Join(p.envPairs(), " ")
	cmd := strings.Join(p.Argv, " ")
	// cd / && env FOO=BAR cmd -arg1 -arg2 foo < /dev/null 1>/dev/null 2>/dev/null
	return fmt.Sprintf("cd %s && env %s %s <%s 1>%s 2>%s",
		p.Chdir, env, cmd, p.Stdin, p.Stdout, p.Stderr)
}

// envPairs converts the key:value map into an array of key=val which
// is what execve(3P) uses.
func (p *Process) envPairs() []string {
	env := make([]string, len(p.Env))
	i := 0
	for key, value := range p.Env {
		env[i] = fmt.Sprintf("%s=%s", key, value)
		i++
	}
	return env
}
