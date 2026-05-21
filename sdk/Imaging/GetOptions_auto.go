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

// Call_GetOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetOptionsResponse.
func Call_GetOptions(ctx context.Context, dev *onvif.Device, request Imaging.GetOptions) (Imaging.GetOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOptionsResponse Imaging.GetOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethodContext(ctx, request); err != nil {
		return reply.Body.GetOptionsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetOptions")
		return reply.Body.GetOptionsResponse, errors.Annotate(err, "reply")
	}
}
