package main

import (
	"fmt"
	"strconv"
)

//init default numbers

//template
/* var items = [9][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
} */
/* var items = [9][9]int{
	{0, 9, 1, 0, 0, 0, 2, 3, 0},
	{7, 0, 0, 4, 0, 5, 0, 0, 6},
	{0, 0, 0, 0, 1, 0, 0, 0, 0},
	{9, 5, 0, 0, 0, 0, 0, 2, 7},
	{0, 0, 0, 8, 0, 2, 0, 0, 0},
	{6, 1, 0, 0, 0, 0, 0, 4, 3},
	{0, 0, 0, 0, 5, 0, 0, 0, 0},
	{8, 0, 0, 9, 0, 1, 0, 0, 4},
	{0, 3, 6, 0, 0, 0, 7, 9, 0},
} */
/* var items = [9][9]int{
	{0, 0, 1, 0, 0, 7, 5, 0, 0},
	{0, 0, 3, 4, 0, 0, 0, 0, 8},
	{2, 0, 0, 0, 1, 0, 0, 3, 0},
	{0, 5, 0, 0, 0, 2, 9, 0, 0},
	{9, 0, 0, 0, 5, 0, 0, 0, 2},
	{0, 0, 6, 3, 0, 0, 0, 7, 0},
	{0, 3, 0, 0, 8, 0, 0, 0, 4},
	{4, 0, 0, 0, 0, 9, 7, 0, 0},
	{0, 0, 7, 2, 0, 0, 1, 0, 0},
} */

/* var items = [9][9]int{
	{0, 5, 0, 0, 0, 0, 0, 6, 0},
	{8, 0, 0, 4, 0, 5, 0, 0, 1},
	{0, 0, 2, 1, 0, 7, 4, 0, 0},
	{0, 6, 3, 0, 0, 0, 9, 2, 0},
	{0, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 9, 7, 0, 0, 0, 5, 1, 0},
	{0, 0, 8, 5, 0, 6, 2, 0, 0},
	{7, 0, 0, 8, 0, 4, 0, 0, 6},
	{0, 4, 0, 0, 0, 0, 0, 7, 0},
} */

var items = [9][9]int{
	{0, 0, 1, 0, 0, 7, 5, 0, 0},
	{0, 0, 3, 4, 0, 0, 0, 0, 8},
	{2, 0, 0, 0, 1, 0, 0, 3, 0},
	{0, 5, 0, 0, 0, 2, 9, 0, 0},
	{9, 0, 0, 0, 5, 0, 0, 0, 2},
	{0, 0, 6, 3, 0, 0, 0, 7, 0},
	{0, 3, 0, 0, 8, 0, 0, 0, 4},
	{4, 0, 0, 0, 0, 9, 7, 0, 0},
	{0, 0, 7, 2, 0, 0, 1, 0, 0},
}

type ItemData struct {
	items [9][9][10]int
}

//itemInfo is item info
var itemInfo [9][9][10]int

//set11 is seting no11
func set11(it ItemData) ItemData {

	for ri, rowItem := range it.items {
		for ci, colItem := range rowItem {
			if colItem[0] >= 1 && colItem[0] <= 9 {
				for i := 0; i < 9; i++ {
					//myself
					it.items[ri][ci][i+1] = 11
					//row set
					it.items[ri][i][colItem[0]] = 11
					//col set
					it.items[i][ci][colItem[0]] = 11
				}
				//box set
				_, startRi, endRi, startCi, endCi := getBox(100, ri, ci)
				for ri := startRi; ri <= endRi; ri++ {
					for ci := startCi; ci <= endCi; ci++ {
						it.items[ri][ci][colItem[0]] = 11
					}
				}
			}
		}
	}

	return it
}

//checkRow2 is searching same number in the one Row
func checkRow2(it ItemData) ItemData {
	for index := 0; index < 9; index++ {
		for number := 1; number <= 9; number++ {
			var tmp [9]int
			_, tmp = getOne(it, "row", "one", index, number)
			b, sameIndex := checkSame(tmp)
			if b == true {
				it.items[index][sameIndex][0] = number
				it = set11(it)
			}
		}
	}

	return it
}

//checkSame is searching same data
func checkSame(items [9]int) (check bool, index int) {
	checkData := [2]int{0, 0}
	for i, item := range items {
		if item == 10 {
			checkData[0] = checkData[0] + 1
			checkData[1] = i
		}
	}
	if checkData[0] == 1 {
		check = true
		index = checkData[1]
	} else {
		check = false
		index = 100
	}
	return
}

//checkSameBox is searching same data
func checkSameBox(it ItemData, index, number int) (check bool, rIndex, cIndex int) {
	checkData := [3]int{0, 0, 0}

	_, startRi, endRi, startCi, endCi := getBox(index, 100, 100)
	for ri := startRi; ri <= endRi; ri++ {
		for ci := startCi; ci <= endCi; ci++ {
			if it.items[ri][ci][number] == 10 {
				checkData[0] = checkData[0] + 1
				checkData[1] = ri
				checkData[2] = ci
			}
		}
	}
	if checkData[0] == 1 {
		check = true
		rIndex = checkData[1]
		cIndex = checkData[2]
	} else {
		check = false
		rIndex = 100
		cIndex = 100
	}
	return
}

//checkCol2 is searching same number in the one Col
func checkCol2(it ItemData) ItemData {
	for index := 0; index < 9; index++ {
		for number := 1; number <= 9; number++ {
			var tmp [9]int
			_, tmp = getOne(it, "col", "one", index, number)
			b, sameIndex := checkSame(tmp)
			if b == true {
				it.items[sameIndex][index][0] = number
				set11(it)
			}
		}
	}

	return it
}

//checkCol2 is searching same number in the one Col
func checkBox2(it ItemData) ItemData {
	for index := 1; index <= 9; index++ {
		for number := 1; number <= 9; number++ {
			b, rIndex, cIndex := checkSameBox(it, index, number)
			if b == true {
				it.items[rIndex][cIndex][0] = number
				set11(it)
			}
		}
	}

	return it
}

//getRow return row data
func getRow(it ItemData) ItemData {
	var rowData ItemData
	for ri, rowItem := range it.items {
		for ci, item := range rowItem {
			for di, d := range item {
				rowData.items[ri][ci][di] = d
			}
		}
	}
	return rowData
}

//getOne return row or col one data
func getOne(it ItemData, typ, sel string, index, number int) (data [9][10]int, oneData [9]int) {

	var tmp ItemData

	switch typ {
	case "row":
		tmp = getRow(it)
	case "col":
		tmp = getCol(it)
	default:
	}

	switch sel {
	case "all":
		for ci, colItem := range tmp.items[index] {
			for vi, value := range colItem {
				data[ci][vi] = value
			}
		}
	case "one":
		for ci, colItem := range tmp.items[index] {
			if ci < 9 {
				oneData[ci] = colItem[number]
			}
		}
	default:
	}
	return
}

//getCol return coll data
func getCol(it ItemData) ItemData {
	var colData ItemData
	for ri, rowItem := range it.items {
		for ci, item := range rowItem {
			for di, d := range item {
				colData.items[ci][ri][di] = d
			}
		}
	}
	return colData
}

//getBox return box data
func getBox(in, ri, ci int) (index, startRi, endRi, startCi, endCi int) {

	switch {
	case (ri >= 0 && ri <= 2 && ci >= 0 && ci <= 2) || in == 1:
		index = 1
		startRi, endRi, startCi, endCi = 0, 2, 0, 2
	case (ri >= 0 && ri <= 2 && ci >= 3 && ci <= 5) || in == 2:
		index = 2
		startRi, endRi, startCi, endCi = 0, 2, 3, 5
	case (ri >= 0 && ri <= 2 && ci >= 6 && ci <= 8) || in == 3:
		index = 3
		startRi, endRi, startCi, endCi = 0, 2, 6, 8
	case (ri >= 3 && ri <= 5 && ci >= 0 && ci <= 2) || in == 4:
		index = 4
		startRi, endRi, startCi, endCi = 3, 5, 0, 2
	case (ri >= 3 && ri <= 5 && ci >= 3 && ci <= 5) || in == 5:
		index = 5
		startRi, endRi, startCi, endCi = 3, 5, 3, 5
	case (ri >= 3 && ri <= 5 && ci >= 6 && ci <= 8) || in == 6:
		index = 6
		startRi, endRi, startCi, endCi = 3, 5, 6, 8
	case (ri >= 6 && ri <= 8 && ci >= 0 && ci <= 2) || in == 7:
		index = 7
		startRi, endRi, startCi, endCi = 6, 8, 0, 2
	case (ri >= 6 && ri <= 8 && ci >= 3 && ci <= 5) || in == 8:
		index = 8
		startRi, endRi, startCi, endCi = 6, 8, 3, 5
	case (ri >= 6 && ri <= 8 && ci >= 6 && ci <= 8) || in == 9:
		index = 9
		startRi, endRi, startCi, endCi = 6, 8, 6, 8
	default:
		fmt.Println("error")
	}

	return
}

//showInfo show Info. type is 'row or col or box or all'
func showInfo(it ItemData, showType string, sele int) {
	switch showType {
	case "row":
		//row item print
		tmp := getRow(it)
		for ri, _ := range tmp.items {
			if sele == 100 {
				for ci := 0; ci < 9; ci++ {
					fmt.Printf("colNo: %v rowNo: %v item: %v\n", ri+1, ci+1, tmp.items[ri][ci])
				}
			} else if sele >= 0 && sele <= 10 {
				var tmpRow [9]int
				for ci := 0; ci < 9; ci++ {
					tmpRow[ci] = tmp.items[ri][ci][sele]
				}
				fmt.Printf("rowNo: %v item: %v\n", ri+1, tmpRow)
			}
		}
	case "col":
		//col item print
		tmp := getCol(it)
		for ci, _ := range tmp.items {
			if sele == 100 {
				for ri := 0; ri < 9; ri++ {
					fmt.Printf("colNo: %v rowNo: %v item: %v\n", ci+1, ri+1, tmp.items[ci][ri])
				}
			} else if sele >= 0 && sele <= 10 {
				var tmpCol [9]int
				for ri := 0; ri < 9; ri++ {
					tmpCol[ri] = tmp.items[ci][ri][sele]
				}
				fmt.Printf("colNo: %v item: %v\n", ci+1, tmpCol)
			}
		}

	case "box":
		//box item print
		if sele == 100 {
			for i := 1; i <= 9; i++ {
				_, startRi, endRi, startCi, endCi := getBox(i, 100, 100)

				for rIndex := startRi; rIndex <= endRi; rIndex++ {
					for cIndex := startCi; cIndex <= endCi; cIndex++ {
						fmt.Printf("showBox: index:%v value:%v \n", i, it.items[rIndex][cIndex])
					}
				}
			}

		} else if sele >= 0 && sele <= 10 {
			for i := 1; i <= 9; i++ {
				fmt.Printf("showInfo-box: box-index: %v ", i)
				_, startRi, endRi, startCi, endCi := getBox(i, 100, 100)

				for rIndex := startRi; rIndex <= endRi; rIndex++ {
					for cIndex := startCi; cIndex <= endCi; cIndex++ {
						fmt.Printf("showBox: index:%v namber:%v value:%v \n", i, sele, it.items[rIndex][cIndex][sele])
					}
				}
			}

		}
	case "all":
		//all item print
		for ri, _ := range it.items {
			for ci := 0; ci < 9; ci++ {
				fmt.Printf("rowNo: %v colNo: %v item: %v\n", ri+1, ci+1, it.items[ri][ci])
			}
		}
	case "all-one":
		//all item print
		var tmp [9]int
		for ri, _ := range it.items {
			for ci := 0; ci < 9; ci++ {
				tmp[ci] = it.items[ri][ci][0]
			}
			fmt.Println(spacePlus(tmp))
		}
	default:
		fmt.Println("nothing")
	}
}

//spacePlus is add space
func spacePlus(items [9]int) (str string) {

	for _, item := range items {
		s := strconv.Itoa(item)
		if item < 10 {
			str += s + " "
		} else {
			str += "-" + " "
		}
	}
	return
}

//check2 is check and set
func check2(it ItemData) ItemData {
	it = checkRow2(it)
	it = checkCol2(it)
	checkBox2(it)

	return it
}

func main() {

	//items print
	for _, item := range items {
		fmt.Printf("%v\n", item)
	}

	//itemInfo set
	for ri, rowItem := range items {
		for ci, item := range rowItem {
			if item > 0 {
				itemInfo[ri][ci][0] = item
				for i := 1; i <= 9; i++ {
					itemInfo[ri][ci][i] = 11
				}
			} else {
				itemInfo[ri][ci][0] = 10
				for i := 1; i <= 9; i++ {
					itemInfo[ri][ci][i] = 10
				}
			}
		}
	}
	fmt.Println("====iteminfo====")
	fmt.Println(itemInfo)
	fmt.Println("====iteminfo====")

	var it ItemData

	//itemInfo set
	for ri, rowItem := range items {
		for ci, item := range rowItem {
			if item > 0 {
				it.items[ri][ci][0] = item
				for i := 1; i <= 9; i++ {
					it.items[ri][ci][i] = 11
				}
			} else {
				it.items[ri][ci][0] = 10
				for i := 1; i <= 9; i++ {
					it.items[ri][ci][i] = 10
				}
			}
		}
	}

	fmt.Println("====struct====")
	fmt.Println(it.items)
	fmt.Println("====struct====")

	//itemInfo set
	it = set11(it)

	for i := 0; i < 10; i++ {
		//itemInfo check2 & set
		it = check2(it)
	}

	//show all
	fmt.Println("show all:")
	showInfo(it, "all-one", 100)
}
