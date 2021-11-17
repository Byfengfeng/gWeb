package enum

type ContentType string

const(
	Json = "application/json; charset=utf-8"
	From = "application/x-www-form-urlencoded"
)

type ReqType string
const(
	Get ReqType = "GET"
	Post ReqType = "POST"
)

func (req ReqType) GetString() string {
	return string(req)
}