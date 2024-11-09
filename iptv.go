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
		name: "gh-proxy",
		m3u:  "https://gh-proxy.com/raw.githubusercontent.com/BurningC4/Chinese-IPTV/master/TV-IPV4.m3u",
		epg:  "https://gh-proxy.com/raw.githubusercontent.com/BurningC4/Chinese-IPTV/master/guide.xml",
	},
	{
		name: "ghp.ci",
		m3u:  "https://ghp.ci/https://raw.githubusercontent.com/BurningC4/Chinese-IPTV/master/TV-IPV4.m3u",
		epg:  "https://ghp.ci/https://raw.githubusercontent.com/BurningC4/Chinese-IPTV/master/guide.xml",
	},
}
