package parallel

import (
	"fmt"
	"sync"
)

func logos1() bool {
	ch := make(chan bool)
	defer close(ch)
	go func() {
		fmt.Println("Hello World")
		ch <- true
	}()
	return <-ch
}

func logos2() []bool {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	defer close(ch1)
	defer close(ch2)
	defer close(ch3)
	res := []bool{false, false, false}

	go func() {
		fmt.Println("Hello World. Ch1 is OK")
		ch1 <- true
	}()
	res[0] = <-ch1

	go func() {
		fmt.Println("Hello World. Ch2 is OK")
		ch2 <- true
	}()
	res[1] = <-ch2

	go func() {
		fmt.Println("Hello World. Ch3 is OK")
		ch3 <- true
	}()
	res[2] = <-ch3

	return res
}

func logos3() []int {
	var wg sync.WaitGroup
	res := []int{0, 1, 2, 3}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			res[i] += i
		}(i)
	}
	wg.Wait()
	return res
}

func logos4() []int {
	var wg sync.WaitGroup
	res := []int{0, 1, 2, 3}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		// You cannot write "res[i] = go thread(&wg, i)"
		go func(i int) {
			// The wg variable is a pointer to pass by reference
			res[i] = thread(&wg, i)
		}(i)
	}
	wg.Wait()
	return res
}

func thread(wg *sync.WaitGroup, i int) int {
	defer wg.Done()
	return i + i
}

func logos5() int {
	c := 0
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ii := 0; ii < 10000; ii++ {
				c++
			}
			fmt.Println(c)
		}()
	}
	wg.Wait()
	return c

}

func logos6() int {
	c := 0
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ii := 0; ii < 10000; ii++ {
				c++
			}
			fmt.Println(c)
			ch <- c
		}()
		c = <-ch
	}
	wg.Wait()
	return c
}

func logos7() int {
	c := 0
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ii := 0; ii < 10000; ii++ {
				m.Lock()
				c++
				m.Unlock()
			}
			fmt.Println(c)
		}()
	}
	wg.Wait()
	return c
}
