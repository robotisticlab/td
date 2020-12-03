// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// HelpGetNearestDcRequest represents TL type `help.getNearestDc#1fb33026`.
//
// See https://core.telegram.org/method/help.getNearestDc for reference.
type HelpGetNearestDcRequest struct {
}

// HelpGetNearestDcRequestTypeID is TL type id of HelpGetNearestDcRequest.
const HelpGetNearestDcRequestTypeID = 0x1fb33026

// Encode implements bin.Encoder.
func (g *HelpGetNearestDcRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getNearestDc#1fb33026 as nil")
	}
	b.PutID(HelpGetNearestDcRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetNearestDcRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getNearestDc#1fb33026 to nil")
	}
	if err := b.ConsumeID(HelpGetNearestDcRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getNearestDc#1fb33026: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetNearestDcRequest.
var (
	_ bin.Encoder = &HelpGetNearestDcRequest{}
	_ bin.Decoder = &HelpGetNearestDcRequest{}
)

// HelpGetNearestDc invokes method help.getNearestDc#1fb33026 returning error if any.
//
// See https://core.telegram.org/method/help.getNearestDc for reference.
func (c *Client) HelpGetNearestDc(ctx context.Context, request *HelpGetNearestDcRequest) (*NearestDc, error) {
	var result NearestDc
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
