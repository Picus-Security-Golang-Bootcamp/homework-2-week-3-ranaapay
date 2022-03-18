# WaitGroups in Golang

Goroutines are a very strong part of Golang.  How can we use them effectively?
Let's look at a simple code example here.

```
package main

import (
	"fmt"
)

func writeNum(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

func runIt() {
	fmt.Println("runIt func")
}

func main() {
	go writeNum(10)
	go runIt()
}
```
What do we expect from the output? Nothing sees in the console. Why? Because the main function itself is Goroutine, and it ends, it doesn't wait for other goroutines to finish.
We can solve this problem with use the Golang standard library ```sync.WaitGroup``` WaitGroup. WaitGroup works like a counter, it blocks some of the code until the Goroutines defined for it are finished.
WaitGroup has three methods that we can use.
* Add(int) → We say that Goroutines works like counter. When we initialize a WaitGroup it came with zero. We should specify how much Goroutine programs will wait. We use Add(int). It increases the counter by a given integer value. Before we execute the goroutine we should call Add(1). Or it can cause a panic.
* Done() → After goroutine execution done we call Done() func and it decrease counter.
* Wait() → It blocks the code until counter value become zero.

Let's write some code.

```
package main

import (
"fmt"
"sync"
)

func print(wg *sync.WaitGroup) {
fmt.Println("second goroutine")
wg.Done()
}

func main() {
wg := sync.WaitGroup{}

	wg.Add(2)

	go func(num int) {
		for i := 0; i < num; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}(5)

	go print(&wg)

	wg.Wait()
	fmt.Println("Done")
}
```
First, we declare WaitGroup with wg := sync.WaitGroup. Then we add two to WaitGroup. We say that we will wait until two goroutines are done. After goroutine execution is done we call Done() functions.
WaitGroup is concurrency safe.  

* If we want to pass the WaitGroup as a parameter to a function, we must do so using a pointer.



________________________________________________
### Source
* https://www.geeksforgeeks.org/using-waitgroup-in-golang/?ref=lbp
* https://gobyexample.com/waitgroups
* https://www.reply.com/alpha-reply/en/content/go-concurrency-with-waitgroup