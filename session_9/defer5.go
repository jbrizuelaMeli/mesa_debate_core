package session_9

func main() {
	defer println("Third")
	defer println("Second")
	defer println("First")
}