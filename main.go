package main

import "fmt"

func main() {
	// Задание 1
	// 1. true
	fmt.Println(5 == 5)
	// 2. true
	fmt.Println(10 != 3)
	// 3. false
	fmt.Println(7 > 12)
	// 4. true
	fmt.Println(15 < 20)
	// 5. true
	fmt.Println(8 >= 8)
	// 6. false
	fmt.Println(6 <= 4)
	// 7. false
	fmt.Println((10 > 5) && (3 < 1))
	// 8. true
	fmt.Println((10 > 5) || (3 < 1))
	// 9. false
	fmt.Println(!(5 == 5))
	// 10. true
	fmt.Println(!(7 < 3))
	// 11. false
	fmt.Println(true && false)
	// 12. false
	fmt.Println(false || false)
	// 13. true
	fmt.Println(true || false)
	// 14. true
	fmt.Println((4 + 6 == 10) && (9 > 2))
	// 15. true
	fmt.Println((12 / 3 == 4) || (8 < 5))
	
	// Задание 2
	age := 17
	var hasTicket bool = true
	var canEnter bool
	if age >= 18 && hasTicket {
		canEnter = true
	} else {
		canEnter = false
	}
	fmt.Println(canEnter)

	// Задание 3
	isLoggedIn := true
	isAdmin := false
	hasAccess := isLoggedIn && (isAdmin || !isAdmin)
	fmt.Println(hasAccess)
	
}
