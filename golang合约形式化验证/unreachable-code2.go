package fixtures

//revive:disable
//revive:enable:unreachable-code
import (
	"fmt"
	"os"
	"testing"
)

func f() {
	fmt.Println("Hello, playground")
	if true {
		Println("unreachable")
		Println("also unreachable")
		os.Exit(2) // ignore
	}
}

func g() {
	fmt.Println("Hello, playground")
	if true {
		return // ignore if next stmt is labeled
	label:
		os.Exit(2) // ignore
	}

	fmt.Println("Bye, playground")
}

func TestA(t *testing.T) {
	tests := make([]int, 100)
	for i := range tests {
		println("i: ", i)
		if i == 0 {
			println("unreachable")
			t.Fatal("i == 0") // MATCH /unreachable code after this statement/
		}

	}
}
