package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	onestcaptcha "onest_captcha_service/onest_captcha"
)

func ImageURLToBase64(url string) (string, error) {
	// Send a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Encode the data as base64
	base64String := base64.StdEncoding.EncodeToString(data)

	return base64String, nil
}

func main() {
	//  init client
	var APIKEY = "ee00fb4aae5e45279cb1d5fe55be865d"
	client := onestcaptcha.OneStCaptchaClient(APIKEY)

	// // solver recaptcha v2:
	// config := onestcaptcha.RecaptchaV2TaskProxylessConfig{
	// 	SiteURL:   "https://www.google.com/recaptcha/api2/demo",
	// 	SiteKey:   "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-",
	// 	Invisible: true,
	// }
	// data, err := client.RecaptchaV2TaskProxyless(config)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(data.Token)

	// // solver recaptcha v2 enterprise:
	// config := onestcaptcha.RecaptchaV2EnterpriseTaskProxylessConfig{
	// 	SiteURL: "https://www.google.com/recaptcha/api2/demo",
	// 	SiteKey: "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-",
	// }
	// data, err := client.RecaptchaV2EnterpriseTaskProxyless(config)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(data.Token)

	// // solver recaptcha v3:
	// config := onestcaptcha.RecaptchaV3TaskProxylessConfig{
	// 	SiteURL:    "https://2captcha.com/demo/recaptcha-v3",
	// 	SiteKey:    "6LfB5_IbAAAAAMCtsjEHEHKqcB9iQocwwxTiihJu",
	// 	PageAction: "demo_action",
	// }
	// data, err := client.RecaptchaV3TaskProxyless(config)
	// if err != nil { // error
	// 	fmt.Println(err)
	// }
	// // success
	// fmt.Println(data.Token)

	// // solver recaptcha v3 enterprise:
	// config := onestcaptcha.RecaptchaV3EnterpriseTaskProxylessConfig{
	// 	SiteURL:    "https://2captcha.com/demo/recaptcha-v3",
	// 	SiteKey:    "6LfB5_IbAAAAAMCtsjEHEHKqcB9iQocwwxTiihJu",
	// 	PageAction: "demo_action",
	// }
	// data, err := client.RecaptchaV3EnterpriseTaskProxyless(config)
	// if err != nil { // error
	// 	fmt.Println(err)
	// }
	// // success
	// fmt.Println(data.Token)
	// fmt.Println(data.UserAgent)

	// // solver image2text:
	// url := "https://upload.wikimedia.org/wikipedia/commons/thumb/2/21/Flag_of_Vietnam.svg/225px-Flag_of_Vietnam.svg.png"
	// base64String, err := ImageURLToBase64(url)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// config := onestcaptcha.ImageToTextConfig{
	// 	Base64Image: base64String,
	// }
	// data, err := client.ImageToText(config)
	// if err != nil { // error
	// 	fmt.Println(err)
	// }
	// // success
	// fmt.Println(data.Token)

	// solve recaptchaClick
	// config := onestcaptcha.RecaptchaClickConfig{
	// 	UrlList: []string{"https://upload.wikimedia.org/wikipedia/commons/thumb/2/21/Flag_of_Vietnam.svg/225px-Flag_of_Vietnam.svg.png"},
	// 	Caption: "Select all squares with motorcycles",
	// }
	// data, err := client.RecaptchaClick(config)
	// if err != nil { // error
	// 	fmt.Println(err.Error())
	// }
	// // success
	// fmt.Println(data.Token)

	// // funcaptcha
	// config := onestcaptcha.FunCaptchaTaskProxylessConfig{
	// 	SiteURL: "https://signup.live.com",
	// 	SiteKey: "B7D8911C-5CC8-A9A3-35B0-554ACEE604DA",
	// }
	// data, err := client.FunCaptchaTaskProxyless(config)
	// if err != nil { // error
	// 	fmt.Println(err)
	// }
	// if data.Code != 0 {
	// 	// error
	// 	fmt.Println(data.Message)
	// } else {
	// 	// success
	// 	fmt.Println(data.Token)
	// }

	// hcaptcha
	config := onestcaptcha.HCaptchaTaskProxylessConfig{
		SiteURL: "https://discord.com/login",
		SiteKey: "f5561ba9-8f1e-40ca-9b5b-a0b3f719ef34",
	}
	data, err := client.HCaptchaTaskProxyless(config)
	if err != nil { // error
		fmt.Println(err)
	}
	if data.Code != 0 {
		// error
		fmt.Println(data.Message)
	} else {
		// success
		fmt.Println(data.Token)
	}

}
