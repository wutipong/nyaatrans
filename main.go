package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apaxa-go/eval"
	"github.com/namsral/flag"
)

func filter(item NyaaTorrentItem, expr *eval.Expression) (result bool, err error) {
	arg := eval.Args{
		"item": eval.MakeDataRegularInterface(item),
	}

	r, err := expr.EvalToInterface(arg)
	if err != nil {
		return
	}

	result = r.(bool)

	return

}

//FilterNyaaItems filter out items that does not match the criteria.
func FilterNyaaItems(items []NyaaTorrentItem, expr string) []NyaaTorrentItem {
	var out []NyaaTorrentItem

	exprObj, err := eval.ParseString(expr, "")
	if err != nil {
		return out
	}

	for _, i := range items {
		if result, e := filter(i, exprObj); e != nil || !result {
			continue
		}

		out = append(out, i)
	}

	return out
}

func main() {
	rssURL := flag.String("rss", "https://sukebei.nyaa.si/?page=rss&c=1_4&f=0", "rss url")
	transURL := flag.String("transmission", "http://localhost:9091/transmission/rpc", "Transmission RPC url")
	condition := flag.String("condition", "item.Seeder > 100", "condition")
	path := flag.String("download_path", "/mnt/storage1/manga", "download path")
	runAt := flag.String("run_at", "", "the time the task is run. download immediately if not specified.")

	help := flag.Bool("help", false, "Print Help Message")

	flag.Parse()

	if *help {
		fmt.Println("Nyaa->Transmission Daemon")
		flag.Usage()
		return
	}

	log.Println("Nyaa->Transmission Daemon")

	var nextRun time.Time
	if *runAt == "" {
		nextRun = time.Now()
	} else {
		t, err := time.Parse(time.Kitchen, *runAt)
		if err != nil {
			log.Println(err)
			return
		}

		now := time.Now()
		nextRun = time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), now.Location())
		log.Printf("next run: %v", nextRun)
	}

	running := true

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		running = false
	}()

	for running {
		now := time.Now()
		if now.After(nextRun) {
			log.Println("begin adding task.")
			downloads(*rssURL, *condition, *transURL, *path)
			log.Println("done adding task.")

			if *runAt == "" {
				break
			}

			nextRun = nextRun.AddDate(0, 0, 1)
			log.Printf("next run: %v", nextRun)

		}

		time.Sleep(time.Minute)
	}
}

func downloads(rssURL, condition, transURL, path string) {
	items, err := FetchTorrentItem(rssURL)

	if err != nil {
		log.Println(err)
		return
	}

	items = FilterNyaaItems(items, condition)

	session, err := TransGetSession(transURL)
	if err != nil {
		log.Println(err)
		return
	}

	for _, i := range items {
		log.Printf("Adding torrent: %s \n", i.Title)
		_, err := TransAddTorrent(transURL, i.Link, session, path)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
