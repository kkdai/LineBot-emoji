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
	"net/http"
	"os"

	"github.com/kkdai/line-bot-sdk-go/linebot"
	"golang.org/x/text/encoding/unicode/utf32"
)

var bot *linebot.Client

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

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

//OldEmojiMsg Please note this way already been deprecated.
func OldEmojiMsg(msg *linebot.TextMessage) *linebot.TextMessage {
	return linebot.NewTextMessage(fmt.Sprintf("%s 你好 \n %s, 這是舊的傳送 Emoji 的方式。", BrownEmoji, msg.Text))
}

//NewEmojiMsg This use linebot.AddEmoji function.
func NewEmojiMsg(msg *linebot.TextMessage) linebot.SendingMessage {
	return linebot.NewTextMessage(fmt.Sprintf("$ 你好 \n %s, 這是新的傳送 Emoji 的方式。", msg.Text)).AddEmoji(linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "086"))
}

//NewEmojiMsgWithEmoji This use linebot.AddEmoji function, also parse original emoji to replace it.
func NewEmojiMsgWithEmoji(msg *linebot.TextMessage) linebot.SendingMessage {
	if len(msg.Emojis) > 0 {
		retObj := linebot.NewTextMessage(fmt.Sprintf("$ 你好 \n %s, 這是新的傳送 Emoji 的方式，如果你有 emoji 這裡會替換。", msg.Text)).AddEmoji(linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "086"))
		for _, v := range msg.Emojis {
			retObj = retObj.AddEmoji(linebot.NewEmoji(v.Index, v.ProductID, v.EmojiID))
		}
		return retObj
	}
	return linebot.NewTextMessage(fmt.Sprintf("$ 你好 \n %s, 這是新的傳送 Emoji 的方式，如果你有 emoji 這裡會替換。", msg.Text)).AddEmoji(linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "086"))
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err != nil {
					log.Println("Quota err:", err)
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, OldEmojiMsg(message), NewEmojiMsg(message), NewEmojiMsgWithEmoji(message)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
