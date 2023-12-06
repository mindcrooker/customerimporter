package customerimporter

import (
	"encoding/csv"
	"log"
	"net/mail"
	"os"
	"sort"
	"strings"
	"time"
)

type Domain struct {
	Address string
	Count   int
}

func timeTrack(start time.Time, name string, path string) {
	elapsed := time.Since(start)
	log.Printf("%s for file '%s' took: %s", name, path, elapsed)
}

func HandleCustomers(path string) []Domain {
	if path == "" {
		path = "customers.csv"
	}

	// Measure execution time
	defer timeTrack(time.Now(), "HandleCustomers", path)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully opened the CSV file:", f.Name())
	defer f.Close()

	fileReader := csv.NewReader(f)
	records, err := fileReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	totalRows := len(records)
	log.Println("Number of rows: ", totalRows)

	domains := make(map[string]int)

	for i, r := range records {
		email, err := mail.ParseAddress(r[2])
		if err != nil {
			log.Printf("Incorrect email address in row %d: '%s'", i, r[2])
			continue
		}

		domain := getDomain(email.Address)
		domains[domain] += 1
	}
	uniqueDomains := len(domains)
	log.Println("Total unique domains:", uniqueDomains)

	sortedDomains := sortDomains(domains, uniqueDomains)

	return sortedDomains
}

func sortDomains(domains map[string]int, uniqueDomains int) []Domain {
	sortedDomains := make([]Domain, 0, uniqueDomains)

	for address, count := range domains {
		sortedDomains = append(sortedDomains, Domain{Address: address, Count: count})
	}

	sort.SliceStable(sortedDomains, func(i int, j int) bool {
		return sortedDomains[i].Address < sortedDomains[j].Address
	})

	return sortedDomains
}

func getDomain(email string) string {
	components := strings.Split(email, "@")
	_, domain := components[0], components[1]

	return domain
}
