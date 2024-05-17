package utils

import "testing"

func TestMD5V(t *testing.T) {
	type args struct {
		str []byte
		b   []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				str: []byte("hello world"),
				b:   []byte("1234567890"),
			},
			want: "313233343536373839305eb63bbbe01eeed093cb22bb8f5acdc3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5V(tt.args.str, tt.args.b...); got != tt.want {
				t.Errorf("MD5V(%v, %v) = %v, want %v", tt.args.str, tt.args.b, got, tt.want)
			}
		})
	}
}
