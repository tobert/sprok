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

examples
========

```sh
sprok_json CassandraDaemon.json
sprok_toml CassandraDaemon.toml
sprok_yaml CassandraDaemon.yaml
```

Why support 3 file formats? Because it makes sense to use whatever config format you
already use in your project. For Apache Cassandra, that's yaml. Some other projects
use JSON and they should use JSON.

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
