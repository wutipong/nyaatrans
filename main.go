package main

import (
	"flag"
	"log"

	"github.com/apaxa-go/eval"
)

func filter(item NyaaTorrentItem) (result bool, err error) {
	arg := eval.Args{
		"item": eval.MakeDataRegularInterface(item),
	}

	src := "item.Seeder > 100"
	expr, err := eval.ParseString(src, "")

	if err != nil {
		return
	}

	r, err := expr.EvalToInterface(arg)
	if err != nil {
		return
	}

	result = r.(bool)

	return

}

func FilterNyaaItems(in []NyaaTorrentItem, min int) []NyaaTorrentItem {
	var out []NyaaTorrentItem

	for _, i := range in {
		//if i.Seeder+i.Leecher < min {
		if r, e := filter(i); e != nil || !r {
			continue
		}

		out = append(out, i)
	}

	return out

}

func main() {
	rssURL := flag.String("rss", "https://sukebei.nyaa.si/?page=rss&c=1_4&f=0", "rss url")
	transURL := flag.String("transmission", "http://localhost:9091/transmission/rpc", "Transmission RPC url")
	minPeer := flag.Int("min_peers", 100, "minimum peers")
	path := flag.String("path", "/mnt/storage1/manga", "download path")

	help := flag.Bool("help", false, "Print Help Message")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	items, err := FetchTorrentItem(*rssURL)

	if err != nil {
		log.Println(err)
		return
	}

	items = FilterNyaaItems(items, *minPeer)

	session, err := TransGetSession(*transURL)
	if err != nil {
		log.Println(err)
		return
	}

	for _, i := range items {
		_, err := TransAddTorrent(*transURL, i.Link, session, *path)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
