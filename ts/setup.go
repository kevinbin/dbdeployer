// DBDeployer - The MySQL Sandbox
// Copyright © 2006-2022 Giuseppe Maxia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ts

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/datacharmer/dbdeployer/common"
	"github.com/rogpeppe/go-internal/testscript"
)

func dbdeployerSetup(t *testing.T, dir string) func(env *testscript.Env) error {
	return func(env *testscript.Env) error {
		readFile := func(fileName string) (string, error) {
			wantedFile := path.Join(dir, fileName)
			if !common.FileExists(wantedFile) {
				return "", fmt.Errorf("no %s file found in %s", fileName, dir)
			}
			text, err := os.ReadFile(wantedFile) // #nosec G304
			if err != nil {
				return "", fmt.Errorf("error reading file %s: %s", wantedFile, err)
			}
			if len(text) == 0 {
				return "", fmt.Errorf("file %s was empty", wantedFile)
			}
			return string(text), nil
		}
		versionText, err := readFile("DB_VERSION")
		if err != nil {
			return fmt.Errorf("error reading version file in %s", dir)
		}
		flavorText, err := readFile("DB_FLAVOR")
		if err != nil {
			return fmt.Errorf("error reading flavor file in %s", dir)
		}
		env.Setenv("db_version", versionText)
		env.Setenv("db_flavor", flavorText)

		env.Values["db_version"] = versionText
		env.Values["db_flavor"] = flavorText
		env.Values["testingT"] = t

		return nil
	}
}
