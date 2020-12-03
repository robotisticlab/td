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

// MessagesGetOldFeaturedStickersRequest represents TL type `messages.getOldFeaturedStickers#5fe7025b`.
//
// See https://core.telegram.org/method/messages.getOldFeaturedStickers for reference.
type MessagesGetOldFeaturedStickersRequest struct {
	// Offset field of MessagesGetOldFeaturedStickersRequest.
	Offset int
	// Limit field of MessagesGetOldFeaturedStickersRequest.
	Limit int
	// Hash field of MessagesGetOldFeaturedStickersRequest.
	Hash int
}

// MessagesGetOldFeaturedStickersRequestTypeID is TL type id of MessagesGetOldFeaturedStickersRequest.
const MessagesGetOldFeaturedStickersRequestTypeID = 0x5fe7025b

// Encode implements bin.Encoder.
func (g *MessagesGetOldFeaturedStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getOldFeaturedStickers#5fe7025b as nil")
	}
	b.PutID(MessagesGetOldFeaturedStickersRequestTypeID)
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetOldFeaturedStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getOldFeaturedStickers#5fe7025b to nil")
	}
	if err := b.ConsumeID(MessagesGetOldFeaturedStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getOldFeaturedStickers#5fe7025b: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getOldFeaturedStickers#5fe7025b: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getOldFeaturedStickers#5fe7025b: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getOldFeaturedStickers#5fe7025b: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetOldFeaturedStickersRequest.
var (
	_ bin.Encoder = &MessagesGetOldFeaturedStickersRequest{}
	_ bin.Decoder = &MessagesGetOldFeaturedStickersRequest{}
)

// MessagesGetOldFeaturedStickers invokes method messages.getOldFeaturedStickers#5fe7025b returning error if any.
//
// See https://core.telegram.org/method/messages.getOldFeaturedStickers for reference.
func (c *Client) MessagesGetOldFeaturedStickers(ctx context.Context, request *MessagesGetOldFeaturedStickersRequest) (MessagesFeaturedStickersClass, error) {
	var result MessagesFeaturedStickersBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.FeaturedStickers, nil
}
