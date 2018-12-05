package functions

func Fib(position int8) int64 {
	var last int64 = 1
	var cur int64 = 1
	for i := 1; i < int(position); i++ {
		var tmp = cur
		cur += last
		last = tmp
	}
	return last
}
