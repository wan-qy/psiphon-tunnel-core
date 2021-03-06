package psiphon

import (
	"encoding/binary"
	"io/ioutil"
	"math/rand"
	"net"
	"testing"
)

var list networkList
var isLocalAddr bool

func Benchmark_NewNetworkList(b *testing.B) {

	routesData, err := ioutil.ReadFile("test_routes.dat")
	if err != nil {
		b.Skipf("can't load test routes file: %s", err)
	}

	for n := 0; n < b.N; n++ {
		list, _ = NewNetworkList(routesData)
	}
}

func Benchmark_containsRandomAddr(b *testing.B) {

	if list == nil {
		b.Skipf("no test routes file")
	}

	rand.Seed(0)
	for n := 0; n < b.N; n++ {
		ip := make([]byte, 4)
		binary.BigEndian.PutUint32(ip, rand.Uint32())
		isLocalAddr = list.ContainsIpAddress(net.IP(ip))
	}
}
