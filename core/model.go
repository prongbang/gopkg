package core

type Model[T any] struct {
	Type T
}

type AccessToken struct {
	Token string `json:"token"`
}

type RequestInfo[T any] struct {
	Claims *T
}

func (r *RequestInfo[T]) SetRequestInfo(info *RequestInfo[T]) {
	r.Claims = info.Claims
}

type Request[T any] interface {
	SetRequestInfo(info *RequestInfo[T])
}
