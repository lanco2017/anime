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
	reg := regexp.MustCompile("^.*(動畫|動畫瘋|巴哈姆特|anime|Anime|ａｎｉｍｅ|Ａｎｉｍｅ|アニメ)(\\s|　|:|;|：|；)([\u4e00-\u9fa5_a-zA-Z0-9]*)\\D*(\\d{1,})") //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	switch reg.ReplaceAllString(text, "$1"){
	case "新番":
		print_string = "最近一期是日本 2016 十月開播的動畫：\n" + 
		"歌之☆王子殿下♪ 真愛 LEGEND STAR\n" +
		"伯納德小姐說。\n" +
		"長騎美眉\n" +
		"3 月的獅子\n" +
		"黑白來看守所\n" +
		"我太受歡迎了該怎麼辦\n" +
		"無畏魔女"
	case "目錄","動畫清單","清單","索引","index","Index","介紹","動漫","動畫介紹","動漫介紹","info","Info":
		print_string = "你可以問我下面這些動畫，我會帶你去看！\n\n" +
		"※ 想知道最近新出的動畫可以輸入：「新番」查詢 \n" +
		"※ 以下是目前能夠查詢的動畫，\n冒號後面是簡短搜尋法。\n當然打跟巴哈姆特一樣的全名也可以。\n\n" +
		"這個美術社大有問題：美術社\n" +
		"歌之☆王子殿下♪ 系列：歌王子\n" +
		"三月的獅子：3月\n" +
		"我太受歡迎了該怎麼辦：我太受歡迎\n" +
		"長騎美眉：長騎\n" +
		"少年阿貝GO！GO！小芝麻\n" +
		"神裝少女小纏\n" +
		"夏目友人帳 伍\n" +
		"黑白來看守所\n" +
		"喵阿愣！\n" +
		"雙星之陰陽師\n" +
		"無畏魔女\n" +
		"伯納德小姐說\n" +
		"漂流武士\n" +
		"JOJO 的奇妙冒險 不滅鑽石：JOJO\n" +
		"影子籃球員\n" +
		"星夢手記\n" +
		"文豪野犬 第二季\n" +
		"機動戰士鋼彈 鐵血孤兒 第二季\n\n" +
		"搜尋方法：\n動畫 動畫名(或短名) 數字\n三個項目中間要用空白或冒號、分號隔開。\n\n例如：\n巴哈姆特　3月　１１\n動畫瘋　我太受歡迎 １\nアニメ;影子籃球員;15\n動畫 雙星 1\nanime：黑白來：5\n\n都可以"
	case "開發者":
		print_string = "你找我主人？OK！\n我跟你講我的夥伴喵在哪，你去加他。\n他跟主人很親近的，跟他說的話主人都會看到。\nhttps://line.me/R/ti/p/%40uwk0684z\n\n\n你也可以從下面這個連結直接去找主人線上對話。\n\n如果他不在線上一樣可以留言給他，\n他會收到的！\n這跟手機、電腦桌面軟體都有同步連線。" +
		"\n\nhttp://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + user_msgid +
		"&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A"
	case "動畫", "動畫瘋","巴哈姆特", "anime", "アニメ","Anime":
		print_string = text + "？\n好像有這個動畫耶，但我找不太到詳細的QQ\n你要手動去「巴哈姆特動畫瘋」找找嗎？\n\nhttps://ani.gamer.com.tw"
		anime_say := "有喔！有喔！你在找這個對吧！？\n"
		switch reg.ReplaceAllString(text, "$3") {
		case "影子籃球員":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3896"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3897"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3898"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3899"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3900"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3901"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3902"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3903"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3904"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3905"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3906"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3907"
			case "13":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3908"
			case "14":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3909"
			case "15":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3910"
			case "16":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3911"
			case "17":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3912"
			case "18":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3913"
			case "19":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3914"
			case "20":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3915"
			case "21":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3916"
			case "22":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3917"
			case "23":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3918"
			case "24":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3919"
			case "25":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3920" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "文豪野犬 第二季","文豪野犬":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6471"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6492"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6550"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6493"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6494"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6495"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6775"
			default:
			}
		case "機動戰士鋼彈 鐵血孤兒 第二季","機動戰士鋼彈 鐵血孤兒","機動戰士鋼彈","鐵血孤兒":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6400"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6401"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6402"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6403"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6404"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6405"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6779"
			default:
			}
		case "星夢手記":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6769"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6770"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6771"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6772"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6773"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6774"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6784"
			default:
			}	
		case "少年阿貝GO！GO！小芝麻","少年阿貝","阿貝","小芝麻","芝麻","GO！":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4999"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5007"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5008"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5120"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5168"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5175"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5176"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5473"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5710"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5711"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5712"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5786"
			case "13":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5882"
			case "14":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6340"
			case "15":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6347"
			case "16":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6348"
			case "17":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6364"
			case "18":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6365"
			case "19":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6366"
			case "20":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6549"
			case "21":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6725"
			case "22":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6726"
			default:
			}
		case "神裝少女小纏","小纏","神裝少女":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6406"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6431"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6432"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6433"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6434"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6435"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6729"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6469"
			default:
			}
		case "夏目友人帳 伍","夏目友人帳","夏目","有人","有人帳","帳":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6425"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6426"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6427"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6428"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6429"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6430"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6732"
			default:
			}
		case "黑白來看守所","黑白來","看守所","黑白":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6473"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6474"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6475"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6476"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6477"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6478"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6735"
			default:
			}
		case "喵阿愣！","喵阿愣","喵啊愣！","阿愣","啊愣":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6850"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6851"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6852"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6853"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6854"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6855"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6856"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6857"
			default:
			}
		case "雙星之陰陽師","雙星":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4998"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5027"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5028"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5029"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5110"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5111"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5169"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5470"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5471"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5472"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5707"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5708"
			case "13":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5709"
			case "14":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5785"
			case "15":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5923"
			case "16":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5924"
			case "17":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6029"
			case "18":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6030"
			case "19":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6175"
			case "20":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6176"
			case "21":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6177"
			case "22":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6307"
			case "23":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6308"
			case "24":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6309"
			case "25":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6370"
			case "26":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6371"
			case "27":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6465"
			case "28":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6466"
			case "29":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6467"
			case "30":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6468"
			case "31":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6762"
			case "32":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6763"
			default:
			}
		case "無畏魔女":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6419"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6420"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6421"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6422"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6423"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6424"
			default:
			}
		case "伯納德小姐說","小姐說","伯納德","伯納":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6710"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6711"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6712"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6713"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6723"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6843"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6859"
			default:
			}
		case "漂流武士":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6485"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6486"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6487"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6488"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6489"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6490"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6871"
			default:
			}
		case "JOJO 的奇妙冒險 不滅鑽石","JOJO","JOJO的奇妙冒險","奇妙冒險":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4994"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5005"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5019"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5020"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5096"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5097"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5165"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5191"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5192"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5508"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5509"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5715"
			case "13":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5776"
			case "14":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5790"
			case "15":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5873"
			case "16":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5874"
			case "17":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5912"
			case "18":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5913"
			case "19":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6154"
			case "20":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6155"
			case "21":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6205"
			case "22":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6206"
			case "23":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6338"
			case "24":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6339"
			case "25":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6356"
			case "26":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6357"
			case "27":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6368"
			case "28":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6369"
			case "29":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6463"
			case "30":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6464"
			case "31":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6706"
			case "32":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6707"
			case "33":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6708"
			case "34":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6709"
			default:
			}
		case "長騎美眉","長騎","單車":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6407"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6408"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6409"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6410"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6411"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6412"
			default:
			}
		case "我太受歡迎了該怎麼辦","我太受歡迎","受歡迎":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6413"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6414"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6415"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6416"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6417"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6418"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6865"
			default:
			}
		case "3月","3月的獅子","三月的獅子","三月":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6479"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6480"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6481"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6482"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6483"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6484"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6889"
			default:
			}
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
			print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我找不到這部動畫，我還沒學呢...\n（可輸入「目錄」查看支援的作品）\n我目前知道的動畫還很少，因為我考試不及格QAQ\n\n（其實是因為開發者半手動輸入更新，沒用自動化爬蟲跟資料庫。才會增加比較慢XD）"
		}
	default:
		print_string = "HI～ 我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這隻喵只提供查詢動畫的功能。\n如有其他建議或想討論，請對這隻貓輸入「開發者」進行聯絡。"
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

				
				//anime
				bot_msg = anime(message.Text,message.ID)
				
				//增加到這
//				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
// 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
// 					log.Print(err)
// 				}
				//https://devdocs.line.me/en/?go#send-message-object
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
					log.Print(err)
				}				
			case *linebot.ImageMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你丟圖")).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.VideoMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你丟影片")).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.AudioMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你丟聲音")).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.LocationMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你丟定位")).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.StickerMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你丟貼圖")).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
