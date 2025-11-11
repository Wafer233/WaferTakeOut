package result

type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func Success() Result[any] {
	return Result[any]{Code: 1, Msg: "success"}
}

func SuccessData[T any](data T) Result[T] {
	return Result[T]{Code: 1, Msg: "success", Data: data}
}

func Error(msg string) Result[any] {
	return Result[any]{Code: 0, Msg: msg}
}
