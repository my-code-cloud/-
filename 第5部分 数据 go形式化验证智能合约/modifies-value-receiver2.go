//revive:disable
package fixtures

type data struct {
	num   int
	key   []string
	items map[string]bool
}

//revive:enable:modifies-value-receiver
func (this data) vmethod() {
	//this.num = 8 // MATCH /suspicious assignment to a by-value method receiver/
	this.key[0] = "v.key"
	this.items["vmethod"] = true
}
