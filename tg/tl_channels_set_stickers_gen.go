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

// ChannelsSetStickersRequest represents TL type `channels.setStickers#ea8ca4f9`.
//
// See https://core.telegram.org/method/channels.setStickers for reference.
type ChannelsSetStickersRequest struct {
	// Channel field of ChannelsSetStickersRequest.
	Channel InputChannelClass
	// Stickerset field of ChannelsSetStickersRequest.
	Stickerset InputStickerSetClass
}

// ChannelsSetStickersRequestTypeID is TL type id of ChannelsSetStickersRequest.
const ChannelsSetStickersRequestTypeID = 0xea8ca4f9

// Encode implements bin.Encoder.
func (s *ChannelsSetStickersRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode channels.setStickers#ea8ca4f9 as nil")
	}
	b.PutID(ChannelsSetStickersRequestTypeID)
	if s.Channel == nil {
		return fmt.Errorf("unable to encode channels.setStickers#ea8ca4f9: field channel is nil")
	}
	if err := s.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.setStickers#ea8ca4f9: field channel: %w", err)
	}
	if s.Stickerset == nil {
		return fmt.Errorf("unable to encode channels.setStickers#ea8ca4f9: field stickerset is nil")
	}
	if err := s.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.setStickers#ea8ca4f9: field stickerset: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *ChannelsSetStickersRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode channels.setStickers#ea8ca4f9 to nil")
	}
	if err := b.ConsumeID(ChannelsSetStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.setStickers#ea8ca4f9: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.setStickers#ea8ca4f9: field channel: %w", err)
		}
		s.Channel = value
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.setStickers#ea8ca4f9: field stickerset: %w", err)
		}
		s.Stickerset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsSetStickersRequest.
var (
	_ bin.Encoder = &ChannelsSetStickersRequest{}
	_ bin.Decoder = &ChannelsSetStickersRequest{}
)

// ChannelsSetStickers invokes method channels.setStickers#ea8ca4f9 returning error if any.
//
// See https://core.telegram.org/method/channels.setStickers for reference.
func (c *Client) ChannelsSetStickers(ctx context.Context, request *ChannelsSetStickersRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
