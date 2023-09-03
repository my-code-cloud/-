//revive:disable
package fixtures

//revive:enable:unconditional-recursion
func ur2tris() {
	for true == false {
		println()
		ur2tris()
	}
}
