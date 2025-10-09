package service

import (
	"context"
	"errors"
	"gymlink/internal/entity"
	"log"
)

type exerciseService struct {
	e ExerciseQueryRepo
}

func NewExerciseService(e ExerciseQueryRepo) (ExerciseService, error) {
	if e == nil {
		return nil, errors.New("nil error: ExerciseQueryRepo or ExerciseCreateRepo")
	}
	return &exerciseService{e: e}, nil
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
