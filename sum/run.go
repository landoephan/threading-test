package main

var length = 10_000_000_000

func evenNumberSum() {
	sum := 0
	for i := 0; i <= length; i++ {
		if i%2 == 0 {
			sum = sum + i
		}
	}
}

func oddNumberSum() {
	sum := 0
	for i := 0; i <= length; i++ {
		if i%2 == 1 {
			sum = sum + i
		}
	}
}

func main() {
	for j := 0; j < 10; j++ {
		evenNumberSum()
		oddNumberSum()
	}
}
