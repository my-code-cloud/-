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
		return // MATCH /unreachable code after this statement/
		Println("unreachable")
		os.Exit(2) // ignore
		Println("also unreachable")
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
			t.Fatal("i == 0") // MATCH /unreachable code after this statement/
			println("unreachable")
			continue
		}
	}
}
