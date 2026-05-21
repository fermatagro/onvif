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

// Call_Move forwards the call to dev.CallMethod() then parses the payload of the reply as a MoveResponse.
func Call_Move(ctx context.Context, dev *onvif.Device, request Imaging.Move) (Imaging.MoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			MoveResponse Imaging.MoveResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethodContext(ctx, request); err != nil {
		return reply.Body.MoveResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "Move")
		return reply.Body.MoveResponse, errors.Annotate(err, "reply")
	}
}
