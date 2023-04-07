package main

import (
	"HW9/Consts"
	"fmt"
	"log"
)

func main() {

	fmt.Println("Введите количество итераций от 1 до 100 включительно:")
	var T int
	var err error
	_, err = fmt.Scan(&T)
	if err != nil {
		log.Println(Consts.InErr)
		return
	}

	for i := 0; i < T; i++ {
		var S string
		//var array [4]rune
		fmt.Println("Введите строку из четырёх букв:")

		fmt.Scan(&S)
		if len(S) != 4 {
			fmt.Println(Consts.InErr)
		} else if len(S) == 4 && S[0] == S[1] || S[0] == S[2] || S[0] == S[3] && S[1] == S[2] || S[1] == S[3] || S[2] == S[3] {
			fmt.Println("Yes")

		} else {
			fmt.Println("NO")
		}

	}
	//var count int

	//var array [4]rune

	/*for _, k := range s {
			array[count] = k
			count++
			fmt.Println(k)
		}

		var (
			array1 [2]rune
			array2 [2]rune
			array3 [2]rune
			array4 [2]rune
		)

		array1[0] = array[0]
		array1[1] = array[1]
		array2[0] = array[2]
		array2[1] = array[3]
		array3[1] = array[0]
		array3[0] = array[1]
		array4[0] = array[3]
		array4[1] = array[2]

		//fmt.Println("Кодовые точки элементов массива рун в Unicode:", array)
		//fmt.Println("Кодовые точки элементов четырёх фрагментов массива рун:", array1, array2, array3, array4)
		if array1 == array2 && array3 == array4 && array1 == array3 && array2 == array4 {
			fmt.Println("No")
		} else if array1 == array2 || array3 == array4 || array1 == array3 || array2 == array4 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
		fmt.Println("---------------")
	}*/
}
