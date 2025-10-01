package exporter

import (
	"encoding/csv"
	"fmt"
	"importer/entity"
	"io"
	"os"
	"strconv"
)

type CsvExporter struct {
	outputPath string
}

// NewCsvExporter returns a new CsvExporter that writes customer domain data to specified file.
func NewCsvExporter(outputPath string) Exporter {
	return &CsvExporter{
		outputPath: outputPath,
	}
}

func (ex CsvExporter) ExportData(data []entity.DomainData) error {
	if data == nil {
		return fmt.Errorf("error provided data is empty (nil)")
	}
	outputFile, err := os.Create(ex.outputPath)
	if err != nil {
		return fmt.Errorf("error creating new file for saving: %v", err)
	}
	defer outputFile.Close()
	return exportCsv(data, outputFile)
}

func exportCsv(data []entity.DomainData, output io.Writer) error {
	headers := []string{"domain", "number_of_customers"}
	csvWriter := csv.NewWriter(output)
	defer csvWriter.Flush()

	if err := csvWriter.Write(headers); err != nil {
		return err
	}
	pair := make([]string, 2)
	for _, v := range data {
		pair[0] = v.Domain
		pair[1] = strconv.FormatUint(v.CustomerQuantity, 10)
		if err := csvWriter.Write(pair); err != nil {
			return err
		}
	}
	return csvWriter.Error()
}
