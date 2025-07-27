package code

import (
	"net"
)

/*
判断是不是合法的ipv4地址
官方一把梭工具没秒掉,只过了90的点,不知道给的啥样例,还有10%连官方的库都过不了?
怀疑是认知层面的问题
*/
func IsValidIPv4Address(str_to_check string) bool {
	// write code here
	//splitPart := strings.Split(str_to_check, ".")
	//if len(splitPart) != 4 {
	//	return false
	//}
	//for _, part := range splitPart {
	//	if strings.Contains(part, "/") {
	//		return false
	//	}
	//	if part >= "0" && part <= "255" {
	//		ip := net.ParseIP(str_to_check)
	//		if ip == nil {
	//			return false
	//		}
	//
	//		continue
	//	} else {
	//		return false
	//	}
	//}
	//return true
	ip := net.ParseIP(str_to_check)
	if ip == nil {
		return false
	}
	return ip.To4() != nil
}
