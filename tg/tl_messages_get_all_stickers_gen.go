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

// MessagesGetAllStickersRequest represents TL type `messages.getAllStickers#1c9618b1`.
//
// See https://core.telegram.org/method/messages.getAllStickers for reference.
type MessagesGetAllStickersRequest struct {
	// Hash field of MessagesGetAllStickersRequest.
	Hash int
}

// MessagesGetAllStickersRequestTypeID is TL type id of MessagesGetAllStickersRequest.
const MessagesGetAllStickersRequestTypeID = 0x1c9618b1

// Encode implements bin.Encoder.
func (g *MessagesGetAllStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAllStickers#1c9618b1 as nil")
	}
	b.PutID(MessagesGetAllStickersRequestTypeID)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetAllStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAllStickers#1c9618b1 to nil")
	}
	if err := b.ConsumeID(MessagesGetAllStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getAllStickers#1c9618b1: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getAllStickers#1c9618b1: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetAllStickersRequest.
var (
	_ bin.Encoder = &MessagesGetAllStickersRequest{}
	_ bin.Decoder = &MessagesGetAllStickersRequest{}
)

// MessagesGetAllStickers invokes method messages.getAllStickers#1c9618b1 returning error if any.
//
// See https://core.telegram.org/method/messages.getAllStickers for reference.
func (c *Client) MessagesGetAllStickers(ctx context.Context, request *MessagesGetAllStickersRequest) (MessagesAllStickersClass, error) {
	var result MessagesAllStickersBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.AllStickers, nil
}
