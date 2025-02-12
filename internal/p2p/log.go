package p2p

import (
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
	astLog "github.com/n42blockchain/N42/log"
	"strconv"
	"strings"
)

var log = astLog.New("prefix", "p2p")

func logIPAddr(id peer.ID, addrs ...ma.Multiaddr) {
	var correctAddr ma.Multiaddr
	for _, addr := range addrs {
		if strings.Contains(addr.String(), "/ip4/") || strings.Contains(addr.String(), "/ip6/") {
			correctAddr = addr
			break
		}
	}
	if correctAddr != nil {
		log.Info("Node started p2p server", "multiAddr", correctAddr.String()+"/p2p/"+id.String())
	}
}

func logExternalIPAddr(id peer.ID, addr string, port int) {
	if addr != "" {
		multiAddr, err := MultiAddressBuilder(addr, uint(port))
		if err != nil {
			log.Error("Could not create multiaddress", "err", err)
			return
		}
		log.Info("Node started external p2p server", "multiAddr", multiAddr.String()+"/p2p/"+id.String())
	}
}

func logExternalDNSAddr(id peer.ID, addr string, port int) {
	if addr != "" {
		p := strconv.FormatUint(uint64(port), 10)
		log.Info("Node started external p2p server", "multiAddr", "/dns4/"+addr+"/tcp/"+p+"/p2p/"+id.String())
	}
}
