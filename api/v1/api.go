package apiv1

import (
	"errors"
	"os"
	"strconv"
	)

type GetStoragePoolsResp_entries struct {
	Entries `json:"entries"`
	Content []struct {
		StoragePoolName string `json:"i$ElementName"`
		InstanceName string `json:"i$InstanceID"`
		}
	}
//Get a list of all Storage Pools associated with SymmID
func (smis *SMIS) GetStoragePools(sid string) (resp *GetStoragePoolsResp, err error){
	err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::SymmStorageSystem,Name::" + sid + "/relationships/CIM_StoragePool", nil, &resp)
	return resp,err
}

//Get a list of Storage Volumes to/associated with a Device Masking Group
func (smis *SMIS) GetDeviceMaskingGroups(sid string) (resp *GetDeviceMaskingGroupsResp, err error){
	err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::Symm_StorageSystem,Name::" + sid + "/relationships/CIM_DeviceMaskingGroup", nil, &resp)
	return resp,err
}

//Get a list of Storage Volumes on/associated with SymmID
func (smis *SMIS) GetStorageVolumes(sid string) (resp *GetStorageVolumesResp, err error){
	err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::Symm_StorageSystem,Name::" + sid + "/relationships/CIM_StorageVolume", nil, &resp)
	return resp,err
}

type PostVolumesReq struct {

}

type PostVolumesContent struct {

	ElementType string `json: "ElementType"`
	Size string `json: "Size"`
	ElementName string `json: "ElementName"`
	EMCNumberOfDevices string `json: "EMCNumberOfDevices"`
	Type string `json: "@type"`
}
//Add goal struct later 

type PostVolumesResp struct {

}

//Create a Storage Volume
func (smis *SMIS) PostVolumes(req *PostVolumesReq, sid string) (resp *PostVolumesResp, err error){
	err = smis.query("POST","/ecom/edaa/root/emc/instances/Symm_StorageConfigurationService/CreationClassName::Symm_StorageConfigurationService,Name::EMCStorageConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::" + sid + "/action/CreateOrModifyElementFromStoragePool", req, &resp)
	return resp,err
}

