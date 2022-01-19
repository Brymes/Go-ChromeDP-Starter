package scraper

import (
	"scraper/utils"
	"time"

	"github.com/chromedp/chromedp"
)

func Scraper() {
	// Pass in the amount of seconds for Chrome Context to lie
	ctxt, _ := utils.GenerateChromeContext(180)

	url := ""

	task1 := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(time.Second * 30),
		// Do something else
	}

	chromedp.Run(ctxt, task1)
}
