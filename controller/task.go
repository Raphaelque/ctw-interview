package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"ctw-interview/model"
	"ctw-interview/response"
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

type taskResponse struct {
	TaskId int64 `json:"task_id"`
}

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
	// 完成文件处理操作，开始将数据传库
	task := model.Task{}
	_, err = task.Save()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"success": false,
		})
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

	// 将翻译后的 JSON 写入输出文件
	_ = os.WriteFile(strconv.FormatInt(task.Id, 10)+".json", inputJSON, 0644)

	// 返回翻译后的 JSON 文件
	c.JSON(http.StatusOK, gin.H{
		"message": "上传成功",
		"success": true,
		"data": taskResponse{
			TaskId: task.Id,
		},
	})
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
			openai.UserMessage("直接返回翻译结果:" + text),
		}),
		Model: openai.F("deepseek-chat"),
	})
	if err != nil {
		panic(err.Error())
	}
	return chatCompletion.Choices[0].Message.Content, err
}

func TaskDownload(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	task, err := model.GetTaskById(int64(taskId))
	if err != nil {
		response.FailWithMessage("task not found", c)
		return
	}

	filePath := "./" + strconv.FormatInt(task.Id, 10) + ".json"
	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file info"})
		return
	}

	// 设置响应头，指定文件名和内容类型
	c.Header("Content-Disposition", "attachment; filename=data.json")
	c.Header("Content-Type", "application/json")
	c.Header("Content-Length", string(fileInfo.Size()))

	// 将文件内容写入响应体
	http.ServeContent(c.Writer, c.Request, fileInfo.Name(), fileInfo.ModTime(), file)
}
