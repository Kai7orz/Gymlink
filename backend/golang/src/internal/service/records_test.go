package service

import (
	"context"
	"gymlink/internal/domain"

	"github.com/stretchr/testify/mock"
)

type mockRecordQueryRepo struct {
	mock.Mock
}

type mockRecordCommandRepo struct {
	mock.Mock
}

func NewMockRecordQueryRepo() *mockRecordQueryRepo {
	return &mockRecordQueryRepo{}
}

func (mq *mockRecordQueryRepo) GetRecords(ctx context.Context) ([]domain.RecordType, error) {
	args := mq.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.RecordType), args.Error(1)
}
