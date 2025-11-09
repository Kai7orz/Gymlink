package service

import (
	"context"
	"gymlink/internal/domain"
	"testing"
	"time"
)

func TestNewRecordService_NilDeps(t *testing.T) {
	_, err := NewRecordService(nil, &fakeRecordQuery{}, &fakeRecordCmd{}, &fakeAuth{}, &fakeGptCli{}, &fakeAwsCli{})
}

func (f *fakeRecordQuery) GetRecordsById(ctx context.Context, id int64) ([]domain.RecordRawType, error) {
	if f.findErr != nil {
		return nil, f.checkErr
	}
	return f.records, nil
}

func (f *fakeRecordQuery) GetRecords(ctx context.Context) ([]domain.RecordRawType, error) {
	if f.findErr != nil {
		return nil, f.checkErr
	}
	return f.records, nil
}

func (f *fakeRecordCmd) CreateRecordById(ctx context.Context, objectKey string, cleanTime int64, date time.Time, comment string, uid string) error {
	if f.createErr != nil {
		return f.createErr
	}

}
