package main

import (
	"fmt"
	"runtime"
	"os/exec"
	"sync"
	"strconv"
	"strings"
)

func main(){
	plain_cipher := []string{"affine", "ihhwvc"}
	affineCmd    := "./affine"
	encType      := []string{"encrypt", "decrypt"}

	//Common Channel for the goroutines
	tasks := make(chan *exec.Cmd, 64)

	//Create a WaitGroup for the goroutines
	var wg sync.WaitGroup
	//Getting Max number of goroutines
	cores := MaxParallelism()

	for i := 0; i < cores; i++ {
		//Add a new goroutine to the Waitgroup
		wg.Add(1)
		//Call a new Goroutine
		go func(num int, w *sync.WaitGroup) {
			//Until done
			defer w.Done()
			//Initialize a byte container
			var (
				out []byte
				err error
			)
			//Execute a passed Command
			for cmd := range tasks {
				//Get its output
				out, err = cmd.Output()
				//If error
				if err != nil {
					fmt.Printf("Can't get stdout:", err)
				}
				//Cast output to string
				s:= string(out)
				res := strings.Fields(s)

				if len(res) == 0 {
					fmt.Println("No result obtained")
				} else if res[13] == plain_cipher[1] {
					fmt.Println("Found keys, a: "+cmd.Args[2]+" b: "+cmd.Args[3]+" that encrypt "+cmd.Args[4]+" into "+res[13])
				}
			}
		}(i, &wg)
	}

	for i := 1; i < 26; i++ {
		if i%2 == 0 {
			continue
		}
		a := i
		for j := 1; j < 26; j++ {
			b := j
			//Execute Encrypt
			tasks <- exec.Command(affineCmd, encType[0], strconv.Itoa(int(a)), strconv.Itoa(int(b)), plain_cipher[0])
		}
	}
}

//Gets Max number of goroutines to get Called
func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU   := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}