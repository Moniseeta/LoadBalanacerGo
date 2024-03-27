package main

import (
	"errors"
	roundrobin "github.com/thegeekyasian/round-robin-go"
)

/* Implementation
-> Input - list of IPs
-> Load balancer, pick the IP from list of IPs using round robin.
-> Output - redirect traffic to that IP

*/

type IPStates struct {
	IPStatesList []*IPState
}
type IPState struct {
	IP        string
	Connected bool
}

func NewIPStates() *IPStates {
	return &IPStates{}
}
func main() {

}

func (ipSt *IPStates) DiscoverIps(ips []string) error {
	if len(ips) == 0 {
		return errors.New("no IPs provided for discovery")
	}
	for _, ip := range ips {
		ipSt.IPStatesList = append(ipSt.IPStatesList, IPState{
			IP:        ip,
			Connected: true,
		})
	}
	return nil
}

func (ipSt *IPStates) SelectIp() (string, error) {

}

func (ipSt *IPStates) AddToRoundRobin() (robin *roundrobin.RoundRobin[IPState], err error) {
	robin, err = roundrobin.New(ipSt.IPStatesList...)
	for _, ip := range ipSt.IPStatesList {

	}
}

func (ipSt *IPStates) RedirectLoad(ip string) error {

}
