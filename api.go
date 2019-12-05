package modbusy

// RawClient a modbus raw client read/write operation collection for modbus instruction with modbus protocol style
type RawClient interface {
	// 寄存器类型			功能				功能码		寄存器地址范围		读取/写入寄存器数量			操作数据
	// 线圈				读取线圈				0x01		0x0000-0xFFFF		0-2000(0x001~0x7D0)			bit
	// 离散输入			读取离散输入	 	 	0x02		0x0000-0xFFFF		1-2000(0x001~0x7D0)	 		bit
	// 保持寄存器		读取保持寄存器			0x03		0x0000-0xFFFF		1-125(0x01~0x7D)		 	2bytes
	// 输入寄存器		读取输入寄存器			0x04		0x0000-0xFFFF		1-125(0x01~0x7D)			2bytes
	// 线圈				写单个线圈				0x05		0x0000-0xFFFF		1							0xFF00/0x0000
	// 保持寄存器		写单个保持寄存器		0x06		0x0000-0xFFFF 		1							2bytes
	// 线圈				写多个线圈				0x0F 		0x0000-0XFFFF		1-2000(0x001~0x7D0)     	Bit
	// 保持寄存器		写多个保持寄存器	 	0X10		0x0000-0xFFFF		1-120(0x01-0x78)			2bytes

	//Read Coils/DiscreteInputs/InputRegisters/HoldingRegisters
	Read(slaveIDAddr uint8, funCode uint8, regAddr, length uint) (result []byte, err error)
	//Read/Write Coils/DiscreteInputs/InputRegisters/HoldingRegisters
	Exec(slaveIDAddr uint8, funCode uint8, regAddr, length uint, data []byte) (result []byte, err error)
}

// HoldRegClient a hold regigter client read/write operation collection for modbus instruction with PLC style
type HoldRegClient interface {
	//Read HoldingRegisters
	Read(slaveIDAddr uint8, regAddr, length uint) (result []byte, err error)
	//Read/Write HoldingRegisters
	Write(slaveIDAddr uint8, regAddr, length uint, data []byte) (result []byte, err error)
}

// PLCClient a client read/write operation collection for modbus instruction with PLC style
type PLCClient interface {
	// Exec execute one modbus instruction with plc mode
	// PLC Mode: false 表示读取，true 表写入
	//
	// PLC Modbus 数据地址、数据长度与功能码关系表
	// 数据读取
	// *******************************************************************
	// 数据起始地址, 	数据长度,	 		功能码
	// 0-9999,	  		1-2000,				01			// 读取离散输出
	// 10001-19999,	  	1-2000,          	02			// 读取离散输入
	// 40001-49999,	  	1-125,          	03			// 读取保持寄存器
	// 400001-465535,	1-125,				03			// 读取保持寄存器
	// 30001-39999,		1-125,				04			// 读取输入寄存器
	// 20001-29999,		1-125,				03/04		// 读取浮点寄存器
	// *******************************************************************
	//
	// 数据写入
	// *******************************************************************
	// 数据起始地址, 	数据长度,	 		功能码
	// 0-9999,	  		1,					05			// 写入单个离散输出
	// 40001-49999,	  	1,          		06			// 写入单个保持寄存器
	// 400001-465535,	1,					06			// 写入单个保持寄存器
	// 0-9999,	  		2-1968,				15			// 写入多个离散输出
	// 40001-49999,	  	2-123,          	16			// 写入多个保持寄存器
	// 400001-465535,	2-123,				16			// 写入单个保持寄存器
	// 20001-29999,		1-125,				06/16		// 写入多个浮点寄存器
	// *******************************************************************

	//Read/Write Coils/DiscreteInputs/InputRegisters/HoldingRegisters with PLC mode
	Exec(SlaveIDaddr uint8, mode bool, addr, length uint, databuf []byte) (result []byte, err error)
}
