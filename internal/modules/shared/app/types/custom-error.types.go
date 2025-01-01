package types

type HttpErrorArgs struct {
	StatusCode int
	Message    string
	Details    *JsonMap
}
