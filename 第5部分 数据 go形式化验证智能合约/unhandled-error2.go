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
func unhandledError2() error {
	_, err := unhandledError1(1)
	err = fmt.Fprintf(nil, "") // MATCH /Unhandled error in call to function fmt.Fprintf/
	err = os.Chdir("..")       // MATCH /Unhandled error in call to function os.Chdir/
}
