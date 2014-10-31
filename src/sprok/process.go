package sprok

import (
	"fmt"
	"log"
	"strings"
	"syscall"
)

type Process struct {
	Env  map[string]string
	Argv []string
	// The file descriptors should be directly wired to these destinations. No attempt
	// should be made to verify permissions or whatever. Fail fast.
	Stdin  string
	Stdout string
	Stderr string
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
	return fmt.Sprintf("%s %s < %s 1>%s 2>%s", env, cmd.p.Stdin, p.Stdout, ps.Stderr)
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
