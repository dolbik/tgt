// Package customerimporter reads from a CSV file and returns a sorted (data
// structure of your choice) of email domains along with the number of customers
// with e-mail addresses for each domain. This should be able to be ran from the
// CLI and output the sorted domains to the terminal or to a file. Any errors
// should be logged (or handled). Performance matters (this is only ~3k lines,
// but could be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"importer/entity"
)

type CustomerImporter struct {
	path *string
}

// NewCustomerImporter returns a new CustomerImporter that reads from file at specified path.
func NewCustomerImporter(filePath *string) *CustomerImporter {
	return &CustomerImporter{
		path: filePath,
	}
}

// ImportDomainData reads and returns sorted customer domain data from CSV file.
func (ci CustomerImporter) ImportDomainData() ([]entity.DomainData, error) {
	file, err := os.Open(*ci.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.ReuseRecord = true

	data := make(map[string]uint64)

	_, err = csvReader.Read()
	if errors.Is(err, io.EOF) {
		return []entity.DomainData{}, nil
	}

	if err != nil {
		return nil, err
	}

	for {
		line, readErr := csvReader.Read()
		if errors.Is(readErr, io.EOF) {
			break
		}
		if readErr != nil {
			return nil, readErr
		}

		if len(line) < 3 {
			return nil, fmt.Errorf("error invalid csv line: %v", line)
		}

		domain, err := ci.getDomainFromEmail(line[2])
		if err != nil {
			return nil, err
		}

		data[domain] += 1
	}

	return ci.sort(data), nil
}

func (ci CustomerImporter) getDomainFromEmail(email string) (string, error) {
	email, domain, found := strings.Cut(email, "@")
	if email == "" || !found {
		return "", fmt.Errorf("error invalid email address: %s", email)
	}
	return domain, nil
}

func (ci CustomerImporter) sort(data map[string]uint64) []entity.DomainData {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	result := make([]entity.DomainData, 0, len(keys))
	for _, k := range keys {
		result = append(result, entity.DomainData{
			Domain:           k,
			CustomerQuantity: data[k],
		})
	}

	return result
}
