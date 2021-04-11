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

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func TestSingleReplaceEmoji(t *testing.T) {
	test0 := "hello"
	want0 := "hello"
	emojis0 := []*linebot.Emoji{}
	if ret := ReplaceEmoji(test0, emojis0); !strings.EqualFold(want0, ret) {
		t.Errorf("[[zero]] Replaced failed, %s", ret)
	}

	test01 := "(telescope)"
	want01 := "$"
	emojis := []*linebot.Emoji{
		&linebot.Emoji{
			Index:  0,
			Length: 11}}

	if ret := ReplaceEmoji(test01, emojis); !strings.EqualFold(want01, ret) {
		t.Errorf("[[single emopji without string]] Replaced failed, %s", ret)
	}

	test1 := "(telescope) hello"
	want1 := "$ hello"
	if ret := ReplaceEmoji(test1, emojis); !strings.EqualFold(want1, ret) {
		t.Errorf("[[single]] Replaced failed, %s", ret)
	}

	test2 := "aa (telescope)"
	want2 := "aa $"
	emojis2 := []*linebot.Emoji{
		&linebot.Emoji{
			Index:  3,
			Length: 11}}
	if ret := ReplaceEmoji(test2, emojis2); !strings.EqualFold(want2, ret) {
		t.Errorf("[[single]] Replaced failed is not begin, %s", ret)
	}

	test3 := "aa (telescope) aabb"
	want3 := "aa $ aabb"
	if ret := ReplaceEmoji(test3, emojis2); !strings.EqualFold(want3, ret) {
		t.Errorf("[[single]] Replaced failed is not begin with string, %s", ret)
	}
}

func TestMultipleReplaceEmoji(t *testing.T) {
	test1 := "(12)(12)"
	want1 := "$$"
	emojis1 := []*linebot.Emoji{
		&linebot.Emoji{
			Index:  0,
			Length: 4},
		&linebot.Emoji{
			Index:  4,
			Length: 4}}

	if ret := ReplaceEmoji(test1, emojis1); !strings.EqualFold(want1, ret) {
		t.Errorf("[[double]] Replaced failed, %s", ret)
	}

	test2 := " (12) (12)"
	want2 := " $ $"
	emojis2 := []*linebot.Emoji{
		&linebot.Emoji{
			Index:  1,
			Length: 4},
		&linebot.Emoji{
			Index:  6,
			Length: 4}}

	if ret := ReplaceEmoji(test2, emojis2); !strings.EqualFold(want2, ret) {
		t.Errorf("[[double]] Replaced failed with space, %s", ret)
	}

	test3 := " (12) (12) (12)"
	want3 := " $ $ $"
	emojis3 := []*linebot.Emoji{
		&linebot.Emoji{
			Index:  1,
			Length: 4},
		&linebot.Emoji{
			Index:  6,
			Length: 4},
		&linebot.Emoji{
			Index:  11,
			Length: 4}}

	if ret := ReplaceEmoji(test3, emojis3); !strings.EqualFold(want3, ret) {
		t.Errorf("[[Third]] Replaced failed with space, %s", ret)
	}
}
func TestCheckProdEmojiID(t *testing.T) {
	//It includes in standard list
	prod, emoji := CheckProdEmojiID("5ac2173d031a6752fb806d56", "001")
	if !strings.EqualFold(prod, "5ac2173d031a6752fb806d56") || !strings.EqualFold("001", emoji) {
		t.Error("Standard list checking failed \n")
	}

	//Missing case which is not in standard list, reply with brown
	prod, emoji = CheckProdEmojiID("5ac309f0031a6752fb806d8d", "002")
	if !strings.EqualFold(prod, "5ac1bfd5040ab15980c9b435") || !strings.EqualFold("086", emoji) {
		t.Error("Missing emoji should be replaced by brown.")
	}
}
