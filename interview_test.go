package customerimporter

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortDomain(t *testing.T) {
	domains := map[string]int{
		"gmail.com": 3,
		"abcd.com":  5,
		"124.org":   2,
	}

	expectedSortedDomains := []Domain{
		{Address: "124.org", Count: 2},
		{Address: "abcd.com", Count: 5},
		{Address: "gmail.com", Count: 3},
	}

	assert.Len(t, sortDomains(domains, len(domains)), 3)
	assert.Equal(t, expectedSortedDomains, sortDomains(domains, len(domains)))
}

func TestSortDomainEmpty(t *testing.T) {
	domains := make(map[string]int)

	expectedSortedDomains := []Domain{}

	assert.Len(t, sortDomains(domains, len(domains)), 0)
	assert.Equal(t, expectedSortedDomains, sortDomains(domains, len(domains)))

}

func TestGetDomain(t *testing.T) {
	type testCase struct {
		email          string
		expectedDomain string
	}

	tests := []testCase{
		{email: "email1@124.org", expectedDomain: "124.org"},
		{email: "email2@abcd.com", expectedDomain: "abcd.com"},
		{email: "email3@gmail.com", expectedDomain: "gmail.com"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedDomain, getDomain(test.email))
	}

}

func TestHandleCustomers(t *testing.T) {
	filePath := "testdata/valid_customers.csv"
	CSVData := [][]string{
		{"first_name", "last_name", "email", "gender", "ip_address"},
		{"Mildred", "Hernandez", "mhernandez0@github.io", "Female", "38.194.51.128"},
		{"Bonnie", "Ortiz", "bortiz1@cyberchimps.com", "Female", "197.54.209.129"},
		{"Dennis", "Henry", "dhenry2@hubpages.com", "Male", "155.75.186.217"},
		{"Justin", "Hansen", "jhansen3@cyberchimps.com", "Male", "251.166.224.119"},
	}
	createTestCSV(filePath, CSVData)

	defer func() {
		if err := os.Remove(filePath); err != nil {
			t.Errorf("Error removing test file: %s", err)
		}
	}()

	result := HandleCustomers(filePath)

	expectedDomains := []Domain{
		{Address: "cyberchimps.com", Count: 2},
		{Address: "github.io", Count: 1},
		{Address: "hubpages.com", Count: 1},
	}

	assert.Equal(t, expectedDomains, result)
}

func TestHandleCustomersInvalidEmails(t *testing.T) {
	filePath := "testdata/invalid_customers.csv"
	CSVData := [][]string{
		{"first_name", "last_name", "email", "gender", "ip_address"},
		{"Mildred", "Hernandez", "github.io", "Female", "38.194.51.128"},
		{"Bonnie", "Ortiz", "", "Female", "197.54.209.129"},
		{"Dennis", "Henry", "email", "Male", "155.75.186.217"},
		{"Justin", "Hansen", "jhansen3@cyberchimps.com", "Male", "251.166.224.119"},
	}
	createTestCSV(filePath, CSVData)

	defer func() {
		if err := os.Remove(filePath); err != nil {
			t.Errorf("Error removing test file: %s", err)
		}
	}()

	result := HandleCustomers(filePath)

	expectedDomains := []Domain{
		{Address: "cyberchimps.com", Count: 1},
	}

	assert.Equal(t, expectedDomains, result)
}

func TestHandleCustomersEmpty(t *testing.T) {
	filePath := "testdata/empty.csv"
	CSVData := [][]string{{}}
	createTestCSV(filePath, CSVData)

	defer func() {
		if err := os.Remove(filePath); err != nil {
			t.Errorf("Error removing test file: %s", err)
		}
	}()

	result := HandleCustomers(filePath)

	expectedDomains := []Domain{}

	assert.Equal(t, expectedDomains, result)
}

func createTestCSV(filePath string, data [][]string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(data)
	if err != nil {
		panic(err)
	}
}
