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

// Call_GetPresets forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPresetsResponse.
func Call_GetPresets(ctx context.Context, dev *onvif.Device, request Imaging.GetPresets) (Imaging.GetPresetsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetsResponse Imaging.GetPresetsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethodContext(ctx, request); err != nil {
		return reply.Body.GetPresetsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetPresets")
		return reply.Body.GetPresetsResponse, errors.Annotate(err, "reply")
	}
}
