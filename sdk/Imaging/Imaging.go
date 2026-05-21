package Imaging

//go:generate go run github.com/fermatagro/onvif/sdk/codegen Imaging Imaging GetServiceCapabilities
//go:generate go run github.com/fermatagro/onvif/sdk/codegen Imaging Imaging GetImagingSettings
//go:generate go run github.com/fermatagro/onvif/sdk/codegen Imaging Imaging SetImagingSettings
//go:generate go run github.com/fermatagro/onvif/sdk/codegen Imaging Imaging GetOptions
//go:generate go run github.com/fermatagro/onvif/sdk/codegen Imaging Imaging Move
//go:generate go run github.com/fermatagro/onvif/sdk/codegen Imaging Imaging GetMoveOptions
//go:generate go run github.com/fermatagro/onvif/sdk/codegen Imaging Imaging Stop
//go:generate go run github.com/fermatagro/onvif/sdk/codegen Imaging Imaging GetStatus
//go:generate go run github.com/fermatagro/onvif/sdk/codegen Imaging Imaging GetPresets
//go:generate go run github.com/fermatagro/onvif/sdk/codegen Imaging Imaging GetCurrentPreset
//go:generate go run github.com/fermatagro/onvif/sdk/codegen Imaging Imaging SetCurrentPreset
