package constants

import (
	"os"
)

func GetCountriesURL() string {
	envUrl := os.Getenv("COUNTRY_URL")
	if envUrl != "" {
		return envUrl
	}
	return "http://services.groupkt.com/country/get/"

}
