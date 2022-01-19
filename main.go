package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/apaxa-go/eval"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	rssURL := os.Getenv("RSS_URL")
	transURL := os.Getenv("TRANSMISSION_URL")
	condition := os.Getenv("CONDITION")
	path := os.Getenv("DOWNLOAD_PATH")
	runAt := os.Getenv("RUN_AT")
	dryRun := os.Getenv("DRY_RUN")

	help := flag.Bool("help", false, "Print Help Message")

	flag.Parse()

	if *help {
		fmt.Println("Nyaa->Transmission Daemon")
		fmt.Println(`
Nyaa->Transmission add torrents from nyaa's RSS feed to transmission. It can run immediately or at a scheduled time.

The configuration can be done through environment variables. '.env' file is also supported.

RSS_URL         : Nyaa's rss feed url.
TRANSMISSION_URL: Transmission RPC url, ie. http://localhost:9091/transmission/rpc
DOWNLOAD_PATH   : Download path. Can be left blank for default location.
RUN_AT          : The scheduled time. Left blank to run the task immediately.
                  Use the format "18:00PM"

CONDITION       : Condition string. The torrent will be added only the condition is met.
                  Condition string should look something like "item.Seeder > 100".

                  item contains Seeder, Leecher, Title and PubDate fields.
		`)
		return
	}

	log.Println("Nyaa->Transmission Daemon")

	log.Printf("RSS URL: %s", rssURL)
	log.Printf("Transmission URL: %s", transURL)
	log.Printf("Condition: %s", condition)
	log.Printf("Download Path: %s", path)
	if runAt != "" {
		log.Printf("Run at: %s", runAt)
	}

	if condition == "" {
		condition = "true"
	}

	if rssURL == "" {
		log.Fatal("RSS URL is required. Terminated.")
	}

	if transURL == "" {
		log.Fatal("Transmission URL is required. Terminated.")
	}

	doDryRun := false
	if b, err := strconv.ParseBool(dryRun); err == nil {
		doDryRun = b
	}

	if runAt == "" {
		log.Println("begin adding task.")
		Perform(rssURL, condition, transURL, path, doDryRun)
		log.Println("done adding task.")
		return
	}

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Days().At(runAt).Do(func() {
		log.Println("begin adding task.")
		Perform(rssURL, condition, transURL, path, doDryRun)
		log.Println("done adding task.")
	})

	s.StartBlocking()

}

// Perform read RSS feeds and add the torrent to transmission.
func Perform(rssURL, condition, transURL, path string, dryRun bool) {
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
		if dryRun {
			log.Println("Skipped")
			continue
		}
		resp, err := TransAddTorrent(transURL, i.Link, session, path)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("Response: %s", resp.Result)
	}
}
