package tyerrors_test

import (
	"fmt"

	"github.com/sptea/tyerrors"
)

func ExampleNew() {
	err := tyerrors.New("test message")
	fmt.Println(err)

	// Output: test message
}

func ExampleNew_printf() {
	err := tyerrors.New("root message")
	fmt.Printf("%+v", err)

	// Example Output:
	// root message
	// github.com/sptea/tyerrors_test.ExampleNew_printf
	// /home/tyerrors/example_test.go:17
	// testing.runExample
	// 		/usr/lib/go/src/testing/run_example.go:63
	// testing.runExamples
	// 		/usr/lib/go/src/testing/example.go:40
	// testing.(*M).Run
	// 		/usr/lib/go/src/testing/testing.go:2036
	// main.main
	// 		_testmain.go:53
	// runtime.main
	// 		/usr/lib/go/src/runtime/proc.go:272
	// runtime.goexit
	//		/usr/lib/go/src/runtime/asm_arm64.s:1223
}
