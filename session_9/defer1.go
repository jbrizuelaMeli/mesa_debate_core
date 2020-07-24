package session_9

func main() {
	defer func() {
		print(" World!")
	}()
	print("Hello")
}
