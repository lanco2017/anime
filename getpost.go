package main

import (
	//"fmt"
	"log"
	"net/http"
	//"os"
	// "regexp"

	//"strconv"
	"strings"

	//"encoding/json" //https://golang.org/pkg/encoding/json/#example_Decoder
	
	// "github.com/line/line-bot-sdk-go/linebot"

	"bytes"

	// "io"
	"io/ioutil"

	// "image/jpeg"
	//"image/png"

	// https://github.com/go-martini/martini/blob/master/translations/README_zh_tw.md
	// "github.com/go-martini/martini"

	//http://ithelp.ithome.com.tw/articles/10159486
	//"github.com/alexcesaro/mail/gomail"

    // "os/exec"
    // "path/filepath"

    //"crypto/md5"
    //"encoding/hex"

    //http://l-lin.github.io/2015/01/31/Golang-Deploy_to_heroku
    // "database/sql"
    // _ "github.com/lib/pq"

)

func HttpPost_Zapier(body , title_text, this_id, codename string) error {
	body = strings.Replace(body,"\n", `\n`, -1)
	title_text = strings.Replace(title_text,"\n", `\n`, -1)
	this_id = strings.Replace(this_id,"\n", `\n`, -1)
	codename = strings.Replace(codename,"\n", `\n`, -1)
	//https://internal-api.ifttt.com/maker
	log.Print("已經進來 Zapier POST")
	log.Print("body = " + body)
	log.Print("title_text = " + title_text)
	log.Print("this_id = " + this_id)

	url := "https://hooks.zapier.com/hooks/catch/132196/txma4i/"
	jsonStr := `{
		"value1":"` + body + `",
		"value2": "` + title_text + `",
		"value3": "` + this_id + `",
		"value4": "` + codename + `"
	}`

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		log.Print(err)
		return err
	}

	// Content-Type 設定
	//req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)		
		return err
	}
	defer resp.Body.Close()

	log.Print(err)

	//http://cepave.com/http-restful-api-with-golang/
    // log.Print("response Status = ")
    // log.Print(resp.Status)
    // log.Print("response Headers = ")
    // log.Print(resp.Header)
    // rebody, _ := ioutil.ReadAll(resp.Body)
    // log.Print("response Body = " +string(rebody))
	//http://cepave.com/http-restful-api-with-golang/

	return err
}

func HttpPost_IFTTT(body , title_text, this_id string) error {
	body = strings.Replace(body,"\n", `\n`, -1)
	title_text = strings.Replace(title_text,"\n", `\n`, -1)
	this_id = strings.Replace(this_id,"\n", `\n`, -1)
	//https://internal-api.ifttt.com/maker
	log.Print("已經進來 IFTTT POST")
	log.Print("body = " + body)
	log.Print("title_text = " + title_text)
	log.Print("this_id = " + this_id)

	url := "https://maker.ifttt.com/trigger/linebot/with/key/WJCRNxQhGJuzPd-sUDext"
	jsonStr := `{
		"value1":"` + body + `",
		"value2": "` + title_text + `",
		"value3": "` + this_id + `"
	}`

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		log.Print(err)
		return err
	}

	// Content-Type 設定
	//req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)		
		return err
	}
	defer resp.Body.Close()

	log.Print(err)

	//http://cepave.com/http-restful-api-with-golang/
    log.Print("response Status = ")
    log.Print(resp.Status)
    log.Print("response Headers = ")
    log.Print(resp.Header)
    rebody, _ := ioutil.ReadAll(resp.Body)
    log.Print("response Body = " +string(rebody))
	//http://cepave.com/http-restful-api-with-golang/

	return err
}

func HttpPost_JANDI(body, connectColor, title, code string) error {
	body = strings.Replace(body,"\n", `\n`, -1)
	title = strings.Replace(title,"\n", `\n`, -1)
	code = strings.Replace(code,"\n", `\n`, -1)
	log.Print("已經進來 JANDI POST")
	log.Print("body = " + body)
	log.Print("connectColor = " + connectColor)
	log.Print("title = " + title)
	log.Print("code = " + code)

	url := "https://wh.jandi.com/connect-api/webhook/11691684/46e7f45fd4f68a021afbd844aed66430"
	jsonStr := `{
		"body":"` + body + `",
		"connectColor":"` + connectColor + `",
		"connectInfo" : [{
				"title" : "` + title + `",
				"description" : "這是來自 巴哈姆特 LINE BOT 的通風報信",
				"imageUrl": "https://line.me/R/ti/p/@pyv6283b、@sjk2434l"
		},{
				"title" : "參考數據",
				"description" : "` + code + `"
		}]
	}`

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		log.Print(err)
		return err
	}

	// Content-Type 設定
	req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)		
		return err
	}
	defer resp.Body.Close()

	log.Print(err)

	//http://cepave.com/http-restful-api-with-golang/
    log.Print("response Status = ")
    log.Print(resp.Status)
    log.Print("response Headers = ")
    log.Print(resp.Header)
    rebody, _ := ioutil.ReadAll(resp.Body)
    log.Print("response Body = " +string(rebody))
	//http://cepave.com/http-restful-api-with-golang/

	return err
}