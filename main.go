package main

import (
	"flag"
	"importer/customerimporter"
	"importer/exporter"
	"log/slog"
)

type Options struct {
	path    *string
	outFile *string
}

func readOptions() *Options {
	opts := &Options{}
	opts.path = flag.String("path", "./customers.csv", "Path to the file with customer data")
	opts.outFile = flag.String("out", "", "Optional: output file path. If empty program will output results to the terminal")
	flag.Parse()
	return opts
}

func main() {
	opts := readOptions()
	importer := customerimporter.NewCustomerImporter(opts.path)
	data, err := importer.ImportDomainData()
	if err != nil {
		slog.Error("error importing customer data: ", err)
		return
	}
	exp := exporter.GetExporter(*opts.outFile)
	if saveErr := exp.ExportData(data); saveErr != nil {
		slog.Error("error saving domain data: ", saveErr)
		return
	}
}
