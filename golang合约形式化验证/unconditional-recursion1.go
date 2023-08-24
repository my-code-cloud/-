//revive:disable
package fixtures

//revive:enable:unconditional-recursion
func ur2tris() {
	for {
		println()
		ur2tris() // MATCH /unconditional recursive call/
	}
}
