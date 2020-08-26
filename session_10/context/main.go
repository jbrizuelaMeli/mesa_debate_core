package main

/*
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	defer func() {
		signal.Stop(channel)
		cancel()
	}()
	go func() {
		select {
			case <-channel:
				fmt.Println("Channel os.interrupt canceled")
				cancel()
			case <-ctx.Done():
				fmt.Printf("Context %s", ctx.Err().Error())
		}
	}()
	time.Sleep(5 * time.Second)
	//fmt.Println("Starting web server")
	//http.HandleFunc("/hi", handlerWeb)
	//http.ListenAndServe(":9090", nil)
}

func handlerWeb(w http.ResponseWriter, r *http.Request){
	defer fmt.Println("Ending handler")
	ctx := r.Context()

	// endpoint logic
	select {
		case <- time.After(5 * time.Second):
			fmt.Println("working")
		case <- ctx.Done():
			fmt.Printf("Ups: %s \n", ctx.Err().Error())
	}
}
 */