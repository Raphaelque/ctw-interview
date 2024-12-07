package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	apiKey       = "sk-210aef40ac9144c78c7fe00975b57371"
	translateURL = "https://api.deepseek.com/"
)

func Task(c *gin.Context) {
	// 读取上传的 JSON 文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not provided"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	inputJSON, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 获取源语言和目标语言
	sourceLang := c.PostForm("source_lang")
	targetLang := c.PostForm("target_lang")

	// 遍历 JSON 文件中的每个字符串字段并翻译
	gjson.ParseBytes(inputJSON).ForEach(func(key, value gjson.Result) bool {
		if value.Type == gjson.String {
			translatedText, err := translateText(c, value.String(), sourceLang, targetLang)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return false
			}
			inputJSON, _ = sjson.SetBytes(inputJSON, key.String(), translatedText)
		}
		return true
	})

	// 返回翻译后的 JSON 文件
	c.Header("Content-Disposition", "attachment; filename=translated.json")
	c.Data(http.StatusOK, "application/json", inputJSON)
}

type TranslationRequest struct {
	Text       string `json:"text"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
}

type TranslationResponse struct {
	TranslatedText string `json:"translated_text"`
}

func translateText(c *gin.Context, text, sourceLang, targetLang string) (string, error) {
	requestBody, err := json.Marshal(TranslationRequest{
		Text:       text,
		SourceLang: sourceLang,
		TargetLang: targetLang,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", translateURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := DoApiRequest(c, text)
	if err != nil {
		return "", err
	}
	//defer resp.Body.Close()

	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	return "", err
	//}
	//
	//var translationResponse TranslationResponse
	//err = json.Unmarshal(body, &translationResponse)
	//if err != nil {
	//	return "", err
	//}

	return resp, nil
}

func DoApiRequest(c *gin.Context, text string) (string, error) {
	client := openai.NewClient(
		option.WithAPIKey(apiKey), // defaults to os.LookupEnv("OPENAI_API_KEY")
		option.WithBaseURL(translateURL),
	)
	chatCompletion, err := client.Chat.Completions.New(c, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("翻译全部：" + text),
		}),
		Model: openai.F("deepseek-chat"),
	})
	if err != nil {
		panic(err.Error())
	}
	return chatCompletion.Choices[0].Message.Content, err
}

//func doRequest(c *gin.Context, req *http.Request) (*http.Response, error) {
//	resp, err := GetHttpClient().Do(req)
//	if err != nil {
//		return nil, err
//	}
//	if resp == nil {
//		return nil, errors.New("resp is nil")
//	}
//	_ = req.Body.Close()
//	_ = c.Request.Body.Close()
//	return resp, nil
//}
