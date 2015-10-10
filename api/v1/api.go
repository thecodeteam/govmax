package apiv1

import (
	"errors"
	"os"
	"strconv"
)

type GetStoragePoolsResp struct {
	Entries GetStoragePoolsResp_entries `json:"entries"`
}

type GetStoragePoolsResp_entries struct {
	Content GetStoragePoolsResp_entries_content `json:"content"`
}

type GetStoragePoolsResp_entries_content struct {
	StoragePoolName string `json:"i$ElementName"`
	InstanceName string `json:"i$InstanceID"`
}

func (smis *SMIS) GetStoragePools(sid string) (resp *GetStoragePoolsResp, err error){
	err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::SymmStorageSystem,Name::" + sid + "/relationships/CIM_StoragePool", nil, &resp)
	return resp,err
}

type PostVolumesReq struct {


}

type PostVolumesResp struct {

}

func (smis *SMIS) PostVolumes(req *PostVolumesReq, sid string) (resp *PostVolumesResp, err error){
	err = smis.query("POST","/ecom/edaa/root/emc/instances/Symm_StorageConfigurationService/CreationClassName::Symm_StorageConfigurationService,Name::EMCStorageConfigurationService,SysytemCreateionClassName::Symm_StorageSystem,SystemName::" + sid + "/action/CreateOrModifyElementFromStoragePool/action/CreateOrModifyElementFromStoragePool",
