package Imaging

import (
	"github.com/fermatagro/onvif/xsd"
	"github.com/fermatagro/onvif/xsd/onvif"
)

type GetServiceCapabilities struct {
	XMLName string `xml:"timg:GetServiceCapabilities"`
}

type GetImagingSettings struct {
	XMLName          string               `xml:"timg:GetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type SetImagingSettings struct {
	XMLName          string                  `xml:"timg:SetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken    `xml:"timg:VideoSourceToken"`
	ImagingSettings  onvif.ImagingSettings20 `xml:"timg:ImagingSettings"`
	ForcePersistence xsd.Boolean             `xml:"timg:ForcePersistence"`
}

type GetOptions struct {
	XMLName          string               `xml:"timg:GetOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type Move struct {
	XMLName          string               `xml:"timg:Move"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	Focus            onvif.FocusMove      `xml:"timg:Focus"`
}

type GetMoveOptions struct {
	XMLName          string               `xml:"timg:GetMoveOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type Stop struct {
	XMLName          string               `xml:"timg:Stop"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetStatus struct {
	XMLName          string               `xml:"timg:GetStatus"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetPresets struct {
	XMLName          string               `xml:"timg:GetPresets"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetCurrentPreset struct {
	XMLName          string               `xml:"timg:GetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type SetCurrentPreset struct {
	XMLName          string               `xml:"timg:SetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	PresetToken      onvif.ReferenceToken `xml:"timg:PresetToken"`
}

// Capabilities — timg:Capabilities
type Capabilities struct {
	ImageStabilization xsd.Boolean `xml:"ImageStabilization,attr"`
	Presets            xsd.Boolean `xml:"Presets,attr"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

type GetImagingSettingsResponse struct {
	ImagingSettings onvif.ImagingSettings20
}

type SetImagingSettingsResponse struct {
}

type GetOptionsResponse struct {
	ImagingOptions onvif.ImagingOptions20
}

type MoveResponse struct {
}

type GetMoveOptionsResponse struct {
	MoveOptions onvif.MoveOptions20
}

type StopResponse struct {
}

type GetStatusResponse struct {
	Status onvif.ImagingStatus20
}

type GetPresetsResponse struct {
	Preset []onvif.ImagingPreset
}

type GetCurrentPresetResponse struct {
	Preset *onvif.ImagingPreset `xml:"Preset,omitempty"`
}

type SetCurrentPresetResponse struct {
}
