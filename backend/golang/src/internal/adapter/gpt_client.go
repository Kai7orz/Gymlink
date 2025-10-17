package adapter

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gymlink/internal/dto"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
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
	// content-type としてアップロードされたimage は png 形式でないとはじく処理
	if image.Header.Get("Content-Type") != "image/png" {
		log.Println("type error : image is not png format")
		return fmt.Errorf("image must be png format")
	}

	// httpリクエストを生成
	// ヘッダーなどの設定
	// http.Client の Do でリクエストを実行
	// multipart/form-data の作成

	/// httpRequest 生成のためのbuffer
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	err := writer.WriteField("model", "gpt-image-1")
	if err != nil {
		log.Println("error write field")
		return err
	}
	err = writer.WriteField("prompt", "Create a illustration based on submitted image")
	if err != nil {
		log.Println("error write field")
		return err
	}

	f, err := image.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/png")
	part.Set("Content-Disposition", fmt.Sprintf(`form-data; name="image"; filename="%s"`, image.Filename))
	wp, err := writer.CreatePart(part)
	if err != nil {
		log.Println("error:", err)
	}

	_, err = io.Copy(wp, f)
	if err != nil {
		log.Println("error: copy image object")
		return err
	}
	log.Println("copy")

	if err := writer.Close(); err != nil {
		log.Println("error: writer cannot close!")
		return err
	}
	req, err := http.NewRequest("POST", gc.baseUrl, &buf)
	if err != nil {
		log.Println("error:", err)
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+gc.apiKey)
	res, err := gc.client.Do(req)
	if err != nil {
		log.Println("error at response")
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("failed to read resopnse body ", err)
		return err
	}

	var response dto.ImageResponseType
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("failed to unmarshal")
		return err
	}
	// response data の中の b_64を取得し，image として保存する処理を記述
	if len(response.Data) == 0 || response.Data[0].B64Json == "" {
		return fmt.Errorf("unmarshal json: response is nothing")
	}
	dec, err := base64.StdEncoding.DecodeString(response.Data[0].B64Json)
	if err != nil {
		log.Println("failed to decode base64_image")
		return err
	}

	outFile, err := os.CreateTemp("", "gpt-image-*.png")
	if err != nil {
		return fmt.Errorf("create tmp file : %w", err)
	}
	defer outFile.Close()

	err = os.WriteFile("test_image.png", dec, 0x2)
	if err != nil {
		log.Println("error: write image", err)
		return err
	}

	if _, err := outFile.Write(dec); err != nil {
		return err
	}

	log.Println("illustration is finished")

	return nil
}
