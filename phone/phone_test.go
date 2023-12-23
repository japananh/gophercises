package phone_test

import (
	"os"
	"testing"

	"github.com/japananh/gophercises/phone"
	"gorm.io/gorm"
)

func TestPhoneNumbers_TableName(t *testing.T) {
	type fields struct {
		Id    string
		Phone string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Get table name successfully",
			fields: fields{},
			want:   phone.TableName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ph := phone.Numbers{
				Id:    tt.fields.Id,
				Phone: tt.fields.Phone,
			}
			if got := ph.TableName(); got != tt.want {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_connectDB(t *testing.T) {
	t.Setenv("DSN", "host=localhost user=postgres password=postgres dbname=gophercises_phone port=5430 sslmode=disable TimeZone=Asia/Bangkok")
	tests := []struct {
		name    string
		wantDb  *gorm.DB
		wantErr bool
	}{
		{
			name:    "Test Postgres connection successfully",
			wantDb:  &gorm.DB{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDb, err := phone.ConnectDB(os.Getenv("DSN"))
			if (err != nil) != tt.wantErr {
				t.Errorf("connectDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDb == nil {
				t.Errorf("connectDB() gotDb = %v, want %v", gotDb, tt.wantDb)
			}
		})
	}
}

func Test_normalize(t *testing.T) {
	type args[T int | string] struct {
		in T
	}
	tests := []struct {
		name    string
		args    args[string]
		wantOut string
	}{
		{
			name:    "Normalize successfully 1",
			args:    args[string]{in: "1234567891"},
			wantOut: "1234567891",
		},
		{
			name:    "Normalize successfully 2",
			args:    args[string]{in: "(123) 456 7892"},
			wantOut: "1234567892",
		},
		{
			name:    "Normalize successfully 3",
			args:    args[string]{in: "(123)456-7892"},
			wantOut: "1234567892",
		},
		{
			name:    "Normalize successfully 4",
			args:    args[string]{in: "123-456-7890"},
			wantOut: "1234567890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := phone.Normalize(tt.args.in); gotOut != tt.wantOut {
				t.Errorf("normalize() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
