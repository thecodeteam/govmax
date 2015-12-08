package apiv1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

var vmh *VMHost

func init() {
	vcenterHost := os.Getenv("GOVMAX_VMHOST_HOST")
	vcenterUsername := os.Getenv("GOVMAX_VMHOST_USERNAME")
	vcenterPassword := os.Getenv("GOVMAX_VMHOST_PASSWORD")
	vcenterInsecure, _ := strconv.ParseBool(os.Getenv("GOVMAX_VMHOST_INSECURE"))

	var err error
	vmh, err = NewVMHost(vcenterInsecure, vcenterHost, vcenterUsername, vcenterPassword)
	if err != nil {
		panic(err)
	}
}

func TestGetLocalMac(*testing.T) {
	mac, err := getLocalMAC()
	if err != nil {
		panic(err)
	}

	fmt.Println(mac)
}

func TestFindVM(*testing.T) {
	vm, err := vmh.findVM(vmh.mac)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vm))
}

func TestFindHosts(*testing.T) {
	hosts, err := vmh.FindHosts(vmh.Vm)
	if err != nil {
		panic(err)
	}

	for _, host := range hosts {
		fmt.Println(fmt.Sprintf("%+v", host))
	}
}

func TestGetHBAWWN(*testing.T) {
	hosts, err := vmh.FindHosts(vmh.Vm)
	if err != nil {
		panic(err)
	}

	wwns, err := vmh.getHBAWWN(hosts)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", strings.Join(wwns, ",")))
}

func TestAttachRDM(*testing.T) {
	if err := vmh.AttachRDM(vmh.Vm, "60000970000196701380533030313142"); err != nil {
		panic(err)
	}
}

func TestDetachRDM(*testing.T) {
	if err := vmh.DetachRDM(vmh.Vm, "60000970000196701380533030313142"); err != nil {
		panic(err)
	}
}
