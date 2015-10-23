package apiv1

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

var smis *SMIS

func init() {
	host := os.Getenv("GOVMAX_SMISHOST")
	port := os.Getenv("GOVMAX_SMISPORT")
	insecure, _ := strconv.ParseBool(os.Getenv("GOVMAX_INSECURE"))
	username := os.Getenv("GOVMAX_USERNAME")
	password := os.Getenv("GOVMAX_PASSWORD")

	var err error
	smis, err = New(host, port, insecure, username, password)
	if err != nil {
		panic(err)
	}
}

/*
func TestGetStoragePools(*testing.T) {

	pools, err := smis.GetStoragePools("000196701380")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", pools))
}

func TestGetDeviceMaskingViews(*testing.T) {

	views, err := smis.GetDeviceMaskingViews("000196701380")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", views))
}


func TestGetStorageGroups(*testing.T) {

	groups, err := smis.GetStorageGroups("000196701380")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", groups))
}

func TestGetStorageVolumes(*testing.T) {

	vols, err := smis.GetStorageVolumes("000196701380")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vols))
}

func TestPortGroups(*testing.T) {

	vols, err := smis.GetPortGroups("000196701380")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vols))
}*/

func TestInitiatorGroups(*testing.T) {

	vols, err := smis.GetInitiatorGroups("000196701380")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vols))
}



