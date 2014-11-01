sprok
=====

Simple Proc - start processes with complex command-lines using a trivial config file

This is a package for launching processes using simple config files
rather than string concatenation and/or shell scripts.

Usage
=====

Each config format is built into a separate binary.

```sh
go get github.com/tobert/sprok/sprok-json
go get github.com/tobert/sprok/sprok-toml
go get github.com/tobert/sprok/sprok-yaml


$GOPATH/bin/sprok-json process.json
$GOPATH/bin/sprok-toml process.toml
$GOPATH/bin/sprok-yaml process.yaml
```

Examples
========

If stdout/stderr/stdin are not set, they are not modified.

```json
{
  "chdir": "/opt/cassandra",
  "env": {
    "FOO": "BAR"
  },
  "argv": [
    "/usr/bin/java",
    "-Xmx8G"
  ],
  "stdout": "/dev/null",
  "stderr": "/dev/null",
  "stdin": "/dev/null"
}
```

```yaml
chdir: /opt/cassandra
env:
  FOO: BAR
argv:
  - /usr/bin/java
  - -Xmx8G
stdin: /dev/null
stdout: /var/log/mything/stdout
stderr: /var/log/mything/stderr
```

Rationale
=========

There are many great process supervisors out there that do a nice job of forking
a process and watching it. If that's what you need, you should use one.

What sprok provides is a way to launch a process with a single command with no
shell scripts and no extra process. The process is defined in a configuration
file and no shell-style string concatenation or parsing is used. The config
data structures map directly onto the operating system's execve(3p) functionality.

TODO
====

* figure out why sprok-toml crashes
