package apiv1

import ()

type GetStoragePoolsResp struct {
	Entries []struct {
		Content struct {
			I_ClientSettableUsage         []interface{} `json:"i$ClientSettableUsage"`
			I_ConsumedResourceUnits       string        `json:"i$ConsumedResourceUnits"`
			I_Description                 interface{}   `json:"i$Description"`
			I_EMCCanBeUsedForRemoteMirror bool          `json:"i$EMCCanBeUsedForRemoteMirror"`
			I_EMCDataFormat               string        `json:"i$EMCDataFormat"`
			I_EMCDefaultSRPforFBAvolumes  bool          `json:"i$EMCDefaultSRPforFBAvolumes"`
			I_EMCDiskDriveType            string        `json:"i$EMCDiskDriveType"`
			I_EMCIsBound                  bool          `json:"i$EMCIsBound"`
			I_EMCLocality                 []int         `json:"i$EMCLocality"`
			I_EMCOversubscribedCapacity   int           `json:"i$EMCOversubscribedCapacity"`
			I_EMCPercentReservedCapacity  int           `json:"i$EMCPercentReservedCapacity"`
			I_EMCPercentSubscribed        int           `json:"i$EMCPercentSubscribed"`
			I_EMCPercentageUsed           int           `json:"i$EMCPercentageUsed"`
			I_EMCRemoteMirrorSpace        int           `json:"i$EMCRemoteMirrorSpace"`
			I_EMCSnapshotSpace            int           `json:"i$EMCSnapshotSpace"`
			I_EMCSubscribedCapacity       int           `json:"i$EMCSubscribedCapacity"`
			I_ElementName                 string        `json:"i$ElementName"`
			I_ElementsShareSpace          bool          `json:"i$ElementsShareSpace"`
			I_HealthState                 int           `json:"i$HealthState"`
			I_InstanceID                  string        `json:"i$InstanceID"`
			I_LowSpaceWarningThreshold    int           `json:"i$LowSpaceWarningThreshold"`
			I_OperationalStatus           []int         `json:"i$OperationalStatus"`
			I_PoolID                      string        `json:"i$PoolID"`
			I_Primordial                  bool          `json:"i$Primordial"`
			I_RemainingManagedSpace       int           `json:"i$RemainingManagedSpace"`
			I_SpaceLimit                  int           `json:"i$SpaceLimit"`
			I_SpaceLimitDetermination     int           `json:"i$SpaceLimitDetermination"`
			I_StatusDescriptions          []string      `json:"i$StatusDescriptions"`
			I_ThinProvisionMetaDataSpace  int           `json:"i$ThinProvisionMetaDataSpace"`
			I_TotalManagedSpace           int           `json:"i$TotalManagedSpace"`
			I_Usage                       int           `json:"i$Usage"`
		} `json:"content"`
	} `json:"entries"`
}

//Get a list of all Storage Pools associated with SymmID
func (smis *SMIS) GetStoragePools(sid string) (resp *GetStoragePoolsResp, err error){
	err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::SymmStorageSystem,Name::" + sid + "/relationships/Symm_SRPStoragePool", nil, &resp)
	return resp,err
}

type GetDeviceMaskingGroupsResp struct {
	Entries []struct{
		Content []struct {
			DeviceMaskingGroupName string `json:"i$ElementName"`
			InstanceName string `json:"i$InstanceID"`
		} `json:"content"`
	} `json:"entries"`
}

//Get a list of Device Masking Groups associated with SymmID
func (smis *SMIS) GetDeviceMaskingGroups(sid string) (resp *GetDeviceMaskingGroupsResp, err error){
	err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::Symm_StorageSystem,Name::" + sid + "/relationships/CIM_DeviceMaskingGroup", nil, &resp)
	return resp,err
}

type GetStorageVolumesResp struct {
	Entries []struct{
		Content []struct {
			DeviceName string `json:"i$DeviceID"`
		} `json:"content"`
	} `json:"entries"`
}

//Get a list of Storage Volumes on/associated with SymmID
func (smis *SMIS) GetStorageVolumes(sid string) (resp *GetStorageVolumesResp, err error){
	err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::Symm_StorageSystem,Name::" + sid + "/relationships/CIM_StorageVolume", nil, &resp)
	return resp,err
}

type PostDeviceMaskingGroupsContent struct {
	Content []struct {
		GroupName string `json: "GroupName"`
		Type int `json: "Type"`
		//Type 2 for Initator Group, Type 4 for Device Masking Group
		atType string `json: "@type"`
	} `json:"content"`
}

//Create a Device Masking Group (Same method call for Create an Initiator Group, but different Type in Content)
//func (smis *SMIS) PostDeviceMaskingGroups (req *PostVolumesReq, sid string) (resp *PostVolumesResp, err error){
//	err = smis.query("POST","/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::" + sid + "/action/CreateGroup", req, &resp)
//	return resp,err
//}

//Add goal struct later 
type PostVolumesContent struct {
	Content []struct {
		ElementType string `json: "ElementType"`
		Size string `json: "Size"`
		ElementName string `json: "ElementName"`
		EMCNumberOfDevices string `json: "EMCNumberOfDevices"`
		atType string `json: "@type"`
		//Add goal stuct later {@type, InstanceID}
	} `json:"content"`
}
