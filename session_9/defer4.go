package session_9

func main() {
	for i := 0; i < 10; i++ {
		//n := i
		defer func() {
			//println(n)
			println(i)
		}()
	}
}
