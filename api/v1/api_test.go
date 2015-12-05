package apiv1

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
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
	arrays, err := smis.GetStorageArrays()
	//Setup array name for rest of tests
	if err != nil {
		panic(err)
	}

	testingSID = arrays.Entries[0].Content.I_ElementName
}

func TestGetStorageArrays(*testing.T) {

	arrays, err := smis.GetStorageArrays()
	//Setup array name for rest of tests
	if err != nil {
		panic(err)
	}

	testingSID = arrays.Entries[0].Content.I_ElementName
	fmt.Println(testingSID)
	fmt.Println(fmt.Sprintf("%+v", arrays))
}

func TestGetStoragePools(*testing.T) {

	pools, err := smis.GetStoragePools(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", pools))
}

func TestGetMaskingViews(*testing.T) {
	maskingViews, err := smis.GetMaskingViews(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", maskingViews))
}

func TestGetStorageGroups(*testing.T) {

	groups, err := smis.GetStorageGroups(testingSID)
	if err != nil {
		panic(err)
	}

	for _, entry := range groups.Entries {
		fmt.Println(fmt.Sprintf("%+v", entry.Content))
	}

}

func TestGetVolumes(*testing.T) {

	vols, err := smis.GetVolumes(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vols))
}

func TestGetVolume(*testing.T) {

	vols, err := smis.GetVolumes(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vols.Entries[0].Content))
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
		PostVolumesRequestContent: &PostVolumesReqContent{
			AtType:             "http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageConfigurationService",
			ElementName:        "test_vol",
			ElementType:        "2",
			EMCNumberOfDevices: "1",
			Size:               "123",
		},
	}
	queuedJob, _, err := smis.PostVolumes(PostVolRequest, testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", queuedJob))
}

func TestPostCreateGroup(*testing.T) {
	curTime := time.Now()
	PostGroupRequest := &PostGroupReq{
		PostGroupRequestContent: &PostGroupReqContent{
			AtType:    "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerConfigurationService",
			GroupName: "TestingSG_" + curTime.Format("Jan-2-2006--15-04-05"),
			Type:      "4",
		},
	}
	storageGroup, err := smis.PostCreateGroup(PostGroupRequest, testingSID)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%+v", storageGroup))
}

func TestGetStoragePoolSettings(*testing.T) {
	storagePools, err := smis.GetStoragePools(testingSID)
	if err != nil {
		panic(err)
	}

	storagePoolSettings, err := smis.GetStoragePoolSettings(storagePools.Entries[0].Content.I_ElementName, testingSID)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%+v", storagePoolSettings))
}

func TestGetSLOs(*testing.T) {

	SLOs, err := smis.GetSLOs(testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", SLOs)
}

func TestPostVolumeToSG(*testing.T) {

	PostVol2SGRequest := &PostVolumesToSGReq{
		PostVolumesToSGRequestContent: &PostVolumesToSGReqContent{
			AtType: "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerconfigurationService",
			PostVolumesToSGRequestContentMG: &PostVolumesToSGReqContentMG{
				AtType: "http://schemas.emc.com/ecom/edaa/root/emc/SE_DeviceMaskingGroup",
				//Change SMI_sg2 to any existing Storage Group ID
				InstanceID: "SYMMETRIX-+-" + testingSID + "-+-SMI_sg2",
			},
			PostVolumesToSGRequestContentMember: []*PostVolumesToSGReqContentMember{
				&PostVolumesToSGReqContentMember{
					AtType:            "http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageVolume",
					CreationClassName: "Symm_StorageVolume",
					//Change DeviceID to existing Volume ID
					DeviceID:                "00051",
					SystemCreationClassName: "Symm_StorageSystem",
					SystemName:              "SYMMETRIX-+-" + testingSID,
				},
			},
		},
	}

	vol2SG, err := smis.PostVolumesToSG(PostVol2SGRequest, testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vol2SG))
}

func TestPostStorageHardwareID(*testing.T) {
	PostSHIDRequest := &PostStorageHardwareIDReq{
		PostStorageHardwareIDRequestContent: &PostStorageHardwareIDReqContent{
			AtType:    "http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageHardwareIDManagementService",
			IDType:    "2",
			StorageID: "10000000c94e5d22",
		},
	}
	storageGroup, err := smis.PostStorageHardwareID(PostSHIDRequest, testingSID)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%+v", storageGroup))
}

func TestPostInitiatorToHG(*testing.T) {

	PostInit2SGRequest := &PostInitiatorToHGReq{
		PostInitiatorToHGRequestContent: &PostInitiatorToHGReqContent{
			AtType: "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerconfigurationService",
			PostInitiatorToHGRequestContentMG: &PostInitiatorToHGReqContentMG{
				AtType: "http://schemas.emc.com/ecom/edaa/root/emc/SE_InitiatorMaskingGroup",
				//Change test_hg to any existing Host Group ID
				InstanceID: "SYMMETRIX-+-" + testingSID + "-+-test_hg",
			},
			PostInitiatorToHGRequestContentMember: []*PostInitiatorToHGReqContentMember{
				&PostInitiatorToHGReqContentMember{
					AtType: "http://schemas.emc.com/ecom/edaa/root/emc/SE_StorageHardwareID",
					//Change InstanceID to existing Initiator
					InstanceID: "10000000C94E5D22",
				},
			},
		},
	}

	init2HG, err := smis.PostInitiatorToHG(PostInit2SGRequest, testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", init2HG))
}

func TestPostPortToPG(*testing.T) {

	PostPort2PGRequest := &PostPortToPGReq{
		PostPortToPGRequestContent: &PostPortToPGReqContent{
			AtType: "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerconfigurationService",
			PostPortToPGRequestContentMG: &PostPortToPGReqContentMG{
				AtType: "http://schemas.emc.com/ecom/edaa/root/emc/SE_TargetMaskingGroup",
				//Change test_pg to any existing Port Group ID
				InstanceID: "SYMMETRIX-+-" + testingSID + "-+-test_pg",
			},
			PostPortToPGRequestContentMember: []*PostPortToPGReqContentMember{
				&PostPortToPGReqContentMember{
					AtType:            "http://schemas.emc.com/ecom/edaa/root/emc/Symm_FCSCSIProtocolEndpoint",
					CreationClassName: "Symm_FCSCSIProtocolEndpoint",
					//Change Name to existing FE port
					Name: "5000097350159009",
					SystemCreationClassName: "Symm_StorageProcessorSystem",
					//Change to existing director and port
					SystemName: "SYMMETRIX-+-" + testingSID + "-+-FA-1D-+-9",
				},
			},
		},
	}

	port2PG, err := smis.PostPortToPG(PostPort2PGRequest, testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", port2PG))
}

func TestPostCreateMaskingView(*testing.T) {

	PostCreateMaskingViewReq := &PostCreateMaskingViewReq{
		PostCreateMaskingViewRequestContent: &PostCreateMaskingViewReqContent{
			AtType:      "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerconfigurationService",
			ElementName: "Sak_MV_TEST",
			PostInitiatorMaskingGroupRequest: &PostInitiatorMaskingGroupReq{
				AtType:     "http://schemas.emc.com/ecom/edaa/root/emc/SE_InitiatorMaskingGroup",
				InstanceID: "SYMMETRIX-+-000196701380-+-tt",
			},
			PostTargetMaskingGroupRequest: &PostTargetMaskingGroupReq{
				AtType:     "http://schemas.emc.com/ecom/edaa/root/emc/SE_TargetMaskingGroup",
				InstanceID: "SYMMETRIX-+-000196701380-+-ttt",
			},
			PostDeviceMaskingGroupRequest: &PostDeviceMaskingGroupReq{
				AtType:     "http://schemas.emc.com/ecom/edaa/root/emc/SE_DeviceMaskingGroup",
				InstanceID: "SYMMETRIX-+-000196701380-+-SMI_sg2",
			},
		},
	}

	fmt.Println(fmt.Sprintf("%+v", PostCreateMaskingViewReq))

	storageGroup, err := smis.PostCreateMaskingView(PostCreateMaskingViewReq, testingSID)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%+v", storageGroup))
}

func TestGetBaremetalHBA(*testing.T) {

	HBAs, err := GetBaremetalHBA()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", HBAs)
}

func TestRemoveVolumeFromSG(*testing.T) {

	RemVol2SGRequest := &PostVolumesToSGReq{
		PostVolumesToSGRequestContent: &PostVolumesToSGReqContent{
			AtType: "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerconfigurationService",
			PostVolumesToSGRequestContentMG: &PostVolumesToSGReqContentMG{
				AtType: "http://schemas.emc.com/ecom/edaa/root/emc/SE_DeviceMaskingGroup",
				//Change SMI_sg2 to any existing Storage Group ID
				InstanceID: "SYMMETRIX-+-" + testingSID + "-+-Kim_SG",
			},
			PostVolumesToSGRequestContentMember: []*PostVolumesToSGReqContentMember{
				&PostVolumesToSGReqContentMember{
					AtType:            "http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageVolume",
					CreationClassName: "Symm_StorageVolume",
					//Change DeviceID to existing Volume ID
					DeviceID:                "0001",
					SystemCreationClassName: "Symm_StorageSystem",
					SystemName:              "SYMMETRIX-+-" + testingSID,
				},
			},
		},
	}

	rmVol, err := smis.RemoveVolumeFromSG(RemVol2SGRequest, testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", rmVol))
}

func TestRemovePortFromPG(*testing.T) {

	RemPort2PGRequest := &PostPortToPGReq{
		PostPortToPGRequestContent: &PostPortToPGReqContent{
			AtType: "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerconfigurationService",
			PostPortToPGRequestContentMG: &PostPortToPGReqContentMG{
				AtType: "http://schemas.emc.com/ecom/edaa/root/emc/SE_TargetMaskingGroup",
				//Change test_pg to any existing Port Group ID
				InstanceID: "SYMMETRIX-+-" + testingSID + "-+-ttt",
			},
			PostPortToPGRequestContentMember: []*PostPortToPGReqContentMember{
				&PostPortToPGReqContentMember{
					AtType:            "http://schemas.emc.com/ecom/edaa/root/emc/Symm_FCSCSIProtocolEndpoint",
					CreationClassName: "Symm_FCSCSIProtocolEndpoint",
					//Change Name to existing FE port
					Name: "500009735015908a",
					SystemCreationClassName: "Symm_StorageProcessorSystem",
					//Change to existing director and port
					SystemName: "SYMMETRIX-+-" + testingSID + "-+-FA-3D-+-10",
				},
			},
		},
	}

	rmPort, err := smis.RemovePortFromPG(RemPort2PGRequest, testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", rmPort))
}

func TestRemoveInitiatorFromHG(*testing.T) {

	RemInit2SGRequest := &PostInitiatorToHGReq{
		PostInitiatorToHGRequestContent: &PostInitiatorToHGReqContent{
			AtType: "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerconfigurationService",
			PostInitiatorToHGRequestContentMG: &PostInitiatorToHGReqContentMG{
				AtType: "http://schemas.emc.com/ecom/edaa/root/emc/SE_InitiatorMaskingGroup",
				//Change test_hg to any existing Host Group ID
				InstanceID: "SYMMETRIX-+-" + testingSID + "-+-test_hg",
			},
			PostInitiatorToHGRequestContentMember: []*PostInitiatorToHGReqContentMember{
				&PostInitiatorToHGReqContentMember{
					AtType: "http://schemas.emc.com/ecom/edaa/root/emc/SE_StorageHardwareID",
					//Change InstanceID to existing Initiator
					InstanceID: "10000000C94E5D22",
				},
			},
		},
	}

	rmInit, err := smis.RemoveInitiatorFromHG(RemInit2SGRequest, testingSID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", rmInit))
}

func TestPostDeleteGroup(*testing.T) {
	DeleteGroupRequest := &DeleteGroupReq{
		DeleteGroupRequestContent: &DeleteGroupReqContent{
			AtType: "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerConfigurationService",
			DeleteGroupRequestContentMaskingGroup: &DeleteGroupReqContentMaskingGroup{
				//Change AtType to type of Group and InstanceID to existing name of Group
				AtType:     "http://schemas.emc.com/ecom/edaa/root/emc/SE_DeviceMaskingGroup",
				InstanceID: "SYMMETRIX-+-" + testingSID + "-+-Test_SG",
			},
		},
	}
	deleteGroup, err := smis.PostDeleteGroup(DeleteGroupRequest, testingSID)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%+v", deleteGroup))
}

func TestPostDeleteVol(*testing.T) {
	DeleteVolumeRequest := &DeleteVolReq{
		DeleteVolRequestContent: &DeleteVolReqContent{
			AtType: "http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageConfigurationService",
			DeleteVolRequestContentElement: &DeleteVolReqContentElement{
				AtType:                  "http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageVolume",
				DeviceID:                "0001F",
				CreationClassName:       "Symm_StorageVolume",
				SystemName:              "SYMMETRIX-+-" + testingSID,
				SystemCreationClassName: "Symm_StorageSystem",
			},
		},
	}
	deleteVol, err := smis.PostDeleteVol(DeleteVolumeRequest, testingSID)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%+v", deleteVol))
}

func TestPostDeleteMV(*testing.T) {
	DeleteMVRequest := &DeleteMaskingViewReq{
		DeleteMaskingViewRequestContent: &DeleteMaskingViewReqContent{
			AtType: "http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerConfigurationService",
			DeleteMaskingViewRequestContentPC: &DeleteMaskingViewReqContentPC{
				AtType:                  "http://schemas.emc.com/ecom/edaa/root/emc/Symm_LunMaskingView",
				DeviceID:                "Kim_MV",
				CreationClassName:       "Symm_LunMaskingView",
				SystemName:              "SYMMETRIX-+-" + testingSID,
				SystemCreationClassName: "Symm_StorageSystem",
			},
		},
	}

	deleteMV, err := smis.PostDeleteMaskingView(DeleteMVRequest, testingSID)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%+v", deleteMV))
}

func TestGetVolumeByID(*testing.T) {

	vol, err := smis.GetVolumeByID(testingSID, "000A5")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vol))
}

func TestGetVolumeByName(*testing.T) {

	vol, err := smis.GetVolumeByName(testingSID, "Volume 00001")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", vol))
}

func TestPostPortLogins(*testing.T) {

	PostPortLoginsReq := &PostPortLoggedInReq{
		PostPortLoggedInRequestContent: &PostPortLoggedInReqContent{
			PostPortLoggedInRequestHardwareID: &PostPortLoggedInReqHardwareID{
				AtType:     "http://schemas.emc.com/ecom/edaa/root/emc/SE_StorageHardwareID",
				InstanceID: "10000000C94E5D22",
			},
			AtType: "http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageHardwareIDManagementService",
		},
	}

	portValues, err := smis.PostPortLogins(PostPortLoginsReq, testingSID)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(portValues); i++ {
		fmt.Println("Port Number:" + portValues[i].PortNumber + " Director:" + portValues[i].Director + " WWN:" + portValues[i].WWN)
	}
}
