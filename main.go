package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"crypto/sha256"

	"math/rand"
)

func TranslateToCn(txt string) (string, error) {

	YOUDAO_URL := "https://openapi.youdao.com/api"
	APP_KEY := " "
	APP_SECRET := "  "

	rand.Seed(time.Now().Unix()) // unix 时间戳，秒
	data := rand.Int63n(1000000000)
	salt := fmt.Sprintf("%x", data)
	curTime := fmt.Sprintf("%d", int32(time.Now().Unix()))
	signStr := fmt.Sprintf("%s%s%s%s%s", APP_KEY, truncate(txt), salt, curTime, APP_SECRET)
	sign := getDigest(signStr)

	resp, err := http.PostForm(YOUDAO_URL, url.Values{
		"from":     {"en"},
		"to":       {"zh-CHS"},
		"signType": {"v3"},
		"curtime":  {curTime},
		"appKey":   {APP_KEY},
		"q":        {txt},
		"salt":     {salt},
		"sign":     {sign},
	})
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	return string(body), nil
}

func truncate(q string) string {
	var len = len(q)
	if len <= 20 {
		return q
	}
	return fmt.Sprintf("%s%d%s", q[:10], len, q[len-10:])
}

func getDigest(q string) string {
	sum := sha256.Sum256([]byte(q))
	return fmt.Sprintf("%x", sum)
}
