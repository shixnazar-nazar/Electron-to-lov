package main

import (
	"github.com/khasanovbi/banks_db/gen"
	"log"
	"sort"
)

const (
	BaseDir       = "banks_db"
	CountriesPath = BaseDir + "/countries.go"
	BanksPath     = BaseDir + "/banks.go"
)

func calculateBanksCount(countryBanksSlice []gen.CountryBanks) int {
	sum := 0
	for _, countryBanks := range countryBanksSlice {
		sum += len(countryBanks.Banks)
	}
	return sum
}

func getCountries(banksByCountry []gen.CountryBanks) []string {
	countries := make([]string, 0, len(banksByCountry))
	for _, countryBanks := range banksByCountry {
		countries = append(countries, countryBanks.Country)
	}
	sort.Strings(countries)
	return countries
}

func main() {
	countryBanksSlice := gen.ParseBanks()
	log.Printf(
		"the banks are parsed: countriesCount=%d, banksCount=%d",
		len(countryBanksSlice),
		calculateBanksCount(countryBanksSlice),
	)
	gen.GenerateCountriesFile(CountriesPath, getCountries(countryBanksSlice))
	gen.GenerateBanksFile(BanksPath, countryBanksSlice)
}
