// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package list

import (
	"errors"
	"testing"

	"github.com/hashicorp/consul/agent"
	"github.com/hashicorp/consul/testrpc"
	"github.com/mitchellh/cli"

	"github.com/stretchr/testify/require"

	apply "github.com/hashicorp/consul/command/resource/apply"
)

func TestResourceListCommand(t *testing.T) {
	if testing.Short() {
		t.Skip("too slow for testing.Short")
	}

	t.Parallel()
	a := agent.NewTestAgent(t, ``)
	defer a.Shutdown()
	testrpc.WaitForTestAgent(t, a.RPC, "dc1")

	applyCli := cli.NewMockUi()

	applyCmd := apply.New(applyCli)
	code := applyCmd.Run([]string{
		"-f=../testdata/demo.hcl",
		"-http-addr=" + a.HTTPAddr(),
		"-token=root",
	})
	require.Equal(t, 0, code)
	require.Empty(t, applyCli.ErrorWriter.String())
	require.Contains(t, applyCli.OutputWriter.String(), "demo.v2.Artist 'korn' created.")

	cases := []struct {
		name      string
		output    string
		extraArgs []string
	}{
		{
			name:   "sample output",
			output: "\"name\": \"korn\"",
			extraArgs: []string{
				"demo.v2.artist",
				"-namespace=default",
				"-peer=local",
				"-partition=default",
			},
		},
		{
			name:   "file input",
			output: "\"name\": \"korn\"",
			extraArgs: []string{
				"-f=../testdata/demo.hcl",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ui := cli.NewMockUi()
			c := New(ui)

			args := []string{
				"-http-addr=" + a.HTTPAddr(),
				"-token=root",
			}

			args = append(args, tc.extraArgs...)

			actualCode := c.Run(args)
			require.Equal(t, 0, actualCode)
			require.Empty(t, ui.ErrorWriter.String())
			require.Contains(t, ui.OutputWriter.String(), tc.output)
		})
	}
}

func TestResourceListInvalidArgs(t *testing.T) {
	t.Parallel()

	a := agent.NewTestAgent(t, ``)
	defer a.Shutdown()
	testrpc.WaitForTestAgent(t, a.RPC, "dc1")

	type tc struct {
		args         []string
		expectedCode int
		expectedErr  error
	}

	cases := map[string]tc{
		"nil args": {
			args:         nil,
			expectedCode: 1,
			expectedErr:  errors.New("Please provide required arguments"),
		},
		"minimum args required": {
			args:         []string{},
			expectedCode: 1,
			expectedErr:  errors.New("Please provide required arguments"),
		},
		"no file path": {
			args: []string{
				"-f",
			},
			expectedCode: 1,
			expectedErr:  errors.New("Failed to parse args: flag needs an argument: -f"),
		},
		"file not found": {
			args: []string{
				"-f=../testdata/test.hcl",
			},
			expectedCode: 1,
			expectedErr:  errors.New("Failed to load data: Failed to read file: open ../testdata/test.hcl: no such file or directory"),
		},
		"file parsing failure": {
			args: []string{
				"-f=../testdata/invalid_type.hcl",
			},
			expectedCode: 1,
			expectedErr:  errors.New("Failed to decode resource from input file"),
		},
		"file argument is ignored": {
			args: []string{
				"demo.v2.artist",
				"-namespace=default",
				"-peer=local",
				"-partition=default",
				"-http-addr=" + a.HTTPAddr(),
				"-token=root",
				"-f=demo.hcl",
			},
			expectedCode: 0,
			expectedErr:  errors.New("File argument is ignored when resource definition is provided with the command"),
		},
		"resource type invalid": {
			args: []string{
				"test",
				"-namespace=default",
				"-peer=local",
				"-partition=default",
			},
			expectedCode: 1,
			expectedErr:  errors.New("Must include resource type argument in group.verion.kind format"),
		},
	}

	for desc, tc := range cases {
		t.Run(desc, func(t *testing.T) {
			ui := cli.NewMockUi()
			c := New(ui)

			code := c.Run(tc.args)

			require.Equal(t, tc.expectedCode, code)
			require.Contains(t, ui.ErrorWriter.String(), tc.expectedErr.Error())
		})
	}
}