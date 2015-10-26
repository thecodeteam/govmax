package apiv1

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

var smis *SMIS


var testingSID string

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

func TestGetStorageArrays(*testing.T) {

	arrays, err := smis.GetStorageArrays()
	//Setup array name for rest of tests
	testingSID = arrays.Entries[0].Content.I_ElementName
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v",arrays))
}

func TestGetStoragePools(*testing.T) {

	pools, err := smis.GetStoragePools(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", pools))
}

func TestGetDeviceMaskingViews(*testing.T) {

	views, err := smis.GetDeviceMaskingViews(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", views))
}


func TestGetStorageGroups(*testing.T) {

	groups, err := smis.GetStorageGroups(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", groups))
}

func TestGetStorageVolumes(*testing.T) {

	vols, err := smis.GetStorageVolumes(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vols))
}

func TestPortGroups(*testing.T) {

	vols, err := smis.GetPortGroups(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vols))
}

func TestInitiatorGroups(*testing.T) {

	vols, err := smis.GetHostGroups(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vols))
}


func TestPostVolumes(*testing.T) {

	PostVolRequest := &PostVolumesReq{
	PostVolumesReqContent : PostVolumesReqContent{
			AtType : "http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageConfigurationService",
			EMCNumberOfDevices : "1",
			ElementType : "2",
			Size : "123",
		},
	}
	//fmt.Println(fmt.Sprintf("%+v",PostVolRequest))
	queuedJob, err := smis.PostVolumes(PostVolRequest,testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v",queuedJob))
}

