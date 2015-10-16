package apiv1

import (
	"errors"
	"os"
	"strconv"
	)

type GetStoragePoolsResp struct {
	Entries `json:"entries"`{
		Content []struct {
			StoragePoolName string `json:"i$ElementName"`
			InstanceName string `json:"i$InstanceID"`
		} `json:"content"`
	} `json:"entries"`
}

//Get a list of all Storage Pools associated with SymmID
func (smis *SMIS) GetStoragePools(sid string) (resp *GetStoragePoolsResp, err error){
	err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::SymmStorageSystem,Name::" + sid + "/relationships/CIM_StoragePool", nil, &resp)
	return resp,err
}

type GetDeviceMaskingGroupsResp struct {
	Entries `json:"entries"`{
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
	Entries `json:"entries"`{
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
		@Type string `json: "@type"`
	} `json:"content"`
}

//Create a Device Masking Group (Same method call for Create an Initiator Group, but different Type in Content)
func (smis *SMIS) PostDeviceMaskingGroups (req *PostVolumesReq, sid string) (resp *PostVolumesResp, err error){
	err = smis.query("POST","/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::" + sid + "/action/CreateGroup", req, &resp)
	return resp,err
}

type PostVolumesReq struct {

}

type PostVolumesContent struct {
	Content []struct {
		ElementType string `json: "ElementType"`
		Size string `json: "Size"`
		ElementName string `json: "ElementName"`
		EMCNumberOfDevices string `json: "EMCNumberOfDevices"`
		@Type string `json: "@type"`
		//Add goal stuct later {@type, InstanceID}
	} `json:"content"`
}

type PostVolumesResp struct {

}

//Create a Storage Volume
func (smis *SMIS) PostVolumes(req *PostVolumesReq, sid string) (resp *PostVolumesResp, err error){
	err = smis.query("POST","/ecom/edaa/root/emc/instances/Symm_StorageConfigurationService/CreationClassName::Symm_StorageConfigurationService,Name::EMCStorageConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::" + sid + "/action/CreateOrModifyElementFromStoragePool", req, &resp)
	return resp,err
}

