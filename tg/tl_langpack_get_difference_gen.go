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

// LangpackGetDifferenceRequest represents TL type `langpack.getDifference#cd984aa5`.
//
// See https://core.telegram.org/method/langpack.getDifference for reference.
type LangpackGetDifferenceRequest struct {
	// LangPack field of LangpackGetDifferenceRequest.
	LangPack string
	// LangCode field of LangpackGetDifferenceRequest.
	LangCode string
	// FromVersion field of LangpackGetDifferenceRequest.
	FromVersion int
}

// LangpackGetDifferenceRequestTypeID is TL type id of LangpackGetDifferenceRequest.
const LangpackGetDifferenceRequestTypeID = 0xcd984aa5

// Encode implements bin.Encoder.
func (g *LangpackGetDifferenceRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode langpack.getDifference#cd984aa5 as nil")
	}
	b.PutID(LangpackGetDifferenceRequestTypeID)
	b.PutString(g.LangPack)
	b.PutString(g.LangCode)
	b.PutInt(g.FromVersion)
	return nil
}

// Decode implements bin.Decoder.
func (g *LangpackGetDifferenceRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode langpack.getDifference#cd984aa5 to nil")
	}
	if err := b.ConsumeID(LangpackGetDifferenceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode langpack.getDifference#cd984aa5: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getDifference#cd984aa5: field lang_pack: %w", err)
		}
		g.LangPack = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getDifference#cd984aa5: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getDifference#cd984aa5: field from_version: %w", err)
		}
		g.FromVersion = value
	}
	return nil
}

// Ensuring interfaces in compile-time for LangpackGetDifferenceRequest.
var (
	_ bin.Encoder = &LangpackGetDifferenceRequest{}
	_ bin.Decoder = &LangpackGetDifferenceRequest{}
)

// LangpackGetDifference invokes method langpack.getDifference#cd984aa5 returning error if any.
//
// See https://core.telegram.org/method/langpack.getDifference for reference.
func (c *Client) LangpackGetDifference(ctx context.Context, request *LangpackGetDifferenceRequest) (*LangPackDifference, error) {
	var result LangPackDifference
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
