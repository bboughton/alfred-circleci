package cli

import (
	"fmt"
	"io"
	"os"
)

type CLI struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	Handler NewHandler
}

func New(handler NewHandler) *CLI {
	return &CLI{
		Stdin:   os.Stdin,
		Stdout:  os.Stdout,
		Stderr:  os.Stderr,
		Handler: handler,
	}
}

func (c *CLI) Run(args []string) int {
	input := &Input{
		Stdin: c.Stdin,
		Args:  args,
	}
	output := &DefaultOutputWriter{
		stdout: c.Stdout,
		stderr: c.Stderr,
	}

	c.Handler.Exec(output, input)

	return output.ExitCode()
}

func Run(handler NewHandler) int {
	app := New(handler)
	return app.Run(os.Args[1:])
}

type NewHandler interface {
	Exec(OutputWriter, *Input)
}

type OutputWriter interface {
	Write([]byte) (int, error)
	WriteError([]byte) (int, error)
	ExitWith(int)
	ExitCode() int
}

type Input struct {
	Args  []string
	Stdin io.Reader
}

type DefaultOutputWriter struct {
	stdout   io.Writer
	stderr   io.Writer
	exitCode int
}

func (w *DefaultOutputWriter) ExitWith(code int) {
	if w.exitCode == 0 {
		w.exitCode = code
	}
}

func (w DefaultOutputWriter) ExitCode() int {
	return w.exitCode
}

func (w *DefaultOutputWriter) Write(data []byte) (int, error) {
	return w.stdout.Write(data)
}

func (w *DefaultOutputWriter) WriteError(data []byte) (int, error) {
	return w.stderr.Write(data)
}

type SubCommandHandler struct {
	cmds map[string]NewHandler
}

func NewSubCommandHandler() *SubCommandHandler {
	return &SubCommandHandler{
		cmds: make(map[string]NewHandler),
	}
}

func (h *SubCommandHandler) Handle(name string, handler NewHandler) {
	h.cmds[name] = handler
}

func (h SubCommandHandler) Exec(out OutputWriter, in *Input) {
	if len(in.Args) == 0 {
		fmt.Fprintln(out, "must have subcommand")
		out.ExitWith(1)
		return
	}

	name := in.Args[0]
	cmd, ok := h.cmds[name]
	if !ok {
		fmt.Fprintln(out, "command not found")
		out.ExitWith(1)
		return
	}

	input := &Input{
		Stdin: in.Stdin,
		Args:  in.Args[1:],
	}
	cmd.Exec(out, input)
}
