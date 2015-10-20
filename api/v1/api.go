package apiv1

import ()

type GetStoragePoolsResp struct {
	Entries []struct {
		Content []struct {
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

type GetStorageGroupsResp struct {
	Entries []struct {
		Content []struct {
			I_ElementName			string		`json:"i$ElementName"`
			I_EMCAssociatedToView	bool		`json:"i$EMCAssociatedToView"`
			I_InstanceID 			string		`json:"i$InstanceID"`
		} `json:"content"`
	} `json:"entries"`
}

//Get a list of Device Masking/Storage Groups associated with SymmID 
func (smis *SMIS) GetStorageGroups(sid string) (resp *GetDeviceMaskingGroupsResp, err error){
	err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::Symm_StorageSystem,Name::" + sid + "/relationships/CIM_DeviceMaskingGroup", nil, &resp)
	return resp,err
}

type GetStorageVolumesResp struct {
	Entries []struct {
		Content []struct {
			I_Access							int 			`json:"i$Access"`
			I_BlockSize							int 			`json:"i$BlockSize"`
			I_Caption 							string 			`json:"i$Caption"`
			I_ClientSettableUsage				[]interface{}	`json:"i$ClientSettableUsage"`
			I_CompressionRate					int 			`json:"i$CompressionRate"`
			I_CompressionState					int 			`json:"i$CompressionState"`
			I_ConsumableBlocks					int 			`json:"i$ConsumableBlocks"`
			I_CreationClassName					string			`json:"i$CreationClassName"`
			I_DataRedundancy					int 			`json:"i$DataRedundancy"`
			I_DeltaReservation					int 			`json:"i$DeltaReservation"`
			I_DeviceID 							string			`json:"i$DeviceID"`
			I_ElementName						string 			`json:"i$ElementName"`
			I_EMCBSPElementType					int				`json:"i$EMCBSPElementType"`
			I_EMCBSPInstanceID					string 			`json:"i$EMCBSPInstanceID"`
			I_EMCConsumableBlocksWritten 		int 			`json:"i$EMCConsumableBlocksWritten"`
			I_EMCDataFormat						string 			`json:"i$EMCDataFormat"`
			I_EMCDIFEnabled						bool			`json:"i$EMCDIFEnabled"`
			I_EMCFullyAllocated					bool			`json:"i$EMCFullyAllocated"`
			I_EMCIsBound						bool			`json:"i$EMCIsBound"`
			I_EMCIsComposite					bool			`json:"i$EMCIsComposite"`
			I_EMCIsDeDuplicated					bool			`json:"i$EMCIsDeDuplicated"`
			I_EMCIsMapped						bool			`json:"i$EMCIsMapped"`
			I_EMCRaidLevel						string 			`json:"i$EMCRaidLevel"`
			I_EMCRecoverPointEnabled			bool			`json:"i$EMCRecoverPointEnabled"`
			I_EMCSpaceConsumed					int 			`json:"i$EMCSpaceConsumed"`
			I_EMCVolumeAttributes				[]int 			`json:"i$EMCVolumeAttributes"`					
			I_EMCVolumeAttributes2				[]int 			`json:"i$EMCVolumeAttributes2"`
			I_EMCVolumeAttributesDescription	[]string		`json:"i$EMCVolumeAttributesDescription"`
			I_EMCVolumeAttributesDescription2	[]string 		`json:"i$EMCVolumeAttributesDescription2"`
			I_EMCVolumeLabel					string 			`json:"i$EMCVolumeLabel"`
			I_EMCWWN							string			`json:"i$EMCWWN"`
			I_EnabledState						int 			`json:"i$EnabledState"`
			I_ExtentDiscriminator				[]string 		`json:"i$ExtentDiscriminator"`														          "SNIA:Allocated",
			I_ExtentStatus						[]int 			`json:"i$ExtentStatus"`
			I_HealthState						int 			`json:"i$HealthState"`
			I_IdentifyingDescriptions			[]string     	`json:"i$IdentifyingDescriptions"`
			I_IsBasedOnUnderlyingRedundancy		bool			`json:"i$IsBasedOnUnderlyingRedundancy"`
			I_IsComposite						bool            `json:"i$IsComposite"`
			I_IsCompressed						bool			`json:"i$IsCompressed"`
			I_Name 								string 			`json:"i$Name"`
			I_NameFormat						int 			`json:"i$NameFormat"`
			I_NameNamespace						int 			`json:"i$NameNamespace"`
			I_NoSinglePointOfFailure			bool			`json:"i$NoSinglePointOfFailure"`
			I_NumberOfBlocks					int 			`json:"i$NumberOfBlocks"`
			I_OperationalStatus					[]int 			`json:"i$OperationalStatus"`
			I_OtherIdentifyingInfo				[]string 		`json:"i$OtherIdentifyingInfo"`
			I_OtherNameFormat					string 			`json:"i$OtherNameFormat"`
			I_PackageRedundancy					int 			`json:"i$PackageRedundancy"`
			I_Primordial						bool			`json:"i$Primordial"`
			I_Purpose							string 			`json:"i$Purpose"`
			I_RequestedState					int 			`json:"i$RequestedState"`
			I_SequentialAccess					bool			`json:"i$SequentialAccess"`
			I_SystemCreationClassName			string 			`json:"i$SystemCreationClassName"`
			I_SystemName 						string 			`json:"i$SystemName"`
			I_StatusDescriptions				[]string 		`json:"i$StatusDescriptions"`
			I_StatusInfo						int 			`json:"i$StatusInfo"`
			I_StorageTieringSelection			int 			`json:"i$StorageTieringSelection"`
			I_ThinlyProvisioned					bool			`json:"i$ThinlyProvisioned"`
			I_TransitioningToState				int 			`json:"i$TransitioningToState"`
			I_Usage 							int 			`json:"i$Usage"`
		} `json:"content"`
	} `json:"entries"`
}

//Get a list of Storage Volumes on/associated with SymmID
func (smis *SMIS) GetStorageVolumes(sid string) (resp *GetStorageVolumesResp, err error){
	err = smis.query("GET","/ecom/edaa/root/emc/instances/Symm_StorageSystem/CreationClassName::Symm_StorageSystem,Name::" + sid + "/relationships/CIM_StorageVolume", nil, &resp)
	return resp,err
}

type PostMaskingGroupsReq struct {
	Content []struct {
		GroupName string `json: "GroupName"`
		Type int `json: "Type"`	
			//Type 2 = Initiator Masking Group
			//Type 3 = Target/Port Masking Group
			//Type 4 = Device Masking/Storage Group
		@Type string `json: "@type"`
			//@type = http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerConfigurationService
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
		ElementType string `json: "ElementType"`
		Size string `json: "Size"`
		ElementName string `json: "ElementName"`
		EMCNumberOfDevices string `json: "EMCNumberOfDevices"`
		@Type string `json: "@type"`
			//@type = http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageConfigurationService
		//Add goal stuct later {@type, InstanceID}
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
			@Type string `json: "@type"`
				//@type = http://schemas.emc.com/ecom/edaa/root/emc/SE_DeviceMaskingGroup
			InstanceName string `json: "InstanceID"`
				//InstanceID = SYMMETRIX-+-sid-+-SG_name
		}`json:"MaskingGroup"`,
		Members [] struct {
			@Type string `json: "@type"`
				//@type = http://schemas.emc.com/ecom/edaa/root/emc/Symm_StorageVolume
			CreationClassID string `json: "CreationClassName"`
				//CreationClassName = Symm_StorageVolume
			DeviceName string `json: "DeviceID"`
			SystemCreationClassID string `json: "SystemCreationClassName"`
				//SystemCreationClassName = Symm_StorageSystem
			SystemID string `json: "SystemName"`
				//SystemName = SYMMETRIX-+-sid
		}`json:"Members"`,
		@Type string `json:"@type"`
			//@type = http://schemas.emc.com/ecom/edaa/root/emc/Symm_ControllerconfigurationService
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