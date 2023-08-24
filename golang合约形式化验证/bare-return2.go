//revive:disable
package fixtures

//revive:enable:bare-return
func f280_4() (err error) {
	func() (r int) {
		return func() (r int) {
			return 0 // MATCH /avoid using bare returns, please add return expressions/
		}()
	}()

	return nil
}
