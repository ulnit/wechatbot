package gpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ulnit/wechatbot/config"
)

// BASEURL 请求地址
var BASEURL = "https://api.openai.com/v1/"

// ChatGPTResponseBody 请求体
type ChatGPTResponseBody struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChoiceItem           `json:"choices"`
	Usage   map[string]interface{} `json:"usage"`
}

type ChoiceItem struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	Logprobs     int    `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatGPTRequestBody 响应体
type ChatGPTRequestBody struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float32 `json:"temperature"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
}

type ChatGPTNewRequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Completions gpt文本模型回复
/*
curl https://api.openai.com/v1/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -d '{
    "model": "text-davinci-003",
    "prompt": "Say this is a test",
    "max_tokens": 7,
    "temperature": 0
  }'

*/
// gpt-3.5 模型回复
// curl https://api.openai.com/v1/chat/completions \
// -H "Content-Type: application/json" \
// -H "Authorization: Bearer $OPENAI_API_KEY" \
//
//	-d '{
//	 "model": "gpt-3.5-turbo",
//	 "messages": [{"role": "system", "content": "You are a helpful assistant."}, {"role": "user", "content": "Hello!"}]
//	}'
//
// 增加了gpt-3.5-turbo、gpt-4和gpt-4-32k模型
func Completions(msg string) (string, error) {
	// 定义请求体对象
	var requestBody interface{}

	//如果model为gpt-3.5-turbo，gpt-3.5-turbo-0613,gpt-3.5-turbo-16k,gpt-3.5-turbo-0301,gpt-3.5-turbo-16k-0613 则请求体对象为ChatGPTRequest;
	if condition := config.LoadConfig().Model == "gpt-3.5-turbo" || config.LoadConfig().Model == "gpt-3.5-turbo-0613" || config.LoadConfig().Model == "gpt-3.5-turbo-16k" || config.LoadConfig().Model == "gpt-3.5-turbo-0301" || config.LoadConfig().Model == "gpt-3.5-turbo-16k-0613"; condition {
		requestBody = ChatGPTNewRequestBody{
			Model:    config.LoadConfig().Model,
			Messages: []Message{{Role: "user", Content: msg}},
		}
		BASEURL = "https://api.openai.com/v1/chat/"
	} else {
		//如果model为text-davinci-003,davinci,text-davinci-001,ada,text-curie-001,text-ada-001,curie-instruct-beta,davinci-instruct-beta,text-babbage-001,text-davinci-002,curie，则请求体对象为ChatGPTRequestBody
		if condition := config.LoadConfig().Model == "text-davinci-003" || config.LoadConfig().Model == "davinci" || config.LoadConfig().Model == "text-davinci-001" || config.LoadConfig().Model == "ada" || config.LoadConfig().Model == "text-curie-001" || config.LoadConfig().Model == "text-ada-001" || config.LoadConfig().Model == "curie-instruct-beta" || config.LoadConfig().Model == "davinci-instruct-beta" || config.LoadConfig().Model == "text-babbage-001" || config.LoadConfig().Model == "text-davinci-002" || config.LoadConfig().Model == "curie"; condition {
			requestBody = ChatGPTRequestBody{
				Model:            config.LoadConfig().Model,
				Prompt:           msg,
				MaxTokens:        1024,
				Temperature:      0.7,
				TopP:             1,
				FrequencyPenalty: 0,
				PresencePenalty:  0,
			}
			BASEURL = "https://api.openai.com/v1/"
		}
	}

	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}
	log.Printf("request gpt json string : %v", string(requestData))
	req, err := http.NewRequest("POST", BASEURL+"completions", bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	apiKey := config.LoadConfig().ApiKey
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("gpt api status code not equals 200,code is %d", response.StatusCode))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	gptResponseBody := &ChatGPTResponseBody{}
	log.Println(string(body))
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return "", err
	}

	var reply string
	if len(gptResponseBody.Choices) > 0 {
		reply = gptResponseBody.Choices[0].Text
	}
	log.Printf("gpt response text: %s \n", reply)
	return reply, nil
}
