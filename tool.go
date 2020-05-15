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
	"fmt"
	"log"
	"strings"

	"github.com/kkdai/line-bot-sdk-go/linebot"
	"golang.org/x/text/encoding/unicode/utf32"
)

var BrownEmoji string

func init() {
	var err error
	utf32BEIB := utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM)
	dec := utf32BEIB.NewDecoder()
	BrownEmoji, err = dec.String("\x00\x10\x00\x84")
	if err != nil {
		log.Print(err)
	}
}

//OldEmojiMsg Please note this way already been deprecated.
func OldEmojiMsg(msg *linebot.TextMessage) *linebot.TextMessage {
	return linebot.NewTextMessage(fmt.Sprintf("%s%s 你好 \n , 這是舊的傳送 Emoji 的方式。", BrownEmoji, msg.Text))
}

//NewEmojiMsg This use linebot.AddEmoji function.
func NewEmojiMsg(msg *linebot.TextMessage) linebot.SendingMessage {
	return linebot.NewTextMessage(fmt.Sprintf("$%s 你好 \n , 這是新的傳送 Emoji 的方式。", msg.Text)).AddEmoji(linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "086"))
}

//ReplaceEmoji Replace original emoji `(brown)` to `$`
func ReplaceEmoji(oriMsg string, emojis []*linebot.Emoji) string {
	//Process correct echo message.
	var prefix int
	var lastLength int
	if len(emojis) > 0 {
		prefix = emojis[0].Index
	}

	workMsg := oriMsg
	log.Println("OriMsg:", oriMsg)
	for k, v := range emojis {
		log.Println("Got each detail emoji:", v, " text:", workMsg)
		msgArray := []byte(workMsg)
		index := v.Index - lastLength + k
		workMsg = fmt.Sprintf("%s%s%s", string(msgArray[:index]), "$", string(msgArray[index+v.Length:]))
		log.Println("BB Work msg:", workMsg, " index:", v.Index, " prefix:", prefix, "Length:", prefix+v.Length)
		lastLength = lastLength + v.Length
		log.Println("FF Work msg:", workMsg, " index:", v.Index, " prefix:", prefix, "Length:", prefix+v.Length)
	}

	return workMsg
}

//NewEmojiMsgWithEmoji This use linebot.AddEmoji function, also parse original emoji to replace it.
func NewEmojiMsgWithEmoji(msg *linebot.TextMessage) linebot.SendingMessage {
	if len(msg.Emojis) > 0 {
		//Replace original emoji `(brown)` to `$`
		workMsg := ReplaceEmoji(msg.Text, msg.Emojis)

		log.Println("Got all detail emoji:", msg.Emojis)
		retObj := linebot.NewTextMessage(fmt.Sprintf("$%s 你好 \n , 這是新的傳送 Emoji 的方式，如果你有 emoji 這裡會替換。", workMsg)).AddEmoji(linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "086"))

		for _, v := range msg.Emojis {
			log.Println("Got emoji detail:", v)
			prodID, emojiID := CheckProdEmojiID(v.ProductID, v.EmojiID)
			retObj = retObj.AddEmoji(linebot.NewEmoji(1+v.Index, prodID, emojiID))
		}
		return retObj
	}
	return linebot.NewTextMessage(fmt.Sprintf("$ 你好 \n %s, 這是新的傳送 Emoji 的方式，如果你有 emoji 這裡會替換。", msg.Text)).AddEmoji(linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "086"))
}

//CheckProdEmojiID :Return an valid product and emoji ID.
//If Emoji product ID is not standard one, it will be replaced by standard brown emoji.
func CheckProdEmojiID(proID, emojiID string) (string, string) {
	//All standard emoji product list from https://d.line-scdn.net/r/devcenter/sendable_line_emoji_list.pdf
	emojiProductIDs := [...]string{"5ac1bfd5040ab15980c9b435",
		"5ac1de17040ab15980c9b438",
		"5ac21184040ab15980c9b43a",
		"5ac21542031a6752fb806d55",
		"5ac2173d031a6752fb806d56",
		"5ac21869040ab15980c9b43b",
		"5ac218e3040ab15980c9b43c",
		"5ac2197b040ab15980c9b43d",
		"5ac21a13031a6752fb806d57",
		"5ac21a18040ab15980c9b43e",
		"5ac21a8c040ab15980c9b43f",
		"5ac21ae3040ab15980c9b440",
		"5ac21b4f031a6752fb806d59",
		"5ac21ba5040ab15980c9b441",
		"5ac21bf9031a6752fb806d5a",
		"5ac21c46040ab15980c9b442",
		"5ac21c4e031a6752fb806d5b",
		"5ac21cc5031a6752fb806d5c",
		"5ac21cce040ab15980c9b443",
		"5ac21d59031a6752fb806d5d",
		"5ac21e6c040ab15980c9b444",
		"5ac21ef5031a6752fb806d5e",
		"5ac21f52040ab15980c9b445",
		"5ac21fda040ab15980c9b446",
		"5ac2206d031a6752fb806d5f",
		"5ac220c2031a6752fb806d60",
		"5ac2211e031a6752fb806d61",
		"5ac2213e040ab15980c9b447",
		"5ac2216f040ab15980c9b448",
		"5ac221ca040ab15980c9b449",
		"5ac22224031a6752fb806d62",
		"5ac22293031a6752fb806d63",
		"5ac222bf031a6752fb806d64",
		"5ac223c6040ab15980c9b44a",
		"5ac2264e040ab15980c9b44b",
		"5ac22775040ab15980c9b44c",
		"5ac2280f031a6752fb806d65",
		"5ac22a8c031a6752fb806d66",
		"5ac22b23040ab15980c9b44d",
		"5ac22bad031a6752fb806d67",
		"5ac22c9e031a6752fb806d68",
		"5ac22d62031a6752fb806d69",
		"5ac22def040ab15980c9b44e",
		"5ac22e85040ab15980c9b44f"}

	for _, v := range emojiProductIDs {
		if strings.EqualFold(v, proID) {
			return proID, emojiID
		}
	}
	return "5ac1bfd5040ab15980c9b435", "086"
}
