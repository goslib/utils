package utlXlsx

import (
	"bytes"

	"github.com/goslib/utils/utlReflect"
	"github.com/tealeg/xlsx"
)

type IHeadersFilter interface {
	IsEnabled() bool
	IsColumnAllowed(ith int) bool
}

func NewXlsxFileWithDataset(sheetName string, dataset [][]string, filter IHeadersFilter) (*xlsx.File, error) {
	file := xlsx.NewFile()

	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		return nil, err
	}

	FillSheetWithRowsOfCells(sheet, dataset, filter)

	return file, nil
}

func NewXlsxFileBytesWithDataset(sheetName string, dataset [][]string, filter IHeadersFilter) ([]byte, error) {
	file, err := NewXlsxFileWithDataset(sheetName, dataset, filter)

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = file.Write(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func WriteXlsxFileWithDataset(xlsxFileLocation, sheetName string, dataset [][]string, filter IHeadersFilter) error {
	file, err := NewXlsxFileWithDataset(sheetName, dataset, filter)

	if err != nil {
		return err
	}

	err = file.Save(xlsxFileLocation)
	if err != nil {
		return err
	}

	return nil
}

func FillSheetWithRowsOfCells(sheet *xlsx.Sheet, dataset [][]string, filter IHeadersFilter) {

	if utlReflect.IsInterfaceValueNil(filter) {

		for _, fields := range dataset {
			row := sheet.AddRow()

			for i := 0; i < len(fields); i++ {
				cell := row.AddCell()
				cell.Value = fields[i]
			}
		}

	} else {

		for _, fields := range dataset {
			row := sheet.AddRow()

			for i := 0; i < len(fields); i++ {
				if !filter.IsColumnAllowed(i) {
					continue
				}
				cell := row.AddCell()
				cell.Value = fields[i]
			}
		}

	}
}
