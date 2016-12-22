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

	//"encoding/json" //https://golang.org/pkg/encoding/json/#example_Decoder
	
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

func anime(text string,user_msgid string,reply_mode string) string {
	print_string := text
	text = real_num(text)
	//	reg := regexp.MustCompile(`^.*(動畫|動畫瘋|巴哈姆特|anime|アニメ).*(這個美術社大有問題|美術社)\D*(\d{1,})`) //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	reg := regexp.MustCompile("^.*(動畫|動畫瘋|懶|巴哈姆特|anime|Anime|ａｎｉｍｅ|Ａｎｉｍｅ|アニメ)(\\s|　|:|;|：|；)([\u4e00-\u9fa5_a-zA-Z0-9]*)\\D*(\\d{1,})") //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	
	log.Print("--抓取分析觀察--")
	log.Print(reg.ReplaceAllString(text, "$1"))
	log.Print(reg.ReplaceAllString(text, "$2"))
	log.Print(reg.ReplaceAllString(text, "$3"))
	log.Print(reg.ReplaceAllString(text, "$4"))
	log.Print(reg.ReplaceAllString(text, "--抓取分析結束--"))
	
	switch reg.ReplaceAllString(text, "$1"){
	case "test":
		print_string = "GOTEST"
	case "新番":
		print_string = "最近一期是日本 2016 十月開播的動畫：\n" + 
		"歌之☆王子殿下♪ 真愛 LEGEND STAR\n" +
		"長騎美眉\n" +
		"3 月的獅子\n" +
		"黑白來看守所\n" +
		"我太受歡迎了該怎麼辦\n" +
		"無畏魔女"
	case "bot","機器人","目錄","動畫清單","清單","索引","ｉｎｄｅｘ","index","Index","簡介","介紹","動漫","動畫介紹","動漫介紹","info","Info","ｉｎｆｏ":
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
		"漂流武士\n" +
		"JOJO 的奇妙冒險 不滅鑽石：JOJO\n" +
		"影子籃球員\n" +
		"星夢手記\n" +
		"文豪野犬 第二季\n" +
		"機動戰士鋼彈 鐵血孤兒 第二季\n" +
		"路人超能 100\n" +
		"釣球\n\n" +
		"搜尋方法：\n動畫 動畫名(或短名) 數字\n三個項目中間要用空白或冒號、分號隔開。\n\n例如：\n巴哈姆特　3月　１１\n動畫瘋　我太受歡迎 １\nアニメ;影子籃球員;15\n動畫 雙星 1\nanime：黑白來：5\n\n都可以"
	case "開發者","admin","Admin","ａｄｍｉｎ":
		print_string = "你找我主人？OK！\n我跟你講我的夥伴喵在哪，你去加他。\n他跟主人很親近的，跟他說的話主人都會看到。\nhttps://line.me/R/ti/p/%40uwk0684z\n\n\n你也可以從下面這個連結直接去找主人線上對話。\n\n如果他不在線上一樣可以留言給他，\n他會收到的！\n這跟手機、電腦桌面軟體都有同步連線。" +
		"\n\nhttp://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + user_msgid +
		"&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A"
	case "動畫", "動畫瘋","巴哈姆特", "anime", "アニメ","Anime","ａｎｉｍｅ","Ａｎｉｍｅ","懶":
		print_string = text + "？\n好像有這個動畫耶，但我找不太到詳細的QQ\n你要手動去「巴哈姆特動畫瘋」找找嗎？\n\nhttps://ani.gamer.com.tw"
		anime_say := "有喔！有喔！你在找這個對吧！？\n"
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
		case "路人超能 100","路人","靈能":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5863"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5881"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5893"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5894"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5895"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5896"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6184"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6185"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6256"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6257"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6258"
			case "12","End","END","end":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6259" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "釣球":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6992"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6993"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6994"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6995"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6996"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6997"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6998"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6999"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7000"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7001"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7002"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7003" + "\n\n等等！這是最後一話！？"
			default:
			}
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
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6130"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6131"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6132"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6133"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6134"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6135"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6136"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6137"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6138"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6139"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6140"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6141"
			case "13":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6471"
			case "14":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6492"
			case "15":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6550"
			case "16":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6493"
			case "17":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6494"
			case "18":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6495"
			case "19":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6775"
			case "20":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6776"
			case "21":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6777"
			case "22":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6778"
			case "23":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6877"
			case "24":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6878" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "機動戰士鋼彈 鐵血孤兒 第二季","機動戰士鋼彈 鐵血孤兒","機動戰士鋼彈","鐵血孤兒":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			case "1":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2543"
			case "2":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2544"
			case "3":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2545"
			case "4":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2546"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2547"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2548"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2549"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2550"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3050"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3051"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4019"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4027"
			case "13":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4028"
			case "14":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4614"
			case "15":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4712"
			case "16":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4738"
			case "17":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4763"
			case "18":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4774"
			case "19":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4775"
			case "20":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4780"
			case "21":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4781"
			case "22":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4843"
			case "23":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4857"
			case "24":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4859"
			case "25":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4865"
			case "26":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6400"
			case "27":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6401"
			case "28":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6402"
			case "29":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6403"
			case "30":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6404"
			case "31":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6405"
			case "32":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6779"
			case "33":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6780"
			case "34":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6781"
			case "35":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6899"
			case "36":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6900"
			case "37":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6901"
			default:
			}
		case "星夢手記","星夢":
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
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6903"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6947"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7016"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7017"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7018"
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
			case "23":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6727"
			case "24":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6728"
			case "25":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6834"
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
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6730"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6731"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6836"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6837"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7025"
			case "特別","OVA":
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
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6845"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6733"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6734"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6838"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6839"
			case "特別","OST":
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
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6736"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6737"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6840"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6841"
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
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6917"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6918"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6919"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6920"
			case "13":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=6921"
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
			case "33":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6764"
			case "34":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6765"
			case "35":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6846"
			case "36":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6847"
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
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6766"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6767"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6848"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6849"
			default:
			}
		case "伯納德小姐說","小姐說","伯納德","伯納":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$4") {
			default:
				print_string = "可能不受歡迎或其他原因，\n很遺憾這部已經下架，請幫QQ"
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
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6872"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6873"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6874"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6875"
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
			case "35":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6879"
			case "36":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6880"
			case "37":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6881"
			case "38":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6882"
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
			case "4.5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6411"
			case "5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6412"
			case "6":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6884"
			case "7":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6885"
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6886"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6887"
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
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6866"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6867"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6868"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6869"
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
			case "8":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6890"
			case "9":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6891"
			case "10":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=6892"
			default:
			}
		case "美術社","美術社大有問題","這個美術社大有問題":
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
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6783" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n雖然還沒播，但有還有前作可以看喔！\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2076\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2063\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2050"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6895" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n雖然還沒播，但有還有前作可以看喔！\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2070\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2064\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2051"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6896" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n雖然還沒播，但有還有前作可以看喔！\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2078\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2065\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2052"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6897" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n雖然還沒播，但有還有前作可以看喔！\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2079\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2066\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2053"
			case "13":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6898" + 
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
		if reply_mode!="" {
			print_string = "HI～ 我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這隻喵只提供查詢動畫的功能。\n如有其他建議或想討論，請對這隻貓輸入「開發者」進行聯絡。"
		} else {
            print_string = "" //安靜模式
		}
	}
	return print_string
}

//http://qiita.com/koki_cheese/items/66980888d7e8755d01ec
// func handleTask(w http.ResponseWriter, r *http.Request) {
// }
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
		if event.Type == linebot.EventTypePostback {
				log.Print("觸發 Postback 功能（不讓使用者察覺的程式利用）")
				log.Print("event.Postback.Data = " + event.Postback.Data)
		}
		//觸發加入好友
		if event.Type == linebot.EventTypeFollow {
				target_user := event.Source.UserID + event.Source.GroupID + event.Source.RoomID	//target_user := ""
				log.Print("觸發與 " + target_user + " 加入好友")
			    imageURL := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/c2704b19816673a30c76cdccf67bcf8f/2016_-_%E8%A4%87%E8%A3%BD.png"
				template := linebot.NewCarouselTemplate(
					linebot.NewCarouselColumn(
						imageURL, "查詢巴哈姆特動畫瘋的功能", "我很愛看巴哈姆特動畫瘋。\n問我動畫可以這樣問：動畫 動畫名稱 集數",
						linebot.NewPostbackTemplateAction("動畫 美術社 12","動畫 美術社 12", "動畫 美術社 12"),
						linebot.NewMessageTemplateAction("アニメ 美術社大有問題 12", "アニメ 美術社大有問題 12"),
						linebot.NewMessageTemplateAction("anime：美術社：１", "anime：美術社：１"),
					),
					linebot.NewCarouselColumn(
						imageURL, "其他使用例", "開頭可以是 動畫 / anime / アニメ / 巴哈姆特",
						linebot.NewMessageTemplateAction("巴哈姆特 三月 ３", "巴哈姆特 三月 ３"),
						linebot.NewMessageTemplateAction("Ａｎｉｍｅ　喵阿愣　５", "Ａｎｉｍｅ　喵阿愣　５"),
						linebot.NewMessageTemplateAction("anime：黑白來：7", "anime：黑白來：7"),
					),
					linebot.NewCarouselColumn(
						imageURL, "其他功能", "新番、可查詢的動畫清單",
						linebot.NewMessageTemplateAction("新番", "新番"),
						linebot.NewMessageTemplateAction("可查詢的動畫清單", "目錄"),
						linebot.NewURITemplateAction("缺漏回報", "http://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + target_user + "&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A"),
					),
					linebot.NewCarouselColumn(
						imageURL, "意見反饋 feedback", "你可以透過此功能\n對 開發者 提出建議",
						linebot.NewURITemplateAction("加開發者 LINE", "https://line.me/R/ti/p/@uwk0684z"),
						linebot.NewURITemplateAction("線上與開發者聊天", "http://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + target_user + "&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A"),
						linebot.NewMessageTemplateAction("聯絡 LINE 機器人開發者", "開發者"),
					),
				)
				t_msg := "我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這隻喵只提供查詢動畫的功能。\n如有其他建議或想討論，請對這隻貓輸入「開發者」進行聯絡。"
				obj_message := linebot.NewTemplateMessage(t_msg, template)

				username := ""
				if target_user == "U6f738a70b63c5900aa2c0cbbe0af91c4"{
					username = "懶懶"
				}
				if target_user == "Uf150a9f2763f5c6e18ce4d706681af7f"{
					username = "包包"
				}
			//reply 的寫法
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好啊！" + username + "～\n想知道我的嗜好，可以說：簡介\nPS：手機上可以看到不一樣的內容喔！"),obj_message).Do(); err != nil {
					log.Print(err)
			}

		}
		//觸發解除好友
		if event.Type == linebot.EventTypeUnfollow {
				log.Print("觸發與 " + event.Source.UserID + event.Source.GroupID + event.Source.RoomID + " 解除好友")
				if _, err = bot.ReplyMessage(event.Source.UserID + event.Source.GroupID + event.Source.RoomID, linebot.NewTextMessage("那我走囉！\n哪天還想用用看歡迎隨時加我！\n\nhttps://line.me/R/ti/p/@pyv6283b\n或用 LINE ID 搜尋 @pyv6283b")).Do(); err != nil {
						log.Print(err)
				}
		}
		//觸發加入群組聊天
		if event.Type == linebot.EventTypeJoin {
				log.Print("觸發加入群組對話")
 				source := event.Source
 				log.Print("觸發加入群組聊天事件 = " + source.GroupID)
 				push_string := "很高興你邀請我進來這裡聊天！"
				if source.GroupID == "Ca78bf89fa33b777e54b4c13695818f81"{
					push_string += "\n你好，主人。"
				}
				//push 的寫法
				// 				if _, err = bot.PushMessage(source.GroupID, linebot.NewTextMessage(push_string)).Do(); err != nil {
				// 					log.Print(err)
				// 				}
				// 				if _, err = bot.PushMessage("Ca78bf89fa33b777e54b4c13695818f81", linebot.NewTextMessage("這裡純測試對嗎？\n只發於測試聊天室「test」")).Do(); err != nil {
				// 					log.Print(err)
				// 				}
				target_user := event.Source.UserID + event.Source.GroupID + event.Source.RoomID	//target_user := ""
			    imageURL := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/c2704b19816673a30c76cdccf67bcf8f/2016_-_%E8%A4%87%E8%A3%BD.png"
				template := linebot.NewCarouselTemplate(
					linebot.NewCarouselColumn(
						imageURL, "查詢巴哈姆特動畫瘋的功能", "我很愛看巴哈姆特動畫瘋。\n問我動畫可以這樣問：動畫 動畫名稱 集數",
						linebot.NewPostbackTemplateAction("動畫 美術社 12","動畫 美術社 12", "動畫 美術社 12"),
						linebot.NewMessageTemplateAction("アニメ 美術社大有問題 12", "アニメ 美術社大有問題 12"),
						linebot.NewMessageTemplateAction("anime：美術社：１", "anime：美術社：１"),
					),
					linebot.NewCarouselColumn(
						imageURL, "其他使用例", "開頭可以是 動畫 / anime / アニメ / 巴哈姆特",
						linebot.NewMessageTemplateAction("巴哈姆特 三月 ３", "巴哈姆特 三月 ３"),
						linebot.NewMessageTemplateAction("Ａｎｉｍｅ　喵阿愣　５", "Ａｎｉｍｅ　喵阿愣　５"),
						linebot.NewMessageTemplateAction("anime：黑白來：7", "anime：黑白來：7"),
					),
					linebot.NewCarouselColumn(
						imageURL, "其他功能", "新番、可查詢的動畫清單",
						linebot.NewMessageTemplateAction("新番", "新番"),
						linebot.NewMessageTemplateAction("可查詢的動畫清單", "目錄"),
						linebot.NewURITemplateAction("缺漏回報", "http://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + target_user + "&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A"),
					),
					linebot.NewCarouselColumn(
						imageURL, "意見反饋 feedback", "你可以透過此功能\n對 開發者 提出建議",
						linebot.NewURITemplateAction("加開發者 LINE", "https://line.me/R/ti/p/@uwk0684z"),
						linebot.NewURITemplateAction("線上與開發者聊天", "http://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + target_user + "&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A"),
						linebot.NewMessageTemplateAction("聯絡 LINE 機器人開發者", "開發者"),
					),
				)
				t_msg := "我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這隻喵只提供查詢動畫的功能。\n如有其他建議或想討論，請對這隻貓輸入「開發者」進行聯絡。"
				obj_message := linebot.NewTemplateMessage(t_msg, template)

			//reply 的寫法
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("群組聊天的各位大家好哇～！\n" + push_string + "\n\n想知道我的嗜好，請說：簡介"),obj_message).Do(); err != nil {
					log.Print(err)
			}
		}
		//觸發離開群組聊天
		if event.Type == linebot.EventTypeLeave {
				log.Print("觸發離開 " + event.Source.UserID + event.Source.GroupID + event.Source.RoomID +  " 群組")
		}
		if event.Type == linebot.EventTypeBeacon {
				log.Print("觸發 Beacon（啥鬼）")
		}
		//觸發收到訊息
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				//http://www.netadmin.com.tw/images/news/NP161004000316100411441903.png
				//userID := event.Source.UserID

 				//message.ID
				//message.Text
				log.Print(message.ID)
				log.Print(message.Text)
				bot_msg := "你是說 " + message.Text + " 嗎？\n\n我看看喔...等我一下..."
				// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
				// 	log.Print(err)
				// }
				
				//2016.12.20+
				//只有在 1 對 1 才能抓到 User ID　在群組才能抓到 event.Source.GroupID
 				log.Print("event.Source.UserID = " + event.Source.UserID)
				log.Print("event.Source.GroupID = " + event.Source.GroupID)
				log.Print("event.Source.RoomID = " + event.Source.RoomID)
				
// 				source := event.Source
// 				log.Print("source.UserID = " + source.UserID)
				
// 				userID := event.Source.UserID
// 				log.Print("userID := event.Source.UserID = " + userID)

				target_user := event.Source.UserID + event.Source.GroupID + event.Source.RoomID	//target_user := ""
				// if event.Source.UserID == ""{
				// 	target_user = event.Source.GroupID
				// } else {
				// 	target_user = event.Source.UserID
				// }
				
				
				//anime
				bot_msg = anime(message.Text,target_user,"")//bot_msg = anime(message.Text,message.ID,"")
				log.Print("根據 anime function 匹配到的回應內容：" + bot_msg)
				
								//增加到這
					//				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
					// 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
					// 					log.Print(err)
					// 				}
								//https://devdocs.line.me/en/?go#send-message-object
				

				//沒辦法建立 function 直接在裡面操作， 只好先用加法，從下游進行正則分析處理 reg  //https://play.golang.org/p/cjO5La2cKR
				reg := regexp.MustCompile("^.*(有喔！有喔！你在找這個對吧！？)\\n(https?.*)(\\n*.*)$")
				log.Print("--抓取［" + bot_msg + "］分析觀察--")
				log.Print("anime 後的 1 = " + reg.ReplaceAllString(bot_msg, "$1"))
				log.Print("anime 後的 2 = " + reg.ReplaceAllString(bot_msg, "$2")) //URL
				log.Print("完結篇廢話 = 3 = " + reg.ReplaceAllString(bot_msg, "$3")) //完結篇的廢話




				//anime url get //2016.12.22+
				anime_url := reg.ReplaceAllString(bot_msg, "$2")

				//判斷得到的 $2 是不是 http 開頭字串
				reg_http := regexp.MustCompile("^(http)s?.*") 

				if reg_http.ReplaceAllString(anime_url,"$1") != "http"{
					log.Print("anime_url = " + anime_url)
					anime_url = ""
				}


				//判斷是不是找不到
				reg_nofind := regexp.MustCompile("^你是要找.*\\n.*\\n.*\\n.*\\n.*\\n.*(才會增加比較慢XD）)$") 





				//2016.12.20+ for test
				if bot_msg != ""{
					if bot_msg == "GOTEST"{
						//bot_msg = "HI～ 我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這隻喵只提供查詢動畫的功能。\n如有其他建議或想討論，請對這隻貓輸入「開發者」進行聯絡。"
						bot_msg = "有喔！有喔！你在找這個對吧！？\n" + "https://ani.gamer.com.tw/animeVideo.php?sn=5863" + "\n\n等等！這是最後一話！？"

						// "http://ani.gamer.com.tw/animeVideo.php?sn=6878",
						//  第？話",
						//  "https://p2.bahamut.com.tw/B/2KU/33/0001485933.PNG",
						//  "查詢結果",
						//  "動畫名稱 ",
						// bot_msg 

						//log.Print("完結篇廢話 = 3 = " + reg.ReplaceAllString(bot_msg, "$3")) //完結篇的廢話

						//Create message
						//https://github.com/line/line-bot-sdk-go
						//https://github.com/line/line-bot-sdk-go/blob/master/linebot/message.go

					//模板成功  //官方範例 https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
						//linebot.NewTemplateMessage
						// //1 confirm 純是否類型的問法
						// //.NewConfirmTemplate 模板，裡面最多只能有兩個動作，按鈕只能左右
						// //.NewMessageTemplateAction 發言動作

						// template := linebot.NewConfirmTemplate(
						// 	"Do it?",
						// 	linebot.NewMessageTemplateAction("Yes", "Yes!"),
						// 	linebot.NewMessageTemplateAction("No", "No!"),
						// )

 					//     leftBtn := linebot.NewMessageTemplateAction("left", "left clicked")// 後面的參數 "left clicked" = 在使用者按下後，自動幫使用者發訊息
 					//     rightBtn := linebot.NewMessageTemplateAction("right", "right clicked")// 後面的參數 "right clicked" = 在使用者按下後，自動幫使用者發訊息
					 //    //.NewMessageTemplateAction("字面按鈕", "設定讓使用者按下後發送內容") 會讓使用者發送那樣的內容給系統
 					//     template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)
 					//     //.NewConfirmTemplate

						//linebot.NewTemplateMessage
 					    //2 buttons
 					    //.NewButtonsTemplate 模板，裡面最多只能有四個動作
 					    //.NewURITemplateAction 開啟指定網址的動作
 					    //.NewPostbackTemplateAction ？？動作
 					    //						第二參數可以讓她　ＰＯＳＴ指定內容（但還不會處理．．．）	第三參數類似於 .NewMessageTemplateAction 的效果
 					//     imageURL := "https://images.gamme.com.tw/news2/2016/51/39/paCYoqCXkqSarqSZ.jpg"
						// template := linebot.NewButtonsTemplate(
						// 	imageURL, "你好歡迎光臨", "這是內文",							//這前三個 分別是圖片(必須https)、標題、內文
						// 	linebot.NewURITemplateAction("來我的網站", "https://synr.github.io"),
						// 	linebot.NewPostbackTemplateAction("目錄查詢", "目錄", "目錄"),
						// 	linebot.NewPostbackTemplateAction("開發者", "開發者", "開發者"),
						// 	linebot.NewMessageTemplateAction("Say message", "Rice=米"),
						// )

						//linebot.NewTemplateMessage
						//3 carousel .NewCarouselTemplate  最多可以並排五個「.NewCarouselColumn」的樣板，
						//「.NewCarouselColumn」裡面最多只能有三個動作按鈕，但並列的其他項目也要一致數量才能。2016.12.22+
						//圖片可以是 PNG
						// imageURL := "https://images.gamme.com.tw/news2/2016/51/39/paCYoqCXkqSarqSZ.jpg"
						// template := linebot.NewCarouselTemplate(
						// 	linebot.NewCarouselColumn(
						// 		"https://p2.bahamut.com.tw/B/2KU/33/0001485933.PNG", "hoge", "fuga",
						// 		linebot.NewURITemplateAction("測試看動畫", "http://ani.gamer.com.tw/animeVideo.php?sn=6878"),
						// 		linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", ""),
						// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
						// 	),
						// 	linebot.NewCarouselColumn(
						// 		"https://p2.bahamut.com.tw/B/2KU/18/0001484818.PNG", "hoge", "fuga",
						// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
						// 		linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
						// 		linebot.NewMessageTemplateAction("Say message", "Rice=米"),
						// 	),
						// 	linebot.NewCarouselColumn(
						// 		imageURL, "hoge", "fuga",
						// 		linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
						// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
						// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
						// 	),
						// 	linebot.NewCarouselColumn(
						// 		imageURL, "hoge", "fuga",
						// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
						// 		linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
						// 		linebot.NewMessageTemplateAction("Say message", "Rice=米"),
						// 	),
						// 	linebot.NewCarouselColumn(
						// 		imageURL, "hoge", "fuga",
						// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
						// 		linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
						// 		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
						// 	),
						// )
						//所以有三種樣板，有三種動作按鈕。兩個樣板可以放圖片，一個單純只能兩個按鈕。


 					    //obj_message := linebot.NewTemplateMessage("HI～ 我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這隻喵只提供查詢動畫的功能。\n如有其他建議或想討論，請對這隻貓輸入「開發者」進行聯絡。", template)//messgage := linebot.NewTemplateMessage("請使用更新 APP 或使用手機 APP 才能看到這個功能。", template)
						//obj_message := linebot.NewTemplateMessage(bot_msg, template)
 					    //.NewTemplateMessage("無法支援按鈕模式時要發出的訊息",Template 物件)

							// 						if _, err = bot.ReplyMessage(event.ReplyToken, message).Do(); err != nil {
							// 							log.Print(err)
							// 						}


						//https://devdocs.line.me/en/?go#send-message-object


						//++ https://github.com/dongri/line-bot-sdk-go KEY:linebot.NewImageMessage

						//.NewImageMessage 發圖片成功
						//originalContentURL := "https://avatars0.githubusercontent.com/u/5731891?v=3&s=96"
    					//previewImageURL := "https://avatars0.githubusercontent.com/u/5731891?v=3&s=96"
    					//obj_message := linebot.NewImageMessage(originalContentURL, previewImageURL)

						//.NewStickerMessage 發貼貼圖成功	 //https://devdocs.line.me/files/sticker_list.pdf					
						//obj_message := linebot.NewStickerMessage("1", "1") //https://devdocs.line.me/en/?go#send-message-object

						//https://devdocs.line.me/en/?go#imagemap-message
						//https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
						// obj_message := linebot.NewImagemapMessage(
						// 	"https://synr.github.io/img/index.jpg",
						// 	"Imagemap alt text",
						// 	linebot.ImagemapBaseSize{1040, 1040},
						// 	linebot.NewURIImagemapAction("https://store.line.me/family/manga/en", linebot.ImagemapArea{0, 0, 520, 520}),
						// 	linebot.NewURIImagemapAction("https://store.line.me/family/music/en", linebot.ImagemapArea{520, 0, 520, 520}),
						// 	linebot.NewURIImagemapAction("https://store.line.me/family/play/en", linebot.ImagemapArea{0, 520, 520, 520}),
						// 	linebot.NewMessageImagemapAction("URANAI!", linebot.ImagemapArea{520, 520, 520, 520}),
						// )

						//Audio //https://github.com/dongri/line-bot-sdk-go
					    // originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/ok.m4a"
					    // duration := 1000
					    // obj_message := linebot.NewAudioMessage(originalContentURL, duration)

 					    //接收各種 message object
						//if _, err = bot.ReplyMessage(event.ReplyToken, obj_message,obj_message,obj_message,obj_message,obj_message).Do(); err != nil { //五聯發
						// if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil { 
						//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("1", "1"),linebot.NewStickerMessage("1", "2"),linebot.NewStickerMessage("2", "19"),linebot.NewStickerMessage("2", "20"),linebot.NewStickerMessage("1", "3")).Do(); err != nil {
						// 	log.Print(err)
						// }
					} else {
						if anime_url!=""{
							//找到的時候的 UI
	 					    imageURL := "https://i2.bahamut.com.tw/anime/FB_anime.png"
							template := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "動畫搜尋結果", "在找" + message.Text + "對吧！？\n建議可以直接在巴哈姆特動畫瘋 APP 裡面播放！",							
									linebot.NewURITemplateAction("點此播放找到的動畫", anime_url),
									linebot.NewURITemplateAction("下載巴哈姆特動畫瘋 APP", "https://prj.gamer.com.tw/app2u/animeapp.html"),
									linebot.NewMessageTemplateAction("查詢其他動畫", "目錄"),
								),
								linebot.NewCarouselColumn(
									"https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/c2704b19816673a30c76cdccf67bcf8f/2016_-_%E8%A4%87%E8%A3%BD.png", "意見反饋 feedback", "你可以透過此功能\n對 開發者 提出建議",
									linebot.NewURITemplateAction("加開發者 LINE", "https://line.me/R/ti/p/@uwk0684z"),
									linebot.NewURITemplateAction("線上與開發者聊天", "http://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + target_user + "&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A"),
									linebot.NewMessageTemplateAction("聯絡 LINE 機器人開發者", "開發者"),
								),
							)
							obj_message := linebot.NewTemplateMessage(bot_msg, template)

							originalContentURL_1 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/722268f159dc640ed1639ffd31b4dd0d/94455.jpg"
	    					previewImageURL_1 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/722268f159dc640ed1639ffd31b4dd0d/94455.jpg"
	    					obj_message_img_1 := linebot.NewImageMessage(originalContentURL_1, previewImageURL_1)

							originalContentURL_2 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/f7e158cdc3f1e9640a5f5cf188c33b13/94454.jpg"
	    					previewImageURL_2 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/f7e158cdc3f1e9640a5f5cf188c33b13/94454.jpg"
	    					obj_message_img_2 := linebot.NewImageMessage(originalContentURL_2, previewImageURL_2)

							if _, err = bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage("可參考以下圖例操作讓搜尋到的影片，直接在巴哈姆特動畫瘋 APP 進行播放。"),obj_message_img_1,obj_message_img_2,obj_message).Do(); err != nil {
								log.Print(err)
							}
						}else{
							//2016.12.22+ 利用正則分析字串結果，來設置觸發找不到的時候要 + 的 UI
							if reg_nofind.ReplaceAllString(bot_msg,"$1") == "才會增加比較慢XD）"{
								//找不到的時候
		 					    imageURL := "https://i2.bahamut.com.tw/anime/FB_anime.png"
								template := linebot.NewCarouselTemplate(
									linebot.NewCarouselColumn(
										imageURL, "找不到 "  +  message.Text   +   " 耶", "有可能打錯字或這真的沒有收錄，\n才會找不到。",							
										linebot.NewMessageTemplateAction("查看新番", "新番"),
										linebot.NewMessageTemplateAction("可查詢的其他動畫目錄", "目錄"),
										linebot.NewURITemplateAction("下載巴哈姆特動畫瘋 APP", "https://prj.gamer.com.tw/app2u/animeapp.html"),
									),
									linebot.NewCarouselColumn(
										"https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/c2704b19816673a30c76cdccf67bcf8f/2016_-_%E8%A4%87%E8%A3%BD.png", "意見反饋 feedback", "你可以透過此功能\n對 開發者 提出建議",
										linebot.NewURITemplateAction("加開發者 LINE", "https://line.me/R/ti/p/@uwk0684z"),
										linebot.NewURITemplateAction("線上與開發者聊天", "http://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + target_user + "&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A"),
										linebot.NewMessageTemplateAction("聯絡 LINE 機器人開發者", "開發者"),
									),
								)
								obj_message := linebot.NewTemplateMessage("除了「目錄」以外，\n你也可以輸入「新番」查詢近期的動畫。", template)
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg),obj_message).Do(); err != nil {
									log.Print(err)
								}
							}else{
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
									log.Print(err)
								}
							}
						}
					}
				}
// 				m := linebot.NewTextMessage("ok")
// 				    if _, err = bot.ReplyMessage(event.ReplyToken, m).Do(); err != nil {

// 				    }
				
				//----------PushMessage-----------這段可以跟 ReplyMessage 同時有效，但是只會在 1 對 1 有效。群組無效。---------
				//------開發者測試方案有效(好友最多50人/訊息無上限)，免費版(好友不限人數/訊息限制1000)、入門版無效，旗艦版、專業版有效。
				
				//http://muzigram.muzigen.net/2016/09/linebot-golang-linebot-heroku.html
				//https://github.com/mogeta/lbot/blob/master/main.go
 				source := event.Source
 				log.Print("source.UserID = " + source.UserID)
 				log.Print("target_user = " + target_user)
				//2016.12.20+//push_string := ""
// 				if source.UserID == "U6f738a70b63c5900aa2c0cbbe0af91c4"{
// 					push_string = "你好，主人。（PUSH_MESSAGE 才可以發）"
// 				}
// 				if source.UserID == "Uf150a9f2763f5c6e18ce4d706681af7f"{
// 					push_string = "唉呦，你是包包吼"
// 				}
//2016.12.20+ close push
// 					if source.Type == linebot.EventSourceTypeUser {
// 						if _, err = bot.PushMessage(source.UserID, linebot.NewTextMessage(push_string)).Do(); err != nil {
// 							log.Print(err)
// 						}
// 					}
// 					if source.Type == linebot.EventSourceTypeUser {
// 						if _, err = bot.PushMessage(source.UserID, linebot.NewTextMessage(push_string)).Do(); err != nil {
// 							log.Print(err)
// 						}
// 					}
					//上面重覆兩段 push 用來證明 push 才可以連發訊息框，re 只能一個框
				//---------------------這段可以跟 ReplyMessage 同時有效，但是只會在 1 對 1 有效。群組無效。---------
			case *linebot.ImageMessage:
						  // 				_, err := bot.SendText([]string{event.RawContent.Params[0]}, "Hi~\n歡迎加入 Delicious!\n\n想查詢附近或各地美食都可以LINE我呦！\n\n請問你想吃什麼?\nex:義大利麵\n\n想不到吃什麼，也可以直接'傳送目前位置訊息'")
					 	 // 				var img = "http://imageshack.com/a/img921/318/DC21al.png"
						// 				_, err = bot.SendImage([]string{content.From}, img, img)
						// 				if err != nil {
						// 					log.Println(err)
						// 				}
									
						// 				if err := bot.handleImage(message, event.ReplyToken); err != nil {
						// 					log.Print(err)
						// 				}
									//https://devdocs.line.me/en/#webhook-event-object
				log.Print("對方丟圖片 message.ID = " + message.ID)

				//log.Print("對方丟圖片 linebot.EventSource = " + linebot.EventSource

				//----------------------------------------------------------------取得使用者資訊的寫法
				source := event.Source

				userID := event.Source.UserID
				groupID := event.Source.GroupID
				RoomID := event.Source.RoomID
				markID := userID + groupID + RoomID
				
				log.Print(source.UserID)
				//----------------------------------------------------------------取得使用者資訊的寫法

				username := ""
				if markID == "U6f738a70b63c5900aa2c0cbbe0af91c4"{//if source.UserID == "U6f738a70b63c5900aa2c0cbbe0af91c4"{
					username = "懶懶 = " + userID + groupID + RoomID //2016.12.20+
				}
				if markID == "Uf150a9f2763f5c6e18ce4d706681af7f"{
					username = "包包"
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這圖片是？\n\n" + username + "你丟給我圖片幹嘛！\n我眼睛還沒長好看不懂XD")).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.VideoMessage:
				//https://github.com/dongri/line-bot-sdk-go
			    originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/video-original.mp4"
			    previewImageURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/video-preview.png"
			    obj_message := linebot.NewVideoMessage(originalContentURL, previewImageURL)
 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這影片是？\n我也給你影片吧！\n\n這只是測試功能"),obj_message).Do(); err != nil {
 					log.Print(err)
 				}
			case *linebot.AudioMessage:
				originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/ok.m4a"
				duration := 1000
				obj_message := linebot.NewAudioMessage(originalContentURL, duration)
 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這是什麼聲音？"),obj_message).Do(); err != nil {
 					log.Print(err)
 				}
			case *linebot.LocationMessage:
				log.Print("message.Title = " + message.Title)
				log.Print("message.Address = " + message.Address)
				log.Print("message.Latitude = ")
				log.Print(message.Latitude)
				log.Print("message.Longitude = ")
				log.Print(message.Longitude)
				obj_message := linebot.NewLocationMessage(message.Title, message.Address, message.Latitude, message.Longitude)
				obj_message_1 := linebot.NewLocationMessage("歡迎光臨", "地球", 25.022413, 121.556427) //麵包店 台北市信義區富陽街46號
				//obj_message_2 := linebot.NewLocationMessage("歡迎光臨", "哪個近", 25.022463, 121.556454) //這個遠
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你在這裡？"),obj_message,obj_message_1).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.StickerMessage:
				log.Print("message.PackageID = " + message.PackageID)
				log.Print("message.StickerID = " + message.StickerID)
				//https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go handleSticker
				//message.PackageID, message.StickerID
				obj_message_moto := linebot.NewStickerMessage(message.PackageID, message.StickerID)
				//https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
				//2016.12.20+ 多次框框的方式成功！（最多可以五個）
				//.NewStickerMessage 發貼貼圖成功	 //https://devdocs.line.me/files/sticker_list.pdf			
				obj_message := linebot.NewStickerMessage("2", "514") //https://devdocs.line.me/en/?go#send-message-object
 				//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("OU<"),linebot.NewTextMessage("0.0"),linebot.NewTextMessage("．ω．"),linebot.NewTextMessage("．ω．")).Do(); err != nil {
 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("．ω．"),obj_message_moto,obj_message).Do(); err != nil {
 					log.Print(err)
 				}
			}
		}
	}
}
