package adapter

import (
	"bytes"
	"context"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

// config をコンストラクタで設定するが，接続はしない点に留意する
type gptClient struct {
	client  *http.Client
	apiKey  string
	baseUrl string
}

func NewGptClient(client *http.Client, apiKey string, baseUrl string) *gptClient {
	return &gptClient{client: client, apiKey: apiKey, baseUrl: baseUrl}
}

func (gc *gptClient) CreateIllustration(ctx context.Context, image *multipart.FileHeader) error {
	// httpリクエストを生成
	// ヘッダーなどの設定
	// http.Client の Do でリクエストを実行
	// multipart/form-data の作成

	/// httpRequest 生成のためのbuffer
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	_ = writer.WriteField("model", "gpt-image-1")
	_ = writer.WriteField("prompt", "Create a illustration based on submitted image")

	req, err := http.NewRequest("POST", gc.baseUrl, &buf)
	if err != nil {
		log.Println("error:", err)
		return err
	}

	fw, err := writer.CreateFormFile("image", image.Filename)
	if err != nil {
		return err
	}
	f, err := image.Open()
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(fw, f); err != nil {
		return err
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+gc.apiKey)

	res, err := gc.client.Do(req)
	if err != nil {
		log.Println("error at response")
		return err
	}
	defer res.Body.Close()

	return nil
}
