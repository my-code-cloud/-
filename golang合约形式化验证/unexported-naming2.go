//revive:disable
package fixtures

const NotExportableConstant = "local"

var S int
var Result bool

//revive:enable:unexported-naming
func unexportednaming(s int) (result bool) {
	var local float32
	S = s
	Result = result
	return
}
