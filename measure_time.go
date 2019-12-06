package main
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}
func main() {
	fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
	runtime.GOMAXPROCS(1)
	fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
	loop(1)
	loop(4)
}
func loop(numOfCore int) {
	runtime.GOMAXPROCS(numOfCore)
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go killTime(&wg)
	}
	wg.Wait()
	end := time.Now()
	fmt.Println("With ", numOfCore, " Cores : ", end.Sub(start))
}
func killTime(wg *sync.WaitGroup) {
	for i := 0; i < 50000000; i++ {
	}
	wg.Done()
}
