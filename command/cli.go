package command

import (
	"io"
	"github.com/continuul/random-names/pkg/stream"
)

// Cli represents the command line client.
type Cli interface {
	Out() *stream.OutStream
	Err() io.Writer
	In() *stream.InStream
	SetIn(in *stream.InStream)
}

// CliInstance is an instance the command line client.
// Instances of the client can be returned from NewCliInstance.
type CliInstance struct {
	in  *stream.InStream
	out *stream.OutStream
	err io.Writer
}

// NewCli returns a Cli instance with IO output and error streams set by in, out and err.
func NewCliInstance(in io.ReadCloser, out, err io.Writer) *CliInstance {
	return &CliInstance{in: stream.NewInStream(in), out: stream.NewOutStream(out), err: err}
}

// Out returns the writer used for stdout
func (cli *CliInstance) Out() *stream.OutStream {
	return cli.out
}

// Err returns the writer used for stderr
func (cli *CliInstance) Err() io.Writer {
	return cli.err
}

// SetIn sets the reader used for stdin
func (cli *CliInstance) SetIn(in *stream.InStream) {
	cli.in = in
}

// In returns the reader used for stdin
func (cli *CliInstance) In() *stream.InStream {
	return cli.in
}
