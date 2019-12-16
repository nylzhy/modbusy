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

func Test_readCoils(t *testing.T) {
	type args struct {
		addr     uint16
		quantity uint16
	}
	tests := []struct {
		name    string
		args    args
		want    respdu
		wantErr bool
	}{
		{
			name: "testReadCoils_1",
			args: args{
				addr:     0,
				quantity: 3,
			},
			want:    respdu{01, 00, 00, 00, 03},
			wantErr: false,
		},
		{
			name: "testReadCoils_2",
			args: args{
				addr:     0,
				quantity: 20001,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "testReadCoils_3",
			args: args{
				addr:     65530,
				quantity: 6,
			},
			want:    respdu{01, 0xFF, 0xFA, 0x00, 0x06},
			wantErr: false,
		},
		{
			name: "testReadCoils_4",
			args: args{
				addr:     65530,
				quantity: 7,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readCoils(tt.args.addr, tt.args.quantity)
			if (err != nil) != tt.wantErr {
				// fmt.Println(tt.name)
				t.Errorf("readCoils() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCoils() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readHoldingReg(t *testing.T) {
	type args struct {
		addr     uint16
		quantity uint16
	}
	tests := []struct {
		name    string
		args    args
		want    respdu
		wantErr bool
	}{
		{
			name: "testReadHold_1",
			args: args{
				addr:     10,
				quantity: 3,
			},
			want:    respdu{03, 00, 0x0A, 00, 03},
			wantErr: false,
		},
		// {
		// 	name: "testReadHold_2",
		// 	args: args{
		// 		addr:     0,
		// 		quantity: 125,
		// 	},
		// 	want:    nil,
		// 	wantErr: true,
		// },
		{
			name: "testReadHold_2_1",
			args: args{
				addr:     0,
				quantity: 126,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "testReadHold_3",
			args: args{
				addr:     65530,
				quantity: 6,
			},
			want:    respdu{03, 0xFF, 0xFA, 0x00, 0x06},
			wantErr: false,
		},
		{
			name: "testReadHold_4",
			args: args{
				addr:     65530,
				quantity: 7,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readHoldingReg(tt.args.addr, tt.args.quantity)
			if (err != nil) != tt.wantErr {
				t.Errorf("readHoldingReg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readHoldingReg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readDiscreteInputs(t *testing.T) {
	type args struct {
		addr     uint16
		quantity uint16
	}
	tests := []struct {
		name    string
		args    args
		want    respdu
		wantErr bool
	}{
		{
			name: "testReadDI_1",
			args: args{
				addr:     1,
				quantity: 3,
			},
			want:    respdu{02, 00, 01, 00, 03},
			wantErr: false,
		},
		{
			name: "testReadDI_2",
			args: args{
				addr:     0,
				quantity: 20001,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "testReadDI_3",
			args: args{
				addr:     65530,
				quantity: 6,
			},
			want:    respdu{02, 0xFF, 0xFA, 0x00, 0x06},
			wantErr: false,
		},
		{
			name: "testReadDI_4",
			args: args{
				addr:     65530,
				quantity: 7,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readDiscreteInputs(tt.args.addr, tt.args.quantity)
			if (err != nil) != tt.wantErr {
				t.Errorf("readDiscreteInputs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readDiscreteInputs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readInputReg(t *testing.T) {
	type args struct {
		addr     uint16
		quantity uint16
	}
	tests := []struct {
		name    string
		args    args
		want    respdu
		wantErr bool
	}{
		{
			name: "testReadInput_1",
			args: args{
				addr:     0x0F,
				quantity: 0x0A,
			},
			want:    respdu{04, 00, 0x0F, 00, 0x0A},
			wantErr: false,
		},
		{
			name: "testReadInput_2",
			args: args{
				addr:     4,
				quantity: 126,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "testReadInput_3",
			args: args{
				addr:     65530,
				quantity: 6,
			},
			want:    respdu{04, 0xFF, 0xFA, 0x00, 0x06},
			wantErr: false,
		},
		{
			name: "testReadInput_4",
			args: args{
				addr:     65530,
				quantity: 7,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readInputReg(tt.args.addr, tt.args.quantity)
			if (err != nil) != tt.wantErr {
				t.Errorf("readInputReg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readInputReg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readReg(t *testing.T) {
	type args struct {
		regType  byte
		addr     uint16
		quantity uint16
	}
	tests := []struct {
		name    string
		args    args
		want    respdu
		wantErr bool
	}{
		{
			name: "testReadCoils_1",
			args: args{
				regType:  FCReadCoils,
				addr:     10,
				quantity: 3,
			},
			want:    respdu{01, 00, 0x0A, 00, 03},
			wantErr: false,
		},
		{
			name: "testReadCoils_2",
			args: args{
				regType:  FCReadCoils,
				addr:     0,
				quantity: 20001,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "testReadCoils_3",
			args: args{
				regType:  FCReadCoils,
				addr:     65530,
				quantity: 6,
			},
			want:    respdu{01, 0xFF, 0xFA, 0x00, 0x06},
			wantErr: false,
		},
		{
			name: "testReadCoils_4",
			args: args{
				regType:  FCReadCoils,
				addr:     65530,
				quantity: 7,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "testReadInput_1",
			args: args{
				regType:  FCReadInputReg,
				addr:     0x0F,
				quantity: 0x0A,
			},
			want:    respdu{04, 00, 0x0F, 00, 0x0A},
			wantErr: false,
		},
		{
			name: "testReadInput_2",
			args: args{
				regType:  FCReadInputReg,
				addr:     4,
				quantity: 126,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "testReadInput_3",
			args: args{
				regType:  FCReadInputReg,
				addr:     65530,
				quantity: 6,
			},
			want:    respdu{04, 0xFF, 0xFA, 0x00, 0x06},
			wantErr: false,
		},
		{
			name: "testReadInput_4",
			args: args{
				regType:  FCReadInputReg,
				addr:     65530,
				quantity: 7,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "testReadDI_1",
			args: args{
				regType:  FCReadDiscreteInputs,
				addr:     1,
				quantity: 3,
			},
			want:    respdu{02, 00, 01, 00, 03},
			wantErr: false,
		},
		{
			name: "testReadDI_2",
			args: args{
				regType:  FCReadDiscreteInputs,
				addr:     0,
				quantity: 20001,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "testReadDI_3",
			args: args{
				regType:  FCReadDiscreteInputs,
				addr:     65530,
				quantity: 6,
			},
			want:    respdu{02, 0xFF, 0xFA, 0x00, 0x06},
			wantErr: false,
		},
		{
			name: "testReadDI_4",
			args: args{
				regType:  FCReadDiscreteInputs,
				addr:     65530,
				quantity: 7,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "testReadHold_1",
			args: args{
				regType:  FCReadHoldingReg,
				addr:     10,
				quantity: 3,
			},
			want:    respdu{03, 00, 0x0A, 00, 03},
			wantErr: false,
		},
		// {
		// 	name: "testReadHold_2",
		// 	args: args{
		// 		addr:     0,
		// 		quantity: 125,
		// 	},
		// 	want:    nil,
		// 	wantErr: true,
		// },
		{
			name: "testReadHold_2_1",
			args: args{
				regType:  FCReadHoldingReg,
				addr:     0,
				quantity: 126,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "testReadHold_3",
			args: args{
				regType:  FCReadHoldingReg,
				addr:     65530,
				quantity: 6,
			},
			want:    respdu{03, 0xFF, 0xFA, 0x00, 0x06},
			wantErr: false,
		},
		{
			name: "testReadHold_4",
			args: args{
				regType:  FCReadHoldingReg,
				addr:     65530,
				quantity: 7,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readReg(tt.args.regType, tt.args.addr, tt.args.quantity)
			if (err != nil) != tt.wantErr {
				t.Errorf("readReg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readReg() = %v, want %v", got, tt.want)
			}
		})
	}
}
