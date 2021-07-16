package constant

type ErrorStatus struct {
    StatusCode int32
    StatusMsg  string
}

var (
    SUCCESS              = &ErrorStatus{StatusCode: 0, StatusMsg: "Success"}
    ERR_WRONG_PASSWORD   = &ErrorStatus{StatusCode: 2, StatusMsg: "Wrong Password"}
    ERR_RECORD_NOT_FOUND = &ErrorStatus{StatusCode: 4, StatusMsg: "Record Not Found"}
    ERR_SERVICE_INTERNAL = &ErrorStatus{StatusCode: 6, StatusMsg: "Internal Error"}
)
