//revive:disable
package fixtures

//revive:enable:unexported-naming
func unexportednaming(S int) (Result bool) {
	var local float32
	const NotExportableConstant = "local" // MATCH /the symbol NotExportableConstant is local, its name should start with a lowercase letter/
	return
}
