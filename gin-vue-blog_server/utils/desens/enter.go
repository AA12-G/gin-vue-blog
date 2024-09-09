package desens

import "strings"

func DesensitizationEmail(email string) string {
	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}

func DesensitizationTel(tel string) string {
	// 18545627894
	// 185****7894
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}
