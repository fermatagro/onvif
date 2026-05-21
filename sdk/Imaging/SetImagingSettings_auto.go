// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package Imaging

import (
	"context"
	"github.com/juju/errors"
	"github.com/fermatagro/onvif"
	"github.com/fermatagro/onvif/sdk"
	"github.com/fermatagro/onvif/Imaging"
)

// Call_SetImagingSettings forwards the call to dev.CallMethod() then parses the payload of the reply as a SetImagingSettingsResponse.
func Call_SetImagingSettings(ctx context.Context, dev *onvif.Device, request Imaging.SetImagingSettings) (Imaging.SetImagingSettingsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetImagingSettingsResponse Imaging.SetImagingSettingsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethodContext(ctx, request); err != nil {
		return reply.Body.SetImagingSettingsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetImagingSettings")
		return reply.Body.SetImagingSettingsResponse, errors.Annotate(err, "reply")
	}
}
