package utils

import "testing"

func TestPathExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PathExists(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PathExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
