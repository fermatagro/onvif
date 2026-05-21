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

// Call_GetImagingSettings forwards the call to dev.CallMethod() then parses the payload of the reply as a GetImagingSettingsResponse.
func Call_GetImagingSettings(ctx context.Context, dev *onvif.Device, request Imaging.GetImagingSettings) (Imaging.GetImagingSettingsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetImagingSettingsResponse Imaging.GetImagingSettingsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethodContext(ctx, request); err != nil {
		return reply.Body.GetImagingSettingsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetImagingSettings")
		return reply.Body.GetImagingSettingsResponse, errors.Annotate(err, "reply")
	}
}
