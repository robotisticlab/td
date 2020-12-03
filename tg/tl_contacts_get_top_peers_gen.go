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

// ContactsGetTopPeersRequest represents TL type `contacts.getTopPeers#d4982db5`.
//
// See https://core.telegram.org/method/contacts.getTopPeers for reference.
type ContactsGetTopPeersRequest struct {
	// Flags field of ContactsGetTopPeersRequest.
	Flags bin.Fields
	// Correspondents field of ContactsGetTopPeersRequest.
	Correspondents bool
	// BotsPm field of ContactsGetTopPeersRequest.
	BotsPm bool
	// BotsInline field of ContactsGetTopPeersRequest.
	BotsInline bool
	// PhoneCalls field of ContactsGetTopPeersRequest.
	PhoneCalls bool
	// ForwardUsers field of ContactsGetTopPeersRequest.
	ForwardUsers bool
	// ForwardChats field of ContactsGetTopPeersRequest.
	ForwardChats bool
	// Groups field of ContactsGetTopPeersRequest.
	Groups bool
	// Channels field of ContactsGetTopPeersRequest.
	Channels bool
	// Offset field of ContactsGetTopPeersRequest.
	Offset int
	// Limit field of ContactsGetTopPeersRequest.
	Limit int
	// Hash field of ContactsGetTopPeersRequest.
	Hash int
}

// ContactsGetTopPeersRequestTypeID is TL type id of ContactsGetTopPeersRequest.
const ContactsGetTopPeersRequestTypeID = 0xd4982db5

// Encode implements bin.Encoder.
func (g *ContactsGetTopPeersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getTopPeers#d4982db5 as nil")
	}
	b.PutID(ContactsGetTopPeersRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.getTopPeers#d4982db5: field flags: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	b.PutInt(g.Hash)
	return nil
}

// SetCorrespondents sets value of Correspondents conditional field.
func (g *ContactsGetTopPeersRequest) SetCorrespondents(value bool) {
	if value {
		g.Flags.Set(0)
	} else {
		g.Flags.Unset(0)
	}
}

// SetBotsPm sets value of BotsPm conditional field.
func (g *ContactsGetTopPeersRequest) SetBotsPm(value bool) {
	if value {
		g.Flags.Set(1)
	} else {
		g.Flags.Unset(1)
	}
}

// SetBotsInline sets value of BotsInline conditional field.
func (g *ContactsGetTopPeersRequest) SetBotsInline(value bool) {
	if value {
		g.Flags.Set(2)
	} else {
		g.Flags.Unset(2)
	}
}

// SetPhoneCalls sets value of PhoneCalls conditional field.
func (g *ContactsGetTopPeersRequest) SetPhoneCalls(value bool) {
	if value {
		g.Flags.Set(3)
	} else {
		g.Flags.Unset(3)
	}
}

// SetForwardUsers sets value of ForwardUsers conditional field.
func (g *ContactsGetTopPeersRequest) SetForwardUsers(value bool) {
	if value {
		g.Flags.Set(4)
	} else {
		g.Flags.Unset(4)
	}
}

// SetForwardChats sets value of ForwardChats conditional field.
func (g *ContactsGetTopPeersRequest) SetForwardChats(value bool) {
	if value {
		g.Flags.Set(5)
	} else {
		g.Flags.Unset(5)
	}
}

// SetGroups sets value of Groups conditional field.
func (g *ContactsGetTopPeersRequest) SetGroups(value bool) {
	if value {
		g.Flags.Set(10)
	} else {
		g.Flags.Unset(10)
	}
}

// SetChannels sets value of Channels conditional field.
func (g *ContactsGetTopPeersRequest) SetChannels(value bool) {
	if value {
		g.Flags.Set(15)
	} else {
		g.Flags.Unset(15)
	}
}

// Decode implements bin.Decoder.
func (g *ContactsGetTopPeersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getTopPeers#d4982db5 to nil")
	}
	if err := b.ConsumeID(ContactsGetTopPeersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: field flags: %w", err)
		}
	}
	g.Correspondents = g.Flags.Has(0)
	g.BotsPm = g.Flags.Has(1)
	g.BotsInline = g.Flags.Has(2)
	g.PhoneCalls = g.Flags.Has(3)
	g.ForwardUsers = g.Flags.Has(4)
	g.ForwardChats = g.Flags.Has(5)
	g.Groups = g.Flags.Has(10)
	g.Channels = g.Flags.Has(15)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetTopPeersRequest.
var (
	_ bin.Encoder = &ContactsGetTopPeersRequest{}
	_ bin.Decoder = &ContactsGetTopPeersRequest{}
)

// ContactsGetTopPeers invokes method contacts.getTopPeers#d4982db5 returning error if any.
//
// See https://core.telegram.org/method/contacts.getTopPeers for reference.
func (c *Client) ContactsGetTopPeers(ctx context.Context, request *ContactsGetTopPeersRequest) (ContactsTopPeersClass, error) {
	var result ContactsTopPeersBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.TopPeers, nil
}
