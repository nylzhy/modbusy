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
			//大小端测试
			name: "testReadDI_SH",
			args: args{
				addr:     1,
				quantity: 300,
			},
			want:    respdu{02, 00, 01, 01, 0x2C},
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

func Test_writeOneCoil(t *testing.T) {
	type args struct {
		addr  uint16
		coilv bool
	}
	tests := []struct {
		name    string
		args    args
		want    respdu
		wantErr bool
	}{
		{
			name: "writecoil_test1",
			args: args{
				addr:  2,
				coilv: false,
			},
			want:    respdu{5, 0, 2, 0, 0},
			wantErr: false,
		},
		{
			name: "writecoil_test1",
			args: args{
				addr:  2,
				coilv: true,
			},
			want:    respdu{5, 0, 2, 0xFF, 0},
			wantErr: false,
		},
		{
			name: "writecoil_test1",
			args: args{
				addr:  65534,
				coilv: true,
			},
			want:    respdu{5, 0xFF, 0xFE, 0xFF, 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := writeOneCoil(tt.args.addr, tt.args.coilv)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeOneCoil() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("writeOneCoil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bool2byte(t *testing.T) {
	type args struct {
		// length int
		values []bool
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test1",
			args: args{
				// length: 1,
				values: []bool{false, true, false, true},
			},
			want: []byte{10},
		},
		{
			name: "test2",
			args: args{
				// length: 1,
				values: []bool{false, true, false, true, false, true, false, true},
			},
			want: []byte{170},
		},
		{
			name: "test3",
			args: args{
				// length: 2,
				values: []bool{true, false, false, true, false, false, false, true, false, false, true, true},
			},
			want: []byte{0x89, 0x0C},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bool2byte(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bool2byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeMulCoils(t *testing.T) {
	type args struct {
		addr     uint16
		quantity uint16
		coilsV   []bool
	}
	tests := []struct {
		name    string
		args    args
		want    respdu
		wantErr bool
	}{
		{
			name: "writeMulCoils_Test1",
			args: args{
				addr:     4,
				quantity: 8,
				coilsV:   []bool{false, true, false, true, false, true, false, true},
			},
			want:    respdu{0x0F, 0, 0x04, 0x00, 0x08, 0x01, 0xAA},
			wantErr: false,
		},
		{
			name: "writeMulCoils_Test2",
			args: args{
				addr:     4,
				quantity: 12,
				coilsV:   []bool{true, false, false, true, false, false, false, true, false, false, true, true},
			},
			want:    respdu{0x0F, 0, 0x04, 0x00, 0x0C, 0x02, 0x89, 0x0C},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := writeMulCoils(tt.args.addr, tt.args.quantity, tt.args.coilsV)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeMulCoils() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("writeMulCoils() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeOneHold(t *testing.T) {
	type args struct {
		addr  uint16
		holdv uint16
	}
	tests := []struct {
		name    string
		args    args
		want    respdu
		wantErr bool
	}{
		{
			name: "writeonehold",
			args: args{
				addr:  2,
				holdv: 65535,
			},
			want:    respdu{6, 0, 2, 0xff, 0xff},
			wantErr: false,
		},
		{
			name: "writeonehold2",
			args: args{
				addr:  2,
				holdv: 0,
			},
			want:    respdu{6, 0, 2, 0, 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := writeOneHold(tt.args.addr, tt.args.holdv)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeOneHold() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("writeOneHold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeMulHold(t *testing.T) {
	var i16 int16 = -1
	type args struct {
		addr     uint16
		quantity uint16
		holdvs   []uint16
	}
	tests := []struct {
		name    string
		args    args
		want    respdu
		wantErr bool
	}{
		{
			name: "writemulReg",
			args: args{
				addr:     0,
				quantity: 2,
				holdvs:   []uint16{65535, 20000},
			},
			want:    respdu{0x10, 0, 0, 0, 2, 4, 0xff, 0xff, 0x4e, 0x20},
			wantErr: false,
		},
		{
			name: "writemulReg2",
			args: args{
				addr:     5,
				quantity: 8,
				holdvs:   []uint16{0, 128, 256, 1024, 2048, 4096, 8182, uint16(i16)},
			},
			want:    respdu{0x10, 0, 5, 0, 8, 0x10, 00, 00, 00, 0x80, 01, 00, 04, 00, 8, 00, 0x10, 0x00, 0x1F, 0xF6, 0xFF, 0xFF},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := writeMulHold(tt.args.addr, tt.args.quantity, tt.args.holdvs)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeMulHold() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("writeMulHold() = %v, want %v", got, tt.want)
			}
		})
	}
}
