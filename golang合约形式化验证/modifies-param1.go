//revive:disable
package fixtures

//revive:enable:modifies-parameter
func two(b, c float32) {
	if c > 0.0 {
		b *= 2 // MATCH /parameter 'b' seems to be modified/
	}
}
