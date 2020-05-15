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

package main

import (
	"strings"
	"testing"
)

func TestReplaceEmoji(t *testing.T) {

}

func TestCheckProdEmojiID(t *testing.T) {
	prod, emoji := CheckProdEmojiID("5ac2173d031a6752fb806d56", "001")
	if !strings.EqualFold(prod, "5ac2173d031a6752fb806d56") || !strings.EqualFold("001", emoji) {
		t.Error("Standard list checking failed \n")
	}

	prod, emoji = CheckProdEmojiID("5ac309f0031a6752fb806d8d", "002")
	if !strings.EqualFold(prod, "5ac1bfd5040ab15980c9b435") || !strings.EqualFold("086", emoji) {
		t.Error("Missing emoji should be replaced by brown.")
	}

}
