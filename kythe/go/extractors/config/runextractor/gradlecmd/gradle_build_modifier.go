/*
 * Copyright 2018 The Kythe Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gradlecmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

// These are the lines necessary for gradle build to use a different javac.
const kytheJavacWrapper = `
allprojects {
  gradle.projectsEvaluated {
    tasks.withType(JavaCompile) {
      options.fork = true
      options.forkOptions.executable = '/opt/kythe/extractors/javac-wrapper.sh'
    }
  }
}
`

// This matches a line which sets the javac to use Kythe's javac-wrapper.sh
// TODO(danielmoy): consider making this not hard-coded to /opt/kythe/extractors/javac-wrapper.sh
var kytheMatcher = regexp.MustCompile(`\n\s*options\.forkOptions\.executable\ =\ '/opt/kythe/extractors/javac-wrapper.sh'`)

// This matches any line which sets a new javac executable, useful for detecting
// edge cases which already modify javac.
var javacMatcher = regexp.MustCompile(`\n\s*options\.forkOptions\.executable\ =`)

// PreProcessGradleBuild takes a gradle.build file and either verifies that it
// already has the bits necessary to run kythe's javac wrapper, or adds that
// functionality.
//
// Note this potentially modifies the input file, so make a copy beforehand if
// you need to keep the original.
func PreProcessGradleBuild(gradleBuildFile string) error {
	k, err := hasKytheWrapper(gradleBuildFile)
	if err != nil {
		return err
	}
	if k {
		// Already has the kythe javac-wrapper.
		return nil
	}
	return appendKytheWrapper(gradleBuildFile)
}

func hasKytheWrapper(gradleBuildFile string) (bool, error) {
	bits, err := ioutil.ReadFile(gradleBuildFile)
	if err != nil {
		return false, fmt.Errorf("reading file %s: %v", gradleBuildFile, err)
	}
	if kytheMatcher.Match(bits) {
		return true, nil
	}
	if javacMatcher.Match(bits) {
		return false, fmt.Errorf("found existing non-kythe javac override for file %s, which we can't handle yet", gradleBuildFile)
	}
	return false, nil
}

func appendKytheWrapper(gradleBuildFile string) error {
	f, err := os.OpenFile(gradleBuildFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("opening file %s for append: %v", gradleBuildFile, err)
	}
	if _, err := f.Write([]byte(kytheJavacWrapper)); err != nil {
		return fmt.Errorf("appending javac-wrapper to %s: %v", gradleBuildFile, err)
	}
	return f.Close()
}
