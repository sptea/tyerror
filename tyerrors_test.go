package tyerrors_test

import (
	"errors"
	"testing"

	"github.com/sptea/tyerrors"
)

func TestNew(t *testing.T) {
	err := tyerrors.New("test error")

	// Error() メソッドが期待通りのメッセージを返すか確認
	if err.Error() != "test error" {
		t.Errorf("expected 'test error', got '%s'", err.Error())
	}

	// Unwrap() が nil を返すことを確認
	if errors.Unwrap(err) != nil {
		t.Errorf("expected nil from Unwrap(), got non-nil")
	}
}

func TestWrap(t *testing.T) {
	// テストケースを定義
	tests := []struct {
		name        string
		baseErr     error
		context     string
		wantMessage string
		wantUnwrap  error
	}{
		{
			name:        "Wrapped error with base error",
			baseErr:     errors.New("base error"),
			context:     "additional context",
			wantMessage: "additional context: base error",
			wantUnwrap:  errors.New("base error"), // 元のエラーが返されることを期待
		},
		{
			name:        "Wrapped error with no base error",
			baseErr:     nil,
			context:     "context only",
			wantMessage: "context only: <nil>", // エラーがnilでもコンテキストが追加される
			wantUnwrap:  nil,                   // Unwrapはnilを返す
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrappedErr := tyerrors.Wrap(tt.baseErr, tt.context)

			// Error() のメッセージをテスト
			if wrappedErr.Error() != tt.wantMessage {
				t.Errorf("expected error message '%s', got '%s'", tt.wantMessage, wrappedErr.Error())
			}

			// Unwrap() の結果をテスト
			if unwrappedErr := errors.Unwrap(wrappedErr); unwrappedErr != tt.wantUnwrap {
				if tt.wantUnwrap == nil && unwrappedErr != nil {
					t.Errorf("expected Unwrap() to return nil, got '%v'", unwrappedErr)
				} else if tt.wantUnwrap != nil && unwrappedErr.Error() != tt.wantUnwrap.Error() {
					t.Errorf("expected Unwrap() to return '%v', got '%v'", tt.wantUnwrap, unwrappedErr)
				}
			}
		})
	}
}
