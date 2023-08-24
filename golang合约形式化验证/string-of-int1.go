//revive:disable
package fixtures

//revive:enable:string-of-int
func StringTest() {
	i := 84
	_ = string(i) // MATCH /dubious conversion of an integer into a string, use strconv.Itoa/
}
