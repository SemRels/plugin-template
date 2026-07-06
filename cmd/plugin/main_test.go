// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The plugin-template Authors

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		env        map[string]string
		wantCode   int
		wantStdout string
		wantStderr string
	}{
		{
			name: "success",
			env: map[string]string{
				"SEMREL_VERSION": "1.2.3",
			},
			wantCode:   0,
			wantStdout: "example plugin invoked for version 1.2.3 (dry-run: false)\n",
			wantStderr: "plugin_schema_version=1\n",
		},
		{
			name: "dry run",
			env: map[string]string{
				"SEMREL_VERSION": "1.2.3",
				"SEMREL_DRY_RUN": "true",
			},
			wantCode:   0,
			wantStdout: "example plugin invoked for version 1.2.3 (dry-run: true)\n",
			wantStderr: "plugin_schema_version=1\n",
		},
		{
			name:       "missing version",
			env:        map[string]string{},
			wantCode:   1,
			wantStdout: "",
			wantStderr: "plugin_schema_version=1\nplugin-template: SEMREL_VERSION is required\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var stdout bytes.Buffer
			var stderr bytes.Buffer

			code := run(&stdout, &stderr, envMap(tt.env))

			require.Equal(t, tt.wantCode, code)
			require.Equal(t, tt.wantStdout, stdout.String())
			require.Equal(t, tt.wantStderr, stderr.String())
			require.True(t, strings.HasPrefix(stderr.String(), "plugin_schema_version=1\n"))
		})
	}
}

func envMap(values map[string]string) func(string) string {
	return func(key string) string {
		return values[key]
	}
}
