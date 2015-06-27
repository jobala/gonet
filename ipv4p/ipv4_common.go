package ipv4p

import (
	"net"
	"time"
	//"github.com/hsheth2/logs"
)

const (
	UDP_PROTO = 17
	TCP_PROTO = 6
)

const (
	IP_HEADER_LEN     = 20
	MAX_IP_PACKET_LEN = 65535
)

const (
	DEFAULT_TTL = 64
)

const (
	MTU              = 1500
	FRAGMENT_TIMEOUT = time.Second * 5
)

func Checksum(data []byte) uint16 {
	//logs.Trace.Println(data)
	totalSum := uint64(0)
	for ind, elem := range data {
		if ind%2 == 0 {
			totalSum += (uint64(elem) << 8)
		} else {
			totalSum += uint64(elem)
		}
	}
	//fmt.Println("Checksum total: ", totalSum)

	for prefix := (totalSum >> 16); prefix != 0; prefix = (totalSum >> 16) {
		//        fmt.Println(prefix)
		//        fmt.Println(totalSum)
		//        fmt.Println(totalSum & 0xffff)
		totalSum = uint64(totalSum&0xffff) + prefix
	}
	//fmt.Println("Checksum after carry: ", totalSum)

	carried := uint16(totalSum)

	flip := ^carried
	//fmt.Println("Checksum: ", flip)

	return flip
}

func calculateIPChecksum(header []byte) uint16 {
	header[10] = 0
	header[11] = 0
	//logs.Trace.Println("Compute IP Checksum")
	return Checksum(header)
}
func verifyIPChecksum(header []byte) bool {
	//logs.Trace.Println("Verify Checksum")
	return Checksum(header) == 0
}

func CalcTransportChecksum(header []byte, srcIP, dstIP string, headerLen uint16, proto uint8) uint16 {
	//logs.Trace.Println("Transport Checksum")
	ips := append(net.ParseIP(srcIP)[12:], net.ParseIP(dstIP)[12:]...)
	return Checksum(append(append(ips, []byte{0, byte(proto), byte(headerLen >> 8), byte(headerLen)}...), header...))
}

func VerifyTransportChecksum(header []byte, srcIP, dstIP string, headerLen uint16, proto uint8) bool {
	// TODO: do TCP/UDP checksum verification
	return true
}

func GetSrcIP(dst string) string {
	return "127.0.0.1" // TODO: fix this function to be dynamic
}
