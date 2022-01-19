package utils

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

var devtoolsWsURL = flag.String("devtoolsWSurl", "http://127.0.0.1:9222/", "DevTools Websocket URL")

func GenerateChromeContext(timed int64) (ctxt context.Context, cancelFunc context.CancelFunc) {
	var allocatorContext context.Context
	allocatorContext, _ =
		chromedp.NewRemoteAllocator(context.Background(),
			*devtoolsWsURL)

	ctx, _ := chromedp.NewContext(allocatorContext)

	dur, err := time.ParseDuration(fmt.Sprintf("%vs", timed))
	if err != nil {
		log.Fatalf("Error Generating Chrome Context ::: %v", err)
	}

	ctxt, cancelFunc = context.WithTimeout(ctx, time.Duration(dur))

	return
}
