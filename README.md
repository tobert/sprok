NOTHING TO SEE HERE MOVE ALONG
==============================

Just an idea. It's pretty trivial so maybe pointless. Stream of consciousness-ish notes follow:

I plan to use this in my Docker containers to start Cassandra with a precise command line instead of
with the shell scripts. It's not a huge difference but it removes some complexity in providing a
consistent way for users to make changes to how Cassandra starts in a container. Because I recommend
always using a volume for Cassandra, I create a new state directory on that volume and store all of
Cassandra's startup configuration there, including cassandra.yaml and a couple env scripts. The way
users set heap memory is by editing a procedural shell script, which usually works out OK but it
does get fouled up from time to time.

This feels a lot more precise and removes a bunch of complexity, so let's see what happens.

sprok
=====

Simple Proc - start complex processes using a config file

This is a package for launching processes using simple structured documents rather than string concatenation.
Sprok strives to provide a bare minimum number of features. There are plenty of other tools to choose from.

```
sprok_json CassandraDaemon.json
sprok_toml CassandraDaemon.toml
sprok_yaml CassandraDaemon.yaml
```

```json
{ "env": { "FOO": "BAR" }, "argv": [ "/usr/bin/java", "-Xmx8G" ], "stdout": "/dev/null", "stderr": "/dev/null", "stdin": "/dev/null" }
```

```yaml
env:
  FOO=BAR
argv:
  - /usr/bin/java
  - -Xmx8G
stdout: /run/{{ .Process.User }}/cassandra-stdout.log
stderr: /run/{{ .Process.User }}/cassandra-stderr.log
stdin: /dev/null
```

```go
package sprok

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
```
