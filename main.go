package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/goamz/aws"
	//	"github.com/segmentio/go-route53"
	"github.com/gherlein/go-route53"
	"log"
	"net"
	"os"
	//	"strings"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	auth, err := aws.EnvAuth()
	check(err)

	dns := route53.New(auth, aws.USWest2)
	//	check(err)

	myIP := fmt.Sprintf("%s", GetOutboundIP())

	fmt.Printf("%s\n", myIP)

	res, err := dns.Zone("Z1Z7LTSNXL1I6Q").Upsert("A", "r2d2.rbot.cloud", myIP)
	check(err)

	b, err := json.MarshalIndent(res, "", "  ")
	check(err)

	os.Stdout.Write(b)

}
