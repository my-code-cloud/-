//revive:disable
package fixtures

//revive:enable:modifies-parameter
func two(b, c float32) float32 {
	x := b
	if c > 0.0 {
		x *= 2 // MATCH /parameter 'b' seems to be modified/
	}
	return x
}
