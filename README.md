# stacktracer

Package for capture and pass around stack traces.

## Installation

`go get -u github.com/palano/stacktracer`

## Quick Start

```go

// Single Stack
stack:= stacktracer.Caller(0)

// Output String
fmt.Print(stack.String())

// Multi Stacks
stacks := stacktracer.Callers(0)

// Output String
fmt.Print(stacks.String())

```

## License
Released under the [MIT License](LICENSE).
