package main

import (
	"fmt"
)

func FilterNyaaItems(in []NyaaTorrentItem) []NyaaTorrentItem {
	var out []NyaaTorrentItem

	for _, i := range in {
		if i.Seeder+i.Leecher < 100 {
			continue
		}

		out = append(out, i)
	}

	return out

}

func main() {
	items, err := FetchTorrentItem("https://sukebei.nyaa.si/?page=rss&f=0&c=1_4&q=")

	if err != nil {
		fmt.Println(err)
		return
	}

	items = FilterNyaaItems(items)

	for _, i := range items {
		fmt.Println(i.Link)
	}
}
