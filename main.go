package main

import "fmt"

//init default numbers
var items = [9][9]int{
	[9]int{0, 5, 0, 0, 0, 0, 0, 6, 0},
	[9]int{8, 0, 0, 4, 0, 5, 0, 0, 1},
	[9]int{0, 0, 2, 1, 0, 7, 4, 0, 0},
	[9]int{0, 6, 3, 0, 0, 0, 9, 2, 0},
	[9]int{0, 0, 0, 0, 1, 0, 0, 0, 0},
	[9]int{0, 9, 7, 0, 0, 0, 5, 1, 0},
	[9]int{0, 0, 8, 5, 0, 6, 2, 0, 0},
	[9]int{7, 0, 0, 8, 0, 4, 0, 0, 6},
	[9]int{0, 4, 0, 0, 0, 0, 0, 7, 0},
}

//rowInfo is row info
var rowInfo [9][9]int

//colInfo is col info
var colInfo [9][9]int

//boxInfo is box info
var boxInfo [9][9]int

//itemInfo is item info
var itemInfo [9][9][10]int

//getBoxNum return item's box number
func getBoxNum(ri, ci int) (index int) {
	switch {
	case ri >= 0 && ri <= 2 && ci >= 0 && ci <= 2:
		index = 0
	case ri >= 0 && ri <= 2 && ci >= 3 && ci <= 5:
		index = 1
	case ri >= 0 && ri <= 2 && ci >= 6 && ci <= 8:
		index = 2
	case ri >= 3 && ri <= 5 && ci >= 0 && ci <= 2:
		index = 3
	case ri >= 3 && ri <= 5 && ci >= 3 && ci <= 5:
		index = 4
	case ri >= 3 && ri <= 5 && ci >= 6 && ci <= 8:
		index = 5
	case ri >= 6 && ri <= 8 && ci >= 0 && ci <= 2:
		index = 6
	case ri >= 6 && ri <= 8 && ci >= 3 && ci <= 5:
		index = 7
	case ri >= 6 && ri <= 8 && ci >= 6 && ci <= 8:
		index = 8
	default:
		fmt.Println("error")
	}
	return
}

//searchRow is searching same number in the Row
func searchRow(ri, ci, item int) {
	if item > 0 && item <= 9 {
		for i := 0; i < 9; i++ {
			itemInfo[ri][i][item] = 11
		}
	} else if item == 10 {
		for _, colItem := range itemInfo[ri] {
			if colItem[0] >= 1 && colItem[0] <= 9 {
				itemInfo[ri][ci][colItem[0]] = 11
			}
		}
	}
}

//searchCol is searching same number in the col
func searchCol(ri, ci, item int) {
	if item > 0 && item <= 9 {
		for i := 0; i < 9; i++ {
			itemInfo[ci][i][item] = 11
		}
	} else if item == 10 {
		for _, rowItem := range itemInfo[ci] {
			if rowItem[0] >= 1 && rowItem[0] <= 9 {
				itemInfo[ci][ri][rowItem[0]] = 11
			}
		}
	}
}

//searchBox is searching same number in the box
func searchBox(ri, ci, item int) {
	bIndex := getBoxNum(ri, ci)
	startRi, endRi, startCi, endCi := getBox(bIndex)

	if item >= 1 && item <= 9 {
		for rIndex := startRi; rIndex <= endRi; rIndex++ {
			for cIndex := startCi; cIndex <= endCi; cIndex++ {
				itemInfo[rIndex][cIndex][item] = 11
			}
		}
	} else if item == 10 {
		for rIndex := startRi; rIndex <= endRi; rIndex++ {
			for cIndex := startCi; cIndex <= endCi; cIndex++ {
				if itemInfo[rIndex][cIndex][0] > 0 && itemInfo[rIndex][cIndex][0] <= 9 {
					itemInfo[ri][ci][itemInfo[rIndex][cIndex][0]] = 11
				}
			}
		}
	}
}

//searchInfoRow is searching same number in the itemInfo's row data
/* func searchInfoRow(ri, ci, item int) {
	var tmp [3]bool
	//9 times search No.1 - No.9
	for index := 1; index <= 9; index++ {
		//3 times search
		for j:=0;j<3;j++{
			//row0 row1 ro2 cheack
			for _,val:=range itemInfo[j]{
				//
				if val == index {
					tmp[0] = true
				}
			}
		}

	}
} */

//getRow return row data
func getRow() [9][9][10]int {
	var rowData [9][9][10]int
	for ri, rowItem := range itemInfo {
		for ci, item := range rowItem {
			for di, d := range item {
				rowData[ri][ci][di] = d
			}
		}
	}
	return rowData
}

//getCol return coll data
func getCol() [9][9][10]int {
	var colData [9][9][10]int
	for ri, rowItem := range itemInfo {
		for ci, item := range rowItem {
			for di, d := range item {
				colData[ci][ri][di] = d
			}
		}
	}
	return colData
}

//getBox return box data
func getBox(index int) (startRi, endRi, startCi, endCi int) {

	switch index {
	case 0:
		startRi, endRi, startCi, endCi = 0, 2, 0, 2
		//startRi=0 , endRi=2 , startCi=0 , endCi=2
	case 1:
		startRi, endRi, startCi, endCi = 0, 2, 3, 5
		//startRi=0 , endRi=2 , startCi=3 , endCi=5
	case 2:
		startRi, endRi, startCi, endCi = 0, 2, 6, 8
		//startRi=0 , endRi=2 , startCi=6 , endCi=8
	case 3:
		startRi, endRi, startCi, endCi = 3, 5, 0, 2
		//startRi=3 , endRi=5 , startCi=0 , endCi=2
	case 4:
		startRi, endRi, startCi, endCi = 3, 5, 3, 5
		//startRi=3 , endRi=5 , startCi=3 , endCi=5
	case 5:
		startRi, endRi, startCi, endCi = 3, 5, 6, 8
		//startRi=3 , endRi=5 , startCi=6 , endCi=8
	case 6:
		startRi, endRi, startCi, endCi = 6, 8, 0, 2
		//startRi=6 , endRi=8 , startCi=0 , endCi=2
	case 7:
		startRi, endRi, startCi, endCi = 6, 8, 3, 5
		//startRi=6 , endRi=8 , startCi=3 , endCi=5
	case 8:
		startRi, endRi, startCi, endCi = 6, 8, 6, 8
		//startRi=6 , endRi=8 , startCi=6 , endCi=8
	default:
		fmt.Println("error")
	}

	return
}

//showInfo show Info. type is 'row or col or box or all'
func showInfo(showType string) {
	switch showType {
	case "row":
		//row item print
		for _, rowItem := range rowInfo {
			fmt.Printf("%v\n", rowItem)
		}
	case "col":
		//col item print
		/* for _, colItem := range colInfo {
			fmt.Printf("%v\n", colItem)
		} */
		//all item print
		tmp := getCol()
		for ci, _ := range tmp {
			for ri := 0; ri < 9; ri++ {
				fmt.Printf("colNo: %v rowNo: %v item: %v\n", ci+1, ri+1, tmp[ci][ri])
			}
		}

	case "box":
		//box item print
		for _, boxItem := range boxInfo {
			fmt.Printf("%v\n", boxItem)
		}
	case "all":
		//all item print
		for ri, _ := range itemInfo {
			for ci := 0; ci < 9; ci++ {
				fmt.Printf("rowNo: %v colNo: %v item: %v\n", ri+1, ci+1, itemInfo[ri][ci])
			}
		}
	default:
		fmt.Println("nothing")
	}
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

	//show
	fmt.Println("show all-1:")
	showInfo("all")

	//itemInfo check & set
	for ri, rowItem := range itemInfo {
		for ci, item := range rowItem {
			//row check
			searchRow(ri, ci, item[0])

			//col check
			searchCol(ri, ci, item[0])

			//box check
			searchBox(ri, ci, item[0])
		}
	}

	//show
	fmt.Println("show all-2:")
	showInfo("all")

	/* //row item set
	for ri, rowItem := range itemInfo {
		for ci, item := range rowItem {
			rowInfo[ri][ci] = item
		}
	}

	//col item set
	for ri, rowItem := range itemInfo {
		for ci, item := range rowItem {
			colInfo[ci][ri] = item
		}
	}

	//box item set
	for ri, rowItem := range itemInfo {
		for ci, item := range rowItem {
			if item > 0 {
				boxInfo[getBoxNum(ri, ci)][item-1] = item
			}
		}
	} */

	fmt.Println("showCol: ")
	showInfo("col")
}
