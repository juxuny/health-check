package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/juxuny/health-check/log"
)

var (
	url              string
	timeoutInSeconds int
	verbose          bool
)

func storeArgs(k string, v string) {
	switch k {
	case "timeout", "t":
		parsedValue, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("invalid timeout in seconds: %s", v))
		}
		timeoutInSeconds = int(parsedValue)
	case "v", "verbose":
		verbose = v == "1" || strings.ToLower(v) == "true"
	}
}

func parseArgs() {
	i := 1
	for i < len(os.Args) {
		k := os.Args[i]
		if strings.Index(k, "-") == 0 {
			if i+1 >= len(os.Args) {
				log.Fatal(fmt.Sprintf("missing argument value for option '%s'", k))
			}
			k = strings.Trim(k, "-")
			v := os.Args[i+1]
			storeArgs(k, v)
			i += 2
			continue
		} else {
			url = k
		}
		i++
	}
	if timeoutInSeconds == 0 {
		timeoutInSeconds = 30
	}
}

func request(ctx context.Context, url string) (code int, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer func () {
		_ = resp.Body.Close()
	} ()
	_, _ = io.ReadAll(resp.Body)
	return resp.StatusCode, nil
}

func main() {
	parseArgs()
	ctx := context.Background()
	var cancel context.CancelFunc = func() {

	}
	if timeoutInSeconds > 0 {
		ctx, cancel = context.WithTimeout(ctx, time.Duration(timeoutInSeconds)*time.Second)
	}
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			os.Exit(255)
		default:
		}
		code, err := request(ctx, url)
		if err != nil {
			log.Error(err)
			time.Sleep(time.Second * 3)
			continue
		}
		if code/100 != 2 {
			log.Info("fail, code: " + fmt.Sprint(code))
			time.Sleep(time.Second * 3)
		} else {
			log.Info("success")
			break
		}

	}
}
