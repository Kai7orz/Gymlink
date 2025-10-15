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
}

func NewExerciseService(e ExerciseQueryRepo, ec ExerciseCreateRepo, a AuthClient, gc GptClient) (ExerciseService, error) {
	if e == nil {
		return nil, errors.New("nil error: ExerciseQueryRepo or ExerciseCreateRepo")
	}
	return &exerciseService{e: e, ec: ec, a: a, gc: gc}, nil
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

func (s *exerciseService) CreateExercise(ctx context.Context, image string, exerciseTime int64, date time.Time, comment string, idToken string) error {
	if s.a == nil {
		return errors.New("error auth nil")
	}
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		log.Println("error: ", err)
		return errors.New("failed to verify user ")
	}
	err = s.ec.CreateExerciseById(ctx, image, exerciseTime, date, comment, token.UID)
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

func (s *exerciseService) GenerateImg(ctx context.Context, image *multipart.FileHeader) error {
	// 画像を読み込み
	// GPT API interface　の　アップロード関数を呼び出す
	// レスポンスの画像を記録する．
	err := s.gc.CreateIllustration(ctx, image)
	if err != nil {
		log.Println("error create illustration")
		return err
	}
	return nil
}
