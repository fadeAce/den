package main

import (
	"fmt"
	"os"
	"os/exec"
	"bytes"
	"time"
)

func main() {

	path := os.Getenv("DEN")

	fmt.Println(path)

	pull := exec.Command("docker-compose", "-f", path, "pull")
	down := exec.Command("docker-compose", "-f", path, "down")
	up := exec.Command("docker-compose", "-f", path, "up")

	ch := make(chan int)
	pass := make(chan int)
	end := make(chan int)
	go func() {
		down.Run()
		ch <- 0
		pass <- 0
	}()
	go func() {
		<-pass
		pull.Run()
		ch <- 0
		end <- 0
	}()
	go func() {
		<-end
		up.Start()
		time.Sleep(8 * time.Second)
		ch <- 0
	}()

	for i := 0; i < 100; i++ {
		if i == 34 || i == 67 || i == 99 {
			<-ch
		}
		fmt.Print("\r")
		for v := 0; v < 100; v++ {
			if v < i {
				fmt.Print("▇")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print(" ", i+1, "%", getStr(i))
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Print("\r")
	fmt.Print("den done")
	fmt.Println("")

	cmd := exec.Command("docker-compose", "-f", path, "ps")
	var out bytes.Buffer
	cmd.Stdout = &out

	cmd.Run()
	fmt.Println(out.String())
	close(ch)
}

func getStr(i int) string {
	if i < 34 {
		return " down"
	}
	if i > 66 {
		return " up  "
	}
	return " pull"
}
