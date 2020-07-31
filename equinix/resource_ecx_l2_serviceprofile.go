package equinix

import (
	"fmt"

	"github.com/equinix/ecx-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

var ecxL2ServiceProfileSchemaNames = map[string]string{
	"UUID":                                "uuid",
	"State":                               "state",
	"AlertPercentage":                     "bandwidth_alert_threshold",
	"AllowCustomSpeed":                    "speed_customization_allowed",
	"AllowOverSubscription":               "oversubscription_allowed",
	"APIAvailable":                        "api_integration",
	"AuthKeyLabel":                        "authkey_label",
	"ConnectionNameLabel":                 "connection_name_label",
	"CTagLabel":                           "ctag_label",
	"Description":                         "description",
	"EnableAutoGenerateServiceKey":        "servicekey_autogenerated",
	"EquinixManagedPortAndVlan":           "equinix_managed_port_vlan",
	"IntegrationID":                       "integration_id",
	"Name":                                "name",
	"OnBandwidthThresholdNotification":    "bandwidth_threshold_notifications",
	"OnProfileApprovalRejectNotification": "profile_statuschange_notifications",
	"OnVcApprovalRejectionNotification":   "vc_statuschange_notifications",
	"OverSubscription":                    "oversubscription",
	"Private":                             "private",
	"PrivateUserEmails":                   "private_user_emails",
	"RequiredRedundancy":                  "redundancy_required",
	"SpeedFromAPI":                        "speed_from_api",
	"TagType":                             "tag_type",
	"VlanSameAsPrimary":                   "secondary_vlan_from_primary",
	"Features":                            "features",
	"Port":                                "port",
	"SpeedBand":                           "speed_band",
}

var ecxL2ServiceProfileFeaturesSchemaNames = map[string]string{
	"CloudReach":  "cloud_reach",
	"TestProfile": "test_profile",
}

var ecxL2ServiceProfilePortSchemaNames = map[string]string{
	"ID":        "uuid",
	"MetroCode": "metro_code",
}

var ecxL2ServiceProfileSpeedBandSchemaNames = map[string]string{
	"Speed":     "speed",
	"SpeedUnit": "speed_unit",
}

func resourceECXL2ServiceProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceECXL2ServiceProfileCreate,
		Read:   resourceECXL2ServiceProfileRead,
		Update: resourceECXL2ServiceProfileUpdate,
		Delete: resourceECXL2ServiceProfileDelete,
		Schema: createECXL2ServiceProfileResourceSchema(),
	}
}

func createECXL2ServiceProfileResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		ecxL2ServiceProfileSchemaNames["UUID"]: {
			Type:     schema.TypeString,
			Computed: true,
		},
		ecxL2ServiceProfileSchemaNames["State"]: {
			Type:     schema.TypeString,
			Computed: true,
		},
		ecxL2ServiceProfileSchemaNames["AlertPercentage"]: {
			Type:     schema.TypeFloat,
			Required: true,
		},
		ecxL2ServiceProfileSchemaNames["AllowCustomSpeed"]: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		ecxL2ServiceProfileSchemaNames["AllowOverSubscription"]: {
			Type:     schema.TypeBool,
			Required: true,
		},
		ecxL2ServiceProfileSchemaNames["APIAvailable"]: {
			Type:         schema.TypeBool,
			Optional:     true,
			RequiredWith: []string{ecxL2ServiceProfileSchemaNames["IntegrationID"]},
		},
		ecxL2ServiceProfileSchemaNames["AuthKeyLabel"]: {
			Type:     schema.TypeString,
			Optional: true,
		},
		ecxL2ServiceProfileSchemaNames["ConnectionNameLabel"]: {
			Type:     schema.TypeString,
			Required: true,
		},
		ecxL2ServiceProfileSchemaNames["CTagLabel"]: {
			Type:     schema.TypeString,
			Optional: true,
		},
		ecxL2ServiceProfileSchemaNames["Description"]: {
			Type:     schema.TypeString,
			Optional: true,
		},
		ecxL2ServiceProfileSchemaNames["EnableAutoGenerateServiceKey"]: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		ecxL2ServiceProfileSchemaNames["EquinixManagedPortAndVlan"]: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		ecxL2ServiceProfileSchemaNames["IntegrationID"]: {
			Type:     schema.TypeString,
			Optional: true,
		},
		ecxL2ServiceProfileSchemaNames["Name"]: {
			Type:     schema.TypeString,
			Required: true,
		},
		ecxL2ServiceProfileSchemaNames["OnBandwidthThresholdNotification"]: {
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		ecxL2ServiceProfileSchemaNames["OnProfileApprovalRejectNotification"]: {
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		ecxL2ServiceProfileSchemaNames["OnVcApprovalRejectionNotification"]: {
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		ecxL2ServiceProfileSchemaNames["OverSubscription"]: {
			Type:     schema.TypeString,
			Optional: true,
		},
		ecxL2ServiceProfileSchemaNames["Private"]: {
			Type:         schema.TypeBool,
			Optional:     true,
			RequiredWith: []string{ecxL2ServiceProfileSchemaNames["PrivateUserEmails"]},
		},
		ecxL2ServiceProfileSchemaNames["PrivateUserEmails"]: {
			Type:     schema.TypeSet,
			Optional: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		ecxL2ServiceProfileSchemaNames["RequiredRedundancy"]: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		ecxL2ServiceProfileSchemaNames["SpeedFromAPI"]: {
			Type:         schema.TypeBool,
			Optional:     true,
			AtLeastOneOf: []string{ecxL2ServiceProfileSchemaNames["SpeedFromAPI"], ecxL2ServiceProfileSchemaNames["SpeedBand"]},
			RequiredWith: []string{ecxL2ServiceProfileSchemaNames["APIAvailable"]},
		},
		ecxL2ServiceProfileSchemaNames["TagType"]: {
			Type:     schema.TypeString,
			Optional: true,
		},
		ecxL2ServiceProfileSchemaNames["VlanSameAsPrimary"]: {
			Type:     schema.TypeBool,
			Required: true,
		},
		ecxL2ServiceProfileSchemaNames["Features"]: {
			Type:     schema.TypeSet,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					ecxL2ServiceProfileFeaturesSchemaNames["CloudReach"]: {
						Type:     schema.TypeBool,
						Required: true,
					},
					ecxL2ServiceProfileFeaturesSchemaNames["TestProfile"]: {
						Type:     schema.TypeBool,
						Required: true,
					},
				},
			},
		},
		ecxL2ServiceProfileSchemaNames["Port"]: {
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					ecxL2ServiceProfilePortSchemaNames["ID"]: {
						Type:     schema.TypeString,
						Required: true,
					},
					ecxL2ServiceProfilePortSchemaNames["MetroCode"]: {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		ecxL2ServiceProfileSchemaNames["SpeedBand"]: {
			Type:         schema.TypeSet,
			MinItems:     1,
			Optional:     true,
			AtLeastOneOf: []string{ecxL2ServiceProfileSchemaNames["SpeedFromAPI"], ecxL2ServiceProfileSchemaNames["SpeedBand"]},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					ecxL2ServiceProfileSpeedBandSchemaNames["Speed"]: {
						Type:     schema.TypeInt,
						Required: true,
					},
					ecxL2ServiceProfileSpeedBandSchemaNames["SpeedUnit"]: {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"MB", "GB"}, false),
					},
				},
			},
		},
	}
}

func resourceECXL2ServiceProfileCreate(d *schema.ResourceData, m interface{}) error {
	conf := m.(*Config)
	profile := createECXL2ServiceProfile(d)
	resp, err := conf.ecx.CreateL2ServiceProfile(*profile)
	if err != nil {
		return err
	}
	d.SetId(resp.UUID)
	return resourceECXL2ServiceProfileRead(d, m)
}

func resourceECXL2ServiceProfileRead(d *schema.ResourceData, m interface{}) error {
	conf := m.(*Config)
	profile, err := conf.ecx.GetL2ServiceProfile(d.Id())
	if err != nil {
		return err
	}
	if err := updateECXL2ServiceProfileResource(profile, d); err != nil {
		return err
	}
	return nil
}

func resourceECXL2ServiceProfileUpdate(d *schema.ResourceData, m interface{}) error {
	conf := m.(*Config)
	profile := createECXL2ServiceProfile(d)
	_, err := conf.ecx.UpdateL2ServiceProfile(*profile)
	if err != nil {
		return err
	}
	return resourceECXL2ServiceProfileRead(d, m)
}

func resourceECXL2ServiceProfileDelete(d *schema.ResourceData, m interface{}) error {
	conf := m.(*Config)
	if err := conf.ecx.DeleteL2ServiceProfile(d.Id()); err != nil {
		ecxRestErr, ok := err.(ecx.RestError)
		if ok {
			//IC-PROFILE-004 =  profile does not exist
			if hasECXErrorCode(ecxRestErr.Errors, "IC-PROFILE-004") {
				return nil
			}
		}
		return err
	}
	return nil
}

func createECXL2ServiceProfile(d *schema.ResourceData) *ecx.L2ServiceProfile {
	profile := ecx.L2ServiceProfile{}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["UUID"]); ok {
		profile.UUID = v.(string)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["State"]); ok {
		profile.State = v.(string)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["AlertPercentage"]); ok {
		profile.AlertPercentage = v.(float64)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["AllowCustomSpeed"]); ok {
		profile.AllowCustomSpeed = v.(bool)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["AllowOverSubscription"]); ok {
		profile.AllowOverSubscription = v.(bool)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["APIAvailable"]); ok {
		profile.APIAvailable = v.(bool)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["AuthKeyLabel"]); ok {
		profile.AuthKeyLabel = v.(string)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["ConnectionNameLabel"]); ok {
		profile.ConnectionNameLabel = v.(string)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["CTagLabel"]); ok {
		profile.CTagLabel = v.(string)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["Description"]); ok {
		profile.Description = v.(string)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["EnableAutoGenerateServiceKey"]); ok {
		profile.EnableAutoGenerateServiceKey = v.(bool)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["EquinixManagedPortAndVlan"]); ok {
		profile.EquinixManagedPortAndVlan = v.(bool)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["IntegrationID"]); ok {
		profile.IntegrationID = v.(string)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["Name"]); ok {
		profile.Name = v.(string)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["OnBandwidthThresholdNotification"]); ok {
		profile.OnBandwidthThresholdNotification = expandSetToStringList(v.(*schema.Set))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["OnProfileApprovalRejectNotification"]); ok {
		profile.OnProfileApprovalRejectNotification = expandSetToStringList(v.(*schema.Set))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["OnVcApprovalRejectionNotification"]); ok {
		profile.OnVcApprovalRejectionNotification = expandSetToStringList(v.(*schema.Set))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["OverSubscription"]); ok {
		profile.OverSubscription = v.(string)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["Private"]); ok {
		profile.Private = v.(bool)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["PrivateUserEmails"]); ok {
		profile.PrivateUserEmails = expandSetToStringList(v.(*schema.Set))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["RequiredRedundancy"]); ok {
		profile.RequiredRedundancy = v.(bool)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["SpeedFromAPI"]); ok {
		profile.SpeedFromAPI = v.(bool)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["TagType"]); ok {
		profile.TagType = v.(string)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["VlanSameAsPrimary"]); ok {
		profile.VlanSameAsPrimary = v.(bool)
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["Features"]); ok {
		featureSet := v.(*schema.Set)
		if featureSet.Len() > 0 {
			profile.Features = expandECXL2ServiceProfileFeatures(featureSet)[0]
		}
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["Port"]); ok {
		profile.Ports = expandECXL2ServiceProfilePorts(v.(*schema.Set))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["SpeedBand"]); ok {
		profile.SpeedBands = expandECXL2ServiceProfileSpeedBands(v.(*schema.Set))
	}
	return &profile
}

func updateECXL2ServiceProfileResource(profile *ecx.L2ServiceProfile, d *schema.ResourceData) error {
	if err := d.Set(ecxL2ServiceProfileSchemaNames["UUID"], profile.UUID); err != nil {
		return fmt.Errorf("error reading UUID: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["State"], profile.State); err != nil {
		return fmt.Errorf("error reading State: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["AlertPercentage"], profile.AlertPercentage); err != nil {
		return fmt.Errorf("error reading AlertPercentage: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["AllowCustomSpeed"], profile.AllowCustomSpeed); err != nil {
		return fmt.Errorf("error reading AllowCustomSpeed: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["AllowOverSubscription"], profile.AllowOverSubscription); err != nil {
		return fmt.Errorf("error reading AllowOverSubscription: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["APIAvailable"], profile.APIAvailable); err != nil {
		return fmt.Errorf("error reading APIAvailable: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["AuthKeyLabel"], profile.AuthKeyLabel); err != nil {
		return fmt.Errorf("error reading AuthKeyLabel: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["ConnectionNameLabel"], profile.ConnectionNameLabel); err != nil {
		return fmt.Errorf("error reading ConnectionNameLabel: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["CTagLabel"], profile.CTagLabel); err != nil {
		return fmt.Errorf("error reading CTagLabel: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["Description"], profile.Description); err != nil {
		return fmt.Errorf("error reading Description: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["EnableAutoGenerateServiceKey"], profile.EnableAutoGenerateServiceKey); err != nil {
		return fmt.Errorf("error reading EnableAutoGenerateServiceKey: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["EquinixManagedPortAndVlan"], profile.EquinixManagedPortAndVlan); err != nil {
		return fmt.Errorf("error reading EquinixManagedPortAndVlan: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["IntegrationID"], profile.IntegrationID); err != nil {
		return fmt.Errorf("error reading IntegrationID: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["Name"], profile.Name); err != nil {
		return fmt.Errorf("error reading Name: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["OnBandwidthThresholdNotification"], profile.OnBandwidthThresholdNotification); err != nil {
		return fmt.Errorf("error reading OnBandwidthThresholdNotification: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["OnProfileApprovalRejectNotification"], profile.OnProfileApprovalRejectNotification); err != nil {
		return fmt.Errorf("error reading OnProfileApprovalRejectNotification: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["OnVcApprovalRejectionNotification"], profile.OnVcApprovalRejectionNotification); err != nil {
		return fmt.Errorf("error reading OnVcApprovalRejectionNotification: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["OverSubscription"], profile.OverSubscription); err != nil {
		return fmt.Errorf("error reading OverSubscription: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["Private"], profile.Private); err != nil {
		return fmt.Errorf("error reading Private: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["PrivateUserEmails"], profile.PrivateUserEmails); err != nil {
		return fmt.Errorf("error reading PrivateUserEmails: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["RequiredRedundancy"], profile.RequiredRedundancy); err != nil {
		return fmt.Errorf("error reading RequiredRedundancy: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["SpeedFromAPI"], profile.SpeedFromAPI); err != nil {
		return fmt.Errorf("error reading SpeedFromAPI: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["TagType"], profile.TagType); err != nil {
		return fmt.Errorf("error reading TagType: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["VlanSameAsPrimary"], profile.VlanSameAsPrimary); err != nil {
		return fmt.Errorf("error reading VlanSameAsPrimary: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["Features"], flattenECXL2ServiceProfileFeatures(profile.Features)); err != nil {
		return fmt.Errorf("error reading Features: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["Port"], flattenECXL2ServiceProfilePorts(profile.Ports)); err != nil {
		return fmt.Errorf("error reading Port: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["SpeedBand"], flattenECXL2ServiceProfileSpeedBands(profile.SpeedBands)); err != nil {
		return fmt.Errorf("error reading SpeedBand: %s", err)
	}
	return nil
}

func flattenECXL2ServiceProfileFeatures(features ecx.L2ServiceProfileFeatures) interface{} {
	transformed := make(map[string]interface{})
	transformed[ecxL2ServiceProfileFeaturesSchemaNames["CloudReach"]] = features.CloudReach
	transformed[ecxL2ServiceProfileFeaturesSchemaNames["TestProfile"]] = features.TestProfile
	return []map[string]interface{}{transformed}
}

func flattenECXL2ServiceProfilePorts(ports []ecx.L2ServiceProfilePort) interface{} {
	transformed := make([]interface{}, 0, len(ports))
	for _, port := range ports {
		transformed = append(transformed, map[string]interface{}{
			ecxL2ServiceProfilePortSchemaNames["ID"]:        port.ID,
			ecxL2ServiceProfilePortSchemaNames["MetroCode"]: port.MetroCode,
		})
	}
	return transformed
}

func flattenECXL2ServiceProfileSpeedBands(bands []ecx.L2ServiceProfileSpeedBand) interface{} {
	transformed := make([]interface{}, 0, len(bands))
	for _, band := range bands {
		transformed = append(transformed, map[string]interface{}{
			ecxL2ServiceProfileSpeedBandSchemaNames["Speed"]:     band.Speed,
			ecxL2ServiceProfileSpeedBandSchemaNames["SpeedUnit"]: band.SpeedUnit,
		})
	}
	return transformed
}

func expandECXL2ServiceProfileFeatures(features *schema.Set) []ecx.L2ServiceProfileFeatures {
	transformed := make([]ecx.L2ServiceProfileFeatures, 0, features.Len())
	for _, feature := range features.List() {
		featureMap := feature.(map[string]interface{})
		transformed = append(transformed, ecx.L2ServiceProfileFeatures{
			CloudReach:  featureMap[ecxL2ServiceProfileFeaturesSchemaNames["CloudReach"]].(bool),
			TestProfile: featureMap[ecxL2ServiceProfileFeaturesSchemaNames["TestProfile"]].(bool),
		})
	}
	return transformed
}

func expandECXL2ServiceProfilePorts(ports *schema.Set) []ecx.L2ServiceProfilePort {
	transformed := make([]ecx.L2ServiceProfilePort, 0, ports.Len())
	for _, port := range ports.List() {
		portMap := port.(map[string]interface{})
		transformed = append(transformed, ecx.L2ServiceProfilePort{
			ID:        portMap[ecxL2ServiceProfilePortSchemaNames["ID"]].(string),
			MetroCode: portMap[ecxL2ServiceProfilePortSchemaNames["MetroCode"]].(string),
		})
	}
	return transformed
}

func expandECXL2ServiceProfileSpeedBands(bands *schema.Set) []ecx.L2ServiceProfileSpeedBand {
	transformed := make([]ecx.L2ServiceProfileSpeedBand, 0, bands.Len())
	for _, band := range bands.List() {
		bandMap := band.(map[string]interface{})
		transformed = append(transformed, ecx.L2ServiceProfileSpeedBand{
			Speed:     bandMap[ecxL2ServiceProfileSpeedBandSchemaNames["Speed"]].(int),
			SpeedUnit: bandMap[ecxL2ServiceProfileSpeedBandSchemaNames["SpeedUnit"]].(string),
		})
	}
	return transformed
}
