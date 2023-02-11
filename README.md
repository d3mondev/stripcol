# stripcol
A command line tool that reads from stdin and strips ANSI color codes from the input.

## Description
`stripcol` is a simple command line utility that removes ANSI color codes from the input. This can be useful for cleaning up output from other programs that use ANSI codes for colorized text.

## Usage
```
echo -e '\033[31mHello\033[0m World' | stripcol
```

This command will output "Hello World" with color codes removed.

## Installation
To install nocol, you need to have [Go](https://go.dev/dl/) installed on your system.

```
go install -v github.com/d3mondev/stripcol/cmd/stripcol
```
