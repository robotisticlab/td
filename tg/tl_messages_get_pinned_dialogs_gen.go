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

// MessagesGetPinnedDialogsRequest represents TL type `messages.getPinnedDialogs#d6b94df2`.
//
// See https://core.telegram.org/method/messages.getPinnedDialogs for reference.
type MessagesGetPinnedDialogsRequest struct {
	// FolderID field of MessagesGetPinnedDialogsRequest.
	FolderID int
}

// MessagesGetPinnedDialogsRequestTypeID is TL type id of MessagesGetPinnedDialogsRequest.
const MessagesGetPinnedDialogsRequestTypeID = 0xd6b94df2

// Encode implements bin.Encoder.
func (g *MessagesGetPinnedDialogsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPinnedDialogs#d6b94df2 as nil")
	}
	b.PutID(MessagesGetPinnedDialogsRequestTypeID)
	b.PutInt(g.FolderID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetPinnedDialogsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPinnedDialogs#d6b94df2 to nil")
	}
	if err := b.ConsumeID(MessagesGetPinnedDialogsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getPinnedDialogs#d6b94df2: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPinnedDialogs#d6b94df2: field folder_id: %w", err)
		}
		g.FolderID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetPinnedDialogsRequest.
var (
	_ bin.Encoder = &MessagesGetPinnedDialogsRequest{}
	_ bin.Decoder = &MessagesGetPinnedDialogsRequest{}
)

// MessagesGetPinnedDialogs invokes method messages.getPinnedDialogs#d6b94df2 returning error if any.
//
// See https://core.telegram.org/method/messages.getPinnedDialogs for reference.
func (c *Client) MessagesGetPinnedDialogs(ctx context.Context, request *MessagesGetPinnedDialogsRequest) (*MessagesPeerDialogs, error) {
	var result MessagesPeerDialogs
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
