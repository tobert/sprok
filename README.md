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

Simple Proc - start complex processes using a simple config file

This is a package for launching processes using simple structured documents rather than string concatenation.
Sprok strives to provide a bare minimum number of features. There are plenty of other tools to choose from.
Sprok will call syscall.Exec(). Your process is expected to run in the foreground.

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
  FOO=BAR
argv:
  - /usr/bin/java
  - -Xmx8G
stdout: /run/{{ .Process.User }}/cassandra-stdout.log
stderr: /run/{{ .Process.User }}/cassandra-stderr.log
stdin: /dev/null
