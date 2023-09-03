//revive:disable
package fixtures

//revive:enable:defer
func deferrer() {

	helper := func(_ interface{}) {}

	defer func() {
		recover()
	}()

}
