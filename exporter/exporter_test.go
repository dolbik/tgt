package exporter

import (
	"fmt"
	"importer/customerimporter"
	"importer/entity"
	"testing"
)

func TestExportData(t *testing.T) {
	path := "./test_output.csv"
	data := []entity.DomainData{
		{
			Domain:           "livejournal.com",
			CustomerQuantity: 12,
		},
		{
			Domain:           "microsoft.com",
			CustomerQuantity: 22,
		},
		{
			Domain:           "newsvine.com",
			CustomerQuantity: 15,
		},
		{
			Domain:           "pinteres.uk",
			CustomerQuantity: 10,
		},
		{
			Domain:           "yandex.ru",
			CustomerQuantity: 43,
		},
	}
	exporter := NewCsvExporter(path)

	err := exporter.ExportData(data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestExportInvalidPath(t *testing.T) {
	path := ""
	exporter := NewCsvExporter(path)

	err := exporter.ExportData([]entity.DomainData{})
	if err == nil {
		t.Fatal(err)
	}
	t.Log(err)
}

func TestExportEmptyData(t *testing.T) {
	path := "./test_output.csv"
	exporter := NewCsvExporter(path)

	err := exporter.ExportData(nil)
	if err == nil {
		t.Fatal(err)
	}
	t.Log(err)
}

func BenchmarkImportDomainData(b *testing.B) {
	b.StopTimer()
	dir := b.TempDir()
	path := fmt.Sprintf("%s/test_output.csv", dir)
	dataPath := "../customerimporter/benchmark10k.csv"
	importer := customerimporter.NewCustomerImporter(&dataPath)
	data, err := importer.ImportDomainData()
	if err != nil {
		b.Error(err)
	}
	exporter := NewCsvExporter(path)

	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := exporter.ExportData(data); err != nil {
			b.Fatal(err)
		}
	}
}
