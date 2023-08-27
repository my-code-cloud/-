//revive:disable
package fixtures

import "fmt"

//revive:enable:range-val-in-closure
func foo() {
	mySlice := []string{"A", "B", "C"}
	for index, value := range mySlice {
		fmt.Printf("Index: %d\n", index)
		fmt.Printf("Value: %s\n", value)
	}
}
