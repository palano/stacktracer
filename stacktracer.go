package stacktracer

import (
	"fmt"
	"runtime"
)

const maxSize = 64

type Frame struct {
	File string
	Line int
	Name string
}

type TracerFrame struct {
	File string
	Line int
	Name string
}

type TracerFrames []TracerFrame

func (f Frame) String() string {
	return fmt.Sprintf("%s:%d %s", f.File, f.Line, f.Name)
}

func (f TracerFrame) String() string {
	return fmt.Sprintf("%s\t\n%s:%d", f.Name, f.File, f.Line)
}

func (sf TracerFrames) String() string {
	var stacks string
	for i, frame := range sf {
		if i == 0 {
			stacks = frame.String()
		} else {
			stacks = stacks + "\n" + frame.String()
		}
	}
	return stacks
}

func Caller(skip int) Frame {
	programCounter, file, line, _ := runtime.Caller(skip + 1)
	fun := runtime.FuncForPC(programCounter)
	callerFrame := Frame{
		File: file,
		Line: line,
		Name: fun.Name(),
	}
	return callerFrame
}

func Callers(skip int) TracerFrames {
	programCounters := make([]uintptr, maxSize)

	var numFrames int
	for {
		numFrames = runtime.Callers(skip+2, programCounters)
		if numFrames < len(programCounters) {
			break
		}
		programCounters = make([]uintptr, len(programCounters)*2)
	}

	var callerFrames TracerFrames
	frames := runtime.CallersFrames(programCounters[:numFrames])
	for frame, next := frames.Next(); next; frame, next = frames.Next() {
		callerFrame := TracerFrame{
			File: frame.File,
			Line: frame.Line,
			Name: frame.Function,
		}
		callerFrames = append(callerFrames, callerFrame)
	}

	return callerFrames
}
