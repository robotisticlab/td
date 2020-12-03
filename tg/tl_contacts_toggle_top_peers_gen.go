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

// ContactsToggleTopPeersRequest represents TL type `contacts.toggleTopPeers#8514bdda`.
//
// See https://core.telegram.org/method/contacts.toggleTopPeers for reference.
type ContactsToggleTopPeersRequest struct {
	// Enabled field of ContactsToggleTopPeersRequest.
	Enabled bool
}

// ContactsToggleTopPeersRequestTypeID is TL type id of ContactsToggleTopPeersRequest.
const ContactsToggleTopPeersRequestTypeID = 0x8514bdda

// Encode implements bin.Encoder.
func (t *ContactsToggleTopPeersRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.toggleTopPeers#8514bdda as nil")
	}
	b.PutID(ContactsToggleTopPeersRequestTypeID)
	b.PutBool(t.Enabled)
	return nil
}

// Decode implements bin.Decoder.
func (t *ContactsToggleTopPeersRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.toggleTopPeers#8514bdda to nil")
	}
	if err := b.ConsumeID(ContactsToggleTopPeersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.toggleTopPeers#8514bdda: %w", err)
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.toggleTopPeers#8514bdda: field enabled: %w", err)
		}
		t.Enabled = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsToggleTopPeersRequest.
var (
	_ bin.Encoder = &ContactsToggleTopPeersRequest{}
	_ bin.Decoder = &ContactsToggleTopPeersRequest{}
)

// ContactsToggleTopPeers invokes method contacts.toggleTopPeers#8514bdda returning error if any.
//
// See https://core.telegram.org/method/contacts.toggleTopPeers for reference.
func (c *Client) ContactsToggleTopPeers(ctx context.Context, request *ContactsToggleTopPeersRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
