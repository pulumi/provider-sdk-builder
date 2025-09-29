package shell

import "strings"

type ShellCommandSequence struct {
	Commands []string
}

func NewShellCommandSequence() ShellCommandSequence {
	return ShellCommandSequence{
		Commands: []string{},
	}
}

func (s *ShellCommandSequence) Append(cmd string) {
	s.Commands = append(s.Commands, cmd)
}

func (s *ShellCommandSequence) AppendAll(seq ShellCommandSequence) {
	s.Commands = append(s.Commands, seq.Commands...)
}

func (scs *ShellCommandSequence) String() string {
	return strings.Join(scs.Commands, "\n")
}
