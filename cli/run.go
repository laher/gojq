package cli

import (
	"io"
	"os"
)

// Run gojq.
func Run() int {
	return (&cli{
		inStream:  os.Stdin,
		outStream: os.Stdout,
		errStream: os.Stderr,
	}).run(os.Args[1:])
}

// Run gojq with specific parameters
func RunWithParams(args []string, inStream io.Reader, outStream io.Writer, errStream io.Writer) int {
	return (&cli{
		inStream:  inStream,
		outStream: outStream,
		errStream: errStream,
	}).run(args)
}
