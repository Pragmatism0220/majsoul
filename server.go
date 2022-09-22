package majsoul

import (
	"fmt"
	"github.com/go-ping/ping"
	"net/url"
	"runtime"
	"strings"
)

func serverPing(target string) bool {
	pinger, err := ping.NewPinger(target)
	if runtime.GOOS == "windows" {
		pinger.SetPrivileged(true) // windows support
	}
	if err != nil {
		return false
	}
	pinger.Count = 3
	err = pinger.Run()
	if err != nil {
		return false
	}
	stats := pinger.Statistics()
	if stats.PacketsRecv >= 1 {
		return true
	}
	return false
}

type majsoulConfig struct {
	IP []struct {
		RegionURLs []struct {
			URL    string `json:"url"`    // https://lb-hw.maj-soul.com/api/v0/recommend_list
			OB_URL string `json:"ob_url"` // wss://live-hw.maj-soul.com/ob
		} `json:"region_urls"`
	} `json:"ip"`
}

func (c *majsoulConfig) apiGetMainlandRecommendListURL() (urls []string) {
	if len(c.IP) == 0 {
		return urls
	}
	for _, e := range c.IP[0].RegionURLs {
		rawURL := e.URL
		u, err := url.Parse(rawURL)
		if err != nil {
			return urls
		}
		if serverPing(u.Hostname()) {
			urls = append(urls, rawURL+"?service=ws-gateway&protocol=ws&ssl=true")
		}
	}
	return urls
}

func (c *majsoulConfig) apiGetMainlandRecommendListURLWithLocation(location string) (urls []string) {
	if len(c.IP) == 0 {
		return urls
	}
	for _, e := range c.IP[0].RegionURLs {
		rawURL := e.URL
		u, err := url.Parse(rawURL)
		if err != nil {
			return urls
		}
		if serverPing(u.Hostname()) {
			urls = append(urls, rawURL+"?service=ws-game-gateway&protocol=ws&ssl=true&location="+location)
		}
	}
	return urls
}

func getConfig(apiGetConfigURL string) (config *majsoulConfig, err error) {
	config = &majsoulConfig{}
	err = get(apiGetConfigURL, config)

	// url replace
	if len(config.IP) == 0 {
		return
	}
	for idx, _ := range config.IP[0].RegionURLs {
		config.IP[0].RegionURLs[idx].URL = strings.Replace(config.IP[0].RegionURLs[idx].URL, "maj-soul.com", "maj-soul.net", -1)
		config.IP[0].RegionURLs[idx].OB_URL = strings.Replace(config.IP[0].RegionURLs[idx].OB_URL, "maj-soul.com", "maj-soul.net", -1)
	}
	return
}

// {"servers":["gateway-hw.maj-soul.com:443", "gateway-v2.maj-soul.com:443", "gateway-cdn.maj-soul.com:443", "gateway-sy.maj-soul.com:443"]}
type recommendServers struct {
	Servers []string `json:"servers"`
}

func getRecommendServers(apiGetRecommendServersURL []string) (servers []string, err error) {
	recommendServers := new(recommendServers)
	for _, e := range apiGetRecommendServersURL {
		if err = get(e, &recommendServers); err != nil {
			continue
		}
		servers = append(servers, recommendServers.Servers...)
	}
	return servers, nil
}

// 获取雀魂 WebSocket 服务器地址
func GetMajsoulWebSocketURLs() (murl []string, err error) {
	version, err := GetMajsoulVersion(ApiGetVersionZH)
	if err != nil {
		return
	}

	apiGetConfigURL := fmt.Sprintf(apiGetConfigFormatZH, version.ResVersion)
	config, err := getConfig(apiGetConfigURL)
	if err != nil {
		return
	}

	apiGetMainlandRecommendListURL := config.apiGetMainlandRecommendListURL()
	for idx, _ := range apiGetMainlandRecommendListURL {
		apiGetMainlandRecommendListURL[idx] = strings.Replace(apiGetMainlandRecommendListURL[idx], "maj-soul.com", "maj-soul.net", -1)
	}
	servers, err := getRecommendServers(apiGetMainlandRecommendListURL)
	if err != nil {
		return
	}
	if len(servers) == 0 {
		return murl, fmt.Errorf("维护中，没有可用的服务器地址")
	}
	for _, host := range servers {
		tmpUrl := "wss://" + host
		u, _ := url.Parse(tmpUrl)
		murl = append(murl, u.Scheme+"://"+u.Hostname()+"/")
	}
	return
}
