package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://twitter.com/?lang=es",
		"https://es.stackoverflow.com/",
		"https://www.amazon.com.mx/",
		"https://pkg.go.dev/std",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://twitter.com/?lang=es",
		"https://es.stackoverflow.com/",
		"https://www.amazon.com.mx/",
		"https://pkg.go.dev/std",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://twitter.com/?lang=es",
		"https://es.stackoverflow.com/",
		"https://www.amazon.com.mx/",
		"https://pkg.go.dev/std",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://twitter.com/?lang=es",
		"https://es.stackoverflow.com/",
		"https://www.amazon.com.mx/",
		"https://pkg.go.dev/std",
	}

	c := make(chan string)

	start := time.Now()
	for _, link := range links {
		go checkLink(link, c)
	}

	/*for _, link := range links {
		go func(linkparam string) {
			time.Sleep(5 * time.Second)
			checkLink(linkparam, c)
		}(link)

	}*/

	/*fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)*/
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	/*for l := range c {
		fmt.Println(l)
	}*/

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "Might be down")
		c <- "Might be down " + link
		return
	}
	c <- "All is ok " + link
	fmt.Println(link, " all is ok")
}
