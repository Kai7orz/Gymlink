package service

import (
	"context"
	"errors"
	"fmt"
	"gymlink/internal/entity"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type recordService struct {
	q  UserQueryRepo
	e  RecordQueryRepo
	ec RecordCommandRepo
	a  AuthClient
	gc GptClient
	ac AwsClient
}

func NewRecordService(q UserQueryRepo, e RecordQueryRepo, ec RecordCommandRepo, a AuthClient, gc GptClient, ac AwsClient) (RecordService, error) {
	if e == nil {
		return nil, errors.New("nil error: RecordQueryRepo or Record CreateRepo")
	}
	return &recordService{q: q, e: e, ec: ec, a: a, gc: gc, ac: ac}, nil
}

func (s *recordService) GetRecordsById(ctx context.Context, id int64) ([]entity.RecordType, error) {
	recordsRaw, err := s.e.GetRecordsById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get records by Id : %w", err)
	}

	records := []entity.RecordType{}

	for _, record := range recordsRaw {
		presignedGetRequest, err := s.ac.GetObject(ctx, record.ObjectKey, 300)
		if err != nil {
			return nil, fmt.Errorf("failed to get request")
		}

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

func (s *recordService) GetRecords(ctx context.Context) ([]entity.RecordType, error) {
	recordsRaw, err := s.e.GetRecords(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get records : %w", err)
	}

	records := []entity.RecordType{}

	for _, record := range recordsRaw {
		presignedGetRequest, err := s.ac.GetObject(ctx, record.ObjectKey, 300)
		if err != nil {
			return nil, fmt.Errorf("failed to get presigned request")
		}

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

func (s *recordService) CreateRecord(ctx context.Context, objectKey string, cleanUpTimeRaw string, cleanUpDateRaw string, comment string, idToken string) error {
	if s.a == nil {
		return fmt.Errorf("auth is nil")
	}
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return fmt.Errorf("failed to verify user : %w", err)
	}
	cleanUpTimeInt, err := strconv.Atoi(cleanUpTimeRaw)
	if err != nil {
		return fmt.Errorf("failed to convert to integer : %w", err)
	}

	cleanUpTime := int64(cleanUpTimeInt)

	cleanUpDate, err := time.Parse("20060102", cleanUpDateRaw)
	if err != nil {
		return fmt.Errorf("failed to convert to date : %w", err)
	}
	err = s.ec.CreateRecordById(ctx, objectKey, cleanUpTime, cleanUpDate, comment, token.UID)
	if err != nil {
		return fmt.Errorf("failed to create record by id : %w", err)
	}

	return nil
}

func (s *recordService) DeleteRecordById(ctx context.Context, userIdRaw int64, recordId int64, idToken string) error {
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return fmt.Errorf("failed to verify user : %w", err)
	}

	user, err := s.q.FindByToken(ctx, token.UID)
	if err != nil {
		return fmt.Errorf("error: find by token : %w", err)
	}
	if user.Id != userIdRaw {
		return fmt.Errorf("invalid: delete command must be used by owner : %w", err)
	}

	err = s.ec.DeleteRecordById(ctx, user.Id, recordId, token.UID)
	if err != nil {
		return fmt.Errorf("failed to delete record : %w", err)
	}
	return nil
}

func (s *recordService) CheckLikeById(ctx context.Context, recordId int64, idToken string) (bool, error) {
	if s.a == nil {
		return false, fmt.Errorf("error auth nil")
	}
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return false, fmt.Errorf("failed to verify user : %w", err)
	}
	liked, err := s.ec.CheckLike(ctx, recordId, token.UID)
	if err != nil {
		return false, fmt.Errorf("failed to check like : %w", err)
	}
	return liked, nil
}

func (s *recordService) CreateLike(ctx context.Context, recordId int64, idToken string) error {
	if s.a == nil {
		return fmt.Errorf("error auth nil")
	}
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return fmt.Errorf("failed to verify user : %w", err)
	}
	err = s.ec.CreateLike(ctx, recordId, token.UID)
	if err != nil {
		return fmt.Errorf("failed to create like : %w", err)
	}
	return nil
}

func (s *recordService) DeleteLikeById(ctx context.Context, recordId int64, idToken string) error {
	if s.a == nil {
		return errors.New("error auth nil")
	}
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return fmt.Errorf("failed to verify user : %w", err)
	}

	err = s.ec.DeleteLike(ctx, recordId, token.UID)
	if err != nil {
		return fmt.Errorf("failed to delete like : %w", err)
	}

	return nil

}

func (s *recordService) UploadIllustration(ctx context.Context, image *multipart.FileHeader, s3KeyRaw string, idToken string) (string, error) {
	if s.a == nil {
		return "", errors.New("error auth nil")
	}
	_, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return "", fmt.Errorf("failed to verify user : %w", err)
	}
	// 画像を読み込み
	// GPT API interface　の　アップロード関数を呼び出す
	// レスポンスの画像を記録する．
	err = s.gc.CreateIllustration(ctx, image)
	if err != nil {
		return "", fmt.Errorf("failed to create illustration: %w", err)
	}

	//	s3Key の名前を UNIQUE 化する処理
	uuidV1, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("failed to generate uuid : %w", err)
	}
	// adapter で S3 へ画像をs3Key に基準でアップロードする処理が必要
	s3Key := uuidV1.String() + s3KeyRaw + ".png"
	//	os.File を S3　のアップロード時使うからメモリに保存して，それを s3 へアップロードする
	// 最初は test_image.png でメモリ保存してそれを上書きする運用 これだと同時に複数人が使えないので，のちにkey を使って，それぞれの保存名は変えていき　アップロード後に削除する運用（今はプログラムでtest_image.png を保存すると ownership の関係上pngをホスト側で確認できず不便なので，同名保存の運用にする
	err = s.ac.UploadImage(ctx, s3Key)
	if err != nil {
		return "", fmt.Errorf("failed to upload image to S3 : %w", err)
	}
	return s3Key, nil
}
