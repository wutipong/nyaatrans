package main

import (
	"testing"
	"time"
)

func TestParseTorrentItem(t *testing.T) {
	given := `
<item>
	<title>(成年コミック) [杏流ゆいと] 修学旅行にエッチなおもちゃ！？消灯中にぶるぶるイクまで【完全版】</title>
	<link>https://sukebei.nyaa.si/download/2806853.torrent</link>
	<guid isPermaLink="true">https://sukebei.nyaa.si/view/2806853</guid>
	<pubDate>Wed, 18 Sep 2019 13:22:35 -0000</pubDate>
	<nyaa:seeders>190</nyaa:seeders>
	<nyaa:leechers>33</nyaa:leechers>
	<nyaa:downloads>1755</nyaa:downloads>
	<nyaa:infoHash>1400bac36e4d91c8d018a5d17d2f8ed2ec1a403d</nyaa:infoHash>
	<nyaa:categoryId>1_4</nyaa:categoryId>
	<nyaa:category>Art - Manga</nyaa:category>
	<nyaa:size>30.7 MiB</nyaa:size>
	<nyaa:comments>0</nyaa:comments>
	<nyaa:trusted>Yes</nyaa:trusted>
	<nyaa:remake>No</nyaa:remake>
	<description><![CDATA[<a href="https://sukebei.nyaa.si/view/2806853">#2806853 | (成年コミック) [杏流ゆいと] 修学旅行にエッチなおもちゃ！？消灯中にぶるぶるイクまで【完全版】</a> | 30.7 MiB | Art - Manga | 1400BAC36E4D91C8D018A5D17D2F8ED2EC1A403D]]></description>
</item>
	`
	output, err := ParseTorrentItem(given)

	if err != nil {
		t.Error(err)
		return
	}

	if output.Title != "(成年コミック) [杏流ゆいと] 修学旅行にエッチなおもちゃ！？消灯中にぶるぶるイクまで【完全版】" {
		t.Errorf("expected : %s\n return : %s\n",
			"(成年コミック) [杏流ゆいと] 修学旅行にエッチなおもちゃ！？消灯中にぶるぶるイクまで【完全版】",
			output.Title)
	}

	if output.Link != "https://sukebei.nyaa.si/download/2806853.torrent" {
		t.Errorf("expected: %s\n return: %s\n",
			"https://sukebei.nyaa.si/download/2806853.torrent",
			output.Link)
	}
	expectedDate, _ := time.Parse(time.RFC1123Z, "Wed, 18 Sep 2019 13:22:35 -0000")

	if time.Time(output.PubDate).String() != expectedDate.String() {
		t.Errorf("expected: %v\n return: %v\n",
			expectedDate,
			time.Time(output.PubDate))
	}

	if output.Seeder != 190 {
		t.Errorf("expected: %v\n return: %v\n",
			190,
			output.Seeder)
	}

	if output.Leecher != 33 {
		t.Errorf("expected: %v\n return: %v\n",
			33,
			output.Leecher)
	}
}

func TestParseRSS(t *testing.T) {
	input := `
<rss xmlns:atom="http://www.w3.org/2005/Atom" xmlns:nyaa="https://sukebei.nyaa.si/xmlns/nyaa" version="2.0">
	<channel>
		<title>Sukebei - Home - Torrent File RSS</title>
		<description>RSS Feed for Home</description>
		<link>https://sukebei.nyaa.si/</link>
		<atom:link href="https://sukebei.nyaa.si/?page=rss" rel="self" type="application/rss+xml" />
		<item>
			<title>(成年コミック) [杏流ゆいと] 修学旅行にエッチなおもちゃ！？消灯中にぶるぶるイクまで【完全版】</title>
				<link>https://sukebei.nyaa.si/download/2806853.torrent</link>
				<guid isPermaLink="true">https://sukebei.nyaa.si/view/2806853</guid>
				<pubDate>Wed, 18 Sep 2019 13:22:35 -0000</pubDate>

				<nyaa:seeders>190</nyaa:seeders>
				<nyaa:leechers>33</nyaa:leechers>
				<nyaa:downloads>1755</nyaa:downloads>
				<nyaa:infoHash>1400bac36e4d91c8d018a5d17d2f8ed2ec1a403d</nyaa:infoHash>
			<nyaa:categoryId>1_4</nyaa:categoryId>
			<nyaa:category>Art - Manga</nyaa:category>
			<nyaa:size>30.7 MiB</nyaa:size>
			<nyaa:comments>0</nyaa:comments>
			<nyaa:trusted>Yes</nyaa:trusted>
			<nyaa:remake>No</nyaa:remake>
			<description><![CDATA[<a href="https://sukebei.nyaa.si/view/2806853">#2806853 | (成年コミック) [杏流ゆいと] 修学旅行にエッチなおもちゃ！？消灯中にぶるぶるイクまで【完全版】</a> | 30.7 MiB | Art - Manga | 1400BAC36E4D91C8D018A5D17D2F8ED2EC1A403D]]></description>
		</item>
		<item>
			<title>[朝峰テル] milking♥ 第1-7話 [中国翻訳].zip</title>
				<link>https://sukebei.nyaa.si/download/2806755.torrent</link>
				<guid isPermaLink="true">https://sukebei.nyaa.si/view/2806755</guid>
				<pubDate>Wed, 18 Sep 2019 10:28:41 -0000</pubDate>

				<nyaa:seeders>24</nyaa:seeders>
				<nyaa:leechers>15</nyaa:leechers>
				<nyaa:downloads>137</nyaa:downloads>
				<nyaa:infoHash>5534b4ecdcc91b46d4e0e66439bca0573e74432b</nyaa:infoHash>
			<nyaa:categoryId>1_4</nyaa:categoryId>
			<nyaa:category>Art - Manga</nyaa:category>
			<nyaa:size>481.8 MiB</nyaa:size>
			<nyaa:comments>0</nyaa:comments>
			<nyaa:trusted>No</nyaa:trusted>
			<nyaa:remake>No</nyaa:remake>
			<description><![CDATA[<a href="https://sukebei.nyaa.si/view/2806755">#2806755 | [朝峰テル] milking♥ 第1-7話 [中国翻訳].zip</a> | 481.8 MiB | Art - Manga | 5534B4ECDCC91B46D4E0E66439BCA0573E74432B]]></description>
		</item>
		<item>
			<title>(成年コミック) [猫男爵] 甘えて♡吸って♡ [DL版]</title>
				<link>https://sukebei.nyaa.si/download/2806724.torrent</link>
				<guid isPermaLink="true">https://sukebei.nyaa.si/view/2806724</guid>
				<pubDate>Wed, 18 Sep 2019 09:57:34 -0000</pubDate>

				<nyaa:seeders>491</nyaa:seeders>
				<nyaa:leechers>105</nyaa:leechers>
				<nyaa:downloads>6645</nyaa:downloads>
				<nyaa:infoHash>8f666a7a63faa181bbf3d87cee78d32b64b86bdf</nyaa:infoHash>
			<nyaa:categoryId>1_4</nyaa:categoryId>
			<nyaa:category>Art - Manga</nyaa:category>
			<nyaa:size>216.4 MiB</nyaa:size>
			<nyaa:comments>0</nyaa:comments>
			<nyaa:trusted>Yes</nyaa:trusted>
			<nyaa:remake>No</nyaa:remake>
			<description><![CDATA[<a href="https://sukebei.nyaa.si/view/2806724">#2806724 | (成年コミック) [猫男爵] 甘えて♡吸って♡ [DL版]</a> | 216.4 MiB | Art - Manga | 8F666A7A63FAA181BBF3D87CEE78D32B64B86BDF]]></description>
		</item>
	</channel>
</rss>
	`

	rss, err := ParseRSS(input)

	if err != nil {
		t.Error(err)
		return
	}

	if len(rss.Channel.Items) == 0 {
		t.Error("Unable to parse items.")
		return
	}
}
