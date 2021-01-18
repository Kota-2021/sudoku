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

//checkRow is searching same number in the Row
func checkRow(ri, ci, item int) {
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

//checkCol is searching same number in the col
func checkCol(ri, ci, item int) {
	if item > 0 && item <= 9 {
		for i := 0; i < 9; i++ {
			itemInfo[i][ci][item] = 11
		}
	} else if item == 10 {
		for _, rowItem := range itemInfo[ci] {
			if rowItem[0] >= 1 && rowItem[0] <= 9 {
				itemInfo[ci][ri][rowItem[0]] = 11
			}
		}
	}
}

//checkBox is searching same number in the box
func checkBox(ri, ci, item int) {

	_, startRi, endRi, startCi, endCi := getBox(0, ri, ci)

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
func showInfo(showType string, sele int) {
	switch showType {
	case "row":
		//row item print
		tmp := getRow()
		for ri, _ := range tmp {
			if sele == 100 {
				for ci := 0; ci < 9; ci++ {
					fmt.Printf("colNo: %v rowNo: %v item: %v\n", ri+1, ci+1, tmp[ri][ci])
				}
			} else if sele >= 0 && sele <= 10 {
				var tmpRow [9]int
				for ci := 0; ci < 9; ci++ {
					tmpRow[ci] = tmp[ri][ci][sele]
				}
				fmt.Printf("rowNo: %v item: %v\n", ri+1, tmpRow)
			}
		}
	case "col":
		//col item print
		tmp := getCol()
		for ci, _ := range tmp {
			if sele == 100 {
				for ri := 0; ri < 9; ri++ {
					fmt.Printf("colNo: %v rowNo: %v item: %v\n", ci+1, ri+1, tmp[ci][ri])
				}
			} else if sele >= 0 && sele <= 10 {
				var tmpCol [9]int
				for ri := 0; ri < 9; ri++ {
					tmpCol[ri] = tmp[ci][ri][sele]
				}
				fmt.Printf("colNo: %v item: %v\n", ci+1, tmpCol)
			}
		}

	case "box":
		//box item print

		for _, boxItem := range boxInfo {
			fmt.Printf("%v\n", boxItem)
		}

		if sele == 100 {
			for i := 1; i <= 9; i++ {
				_, startRi, endRi, startCi, endCi := getBox(i, 100, 100)

				for rIndex := startRi; rIndex <= endRi; rIndex++ {
					for cIndex := startCi; cIndex <= endCi; cIndex++ {
						fmt.Printf("showBox: index:%v value:%v \n", i, itemInfo[rIndex][cIndex])
					}
				}
			}

		} else if sele >= 0 && sele <= 10 {
			for i := 1; i <= 9; i++ {
				_, startRi, endRi, startCi, endCi := getBox(i, 100, 100)

				for rIndex := startRi; rIndex <= endRi; rIndex++ {
					for cIndex := startCi; cIndex <= endCi; cIndex++ {
						fmt.Printf("showBox: index:%v namber:%v value:%v \n", i, sele, itemInfo[rIndex][cIndex][sele])
					}
				}
			}

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

	//show all
	fmt.Println("show all-1:")
	showInfo("all", 100)

	//itemInfo check & set
	for ri, rowItem := range itemInfo {
		for ci, item := range rowItem {

			//row check
			checkRow(ri, ci, item[0])

			//col check
			checkCol(ri, ci, item[0])

			//box check
			checkBox(ri, ci, item[0])
		}
	}

	//show all
	fmt.Println("show all-2:")
	showInfo("all", 100)

	//show col
	fmt.Println("showCol: ")
	showInfo("col", 1)

	//show row
	fmt.Println("showRow: ")
	showInfo("row", 1)

	//show box
	fmt.Println("showBox: ")
	showInfo("box", 1)
}
