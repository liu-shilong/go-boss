package util

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"strings"
)

func Ip2region(ip string) any {
	var dbPath = "public/assets/geo/ip2region.xdb"
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return nil
	}

	defer searcher.Close()

	geo, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
		return nil
	}

	location := strings.Split(geo, "|")
	info := map[string]string{
		"country":  location[0],
		"province": location[2],
		"city":     location[3],
		"isp":      location[4],
	}

	return info
}
