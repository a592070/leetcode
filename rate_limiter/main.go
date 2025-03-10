package main

import (
	"fmt"
	"net/http"
)

func main() {
	rs := rateLimit([]string{
		"https://www.baidu.com",
		"https://www.taobao.com",
		"https://www.taobao.com",
		"https://www.google.com",
		"https://www.google.com",
		"https://www.google.com", //x
		"https://www.yahoo.com",
		"https://www.taobao.com",
		"https://www.taobao.com", //x
		"https://www.taobao.com", //x
		"https://www.taobao.com", //x
	})
	fmt.Printf("%+v\n", rs)

}

type rateLimiter struct {
	ShortWindowSize int // Number of requests in short window
	ShortWindowTime int // Short window duration
	LongWindowSize  int // Number of requests in long window
	LongWindowTime  int // Long window duration
}

type hostLimit struct {
	shortWindow []int
	longWindow  []int
}

func rateLimit(hosts []string) []string {
	rl := rateLimiter{
		ShortWindowSize: 2,
		ShortWindowTime: 5,
		LongWindowSize:  3,
		LongWindowTime:  10,
	}
	hostsLimit := map[string]*hostLimit{}
	result := []string{}
	for i, host := range hosts {
		if _, exists := hostsLimit[host]; !exists {
			hostsLimit[host] = &hostLimit{
				shortWindow: make([]int, 0),
				longWindow:  make([]int, 0),
			}
		}

		now := i
		hostsLimit[host].shortWindow = cleanExpired(hostsLimit[host].shortWindow, now, rl.ShortWindowTime)
		hostsLimit[host].longWindow = cleanExpired(hostsLimit[host].longWindow, now, rl.LongWindowTime)

		resp := fmt.Sprintf("[%d-%d]", i, http.StatusOK)
		if len(hostsLimit[host].shortWindow) >= rl.ShortWindowSize {
			resp = fmt.Sprintf("[%d-%d too many request within %d]", i, http.StatusTooManyRequests, rl.ShortWindowTime)
		} else if len(hostsLimit[host].longWindow) >= rl.LongWindowSize {
			resp = fmt.Sprintf("[%d-%d too many request within %d]", i, http.StatusTooManyRequests, rl.LongWindowTime)
		}

		hostsLimit[host].shortWindow = append(hostsLimit[host].shortWindow, now)
		hostsLimit[host].longWindow = append(hostsLimit[host].longWindow, now)
		result = append(result, resp)

	}
	return result
}

func cleanExpired(window []int, now int, duration int) []int {
	cutoff := now - duration
	valid := make([]int, 0)
	for _, t := range window {
		if t > cutoff {
			valid = append(valid, t)
		}
	}
	return valid
}
