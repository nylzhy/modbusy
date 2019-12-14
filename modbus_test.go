package modbusy

import (
	"reflect"
	"testing"
)

func Test_genStdByte(t *testing.T) {
	type args struct {
		values []uint16
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test1",
			args: args{
				values: []uint16{3, 2, 244},
			},
			want: []byte{0, 3, 0, 2, 0, 244},
		},
		{
			name: "test2",
			args: args{
				values: []uint16{0x0D47, 34, 278},
			},
			want: []byte{13, 71, 0, 34, 1, 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genStdByte(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genStdByte() = %v, want %v", got, tt.want)
			}
		})
	}
}
