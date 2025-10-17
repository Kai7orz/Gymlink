package service

import (
	"context"
	"errors"
	"gymlink/internal/entity"
	"log"
	"mime/multipart"
	"time"
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

func (s *exerciseService) GetExercisesById(ctx context.Context, id int64) ([]entity.ExerciseRecordType, error) {
	exercises, err := s.e.GetExercisesById(ctx, id)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return exercises, nil
}

func (s *exerciseService) GetExercises(ctx context.Context) ([]entity.ExerciseRecordType, error) {
	exercises, err := s.e.GetExercises(ctx)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return exercises, nil
}

func (s *exerciseService) CreateRecord(ctx context.Context, image string, exerciseTime int64, date time.Time, comment string, idToken string) error {
	if s.a == nil {
		return errors.New("error auth nil")
	}
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		log.Println("error: ", err)
		return errors.New("failed to verify user ")
	}
	err = s.ec.CreateRecordById(ctx, image, exerciseTime, date, comment, token.UID)
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

func (s *exerciseService) GenerateImg(ctx context.Context, image *multipart.FileHeader, s3KeyRaw string) error {
	// 画像を読み込み
	// GPT API interface　の　アップロード関数を呼び出す
	// レスポンスの画像を記録する．
	// err := s.gc.CreateIllustration(ctx, image)
	// if err != nil {
	// 	log.Println("error create illustration")
	// 	return err
	// }

	//	s3Key の名前を UNIQUE 化する処理
	// uuidV1, err := uuid.NewRandom()
	// if err != nil {
	// 	log.Println("failed to generate uuid")
	// 	return err
	// }
	// adapter で S3 へ画像をs3Key に基準でアップロードする処理が必要
	// s3Key := uuidV1.String() + s3KeyRaw + ".png"
	//	os.File を S3　のアップロード時使うからメモリに保存して，それを s3 へアップロードする
	// 最初は test_image.png でメモリ保存してそれを上書きする運用 これだと同時に複数人が使えないので，のちにkey を使って，それぞれの保存名は変えていき　アップロード後に削除する運用（今はプログラムでtest_image.png を保存すると ownership の関係上pngをホスト側で確認できず不便なので，同名保存の運用にする
	s3Key := "katazuke.png"
	s.ac.UploadImage(ctx, s3Key)

	// key をもとにS3 からイラストをとってこれるか確認したい
	presignedGetRequest, err := s.ac.GetObject(ctx, s3Key, 60)
	if err != nil {
		log.Println("error presigned get request")
		return err
	}
	// presigned url 取得から実装始める
	log.Println("URL==>", presignedGetRequest.URL)
	// アップロード後に key と exercise record のデータ、メタ情報を　DB へ INSERT する処理を記述
	//
	return nil
}
