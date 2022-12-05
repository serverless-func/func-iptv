package main

type channel struct {
	name string
	m3u  string
	epg  string
}

var cctv = []channel{
	{
		name: "Github",
		m3u:  "https://raw.githubusercontent.com/BurningC4/Chinese-IPTV/master/TV-IPV4.m3u",
		epg:  "https://raw.githubusercontent.com/BurningC4/Chinese-IPTV/master/guide.xml",
	},
	{
		name: "FastGit",
		m3u:  "https://raw.fastgit.org/BurningC4/Chinese-IPTV/master/TV-IPV4.m3u",
		epg:  "https://raw.fastgit.org/BurningC4/Chinese-IPTV/master/guide.xml",
	},
	{
		name: "GitHub Proxy",
		m3u:  "https://ghproxy.com/https://raw.githubusercontent.com/BurningC4/Chinese-IPTV/master/TV-IPV4.m3u",
		epg:  "https://ghproxy.com/https://raw.githubusercontent.com/BurningC4/Chinese-IPTV/master/guide.xml",
	},
}
