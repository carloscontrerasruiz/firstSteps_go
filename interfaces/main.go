package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnglishBot struct{}

type SpanishBot struct{}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	//printGreetingS(sb)
	printGreeting(sb)
}

func (EnglishBot) getGreeting() string {
	//if we dont use the recevier value
	//we can just put the type
	//custom logic for englishbot
	return "Hi There"
}

func (SpanishBot) getGreeting() string {
	//custom logic for spanishbot
	return "Hola !!!!!"
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

/*func printGreeting(eb EnglishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingS(sp SpanishBot) {
	fmt.Println(sp.getGreeting())
}*/

type Shape interface {
	getArea() float64
}

type Triangle struct {
	base   float64
	height float64
}

type Square struct {
	sideLength float64
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}
