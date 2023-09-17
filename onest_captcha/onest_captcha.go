package onestcaptcha

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type OneStCaptcha struct {
	apikey   string
	BASE_URL string
}

type BalanceData struct {
	Code         int
	Message      string
	KeyType      string
	Balance      float64
	Subscription any
}

type GetResultResponse struct {
	Code    int
	Status  string
	Message string
	Data    struct {
		Token     string
		UserAgent string
	}
}
type RecaptchaV2TaskProxylessConfig struct {
	SiteURL   string
	SiteKey   string
	Invisible bool
	Timeout   int
	TimeSleep int
}

type RecaptchaV2EnterpriseTaskProxylessConfig struct {
	SiteURL   string
	SiteKey   string
	Invisible bool
	SPayload  string
	Timeout   int
	TimeSleep int
}

type RecaptchaV3TaskProxylessConfig struct {
	SiteURL    string
	SiteKey    string
	PageAction string
	MinScore   float64
	Timeout    int
	TimeSleep  int
}

type RecaptchaV3EnterpriseTaskProxylessConfig struct {
	SiteURL    string
	SiteKey    string
	Invisible  bool
	SPayload   string
	PageAction string
	MinScore   float64
	Timeout    int
	TimeSleep  int
}

// get task_id
type TaskIDResponse struct {
	Code    int
	Message string
	TaskId  int
}

// Return for all func
type RecaptchaReturn struct {
	Code    int
	Message string
	Token   string
}

type RecaptchaUserAgentReturn struct {
	Code      int
	Message   string
	Token     string
	UserAgent string
}

const TIMEOUT_DEFAULT = 180  // seconds
const TIME_SLEEP_DEFAULT = 2 // seconds

func OneStCaptchaClient(apikey string) *OneStCaptcha {
	return &OneStCaptcha{
		apikey:   apikey,
		BASE_URL: "https://api.1stcaptcha.com",
	}
}
func (c *OneStCaptcha) GetBalance() (float64, error) {
	url := fmt.Sprintf("%s/user/balance?apikey=%s", c.BASE_URL, c.apikey)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data BalanceData
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return 0, err
		}
		if data.Code == 0 {
			return data.Balance, nil
		} else {
			return 0, errors.New(data.Message)
		}
	} else {
		return 0, errors.New(resp.Status)
	}
}

func (c *OneStCaptcha) GetResult(taskID int, timeout int, timeSleep int, typeCaptcha string) (any, error) {
	tStart := time.Now()
	for time.Since(tStart).Seconds() < float64(timeout) {
		url := fmt.Sprintf("%s/getresult?apikey=%s&taskid=%d", c.BASE_URL, c.apikey, taskID)
		resp, err := http.Get(url)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			var data GetResultResponse
			err = json.NewDecoder(resp.Body).Decode(&data)
			if err != nil {
				return "", err
			}
			if data.Code == 0 {
				status := data.Status
				if status == "SUCCESS" {
					if typeCaptcha == "image2text" || typeCaptcha == "recaptcha_click" || typeCaptcha == "v3_enterprise" {
						return data.Data, nil
					}
					return data.Data.Token, nil
				} else if status == "ERROR" {
					return "", errors.New(data.Message)
				}
				time.Sleep(time.Duration(timeSleep) * time.Second)
			} else {
				return "", errors.New("Error " + data.Message)
			}
		} else {
			return "", errors.New("Error " + resp.Status)
		}
	}
	return "", errors.New("TIMEOUT")
}

func (c *OneStCaptcha) RecaptchaV2TaskProxyless(config RecaptchaV2TaskProxylessConfig) (RecaptchaReturn, error) {
	siteKey := config.SiteKey
	siteURL := config.SiteURL
	invisible := config.Invisible
	timeout := config.Timeout
	timeSleep := config.TimeSleep

	// set default value
	if timeout == 0 {
		timeout = TIMEOUT_DEFAULT
	}
	if timeSleep == 0 {
		timeSleep = TIME_SLEEP_DEFAULT
	}
	// invisible default false

	params := url.Values{
		"apikey":    {c.apikey},
		"sitekey":   {siteKey},
		"siteurl":   {siteURL},
		"version":   {"v2"},
		"invisible": {strconv.FormatBool(invisible)},
	}

	url := fmt.Sprintf("%s/recaptchav2", c.BASE_URL) + "?" + params.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return RecaptchaReturn{Code: 1, Message: err.Error()}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data TaskIDResponse
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return RecaptchaReturn{Code: 1, Message: err.Error()}, nil
		}
		if data.Code == 0 {
			taskID := data.TaskId
			token, err := c.GetResult(taskID, timeout, timeSleep, "")
			if err != nil {
				return RecaptchaReturn{Code: 1, Message: err.Error()}, nil
			}
			tokenStr, ok := token.(string)
			if !ok {
				return RecaptchaReturn{}, errors.New("Error Token must be String")
			}
			return RecaptchaReturn{Code: 0, Message: "Success", Token: tokenStr}, nil
		} else {
			return RecaptchaReturn{}, errors.New("Error " + data.Message)
		}
	} else {
		return RecaptchaReturn{}, errors.New("Error " + resp.Status)
	}
}

func (c *OneStCaptcha) RecaptchaV2EnterpriseTaskProxyless(config RecaptchaV2EnterpriseTaskProxylessConfig) (RecaptchaReturn, error) {
	siteKey := config.SiteKey
	siteURL := config.SiteURL
	s_payload := config.SPayload
	timeout := config.Timeout
	timeSleep := config.TimeSleep

	// set default value
	if timeout == 0 {
		timeout = TIMEOUT_DEFAULT
	}
	if timeSleep == 0 {
		timeSleep = TIME_SLEEP_DEFAULT
	}

	params := url.Values{
		"apikey":  {c.apikey},
		"sitekey": {siteKey},
		"siteurl": {siteURL},
	}
	if s_payload != "" {
		params.Add("s", s_payload)
	}

	url := fmt.Sprintf("%s/recaptchav2_enterprise", c.BASE_URL) + "?" + params.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return RecaptchaReturn{Code: 1, Message: err.Error()}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data TaskIDResponse
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return RecaptchaReturn{Code: 1, Message: err.Error()}, nil
		}
		if data.Code == 0 {
			taskID := data.TaskId
			token, err := c.GetResult(taskID, timeout, timeSleep, "")
			if err != nil {
				return RecaptchaReturn{Code: 1, Message: err.Error()}, nil
			}
			tokenStr, ok := token.(string)
			if !ok {
				return RecaptchaReturn{}, errors.New("Error Token must be String")
			}
			return RecaptchaReturn{Code: 0, Message: "Success", Token: tokenStr}, nil
		} else {
			return RecaptchaReturn{}, errors.New("Error " + data.Message)
		}
	} else {
		return RecaptchaReturn{}, errors.New("Error " + resp.Status)
	}
}

func (c *OneStCaptcha) RecaptchaV3TaskProxyless(config RecaptchaV3TaskProxylessConfig) (RecaptchaReturn, error) {
	siteKey := config.SiteKey
	siteURL := config.SiteURL
	pageAction := config.PageAction
	minScore := config.MinScore
	timeout := config.Timeout
	timeSleep := config.TimeSleep

	// set default value
	if minScore == 0 {
		minScore = 0.3
	}
	if timeout == 0 {
		timeout = TIMEOUT_DEFAULT
	}
	if timeSleep == 0 {
		timeSleep = TIME_SLEEP_DEFAULT
	}

	params := url.Values{
		"apikey":     {c.apikey},
		"sitekey":    {siteKey},
		"siteurl":    {siteURL},
		"pageaction": {pageAction},
		"minscore":   {strconv.FormatFloat(minScore, 'f', -1, 64)},
		"version":    {"v3"},
	}

	url := fmt.Sprintf("%s/recaptchav3", c.BASE_URL) + "?" + params.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return RecaptchaReturn{Code: 1, Message: err.Error()}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data TaskIDResponse
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return RecaptchaReturn{Code: 1, Message: err.Error()}, nil
		}
		if data.Code == 0 {
			taskID := data.TaskId
			token, err := c.GetResult(taskID, timeout, timeSleep, "")
			if err != nil {
				return RecaptchaReturn{Code: 1, Message: err.Error()}, nil
			}
			tokenStr, ok := token.(string)
			if !ok {
				return RecaptchaReturn{}, errors.New("Error Token must be String")
			}
			return RecaptchaReturn{Code: 0, Message: "Success", Token: tokenStr}, nil
		} else {
			return RecaptchaReturn{}, errors.New("Error " + data.Message)
		}
	} else {
		return RecaptchaReturn{}, errors.New("Error " + resp.Status)
	}
}

func (c *OneStCaptcha) RecaptchaV3EnterpriseTaskProxyless(config RecaptchaV3EnterpriseTaskProxylessConfig) (RecaptchaUserAgentReturn, error) {
	siteKey := config.SiteKey
	siteURL := config.SiteURL
	s_payload := config.SPayload
	pageAction := config.PageAction
	minScore := config.MinScore
	timeout := config.Timeout
	timeSleep := config.TimeSleep

	// set default value
	if minScore == 0 {
		minScore = 0.3
	}
	if timeout == 0 {
		timeout = TIMEOUT_DEFAULT
	}
	if timeSleep == 0 {
		timeSleep = TIME_SLEEP_DEFAULT
	}

	params := url.Values{
		"apikey":     {c.apikey},
		"sitekey":    {siteKey},
		"siteurl":    {siteURL},
		"pageaction": {pageAction},
		"minscore":   {strconv.FormatFloat(minScore, 'f', -1, 64)},
	}
	if s_payload != "" {
		params.Add("s", s_payload)
	}

	url := fmt.Sprintf("%s/recaptchav3_enterprise", c.BASE_URL) + "?" + params.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return RecaptchaUserAgentReturn{Code: 1, Message: err.Error()}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data TaskIDResponse
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return RecaptchaUserAgentReturn{Code: 1, Message: err.Error()}, nil
		}
		if data.Code == 0 {
			taskID := data.TaskId
			data_return, err := c.GetResult(taskID, timeout, timeSleep, "v3_enterprise")
			if err != nil {
				return RecaptchaUserAgentReturn{Code: 1, Message: err.Error()}, nil
			}

			data_map, ok := data_return.(struct {
				Token     string
				UserAgent string
			})
			if !ok {
				return RecaptchaUserAgentReturn{}, errors.New("Error struct for Token and UserAgent")
			}
			return RecaptchaUserAgentReturn{Code: 0, Message: "Success", Token: data_map.Token, UserAgent: data_map.UserAgent}, nil
		} else {
			return RecaptchaUserAgentReturn{}, errors.New("Error " + data.Message)
		}
	} else {
		return RecaptchaUserAgentReturn{}, errors.New("Error " + resp.Status)
	}
}
