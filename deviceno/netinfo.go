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
	Info struct {
		IP  []string
		Mac string
	}
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
		var NIP []string
		addrs, err := i.Addrs()
		if err != nil {
			log.Error(err)
			nInfo.Error = nInfo.Error + err.Error()
			continue
		}
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if ok && !ipnet.IP.IsLoopback() {
				NIP = append(NIP, ipnet.IP.String())
			}
		}
		nif.Info.IP = NIP
		nif.Info.Mac = i.HardwareAddr.String()
		nInfo.NetInterfaces = append(nInfo.NetInterfaces, nif)
	}
	return nInfo
}
