// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The plugin-template Authors

package plugin

import "fmt"

// Config contains the semrel release context used by the example plugin.
type Config struct {
	Version string
	DryRun  bool
}

// Message returns the example plugin output for the provided release context.
func Message(cfg Config) (string, error) {
	if cfg.Version == "" {
		return "", fmt.Errorf("SEMREL_VERSION is required")
	}

	return fmt.Sprintf("example plugin invoked for version %s (dry-run: %t)", cfg.Version, cfg.DryRun), nil
}
