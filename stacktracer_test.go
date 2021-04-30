package stacktracer

import (
	"path"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

func caller1(skip int) Frame {
	frame := Caller(skip)
	frame.File = cleanWorkingDir(frame.File)
	return frame
}

func caller2(skip int) Frame {
	return caller1(skip)
}

func callers1(skip int) TracerFrames {
	frames := Callers(skip)
	for i, frame := range frames {
		frames[i].File = cleanWorkingDir(frame.File)
	}
	return frames
}

func callers2(skip int) TracerFrames {
	return callers1(skip)
}

func TestTakeStacktracerCaller1(t *testing.T) {
	frame := caller1(0)

	expected := "stacktracer_test.go:12 github.com/petrosea/stacktracer.caller1"
	result := frame.String()

	if !regexp.MustCompile(expected).MatchString(result) {
		t.Fatalf("Not match %s with %s", expected, result)
	}
}

func TestTakeStacktracerCaller2(t *testing.T) {
	frame := caller2(0)

	expected := "stacktracer_test.go:12 github.com/petrosea/stacktracer.caller1"
	result := frame.String()

	if !regexp.MustCompile(expected).MatchString(result) {
		t.Fatalf("Not match %s with %s", expected, result)
	}
}

func TestTakeStacktracerCallerSkip1(t *testing.T) {
	frame := caller2(1)

	expected := "stacktracer_test.go:18 github.com/petrosea/stacktracer.caller2"
	result := frame.String()

	if !regexp.MustCompile(expected).MatchString(result) {
		t.Fatalf("Not match %s with %s", expected, result)
	}
}

func TestTakeStacktracerCallerSkip2(t *testing.T) {
	frame := caller2(2)

	expected := "stacktracer_test.go:67 github.com/petrosea/stacktracer.TestTakeStacktracerCallerSkip2"
	result := frame.String()

	if !regexp.MustCompile(expected).MatchString(result) {
		t.Fatalf("Not match %s with %s", expected, result)
	}
}

func TestTakeStacktracerCallers(t *testing.T) {
	frames := callers2(0)

	expected := []string{
		"github.com/petrosea/stacktracer.callers1\t\nstacktracer_test.go:22",
		"github.com/petrosea/stacktracer.callers2\t\nstacktracer_test.go:30",
		"github.com/petrosea/stacktracer.TestTakeStacktracerCallers\t\nstacktracer_test.go:78",
	}

	for i, expect := range expected {
		result := frames[i].String()
		if !regexp.MustCompile(expect).MatchString(result) {
			t.Fatalf("Not match %s with %s", expect, result)
		}
	}
}

func TestTakeStacktracerCallersSkip1(t *testing.T) {
	frames := callers2(1)

	expected := []string{
		"github.com/petrosea/stacktracer.callers2\t\nstacktracer_test.go:30",
		"github.com/petrosea/stacktracer.TestTakeStacktracerCallersSkip1\t\nstacktracer_test.go:95",
	}

	for i, expect := range expected {
		result := frames[i].String()
		if !regexp.MustCompile(expect).MatchString(result) {
			t.Fatalf("Not match %s with %s", expect, result)
		}
	}
}

func TestTakeStacktracerCallersSkip2(t *testing.T) {
	frames := callers2(2)

	expected := []string{
		"github.com/petrosea/stacktracer.TestTakeStacktracerCallersSkip2\t\nstacktracer_test.go:111",
	}

	for i, expect := range expected {
		result := frames[i].String()
		if !regexp.MustCompile(expect).MatchString(result) {
			t.Fatalf("Not match %s with %s", expect, result)
		}
	}
}

func BenchmarkTakeStacktracerCaller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Caller(0)
	}
}

func BenchmarkTakeStacktracerCallers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Callers(0)
	}
}

func cleanWorkingDir(file string) string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Dir(filename)
	return strings.ReplaceAll(file, dir+"/", "")
}
