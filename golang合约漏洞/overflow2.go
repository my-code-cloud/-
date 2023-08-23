package fixtures

func main() {
	var a int32 = 3000000000
	var b int8
	b = int8(a) // 可能导致溢出

	var x int16 = 20000
	var y int8
	y = int8(x) // 可能导致溢出

	var u uint32 = 4294967295
	var v int32
	v = int32(u) // 可能导致负数
}