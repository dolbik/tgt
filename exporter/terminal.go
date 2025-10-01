package exporter

import (
	"fmt"
	"importer/entity"
)

type TerminalExporter struct {
}

func NewTerminalExporter() Exporter {
	return &TerminalExporter{}
}

func (ex TerminalExporter) ExportData(data []entity.DomainData) error {
	if data == nil {
		return fmt.Errorf("error provided data is empty (nil)")
	}

	fmt.Println("domain,number_of_customers")
	for _, v := range data {
		fmt.Printf("%s,%v\n", v.Domain, v.CustomerQuantity)
	}

	return nil
}
