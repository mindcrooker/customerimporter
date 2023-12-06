## Customer Importer

### Task description:

Package customerimporter reads from the given customers.csv file and returns a
sorted (data structure of your choice) of email domains along with the number
 of customers with e-mail addresses for each domain.  Any errors should be
 logged (or handled). Performance matters (this is only ~3k lines, but *could*
 be 1m lines or run on a small machine).


### Use
To use the package locally, download it, import in your project, and add a line to your `go.mod` file that will link to directory with the package, e.g.:
```
replace github.com/mindcrooker/customerimporter => <path to the downloaded package>

```
After importing package into the project, simply use `HandleCustomers(path string)` function and pass in the path to `.csv` file with customer data. Function returns a slice of Domains (ordered by domain name):
```
type Domain struct {
	Address string
	Count   int
}

[]Domain{
	{Address: "124.org", Count: 2},
	{Address: "abcd.com", Count: 5},
	{Address: "gmail.com", Count: 3},
}

```
### Performance
Provided 'customers.csv' file, which has 3k lines, was processed in ~6ms on my machine. For testing purposes I also generated files with 1m and 2m lines and processing them took <2s and <4s respectively. If better performance is required and multi-core machine is available, this time could be brought down by loading rows in batches and processing them in parallel. 
