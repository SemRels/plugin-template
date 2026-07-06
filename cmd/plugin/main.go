// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The plugin-template Authors

package main

import (
	"fmt"
	"io"
	"os"

	plugin "github.com/SemRels/plugin-template/internal/plugin"
)

const pluginSchemaVersion = 1

func main() {
	os.Exit(run(os.Stdout, os.Stderr, os.Getenv))
}

func run(stdout, stderr io.Writer, getenv func(string) string) int {
	_, _ = fmt.Fprintf(stderr, "plugin_schema_version=%d\n", pluginSchemaVersion)

	message, err := plugin.Message(plugin.Config{
		Version: getenv("SEMREL_VERSION"),
		DryRun:  getenv("SEMREL_DRY_RUN") == "true",
	})
	if err != nil {
		fmt.Fprintln(stderr, "plugin-template:", err)
		return 1
	}

	fmt.Fprintln(stdout, message)
	return 0
}
