package footsites

import (
	client "Golang-Sitescripts/client"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

var EmailToUse = ""
var UsernameToUse = ""
var PasswordToUse = ""
var captchas []string

func GetCrsfToken() (r string) {
	//fmt.Println("Getting CRSF Token")
	//Setup our GET request obj
	get := client.GET{
		Endpoint: "https://club.pokemon.com/us/pokemon-trainer-club/sign-up/",
	}
	//Retrieve a configured HTTP Request obj
	request := client.NewRequest(get)
	//Add our headers to the HTTP Request obj
	//	request.Header = AddHeaders(Header{cookie: []string{}, content: nil}, host)
	//Obtain the responsey
	respBytes, _ := client.NewResponse(request)
	result := string(respBytes)
	resString := GetStringInBetween(result, `value='`, ">")
	test := strings.Split(resString, "'")
	//	fmt.Println(`resstring`)
	//fmt.Println(test[0])

	return test[0]
}
func Test() string {
	get := client.GET{
		Endpoint: "https://club.pokemon.com/us/pokemon-trainer-club/parents/requestProvider.js.map",
	}
	request := client.NewRequest(get)
	respBytes, _ := client.NewResponse(request)
	result := string(respBytes)
	//	fmt.Println(result)
	return result
}
func SendDob(token string) (r string) {
	//fmt.Println("Sending DOB")
	payload := url.Values{
		"csrfmiddlewaretoken": {token},
		"picker__year":        {"1994"},
		"picker__month":       {"2"},
		"dob":                 {"1994-03-14"},
		"country":             {"US"},
	}

	post := client.POSTUrlEncoded{
		Endpoint:       "https://club.pokemon.com/us/pokemon-trainer-club/sign-up/",
		EncodedPayload: payload.Encode(),
	}

	request := client.NewRequest(post)
	request.Header = AddHeaders(Header{additionalHeaders: []additionalHeaders{
		{key: "referer", value: "https://club.pokemon.com/us/pokemon-trainer-club/sign-up/"},
	}, cookie: []string{}, content: bytes.NewReader([]byte(payload.Encode()))}, "localhost")
	respBytes, _ := client.NewResponse(request)
	result := string(respBytes)
	resString := GetStringInBetween(result, `value='`, ">")
	test := strings.Split(resString, "'")
	//	fmt.Println(`resstring`)
	//	fmt.Println(test[0])

	return test[0]
}
func SendReg(token string, cap string) bool {
	//	fmt.Println("Sending REG")
	rand.Seed(time.Now().UnixNano())

	min := 1111111
	max := 9999999
	ran := rand.Intn(max-min) + min
	EmailToUse = "antreeder" + fmt.Sprint(ran) + "@fortuna7.com"
	UsernameToUse = "ASDasfg" + fmt.Sprint(ran)
	PasswordToUse = "sFG1!" + fmt.Sprint(ran)

	payload := url.Values{
		"csrfmiddlewaretoken":   {token},
		"username":              {"ASDasfg" + fmt.Sprint(ran)},
		"password":              {"sFG1!" + fmt.Sprint(ran)},
		"confirm_password":      {"sFG1!" + fmt.Sprint(ran)},
		"email":                 {EmailToUse},
		"confirm_email":         {EmailToUse},
		"public_profile_opt_in": {"False"},
		"screen_name":           {"ASDasfg" + fmt.Sprint(ran)},
		"terms":                 {"on"},
		"g-recaptcha-response":  {cap},
	}

	post := client.POSTUrlEncoded{
		Endpoint:       "https://club.pokemon.com/us/pokemon-trainer-club/parents/sign-up",
		EncodedPayload: payload.Encode(),
	}

	//cache-control	max-age=0
	//sec-ch-ua	" Not A;Brand";v="99", "Chromium";v="98", "Google Chrome";v="98"
	//sec-ch-ua-mobile	?0
	//sec-ch-ua-platform	"Windows"
	//upgrade-insecure-requests	1
	//dnt	1
	//content-type	application/x-www-form-urlencoded
	//user-agent	Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36
	//accept	text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9

	//sec-fetch-site	same-origin

	request := client.NewRequest(post)
	request.Header = AddHeaders(Header{additionalHeaders: []additionalHeaders{
		{
			key: "referer", value: "https://club.pokemon.com/us/pokemon-trainer-club/parents/sign-up",
		}, {
			key: "referer", value: "https://club.pokemon.com/us/pokemon-trainer-club/parents/sign-up",
		},
		{
			key: "accept-language", value: "en-GB,en-US;q=0.9,en;q=0.8",
		}, {
			key: "sec-fetch-dest", value: "document",
		}, {
			key: "sec-fetch-user", value: "?1",
		}, {
			key: "sec-fetch-mode", value: "navigate",
		},
	}, cookie: []string{}, content: bytes.NewReader([]byte(payload.Encode()))}, "localhost")
	_, resp := client.NewResponse(request)
	//	fmt.Println(result)
	if strings.Contains(resp.Request.URL.String(), "rate_limit_exceeded") {
		fmt.Println("Rate limit")
		return false
	}
	return true
}
func startCap() string {
	//fmt.Println("Prepareing cap")
	get := client.GET{
		Endpoint: "http://2captcha.com/in.php?key=2bc08ce6750eb3ff4b9a6615529e8213&method=userrecaptcha&googlekey=6LdpuiYTAAAAAL6y9JNUZzJ7cF3F8MQGGKko1bCy&pageurl=https://club.pokemon.com/us/pokemon-trainer-club/parents/sign-up",
	}

	request := client.NewRequest(get)
	respBytes, _ := client.NewResponse(request)
	result := string(respBytes)
	//fmt.Println(result)

	return strings.Split(result, "|")[1]
}
func refreshCap(key string) string {
	//fmt.Println("Refreshing cap")
	for {
		//fmt.Println("Refresh")
		get := client.GET{
			Endpoint: "http://2captcha.com/res.php?key=2bc08ce6750eb3ff4b9a6615529e8213&action=get&id=" + key,
		}

		request := client.NewRequest(get)
		respBytes, _ := client.NewResponse(request)
		result := string(respBytes)
		//	fmt.Println(result)

		if strings.Split(result, "|")[0] == "OK" {
			captchas = append(captchas, strings.Split(result, "|")[1])
			return strings.Split(result, "|")[1]
		}
		time.Sleep(10 * time.Second)

	}
}
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	e += s + e - 1
	return str[s:e]
}
func GetTempMail() string {
	//"@mailkept.com"
	//"@promail1.net"
	//"@rcmails.com"
	//"@relxv.com"
	//"@folllo.com"
	//"@fortuna7.com"
	//"@invecra.com"
	//"@linodg.com"
	//"@awiners.com"
	//"@subcaro.com"
	fmt.Println("Checking mail")
	for {
		//fmt.Println("Mail refresh")
		hash := md5.Sum([]byte(EmailToUse))
		hashString := hex.EncodeToString(hash[:])

		get := client.GET{
			Endpoint: "https://privatix-temp-mail-v1.p.rapidapi.com/request/mail/id/" + hashString + "/",
		}
		request := client.NewRequest(get)
		request.Header.Add("x-rapidapi-host", "privatix-temp-mail-v1.p.rapidapi.com")
		request.Header.Add("x-rapidapi-key", "3e03f5891fmsh00b4925fa5d701fp13cb8ajsn06ae8b38cc52")

		respBytes, _ := client.NewResponse(request)
		result := string(respBytes)

		if !strings.Contains(result, "There are no emails yet") {
			fmt.Println("No emails waiting 5 seconds")
			urlTOActivate := GetStringInBetween(result, "https://club.pokemon.com/us/pokemon-trainer-club/activated/", `\"\n`)
			//fmt.Println(result)
			//	fmt.Println(strings.Split(urlTOActivate, `\`)[0])
			Verify(strings.Split(urlTOActivate, `\`)[0])
			return ""
		}
		time.Sleep(5 * time.Second)

	}
}
func Verify(key string) {
	get := client.GET{
		Endpoint: "https://club.pokemon.com/us/pokemon-trainer-club/activated/" + key,
	}
	request := client.NewRequest(get)
	respBytes, _ := client.NewResponse(request)
	result := string(respBytes)
	//fmt.Println(result)
	if strings.Contains(result, "Your account is now active") {
		fmt.Printf("Account activated - User: %s Password: %s", UsernameToUse, PasswordToUse)

	} else {
		fmt.Println("failed")
	}
}
func GenerateCaptchas() {
	key := startCap()
	refreshCap(key)
}
func Start() {
	client.SetupClient()
	for i := 0; i < 5; i++ {
		go GenerateCaptchas()
	}
	for i := 0; i < 5; i++ {
		res := GetCrsfToken()
		Test()
		res2 := SendDob(res)
		Test()
		//check if there is a key in our slice
		for {

			if len(captchas) > 0 {

				cap := captchas[0]
				captchas[0] = captchas[len(captchas)-1]
				captchas[len(captchas)-1] = ""
				captchas = captchas[:len(captchas)-1]

				//	cap := captchas.
				if SendReg(res2, cap) {
					time.Sleep(5 * time.Second)
					GetTempMail()
					fmt.Println(" Account activated")
				}
				client.SetupClient()
				break
			}

		}

	}

}
