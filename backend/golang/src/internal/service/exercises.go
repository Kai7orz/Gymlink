package service

import (
	"context"
	"errors"
	"gymlink/internal/entity"
	"log"
	"time"
)

type exerciseService struct {
	e  ExerciseQueryRepo
	ec ExerciseCreateRepo
	a  AuthClient
}

func NewExerciseService(e ExerciseQueryRepo, ec ExerciseCreateRepo, a AuthClient) (ExerciseService, error) {
	if e == nil {
		return nil, errors.New("nil error: ExerciseQueryRepo or ExerciseCreateRepo")
	}
	return &exerciseService{e: e, ec: ec, a: a}, nil
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
		log.Println("auth :nil ✅")
		return errors.New("eror auth nil")
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
