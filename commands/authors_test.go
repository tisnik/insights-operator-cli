/*
Copyright © 2019, 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands_test

import (
	"github.com/redhatinsighs/insights-operator-cli/commands"
	"strings"
	"testing"
)

// TestCommandAuthors check if the command 'authors' displays list of authors
func TestCommandAuthors(t *testing.T) {
	configureColorizer()
	captured, err := captureStandardOutput(func() {
		commands.PrintAuthors()
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured == "" {
		t.Fatal("Standard output is empty")
	}
	if !strings.HasPrefix(captured, "Authors") {
		t.Fatal("Unexpected output:\n", captured)
	}
	numlines := strings.Count(captured, "\n")
	if numlines <= 1 {
		t.Fatal("Authors are not listed in the output:\n", captured)
	}
}
