// Copyright 2015 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transform

import (
	"bytes"
	"github.com/spf13/hugo/helpers"
	"testing"
)

func TestLiveReloadInject(t *testing.T) {
	out := new(bytes.Buffer)
	in := helpers.StringToReader("</body>")

	tr := NewChain(LiveReloadInject)
	tr.Apply(out, in, []byte("path"))

	expected := `<script data-no-instant>document.write('<script src="/livereload.js?mindelay=10"></' + 'script>')</script></body>`
	if string(out.Bytes()) != expected {
		t.Errorf("Expected %s got %s", expected, string(out.Bytes()))
	}
}