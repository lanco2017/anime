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
// https://www.evanlin.com/create-your-line-bot-golang/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func real_num(text string) string {
	text = strings.Replace(text, "１", "1", -1)
	text = strings.Replace(text, "２", "2", -1)
	text = strings.Replace(text, "３", "3", -1)
	text = strings.Replace(text, "４", "4", -1)
	text = strings.Replace(text, "５", "5", -1)
	text = strings.Replace(text, "６", "6", -1)
	text = strings.Replace(text, "７", "7", -1)
	text = strings.Replace(text, "８", "8", -1)
	text = strings.Replace(text, "９", "9", -1)
	text = strings.Replace(text, "０", "0", -1)
	return text
}

func anime(text string,user_msgid string) string {
	print_string := text
	text = real_num(text)
//	reg := regexp.MustCompile(`^.*(動畫|動畫瘋|巴哈姆特|anime|アニメ).*(這個美術社大有問題|美術社)\D*(\d{1,})`) //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	reg := regexp.MustCompile("^.*(動畫|動畫瘋|巴哈姆特|anime|アニメ)(\\s|　|:|;|：|；)([\u4e00-\u9fa5_a-zA-Z0-9]*)\\D*(\\d{1,})") //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	switch reg.ReplaceAllString(text, "$1"){
	case "開發者":
		print_string = "你找我主人？OK！\n我跟你講我的夥伴喵在哪，你去加他。\n他跟主人很親近的，跟他說的話主人都會看到。\nhttps://line.me/R/ti/p/%40uwk0684z\n\n\n你也可以從下面這個連結直接跟他線上對話。\n\n如果他不在線上一樣可以留言給他，\n他會收到的！\n這有跟手機連線會自動同步。" +
		"\n\nhttp://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + user_msgid +
		"&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A"
	case "動畫", "動畫瘋", "巴哈姆特", "anime", "アニメ":
		print_string = text + "？\n好像有這個動畫耶，但我找不太到詳細的QQ\n你要手動去「巴哈姆特動畫瘋」找找嗎？\n\nhttps://ani.gamer.com.tw"
		anime_say := "有喔！有喔！你在找這個對吧！？\n"
		switch reg.ReplaceAllString(text, "$3") {
		case "美術社","這個美術社大有問題":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=5871"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5918"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5919"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6038"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6039"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6040"
			case "7":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=6207"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6208"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6295"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6296"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6297"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6298" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "歌之☆王子殿下♪ 真愛","歌王子","uta":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6436" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2068\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2055\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2042"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6470" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2069\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2056\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2043"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6496" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2070\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2057\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2044"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6497" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2071\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2058\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2045"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6498" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2072\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2059\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2046"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6499" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2073\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2060\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2047"
			case "7":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=6724" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2074\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2061\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2048"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6782" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2075\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2062\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2049"
			case "9":
				print_string = anime_say + "不對，我搞錯了這下週才有吧！" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n雖然還沒播，但有還有前作可以看喔！\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2076\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2063\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2050"
			case "10":
				print_string = anime_say + "不對，我搞錯了這集根本還沒播XD" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n雖然還沒播，但有還有前作可以看喔！\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2070\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2064\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2051"
			case "11":
				print_string = anime_say + "不對，我搞錯了這集根本還沒播XD" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n雖然還沒播，但有還有前作可以看喔！\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2078\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2065\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2052"
			case "12":
				print_string = anime_say + "不對，我搞錯了這集根本還沒播XD" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n雖然還沒播，但有還有前作可以看喔！\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2079\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2066\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2053"
			case "13":
				print_string = anime_say + "不對，我搞錯了這集根本還沒播XD" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n雖然還沒播，但有還有前作可以看喔！\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2080\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2067\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2054"
			default:
			}
		default:
			print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我找不到這部動畫，我還沒學呢...（目前只會歌王子跟美術社）\n我目前知道的動畫還很少，因為我考試不及格QAQ\n\n（其實是因為開發者純手動輸入，沒用資料庫跟爬蟲。才會增加比較慢XD）"
		}
	default:
		print_string = "HI～ 我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂＞A＜\n\nPS：目前這隻喵只提供查詢動畫的功能"
	}
	return print_string
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
 				//message.ID
				//message.Text
				bot_msg := "你是說 " + message.Text + " 嗎？\n\n我看看喔...等我一下..."
				// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
				// 	log.Print(err)
				// }

				//這裡開始增
				//switch 測試
				switch message.Text {
				case "0":
					bot_msg = "1"
				case "1":
					bot_msg = "2"
				case "2":
					bot_msg = "3"
				case "3":
					bot_msg = "4"
				case "4":
					bot_msg = "5"
				case "5":
					bot_msg = "6"
				default:
				}
				
				//anime
				bot_msg = anime(message.Text,message.ID)
				
				//增加到這
//				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
					log.Print(err)
				}
				//https://devdocs.line.me/en/?go#send-message-object
			}
		}
	}
}
