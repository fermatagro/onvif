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

// Call_GetCurrentPreset forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCurrentPresetResponse.
func Call_GetCurrentPreset(ctx context.Context, dev *onvif.Device, request Imaging.GetCurrentPreset) (Imaging.GetCurrentPresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCurrentPresetResponse Imaging.GetCurrentPresetResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethodContext(ctx, request); err != nil {
		return reply.Body.GetCurrentPresetResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCurrentPreset")
		return reply.Body.GetCurrentPresetResponse, errors.Annotate(err, "reply")
	}
}
