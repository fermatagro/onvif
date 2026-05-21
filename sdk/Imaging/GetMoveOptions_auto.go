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

// Call_GetMoveOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetMoveOptionsResponse.
func Call_GetMoveOptions(ctx context.Context, dev *onvif.Device, request Imaging.GetMoveOptions) (Imaging.GetMoveOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMoveOptionsResponse Imaging.GetMoveOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethodContext(ctx, request); err != nil {
		return reply.Body.GetMoveOptionsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetMoveOptions")
		return reply.Body.GetMoveOptionsResponse, errors.Annotate(err, "reply")
	}
}
