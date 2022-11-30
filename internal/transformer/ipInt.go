package transformer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

func Ip2int(ip net.IP) (uint32, error) {
	if len(ip) == 16 {
		return 0, errors.New("no sane way to convert ipv6 into uint32")
	}
	return binary.BigEndian.Uint32(ip), nil
}

func Int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func IpAsStringToInt(ip string) (uint32, error) {
	var long uint32
	err := binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	if err != nil {
		return 0, errors.New("can't convert ip to int32")
	}
	return long, nil
}
