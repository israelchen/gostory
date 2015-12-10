package util

func Require(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}
