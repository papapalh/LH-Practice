package main

import "fmt"

func main() {
	//sigRecv := make(chan os.Signal, 1)
	//
	//sigs := []os.Signal{syscall.SIGINT,syscall.SIGQUIT}
	//
	//signal.Notify(sigRecv, sigs...)
	//
	//for sig := range sigRecv {
	//	fmt.Println(sig)
	//}


	for i := 0; i < 10; i++ {
		go a(i)
	}

	for i := 20; i < 30; i++ {
		go func(a int) {
			fmt.Println(a)
		}(i)
	}

	for {

	}
}

func a(i int) {
	fmt.Println(i)
}
