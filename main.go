// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	 http://www.apache.org/licenses/LICENSE-2.0
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

	"strconv"
	"strings"

	//"encoding/json" //https://golang.org/pkg/encoding/json/#example_Decoder
	
	"github.com/line/line-bot-sdk-go/linebot"

		//通訊
			//"bytes"

	// "io"
		//通訊//"io/ioutil"

	"image/jpeg"
	//"image/png"

	// https://github.com/go-martini/martini/blob/master/translations/README_zh_tw.md
	// "github.com/go-martini/martini"

	//http://ithelp.ithome.com.tw/articles/10159486
	//"github.com/alexcesaro/mail/gomail"

	// "os/exec"
	// "path/filepath"

	"crypto/md5"
	"encoding/hex"

	//http://l-lin.github.io/2015/01/31/Golang-Deploy_to_heroku
	// "database/sql"
	// _ "github.com/lib/pq"

)

var bot *linebot.Client

func main() {

	//http://www.qetee.com/exp/golang/golang-get-file-path/
 // 	execFileRelativePath, _ := exec.LookPath(os.Args[0])
 //	log.Println("执行程序与命令执行目录的相对路径　　　　:", execFileRelativePath)

	// execFileAbsPath, _ := filepath.Abs(execFileRelativePath)
 //	log.Println("执行程序的绝对路径　　　　　　　　　　　:", execFileAbsPath)

	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)

		// // // http://cepave.com/http-restful-api-with-golang/
	   //  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
				// w.Header().Set("Access-Control-Allow-Origin", "*")
				// w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
				// //http://qiita.com/futosu/items/b49f7d9e28101daaa99e
				// //https://play.golang.org/p/xHp44c_pJm
				// w.Header().Set("Access-Control-Allow-Headers","Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
				// log.Print(req)
	   //  	log.Print("進來了")
	   //	  req.ParseForm()
	   //	  if req.Method == "GET" || req.Method == "POST" {
	   //	  	log.Print("GP了")
	   //		  fmt.Println(req.ContentLength)
	   //		  //firstname := req.FormValue("type")
	   //		  //lastname := req.FormValue("text")
	   //		  //w.Write([]byte(fmt.Sprintf("[%s] Hello, %s %s!", req.Method, firstname, lastname)))
	   //	  } else {
	   //		  //http.Error(w, "The method is not allowed.", http.StatusMethodNotAllowed)
	   //	  }
	   //  })

	 //  m := martini.Classic()
		// m.Post("/", func() (int, string) {
		//   log.Print("POST")		//  return 418, "我是一個茶壺" // HTTP 418 : "我是一個茶壺"
		// })

		// m.NotFound(func() {
		//   log.Print("404")// handle 404
		// })

	 //  m.Run()

}

//https://gist.github.com/synr/d3d68d42b12204d981b39203a0b16762
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// func HttpGET_(url string) string {
// 	//https://internal-api.ifttt.com/maker
// 	log.Print("已經進來 GET")
// 	log.Print("url = " + url)

// 	// url := "https://hooks.zapier.com/hooks/catch/132196/txma4i/"

// 	req, err := http.NewRequest(
// 		"GET",
// 		url,
// 		nil,
// 	)
// 	if err != nil {
// 		log.Print(err)
// 		return ""
// 	}

// 	// Content-Type 設定
// 	//req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
// 	//req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Print(err)		
// 		return ""
// 	}
// 	defer resp.Body.Close()

// 	log.Print(err)

// 	//http://cepave.com/http-restful-api-with-golang/
//	 log.Print("HttpGET_response Status = ")
//	 log.Print(resp.Status)
//	 log.Print("HttpGET_response Headers = ")
//	 log.Print(resp.Header)
//	 rebody, _ := ioutil.ReadAll(resp.Body)
//	 log.Print("response Body = " + string(rebody))
// 	//http://cepave.com/http-restful-api-with-golang/

// 	return string(rebody) //return err
// }

// func HttpPost_Zapier(body , title_text, this_id, codename string) error {
// 	body = strings.Replace(body,"\n", `\n`, -1)
// 	title_text = strings.Replace(title_text,"\n", `\n`, -1)
// 	this_id = strings.Replace(this_id,"\n", `\n`, -1)
// 	codename = strings.Replace(codename,"\n", `\n`, -1)
// 	//https://internal-api.ifttt.com/maker
// 	log.Print("已經進來 Zapier POST")
// 	log.Print("body = " + body)
// 	log.Print("title_text = " + title_text)
// 	log.Print("this_id = " + this_id)

// 	url := "https://hooks.zapier.com/hooks/catch/132196/txma4i/"
// 	jsonStr := `{
// 		"value1":"` + body + `",
// 		"value2": "` + title_text + `",
// 		"value3": "` + this_id + `",
// 		"value4": "` + codename + `"
// 	}`

// 	req, err := http.NewRequest(
// 		"POST",
// 		url,
// 		bytes.NewBuffer([]byte(jsonStr)),
// 	)
// 	if err != nil {
// 		log.Print(err)
// 		return err
// 	}

// 	// Content-Type 設定
// 	//req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Print(err)		
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	log.Print(err)

// 	//http://cepave.com/http-restful-api-with-golang/
//	 // log.Print("response Status = ")
//	 // log.Print(resp.Status)
//	 // log.Print("response Headers = ")
//	 // log.Print(resp.Header)
//	 // rebody, _ := ioutil.ReadAll(resp.Body)
//	 // log.Print("response Body = " +string(rebody))
// 	//http://cepave.com/http-restful-api-with-golang/

// 	return err
// }

// func HttpPost_IFTTT(body , title_text, this_id string) error {
// 	body = strings.Replace(body,"\n", `\n`, -1)
// 	title_text = strings.Replace(title_text,"\n", `\n`, -1)
// 	this_id = strings.Replace(this_id,"\n", `\n`, -1)
// 	//https://internal-api.ifttt.com/maker
// 	log.Print("已經進來 IFTTT POST")
// 	log.Print("body = " + body)
// 	log.Print("title_text = " + title_text)
// 	log.Print("this_id = " + this_id)

// 	url := "https://maker.ifttt.com/trigger/linebot/with/key/WJCRNxQhGJuzPd-sUDext"
// 	jsonStr := `{
// 		"value1":"` + body + `",
// 		"value2": "` + title_text + `",
// 		"value3": "` + this_id + `"
// 	}`

// 	req, err := http.NewRequest(
// 		"POST",
// 		url,
// 		bytes.NewBuffer([]byte(jsonStr)),
// 	)
// 	if err != nil {
// 		log.Print(err)
// 		return err
// 	}

// 	// Content-Type 設定
// 	//req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Print(err)		
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	log.Print(err)

// 	//http://cepave.com/http-restful-api-with-golang/
//	 log.Print("response Status = ")
//	 log.Print(resp.Status)
//	 log.Print("response Headers = ")
//	 log.Print(resp.Header)
//	 rebody, _ := ioutil.ReadAll(resp.Body)
//	 log.Print("response Body = " +string(rebody))
// 	//http://cepave.com/http-restful-api-with-golang/

// 	return err
// }

// func HttpPost_JANDI(body, connectColor, title, code string) error {
// 	body = strings.Replace(body,"\n", `\n`, -1)
// 	title = strings.Replace(title,"\n", `\n`, -1)
// 	code = strings.Replace(code,"\n", `\n`, -1)
// 	log.Print("已經進來 JANDI POST")
// 	log.Print("body = " + body)
// 	log.Print("connectColor = " + connectColor)
// 	log.Print("title = " + title)
// 	log.Print("code = " + code)

// 	url := "https://wh.jandi.com/connect-api/webhook/11691684/46e7f45fd4f68a021afbd844aed66430"
// 	jsonStr := `{
// 		"body":"` + body + `",
// 		"connectColor":"` + connectColor + `",
// 		"connectInfo" : [{
// 				"title" : "` + title + `",
// 				"description" : "這是來自 LINE BOT 的通風報信",
// 				"imageUrl": "https://line.me/R/ti/p/@pyv6283b、@sjk2434l"
// 		},{
// 				"title" : "參考數據",
// 				"description" : "` + code + `"
// 		}]
// 	}`

// 	req, err := http.NewRequest(
// 		"POST",
// 		url,
// 		bytes.NewBuffer([]byte(jsonStr)),
// 	)
// 	if err != nil {
// 		log.Print(err)
// 		return err
// 	}

// 	// Content-Type 設定
// 	req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Print(err)		
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	log.Print(err)

// 	//http://cepave.com/http-restful-api-with-golang/
//	 log.Print("response Status = ")
//	 log.Print(resp.Status)
//	 log.Print("response Headers = ")
//	 log.Print(resp.Header)
//	 rebody, _ := ioutil.ReadAll(resp.Body)
//	 log.Print("response Body = " +string(rebody))
// 	//http://cepave.com/http-restful-api-with-golang/

// 	return err
// }



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


// //----javascript 代勞區	https://trello.com/c/OXnZDLzG/86-javascript-golang

// 	function get_anime(ad=''){
// 		var output_string = "		case \"" + document.title.replace(/(.*)\[.*/gi,"$1") + "\":\n		   \/\/reg.ReplaceAllString(text, \"$2\")\n			switch reg.ReplaceAllString(text, \"$4\") {\n";
// 		var num = ( ( (ad=='') || (ad==0) )   ?   1  : 0  );
// 		for (var i = 0; i < document.getElementsByTagName('a').length; i++) {
// 			if(document.getElementsByTagName('a')[i].href.indexOf('ani.gamer.com.tw\/animeVideo') != -1){
// 				if(num>0){
// 							output_string += "			case \"" + num + "\":\n" + "					print_string = anime_say + \"" + document.getElementsByTagName('a')[i].href + "\"\n";
// 				}
// 				num++;
// 			}
// 		}
// 		output_string += "			default:\n			}";
// 		console.log(output_string)
// 		//return output_string;
// 	}

// 	get_anime(0);//get_anime(); //沒廣告的時候
// 	//get_anime(1); //有廣告的時候

// 	//2017
// 	function get_anime(ad=''){
// 		var output_string = "		case \"" + document.title.replace(/(.*)\[.*/gi,"$1") + "\":\n		   \/\/reg.ReplaceAllString(text, \"$2\")\n			switch reg.ReplaceAllString(text, \"$4\") {\n";
// 		var num = ( ( (ad=='') || (ad==0) )   ?   1  : 0  );
// 		for (var i = 0; i < document.getElementsByTagName('a').length; i++) {
// 			if(document.getElementsByTagName('a')[i].href.indexOf('ani.gamer.com.tw\/animeVideo') != -1){
// 				if(num>0){
// 							if(num<10){
// 								output_string += "			case \"" + num + "\",\"0" + num + "\":\n" + "					print_string = anime_say + \"" + document.getElementsByTagName('a')[i].href + "\"\n";
// 							}else{
// 								output_string += "			case \"" + num + "\":\n" + "					print_string = anime_say + \"" + document.getElementsByTagName('a')[i].href + "\"\n";
// 							}
// 				}
// 				num++;
// 			}
// 		}
// 		output_string += "			default:\n			}";
// 		console.log(output_string)
// 		//return output_string;
// 	}

// 	get_anime(0);//get_anime(); //沒廣告的時候
// 	//get_anime(1); //有廣告的時候

// 	//起始不為 1 的時候，以 13 起始為例
// 	function get_anime(ad=''){
// 		var output_string = "		case \"" + document.title.replace(/(.*)\[.*/gi,"$1") + "\":\n		   \/\/reg.ReplaceAllString(text, \"$2\")\n			switch reg.ReplaceAllString(text, \"$4\") {\n";
// 		var num = ( ( (ad=='') || (ad==0) )   ?   1  : 0  );
// 		for (var i = 0; i < document.getElementsByTagName('a').length; i++) {
// 			if(document.getElementsByTagName('a')[i].href.indexOf('ani.gamer.com.tw\/animeVideo') != -1){
// 				if(num>0){
// 							output_string += "			case \"" + (num+13) + "\":\n" + "					print_string = anime_say + \"" + document.getElementsByTagName('a')[i].href + "\"\n";
// 				}
// 				num++;
// 			}
// 		}
// 		output_string += "			default:\n			}";
// 		console.log(output_string)
// 		//return output_string;
// 	}

// 	//get_anime(0);//get_anime(); //沒廣告的時候
// 	get_anime(1); //有廣告的時候


//new

// ----javascript 代勞區	https://trello.com/c/OXnZDLzG/86-javascript-golang
	// //2017 +1
	// function get_anime(ad=''){
	// 	 var output_string = "		case \"" + document.title.replace(/(.*)\[.*/gi,"$1") + "\":\n			\/\/reg.ReplaceAllString(text, \"$2\")\n			switch reg.ReplaceAllString(text, \"$5\") {\n";
	// 	 var num = ( ( (ad=='') || (ad==0) )   ?   1  : 0  );
	// 	 for (var i = 0; i < document.getElementsByTagName('a').length; i++) {
	// 		 if(document.getElementsByTagName('a')[i].href.indexOf('ani.gamer.com.tw\/animeVideo') != -1){
	// 			 if(num>0){
	// 						 if(num<10){
	// 						 	output_string += "			case \"" + num + "\",\"0" + num + "\":\n" + "				print_string = anime_say + \"" + document.getElementsByTagName('a')[i].href + "\"\n";
	// 						 }else{
	// 							output_string += "			case \"" + num + "\":\n" + "				print_string = anime_say + \"" + document.getElementsByTagName('a')[i].href + "\"\n";
	// 						}
	// 			 }
	// 			 num++;
	// 		 }
	// 	 }
	// 	 output_string += "			default:\n			}";
	// 	 console.log(output_string)
	// 	 //return output_string;
	// }

	// get_anime(0);//get_anime(); //沒廣告的時候
	// //get_anime(1); //有廣告的時候
// //----javascript 代勞區


func anime(text string,user_msgid string,reply_mode string) string {
	//https://gitter.im/kkdai/LineBotTemplate?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge：也可以透過 string.Contains("我要找的字", 原始字串) 來判斷
	print_string := text
	text = real_num(text)

	if GetMD5Hash(text) == "c38b3100b02ef42411a99b7975e4ff47" {
		print_string = "c38b3100b02ef42411a99b7975e4ff47"
		return print_string
	}

	//	reg := regexp.MustCompile(`^.*(動畫|動畫瘋|巴哈姆特|anime|アニメ).*(這個美術社大有問題|美術社)\D*(\d{1,})`) //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//2016
	//reg := regexp.MustCompile("^.*(動畫|動畫瘋|懶|巴哈|巴哈姆特|anime|Anime|ａｎｉｍｅ|Ａｎｉｍｅ|アニメ)(\\s|　|:|;|：|；)([\u4e00-\u9fa5_a-zA-Z0-9]*)\\D*(\\d{1,})") //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//2017
	//reg := regexp.MustCompile("^(動畫|動畫瘋|懶|巴哈|巴哈姆特|anime|Anime|ａｎｉｍｅ|Ａｎｉｍｅ|アニメ)(\\s|　|:|;|：|；)([\u4e00-\u9fa5_a-zA-Z0-9]*)\\D*(\\d{1,})") //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//2017.01.03+
	//reg := regexp.MustCompile("^(動畫|動畫瘋|懶|巴哈|巴哈姆特|anime|Anime|ａｎｉｍｅ|Ａｎｉｍｅ|アニメ)(\\s|　|:|;|：|；)([\u4e00-\u9fa5_a-zA-Z0-9]*)\\D*([0-9.]{1,})") //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//2017.01.05+ 日文（\u30a0-\u30ff\u3040-\u309f）   全形英文數字（\uff21-\uff3a\uff41-\uff5a\uff10-\uff19）
	//reg := regexp.MustCompile("^(動畫|動畫瘋|懶|巴哈|巴哈姆特|anime|Anime|ａｎｉｍｅ|Ａｎｉｍｅ|アニメ)(\\s|　|:|;|：|；)([\u30a0-\u30ff\u3040-\u309f\u4e00-\u9fd5_a-zA-Z0-9]*)\\D*([0-9.]{1,})") //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//2017.01.12+ 1A 1B 這種話數
	//reg := regexp.MustCompile("^(動畫|動畫瘋|懶|巴哈|巴哈姆特|anime|Anime|ａｎｉｍｅ|Ａｎｉｍｅ|アニメ)(\\s|　|:|;|：|；)([\u30a0-\u30ff\u3040-\u309f\u4e00-\u9fd5_a-zA-Z0-9]*)\\D*([0-9.A-Za-z]{1,})") //fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//2017.02.10+ 動畫名稱+捕捉空格...(沒關係反正後面的東西不是數字就可以對到)
	//reg := regexp.MustCompile("^(動畫|動畫瘋|懶|巴哈|巴哈姆特|anime|Anime|ａｎｉｍｅ|Ａｎｉｍｅ|アニメ)(\\s|　|:|;|：|；)([\u30a0-\u30ff\u3040-\u309f\u4e00-\u9fd5_a-zA-Z0-9\\s]*)\\D*([0-9.A-Za-z]{1,})") //fmt.Printf("%q\n", reg.FindAllString(text, -1))	
	reg := regexp.MustCompile("^(動畫|動畫瘋|懶|巴哈|巴哈姆特|anime|Anime|ａｎｉｍｅ|Ａｎｉｍｅ|アニメ)(\\s|　|:|;|：|；)([\u30a0-\u30ff\u3040-\u309f\u4e00-\u9fd5_a-zA-Z0-9\\s#！]*)(\\s|　|:|;|：|；)([0-9.A-Za-z]{1,})") //fmt.Printf("%q\n", reg.FindAllString(text, -1))	

	log.Print("--抓取分析觀察--")
	log.Print(reg.ReplaceAllString(text, "$1"))
	log.Print(reg.ReplaceAllString(text, "$2"))
	log.Print(reg.ReplaceAllString(text, "$3"))
	log.Print(reg.ReplaceAllString(text, "$4"))
	log.Print(reg.ReplaceAllString(text, "$5"))
	log.Print(reg.ReplaceAllString(text, "--抓取分析結束--"))
	
	switch reg.ReplaceAllString(text, "$1"){
	case "動漫通","今日動漫通","動畫瘋答案","今天答案","動畫瘋問題","巴哈答案":
		log.Print("有走到動漫通")
		print_string = "今日動漫通"
		log.Print("print_string =" + print_string)
	case "臉書","FB","ＦＢ","Fb","Ｆｂ","fb","ｆｂ","FACEBOOK","ＦＡＣＥＢＯＯＫ","Facebook","Ｆａｃｅｂｏｏｋ","facebook","ｆａｃｅｂｏｏｋ":
		print_string = "臉書"		
	case "主選單","選單","簡介","教學","help","Help","Ｈｅｌｐ","ｈｅｌｐ","ＨＥＬＰ","HELP":
		print_string = "選單"
	case "動畫瘋88":
		print_string = "動畫瘋88"
	case "test","測試":
		print_string = "測試"
	case "新番":
		print_string = "近期日本新開播的動畫（2016 十月、2017 一月）：\n" + 
		"\n" + 
		"【2016 十月新番】\n" + 
		"\n" + 
		"Yuri on Ice：Yuri、yuri\n" +
		"歌之☆王子殿下♪ 真愛 LEGEND STAR\n" +
		"夏目友人帳 伍\n" +
		"長騎美眉\n" +
		"3 月的獅子\n" +
		"黑白來看守所\n" +
		"我太受歡迎了該怎麼辦\n" +
		"無畏魔女\n" + 
		"殺老師 Q（2016 11 月）\n" + 
		"\n" + 
		"【2017 一月新番】\n" + 
		"\n" + 
		"幼女戰記\n" + 
		"青之驅魔師 京都不淨王篇\n" + 
		"亞人醬有話要說\n" +
		"南鎌倉高校女子自行車社：女子自行車\n" +
		"Hand Shakers\n" +
		"為美好的世界獻上祝福 第二季：美好的世界、獻上祝福\n" +
		"ACCA13 區監察課"
	case "bot","機器人","目錄","動畫清單","清單","索引","ｉｎｄｅｘ","index","Index","介紹","動漫","動畫介紹","動漫介紹","info","Info","ｉｎｆｏ":
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
		"釣球\n" +
		"進擊的巨人：巨人、進擊\n" +
		"Re：從零開始的異世界生活：異世界生活、從零\n" +
		"線上遊戲的老婆不可能是女生：不可能是女生\n" +
		"Thunderbolt Fantasy 東離劍遊紀：東離\n" +
		"來自風平浪靜的明日：風平浪靜\n" +
		"在那個夏天等待：那個夏天\n" +
		"羅馬浴場 THERMAE ROMAE：羅馬浴場、浴場\n" +
		"殺老師 Q\n" +
		"發條精靈戰記 天鏡的極北之星：發條精靈\n" +
		"男子啦啦隊：男子啦啦\n" +
		"月歌\n" +
		"藍海少女\n" + 
		"偶像大師 灰姑娘女孩：灰姑娘、灰姑娘女孩\n" +
		"乒乓\n" + 
		"幼女戰記：幼女\n" + 
		"花舞少女：花舞\n" +
		"夏雪之約：夏雪\n" +
		"混沌之子\n" + 
		"亞人醬有話要說：亞人、亞人醬、有話、有話要說\n" + 
		"南鎌倉高校女子自行車社：女子自行車、南鎌倉自行車、南鎌倉高校\n" +
		"青之驅魔師 京都不淨王篇：青之、不淨王\n" +
		"Hand Shakers\n" +
		"MARGINAL#4\n" +
		"終物語\n" +
		"BanG Dream：Bang、BanG、bang\n" + 
		"Yuri on Ice：Yuri、yuri\n" +
		"白銀的意志：白銀\n" +
		"坂道上的阿波羅：阿波羅\n" +
		"槍與假面舞會：假面舞會\n" +
		"為美好的世界獻上祝福 第二季：美好的世界、獻上祝福\n" +
		"魔物娘的同居日常：魔物娘\n" +
		"ACCA13 區監察課：監察課、ACCA13\n" +
		"史上最強弟子兼一：兼一\n" +
		"青年黑傑克：黑傑克、BJ、少年黑傑克\n" +
		"獵人 Hunter x Hunter：獵人、Hunter、hunter、hunter x hunter、Hunter x Hunter\n\n" +
		"搜尋方法：\n動畫 動畫名(或短名) 數字\n三個項目中間要用空白或冒號、分號隔開。\n\n例如：\n巴哈姆特　3月　１１\n動畫瘋　我太受歡迎 １\nアニメ;影子籃球員;15\n動畫 雙星 1\nanime：黑白來：5\n\n都可以"
	case "開發者","admin","Admin","ａｄｍｉｎ":
		print_string = "你找我主人？OK！\n我跟你講我的夥伴喵在哪，你去加他。\n他跟主人很親近的，跟他說的話主人都會看到。\nhttps://line.me/R/ti/p/%40uwk0684z\n\n\n你也可以從下面這個連結直接去找主人線上對話。\n\n如果他不在線上一樣可以留言給他，\n他會收到的！\n這跟手機、電腦桌面軟體都有同步連線。" +
		"\n\nhttp://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + user_msgid +
		"&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A"
	case "動畫", "動畫瘋","巴哈","巴哈姆特", "anime", "アニメ","Anime","ａｎｉｍｅ","Ａｎｉｍｅ","懶":
		print_string = text + "？\n好像有這個動畫耶，但我找不到詳細的QQ\n有可能你查詢的集數超前，查小一點的數字吧！\n或是你要手動去「巴哈姆特動畫瘋」找找嗎？\n\nhttps://ani.gamer.com.tw"
		anime_say := "有喔！有喔！你在找這個對吧！？\n"
		log.Print(reg.ReplaceAllString(text, "$3"))
		switch reg.ReplaceAllString(text, "$3") {
		case "青年黑傑克","黑傑克","BJ","少年黑傑克":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7232"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7233"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7234"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7235"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7236"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7237"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7238"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7239"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7240"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7241"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7242"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7243" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "史上最強弟子兼一","兼一":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1191"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1208"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1161"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1162"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1163"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1164"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1165"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1166"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1167"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1168"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1169"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1170"
			case "13":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1171"
			case "14":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1172"
			case "15":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1173"
			case "16":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1174"
			case "17":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1175"
			case "18":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1209"
			case "19":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1176"
			case "20":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1177"
			case "21":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1178"
			case "22":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1179"
			case "23":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1180"
			case "24":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1181"
			case "25":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1182"
			case "26":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1183"
			case "27":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1184"
			case "28":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1185"
			case "29":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1186"
			case "30":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1187"
			case "31":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1188"
			case "32":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1189"
			case "33":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1190"
			case "34":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1210"
			case "35":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1192"
			case "36":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1193"
			case "37":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1194"
			case "38":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1195"
			case "39":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1196"
			case "40":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1197"
			case "41":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1198"
			case "42":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1199"
			case "43":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1200"
			case "44":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1201"
			case "45":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1202"
			case "46":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1203"
			case "47":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1204"
			case "48":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1205"
			case "49":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1206"
			case "50":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=1207" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "獵人 Hunter x Hunter","獵人","Hunter","hunter","hunter x hunter","Hunter x Hunter":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5575"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5576"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5577"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5578"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5579"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5580"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5581"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5582"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5583"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5584"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5585"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5586"
			case "13":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5587"
			case "14":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5588"
			case "15":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5589"
			case "16":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5590"
			case "17":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5591"
			case "18":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5592"
			case "19":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5593"
			case "20":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5594"
			case "21":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5595"
			case "22":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5596"
			case "23":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5597"
			case "24":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5598"
			case "25":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5599"
			case "26":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5600"
			case "27":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5601"
			case "28":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5602"
			case "29":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5603"
			case "30":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5604"
			case "31":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5605"
			case "32":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5606"
			case "33":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5607"
			case "34":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5608"
			case "35":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5609"
			case "36":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5610"
			case "37":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5611"
			case "38":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5612"
			case "39":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5613"
			case "40":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5614"
			case "41":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5615"
			case "42":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5616"
			case "43":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5617"
			case "44":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5618"
			case "45":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5619"
			case "46":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5620"
			case "47":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5621"
			case "48":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5622"
			case "49":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5623"
			case "50":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5624"
			case "51":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5625"
			case "52":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5626"
			case "53":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5627"
			case "54":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5628"
			case "55":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5629"
			case "56":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5630"
			case "57":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5631"
			case "58":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5632"
			case "59":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5633"
			case "60":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5634"
			case "61":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5635"
			case "62":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5636"
			case "63":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5637"
			case "64":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5638"
			case "65":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5639"
			case "66":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5640"
			case "67":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5641"
			case "68":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5642"
			case "69":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5643"
			case "70":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5644"
			case "71":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5645"
			case "72":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5646"
			case "73":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5647"
			case "74":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5648"
			case "75":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5534"
			case "76":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5535"
			case "77":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5536"
			case "78":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5537"
			case "79":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5538"
			case "80":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5539"
			case "81":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5540"
			case "82":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5541"
			case "83":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5542"
			case "84":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5543"
			case "85":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5544"
			case "86":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5545"
			case "87":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5546"
			case "88":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5547"
			case "89":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5548"
			case "90":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5549"
			case "91":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5550"
			case "92":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5551"
			case "93":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5552"
			case "94":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5553"
			case "95":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5554"
			case "96":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5555"
			case "97":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5556"
			case "98":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5557"
			case "99":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5558"
			case "100":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5559"
			case "101":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5560"
			case "102":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5561"
			case "103":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5562"
			case "104":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5563"
			case "105":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5564"
			case "106":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5565"
			case "107":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5566"
			case "108":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5567"
			case "109":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5568"
			case "110":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5569"
			case "111":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5570"
			case "112":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5571"
			case "113":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5572"
			case "114":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5573"
			case "115":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5574"
			case "116":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5660"
			case "117":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5661"
			case "118":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5662"
			case "119":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5663"
			case "120":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5664"
			case "121":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5665"
			case "122":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5666"
			case "123":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5667"
			case "124":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5668"
			case "125":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5669"
			case "126":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5670"
			case "127":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5671"
			case "128":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5672"
			case "129":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5673"
			case "130":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5674"
			case "131":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5675"
			case "132":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5676"
			case "133":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5677"
			case "134":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5678"
			case "135":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5649"
			case "136":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5650"
			case "137":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5651"
			case "138":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5679"
			case "139":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5680"
			case "140":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5681"
			case "141":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5652"
			case "142":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5653"
			case "143":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5654"
			case "144":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5655"
			case "145":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5656"
			case "146":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5657"
			case "147":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5658"
			case "148":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5659" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "魔物娘的同居日常","魔物娘":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7244"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7245"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7246"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7247"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7248"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7249"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7250"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7251"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7252"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7253"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7254"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7255"
			default:
			}
		case "坂道上的阿波羅","阿波羅":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6948"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6949"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6950"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6951"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6952"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6953"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6954"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6955"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6956"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6957"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6958"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6959" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "白銀的意志","白銀":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7184"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7185"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7186"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7187"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7188"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7189"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7190"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7191"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7192"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7193"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7194"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7195"
			case "13":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7196"
			case "14":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7197"
			case "15":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7198"
			case "16":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7199"
			case "17":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7200"
			case "18":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7201"
			case "19":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7202"
			case "20":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7203"
			case "21":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7204"
			case "22":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7205"
			case "23":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7206"
			case "24":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7207" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "槍與假面舞會","假面舞會":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7172"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7176"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7177"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7178"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7179"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7180"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7181"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7182"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7183"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7173"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7174"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7175" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "BanG Dream","Bang","BanG","bang":
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7147"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7321"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7322"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7395"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7396"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7397"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7398"
			default:
			}
		//2016.10?
		case "Yuri on Ice","Yuri","yuri","ｙｕｒｉ","Ｙｕｒｉ","YURI!!!on ICE":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7265"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7266"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7293"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7329"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7344"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7345"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7346"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7347"
			default:
			}
		case "終物語":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7220"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7224"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7225"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7226"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7227"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7228"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7229"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7230"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7231"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7221"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7222"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7223" + "\n\n等等！這是最後一話！？"
			default:
			}
		//2017
		case "MARGINAL","MARGINAL#","MARGINAL#4":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7158"
			case "2","02":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7282"
			case "3","03":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7313"
			case "4","04":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7314"
			case "5","05":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7379"
			case "6","06":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7380"
			case "7","07":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7381"
			case "8","08":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7382"
			default:
			}
		//2017
		case "混沌之子":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01","1A":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7156"
			case "1.5","1B":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7264"
			case "2","02":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7278"
			case "3","03":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7279"
			case "4","04":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7311"
			case "5","05":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7371"
			case "6","06":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7372"
			case "7","07":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7373"
			case "8","08":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7374"
			default:
			}
		//2017
		case "為美好的世界獻上祝福！ 第二季","為美好的世界獻上祝福","獻上祝福！ 第二季","美好的世界":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
				case "1","01":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7157"
				case "2","02":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7280"
				case "3","03":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7281"
				case "4","04":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7312"
				case "5","05":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7375"
				case "6","06":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7376"
				case "7","07":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7377"
				case "8","08":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7378"
			default:
			}
		case "夏雪之約","夏雪":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7116"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7117"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7118"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7119"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7120"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7121"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7122"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7123"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7124"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7125"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7126" + "\n\n等等！這是最後一話！？"
			default:
			}
		//2017.01
		case "Hand Shakers","hand shakers","hand","Hand Shakers ":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7155"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7275"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7276"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7308"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7359"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7360"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7361"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7362"
			default:
			}
		//2017.01
		case "ACCA13 區監察課","監察課","ACCA","ACCA13","Acca13","Ａｃｃａ13","ＡＣＣＡ13":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7154"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7273"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7274"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7307"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7355"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7356"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7357"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7362"
			default:
			}
		//2017.01
		case "南鎌倉高校女子自行車社","南鎌倉自行車","南鎌倉高校","女子自行車":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7127"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7267"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7268"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7304"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7348"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7349"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7350"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7351"
			default:
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7127"
			}
		case "花舞少女","花舞":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7127"
			case "2","02":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7128"
			case "3","03":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7129"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7130"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7131"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7132"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7133"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7134"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7135"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7136"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7137"
			case "12":
					print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7138" + "\n\n等等！這是最後一話！？"
			default:
			}
		//2017.01
		case "亞人醬有話要說","亞人","亞人醬","有話","有話要說":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7151"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7260"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7287"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7325"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7326"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7403"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7404"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7405"
			default:
			}
		//2017.01
		case "青之驅魔師","青之驅魔師 京都不淨王篇","青之","不淨王":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3524\n\n京都不淨王篇\nhttp://ani.gamer.com.tw/animeVideo.php?sn=7149"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3525\n\n京都不淨王篇\nhttp://ani.gamer.com.tw/animeVideo.php?sn=7258"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3526\n\n京都不淨王篇\nhttp://ani.gamer.com.tw/animeVideo.php?sn=7285"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3527\n\n京都不淨王篇\nhttp://ani.gamer.com.tw/animeVideo.php?sn=7319"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3528\n\n京都不淨王篇\nhttp://ani.gamer.com.tw/animeVideo.php?sn=7320"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3529\n\n京都不淨王篇\nhttp://ani.gamer.com.tw/animeVideo.php?sn=7391"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3530\n\n京都不淨王篇\nhttp://ani.gamer.com.tw/animeVideo.php?sn=7392"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3531\n\n京都不淨王篇\nhttp://ani.gamer.com.tw/animeVideo.php?sn=7393"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3532\n\n京都不淨王篇\nhttp://ani.gamer.com.tw/animeVideo.php?sn=7394"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3533"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3534"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3535"
			case "13":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3536"
			case "14":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3537"
			case "15":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3538"
			case "16":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3539"
			case "17":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3540"
			case "18":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3541"
			case "19":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3542"
			case "20":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3543"
			case "21":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3544"
			case "22":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3545"
			case "23":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3546"
			case "24":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3547"
			case "25":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3548"
			case "movie":
				print_string = anime_say + "http://ani.gamer.com.tw/animeRef.php?sn=52194"
			default:
				print_string = anime_say + "http://ani.gamer.com.tw/animeRef.php?sn=52194"
			}
		//2017.01
		case "幼女戰記","幼女":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7148"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7257"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7284"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7317"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7318"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7387"
			case "6.5","06.5":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7388"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7389"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7390"
			default:
			}
		case "偶像大師 灰姑娘女孩","灰姑娘","灰姑娘女孩":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7091"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7092"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7093"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7094"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7095"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7096"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7097"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7098"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7099"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7100"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7101"
			case "12":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7102"
			case "13":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7103"
			case "14":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7104"
			case "15":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7105"
			case "16":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7106"
			case "17":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7107"
			case "18":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7108"
			case "19":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7109"
			case "20":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7110"
			case "21":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7111"
			case "22":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7112"
			case "23":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7113"
			case "24":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7114"
			case "25":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7115" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "藍海少女":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5867"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5914"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5915"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6127"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6128"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6129"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6209"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6210"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6323"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6324"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6325"
			case "12":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6326" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "月歌。","月歌":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5792"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5921"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5922"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6023"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6024"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6025"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6173"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6174"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6303"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6304"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6305"
			case "12":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6306"
			case "13":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6399" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "男子啦啦隊","男子啦啦","男子拉拉隊","男子拉拉":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5860"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5861"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5885"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5886"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5887"
			case "5.5":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5888"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6181"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6182"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6183"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6310"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6311"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6312"
			case "12":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6313" + "\n\n等等！這是最後一話！？"
			default:
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5888"
			}
		case "發條精靈戰記 天鏡的極北之星","發條精靈","精靈戰記":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5872"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5910"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5911"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6034"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6035"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6036"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6203"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6204"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6321"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6322"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6345"
			case "12":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6346"
			case "13":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6437" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "乒乓":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7058"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7059"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7060"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7061"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7062"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7063"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7064"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7065"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7066"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7067"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7068"
			default:
			}
		case "鎖鏈戰記 赫克瑟塔斯之光","鎖鏈戰記","赫克瑟塔斯之光":
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7027"
			case "2","02":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7292"
			case "3","03":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7445"
			default:
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7027"
			}
		case "在那個夏天等待","那個夏天":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4196"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4200"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4201"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4202"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4203"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4204"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4205"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4206"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4207"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4197"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4198"
			case "12":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4199" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "來自風平浪靜的明日","風平浪靜","風平浪靜的明日","浪靜的明日":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=196"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=197"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=198"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=199"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=200"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=201"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=202"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=203"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=204"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=205"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=206"
			case "12":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=207"
			case "13":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=208"
			case "14":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=209"
			case "15":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=210"
			case "16":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=211"
			case "17":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=212"
			case "18":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=213"
			case "19":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=214"
			case "20":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=215"
			case "21":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=216"
			case "22":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=217"
			case "23":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=218"
			case "24":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=219"
			case "25":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=220"
			case "26":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=221" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "羅馬浴場 THERMAE ROMAE","羅馬浴場","浴場":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6987"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6988"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6989"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6990"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6991"
			default:
			}
		case "Thunderbolt Fantasy 東離劍遊紀","東離劍遊紀","東離","Thunderbolt Fantasy","Thunderbolt":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5884"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6001"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6037"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6196"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6197"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6198"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6267"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6268"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6334"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6335"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6336"
			case "12":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6337"
			case "13":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6472" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "Re：從零開始的異世界生活","從零","異世界生活","re","Re":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4996"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5003"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5025"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5026"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5106"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5107"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5159"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5178"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5179"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5505"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5506"
			case "12":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5507"
			case "13":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5777"
			case "14":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5787"
			case "15":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5791"
			case "16":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5870"
			case "17":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5902"
			case "18":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5903"
			case "19":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6146"
			case "20":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6147"
			case "21":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6188"
			case "22":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6189"
			case "23":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6317"
			case "24":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6318"
			case "25":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6355" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "線上遊戲的老婆不可能是女生","線上遊戲的老婆不可能是女生？","老婆不可能","線上遊戲的老婆","不可能是女生":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5012"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5030"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5031"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5115"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5116"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5117"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5199"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5200"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5201"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5499"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5500"
			case "12":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5501" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "進擊的巨人","進擊","巨人":
		   //reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3499"
			case "2","02":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3500"
			case "3","03":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3514"
			case "4","04":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3515"
			case "5","05":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3501"
			case "6","06":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3502"
			case "7","07":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3503"
			case "8","08":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3504"
			case "9","09":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3505"
			case "10":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3516"
			case "11":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3517"
			case "12":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3506"
			case "13":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3507"
			case "14":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3518"
			case "15":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3508"
			case "16":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3519"
			case "17":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3509"
			case "18":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3510"
			case "19":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3520"
			case "20":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3521"
			case "21":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3511"
			case "22":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3512"
			case "23":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3522"
			case "24":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3523"
			case "25":
					print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3513" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "殺老師","殺老師 Q","殺老師Q","殺老師 Q！":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7057"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7141"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7152"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7256"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7283"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7315"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7316"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7383"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7384"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7385"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7386"
			default:
			}
		case "路人超能 100","路人","靈能":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5863"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5881"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5893"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5894"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5895"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5896"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6184"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6185"
			case "9","09":
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
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6992"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6993"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6994"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6995"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6996"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6997"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6998"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6999"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7000"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7001"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7002"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7003" + "\n\n等等！這是最後一話！？"
			default:
			}
		//還沒跟其他部合併
		case "影子籃球員","影子籃球","影子籃":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3896"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3897"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3898"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3899"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3900"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3901"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3902"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=3903"
			case "9","09":
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
			case "movie":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7054" //影子籃球員 冬季選拔賽總集篇 門的彼端 [電影]
			default:
			}
		case "文豪野犬 第二季","文豪野犬","文豪","野犬":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6130"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6131"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6132"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6133"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6134"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6135"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6136"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6137"
			case "9","09":
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
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2543"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2544"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2545"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2546"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2547"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2548"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2549"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=2550"
			case "9","09":
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
			case "38":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6902"
			case "39":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7261"
			case "40":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7288"
			case "41":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7327"
			case "42":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7328"
			case "43":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7407"
			case "44":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7408"
			case "45":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7409"
			default:
			}
		case "星夢手記","星夢":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6769"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6770"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6771"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6772"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6773"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6774"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6784"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6903"
			case "9","09":
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
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4999"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5007"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5008"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5120"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5168"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5175"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5176"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5473"
			case "9","09":
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
			case "26":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6835"
			case "27":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7056"
			case "28":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7271"
			case "29":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7272"
			case "30":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7306"
			case "31":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7352"
			case "32":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7353"
			default:
			}
		case "神裝少女小纏","小纏","神裝少女","神裝":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6406"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6431"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6432"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6433"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6434"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6435"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6729"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6730"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6731"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6836"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6837"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7025"
			case "13":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7139"
			case "特別","SP","OVA":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6469"
			default:
			}
		case "夏目友人帳 伍","夏目友人帳","夏目","有人","有人帳","帳":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6425"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6426"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6427"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6428"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6429"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6430"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6845"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6733"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6734"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6838"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6839"
			case "特別","SP","OVA":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6732"
			default:
			}
		case "黑白來看守所","黑白來","看守所","黑白":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6473"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6474"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6475"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6476"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6477"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6478"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6735"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6736"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6737"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6840"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6841"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6842"
			case "13":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7055"
			case "14":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7144"
			case "15":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7145"
			case "16":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7146"
			case "17":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7277"
			case "18":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7309"
			case "19":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7363"
			case "20":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7364"
			case "21":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7365"
			case "22":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7366"
			default:
			}
		case "喵阿愣！","喵阿愣","喵啊愣！","阿愣","啊愣","喵阿楞！","喵阿楞":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6850"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6851"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6852"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6853"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6854"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6855"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6856"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6857"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6917"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6918"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6919"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6920"
			case "13":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=6921"
			case "14":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=6922"
			case "15":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7143"
			case "16":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7269"
			case "17":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7270"
			case "18":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7305"
			case "19":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7413"
			case "20":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7414"
			case "21":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7415"
			case "22":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7416"
			default:
			}
		case "雙星之陰陽師","雙星":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4998"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5027"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5028"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5029"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5110"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5111"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5169"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5470"
			case "9","09":
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
			case "37":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7023"
			case "38":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7023"
			case "39":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7262"
			case "40":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7289"
			case "41":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7290"
			case "42":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7310"
			case "43":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7367"
			case "44":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7368"
			case "45":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7369"
			case "46":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7370"
			default:
			}
		case "無畏魔女":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6419"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6420"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6421"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6422"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6423"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6424"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6766"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6767"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6848"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6849"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7021"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7022"
			default:
			}
		case "伯納德小姐說","小姐說","伯納德","伯納":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			default:
				print_string = "可能不受歡迎或其他原因，\n很遺憾這部已經下架，請幫QQ"
			}
		case "漂流武士":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6485"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6486"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6487"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6488"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6489"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6490"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6871"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6872"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6873"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6874"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6875"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6876" + "\n\n等等！這是最後一話！？"
			default:
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6485"
			}
		//還沒合其他部
		case "JOJO 的奇妙冒險 不滅鑽石","Jojo","jojo","JOJO","JOJO的奇妙冒險","奇妙冒險":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=4994"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5005"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5019"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5020"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5096"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5097"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5165"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5191"
			case "9","09":
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
			case "39":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6883" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "長騎美眉","長騎","單車":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6407"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6408"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6409"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6410"
			case "4.5":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6411"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6412"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6884"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6885"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6886"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6887"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6888"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7342"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=7343"
			default:
			}
		case "我太受歡迎了該怎麼辦","我太受歡迎","受歡迎":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6413"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6414"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6415"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6416"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6417"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6418"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6865"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6866"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6867"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6868"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6869"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6870" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "3月","3 月的獅子","3月的獅子","三月的獅子","三月":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6479"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6480"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6481"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6482"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6483"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6484"
			case "7","07":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6889"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6890"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6891"
			case "10":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=6892"
			case "11":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=6893"
			case "11.5":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7142"
			case "12":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7150"
			case "13":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7259"
			case "14":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7286"
			case "15":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7323"
			case "16":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7324"
			case "17":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7399"
			case "18":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7400"
			case "19":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=7401"
			default:
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6479"
			}
		case "美術社","美術社大有問題","這個美術社大有問題":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=5871"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5918"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=5919"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6038"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6039"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6040"
			case "7","07":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=6207"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6208"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6295"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6296"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6297"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6298" + "\n\n等等！這是最後一話！？"
			default:
			}
		case "歌之☆王子殿下♪ 真愛","歌王子","uta","哥王子":
			//reg.ReplaceAllString(text, "$2")
			switch reg.ReplaceAllString(text, "$5") {
			case "1","01":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6436" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2068\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2055\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2042"
			case "2","02":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6470" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2069\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2056\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2043"
			case "3","03":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6496" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2070\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2057\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2044"
			case "4","04":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6497" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2071\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2058\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2045"
			case "5","05":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6498" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2072\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2059\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2046"
			case "6","06":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6499" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2073\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2060\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2047"
			case "7","07":
				print_string = anime_say + "https://ani.gamer.com.tw/animeVideo.php?sn=6724" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2074\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2061\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2048"
			case "8","08":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6782" +
				"\n\n上面查到的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2075\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2062\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2049"
			case "9","09":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6783" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2076\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2063\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2050"
			case "10":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6895" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2070\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2064\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2051"
			case "11":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6896" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2078\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2065\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2052"
			case "12":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6897" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2079\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2066\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2053"
			case "13":
				print_string = anime_say + "http://ani.gamer.com.tw/animeVideo.php?sn=6898" + 
				"\n\n上面查的這是第四部！日本 2016 年十月才開播。\n還是你要看其他的呢？\n\n" +
				"第三部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2080\n\n" + 
				"第二部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2067\n\n" + 
				"第一部：\nhttp://ani.gamer.com.tw/animeVideo.php?sn=2054" + "\n\n等等！這是最後一話！？"
			default:
			}
		default:
			print_string = "你是要找 " +  reg.ReplaceAllString(text, "$3") + " 對嗎？\n對不起，我找不到這部動畫，我還沒學呢...\n（可輸入「目錄」查看支援的作品）\n我目前知道的動畫還很少，因為我考試不及格QAQ\n\n（其實是因為開發者半手動輸入更新，沒用自動化爬蟲跟資料庫。才會增加比較慢XD）"
		}
	default:
		if reply_mode!="" {
			print_string = "HI～ 我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這只提供查詢動畫的功能。\n如有其他建議或想討論，請對我輸入「開發者」進行聯絡。"
		} else {
			print_string = "" //安靜模式
		}
	}
	return print_string
}

//http://qiita.com/koki_cheese/items/66980888d7e8755d01ec
// func handleTask(w http.ResponseWriter, r *http.Request) {
// }

	//修改時主要參考官方文件以及：
	// https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
		// KEY = handleText 等
	// https://github.com/dongri/line-bot-sdk-go
		// KEY = linebot.NewAudioMessage(originalContentURL, duration)
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// allow cross domain AJAX requests
	// http://stackoverflow.com/questions/12830095/setting-http-headers-in-golang/
	//	https://developer.mozilla.org/ja/docs/Web/HTTP/HTTP_access_control
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
	//http://qiita.com/futosu/items/b49f7d9e28101daaa99e
	//https://play.golang.org/p/xHp44c_pJm
	w.Header().Set("Access-Control-Allow-Headers","Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// https://groups.google.com/forum/#!topic/golang-nuts/-Sh616lXNRE

	//-----------------------------------------------

	// log.Print("r")
	// log.Print(r)

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

		//2016.12.23+ 統一基本資訊集中
		//2016.12.24+ 嘗試抓使用者資訊 https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
		target_id_code := event.Source.UserID + event.Source.GroupID + event.Source.RoomID//target_id_code := ""
 		log.Print("event.Source.UserID = " + event.Source.UserID)
		log.Print("event.Source.GroupID = " + event.Source.GroupID)
		log.Print("event.Source.RoomID = " + event.Source.RoomID)
		log.Print("target_id_code = " + target_id_code)
		target_item := ""
		if event.Source.UserID!="" {
			target_item = "好友"
		}
		if event.Source.GroupID!="" {
			target_item = "群組對話"
		}
		if event.Source.RoomID!="" {
			target_item = "房間"
		}
		log.Print("target_item = " + target_item)

		username := ""
		userStatus := ""
		userImageUrl := ""
																				//userLogo_url := ""
		switch target_id_code{
			case "U6f738a70b63c5900aa2c0cbbe0af91c4":
				username = "懶懶"
			case "Uf150a9f2763f5c6e18ce4d706681af7f":
				username = "包包"
			case "Ca78bf89fa33b777e54b4c13695818f81":
				username = "測試用全開群組 test"
			case "C717159d4582434c603de3cad7e0b4373":
				username = "跟ㄅㄅ測試的群組"
			case "Cf9842427f0517899f9e3607f15be25c1":
				username ="白白測試群組"
		}
		log.Print("username = " + username)

		//如果是群組會出錯，只能 1 對 1的時候。
		//if target_item == "好友"{
		if event.Source.UserID!="" {
			//2016.12.24+ 嘗試抓使用者資訊 https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
			profile, err := bot.GetProfile(event.Source.UserID).Do()
			if err != nil {
				log.Print(1162)
				log.Print(err)
			}
			log.Print("profile.DisplayName = " + profile.DisplayName)			// println(res.Displayname)
			log.Print("profile.StatusMessage " + profile.StatusMessage)			// println(res.StatusMessage)
			log.Print("profile.PictureURL = " + profile.PictureURL)

														// log.Print("userLogo_url = " +  userLogo_url)
			//如果不是認識的 ID，就取得對方的名
			if username == ""{
				username = profile.DisplayName
			}
			userStatus = profile.StatusMessage
			userImageUrl = profile.PictureURL

			log.Print("username = " + username)
			log.Print("userStatus = " + userStatus)
			log.Print("userImageUrl = " + userImageUrl)

		}

		user_talk := ""
		if username == ""{
			user_talk = "【" + target_item + "】 " + target_id_code
		}else{
			user_talk = username
		}
		log.Print("※ user_talk = " + user_talk)

		//2016.12.27+

		SystemImageURL := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/c2704b19816673a30c76cdccf67bcf8f/2016_-_%E8%A4%87%E8%A3%BD.png"
		imageURL := SystemImageURL

		//共用模板
		LineTemplate_chat := linebot.NewURITemplateAction("線上與開發者聊天", "http://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + target_id_code + "&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A")
		LineTemplate_addme := linebot.NewURITemplateAction("加開發者 LINE", "https://line.me/R/ti/p/@uwk0684z")
		LineTemplate_download_app := linebot.NewURITemplateAction("下載巴哈姆特動畫瘋 APP", "https://prj.gamer.com.tw/app2u/animeapp.html")

		LineTemplate_feedback := linebot.NewCarouselColumn(
			imageURL, "意見反饋 feedback", "你可以透過此功能\n對 開發者 提出建議",
			LineTemplate_addme,
			LineTemplate_chat,
			linebot.NewMessageTemplateAction("聯絡 LINE 機器人開發者", "開發者"),
		)

		LineTemplate_other := linebot.NewCarouselColumn(
			imageURL, "其他功能", "新番、可查詢的動畫清單",
			linebot.NewMessageTemplateAction("可查詢的動畫清單", "目錄"),
			linebot.NewMessageTemplateAction("新番、當季動畫", "新番"),
			linebot.NewMessageTemplateAction("今天動漫通答案", "今日動漫通"),
		)
		
		LineTemplate_other_example := linebot.NewCarouselColumn(
			imageURL, "其他使用例", "開頭可以是 動畫 / anime / アニメ / 巴哈姆特",
			linebot.NewMessageTemplateAction("巴哈姆特 三月 ３", "巴哈姆特 三月 ３"),
			linebot.NewMessageTemplateAction("Ａｎｉｍｅ　喵阿愣　５", "Ａｎｉｍｅ　喵阿愣　５"),
			linebot.NewMessageTemplateAction("anime：黑白來：7", "anime：黑白來：7"),
		)

		LineTemplate_firstinfo := linebot.NewCarouselTemplate(
			linebot.NewCarouselColumn(
				imageURL, "查詢巴哈姆特動畫瘋的功能", "我很愛看巴哈姆特動畫瘋。\n問我動畫可以這樣問：動畫 動畫名稱 集數",
				linebot.NewPostbackTemplateAction("動畫 美術社 12","測試 POST", "動畫 美術社 12"),
				linebot.NewMessageTemplateAction("アニメ 美術社大有問題 12", "アニメ 美術社大有問題 12"),
				linebot.NewMessageTemplateAction("anime：美術社：１", "anime：美術社：１"),
			),
			LineTemplate_other_example,
			LineTemplate_other,
			LineTemplate_feedback,
		)

							fb_msg := "\n\n答案請上 FB 查詢大家意見。\n" + "巴哈姆特動畫瘋 FB：\nhttps://www.facebook.com/animategamer/posts/1343665632362471"
							fb_q_msg := "2017/03/05 動漫通\n" +
"關聯：魔法禁書目錄\n" +
"問題：請問 學園城市 的火車種類是什麼？\n" +
"1.雲霄飛車\n" +
"2.五分車\n" +
"3.磁浮列車\n" +
"4.懸吊列車\n" +
"小提示：利用磁力運作的\n" +
"出題者：AVA2011\n" +
								fb_msg

							LineTemplate_today_q := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "2017/03/05 動漫通", "答案請參考 FB 討論，可能是 3 ",
									linebot.NewURITemplateAction("來看 FB 上的答案！","https://www.facebook.com/animategamer/posts/1343665632362471"),
									linebot.NewURITemplateAction("巴哈姆特動畫瘋 官網","http://ani.gamer.com.tw"),
									LineTemplate_download_app,
								),
								LineTemplate_feedback,
								LineTemplate_other,
								LineTemplate_other_example,
							)

		//正題

		//只會抓到透過按鈕按下去的東西。方便做新的觸發點。(缺點是沒有 UI 介面的時候會無法使用)
		if event.Type == linebot.EventTypePostback {
				//這裡用來設計按下某按鈕後要做什麼事情
				log.Print("觸發 Postback 功能（不讓使用者察覺的程式利用）")
				log.Print("event.Postback.Data = " + event.Postback.Data)
				HttpPost_JANDI("[" + user_talk + "](" + userImageUrl + ") 觸發了按鈕並呼了 event.Postback.Data = " + event.Postback.Data + `\n` + userStatus, "brown" , "LINE 程式觀察",target_id_code)
				HttpPost_IFTTT(user_talk + " 觸發了按鈕並呼了 event.Postback.Data = " + event.Postback.Data + `\n<br>` + userImageUrl + `\n<br>` + userStatus , "LINE 程式觀察" ,target_id_code)
				HttpPost_Zapier("[" + user_talk + "](" + userImageUrl + ") 觸發了按鈕並呼了 event.Postback.Data = " + event.Postback.Data + `\n` + userStatus, "LINE 程式觀察" ,target_id_code,user_talk)

				// if event.Postback.Data == "測試"{

				// }

				if event.Postback.Data == "test"{


					// https://devdocs.line.me/en/#imagemap-message
					// "x": 0,
	 				//	"y": 0,
		   			// "width": 520,
		   			// "height": 1040

		   			//log.Print("test MD5 = " + GetMD5Hash(event.Postback.Data))

		   			//測試圖片地圖
					obj_message := linebot.NewImagemapMessage(
							"https://synr.github.io/test",
							"Imagemap alt text",
							linebot.ImagemapBaseSize{1040, 1040},
							linebot.NewURIImagemapAction("https://store.line.me/family/manga/en", linebot.ImagemapArea{0, 0, 520, 520}),
							linebot.NewURIImagemapAction("https://store.line.me/family/music/en", linebot.ImagemapArea{520, 0, 520, 520}),
							linebot.NewURIImagemapAction("https://store.line.me/family/play/en", linebot.ImagemapArea{0, 520, 520, 520}),
							linebot.NewMessageImagemapAction("URANAI!", linebot.ImagemapArea{520, 520, 520, 520}),	//上限 400 字
					)

					if _, err := bot.ReplyMessage(event.ReplyToken,obj_message).Do(); err != nil {
						log.Print(1586)
						log.Print(err)
					}
				}

				if event.Postback.Data == "line_to_me"{
					switch target_id_code {
						case "U6f738a70b63c5900aa2c0cbbe0af91c4":
							//by uid 一定要 .PushMessage （ event.ReplyToken 隨時在變化不可能可以。）
							log.Print("event.ReplyToken = " + event.ReplyToken)
							// if _, err = bot.PushMessage("U6f738a70b63c5900aa2c0cbbe0af91c4", linebot.NewTextMessage("event.ReplyToken = " + event.ReplyToken)).Do(); err != nil {
							// 		log.Print(1944)
							// 		log.Print(err)
							// }
							// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("發射成功")).Do(); err != nil {
							// 		log.Print(1944)
							// 		log.Print(err)
							// }
						default:
					}
				}

				if event.Postback.Data == "開啟管理者選單"{
					switch target_id_code {
						case "U6f738a70b63c5900aa2c0cbbe0af91c4":
							LineTemplate_test := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "管理模式", "測試中",
									linebot.NewPostbackTemplateAction("指定對象發 LINE","line_to_me", ""),
									linebot.NewPostbackTemplateAction("無效選項","admin", ""),
									linebot.NewPostbackTemplateAction("登出","登出管理者", ""),
								),
								// LineTemplate_other_example,
								// LineTemplate_other,
								LineTemplate_feedback,
							)
							no_temp_msg := "請更新使用最新版本的 LINE APP 才能查看內容 。"
							obj_message := linebot.NewTemplateMessage(no_temp_msg, LineTemplate_test)
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
									log.Print(2148)
									log.Print(err)
							}
						default:
					}
				}

				if event.Postback.Data == "passcheck"{
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("請輸入暗號驗證管理者身分")).Do(); err != nil {
							log.Print(1929)
							log.Print(err)
					}
				}

				//2017.01.03+
				if event.Postback.Data == "admin"{
					switch target_id_code {
						case "U6f738a70b63c5900aa2c0cbbe0af91c4":
							LineTemplate_test := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "管理模式", "For ADMIN mode.",
									linebot.NewPostbackTemplateAction("核對「暗號」","passcheck", ""),
									linebot.NewPostbackTemplateAction("管理模式","admin", ""),
									linebot.NewPostbackTemplateAction("測試","test", ""),
								),
								// LineTemplate_other_example,
								// LineTemplate_other,
								LineTemplate_feedback,
							)
							no_temp_msg := "請更新使用最新版本的 LINE APP 才能查看內容 。"
							obj_message := linebot.NewTemplateMessage(no_temp_msg, LineTemplate_test)
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
									log.Print(2148)
									log.Print(err)
							}
						default:
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您無法使用此功能。")).Do(); err != nil {
									log.Print(1955)
									log.Print(err)
							}
					}
				}

				if event.Postback.Data == "登出管理者"{
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你已登出管理模式")).Do(); err != nil {
						log.Print(1965)
						log.Print(err)
					}
				}

				if event.Postback.Data == "取消離開群組"{
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你已經取消請我離開 :)")).Do(); err != nil {
						log.Print(1207)
						log.Print(err)
					}
				}

				//2016.12.26+
				if event.Postback.Data == "按下確定離開群組對話"{
					template := linebot.NewCarouselTemplate(
						linebot.NewCarouselColumn(
							SystemImageURL, "請機器人離開群組", "你確定要請我離開嗎QAQ？\n如果確定請按下方按鈕 QQ",
							linebot.NewPostbackTemplateAction("請機器人離開群組","離開群組", "機器人已經自動離開。\n如要加回來請找：\nhttps://line.me/R/ti/p/@sjk2434l\n如要聯絡開發者請找：\nhttps://line.me/R/ti/p/@uwk0684z"),
							//linebot.NewPostbackTemplateAction("請機器人離開群組","離開群組", "機器人已經自動離開。\n如要加回來請找：\nhttps://line.me/R/ti/p/@sjk2434l\n如要聯絡開發者請找：\nhttps://line.me/R/ti/p/@uwk0684z"),
							LineTemplate_addme,
							LineTemplate_chat,
						),
					)
					obj_message := linebot.NewTemplateMessage("這是命令機器人自己離開群組的方法。\n這功能只支援 APP 使用。\n請用 APP 端查看下一步。", template)
					if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
						log.Print(1225)
						log.Print(err)
					}
				}

				if event.Postback.Data == "離開群組"{
					if target_item == "群組對話" {
						if _, err := bot.LeaveGroup(target_id_code).Do(); err != nil {
							log.Print(1233)
							log.Print(err)
						}
						HttpPost_JANDI("自動離開 "  + user_talk , "gray" , "LINE 離開群組",target_id_code)
						HttpPost_IFTTT("自動離開 "  + user_talk , "LINE 離開群組",target_id_code)
						HttpPost_Zapier("自動離開 "  + user_talk , "LINE 離開群組",target_id_code,user_talk)
						log.Print("觸發自動離開 " + user_talk +  " 群組")
					}
				}
		}
		//觸發加入好友
		if event.Type == linebot.EventTypeFollow {
				HttpPost_JANDI("有新的好朋友：["  + user_talk + "](" + userImageUrl  + ")" + `\n` + userStatus, "blue" , "LINE 新好友",target_id_code)
				HttpPost_IFTTT("有新的好朋友："  + user_talk  + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 新好友" ,target_id_code)
				HttpPost_Zapier("有新的好朋友：["  + user_talk + "](" + userImageUrl  + ")" + `\n` + userStatus, "LINE 新好友" ,target_id_code,user_talk)
				//target_id_code := event.Source.UserID + event.Source.GroupID + event.Source.RoomID	//target_id_code := ""
				log.Print("觸發與 " + user_talk + " 加入好友")

				imageURL = SystemImageURL
				//template := LineTemplate_firstinfo
				t_msg := "我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這只提供查詢動畫的功能。\n如有其他建議或想討論，請對我輸入「開發者」進行聯絡。"
				obj_message := linebot.NewTemplateMessage(t_msg, LineTemplate_firstinfo)

				// username := ""
				// if target_id_code == "U6f738a70b63c5900aa2c0cbbe0af91c4"{
				// 	username = "懶懶"
				// }
				// if target_id_code == "Uf150a9f2763f5c6e18ce4d706681af7f"{
				// 	username = "包包"
				// }
				//reply 的寫法
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好啊！" + username + "～\n想知道我的嗜好，可以說：簡介\n\nPS：手機上可以看到不一樣的內容喔！"),obj_message).Do(); err != nil {
						log.Print(1288)
						log.Print(err)
				}
		}
		//觸發解除好友
		if event.Type == linebot.EventTypeUnfollow {
				HttpPost_JANDI("與 ["  + user_talk + "](" + userImageUrl + ") 解除好友" + `\n` + userStatus, "gray" , "LINE 被解除好友",target_id_code)
				HttpPost_IFTTT("與 "  + user_talk + " 解除好友" + `\n<br>` + userImageUrl + `\n<br>` + userStatus , "LINE 被解除好友" ,target_id_code)
				HttpPost_Zapier("與 ["  + user_talk + "](" + userImageUrl + ") 解除好友" + `\n` + userStatus , "LINE 被解除好友" ,target_id_code,user_talk)
				log.Print("觸發與 " + user_talk + " 解除好友")
		}
		//觸發加入群組聊天
		if event.Type == linebot.EventTypeJoin {
				HttpPost_JANDI("加入了 "  + user_talk , "blue" , "LINE 已加入群組",target_id_code)
				HttpPost_IFTTT("加入了 "  + user_talk , "LINE 已加入群組" ,target_id_code)
				HttpPost_Zapier("加入了 "  + user_talk , "LINE 已加入群組" ,target_id_code,user_talk)
				log.Print("觸發加入" + user_talk)
 				//source := event.Source
 				//log.Print("觸發加入群組聊天事件 = " + source.GroupID)
 				push_string := "很高興你邀請我進來這裡聊天！"

				//if source.GroupID == "Ca78bf89fa33b777e54b4c13695818f81"{
				if target_id_code == "Ca78bf89fa33b777e54b4c13695818f81"{
					push_string += "\n你好，" + user_talk + "。"
				}
					//push 的寫法
					// 				if _, err = bot.PushMessage(source.GroupID, linebot.NewTextMessage(push_string)).Do(); err != nil {
					// 					log.Print(err)
					// 				}
					// 				if _, err = bot.PushMessage("Ca78bf89fa33b777e54b4c13695818f81", linebot.NewTextMessage("這裡純測試對嗎？\n只發於測試聊天室「test」")).Do(); err != nil {
					// 					log.Print(err)
					// 				}
					//target_id_code := event.Source.UserID + event.Source.GroupID + event.Source.RoomID	//target_id_code := ""
				imageURL = SystemImageURL
				//template := LineTemplate_firstinfo
				t_msg := "我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這只提供查詢動畫的功能。\n如有其他建議或想討論，請對我輸入「開發者」進行聯絡。"
				obj_message := linebot.NewTemplateMessage(t_msg, LineTemplate_firstinfo)

				//reply 的寫法
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("群組聊天的各位大家好哇～！\n" + push_string + "\n\n想知道我的嗜好，請說：簡介"),obj_message).Do(); err != nil {
						log.Print(1351)
						log.Print(err)
				}
		}
		//觸發離開群組聊天
		if event.Type == linebot.EventTypeLeave {
				HttpPost_JANDI("被請離開 "  + user_talk , "gray" , "LINE 離開群組",target_id_code)
				HttpPost_IFTTT("被請離開 "  + user_talk , "LINE 離開群組",target_id_code)
				HttpPost_Zapier("被請離開 "  + user_talk , "LINE 離開群組",target_id_code,user_talk)
				log.Print("觸發被踢出 " + user_talk +  " 群組")
		}
		//？？？？？
			//https://admin-official.line.me/beacon/register
			//https://devdocs.line.me/en/#line-beacon
			//https://devdocs.line.me/ja/#line-beacon
		if event.Type == linebot.EventTypeBeacon {
			HttpPost_JANDI("[" + user_talk + "](" + userImageUrl + ") 觸發 Beacon（啥鬼）" + `\n` + userStatus, "yellow" , "LINE 對話同步",target_id_code)
			HttpPost_IFTTT(user_talk + " 觸發 Beacon（啥鬼）" + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 對話同步",target_id_code)
			HttpPost_Zapier("[" + user_talk + "](" + userImageUrl + ") 觸發 Beacon（啥鬼）" + `\n` + userStatus, "LINE 對話同步",target_id_code,user_talk)
			log.Print(user_talk + " 觸發 Beacon（啥鬼）")
		}
		//觸發收到訊息
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				//target_id_code := event.Source.UserID + event.Source.GroupID + event.Source.RoomID	//target_id_code := ""

				//測試群組跳過
				// if target_id_code == "Ca78bf89fa33b777e54b4c13695818f81" {

				// }else{
				// 	HttpPost_JANDI(target_item + " " + user_talk + "：" + message.Text, "yellow" , "LINE 對話同步")
				// 	HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text, "LINE 對話同步",target_id_code)
				// }

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

		 				// log.Print("event.Source.UserID = " + event.Source.UserID)
						// log.Print("event.Source.GroupID = " + event.Source.GroupID)
						// log.Print("event.Source.RoomID = " + event.Source.RoomID)
						
						// 				source := event.Source
						// 				log.Print("source.UserID = " + source.UserID)
										
						// 				userID := event.Source.UserID
						// 				log.Print("userID := event.Source.UserID = " + userID)

						// target_id_code := event.Source.UserID + event.Source.GroupID + event.Source.RoomID	//target_id_code := ""
						// if event.Source.UserID == ""{
						// 	target_id_code = event.Source.GroupID
						// } else {
						// 	target_id_code = event.Source.UserID
						// }
						//都提到最外面去了 for _, event := range events { 的下面

				//anime()
				bot_msg = anime(message.Text,target_id_code,"")//bot_msg = anime(message.Text,message.ID,"")
					log.Print("根據 anime function 匹配到的回應內容：" + bot_msg)
				
								//增加到這
					//				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
					// 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
					// 					log.Print(err)
					// 				}
								//https://devdocs.line.me/en/?go#send-message-object
				

				//沒辦法建立 anime function 直接在裡面操作（因為用途不一樣當然不能）。 只好先用加法，從下游進行正則分析處理 reg  //https://play.golang.org/p/cjO5La2cKR
					//anime() 只是負責處理字串，理所當然裡面無法做任何的發言動作。（除非把可以發言的相關物件傳進去？）
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

				//這是從字串結果來判斷的方式，但發現有其他方式判斷（直接 bot_msg==開發者）所以這個暫時不用				
				reg_loking_for_admin := regexp.MustCompile("^(你找我主人？OK！).*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*\\n.*") 
				//reg_loking_for_admin := regexp.MustCompile("^(你找我主人？OK！).*") 

				//2016.12.26:這裡的 bot_msg 已經是下游，經過 anime() 處理過了，沒有匹配的發言內容都會被濾掉。
				
				if bot_msg != ""{
					//2016.12.20+ for test	
					switch bot_msg{
						case "c38b3100b02ef42411a99b7975e4ff47":
							// if username == "懶懶" {
							// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("登入成功！")).Do(); err != nil {
							// 		log.Print(2162)
							// 		log.Print(err)
							// 	}
							// }
							switch username{
							case "懶懶":
								if target_id_code == "U6f738a70b63c5900aa2c0cbbe0af91c4"{
									template := linebot.NewConfirmTemplate(
										"驗證成功！\n要現在進入管理介面嗎？",
										linebot.NewPostbackTemplateAction("是","開啟管理者選單", ""),
										linebot.NewPostbackTemplateAction("否","登出管理者", ""),
									)
									obj_message := linebot.NewTemplateMessage("這功能只支援最新版本 APP 使用。\n請用 APP 端查看下一步。", template)
									if _, err = bot.ReplyMessage(event.ReplyToken,obj_message).Do(); err != nil {
										log.Print(2162)
										log.Print(err)
									}
								}
							}
							return
						case "測試":
							if target_id_code == "U6f738a70b63c5900aa2c0cbbe0af91c4"{
								LineTemplate_test := linebot.NewCarouselTemplate(
									linebot.NewCarouselColumn(
										imageURL, "test", "For test mode.",
										linebot.NewPostbackTemplateAction("測試","test", ""),
										linebot.NewPostbackTemplateAction("管理模式","admin", ""),
										linebot.NewPostbackTemplateAction("申請使用管理者","開發者", "開發者"),
									),
									LineTemplate_other_example,
									LineTemplate_other,
									LineTemplate_feedback,
								)
								no_temp_msg := "你已觸發測試模式，請更新最新版本的 LINE 查看內容 。"
								obj_message := linebot.NewTemplateMessage(no_temp_msg, LineTemplate_test)
								if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
										log.Print(2148)
										log.Print(err)
								}
							}
							return
						case "GOTEST":
							//簡單說模板有三種（Y/N[1~2動]、Bottons[最多4個動作]、carousel[3個動作 && 並排最多五個(每個動作數量要一致)]），動作也有三種（操作使用者發言、POST兼使用者發言(使用者發言可為空)、URI 可連網址或 tel: 等協定）
								//bot_msg = "HI～ 我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這只提供查詢動畫的功能。\n如有其他建議或想討論，請對我輸入「開發者」進行聯絡。"
								//bot_msg = "有喔！有喔！你在找這個對吧！？\n" + "https://ani.gamer.com.tw/animeVideo.php?sn=5863" + "\n\n等等！這是最後一話！？"

								//2016.12.22+ free POST
								//func HttpPost_JANDI(body, connectColor, title, --url--) error  
								//http://nipponcolors.com/#matsuba
								// HttpPost_JANDI("test for LINE BOT", "#42602D" , "test")
								//HttpPost_IFTTT("test for line bot", "純測試",target_id_code) //2016.12.22+ 成功！！！
								//HttpPost_LINE_notify("test")
								
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

			 					//	 leftBtn := linebot.NewMessageTemplateAction("left", "left clicked")// 後面的參數 "left clicked" = 在使用者按下後，自動幫使用者發訊息
			 					//	 rightBtn := linebot.NewMessageTemplateAction("right", "right clicked")// 後面的參數 "right clicked" = 在使用者按下後，自動幫使用者發訊息
								 //	//.NewMessageTemplateAction("字面按鈕", "設定讓使用者按下後發送內容") 會讓使用者發送那樣的內容給系統
			 					//	 template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)
			 					//	 //.NewConfirmTemplate

								//linebot.NewTemplateMessage
		 						//2 buttons
		 						//.NewButtonsTemplate 模板，裡面最多只能有四個動作
		 						//.NewURITemplateAction 開啟指定網址的動作
		 						//.NewPostbackTemplateAction ？？動作
		 						//						第二參數可以讓她　ＰＯＳＴ指定內容（但還不會處理．．．）	第三參數類似於 .NewMessageTemplateAction 的效果
			 					//	 imageURL := "https://images.gamme.com.tw/news2/2016/51/39/paCYoqCXkqSarqSZ.jpg"
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


			 						//obj_message := linebot.NewTemplateMessage("HI～ 我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這只提供查詢動畫的功能。\n如有其他建議或想討論，請對我輸入「開發者」進行聯絡。", template)//messgage := linebot.NewTemplateMessage("請使用更新 APP 或使用手機 APP 才能看到這個功能。", template)
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

								//這是個謎
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
								//func NewImagemapMessage
								//https://github.com/line/line-bot-sdk-go/blob/master/linebot/message.go > Actions:  actions
								//看起來好像可以有動作

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
							return
						case "今日動漫通":
							log.Print("今日動漫通")
							//imageURL = SystemImageURL
							//template := LineTemplate_today_q
							obj_message := linebot.NewTemplateMessage(fb_q_msg, LineTemplate_today_q)
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
									log.Print(1630)
									log.Print(err)
							}
							return
						case "臉書":
							imageURL = SystemImageURL
							template := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "巴哈姆特動畫瘋相關網站", "動畫瘋官網、APP 載點、Facbook",
									linebot.NewURITemplateAction("巴哈姆特動畫瘋 官網","http://ani.gamer.com.tw"),
									LineTemplate_download_app,
									linebot.NewURITemplateAction("巴哈姆特動畫瘋 FB","https://www.facebook.com/animategamer"),
								),
								LineTemplate_feedback,
							)
							t_msg := "這是 巴哈姆特動畫瘋 的 Facebook：\nhttps://www.facebook.com/animategamer"
							obj_message := linebot.NewTemplateMessage(t_msg, template)
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
									log.Print(1630)
									log.Print(err)
							}
							return
						case "選單":
							imageURL = SystemImageURL
							//template := LineTemplate_firstinfo
							t_msg := "我最近很喜歡看巴哈姆特動畫瘋。\nhttp://ani.gamer.com.tw/\n\n你也可以問我動畫，我可以帶你去看！\n要問我動畫的話可以這樣問：\n動畫 動畫名稱 集數\n\n例如：\n動畫 美術社 12\nアニメ 美術社大有問題 12\nanime 美術社 １\n巴哈姆特 美術社 12\n以上這些都可以\n\n但中間要用空白或冒號、分號隔開喔！\n不然我會看不懂 ＞A＜\n\nPS：目前這只提供查詢動畫的功能。\n如有其他建議或想討論，請對我輸入「開發者」進行聯絡。"
							obj_message := linebot.NewTemplateMessage(t_msg, LineTemplate_firstinfo)
							if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
									log.Print(1639)
									log.Print(err)
							}
							return
						case "動畫瘋88":
							if target_item == "群組對話" {
								log.Print("觸發離開群組，APP 限定")
								//post KEY = 離開群組
								template := linebot.NewConfirmTemplate(
									"你確定要請我離開嗎QAQ？",
									//.NewPostbackTemplateAction(按鈕字面,post,替使用者發言)
									linebot.NewPostbackTemplateAction("是","按下確定離開群組對話", ""),
									linebot.NewPostbackTemplateAction("否", "取消離開群組",""),
								)
								obj_message := linebot.NewTemplateMessage("你確定要請我離開嗎QAQ？\n這功能只支援 APP 使用。\n請用 APP 端查看下一步。", template)
								if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
									log.Print(1654)
									log.Print(err)
								}
							}
							return
					}
					//2016.12.22+ 利用正則分析字串結果，來設置觸發找開發者的時候要 + 的 UI  //不能用 bot_msg == 開發者，因為 bot_msg 早就被改寫成一串廢話。
					if reg_loking_for_admin.ReplaceAllString(bot_msg,"$1") == "你找我主人？OK！"{
						log.Print("觸發找主人")
						template := linebot.NewCarouselTemplate(
							linebot.NewCarouselColumn(
								SystemImageURL, "開發者相關資訊", "你可以透過此功能\n聯絡 開發者",
								LineTemplate_addme,
								LineTemplate_chat,
								linebot.NewPostbackTemplateAction("聯絡 LINE 機器人開發者", "開發者", "開發者"),
							),
						)
						obj_message := linebot.NewTemplateMessage("上面這些都是聯絡開發者的相關方法。", template)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg),obj_message).Do(); err != nil {
							log.Print(1672)
							log.Print(err)
						}
						HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "yellow" , "LINE 同步：執行找開發者",target_id_code)
						HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：執行找開發者",target_id_code)
						HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：執行找開發者",target_id_code,user_talk)
						return
					}





					//因為 bot_msg==GOTEST 的時候，不可能會找到 anime_url。所以不用在 else 裡面。
					if anime_url!=""{
						//找到的時候的 UI
						imageURL = "https://i2.bahamut.com.tw/anime/FB_anime.png"
						template := linebot.NewCarouselTemplate(
							linebot.NewCarouselColumn(
								imageURL, "動畫搜尋結果", "在找" + message.Text + "對吧！？\n建議可以直接在巴哈姆特動畫瘋 APP 裡面播放！",							
								linebot.NewURITemplateAction("點此播放找到的動畫", anime_url),
								LineTemplate_download_app,
								linebot.NewMessageTemplateAction("查詢其他動畫", "目錄"),
							),
							LineTemplate_feedback,
							LineTemplate_other_example,
							LineTemplate_other,
						)
						obj_message := linebot.NewTemplateMessage(bot_msg, template)

						originalContentURL_1 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/722268f159dc640ed1639ffd31b4dd0d/94455.jpg"
	   					previewImageURL_1 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/722268f159dc640ed1639ffd31b4dd0d/94455.jpg"
	   					obj_message_img_1 := linebot.NewImageMessage(originalContentURL_1, previewImageURL_1)

						originalContentURL_2 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/f7e158cdc3f1e9640a5f5cf188c33b13/94454.jpg"
	   					previewImageURL_2 := "https://trello-attachments.s3.amazonaws.com/52ff05f27a3c676c046c37f9/5831e5e304f9fac88ac50a23/f7e158cdc3f1e9640a5f5cf188c33b13/94454.jpg"
	   					obj_message_img_2 := linebot.NewImageMessage(originalContentURL_2, previewImageURL_2)

						if _, err = bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage("可參考以下圖例操作讓搜尋到的影片，直接在巴哈姆特動畫瘋 APP 進行播放。"),obj_message_img_1,obj_message_img_2,obj_message).Do(); err != nil {
							log.Print(1724)
							log.Print(err)
						}
						HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "yellow" , "LINE 同步：查詢成功" + `\n` + anime_url,target_id_code)
						HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：查詢成功" + `\n` + anime_url,target_id_code)
						HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：查詢成功" + `\n` + anime_url,target_id_code,user_talk)
						log.Print("target_id_code +  anime_url = " + target_id_code + "\n" + anime_url)
					}else{
						//2016.12.22+ 利用正則分析字串結果，來設置觸發找不到的時候要 + 的 UI
						if reg_nofind.ReplaceAllString(bot_msg,"$1") == "才會增加比較慢XD）"{
							//找不到的時候
	 						imageURL = "https://i2.bahamut.com.tw/anime/FB_anime.png"
							template := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									imageURL, "找不到 "  +  message.Text   +   " 耶", "有可能打錯字或這真的沒有收錄，\n才會找不到。",							
									linebot.NewMessageTemplateAction("查看新番", "新番"),
									linebot.NewMessageTemplateAction("可查詢的其他動畫目錄", "目錄"),
									LineTemplate_download_app,
								),
								LineTemplate_feedback,
								LineTemplate_other_example,
								LineTemplate_other,
							)
							obj_message := linebot.NewTemplateMessage("除了「目錄」以外，\n你也可以輸入「新番」查詢近期的動畫。", template)
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg),obj_message).Do(); err != nil {
								log.Print(1763)
								log.Print(err)
							}
							HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "orange" , "LINE 同步：查詢失敗",target_id_code)
							HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：查詢失敗",target_id_code)
							HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：查詢失敗",target_id_code,user_talk)
						}else{
							//這是最原始的動作部分，還沒改寫 UI 模式的時候就靠這裡直接回傳結果就好。至於要傳什麼內容已經在 anime() 裡面處理好了。
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg)).Do(); err != nil {
								log.Print(1771)
								log.Print(err)
							}
						}
					}
				}
					// 				m := linebot.NewTextMessage("ok")
					// 					if _, err = bot.ReplyMessage(event.ReplyToken, m).Do(); err != nil {

					// 					}
									
									//----------PushMessage-----------這段可以跟 ReplyMessage 同時有效，但是只會在 1 對 1 有效。群組無效。---------
									//------開發者測試方案有效(好友最多50人/訊息無上限)，免費版(好友不限人數/訊息限制1000)、入門版無效，旗艦版、專業版有效。
									
									//http://muzigram.muzigen.net/2016/09/linebot-golang-linebot-heroku.html
									//https://github.com/mogeta/lbot/blob/master/main.go
					 		// source := event.Source
					 		// log.Print("source.UserID = " + source.UserID)
					 		// log.Print("target_id_code = " + target_id_code)
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
				// source := event.Source

				// userID := event.Source.UserID
				// groupID := event.Source.GroupID
				// RoomID := event.Source.RoomID
				// markID := userID + groupID + RoomID
				
				// log.Print(source.UserID)
				//----------------------------------------------------------------取得使用者資訊的寫法

				// username := ""
				// if markID == "U6f738a70b63c5900aa2c0cbbe0af91c4"{//if source.UserID == "U6f738a70b63c5900aa2c0cbbe0af91c4"{
				// 	username = "懶懶 = " + userID + groupID + RoomID //2016.12.20+
				// }
				// if markID == "Uf150a9f2763f5c6e18ce4d706681af7f"{
				// 	username = "包包"
				// }

				//https://devdocs.line.me/en/#get-content
				//[GAE/GoでLineBotをつくったよ〜 - ベーコンの裏](http://sun-bacon.hatenablog.com/entry/2016/10/10/233520)
				content, err := bot.GetMessageContent(message.ID).Do()
				if err != nil {
					log.Print(2141)
					log.Print(err)
				}
				defer content.Content.Close()
				log.Print("content.ContentType = " + content.ContentType)
				log.Print("content.ContentLength = ")
				log.Print(content.ContentLength) //檔案大小??
				log.Print("content.Content = ")
				log.Print(content.Content)

				//https://github.com/line/line-bot-sdk-go/blob/master/linebot/get_content_test.go
				//ContentLength
				//https://golang.org/pkg/image/jpeg/

				//目標是把 content.Content 存起來

				image, err := jpeg.Decode(content.Content)
				if err != nil {
					log.Print(2167)
					log.Print(err)
				}
				log.Printf("image %v", image.Bounds())
				//http://ithelp.ithome.com.tw/articles/10161612
				//https://webcache.googleusercontent.com/search?q=cache:cLTwZS5RNmMJ:https://libraries.io/go/github.com%252Fline%252Fline-bot-sdk-go%252Flinebot+&cd=6&hl=zh-TW&ct=clnk&gl=tw

				//暫時放棄 = =

									// file, err := ioutil.TempFile("temp.jpg", "")
									// if err != nil {
									// 	log.Print(2175)
									// 	log.Print(err)
									// }
									// defer file.Close()
									
									// _, err = ioutil.WriteFile("temp.jpg", []byte(image.Bounds()), 0600)//io.Copy(file, content.Content)
									// if err != nil {
									// 	log.Print(2182)
									// 	log.Print(err)
									// }
									// log.Printf("Saved %s", file.Name())


				//可以
				// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這圖片是？\n\n" + username + "你丟給我圖片幹嘛！\n我眼睛還沒長好看不懂XD")).Do(); err != nil {
				// 	log.Print(1845)
				// 	log.Print(err)
				// }
			case *linebot.VideoMessage:
				//https://github.com/dongri/line-bot-sdk-go
				originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/video-original.mp4"
				previewImageURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/video-preview.png"
				obj_message := linebot.NewVideoMessage(originalContentURL, previewImageURL)
 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這影片是？\n我也給你影片吧！\n\n這只是測試功能"),obj_message).Do(); err != nil {
 					log.Print(1854)
 					log.Print(err)
 				}
			case *linebot.AudioMessage:
				//下面都是 OK 的寫法，但是還是沒辦法取得...........
				//另外因為現在這個專案不適合這樣玩
				// originalContentURL := "https://dl.dropboxusercontent.com/u/358152/linebot/resource/ok.m4a"
				// duration := 1000
				// obj_message := linebot.NewAudioMessage(originalContentURL, duration)
 				//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這是什麼聲音？"),obj_message).Do(); err != nil {
 				//	log.Print(1862)
 				//	log.Print(err)
 				//}
			case *linebot.LocationMessage:
				log.Print("message.Title = " + message.Title)
				log.Print("message.Address = " + message.Address)
				log.Print("message.Latitude = ")
				log.Print(message.Latitude)
				log.Print("message.Longitude = ")
				log.Print(message.Longitude)
				obj_message := linebot.NewLocationMessage(message.Title, message.Address, message.Latitude, message.Longitude)

				//case 1
				//obj_message_1 := linebot.NewLocationMessage("歡迎光臨", "地球", 25.022413, 121.556427) //台北市信義區富陽街46號
					//obj_message_2 := linebot.NewLocationMessage("歡迎光臨", "哪個近", 25.022463, 121.556454) //這個遠

				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你在這裡？"),obj_message).Do(); err != nil {
					log.Print(1876)
					log.Print(err)
				}
			case *linebot.StickerMessage:
				log.Print("message.PackageID = " + message.PackageID)
				log.Print("message.StickerID = " + message.StickerID)
					//https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go handleSticker
					//message.PackageID, message.StickerID
				//丟跟對方一樣的貼圖回他
				obj_message_moto := linebot.NewStickerMessage(message.PackageID, message.StickerID)
					//https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
					//2016.12.20+ 多次框框的方式成功！（最多可以五個）
					//.NewStickerMessage 發貼貼圖成功	 //https://devdocs.line.me/files/sticker_list.pdf			
				obj_message := linebot.NewStickerMessage("2", "514") //https://devdocs.line.me/en/?go#send-message-object
 				//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("OU<"),linebot.NewTextMessage("0.0"),linebot.NewTextMessage("．ω．"),linebot.NewTextMessage("．ω．")).Do(); err != nil {

				PackageID_int := 0
				StickerID_int := 0
				if PackageID_int, err = strconv.Atoi(message.PackageID); err != nil {
					log.Print("7793 字串轉整數失敗")
					log.Print(PackageID_int)
					log.Print(err.Error())
				}

				if StickerID_int, err = strconv.Atoi(message.StickerID); err != nil {
					log.Print("7798 字串轉整數失敗")
					log.Print(StickerID_int)
					log.Print(err.Error())
				}

				//特別處理過貼圖範圍外的貼圖
				if (PackageID_int!=0) && (PackageID_int<=4){
					// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("．ω．"),obj_message_moto,obj_message).Do(); err != nil {
					// 	log.Print(7806)
					// 	log.Print(err)
					// }
					if _, err = bot.ReplyMessage(event.ReplyToken, obj_message_moto,obj_message).Do(); err != nil {
						log.Print(7806)
						log.Print(err)
					}
				}else{
					// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("．ω．"),obj_message).Do(); err != nil {
					// 	log.Print(7811)
					// 	log.Print(err)
					// }
					//if(username=="LL"){
						//https://store.line.me/stickershop/product/1021884/zh-Hant
						//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("PackageID_int = " + strconv.Itoa(PackageID_int) + "\nStickerID_int = " + strconv.Itoa(StickerID_int)), obj_message).Do(); err != nil {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("https://store.line.me/stickershop/product/" + strconv.Itoa(PackageID_int) + "/zh-Hant"), obj_message).Do(); err != nil {	
							log.Print(7806)
							log.Print(err)
						}
					//}else{

					//}
				}

					// if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
					// 	log.Print(7806)
					// 	log.Print(err)
					// }
			}
		}
	}
}
