# 1stCaptcha package for Golang

[1stcaptcha.com](https://1stcaptcha.com) package for Golang

Solver recaptchaV2, recaptchaV3, hcaptcha, funcaptcha, imageToText, Zalo Captcha,.... Super fast and cheapest.

# Install

```bash
go get github.com/1stcaptcha/1stcaptcha-golang
```

# Usage

## init client

```golang
import "onestcaptcha"

var APIKEY = "0aa92cd8393a49698c408ea0ee56c2a5"
client := onestcaptcha.OneStCaptchaClient(APIKEY)
```

## solver recaptcha v2:

```golang
config := onestcaptcha.RecaptchaV2TaskProxylessConfig{
    SiteURL: "https://www.google.com/recaptcha/api2/demo",
    SiteKey: "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJw",
    Invisible: true,
}
data, err := client.RecaptchaV2TaskProxyless(config)
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
```

## solver recaptcha v2 enterprise:

```golang
config := onestcaptcha.RecaptchaV2EnterpriseTaskProxylessConfig{
    SiteURL: "https://www.google.com/recaptcha/api2/demo",
    SiteKey: "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-",
}
data, err := client.RecaptchaV2EnterpriseTaskProxyless(config)
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
```

## solver recaptcha v3:

```golang
config := onestcaptcha.RecaptchaV3TaskProxylessConfig{
    SiteURL:    "https://2captcha.com/demo/recaptcha-v3",
    SiteKey:    "6LfB5_IbAAAAAMCtsjEHEHKqcB9iQocwwxTiihJu",
    PageAction: "demo_action",
}
data, err := client.RecaptchaV3TaskProxyless(config)
if err != nil { // error
    fmt.Println(err)
}
// success
fmt.Println(data.Token)
```

## solver recaptcha enterprise:

```golang
config := onestcaptcha.RecaptchaV3EnterpriseTaskProxylessConfig{
    SiteURL:    "https://2captcha.com/demo/recaptcha-v3",
    SiteKey:    "6LfB5_IbAAAAAMCtsjEHEHKqcB9iQocwwxTiihJu",
    PageAction: "demo_action",
}
data, err := client.RecaptchaV3EnterpriseTaskProxyless(config)
if err != nil { // error
    fmt.Println(err)
}
if data.Code != 0 {
    // error
    fmt.Println(data.Message)
} else {
    // success
    fmt.Println(data.Token)
    fmt.Println(data.UserAgent)
}

```

## solve image2text

```golang
// input string
config := onestcaptcha.ImageToTextConfig{
    Base64Image: "BASE_64_STRING",
}
//input bytes file
config := onestcaptcha.ImageToTextConfig{
    File: "File",
}
data, err := client.ImageToText(config)
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
```

## solve recaptchaClick

```golang
config := onestcaptcha.RecaptchaClickConfig{
    UrlList: []string{"LIST_1_TO_9_IMAGE"},
    Caption: "IMAGES_CAPTION",
}
data, err := client.RecaptchaClick(config)
if err != nil { // error
    fmt.Println(err.Error())
}
if data.Code != 0 {
    // error
    fmt.Println(data.Message)
} else {
    // success
    fmt.Println(data.Token)
}
```

## funcaptcha

```golang
config := onestcaptcha.FunCaptchaTaskProxylessConfig{
    SiteURL: "https://signup.live.com",
    SiteKey: "B7D8911C-5CC8-A9A3-35B0-554ACEE604DA",
}
data, err := client.FunCaptchaTaskProxyless(config)
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
```

## hcaptcha

```golang
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
```
