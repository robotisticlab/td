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

// MessagesGetChatsRequest represents TL type `messages.getChats#3c6aa187`.
//
// See https://core.telegram.org/method/messages.getChats for reference.
type MessagesGetChatsRequest struct {
	// ID field of MessagesGetChatsRequest.
	ID []int
}

// MessagesGetChatsRequestTypeID is TL type id of MessagesGetChatsRequest.
const MessagesGetChatsRequestTypeID = 0x3c6aa187

// Encode implements bin.Encoder.
func (g *MessagesGetChatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getChats#3c6aa187 as nil")
	}
	b.PutID(MessagesGetChatsRequestTypeID)
	b.PutVectorHeader(len(g.ID))
	for _, v := range g.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetChatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getChats#3c6aa187 to nil")
	}
	if err := b.ConsumeID(MessagesGetChatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getChats#3c6aa187: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getChats#3c6aa187: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.getChats#3c6aa187: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetChatsRequest.
var (
	_ bin.Encoder = &MessagesGetChatsRequest{}
	_ bin.Decoder = &MessagesGetChatsRequest{}
)

// MessagesGetChats invokes method messages.getChats#3c6aa187 returning error if any.
//
// See https://core.telegram.org/method/messages.getChats for reference.
func (c *Client) MessagesGetChats(ctx context.Context, request *MessagesGetChatsRequest) (MessagesChatsClass, error) {
	var result MessagesChatsBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
