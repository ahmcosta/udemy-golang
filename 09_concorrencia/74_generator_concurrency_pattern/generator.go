/*
	Google I/O 2012 - Go Concurrency Patterns
	Generator Pattern: function that returns a channel. That encapsules a goroutine call!
	More info in: https://talks.golang.org/2012/concurrency.slide#25
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

/*
	This function expects a URLs array. For each URL, it gets the page, cuts its body, search its title using regex
*/
func titulo(urls ...string) <-chan string { // <-chan - canal somente-leitura
	c := make(chan string)
	for _, url := range urls {
		go func(url string) { // anonymous function
			beginTime := time.Now()
			resp, _ := http.Get(url)             // get the page by url
			html, _ := ioutil.ReadAll(resp.Body) // get the body of the page gotten

			r, _ := regexp.Compile("<title>(.*?)<\\/title>") // compile the regex
			title := r.FindStringSubmatch(string(html))[1]
			finishTime := time.Now()
			totalTime := finishTime.Sub(beginTime)
			c <- fmt.Sprintf("Elapsed Time: %d\tTitle: %s", totalTime, title) // return the string from compiled regex
		}(url) // It's needed because it's a anonymous functions. So it must be called
	}
	return c
}

// This function doesn't worry about the goroutines. This is the generator patterns
func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
