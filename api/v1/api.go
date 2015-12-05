package apiv1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////
//             RESPONSE Struct used for                   //
//        list of storage arrays on the VMAX3.            //
////////////////////////////////////////////////////////////

type GetStorageArraysResp struct {
	Entries []struct {
		Content struct {
			AtType                         string      `json:"@type"`
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

///////////////////////////////////////////////////////////////
//            GET a list of Storage Arrays                   //
///////////////////////////////////////////////////////////////

func (smis *SMIS) GetStorageArrays() (resp *GetStorageArraysResp, err error) {
	err = smis.query("GET", "/ecom/edaa/root/emc/types/Symm_StorageSystem/instances", nil, &resp)
	return resp, err
}

////////////////////////////////////////////////////////////
//             RESPONSE Struct used for                   //
//        list of storage pools on the VMAX3.             //
////////////////////////////////////////////////////////////

type GetStoragePoolsResp struct {
	Entries []struct {
		Content struct {
			AtType                        string        `json:"@type"`
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

///////////////////////////////////////////////////////////////
//            GET a list of Storage Pools                    //
///////////////////////////////////////////////////////////////

func (smis *SMIS) GetStoragePools(sid string) (resp *GetStoragePoolsResp, err error) {
	err = smis.query("GET", "/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::SymmStorageSystem,Name::"+sid+"/relationships/Symm_SRPStoragePool", nil, &resp)
	return resp, err
}

////////////////////////////////////////////////////////////
//             RESPONSE Struct used for                   //
//        list of masking views on the VMAX3.             //
////////////////////////////////////////////////////////////

type GetMaskingViewsResp struct {
	Entries []struct {
		Content struct {
			AtType                    string   `json:"@type"`
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

///////////////////////////////////////////////////////////////
//            GET a list of Masking Views                    //
///////////////////////////////////////////////////////////////

func (smis *SMIS) GetMaskingViews(sid string) (resp *GetMaskingViewsResp, err error) {
	err = smis.query("GET", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/relationships/Symm_LunMaskingView", nil, &resp)
	return resp, err
}

////////////////////////////////////////////////////////////
//             RESPONSE Struct used for                   //
//        list of storage groups on the VMAX3.            //
////////////////////////////////////////////////////////////

type GetStorageGroupsResp struct {
	Entries []struct {
		Content struct {
			AtType                            string      `json:"@type"`
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

///////////////////////////////////////////////////////////////
//         GET a list of Storage (Device) Groups             //
///////////////////////////////////////////////////////////////

func (smis *SMIS) GetStorageGroups(sid string) (resp *GetStorageGroupsResp, err error) {
	err = smis.query("GET", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::SYMMETRIX-+-"+sid+"/relationships/SE_DeviceMaskingGroup", nil, &resp)
	return resp, err
}

////////////////////////////////////////////////////////////
//             RESPONSE Struct used for                   //
//         list of port groups on the VMAX3.              //
////////////////////////////////////////////////////////////

type GetPortGroupsResp struct {
	Entries []struct {
		Content struct {
			AtType                          string `json:"@type"`
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

///////////////////////////////////////////////////////////////
//         GET a list of Port (Target) Groups                //
///////////////////////////////////////////////////////////////

func (smis *SMIS) GetPortGroups(sid string) (resp *GetPortGroupsResp, err error) {
	err = smis.query("GET", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::SYMMETRIX-+-"+sid+"/relationships/SE_TargetMaskingGroup", nil, &resp)
	return resp, err
}

////////////////////////////////////////////////////////////
//             RESPONSE Struct used for                   //
//         list of host groups on the VMAX3.              //
////////////////////////////////////////////////////////////

type GetHostGroupsResp struct {
	Entries []struct {
		Content struct {
			AtType                          string `json:"@type"`
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

///////////////////////////////////////////////////////////////
//         GET a list of Host (Initiator) Groups             //
///////////////////////////////////////////////////////////////

func (smis *SMIS) GetHostGroups(sid string) (resp *GetHostGroupsResp, err error) {
	err = smis.query("GET", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::SYMMETRIX-+-"+sid+"/relationships/SE_InitiatorMaskingGroup", nil, &resp)
	return resp, err
}

////////////////////////////////////////////////////////////
//            RESPONSE Struct used for                    //
//        list of storage volumes on the VMAX3.           //
////////////////////////////////////////////////////////////

type GetVolumesResp struct {
	Entries []struct {
		Content struct {
			AtType                            string        `json:"@type"`
			I_Access                          int           `json:"i$Access"`
			I_BlockSize                       int           `json:"i$BlockSize"`
			I_Caption                         string        `json:"i$Caption"`
			I_ClientSettableUsage             []interface{} `json:"i$ClientSettableUsage"`
			I_CompressionRate                 int           `json:"i$CompressionRate"`
			I_CompressionState                int           `json:"i$CompressionState"`
			I_ConsumableBlocks                int           `json:"i$ConsumableBlocks"`
			I_CreationClassName               string        `json:"i$CreationClassName"`
			I_DataRedundancy                  int           `json:"i$DataRedundancy"`
			I_DeltaReservation                int           `json:"i$DeltaReservation"`
			I_DeviceID                        string        `json:"i$DeviceID"`
			I_ElementName                     string        `json:"i$ElementName"`
			I_EMCBSPElementType               int           `json:"i$EMCBSPElementType"`
			I_EMCBSPInstanceID                string        `json:"i$EMCBSPInstanceID"`
			I_EMCConsumableBlocksWritten      int           `json:"i$EMCConsumableBlocksWritten"`
			I_EMCDataFormat                   string        `json:"i$EMCDataFormat"`
			I_EMCDIFEnabled                   bool          `json:"i$EMCDIFEnabled"`
			I_EMCFullyAllocated               bool          `json:"i$EMCFullyAllocated"`
			I_EMCIsBound                      bool          `json:"i$EMCIsBound"`
			I_EMCIsComposite                  bool          `json:"i$EMCIsComposite"`
			I_EMCIsDeDuplicated               bool          `json:"i$EMCIsDeDuplicated"`
			I_EMCIsMapped                     bool          `json:"i$EMCIsMapped"`
			I_EMCRaidLevel                    string        `json:"i$EMCRaidLevel"`
			I_EMCRecoverPointEnabled          bool          `json:"i$EMCRecoverPointEnabled"`
			I_EMCSpaceConsumed                int           `json:"i$EMCSpaceConsumed"`
			I_EMCVolumeAttributes             []int         `json:"i$EMCVolumeAttributes"`
			I_EMCVolumeAttributes2            []int         `json:"i$EMCVolumeAttributes2"`
			I_EMCVolumeAttributesDescription  []string      `json:"i$EMCVolumeAttributesDescription"`
			I_EMCVolumeAttributesDescription2 []string      `json:"i$EMCVolumeAttributesDescription2"`
			I_EMCVolumeLabel                  string        `json:"i$EMCVolumeLabel"`
			I_EMCWWN                          string        `json:"i$EMCWWN"`
			I_EnabledState                    int           `json:"i$EnabledState"`
			I_ExtentDiscriminator             []string      `json:"i$ExtentDiscriminator"`
			I_ExtentStatus                    []int         `json:"i$ExtentStatus"`
			I_HealthState                     int           `json:"i$HealthState"`
			I_IdentifyingDescriptions         []string      `json:"i$IdentifyingDescriptions"`
			I_IsBasedOnUnderlyingRedundancy   bool          `json:"i$IsBasedOnUnderlyingRedundancy"`
			I_IsComposite                     bool          `json:"i$IsComposite"`
			I_IsCompressed                    bool          `json:"i$IsCompressed"`
			I_Name                            string        `json:"i$Name"`
			I_NameFormat                      int           `json:"i$NameFormat"`
			I_NameNamespace                   int           `json:"i$NameNamespace"`
			I_NoSinglePointOfFailure          bool          `json:"i$NoSinglePointOfFailure"`
			I_NumberOfBlocks                  int           `json:"i$NumberOfBlocks"`
			I_OperationalStatus               []int         `json:"i$OperationalStatus"`
			I_OtherIdentifyingInfo            []string      `json:"i$OtherIdentifyingInfo"`
			I_OtherNameFormat                 string        `json:"i$OtherNameFormat"`
			I_PackageRedundancy               int           `json:"i$PackageRedundancy"`
			I_Primordial                      bool          `json:"i$Primordial"`
			I_Purpose                         string        `json:"i$Purpose"`
			I_RequestedState                  int           `json:"i$RequestedState"`
			I_SequentialAccess                bool          `json:"i$SequentialAccess"`
			I_SystemCreationClassName         string        `json:"i$SystemCreationClassName"`
			I_SystemName                      string        `json:"i$SystemName"`
			I_StatusDescriptions              []string      `json:"i$StatusDescriptions"`
			I_StatusInfo                      int           `json:"i$StatusInfo"`
			I_StorageTieringSelection         int           `json:"i$StorageTieringSelection"`
			I_ThinlyProvisioned               bool          `json:"i$ThinlyProvisioned"`
			I_TransitioningToState            int           `json:"i$TransitioningToState"`
			I_Usage                           int           `json:"i$Usage"`
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

///////////////////////////////////////////////////////////////
//            GET a list of Storage Volumes                  //
///////////////////////////////////////////////////////////////

func (smis *SMIS) GetVolumes(sid string) (resp *GetVolumesResp, err error) {
	err = smis.query("GET", "/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::Symm_StorageSystem,Name::"+sid+"/relationships/CIM_StorageVolume", nil, &resp)
	return resp, err
}

///////////////////////////////////////////////////////////
//            GET a Storage Volume by ID                 //
///////////////////////////////////////////////////////////

func (smis *SMIS) GetVolumeByID(sid string, volumeID string) (resp *GetVolumesResp, err error) {
	err = smis.query("GET", "/ecom/edaa/root/emc/instances/Symm_StorageVolume/CreationClassName::Symm_StorageVolume,DeviceID::"+volumeID+",SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid, nil, &resp)
	return resp, err
}

///////////////////////////////////////////////////////////
//            GET a Storage Volume by Name               //
///////////////////////////////////////////////////////////

func (smis *SMIS) GetVolumeByName(sid string, volumeName string) (resp *GetVolumesResp, err error) {
	err = smis.query("GET", `/ecom/edaa/root/emc/types/CIM_StorageVolume/instances?filter=SystemName lk "`+sid+`" and ElementName lk "`+volumeName+`"`, nil, &resp)
	return resp, err
}

////////////////////////////////////////////////////////////
//            RESPONSE Struct used for                    //
//              getting a Job Status                      //
////////////////////////////////////////////////////////////

type GetJobStatusResp struct {
	Entries []struct {
		Content struct {
			AtType                  string   `json:"@type"`
			I_DeleteOnCompletion    bool     `json:"i$DeleteOnCompletion"`
			I_Description           string   `json:"i$Description"`
			I_ElapsedTime           string   `json:"i$ElapsedTime"`
			I_ErrorCode             int      `json:"i$ErrorCode"`
			I_ErrorDescription      string   `json:"i$ErrorDescription"`
			I_HealthState           int      `json:"i$HealthState"`
			I_JobRunTimes           int      `json:"i$JobRunTimes"`
			I_JobState              int      `json:"i$JobState"`
			I_JobStatus             string   `json:"i$JobStatus"`
			I_LocalOrUtcTime        int      `json:"i$LocalOrUtcTime"`
			I_Name                  string   `json:"i$Name"`
			I_OperationalStatus     []int    `json:"i$OperationalStatus"`
			I_Owner                 string   `json:"i$Owner"`
			I_PercentComplete       int      `json:"i$PercentComplete"`
			I_StartTime             string   `json:"i$StartTime"`
			I_Status                string   `json:"i$Status"`
			I_StatusDescriptions    []string `json:"i$StatusDescriptions"`
			I_TimeBeforeRemoval     string   `json:"i$TimeBeforeRemoval"`
			I_TimeOfLastStateChange string   `json:"i$TimeOfLastStateChange"`
			I_TimeSubmitted         string   `json:"i$TimeSubmitted"`
			I_UntilTime             string   `json:"i$UntilTime"`
			Links                   []struct {
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

//////////////////////////////////////////////////////////////////////////////////////////////////
//      GET a Job Status                                                                        //
//                                                                                              //
//  2 - New: job has not been started                                                           //
//  3 - Starting: job is moving into running state                                              //
//  4 - Running: job is running                                                                 //
//  5 - Suspended: job is stopped, but can be restarted                                         //
//  6 - Shutting Down: job is moving to an completed/terminated/killed state                    //
//  7 - Completed: job has been completed normally                                              //
//  8 - Terminated: job has been stopped by a terminate state change request                    //
//  9 - Killed: job has stopped by a kill state change request                                  //
//  10 - Exception: job is in an abnormal state due to an error condition                       //
//  11 - Service: job is in a vendor-specific state that supports problem discovery/resolution  //
//  12 - Query Pending: job is waiting for a client to resolve a query                          //
//////////////////////////////////////////////////////////////////////////////////////////////////

func (smis *SMIS) GetJobStatus(instanceID string) (resp *GetJobStatusResp, JobStatus string, err error) {
	err = smis.query("GET", "/ecom/edaa/root/emc/instances/SE_ConcreteJob/InstanceID::"+instanceID, nil, &resp)
	if err != nil {
		return nil, "", err
	}
	jobStatusMap := map[int]string{
		2:  "NEW",
		3:  "STARTING",
		4:  "RUNNING",
		5:  "SUSPENDED",
		6:  "SHUTTING_DOWN",
		7:  "COMPLETED",
		8:  "TERMINATED",
		9:  "KILLED",
		10: "EXCEPTION",
		11: "SERVICE",
		12: "QUERY_PENDING",
	}

	var jobStatus string
	var ok bool
	if jobStatus, ok = jobStatusMap[resp.Entries[0].Content.I_JobState]; !ok {
		jobStatus = "UNKNOWN"
	}
	return resp, jobStatus, err
}

//////////////////////////////////////
//    REQUEST Structs used for      //
//   volume creation on the VMAX3.  //
//////////////////////////////////////

type PostVolumesReq struct {
	PostVolumesRequestContent *PostVolumesReqContent `json:"content"`
}

type PostVolumesReqContent struct {
	AtType             string `json:"@type"`
	ElementName        string `json:"ElementName"`
	ElementType        string `json:"ElementType"`
	EMCNumberOfDevices string `json:"EMCNumberOfDevices"`
	Size               string `json:"Size"`
}

////////////////////////////////////////////////////////////
//            RESPONSE Struct used for                    //
//          volume creation on the VMAX3.                 //
////////////////////////////////////////////////////////////

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

///////////////////////////////////////////////////////////
//              CREATE a Storage Volume                  //
//     and check for Volume Creation Completion          //
///////////////////////////////////////////////////////////

func (smis *SMIS) PostVolumes(req *PostVolumesReq, sid string) (resp1 *PostVolumesResp, resp2 *GetJobStatusResp, err error) {
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_StorageConfigurationService/CreationClassName::Symm_StorageConfigurationService,Name::EMCStorageConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/action/CreateOrModifyElementFromStoragePool", req, &resp1)

	err = smis.query("GET", "/ecom/edaa/root/emc/instances/SE_ConcreteJob/InstanceID::"+resp1.Entries[0].Content.I_Parameters.I_Job.E0_InstanceID, nil, &resp2)
	JobStatus := resp2.Entries[0].Content.I_JobState
	for JobStatus == 2 || JobStatus == 3 || JobStatus == 4 || JobStatus == 5 || JobStatus == 6 || JobStatus == 12 {
		err = smis.query("GET", "/ecom/edaa/root/emc/instances/SE_ConcreteJob/InstanceID::"+resp1.Entries[0].Content.I_Parameters.I_Job.E0_InstanceID, nil, &resp2)
		JobStatus = resp2.Entries[0].Content.I_JobState
	}
	if JobStatus != 7 {
		fmt.Println("Error: Volume creation incomplete")
	}
	return resp1, resp2, err
}

//////////////////////////////////////
//   REQUEST Structs used for any   //
//   group creation on the VMAX3.   //
//                                  //
//    Storage Group (SG) - Type 4   //
//     Port Group (PG) - Type 3     //
//   Initiator Group (IG) - Type 2  //
//////////////////////////////////////

type PostGroupReq struct {
	PostGroupRequestContent *PostGroupReqContent `json:"content"`
}

type PostGroupReqContent struct {
	AtType    string `json:"@type"`
	GroupName string `json:"GroupName"`
	Type      string `json:"Type"`
}

////////////////////////////////////////////////////////////
//           RESPONSE Struct  used for any                //
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
//                  CREATE an Array Group                    //
// Type Depends on Type field specified in requesting struct //
///////////////////////////////////////////////////////////////

func (smis *SMIS) PostCreateGroup(req *PostGroupReq, sid string) (resp *PostGroupResp, err error) {
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::SYMMETRIX-+-"+sid+"/action/CreateGroup", req, &resp)
	return resp, err
}

////////////////////////////////////////////////////////////
//      RESPONSE Struct used for each SRP settings        //
////////////////////////////////////////////////////////////

type GetStoragePoolSettingsResp struct {
	Entries []struct {
		Content struct {
			AtType                           string  `json:"@type"`
			I_Changeable                     bool    `json:"i$Changeable"`
			I_ChangeableType                 int     `json:"i$ChangeableType"`
			I_CompressedElement              bool    `json:"i$CompressedElement"`
			I_CompressionRate                int     `json:"i$CompressionRate"`
			I_DataRedundancyGoal             int     `json:"i$DataRedundancyGoal"`
			I_DataRedundancyMax              int     `json:"i$DataRedundancyMax"`
			I_DataRedundancyMin              int     `json:"i$DataRedundancyMin"`
			I_DeltaReservationGoal           int     `json:"i$DeltaReservationGoal"`
			I_DeltaReservationMax            int     `json:"i$DeltaReservationMax"`
			I_DeltaReservationMin            int     `json:"i$DeltaReservationMin"`
			I_ElementName                    string  `json:"i$ElementName"`
			I_EMCApproxAverageResponseTime   float64 `json:"i$EMCApproxAverageResponseTime"`
			I_EMCDeduplicationRate           int     `json:"i$EMCDeduplicationRate"`
			I_EMCEnableDIF                   int     `json:"i$EMCEnableDIF"`
			I_EMCEnableEFDCache              int     `json:"i$EMCEnableEFDCache"`
			I_EMCFastSetting                 string  `json:"i$EMCFastSetting"`
			I_EMCParticipateInPowerSavings   int     `json:"i$EMCParticipateInPowerSavings"`
			I_EMCPoolCompressionState        int     `json:"i$EMCPoolCompressionState"`
			I_EMCPottedSetting               bool    `json:"i$EMCPottedSetting"`
			I_EMCRaidGroupLUN                bool    `json:"i$EMCRaidGroupLUN"`
			I_EMCRaidLevel                   string  `json:"i$EMCRaidLevel"`
			I_EMCSLO                         string  `json:"i$EMCSLO"`
			I_EMCSLOBaseName                 string  `json:"i$EMCSLOBaseName"`
			I_EMCSLOdescription              string  `json:"i$EMCSLOdescription"`
			I_EMCSRP                         string  `json:"i$EMCSRP"`
			I_EMCStorageSettingType          int     `json:"i$EMCStorageSettingType"`
			I_EMCUniqueID                    string  `json:"i$EMCUniqueID"`
			I_EMCWorkload                    string  `json:"i$EMCWorkload"`
			I_ExtentStripeLength             int     `json:"i$ExtentStripeLength"`
			I_ExtentStripeLengthMax          int     `json:"i$ExtentStripeLengthMax"`
			I_ExtentStripeLengthMin          int     `json:"i$ExtentStripeLengthMin"`
			I_InitialStorageTierMethodology  int     `json:"i$InitialStorageTierMethodology"`
			I_InitialStorageTieringSelection int     `json:"i$InitialStorageTieringSelection"`
			I_InitialSynchronization         int     `json:"i$InitialSynchronization"`
			I_InstanceID                     string  `json:"i$InstanceID"`
			I_NoSinglePointOfFailure         bool    `json:"i$NoSinglePointOfFailure"`
			I_PackageRedundancyGoal          int     `json:"i$PackageRedundancyGoal"`
			I_PackageRedundancyMax           int     `json:"i$PackageRedundancyMax"`
			I_PackageRedundancyMin           int     `json:"i$PackageRedundancyMin"`
			I_SpaceLimit                     int     `json:"i$SpaceLimit"`
			I_StorageExtentInitialUsage      int     `json:"i$StorageExtentInitialUsage"`
			I_StoragePoolInitialUsage        int     `json:"i$StoragePoolInitialUsage"`
			I_ThinProvisionedPoolType        int     `json:"i$ThinProvisionedPoolType"`
			I_UseReplicationBuffer           int     `json:"i$UseReplicationBuffer"`
			Links                            []struct {
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

///////////////////////////////////////////////////////////////
//                GET Storage Pool Settings                  //
///////////////////////////////////////////////////////////////

func (smis *SMIS) GetStoragePoolSettings(srp_name, sid string) (resp *GetStoragePoolSettingsResp, err error) {
	err = smis.query("GET", "/ecom/edaa/root/emc/instances/Symm_StoragePoolCapabilities/InstanceID::SYMMETRIX-+-"+sid+"-+-SR-+-"+srp_name+"/relationships/CIM_StorageSetting", nil, &resp)
	return resp, err
}

///////////////////////////////////////////////////////////////
//        Struct used to store all SLO information           //
///////////////////////////////////////////////////////////////

type SLO_Struct struct {
	SLO_Name    string
	respTime    float64
	SRP         string
	Workload    string
	ElementName string
	InstanceID  string
}

////////////////////////////////////////////////////////////////
//               GET Storage Level Objectives                 //
//                                                            //
//             1 -> Get Storage Pools of VMAX3                //
//    2 -> Get Storage Pool Settings of each Storage Pool     //
//   3 -> Parse out SLO information of VMAX3 and return it    //
////////////////////////////////////////////////////////////////

func (smis *SMIS) GetSLOs(sid string) (SLOs []SLO_Struct, err error) {
	storagePools, err := smis.GetStoragePools(sid)
	if err != nil {
		return nil, err
	}

	for _, SRP := range storagePools.Entries {
		storagePoolSetting, err := smis.GetStoragePoolSettings(SRP.Content.I_ElementName, sid)
		if err != nil {
			return nil, err
		}
		for _, storagePoolSetting := range storagePoolSetting.Entries {
			if storagePoolSetting.Content.I_EMCSLO != "" {
				newSLO := SLO_Struct{
					SLO_Name:    storagePoolSetting.Content.I_EMCSLOBaseName,
					respTime:    storagePoolSetting.Content.I_EMCApproxAverageResponseTime,
					SRP:         storagePoolSetting.Content.I_EMCSRP,
					Workload:    storagePoolSetting.Content.I_EMCWorkload,
					ElementName: storagePoolSetting.Content.I_ElementName,
					InstanceID:  storagePoolSetting.Content.I_InstanceID,
				}
				SLOs = append(SLOs, newSLO)
			}
		}
	}
	return SLOs, nil
}

/////////////////////////////////////////////////////////
//               REQUEST Structs used for              //
//         adding AND removing a volume to/from        //
//             a storage group on the VMAX3.           //
/////////////////////////////////////////////////////////

type PostVolumesToSGReq struct {
	PostVolumesToSGRequestContent *PostVolumesToSGReqContent `json:"content"`
}

type PostVolumesToSGReqContent struct {
	AtType                              string                             `json:"@type"`
	PostVolumesToSGRequestContentMG     *PostVolumesToSGReqContentMG       `json:"MaskingGroup"`
	PostVolumesToSGRequestContentMember []*PostVolumesToSGReqContentMember `json:"Members"`
}

type PostVolumesToSGReqContentMG struct {
	AtType     string `json:"@type"`
	InstanceID string `json:"InstanceID"`
}

type PostVolumesToSGReqContentMember struct {
	AtType                  string `json:"@type"`
	CreationClassName       string `json:"CreationClassName"`
	DeviceID                string `json:"DeviceID"`
	SystemCreationClassName string `json:"SystemCreationClassName"`
	SystemName              string `json:"SystemName"`
}

////////////////////////////////////////////////////////////
//               RESPONSE Struct used for                 //
//         adding AND removing a volume to/from           //
//             a storage group on the VMAX3.              //
////////////////////////////////////////////////////////////

type PostVolumesToSGResp struct {
	Entries []struct {
		Content struct {
			AtType       string `json:"@type"`
			I_Parameters struct {
				I_Job struct {
					AtType        string `json:"@type"`
					E0_InstanceID string `json:"e0$InstanceID"`
					Xmlns_e0      string `json:"xmlns$e0"`
				} `json:"i$Job"`
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

///////////////////////////////////////////////////////////////
//             ADD Volumes to a Storage Group                //
///////////////////////////////////////////////////////////////

func (smis *SMIS) PostVolumesToSG(req *PostVolumesToSGReq, sid string) (resp *PostVolumesToSGResp, err error) {
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/action/AddMembers", req, &resp)
	return resp, err
}

///////////////////////////////////////////////////////////////
//          REMOVE Volumes from a Storage Group              //
///////////////////////////////////////////////////////////////

func (smis *SMIS) RemoveVolumeFromSG(req *PostVolumesToSGReq, sid string) (resp *PostVolumesToSGResp, err error) {
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/action/RemoveMembers", req, &resp)
	return resp, err
}

/////////////////////////////////////////////////////////
//               REQUEST Structs used for              //
//   creating a storage hardware ID for an initiator   //
/////////////////////////////////////////////////////////

type PostStorageHardwareIDReq struct {
	PostStorageHardwareIDRequestContent *PostStorageHardwareIDReqContent `json:"content"`
}

type PostStorageHardwareIDReqContent struct {
	AtType    string `json:"@type"`
	IDType    string `json:"IDType"`
	StorageID string `json:"StorageID"`
}

////////////////////////////////////////////////////////////
//            RESPONSE Struct used for                    //
//   creating a storage hardware ID for an initiator      //
////////////////////////////////////////////////////////////

type PostStorageHardwareIDResp struct {
	Entries []struct {
		Content struct {
			AtType       string `json:"@type"`
			I_Parameters struct {
				I_HardwareID struct {
					AtType        string `json:"@type"`
					E0_InstanceID string `json:"e0$InstanceID"`
					Xmlns_e0      string `json:"xmlns$e0"`
				} `json:"i$HardwareID"`
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

/////////////////////////////////////////////////////////
//               REQUEST Struct used for               //
//       adding AND removing an initiator to/from      //
//              a host group on the VMAX3.             //
/////////////////////////////////////////////////////////

type PostInitiatorToHGReq struct {
	PostInitiatorToHGRequestContent *PostInitiatorToHGReqContent `json:"content"`
}

type PostInitiatorToHGReqContent struct {
	AtType                                string                               `json:"@type"`
	PostInitiatorToHGRequestContentMG     *PostInitiatorToHGReqContentMG       `json:"MaskingGroup"`
	PostInitiatorToHGRequestContentMember []*PostInitiatorToHGReqContentMember `json:"Members"`
}

type PostInitiatorToHGReqContentMG struct {
	AtType     string `json:"@type"`
	InstanceID string `json:"InstanceID"`
}

type PostInitiatorToHGReqContentMember struct {
	AtType     string `json:"@type"`
	InstanceID string `json:"InstanceID"`
}

////////////////////////////////////////////////////////////
//                RESPONSE Struct used for                //
//        adding AND removing an initiator to/from        //
//             a host group on the VMAX3.                 //
////////////////////////////////////////////////////////////

type PostInitiatorToHGResp struct {
	Entries []struct {
		Content struct {
			AtType       string `json:"@type"`
			I_Parameters struct {
				I_Job struct {
					AtType        string `json:"@type"`
					E0_InstanceID string `json:"e0$InstanceID"`
					Xmlns_e0      string `json:"xmlns$e0"`
				} `json:"i$Job"`
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

///////////////////////////////////////////////////////////////
//             ADD Initiators to a Host Group                //
//                                                           //
//     1 -> Create Storage Hardware ID for the Initiator     //
//     2 -> Add Storage Hardware ID to Initiator Group       //
///////////////////////////////////////////////////////////////

func (smis *SMIS) PostStorageHardwareID(req *PostStorageHardwareIDReq, sid string) (resp *PostStorageHardwareIDResp, err error) {
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_StorageHardwareIDManagementService/CreationClassName::Symm_StorageHardwareIDManagementService,Name::EMCStorageHardwareIDManagementService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/action/CreateStorageHardwareID", req, &resp)
	return resp, err
}

func (smis *SMIS) PostInitiatorToHG(req *PostInitiatorToHGReq, sid string) (resp *PostInitiatorToHGResp, err error) {
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/action/AddMembers", req, &resp)
	return resp, err
}

///////////////////////////////////////////////////////////////
//          REMOVE Initiators from a Host Group              //
//     (Requires a Storage Hardware ID from the Initiator)   //
///////////////////////////////////////////////////////////////

func (smis *SMIS) RemoveInitiatorFromHG(req *PostInitiatorToHGReq, sid string) (resp *PostInitiatorToHGResp, err error) {
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/action/RemoveMembers", req, &resp)
	return resp, err
}

/////////////////////////////////////////////////////////
//               REQUEST Structs used for              //
//         adding AND removing a port to/from          //
//           a port group on the VMAX3.                //
/////////////////////////////////////////////////////////

type PostPortToPGReq struct {
	PostPortToPGRequestContent *PostPortToPGReqContent `json:"content"`
}

type PostPortToPGReqContent struct {
	AtType                           string                          `json:"@type"`
	PostPortToPGRequestContentMG     *PostPortToPGReqContentMG       `json:"MaskingGroup"`
	PostPortToPGRequestContentMember []*PostPortToPGReqContentMember `json:"Members"`
}

type PostPortToPGReqContentMG struct {
	AtType     string `json:"@type"`
	InstanceID string `json:"InstanceID"`
}

type PostPortToPGReqContentMember struct {
	AtType                  string `json:"@type"`
	CreationClassName       string `json:"CreationClassName"`
	Name                    string `json:"Name"`
	SystemCreationClassName string `json:"SystemCreationClassName"`
	SystemName              string `json:"SystemName"`
}

///////////////////////////////////////////////////////
//            RESPONSE Struct used for               //
//       adding AND removing a port to/from          //
//            a port group on the VMAX3.             //
///////////////////////////////////////////////////////

type PostPortToPGResp struct {
	Entries []struct {
		Content struct {
			AtType       string `json:"@type"`
			I_Parameters struct {
				I_Job struct {
					AtType        string `json:"@type"`
					E0_InstanceID string `json:"e0$InstanceID"`
					Xmlns_e0      string `json:"xmlns$e0"`
				} `json:"i$Job"`
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

/////////////////////////////////////////////////////////////////////
//                     ADD Ports to a Port Group                   //
//                                                                 //
//    1 -> GET a list of Available Interfaces (aka FE Directors)   //
// 2 -> GET a list of Front End Adapter Endpoints (aka FE Ports)   //
//                  3 -> ADD Ports to Port Groups                  //
/////////////////////////////////////////////////////////////////////

func (smis *SMIS) PostPortToPG(req *PostPortToPGReq, sid string) (resp *PostPortToPGResp, err error) {
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/action/AddMembers", req, &resp)
	return resp, err
}

///////////////////////////////////////////////////////////////
//             REMOVE Ports from a Port Group                //
///////////////////////////////////////////////////////////////

func (smis *SMIS) RemovePortFromPG(req *PostPortToPGReq, sid string) (resp *PostPortToPGResp, err error) {
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/action/RemoveMembers", req, &resp)
	return resp, err
}

/////////////////////////////////////////////////////////
//               REQUEST Structs used for              //
//        creating a masking view on the VMAX3.        //
/////////////////////////////////////////////////////////

type PostCreateMaskingViewReq struct {
	PostCreateMaskingViewRequestContent *PostCreateMaskingViewReqContent `json:"content"`
}

type PostCreateMaskingViewReqContent struct {
	AtType                           string                        `json:"@type"`
	ElementName                      string                        `json:"ElementName"`
	PostInitiatorMaskingGroupRequest *PostInitiatorMaskingGroupReq `json:"InitiatorMaskingGroup"`
	PostTargetMaskingGroupRequest    *PostTargetMaskingGroupReq    `json:"TargetMaskingGroup"`
	PostDeviceMaskingGroupRequest    *PostDeviceMaskingGroupReq    `json:"DeviceMaskingGroup"`
}

type PostInitiatorMaskingGroupReq struct {
	AtType     string `json:"@type"`
	InstanceID string `json:"InstanceID"`
}
type PostTargetMaskingGroupReq struct {
	AtType     string `json:"@type"`
	InstanceID string `json:"InstanceID"`
}
type PostDeviceMaskingGroupReq struct {
	AtType     string `json:"@type"`
	InstanceID string `json:"InstanceID"`
}

////////////////////////////////////////////////////////////
//            RESPONSE Struct used for                    //
//        creating a masking view on the VMAX3.           //
////////////////////////////////////////////////////////////

type PostCreateMaskingViewResp struct {
	Xmlns_gd string `json:"xmlns$gd"`
	Updated  string `json:"updated"`
	ID       string `json:"id"`

	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`

	Entries []struct {
		Updated string `json:"updated"`

		Links []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"links"`

		Content_type string `json:"content-type"`

		Content struct {
			AtType       string `json:"@type"`
			Xmlns_i      string `json:"xmlns$i"`
			I_Parameters struct {
				I_Job struct {
					AtType        string `json:"@type"`
					Xmlns_e0      string `json:"xmlns$e0"`
					E0_InstanceID string `json:"e0$InstanceID"`
				} `json:"i$Job"`
			} `json:"i$parameters"`
			I_ReturnValue int `json:"i$returnValue"`
		} `json:"content"`
	} `json:"entries"`
}

///////////////////////////////////////////////////////////////
//                  CREATE a Masking View                    //
///////////////////////////////////////////////////////////////

func (smis *SMIS) PostCreateMaskingView(req *PostCreateMaskingViewReq, sid string) (resp *PostCreateMaskingViewResp, err error) {
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/action/CreateMaskingView", req, &resp)
	return resp, err
}

////////////////////////////////////////////////////////////////
//     Struct used to store all Baremetal HBA Information     //
////////////////////////////////////////////////////////////////

type HostAdapter struct {
	HostID string
	WWN    string
}

////////////////////////////////////////////////////////////////
//             GET Baremetal HBA Information                  //
////////////////////////////////////////////////////////////////

func GetBaremetalHBA() (myHosts []HostAdapter, err error) {
	//Works for RedHat 5 and above (including CentOS and SUSE Linux)
	hostDir, _ := ioutil.ReadDir("/sys/class/scsi_host/")

	for _, host := range hostDir {
		if _, err := os.Stat("/sys/class/scsi_host/" + host.Name() + "/device/fc_host/" + host.Name() + "/port_name"); err == nil {
			byteArray, _ := ioutil.ReadFile("/sys/class/scsi_host/" + host.Name() + "/device/fc_host/" + host.Name() + "/port_name")
			newHost := HostAdapter{
				HostID: host.Name(),
				WWN:    string(byteArray),
			}
			myHosts = append(myHosts, newHost)
		}
	}
	return myHosts, nil
}

//////////////////////////////////////////////////////////////
//             REQUEST Structs used for any                 //
//             group deletion on the VMAX3.                 //
//                                                          //
//   Storage Group (SG) - AtType SE_DeviceMaskingGroup      //
//      Port Group (PG) - AtType SE_TargetMaskingGroup      //
//  Initiator Group (IG) - AtType SE_InitiatorMaskingGroup  //
//////////////////////////////////////////////////////////////

type DeleteGroupReq struct {
	DeleteGroupRequestContent *DeleteGroupReqContent `json:"content"`
}

type DeleteGroupReqContent struct {
	AtType                                string                             `json:"@type"`
	DeleteGroupRequestContentMaskingGroup *DeleteGroupReqContentMaskingGroup `json:"MaskingGroup"`
}

type DeleteGroupReqContentMaskingGroup struct {
	AtType     string `json:"@type"`
	InstanceID string `json:"InstanceID"`
}

////////////////////////////////////////////////////////////
//           RESPONSE Struct used for any                 //
//           group deletion on the VMAX3.                 //
////////////////////////////////////////////////////////////

type DeleteGroupResp struct {
	Entries []struct {
		Content struct {
			AtType       string `json:"@type"`
			I_parameters struct {
				I_Job struct {
					AtType        string `json:"@type"`
					E0_InstanceID string `json:"e0$InstanceID"`
					Xmlns_e0      string `json:"xmlns$e0"`
				} `json:"i$Job"`
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

/////////////////////////////////////////////////////////////////
//                  DELETE an Array Group                      //
// Type Depends on AtType field specified in requesting struct //
/////////////////////////////////////////////////////////////////

func (smis *SMIS) PostDeleteGroup(req *DeleteGroupReq, sid string) (resp *DeleteGroupResp, err error) {
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/action/DeleteGroup", req, &resp)
	return resp, err
}

//////////////////////////////////////////////////////////////
//             REQUEST Structs used for any                 //
//             volume deletion on the VMAX3.                //
//////////////////////////////////////////////////////////////

type DeleteVolReq struct {
	DeleteVolRequestContent *DeleteVolReqContent `json:"content"`
}

type DeleteVolReqContent struct {
	AtType                         string                      `json:"@type"`
	DeleteVolRequestContentElement *DeleteVolReqContentElement `json:"TheElement"`
}

type DeleteVolReqContentElement struct {
	AtType                  string `json:"@type"`
	DeviceID                string `json:"DeviceID"`
	CreationClassName       string `json:"CreationClassName"`
	SystemName              string `json:"SystemName"`
	SystemCreationClassName string `json:"SystemCreationClassName"`
}

////////////////////////////////////////////////////////////
//           RESPONSE Struct used for any                 //
//           volume deletion on the VMAX3.                //
////////////////////////////////////////////////////////////

type DeleteVolResp struct {
	Entries []struct {
		Content struct {
			AtType       string `json:"@type"`
			I_Parameters struct {
				I_Job struct {
					AtType        string `json:"@type"`
					E0_InstanceID string `json:"e0$InstanceID"`
					Xmlns_e0      string `json:"xmlns$e0"`
				} `json:"i$Job"`
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

/////////////////////////////////////////////////////////////////
//                  DELETE a Volume                            //
/////////////////////////////////////////////////////////////////

func (smis *SMIS) PostDeleteVol(req *DeleteVolReq, sid string) (resp *DeleteVolResp, err error) {
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_StorageConfigurationService/CreationClassName::Symm_StorageConfigurationService,Name::EMCStorageConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/action/ReturnToStoragePool", req, &resp)
	return resp, err
}

//////////////////////////////////////////////////////////////
//             REQUEST Structs used for any                 //
//          masking view deletion on the VMAX3.             //
//////////////////////////////////////////////////////////////

type DeleteMaskingViewReq struct {
	DeleteMaskingViewRequestContent *DeleteMaskingViewReqContent `json:"content"`
}

type DeleteMaskingViewReqContent struct {
	AtType                            string                         `json:"@type"`
	DeleteMaskingViewRequestContentPC *DeleteMaskingViewReqContentPC `json:"ProtocolController"`
}

type DeleteMaskingViewReqContentPC struct {
	AtType                  string `json:"@type"`
	DeviceID                string `json:"DeviceID"`
	CreationClassName       string `json:"CreationClassName"`
	SystemName              string `json:"SystemName"`
	SystemCreationClassName string `json:"SystemCreationClassName"`
}

////////////////////////////////////////////////////////////
//           RESPONSE Struct used for any                 //
//        masking view deletion on the VMAX3.             //
////////////////////////////////////////////////////////////

type DeleteMaskingViewResp struct {
	Entries []struct {
		Content struct {
			AtType       string `json:"@type"`
			I_parameters struct {
				I_Job struct {
					AtType        string `json:"@type"`
					E0_InstanceID string `json:"e0$InstanceID"`
					Xmlns_e0      string `json:"xmlns$e0"`
				} `json:"i$Job"`
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

/////////////////////////////////////////////////////////////////
//               DELETE a Masking View                         //
/////////////////////////////////////////////////////////////////

func (smis *SMIS) PostDeleteMaskingView(req *DeleteMaskingViewReq, sid string) (resp *DeleteMaskingViewResp, err error) {
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_ControllerConfigurationService/CreationClassName::Symm_ControllerConfigurationService,Name::EMCControllerConfigurationService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/action/DeleteMaskingView", req, &resp)
	return resp, err
}

/////////////////////////////////////////////////////////
//               REQUEST Structs used for              //
//        getting Ports logged in, on the VMAX3.       //
/////////////////////////////////////////////////////////

type PostPortLoggedInReq struct {
	PostPortLoggedInRequestContent *PostPortLoggedInReqContent `json:"content"`
}

type PostPortLoggedInReqContent struct {
	PostPortLoggedInRequestHardwareID *PostPortLoggedInReqHardwareID `json:"HardwareID"`
	AtType                            string                         `json:"@type"`
}

type PostPortLoggedInReqHardwareID struct {
	AtType     string `json:"@type"`
	InstanceID string `json: "InstanceID"`
}

/////////////////////////////////////////////////////////
//               RESPONSE Structs used for             //
//        getting Ports logged in, on the VMAX3.       //
/////////////////////////////////////////////////////////

type PostPortLoginResp struct {
	Entries []struct {
		Content struct {
			_type        string `json:"@type"`
			I_parameters struct {
				I_TargetEndpoints []map[string]string `json:"i$TargetEndpoints"`
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

type PortValues struct {
	WWN        string
	PortNumber string
	Director   string
}

///////////////////////////////////////////////////////////////
//           Getting Ports Logged In                         //
///////////////////////////////////////////////////////////////

func (smis *SMIS) PostPortLogins(req *PostPortLoggedInReq, sid string) (portvalues1 []PortValues, err error) {

	var resp *PostPortLoginResp
	var wwn string = req.PostPortLoggedInRequestContent.PostPortLoggedInRequestHardwareID.InstanceID
	wwn = "W-+-" + wwn
	req.PostPortLoggedInRequestContent.PostPortLoggedInRequestHardwareID.InstanceID = wwn
	err = smis.query("POST", "/ecom/edaa/root/emc/instances/Symm_StorageHardwareIDManagementService/CreationClassName::Symm_StorageHardwareIDManagementService,Name::EMCStorageHardwareIDManagementService,SystemCreationClassName::Symm_StorageSystem,SystemName::"+sid+"/action/EMCGetTargetEndpoints", req, &resp)
	var portValues []PortValues
	var length = len(resp.Entries[0].Content.I_parameters.I_TargetEndpoints)
	for i := 0; i < length; i++ {
		var m map[string]string
		var name string
		m = resp.Entries[0].Content.I_parameters.I_TargetEndpoints[i]
		name = "e" + strconv.Itoa(i) + "$SystemName"
		var eSystemName string = m[name]
		eSystemNameSplit := strings.Split(eSystemName, "-+-")
		PortAndDirector := strings.Split(eSystemNameSplit[2], "-")
		portNumber := PortAndDirector[0]
		director := PortAndDirector[1]
		PV := PortValues{
			WWN:        wwn,
			PortNumber: portNumber,
			Director:   director,
		}
		portValues = append(portValues, PV)
	}
	return portValues, err
}
