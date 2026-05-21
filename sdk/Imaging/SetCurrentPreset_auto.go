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

// Call_SetCurrentPreset forwards the call to dev.CallMethod() then parses the payload of the reply as a SetCurrentPresetResponse.
func Call_SetCurrentPreset(ctx context.Context, dev *onvif.Device, request Imaging.SetCurrentPreset) (Imaging.SetCurrentPresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetCurrentPresetResponse Imaging.SetCurrentPresetResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethodContext(ctx, request); err != nil {
		return reply.Body.SetCurrentPresetResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetCurrentPreset")
		return reply.Body.SetCurrentPresetResponse, errors.Annotate(err, "reply")
	}
}
