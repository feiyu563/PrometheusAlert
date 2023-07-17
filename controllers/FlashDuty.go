package controllers

import (
	"github.com/astaxie/beego"
	"net"
	"strings"
	"time"
)

type FlashDuty struct {
	beego.Controller
}

type PrometheusAlertInput struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"` // key identifying the group of alerts (e.g. to deduplicate)
	Status            string            `json:"status"`
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"` // backlink to the Alertmanager.
	Alerts            []Alert           `json:"alerts" binding:"required"`
}

type Alert struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
}

// PrometheusAlertPush 接受http请求

func parseIpFromAddr(addr string) string {
	strs := strings.Split(addr, ":")
	if len(strs) == 2 {
		ip := net.ParseIP(strs[0])
		if ip != nil {
			return ip.String()
		}
	}

	return ""
}
