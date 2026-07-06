// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The plugin-template Authors

package plugin

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMessage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		cfg     Config
		want    string
		wantErr string
	}{
		{
			name: "success",
			cfg: Config{
				Version: "1.2.3",
				DryRun:  false,
			},
			want: "example plugin invoked for version 1.2.3 (dry-run: false)",
		},
		{
			name: "dry run",
			cfg: Config{
				Version: "v2.0.0",
				DryRun:  true,
			},
			want: "example plugin invoked for version v2.0.0 (dry-run: true)",
		},
		{
			name:    "missing version",
			cfg:     Config{},
			wantErr: "SEMREL_VERSION is required",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Message(tt.cfg)
			if tt.wantErr != "" {
				require.EqualError(t, err, tt.wantErr)
				require.Empty(t, got)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
