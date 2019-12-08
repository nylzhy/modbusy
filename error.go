package modbusy

import "fmt"

// Error implements error interface
type Error struct {
	//FunctionCode 功能码
	FunctionCode byte
	//ExceptionCode 错误原因码
	ExceptionCode byte
}

// Error Converts Modbus Exception Code to error message
func (e *Error) Error() error {
	var name string
	switch e.ExceptionCode {
	case ECIllegalFunction:
		name = "ILLEGAL FUNCTION/非法功能"
	case ECIllegalDataAddress:
		name = "ILLEGAL DATA ADDRESS/非法数据地址"
	case ECIllegalDataValue:
		name = "ILLEGAL DATA VALUE/非法数据值"
	case ECDeviceFailure:
		name = "SERVER DEVICE FAILURE/从站设备故障"
	case ECAcknowledge:
		name = "ACKNOWLEDGE/确认"
	case ECDeviceBusy:
		name = "SERVER DEVICE BUSY/从属设备忙"
	case ECNegativeAcknowledge:
		name = "GEGATIVEACKNOWLEDGE/未确认"
	case ECMemoryParityError:
		name = "MEMORY PARITY ERROR/存储奇偶性差错"
	case ECNotDefined:
		name = "NOT DEFINED 未定义"
	case ECGatewayPathUnavailable:
		name = "GATEWAY PATH UNAVAILABLE/不可用网关路径"
	case ECGatewayTargetDeviceFailedToRespond:
		name = "GATEWAY TARGET DEVICE FAILED TO RESPOND/网关目标设备响应失败"
	default:
		name = "UNKNOWN/未知错误"
	}
	return fmt.Errorf("'%v' Exception: %s", e.ExceptionCode, name)
}
