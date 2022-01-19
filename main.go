package main

import (
	"scraper/scraper"
	"scraper/utils"
)

func main() {
	//Launch Chrome
	utils.LaunchChrome()

	// Launch Main Scraper entry point function
	scraper.Scraper()
}
