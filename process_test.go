package sprok

import "testing"

func TestNewProcess(t *testing.T) {
	p := NewProcess()

	if p.Stdin != "/dev/null" {
		t.Error("stdin initialization failed")
	}
}

func TestProcessString(t *testing.T) {
	want := "TESTING=true /bin/cat process.go <a 1>b 2>c"

	p := NewProcess()
	p.Argv[0] = "/bin/cat"
	p.Argv = append(p.Argv, "process.go")
	p.Env["TESTING"] = "true"
	p.Stdin = "a"
	p.Stdout = "b"
	p.Stderr = "c"

	if p.String() != want {
		t.Error("configured process did not return the expected string")
	}
}
