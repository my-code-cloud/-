//revive:disable
package fixtures

import "strconv"

//revive:enable:string-of-int
func StringTest() {
	i := 84
	_ = strconv.Itoa(i) // MATCH /dubious conversion of an integer into a string, use strconv.Itoa/
}
