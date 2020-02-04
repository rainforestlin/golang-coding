package main

import (
	"./runner"
	"log"
	"os"
	"time"
)

const timeout = 50 * time.Second

func main() {
	log.Print("Starting work")
	r := runner.New(timeout)
	r.Add(createTask(),createTask(),createTask(),createTask(),createTask(),createTask(),createTask(),createTask(),createTask(),createTask())
	if err:= r.Start();err!=nil{
		switch err {
		case runner.ErrInterrupt:
			log.Printf("Terminating due to interrupt\n")
			os.Exit(1)
		case runner.ErrTimeout:
			log.Printf("Terminating due to timeout\n")
			os.Exit(2)

		}
	}
	log.Println("Process ended")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor-Task ### %d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
