package util

import (
    "reflect"

    "github.com/ameidance/paster_core/constant"
)

func FillBizResp(resp interface{}, status *constant.PasterStatus) {
    if resp == nil {
        return
    }
    if status == nil {
        status = constant.ERR_SERVICE_INTERNAL
    }

    respType := reflect.TypeOf(resp)
    respVal := reflect.ValueOf(resp)
    if respType.Kind() == reflect.Ptr {
        respVal = respVal.Elem()
    }

    respVal.FieldByName("StatusCode").SetInt(int64(status.StatusCode))
    respVal.FieldByName("StatusMessage").SetString(status.StatusMsg)
}
