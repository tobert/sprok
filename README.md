sprok
=====

Simple Proc - start processes with complex command-lines using a simple config file

This is a package for launching processes using simple config files
rather than string concatenation and/or shell scripts.

why?
====

There are lots of great process supervisors available. This is not
a process manager and never will be.  Sprok exists to provide exactly
one feature: defining a process using a simple config file that maps
directly onto how the OS launches a process rather than relying on
string parsing (e.g. sh -c).

Why support 3 file formats? Because it makes sense to use whatever config format you
already use in your project. For example, when I'm working with Apache Cassandra,
it uses YAML so I want to specify the process in YAML as well.

Usage
=====

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

```json
{
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
env:
  FOO: BAR
argv:
  - /usr/bin/java
  - -Xmx8G
stdin: /dev/stdin
stdout: /dev/stdout
stderr: /dev/stderr
```

TODO
====

* implement stdio fd changes
