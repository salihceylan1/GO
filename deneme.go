// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	//first code
	fmt.Println("Hello, 世界")

	//variables
	var x int = 10
	y := 20 // Tipi otomatik belirler

	fmt.Println(y)

	//loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//conditions
	if x > 5 {
		fmt.Println("x büyük")
	} else {
		fmt.Println("x küçük veya eşit")
	}

	//usage of struct ve objects
	kisi := Kisi{Ad: "Ahmet", Yas: 25}
	fmt.Println(kisi)

	go goroutineDemo()
	fmt.Println("Ana program")

}

//Struct and object
type Kisi struct {
	Ad  string
	Yas int
}

//functions
func toplama(a int, b int) int {
	return a + b
}

func goroutineDemo() {
	fmt.Println("Goroutine çalışıyor")
}
