package session_9

func main() {
	defer println("6") //Primera
	defer func() { //Segunda
		println("1")
		defer println("5") //Tercera
		defer func() { //Cuarta
			println("3")
			defer println("4") //Quinta
		}()
		defer println("2") //Sexta
	}()
}
