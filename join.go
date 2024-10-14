package tyerrors

import "unsafe"

type joinError struct {
	errs []error
}

func Join(errs ...error) error {
	// TODO errsが空の場合
	// TODO errsの要素がnilの場合
	return &joinError{errs: errs}
}

// TODO format指定されている場合の実装
// 中身がtyerrorsかどうかを判断する必要があるのかな
func (e *joinError) Error() string {
	if len(e.errs) == 1 {
		return e.errs[0].Error()
	}

	b := []byte(e.errs[0].Error())
	for _, err := range e.errs[1:] {
		b = append(b, '\n')
		b = append(b, err.Error()...)
	}
	return unsafe.String(&b[0], len(b))
}

func (e *joinError) Unwrap() []error {
	return e.errs
}
