package deviceno

import (
	"github.com/golangrustnode/log"
	"net"
)

type NetInfo struct {
	NetInterfaces []NetInterface
	Error         string
}

type NetInterface struct {
	Name string
	IP   []string
}

func GetNetInfo() NetInfo {
	nInfo := NetInfo{}
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Error(err)
		return NetInfo{Error: err.Error()}
	}
	for _, i := range ifaces {
		nif := NetInterface{
			Name: i.Name,
		}

		addrs, err := i.Addrs()
		if err != nil {
			log.Error(err)
			nInfo.Error = nInfo.Error + err.Error()
			continue
		}
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if ok && !ipnet.IP.IsLoopback() {
				nif.IP = append(nif.IP, ipnet.IP.String())
			}
		}
		nInfo.NetInterfaces = append(nInfo.NetInterfaces, nif)
	}
	return nInfo
}
