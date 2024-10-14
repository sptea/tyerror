package tyerrors

import (
	"errors"
	"testing"
)

func TestJoin(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")

	joinedErr := Join(err1, err2, err3)

	if joinedErr == nil {
		t.Fatalf("expected non-nil error, got nil")
	}

	if joinedErr.Error() != "multiple errors occurred" {
		t.Errorf("expected 'multiple errors occurred', got '%s'", joinedErr.Error())
	}

	unwrappedErrs := joinedErr.(*joinError).Unwrap()
	if len(unwrappedErrs) != 3 {
		t.Fatalf("expected 3 errors, got %d", len(unwrappedErrs))
	}

	expectedErrs := []error{err1, err2, err3}
	for i, err := range unwrappedErrs {
		if err != expectedErrs[i] {
			t.Errorf("expected error %d to be '%v', got '%v'", i, expectedErrs[i], err)
		}
	}
}
