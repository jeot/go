package main

import "fmt"

func main() {
	fmt.Println("hello, world\n")

	const pi float64 = 3.1415
	var favoritNumber float32 = 3.14
	var age = 31
	var hight int = 171 // cm
	randNum := 1.33;
	fmt.Println(hight, favoritNumber, randNum, age)
	fmt.Printf("PI = %f \n", pi)
	fmt.Printf("PI = %.2f \n", pi)
	fmt.Printf("Type of pi var is %T \n", pi)

	var name string = "Shamim Keshani"
	fmt.Println(name, "Length:", len(name))
	fmt.Printf("Type of name var is %T \n", name)
	
	var isOver40 bool = false;
	fmt.Printf("Type of isOver40=%t var is %T \n", isOver40, isOver40)
	
	fmt.Printf("dec:  %d \n", 100)
	fmt.Printf("bin:  %b \n", 100)
	fmt.Printf("char: %c \n", 100)
	fmt.Printf("hex:  %x \n", 100)
	fmt.Printf("sci:  %e \n", 100.0)

	i := 0;
	for i < 10 {
		fmt.Println(i);
		i++
	}
	for j := 0; j != 5; j++ {
		fmt.Println(j)
	}

	var yourAge int = 5
	switch yourAge {
		case 16 : fmt.Println("go drive")
		case 18 : fmt.Println("go vote")
		default : fmt.Println("go have fun")
	}

	// Arrays:
	var favNumsArray1[3] float64
	favNumsArray1[0] = 16
	favNumsArray1[1] = 3.3
	favNumsArray1[2] = 3.32

	favNumsArray2 := [5]int {11, 22, 33, 44, 55}
	for i, value := range favNumsArray2 {
		fmt.Printf("i=%d, value=%d\n", i, value)
	}
}
