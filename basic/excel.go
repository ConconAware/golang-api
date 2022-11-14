package test

import (
	"bytes"
	"fmt"
	"io"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func Excel(c *gin.Context) {
	startTime := time.Now()
	fmt.Println("8888888888888888")
	uploadExcel(c)
	// createExcel()
	// readExcel()
	// importExcel()
	endTime := time.Since(startTime)
	fmt.Printf("\n excel time : %g", endTime.Seconds())
	fmt.Println("\n end uplaod excel")
}

func uploadExcel(c *gin.Context) {
	// form, _ := c.MultipartForm()
	// files := form.File["file"]

	// for _, file := range files {
	// 	fmt.Println(file.Filename)

	// }

	file, _, err := c.Request.FormFile("file")
	defer file.Close()
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return
	}

	f, err := excelize.OpenReader(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	f1 := excelize.NewFile()
	index := f1.NewSheet("Sheet2")
	for i, row := range rows {

		for j, colCell := range row {
			// if i == len(row) {
			// 	fmt.Println(colCell, "\t")
			if j == 0 {
				f1.SetCellValue("Sheet2", "A"+strconv.Itoa(i), colCell)
			} else {
				f1.SetCellValue("Sheet2", "B"+strconv.Itoa(i), colCell)
			}

			// }

		}
		// fmt.Println()
	}
	f1.SetActiveSheet(index)
	if err := f1.SaveAs("Book3.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func createExcel() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func readExcel() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	//ex.1
	// Get value from cell by given worksheet name and cell reference.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)

	//ex.2
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func importExcel() {
	for i := 0; i < 103589; i++ {
		for j := 0; j < 1; j++ {
			if i == 103588 {
				fmt.Println("===============")
			}
		}
	}
}
