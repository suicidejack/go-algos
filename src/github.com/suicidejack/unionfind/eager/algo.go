package eager

/************************************************************************
 * Author: Matt Price
 * Last Modified: Thu Jun 19, 2014 18:41:53 EDT
 * Purpose: Algorithm for the quick-find union find
 ************************************************************************/

import (
	"fmt"
	"strconv"
)

var id []int

func Init(numElements int) {
	id = make([]int, numElements)
	for i := 0; i < numElements; i++ {
		id[i] = i
	}
}

func String() (str string) {
	str = ""
	for _, val := range id {
		str += strconv.Itoa(val) + " "
	}
	return fmt.Sprintf("%s", str)
}