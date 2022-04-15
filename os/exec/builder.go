package exec

import (
	"context"
	"io"
	"os/exec"
)

type Builder struct {
	ctx  context.Context
	name string
	args []string
	dir  string
	env  []string
	path string
	out  io.Writer
	err  io.Writer
	in   io.Reader
}

func NewBuilder(name string, args ...string) *Builder {
	return NewBuilderContext(nil, name, args...)
}

func NewBuilderContext(ctx context.Context, name string, args ...string) *Builder {
	return &Builder{
		ctx:  ctx,
		name: name,
		args: args,
	}
}

func (b *Builder) Arg(arg string) *Builder {
	b.args = append(b.args, arg)
	return b
}

func (b *Builder) Args(args ...string) *Builder {
	b.args = append(b.args, args...)
	return b
}

func (b *Builder) Dir(dir string) *Builder {
	b.dir = dir
	return b
}

func (b *Builder) Env(key, value string) *Builder {
	b.env = append(b.env, key+"="+value)
	return b
}

func (b *Builder) Envs(envs map[string]string) *Builder {
	for key, value := range envs {
		b.env = append(b.env, key+"="+value)
	}
	return b
}

func (b *Builder) Path(path string) *Builder {
	b.path = path
	return b
}

func (b *Builder) Stderr(w io.Writer) *Builder {
	b.err = w
	return b
}

func (b *Builder) Stdout(w io.Writer) *Builder {
	b.out = w
	return b
}

func (b *Builder) Stdin(r io.Reader) *Builder {
	b.in = r
	return b
}

func (b *Builder) Build() *exec.Cmd {
	cmd := exec.CommandContext(b.ctx, b.name, b.args...)
	cmd.Dir = b.dir
	cmd.Env = b.env
	cmd.Path = b.path
	cmd.Stderr = b.err
	cmd.Stdin = b.in
	cmd.Stdout = b.out
	return cmd
}
