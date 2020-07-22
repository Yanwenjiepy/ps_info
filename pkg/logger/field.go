package logger

import (
	"reflect"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogDetail expanded more fields to meet actual needs
type LogDetail struct {
	// current host server ip
	// Host string

	// current application id
	// Aid string

	// current instance id
	// Iid string

	// time consume, use Nanosecond
	Cost int64

	// request application id
	ReqID string

	// request method
	ReqMethod string

	// request path
	ReqPath string

	// client ip
	ClientAddr string

	// response status code
	RespCode int64

	// response error msg
	RespError string
}

// you can add host applicationID instanceID fields to logger recorder when init logger

// func (ld LogDetail)addHost() zap.Field {
//
// 	return zapcore.Field{
// 		Key:       "host",
// 		Type:      zapcore.StringType,
// 		Integer:   0,
// 		String:    ld.Host,
// 		Interface: nil,
// 	}
// }

// func (ld LogDetail)addApplicationID() zap.Field {
//
// 	return zapcore.Field{
// 		Key:       "aid",
// 		Type:      zapcore.StringType,
// 		Integer:   0,
// 		String:    ld.Aid,
// 		Interface: nil,
// 	}
// }

// func (ld LogDetail)addInstanceId() zap.Field {
//
// 	return zapcore.Field{
// 		Key:       "iid",
// 		Type:      zapcore.StringType,
// 		Integer:   0,
// 		String:    ld.Iid,
// 		Interface: nil,
// 	}
// }

func (ld LogDetail) addTimeCost() zap.Field {

	return zapcore.Field{
		Key:       "cost",
		Type:      zapcore.Int64Type,
		Integer:   ld.Cost,
		String:    "",
		Interface: nil,
	}
}

func (ld LogDetail) addRequestId() zap.Field {

	return zapcore.Field{
		Key:       "req_id",
		Type:      zapcore.StringType,
		Integer:   0,
		String:    ld.ReqID,
		Interface: nil,
	}
}

func (ld LogDetail) addRequestMethod() zap.Field {

	return zapcore.Field{
		Key:       "req_method",
		Type:      zapcore.StringType,
		Integer:   0,
		String:    ld.ReqMethod,
		Interface: nil,
	}
}

func (ld LogDetail) addRequestPath() zap.Field {

	return zapcore.Field{
		Key:       "req_path",
		Type:      zapcore.StringType,
		Integer:   0,
		String:    ld.ReqPath,
		Interface: nil,
	}
}

func (ld LogDetail) addClientAddr() zap.Field {

	return zapcore.Field{
		Key:       "client_addr",
		Type:      zapcore.StringType,
		Integer:   0,
		String:    ld.ClientAddr,
		Interface: nil,
	}
}

func (ld LogDetail) addResponseCode() zap.Field {

	return zapcore.Field{
		Key:       "resp_code",
		Type:      zapcore.Int64Type,
		Integer:   ld.RespCode,
		String:    "",
		Interface: nil,
	}
}

func (ld LogDetail) addResponseErrorMsg() zap.Field {

	return zapcore.Field{
		Key:       "resp_err",
		Type:      zapcore.StringType,
		Integer:   0,
		String:    ld.RespError,
		Interface: nil,
	}
}

// AddFields add extend fields to log record
// Generally speaking, the host applicationID instanceId are unchanged,
// so add host applicationID instanceID to logger recorder when init logger.

func (ld LogDetail) AddFields() []zap.Field {

	ldVal := reflect.ValueOf(ld)
	fieldNum := ldVal.NumField()

	ldFields := make([]zapcore.Field, 0, fieldNum)

	// hostField := ld.addHost()
	// ldFields = append(ldFields, hostField)

	// aidField := ld.addApplicationID()
	// ldFields = append(ldFields, aidField)

	// iidField := ld.addInstanceId()
	// ldFields = append(ldFields, iidField)

	costField := ld.addTimeCost()
	ldFields = append(ldFields, costField)

	reqIDField := ld.addRequestId()
	ldFields = append(ldFields, reqIDField)

	reqMethodField := ld.addRequestMethod()
	ldFields = append(ldFields, reqMethodField)

	reqPathField := ld.addRequestPath()
	ldFields = append(ldFields, reqPathField)

	clientAddrField := ld.addClientAddr()
	ldFields = append(ldFields, clientAddrField)

	respCodeField := ld.addResponseCode()
	ldFields = append(ldFields, respCodeField)

	respErrorField := ld.addResponseErrorMsg()
	ldFields = append(ldFields, respErrorField)

	return ldFields
}

// According to actual needs, add new k-v to the context of the logger recorder
// you can set key keyType and value(according to your value)

// NewFiled according to key, valueType, *value return a filed
func NewFiled(key string, valueType zapcore.FieldType, intValue int64, strValue string, otherValue interface{}) zap.Field {
	return zapcore.Field{
		Key:       key,
		Type:      valueType,
		Integer:   intValue,
		String:    strValue,
		Interface: otherValue,
	}
}
