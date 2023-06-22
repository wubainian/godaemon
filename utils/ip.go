package utils

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

func CheckIp(ip string) bool {
	re := regexp.MustCompile(`((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)`)
	return re.MatchString(ip)
}

// ResolveDomain 解析域名为ip
func ResolveDomain(name string) (string, error) {
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		return "", err
	}
	return addr.String(), nil
}

func CheckLocalIp(ip string) bool {
	ip = strings.TrimSpace(ip)
	if strings.HasPrefix(ip, "10.") {
		return true
	}
	ipObj := net.ParseIP(ip)
	ipV4Obj := ipObj.To4()
	if ipV4Obj == nil {
		return false
	}
	if len(ipV4Obj) > 2 {
		if ipV4Obj[0] == 172 && (ipV4Obj[1] >= 16 && ipV4Obj[1] <= 31) {
			return true
		}
	}
	if strings.HasPrefix(ip, "192.168.") {
		return true
	}

	return false
}
