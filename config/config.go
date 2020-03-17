package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	AsUpstream              bool              `json:"as-upstream"`
	LHost                   string            `json:"listen-host"`
	LPort                   int               `json:"listen-port"`
	ProxyType               string            `json:"proxy-type"`
	ProxyTimeout            int               `json:"proxy-timeout"`
	ProxyScope              string            `json:"proxy-scope"`
	BypassHosts             []string          `json:"bypass-hosts"`
	BypassSrcIPs            []string          `json:"bypass-src-ips"`
	HTTPProxyHost           string            `json:"http-proxy-host"`
	HTTPProxyPort           int               `json:"http-proxy-port"`
	HTTPProxyAuthUser       string            `json:"http-proxy-auth-user"`
	HTTPProxyAuthPassword   string            `json:"http-proxy-auth-password"`
	SSHost                  string            `json:"ss-host"`
	SSPort                  int               `json:"ss-port"`
	SSCphierMethod          string            `json:"ss-chpier-method"`
	SSPasswd                string            `json:"ss-passwd"`
	TLSHost                 string            `json:"tls-host"`
	TLSPort                 int               `json:"tls-port"`
	TLSToken                string            `json:"tls-token"`
	SOCKS5Host              string            `json:"socks5-host"`
	SOCKS5Port              int               `json:"socks5-port"`
	SOCKS5AuthUser          string            `json:"socks5-auth-user"`
	SOCKS5AuthPassword      string            `json:"socks5-auth-password"`
	DNSLoggingFile          string            `json:"dns-logging-file"`
	CNDNS                   string            `json:"cn-dns"`
	FQDNS                   string            `json:"fq-dns"`
	EnableDNSCache          bool              `json:"enable-dns-cache"`
	EnforceTTL              uint32            `json:"enforce-ttl"`
	DNSPrefetchEnable       bool              `json:"dns-prefetch-enable"`
	DNSPrefetchCount        int               `json:"dns-prefetch-count"`
	DNSPrefetchInterval     int               `json:"dns-prefetch-interval"`
	DisableQTypes           []string          `json:"disable-qtypes"`
	ForceFQ                 []string          `json:"force-fq"`
	HostMap                 map[string]string `json:"host-map"`
	BlockHostFile           string            `json:"block-host-file"`
	BlockHosts              []string          `json:"block-hosts"`
	Mode                    string            `json:"mode"`
	UpstreamType            string            `json:"upstream-type"`
	UpstreamTLSServerListen string            `json:"upstream-tls-server-listen"`
	UpstreamTLSKey          string            `json:"upstream-tls-key"`
	UpstreamTLSCRT          string            `json:"upstream-tls-crt"`
	UpstreamTLSToken        string            `json:"upstream-tls-token"`
}

func LoadConfig(configPath string) (*Config, error) {
	var config Config
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}