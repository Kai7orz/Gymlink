package service

import (
	"context"
	"errors"
	"gymlink/internal/entity"
	"log"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type exerciseService struct {
	e  ExerciseQueryRepo
	ec ExerciseCreateRepo
	a  AuthClient
	gc GptClient
	ac AwsClient
}

func NewExerciseService(e ExerciseQueryRepo, ec ExerciseCreateRepo, a AuthClient, gc GptClient, ac AwsClient) (ExerciseService, error) {
	if e == nil {
		return nil, errors.New("nil error: ExerciseQueryRepo or ExerciseCreateRepo")
	}
	return &exerciseService{e: e, ec: ec, a: a, gc: gc, ac: ac}, nil
}

func (s *exerciseService) GetRecordsById(ctx context.Context, id int64) ([]entity.RecordType, error) {
	recordsRaw, err := s.e.GetRecordsById(ctx, id)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	records := []entity.RecordType{}

	for _, record := range recordsRaw {
		presignedGetRequest, err := s.ac.GetObject(ctx, record.ObjectKey, 60)
		if err != nil {
			log.Println("error presigned get request")
			return nil, err
		}
		// presigned url 取得から実装始める
		log.Println("URL==>", presignedGetRequest.URL)

		newRecord := entity.RecordType{
			Id:             record.Id,
			UserId:         record.UserId,
			UserName:       record.UserName,
			PresignedImage: presignedGetRequest.URL,
			CleanUpTime:    record.CleanUpTime,
			CleanUpDate:    record.CleanUpDate,
			Comment:        record.Comment,
			LikesCount:     record.LikesCount,
		}

		records = append(records, newRecord)
	}

	return records, nil
}

func (s *exerciseService) GetRecords(ctx context.Context) ([]entity.RecordType, error) {
	recordsRaw, err := s.e.GetRecords(ctx)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	records := []entity.RecordType{}

	for _, record := range recordsRaw {
		presignedGetRequest, err := s.ac.GetObject(ctx, record.ObjectKey, 60)
		if err != nil {
			log.Println("error presigned get request")
			return nil, err
		}
		// presigned url 取得から実装始める
		log.Println("URL==>", presignedGetRequest.URL)

		newRecord := entity.RecordType{
			Id:             record.Id,
			UserId:         record.UserId,
			UserName:       record.UserName,
			PresignedImage: presignedGetRequest.URL,
			CleanUpTime:    record.CleanUpTime,
			CleanUpDate:    record.CleanUpDate,
			Comment:        record.Comment,
			LikesCount:     record.LikesCount,
		}

		records = append(records, newRecord)
	}

	for _, rec := range records {
		log.Println("record list:", rec.Id)
		log.Println("record url:", rec.PresignedImage)
	}

	return records, nil
}

func (s *exerciseService) CreateRecord(ctx context.Context, objectKey string, cleanUpTimeRaw string, cleanUpDateRaw string, comment string, idToken string) error {
	if s.a == nil {
		return errors.New("error auth nil")
	}
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		log.Println("error: ", err)
		return errors.New("failed to verify user ")
	}
	cleanUpTimeInt, err := strconv.Atoi(cleanUpTimeRaw)
	if err != nil {
		log.Println("failed to convert to integer")
		return errors.New("failed to convert to integer")
	}

	cleanUpTime := int64(cleanUpTimeInt)

	cleanUpDate, err := time.Parse("20060102", cleanUpDateRaw)
	if err != nil {
		log.Println("failed to convert string to date")
		return errors.New("failed to convert to date")
	}
	err = s.ec.CreateRecordById(ctx, objectKey, cleanUpTime, cleanUpDate, comment, token.UID)
	if err != nil {
		log.Println("error: ", err)
		return err
	}

	return nil
}

func (s *exerciseService) CreateLike(ctx context.Context, exerciseRecordId int64, idToken string) error {
	if s.a == nil {
		return errors.New("error auth nil")
	}
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		log.Println("error: ", err)
		return errors.New("failed to verify user ")
	}
	err = s.ec.CreateLike(ctx, exerciseRecordId, token.UID)
	if err != nil {
		log.Println("error create like: ", err)
		return err
	}
	return nil
}

func (s *exerciseService) DeleteLikeById(ctx context.Context, exerciseRecordId int64, idToken string) error {
	if s.a == nil {
		return errors.New("error auth nil")
	}
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		log.Println("failed to verify user ")
		return err
	}
	err = s.ec.DeleteLike(ctx, exerciseRecordId, token.UID)
	if err != nil {
		log.Println("error delete like ")
		return err
	}

	return nil

}

func (s *exerciseService) GenerateImgAndUpload(ctx context.Context, image *multipart.FileHeader, s3KeyRaw string) (string, error) {
	// 画像を読み込み
	// GPT API interface　の　アップロード関数を呼び出す
	// レスポンスの画像を記録する．
	err := s.gc.CreateIllustration(ctx, image)
	if err != nil {
		log.Println("error create illustration")
		return "", err
	}

	//	s3Key の名前を UNIQUE 化する処理
	uuidV1, err := uuid.NewRandom()
	if err != nil {
		log.Println("failed to generate uuid")
		return "", err
	}
	// adapter で S3 へ画像をs3Key に基準でアップロードする処理が必要
	s3Key := uuidV1.String() + s3KeyRaw + ".png"
	//	os.File を S3　のアップロード時使うからメモリに保存して，それを s3 へアップロードする
	// 最初は test_image.png でメモリ保存してそれを上書きする運用 これだと同時に複数人が使えないので，のちにkey を使って，それぞれの保存名は変えていき　アップロード後に削除する運用（今はプログラムでtest_image.png を保存すると ownership の関係上pngをホスト側で確認できず不便なので，同名保存の運用にする

	err = s.ac.UploadImage(ctx, s3Key)
	if err != nil {
		log.Println("error upload image to S3 ", err)
	}
	// key をもとにS3 からイラストをとってこれるか確認したい
	presignedGetRequest, err := s.ac.GetObject(ctx, s3Key, 60)
	if err != nil {
		log.Println("error presigned get request")
		return "", err
	}
	// presigned url 取得から実装始める
	log.Println("URL==>", presignedGetRequest.URL)
	// アップロード後に key と exercise record のデータ、メタ情報を　DB へ INSERT する処理を記述
	//
	return s3Key, nil
}
