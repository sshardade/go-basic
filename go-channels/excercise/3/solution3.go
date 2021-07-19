package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	letters = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	numbers = `0123456789`
	special = `-=~!@#$%^&*()_+[]\{}|;':",./<>?`
)

var (
	l    int  //length for password
	help bool // to print help
	s    = rand.NewSource(time.Now().Unix())
	r    = rand.New(s)
)

func init() {
	flag.BoolVar(&help, "help", false, "show help")
	flag.IntVar(&l, "l", 16, "Length for random password generation [8-255]..")
	flag.Parse()
}

func main() {
	fmt.Println("Random password Generation..!!")
	if help || l < 8 || l > 255 {
		flag.Usage()
		return
	}

	ch := passwordGen(l)
	fmt.Printf("Random password: ")
	for v := range ch {
		fmt.Printf("%c", v)
	}
	fmt.Println()

}

func passwordGen(l int) (out chan rune) {
	out = make(chan rune, l)

	var (
		lettersRune = []rune(letters)
		numbers     = []rune(numbers)
		special     = []rune(special)

		lenLettersRune = len(lettersRune)
		lenNumbers     = len(numbers)
		lenSpecial     = len(special)
	)

	for i := 0; i < l; i++ {
		select {
		case out <- lettersRune[r.Int()%lenLettersRune]:
		case out <- lettersRune[r.Int()%lenLettersRune]:
		case out <- lettersRune[r.Int()%lenLettersRune]:
		case out <- lettersRune[r.Int()%lenLettersRune]:
		case out <- lettersRune[r.Int()%lenLettersRune]:
		case out <- lettersRune[r.Int()%lenLettersRune]:
		case out <- lettersRune[r.Int()%lenLettersRune]:
		case out <- lettersRune[r.Int()%lenLettersRune]:

		case out <- numbers[r.Int()%lenNumbers]:
		case out <- numbers[r.Int()%lenNumbers]:
		case out <- numbers[r.Int()%lenNumbers]:

		case out <- special[r.Int()%lenSpecial]:
		}
	}
	close(out)
	return
}
