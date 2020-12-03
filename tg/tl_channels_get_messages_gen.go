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

// ChannelsGetMessagesRequest represents TL type `channels.getMessages#ad8c9a23`.
//
// See https://core.telegram.org/method/channels.getMessages for reference.
type ChannelsGetMessagesRequest struct {
	// Channel field of ChannelsGetMessagesRequest.
	Channel InputChannelClass
	// ID field of ChannelsGetMessagesRequest.
	ID []InputMessageClass
}

// ChannelsGetMessagesRequestTypeID is TL type id of ChannelsGetMessagesRequest.
const ChannelsGetMessagesRequestTypeID = 0xad8c9a23

// Encode implements bin.Encoder.
func (g *ChannelsGetMessagesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getMessages#ad8c9a23 as nil")
	}
	b.PutID(ChannelsGetMessagesRequestTypeID)
	if g.Channel == nil {
		return fmt.Errorf("unable to encode channels.getMessages#ad8c9a23: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getMessages#ad8c9a23: field channel: %w", err)
	}
	b.PutVectorHeader(len(g.ID))
	for idx, v := range g.ID {
		if v == nil {
			return fmt.Errorf("unable to encode channels.getMessages#ad8c9a23: field id element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.getMessages#ad8c9a23: field id element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetMessagesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getMessages#ad8c9a23 to nil")
	}
	if err := b.ConsumeID(ChannelsGetMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getMessages#ad8c9a23: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getMessages#ad8c9a23: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getMessages#ad8c9a23: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.getMessages#ad8c9a23: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetMessagesRequest.
var (
	_ bin.Encoder = &ChannelsGetMessagesRequest{}
	_ bin.Decoder = &ChannelsGetMessagesRequest{}
)

// ChannelsGetMessages invokes method channels.getMessages#ad8c9a23 returning error if any.
//
// See https://core.telegram.org/method/channels.getMessages for reference.
func (c *Client) ChannelsGetMessages(ctx context.Context, request *ChannelsGetMessagesRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
