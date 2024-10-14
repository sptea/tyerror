package tyerrors

import (
	"fmt"
	"runtime"
)

// TODO プロパティがこれで良いかの検討
type TyError struct {
	msg   string
	cause error // 元のエラーを保持するフィールドを追加
	trace []uintptr
}

func New(message string) *TyError {
	return &TyError{
		msg:   message,
		trace: captureStackTrace(),
	}
}

func (e *TyError) Error() string {
	return e.msg
}

func (e *TyError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%s\n", e.msg)
			frames := runtime.CallersFrames(e.trace)
			for {
				frame, more := frames.Next()
				fmt.Fprintf(s, "%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)
				if !more {
					break
				}
			}
			return
		}
		fallthrough
	case 's':
		fmt.Fprint(s, e.msg)
	case 'q':
		fmt.Fprintf(s, "%q", e.msg)
	}
}

func (e *TyError) Unwrap() error {
	return e.cause
}

func Wrap(err error, message string) *TyError {
	return &TyError{
		msg:   fmt.Sprintf("%s: %v", message, err),
		cause: err,
		trace: captureStackTrace(),
	}
}

func captureStackTrace() []uintptr {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // 呼び出し元のフレームをスキップ
	return pcs[:n]
}
