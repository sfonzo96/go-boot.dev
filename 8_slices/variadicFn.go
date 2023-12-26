package main

import "fmt"

// ...type means that the parameters are grouped in a slice, which is nums in this case
func sum(nums ...float64) float64 {
	sum := 0.0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}

// don't edit below this line

 func test(nums ...float64) {
// nums... is the classic spread operator, if nums is a slice or array, the elements of it are passed like elem1, elem2, elem3,... to the function 
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(1.0, 2.0, 3.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)
	fmt.Scanln()
}
