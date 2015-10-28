package apiv1

import ()

type GetStorageArraysResp struct {
	Entries []struct {
		Content struct {
			AtType                          string      `json:"@type"`
			I_Caption                      string      `json:"i$Caption"`
			I_CreationClassName            string      `json:"i$CreationClassName"`
			I_Dedicated                    []int       `json:"i$Dedicated"`
			I_Description                  string      `json:"i$Description"`
			I_EMCAutoMetaConfiguration     []int       `json:"i$EMCAutoMetaConfiguration"`
			I_EMCAutoMetaEnabled           bool        `json:"i$EMCAutoMetaEnabled"`
			I_EMCAutoMetaMemberSize        int         `json:"i$EMCAutoMetaMemberSize"`
			I_EMCAutoMetaMinSize           int         `json:"i$EMCAutoMetaMinSize"`
			I_EMCBSPElementType            int         `json:"i$EMCBSPElementType"`
			I_EMCBSPInstanceID             string      `json:"i$EMCBSPInstanceID"`
			I_EMCCacheHiWatermark          int         `json:"i$EMCCacheHiWatermark"`
			I_EMCCacheLoWatermark          int         `json:"i$EMCCacheLoWatermark"`
			I_EMCEnclosureCount            int         `json:"i$EMCEnclosureCount"`
			I_EMCLastSyncTime              string      `json:"i$EMCLastSyncTime"`
			I_EMCLastSyncType              []int       `json:"i$EMCLastSyncType"`
			I_EMCLastSyncTypeDesc          []string    `json:"i$EMCLastSyncTypeDesc"`
			I_EMCLicenseActivationType     []int       `json:"i$EMCLicenseActivationType"`
			I_EMCLicenseCapacityType       []int       `json:"i$EMCLicenseCapacityType"`
			I_EMCLicenseConversionFactor   []int       `json:"i$EMCLicenseConversionFactor"`
			I_EMCLicenseExpirationDate     []string    `json:"i$EMCLicenseExpirationDate"`
			I_EMCLicenseFeatureName        []string    `json:"i$EMCLicenseFeatureName"`
			I_EMCLicenseStatus             []int       `json:"i$EMCLicenseStatus"`
			I_EMCLocality                  int         `json:"i$EMCLocality"`
			I_EMCMaskingEnabled            bool        `json:"i$EMCMaskingEnabled"`
			I_EMCMemorySize                int         `json:"i$EMCMemorySize"`
			I_EMCNonSataActualCapacity     []int       `json:"i$EMCNonSataActualCapacity"`
			I_EMCNonSataLicenseCapacity    []int       `json:"i$EMCNonSataLicenseCapacity"`
			I_EMCNumberOfDisks             int         `json:"i$EMCNumberOfDisks"`
			I_EMCNumberOfUnconfiguredDisks int         `json:"i$EMCNumberOfUnconfiguredDisks"`
			I_EMCNumberofFrontEndPorts     int         `json:"i$EMCNumberofFrontEndPorts"`
			I_EMCNumberofSymDevs           int         `json:"i$EMCNumberofSymDevs"`
			I_EMCReadCacheSize             int         `json:"i$EMCReadCacheSize"`
			I_EMCSataActualCapacity        []int       `json:"i$EMCSataActualCapacity"`
			I_EMCSataLicenseCapacity       []int       `json:"i$EMCSataLicenseCapacity"`
			I_EMCUpTime                    string      `json:"i$EMCUpTime"`
			I_EMCWriteCacheSize            int         `json:"i$EMCWriteCacheSize"`
			I_EMCisDARE                    bool        `json:"i$EMCisDARE"`
			I_ElementName                  string      `json:"i$ElementName"`
			I_EnabledDefault               int         `json:"i$EnabledDefault"`
			I_EnabledState                 int         `json:"i$EnabledState"`
			I_HealthState                  int         `json:"i$HealthState"`
			I_IdentifyingDescriptions      []string    `json:"i$IdentifyingDescriptions"`
			I_Name                         string      `json:"i$Name"`
			I_NameFormat                   string      `json:"i$NameFormat"`
			I_OperationalStatus            []int       `json:"i$OperationalStatus"`
			I_OtherIdentifyingInfo         []string    `json:"i$OtherIdentifyingInfo"`
			I_PowerManagementCapabilities  []int       `json:"i$PowerManagementCapabilities"`
			I_PrimaryOwnerContact          interface{} `json:"i$PrimaryOwnerContact"`
			I_PrimaryOwnerName             interface{} `json:"i$PrimaryOwnerName"`
			I_RequestedState               int         `json:"i$RequestedState"`
			I_ResetCapability              int         `json:"i$ResetCapability"`
			I_StatusDescriptions           []string    `json:"i$StatusDescriptions"`
			I_TransitioningToState         int         `json:"i$TransitioningToState"`
			Links                          []struct {
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

func (smis *SMIS) GetStorageArrays() (resp *GetStorageArraysResp, err error){
    err = smis.query("GET","/ecom/edaa/root/emc/types/Symm_StorageSystem/instances", nil, &resp)
    return resp,err
}

type GetStoragePoolsResp  struct {
    Entries []struct {
        Content struct {
            AtType                         string        `json:"@type"`
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
            AtType                     string   `json:"@type"`
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
			AtType                             string      `json:"@type"`
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
    err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::SYMMETRIX-+-" + sid + "/relationships/SE_DeviceMaskingGroup", nil, &resp)
    return resp,err
}

type GetPortGroupsResp struct {
	Entries []struct {
		Content struct {
			AtType                           string `json:"@type"`
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
    err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::SYMMETRIX-+-" + sid + "/relationships/SE_TargetMaskingGroup", nil, &resp)
    return resp,err
}

type GetHostGroupsResp struct {
	Entries []struct {
		Content struct {
			AtType                           string `json:"@type"`
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
    err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::SYMMETRIX-+-" + sid + "/relationships/SE_InitiatorMaskingGroup", nil, &resp)
    return resp,err
}


type GetStorageVolumesResp struct {
    Entries []struct {
        Content struct {
            AtType                               string          `json:"@type"`
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


type PostVolumesReq struct {
	PostVolumesRequestContent PostVolumesReqContent `json:"content"`
}

type PostVolumesReqContent struct {
	AtType                 string `json:"@type"`
    ElementName            string `json:"ElementName"`
    ElementType            string `json:"ElementType"`
	EMCNumberOfDevices     string `json:"EMCNumberOfDevices"`
	Size                   string `json:"Size"`
}

type PostVolumesResp struct {
	Entries []struct {
		Content struct {
			AtType       string `json:"@type"`
			I_Parameters struct {
				I_Job struct {
					AtType        string `json:"@type"`
					E0_InstanceID string `json:"e0$InstanceID"`
					Xmlns_e0      string `json:"xmlns$e0"`
				} `json:"i$Job"`
				I_Size int `json:"i$Size"`
			} `json:"i$parameters"`
			I_ReturnValue int    `json:"i$returnValue"`
			Xmlns_i       string `json:"xmlns$i"`
		} `json:"content"`
		Content_type string `json:"content-type"`
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


//Create a Storage Volume
func (smis *SMIS) PostVolumes(req *PostVolumesReq, sid string) (resp *PostVolumesResp, err error){
    err = smis.query("POST","/ecom/edaa/root/emc/instances/Symm_StorageConfigurationService/CreationClassName::Symm_StorageConfigurationService,Name::EMCStorageConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::" + sid + "/action/CreateOrModifyElementFromStoragePool", req, &resp)
    return resp,err
}

//////////////////////////////////////
//   REQUEST Structs used for any   //
//   group creation on the VMAX3.   //
//  				    //
//    Storage Group (SG) - Type 4   //
//     Port Group (PG) - Type 3     //
//   Initiator Group (IG) - Type 2  //
//////////////////////////////////////

type PostGroupReq struct {
        PostGroupRequestContent PostGroupReqContent `json:"content"`
}

type PostGroupReqContent struct {
        AtType           string `json:"@type"`
        SG_Name		     string `json:"GroupName"`
        Type             string `json:"Type"`
}

////////////////////////////////////////////////////////////
//           RESPONSE Structs used for any                //
//           group creation on the VMAX3.                 //
//                                                        //
//   Storage Group (SG) - Type SE_DeviceMaskingGroup      //
//      Port Group (PG) - Type SE_TargetMaskingGroup      //
//  Initiator Group (IG) - Type SE_InitiatorMaskingGroup  //
////////////////////////////////////////////////////////////


type PostGroupResp struct {
	Entries []struct {
		Content struct {
			_type        string `json:"@type"`
			I_parameters struct {
				I_MaskingGroup struct {
					_type         string `json:"@type"`
					E0_InstanceID string `json:"e0$InstanceID"`
					Xmlns_e0      string `json:"xmlns$e0"`
				} `json:"i$MaskingGroup"`
			} `json:"i$parameters"`
			I_returnValue int    `json:"i$returnValue"`
			Xmlns_i       string `json:"xmlns$i"`
		} `json:"content"`
		Content_type string `json:"content-type"`
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


///////////////////////////////////////////////////////////////
//                  Create an Array Group                    //
// Type Depends on Type field specified in requesting struct //
///////////////////////////////////////////////////////////////

func (smis *SMIS) PostCreateGroup(req *PostGroupReq, sid string) (resp *PostGroupResp, err error){
	err = smis.query("POST","/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::SYMMETRIX-+-" + sid + "/action/CreateGroup", req, &resp)
	return resp, err
}

/*
type PostVolumesToSGReq struct {
    PostVolumesToSGReqContent `json:"content"` 
}

type PostVolumesToSGReqContent struct {
    MaskingGroup `json:"MaskingGroup"`
    Members `json:"Members"`
    //@type = http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerconfigurationService
    AtType string `json:"@type"`
}

type PostVolumesToSGReqContentMaskingGroup struct {
    //@type = http://schemas.emc.com/ecom/edaa/root/emc/SE_DeviceMaskingGroup
    //InstanceID = SYMMETRIX-+-sid-+-SG_name  
    AtType          string  `json: "@type"`
    InstanceName    string  `json: "InstanceID"`
} 

type PostVolumesToSGReqContentMembers []struct { 
        //@type = http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageVolume
        //CreationClassName = Symm_StorageVolume
        //SystemCreationClassName = Symm_StorageSystem
        //SystemName = SYMMETRIX-+-sid
        AtType                  string      `json: "@type"`
        CreationClassID         string      `json: "CreationClassName"`
        DeviceName              string      `json: "DeviceID"`
        SystemCreationClassID   string      `json: "SystemCreationClassName"`               
        SystemID                string      `json: "SystemName"`
} 

type PostVolumesToSGResp struct {
    Entries []struct {
        Content []struct {
            I_Parameters []struct{
                I_Job []struct {
                    AtType        string `json:"@type"`
                    E0_InstanceID string `json:"e0$InstanceID"`
                    Xmlns_e0      string `json:"xmlns$e0"`
                } `json:"i$Job"`
                I_Size int `json:"i$Size"`
            } `json:"i$parameters"`
            I_ReturnValue int    `json:"i$returnValue"`
            Xmlns_i       string `json:"xmlns$i"`
        } `json:"content"`
        Content_type string `json:"content-type"`
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

//Add Volume to a Device Masking/Storage Group
func (smis *SMIS) PostVolumesToSG(req *PostVolumesToSGReq, sid string) (resp *PostVolumesToSGResp, err error){
    err = smis.query("POST","http://10.246.125.104:5888/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::" + sid + "/action/AddMembers", req, &resp)
    return resp,err
} 

type GetSLOsResp struct {
    Entries []struct {
        Content struct {
            AtType                              string          `json:"@type"`
            I_InstanceID                        string          `json:"i$InstanceID"`
            I_Changeable                        bool            `json:"i$Changeable"`
            I_ChangeableType                    int             `json:"i$ChangeableType"`
            I_CompressedElement                 bool            `json:"i$CompressedElement"`
            I_CompressionRate                   int             `json:"i$CompressionRate"`
            I_DataRedundancyGoal                int             `json:"i$DataRedundancyGoal"`
            I_DataRedundancyMax                 int             `json:"i$DataRedundancyMax"`
            I_DataRedundancyMin                 int             `json:"i$DataRedundancyMin"`
            I_DeltaReservationGoal              int             `json:"i$DeltaReservationGoal"`
            I_DeltaReservationMax               int             `json:"i$DeltaReservationMax"`
            I_DeltaReservationMin               int             `json:"i$DeltaReservationMin"`
            I_ElementName                       string          `json:"i$ElementName"`
            I_EMCApproxAverageResponseTime      float32         `json:"i$EMCApproxAverageResponseTime"`
            I_EMCDeduplicationRate              int             `json:"i$EMCDeduplicationRate"`
            I_EMCEnableDIF                      int             `json:"i$EMCEnableDIF"`
            I_EMCEnableEFDCache                 int             `json:"i$EMCEnableEFDCache"`
            I_EMCFastSetting                    string          `json:"i$EMCFastSetting"`
            I_EMCParticipateInPowerSavings      int             `json:"i$EMCParticipateInPowerSavings"`
            I_EMCPoolCompressionState           int             `json:"i$EMCPoolCompressionState"`
            I_EMCPottedSetting                  bool            `json:"i$EMCPottedSetting"`
            I_EMCRaidGroupLUN                   bool            `json:"i$EMCRaidGroupLUN"`
            I_EMCRaidLevel                      string          `json:"i$EMCRaidLevel"`
            I_EMCSLO                            string          `json:"i$EMCSLO"`
            I_EMCSLOBaseName                    string          `json:"i$EMCSLOBaseName"`
            I_EMCSLOdescription                 string          `json:"i$EMCSLOdescription"`
            I_EMCSRP                            string          `json:"i$EMCSRP"`
            I_EMCStorageSettingType             int             `json:"i$EMCStorageSettingType"`
            I_EMCUniqueID                       string          `json:"i$EMCUniqueID"`
            I_EMCWorkload                       string          `json:"i$EMCWorkload"`
            I_ExtentStripeLength                int             `json:"i$ExtentStripeLength"`
            I_ExtentStripeLengthMax             int             `json:"i$ExtentStripeLengthMax"`
            I_ExtentStripeLengthMin             int             `json:"i$ExtentStripeLengthMin"`
            I_InitialStorageTieringSelection    int             `json:"i$InitialStorageTieringSelection"`
            I_InitialStorageTierMethodology     int             `json:"i$InitialStorageTierMethodology"`
            I_InitialSynchronization            int             `json:"i$InitialSynchronization"`
            I_NoSinglePointOfFailure            bool            `json:"i$NoSinglePointOfFailure"`
            I_PackageRedundancyGoal             int             `json:"i$PackageRedundancyGoal"`
            I_PackageRedundancyMax              int             `json:"i$PackageRedundancyMax"`
            I_PackageRedundancyMin              int             `json:"i$PackageRedundancyMin"`
            I_SpaceLimit                        int             `json:"i$SpaceLimit"`
            I_StorageExtentInitialUsage         int             `json:"i$StorageExtentInitialUsage"`
            I_StoragePoolInitialUsage           int             `json:"i$StoragePoolInitialUsage"`
            I_ThinProvisionedPoolType           int             `json:"i$ThinProvisionedPoolType"`
            I_UseReplicationBuffer              int             `json:"i$UseReplicationBuffer"`
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

//Get a list of SLOs on/associated with SRP and SymmId
func (smis *SMIS) GetSLOs (SRP string, sid string) (resp *GetSLOsResp, err error){
    err = smis.query("GET",`/ecom/edaa/root/emc/types/Symm_StoragePoolSetting/instances?filter=EMCSRP lk eq "` + SRP + `" and InstanceID lk "` + sid + `"`, nil, &resp)
    return resp,err
}*/
