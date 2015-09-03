package command

import (
	"github.com/mitchellh/cli"
	"strings"
	"testing"
)

func TestSubCommand_implements(t *testing.T) {
	var _ cli.Command = &SubCommand{}
}

func TestSubCommandRun(t *testing.T) {

	ui := new(cli.MockUi)
	c := &SubCommand{UI: ui}
	args := []string{"key=val"}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if !strings.Contains(ui.OutputWriter.String(), "Subcommand Complete") {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}
