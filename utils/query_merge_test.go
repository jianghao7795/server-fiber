package utils

import (
	"testing"

	"gorm.io/gorm"
)

func TestMergeQuery(t *testing.T) {
	type args struct {
		db    *gorm.DB
		query interface{}
		args  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestMergeQuery",
			args: args{
				db:    nil,
				query: nil,
				args:  nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MergeQuery(tt.args.db, tt.args.query, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("MergeQuery(%v, %v, %v) error = %v, wantErr %v", tt.args.db, tt.args.query, tt.args.args, err, tt.wantErr)
			}
		})
	}
}
