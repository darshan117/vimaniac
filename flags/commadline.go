package flags

import (
	"errors"
	"flag"
)

// server flag  bool
// connect flag takes ip
// no flags = filename
func CheckFlags() (bool, string, []string, error) {

	var isClient bool
	var ip string
	flag.BoolVar(&isClient, "client", false, "to specify this is the client")
	flag.StringVar(&ip, "connect", "", "tell the ip address along with port number")
	flag.Parse()
	if isClient == true && ip == "" {
		return false, "", []string{""}, errors.New("Please provide the ipaddr as well")
	}
	FileName := flag.Args()
	return isClient, ip, FileName, nil
}
