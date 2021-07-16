package constant

type PasterStatus struct {
    StatusCode int32
    StatusMsg  string
}

var (
    SUCCESS              = &PasterStatus{StatusCode: 0, StatusMsg: "Success"}
    ERR_WRONG_PASSWORD   = &PasterStatus{StatusCode: 2, StatusMsg: "Wrong Password"}
    ERR_RECORD_NOT_FOUND = &PasterStatus{StatusCode: 4, StatusMsg: "Record Not Found"}
    ERR_SERVICE_INTERNAL = &PasterStatus{StatusCode: 6, StatusMsg: "Internal Error"}
)
