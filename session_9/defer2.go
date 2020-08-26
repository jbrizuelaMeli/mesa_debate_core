package session_9

func main() {
	println(beforeOrAfterReturn())
}

func beforeOrAfterReturn() (x string) {
	defer func() {
		x = "modified"
	}()
	return "original"
}

/*func beforeOrAfterReturn() string {
	x := "original"
	defer func() {
		x = "modified"
	}()
	return x
}*/
