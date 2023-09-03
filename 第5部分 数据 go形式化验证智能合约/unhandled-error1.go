//revive:disable
package fixtures

import (
	"fmt"
	"os"
)

func unhandledError1(a int) (int, error) {
	return a, nil
}

//revive:enable:unhandled-error
func unhandledError2() {
	unhandledError1(1)   // MATCH /Unhandled error in call to function unhandledError1/
	fmt.Fprintf(nil, "") // MATCH /Unhandled error in call to function fmt.Fprintf/
	os.Chdir("..")       // MATCH /Unhandled error in call to function os.Chdir/
}
