package main

import (
	"fmt"
	onestcaptcha "onest_captcha_service/onest_captcha"
)

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

	// solver recaptcha v3 enterprise:
	config := onestcaptcha.RecaptchaV3EnterpriseTaskProxylessConfig{
		SiteURL:    "https://2captcha.com/demo/recaptcha-v3",
		SiteKey:    "6LfB5_IbAAAAAMCtsjEHEHKqcB9iQocwwxTiihJu",
		PageAction: "demo_action",
	}
	data, err := client.RecaptchaV3EnterpriseTaskProxyless(config)
	if err != nil { // error
		fmt.Println(err)
	}
	// success
	fmt.Println(data.Token)
	fmt.Println(data.UserAgent)

}
