package exporter

import (
	"importer/entity"
)

type Exporter interface {
	ExportData(data []entity.DomainData) error
}

func GetExporter(outputPath string) Exporter {
	if outputPath == "" {
		return NewTerminalExporter()
	}
	return NewCsvExporter(outputPath)
}
