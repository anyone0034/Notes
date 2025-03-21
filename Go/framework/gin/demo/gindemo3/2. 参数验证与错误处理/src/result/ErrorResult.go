package result

import "fmt"

type ErrorResult struct {
	data interface{}
	err  error
}

func (this *ErrorResult) Unwrap() interface{} {
	if this.err != nil {
		panic(this.err.Error())
	}
	return this.data

}
func Result(vs ...interface{}) *ErrorResult {
	if len(vs) == 1 {
		if vs[0] == nil {
			return &ErrorResult{nil, nil}
		}
		if e, ok := vs[0].(error); ok {
			return &ErrorResult{nil, e}
		}
	}
	if len(vs) == 2 {
		if vs[1] == nil {
			return &ErrorResult{vs[0], nil}
		}
		if e, ok := vs[1].(error); ok {
			return &ErrorResult{vs[0], e}
		}
	}
	return &ErrorResult{nil, fmt.Errorf("error result format")}
}
