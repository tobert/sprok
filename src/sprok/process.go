package sprok

import (
    "log"
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

// Exec executes Argv with environment Env and file descriptors 1, 2, and 3 open on the
// files specified in Stdin, Stdout, and Stderr. They all default to /dev/null if left unspecified
// in the config or for empty string "".
func (p *Process) Exec() error {
    return os.Exec( ... )
}

// String returns a stringified copy of the command in `ENV=VAL exe -args` form.
// No quote modification is made; it is a simple concatenation of whatever is in Env and Argv.
func (p *Process) String() string {
}
