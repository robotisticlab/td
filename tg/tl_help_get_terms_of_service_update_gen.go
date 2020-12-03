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

// HelpGetTermsOfServiceUpdateRequest represents TL type `help.getTermsOfServiceUpdate#2ca51fd1`.
//
// See https://core.telegram.org/method/help.getTermsOfServiceUpdate for reference.
type HelpGetTermsOfServiceUpdateRequest struct {
}

// HelpGetTermsOfServiceUpdateRequestTypeID is TL type id of HelpGetTermsOfServiceUpdateRequest.
const HelpGetTermsOfServiceUpdateRequestTypeID = 0x2ca51fd1

// Encode implements bin.Encoder.
func (g *HelpGetTermsOfServiceUpdateRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getTermsOfServiceUpdate#2ca51fd1 as nil")
	}
	b.PutID(HelpGetTermsOfServiceUpdateRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetTermsOfServiceUpdateRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getTermsOfServiceUpdate#2ca51fd1 to nil")
	}
	if err := b.ConsumeID(HelpGetTermsOfServiceUpdateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getTermsOfServiceUpdate#2ca51fd1: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetTermsOfServiceUpdateRequest.
var (
	_ bin.Encoder = &HelpGetTermsOfServiceUpdateRequest{}
	_ bin.Decoder = &HelpGetTermsOfServiceUpdateRequest{}
)

// HelpGetTermsOfServiceUpdate invokes method help.getTermsOfServiceUpdate#2ca51fd1 returning error if any.
//
// See https://core.telegram.org/method/help.getTermsOfServiceUpdate for reference.
func (c *Client) HelpGetTermsOfServiceUpdate(ctx context.Context, request *HelpGetTermsOfServiceUpdateRequest) (HelpTermsOfServiceUpdateClass, error) {
	var result HelpTermsOfServiceUpdateBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.TermsOfServiceUpdate, nil
}
