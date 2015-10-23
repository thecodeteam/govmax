package apiv1

import ()

type GetStoragePoolsResp  struct {
    Entries []struct {
        Content struct {
            _type                         string        `json:"@type"`
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
            Links                         []struct {
                Href string `json:"href"`
                Rel  string `json:"rel"`
            } `json:"links"`
            Xmlns_i string `json:"xmlns$i"`
        } `json:"content"`
        Content_type string `json:"content-type"`
        Gd_etag      string `json:"gd$etag"`
        Links        []struct {
            Href string `json:"href"`
            Rel  string `json:"rel"`
        } `json:"links"`
        Updated string `json:"updated"`
    } `json:"entries"`
    ID    string `json:"id"`
    Links []struct {
        Href string `json:"href"`
        Rel  string `json:"rel"`
    } `json:"links"`
    Updated  string `json:"updated"`
    Xmlns_gd string `json:"xmlns$gd"`
}

//Get a list of all Storage Pools associated with SymmID
func (smis *SMIS) GetStoragePools(sid string) (resp *GetStoragePoolsResp, err error){
    err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::SymmStorageSystem,Name::" + sid + "/relationships/Symm_SRPStoragePool", nil, &resp)
    return resp,err
}

type GetDeviceMaskingViewsResp struct {
    Entries []struct {
        Content struct {
            _type                     string   `json:"@type"`
            I_Caption                 string   `json:"i$Caption"`
            I_CreationClassName       string   `json:"i$CreationClassName"`
            I_Description             string   `json:"i$Description"`
            I_DeviceID                string   `json:"i$DeviceID"`
            I_EMCAdapterRole          string   `json:"i$EMCAdapterRole"`
            I_ElementName             string   `json:"i$ElementName"`
            I_EnabledDefault          int      `json:"i$EnabledDefault"`
            I_EnabledState            int      `json:"i$EnabledState"`
            I_Generation              int      `json:"i$Generation"`
            I_HealthState             int      `json:"i$HealthState"`
            I_MaxUnitsControlled      int      `json:"i$MaxUnitsControlled"`
            I_Name                    string   `json:"i$Name"`
            I_NameFormat              int      `json:"i$NameFormat"`
            I_OperationalStatus       []int    `json:"i$OperationalStatus"`
            I_OtherNameFormat         string   `json:"i$OtherNameFormat"`
            I_RequestedState          int      `json:"i$RequestedState"`
            I_StatusDescriptions      []string `json:"i$StatusDescriptions"`
            I_SystemCreationClassName string   `json:"i$SystemCreationClassName"`
            I_SystemName              string   `json:"i$SystemName"`
            I_TransitioningToState    int      `json:"i$TransitioningToState"`
            Links                     []struct {
                Href string `json:"href"`
                Rel  string `json:"rel"`
            } `json:"links"`
            Xmlns_i string `json:"xmlns$i"`
        } `json:"content"`
        Content_type string `json:"content-type"`
        Gd_etag      string `json:"gd$etag"`
        Links        []struct {
            Href string `json:"href"`
            Rel  string `json:"rel"`
        } `json:"links"`
        Updated string `json:"updated"`
    } `json:"entries"`
    ID    string `json:"id"`
    Links []struct {
        Href string `json:"href"`
        Rel  string `json:"rel"`
    } `json:"links"`
    Updated  string `json:"updated"`
    Xmlns_gd string `json:"xmlns$gd"`
}

//Get a list of Device Masking Views associated with SymmID
func (smis *SMIS) GetDeviceMaskingViews(sid string) (resp *GetDeviceMaskingViewsResp, err error){
    err = smis.query("GET","/ecom/edaa/root/emc/types/Symm_LunMaskingView/instances", nil, &resp)
    return resp,err
}

type GetStorageGroupsResp struct {
	Entries []struct {
		Content struct {
			_type                             string      `json:"@type"`
			I_Caption                         string      `json:"i$Caption"`
			I_DeleteOnEmpty                   bool        `json:"i$DeleteOnEmpty"`
			I_DeleteWhenBecomesUnassociated   bool        `json:"i$DeleteWhenBecomesUnassociated"`
			I_EMCAssociatedToView             bool        `json:"i$EMCAssociatedToView"`
			I_EMCFastSetting                  string      `json:"i$EMCFastSetting"`
			I_EMCMaxIOBandwidthInherited      bool        `json:"i$EMCMaxIOBandwidthInherited"`
			I_EMCMaxIODynamicDistributionType int         `json:"i$EMCMaxIODynamicDistributionType"`
			I_EMCMaximumBandwidth             int         `json:"i$EMCMaximumBandwidth"`
			I_EMCMaximumIO                    int         `json:"i$EMCMaximumIO"`
			I_EMCRecoverPointEnabled          bool        `json:"i$EMCRecoverPointEnabled"`
			I_EMCSLO                          string      `json:"i$EMCSLO"`
			I_EMCSLOBaseName                  string      `json:"i$EMCSLOBaseName"`
			I_EMCSRP                          string      `json:"i$EMCSRP"`
			I_EMCWorkload                     interface{} `json:"i$EMCWorkload"`
			I_ElementName                     string      `json:"i$ElementName"`
			I_Generation                      int         `json:"i$Generation"`
			I_InstanceID                      string      `json:"i$InstanceID"`
			Links                             []struct {
				Href string `json:"href"`
				Rel  string `json:"rel"`
			} `json:"links"`
			Xmlns_i string `json:"xmlns$i"`
		} `json:"content"`
		Content_type string `json:"content-type"`
		Gd_etag      string `json:"gd$etag"`
		Links        []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"links"`
		Updated string `json:"updated"`
	} `json:"entries"`
	ID    string `json:"id"`
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Updated  string `json:"updated"`
	Xmlns_gd string `json:"xmlns$gd"`
}



//Get a list of Device Masking/Storage Groups associated with SymmID 
func (smis *SMIS) GetStorageGroups(sid string) (resp *GetStorageGroupsResp, err error){
    err = smis.query("GET","/ecom/edaa/root/emc/types/SE_DeviceMaskingGroup/instances", nil, &resp)
    return resp,err
}

type GetPortGroupsResp struct {
	Entries []struct {
		Content struct {
			_type                           string `json:"@type"`
			I_Caption                       string `json:"i$Caption"`
			I_DeleteOnEmpty                 bool   `json:"i$DeleteOnEmpty"`
			I_DeleteWhenBecomesUnassociated bool   `json:"i$DeleteWhenBecomesUnassociated"`
			I_ElementName                   string `json:"i$ElementName"`
			I_Generation                    int    `json:"i$Generation"`
			I_InstanceID                    string `json:"i$InstanceID"`
			Links                           []struct {
				Href string `json:"href"`
				Rel  string `json:"rel"`
			} `json:"links"`
			Xmlns_i string `json:"xmlns$i"`
		} `json:"content"`
		Content_type string `json:"content-type"`
		Gd_etag      string `json:"gd$etag"`
		Links        []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"links"`
		Updated string `json:"updated"`
	} `json:"entries"`
	ID    string `json:"id"`
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Updated  string `json:"updated"`
	Xmlns_gd string `json:"xmlns$gd"`
}


//Get a list of Target/Port Groups associated with SymmID 
func (smis *SMIS) GetPortGroups(sid string) (resp *GetPortGroupsResp, err error){
    err = smis.query("GET","/ecom/edaa/root/emc/types/SE_TargetMaskingGroup/instances", nil, &resp)
    return resp,err
}

type GetHostGroupsResp struct {
	Entries []struct {
		Content struct {
			_type                           string `json:"@type"`
			I_Caption                       string `json:"i$Caption"`
			I_ConsistentLogicalUnitNumber   bool   `json:"i$ConsistentLogicalUnitNumber"`
			I_DeleteOnEmpty                 bool   `json:"i$DeleteOnEmpty"`
			I_DeleteWhenBecomesUnassociated bool   `json:"i$DeleteWhenBecomesUnassociated"`
			I_ElementName                   string `json:"i$ElementName"`
			I_Generation                    int    `json:"i$Generation"`
			I_InstanceID                    string `json:"i$InstanceID"`
			Links                           []struct {
				Href string `json:"href"`
				Rel  string `json:"rel"`
			} `json:"links"`
			Xmlns_i string `json:"xmlns$i"`
		} `json:"content"`
		Content_type string `json:"content-type"`
		Gd_etag      string `json:"gd$etag"`
		Links        []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"links"`
		Updated string `json:"updated"`
	} `json:"entries"`
	ID    string `json:"id"`
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Updated  string `json:"updated"`
	Xmlns_gd string `json:"xmlns$gd"`
}


//Get a list of Initiator Groups associated with SymmID 
func (smis *SMIS) GetHostGroups(sid string) (resp *GetHostGroupsResp, err error){
    err = smis.query("GET","/ecom/edaa/root/emc/types/SE_InitiatorMaskingGroup/instances", nil, &resp)
    return resp,err
}


type GetStorageVolumesResp struct {
    Entries []struct {
        Content struct {
            _type                               string          `json:"@type"`
            I_Access                            int             `json:"i$Access"`
            I_BlockSize                         int             `json:"i$BlockSize"`
            I_Caption                           string          `json:"i$Caption"`
            I_ClientSettableUsage               []interface{}   `json:"i$ClientSettableUsage"`
            I_CompressionRate                   int             `json:"i$CompressionRate"`
            I_CompressionState                  int             `json:"i$CompressionState"`
            I_ConsumableBlocks                  int             `json:"i$ConsumableBlocks"`
            I_CreationClassName                 string          `json:"i$CreationClassName"`
            I_DataRedundancy                    int             `json:"i$DataRedundancy"`
            I_DeltaReservation                  int             `json:"i$DeltaReservation"`
            I_DeviceID                          string          `json:"i$DeviceID"`
            I_ElementName                       string          `json:"i$ElementName"`
            I_EMCBSPElementType                 int             `json:"i$EMCBSPElementType"`
            I_EMCBSPInstanceID                  string          `json:"i$EMCBSPInstanceID"`
            I_EMCConsumableBlocksWritten        int             `json:"i$EMCConsumableBlocksWritten"`
            I_EMCDataFormat                     string          `json:"i$EMCDataFormat"`
            I_EMCDIFEnabled                     bool            `json:"i$EMCDIFEnabled"`
            I_EMCFullyAllocated                 bool            `json:"i$EMCFullyAllocated"`
            I_EMCIsBound                        bool            `json:"i$EMCIsBound"`
            I_EMCIsComposite                    bool            `json:"i$EMCIsComposite"`
            I_EMCIsDeDuplicated                 bool            `json:"i$EMCIsDeDuplicated"`
            I_EMCIsMapped                       bool            `json:"i$EMCIsMapped"`
            I_EMCRaidLevel                      string          `json:"i$EMCRaidLevel"`
            I_EMCRecoverPointEnabled            bool            `json:"i$EMCRecoverPointEnabled"`
            I_EMCSpaceConsumed                  int             `json:"i$EMCSpaceConsumed"`
            I_EMCVolumeAttributes               []int           `json:"i$EMCVolumeAttributes"`
            I_EMCVolumeAttributes2              []int           `json:"i$EMCVolumeAttributes2"`
            I_EMCVolumeAttributesDescription    []string        `json:"i$EMCVolumeAttributesDescription"`
            I_EMCVolumeAttributesDescription2   []string        `json:"i$EMCVolumeAttributesDescription2"`
            I_EMCVolumeLabel                    string          `json:"i$EMCVolumeLabel"`
            I_EMCWWN                            string          `json:"i$EMCWWN"`
            I_EnabledState                      int             `json:"i$EnabledState"`
            I_ExtentDiscriminator               []string        `json:"i$ExtentDiscriminator"`
            I_ExtentStatus                      []int           `json:"i$ExtentStatus"`
            I_HealthState                       int             `json:"i$HealthState"`
            I_IdentifyingDescriptions           []string        `json:"i$IdentifyingDescriptions"`
            I_IsBasedOnUnderlyingRedundancy     bool            `json:"i$IsBasedOnUnderlyingRedundancy"`
            I_IsComposite                       bool            `json:"i$IsComposite"`
            I_IsCompressed                      bool            `json:"i$IsCompressed"`
            I_Name                              string          `json:"i$Name"`
            I_NameFormat                        int             `json:"i$NameFormat"`
            I_NameNamespace                     int             `json:"i$NameNamespace"`
            I_NoSinglePointOfFailure            bool            `json:"i$NoSinglePointOfFailure"`
            I_NumberOfBlocks                    int             `json:"i$NumberOfBlocks"`
            I_OperationalStatus                 []int           `json:"i$OperationalStatus"`
            I_OtherIdentifyingInfo              []string        `json:"i$OtherIdentifyingInfo"`
            I_OtherNameFormat                   string          `json:"i$OtherNameFormat"`
            I_PackageRedundancy                 int             `json:"i$PackageRedundancy"`
            I_Primordial                        bool            `json:"i$Primordial"`
            I_Purpose                           string          `json:"i$Purpose"`
            I_RequestedState                    int             `json:"i$RequestedState"`
            I_SequentialAccess                  bool            `json:"i$SequentialAccess"`
            I_SystemCreationClassName           string          `json:"i$SystemCreationClassName"`
            I_SystemName                        string          `json:"i$SystemName"`
            I_StatusDescriptions                []string        `json:"i$StatusDescriptions"`
            I_StatusInfo                        int             `json:"i$StatusInfo"`
            I_StorageTieringSelection           int             `json:"i$StorageTieringSelection"`
            I_ThinlyProvisioned                 bool            `json:"i$ThinlyProvisioned"`
            I_TransitioningToState              int             `json:"i$TransitioningToState"`
            I_Usage                             int             `json:"i$Usage"`
            Links []struct {
                Href string `json:"href"`
                Rel  string `json:"rel"`
            } `json:"links"`
            Xmlns_i string `json:"xmlns$i"`
        } `json:"content"`
        Content_type    string      `json:"content-type"`
        Gd_etag         string      `json:"gd$etag"`
        Links           []struct {
            Href string `json:"href"`
            Rel  string `json:"rel"`
        } `json:"links"`
        Updated         string      `json:"updated"`
    } `json:"entries"`
    ID    string `json:"id"`
    Links []struct {
        Href string `json:"href"`
        Rel  string `json:"rel"`
    } `json:"links"`
    Updated  string `json:"updated"`
    Xmlns_gd string `json:"xmlns$gd"`
}

//Get a list of Storage Volumes on/associated with SymmID
func (smis *SMIS) GetStorageVolumes(sid string) (resp *GetStorageVolumesResp, err error){
    err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::Symm_StorageSystem,Name::" + sid + "/relationships/CIM_StorageVolume", nil, &resp)
    return resp,err
}

/*
type PostMaskingGroupsReq struct {
    Content []struct {
        //@type = http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerConfigurationService
        _type       string      `json: "@type"`
        GroupName   string      `json: "GroupName"`
        Type        int         `json: "Type"` 
        //Type 2 = Initiator Masking Group
        //Type 3 = Target/Port Masking Group
        //Type 4 = Device Masking/Storage Group         
    } `json:"content"`
}

type PostMaskingGroupsResp struct {
    Entries []struct {
        Content []struct {
            I_Parameters []struct {
                I_MaskingGroup []struct {
                    I_InstanceID string `json: "e0$InstanceID"`
                } `json:"i$MaskingGroup"`
            } `json:"i$parameters"`
        } `json:"content"`
    } `json:"entries"`
}

//Create a Masking Group (Same method call for Initator, Target/Port, and Device/Storage)
func (smis *SMIS) PostMaskingGroupsReq (req *PostMaskingGroupsReq, sid string) (resp *PostVolumesResp, err error){
    err = smis.query("POST","/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::" + sid + "/action/CreateGroup", req, &resp)
    return resp,err
}

type PostVolumesReq struct {
    Content []struct {
        //@type = http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageConfigurationService
        //Add goal stuct later {@type, InstanceID}
        _type               string      `json: "@type"`
        ElementName         string      `json: "ElementName"`
        ElementType         string      `json: "ElementType"`
        EMCNumberOfDevices  string      `json: "EMCNumberOfDevices"`
        Size                string      `json: "Size"`
    } `json:"content"`

}
type PostVolumesResp struct {
    Entries []struct {
        Content []struct {
            I_Parameters []struct {
                I_Job []struct {
                    I_InstanceID string `json: "e0$InstanceID"`
                } `json:"i$Job"`
            } `json:"i$parameters"`
        } `json:"content"`
    } `json:"entries"`
}

//Create a Storage Volume
func (smis *SMIS) PostVolumes(req *PostVolumesReq, sid string) (resp *PostVolumesResp, err error){
    err = smis.query("POST","/ecom/edaa/root/emc/instances/Symm_StorageConfigurationService/CreationClassName::Symm_StorageConfigurationService,Name::EMCStorageConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::" + sid + "/action/CreateOrModifyElementFromStoragePool", req, &resp)
    return resp,err
}

type PostVolumesToSG struct {
    Content []struct {
        MaskingGroup [] struct {
            //@type = http://schemas.emc.com/ecom/edaa/root/emc/SE_DeviceMaskingGroup
            //InstanceID = SYMMETRIX-+-sid-+-SG_name  
            _type           string  `json: "@type"`
            InstanceName    string  `json: "InstanceID"`
        } `json:"MaskingGroup"`
        Members []struct { 
            //@type = http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageVolume
            //CreationClassName = Symm_StorageVolume
            //SystemCreationClassName = Symm_StorageSystem
            //SystemName = SYMMETRIX-+-sid
            _type                   string      `json: "@type"`
            CreationClassID         string      `json: "CreationClassName"`
            DeviceName              string      `json: "DeviceID"`
            SystemCreationClassID   string      `json: "SystemCreationClassName"`               
            SystemID                string      `json: "SystemName"`
        } `json:"Members"`
        //@type = http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerconfigurationService
        _type string `json:"@type"`
    } `json:"content"`
}

type PostVolumesToSGResp struct {
    Entries []struct {
        Content []struct {
            I_Parameters []struct{
                I_Job []struct {
                    I_InstanceID string `json: "e0$InstanceID"`
                } `json:"i$Job"`
            } `json:"i$parameters"`
        } `json:"content"`
    } `json:"entries"`
}

//Add Volume to a Device Masking/Storage Group
func (smis *SMIS) PostVolumesToSG(req *PostVolumesReq, sid string) (resp *PostVolumesResp, err error){
    err = smis.query("POST","http://10.246.125.104:5888/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::" + sid + "/action/AddMembers", req, &resp)
    return resp,err
} 
*/
