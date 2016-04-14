package commands_test

import (
	"bytes"
	"testing"

	"github.com/bboughton/alfred-circleci/commands"
)

func TestFilterCommandRuns(t *testing.T) {
	// given
	var (
		actual []byte
		buf    bytes.Buffer
	)
	expected := []byte(`foo`)
	cmd := commands.Filter{}
	req := commands.Input{}
	res := commands.Output{
		Stdout: &buf,
	}

	// when
	cmd.Run(&req, &res)
	actual = buf.Bytes()

	// then
	if !bytes.Equal(actual, expected) {
		t.Fail()
	}
}
