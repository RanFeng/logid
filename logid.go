package logid

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var (
	layout     = "20060102150405"
	localIP    = "000000000000"
	length     = len(layout) + len(localIP) + len("abcdef")
	maxRandNum = 1<<24 - 1<<20
)

func init() {
	getClientIp()
}

func GenLogID() (logID string) {
	r := Uint32n(uint32(maxRandNum)) + 1<<20
	sb := strings.Builder{}
	sb.Grow(length)
	sb.WriteString(time.Now().Format(layout))
	sb.WriteString(localIP)
	sb.WriteString(strconv.FormatUint(uint64(r), 16))
	return sb.String()
}

func getClientIp() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIP = fmt.Sprintf("%03d%03d%03d%03d", ipnet.IP[12], ipnet.IP[13], ipnet.IP[14], ipnet.IP[15])
				return
			}
		}
	}
	panic("未找到本地IPV4地址")
}
