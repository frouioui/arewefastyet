/*
 *
 * Copyright 2021 The Vitess Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * /
 */

package git

import (
	qt "github.com/frankban/quicktest"
	"testing"
)

func TestShortenSHA(t *testing.T) {
	tests := []struct {
		name string
		sha string
		want string
	}{
		{name: "Regular SHA", sha: "5a504473aec6176b2523bf935ffe4217f61e9928", want: "5a50447"},
		{name: "Short SHA", sha: "5a504473", want: "5a50447"},
		{name: "Tiny SHA", sha: "5a50", want: "5a50"},
		{name: "Empty", sha: "", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := qt.New(t)
			gotSHA := ShortenSHA(tt.sha)
			c.Assert(gotSHA, qt.Equals, tt.want)
		})
	}
}
