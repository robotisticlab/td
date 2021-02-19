// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// AccountGetNotifySettingsRequest represents TL type `account.getNotifySettings#12b3ad31`.
// Gets current notification settings for a given user/group, from all users/all groups.
//
// See https://core.telegram.org/method/account.getNotifySettings for reference.
type AccountGetNotifySettingsRequest struct {
	// Notification source
	Peer InputNotifyPeerClass
}

// AccountGetNotifySettingsRequestTypeID is TL type id of AccountGetNotifySettingsRequest.
const AccountGetNotifySettingsRequestTypeID = 0x12b3ad31

func (g *AccountGetNotifySettingsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetNotifySettingsRequest) String() string {
	if g == nil {
		return "AccountGetNotifySettingsRequest(nil)"
	}
	type Alias AccountGetNotifySettingsRequest
	return fmt.Sprintf("AccountGetNotifySettingsRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetNotifySettingsRequest from given interface.
func (g *AccountGetNotifySettingsRequest) FillFrom(from interface {
	GetPeer() (value InputNotifyPeerClass)
}) {
	g.Peer = from.GetPeer()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *AccountGetNotifySettingsRequest) TypeID() uint32 {
	return AccountGetNotifySettingsRequestTypeID
}

// Encode implements bin.Encoder.
func (g *AccountGetNotifySettingsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getNotifySettings#12b3ad31 as nil")
	}
	b.PutID(AccountGetNotifySettingsRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode account.getNotifySettings#12b3ad31: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getNotifySettings#12b3ad31: field peer: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *AccountGetNotifySettingsRequest) GetPeer() (value InputNotifyPeerClass) {
	return g.Peer
}

// Decode implements bin.Decoder.
func (g *AccountGetNotifySettingsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getNotifySettings#12b3ad31 to nil")
	}
	if err := b.ConsumeID(AccountGetNotifySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getNotifySettings#12b3ad31: %w", err)
	}
	{
		value, err := DecodeInputNotifyPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getNotifySettings#12b3ad31: field peer: %w", err)
		}
		g.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetNotifySettingsRequest.
var (
	_ bin.Encoder = &AccountGetNotifySettingsRequest{}
	_ bin.Decoder = &AccountGetNotifySettingsRequest{}
)

// AccountGetNotifySettings invokes method account.getNotifySettings#12b3ad31 returning error if any.
// Gets current notification settings for a given user/group, from all users/all groups.
//
// Possible errors:
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/account.getNotifySettings for reference.
func (c *Client) AccountGetNotifySettings(ctx context.Context, peer InputNotifyPeerClass) (*PeerNotifySettings, error) {
	var result PeerNotifySettings

	request := &AccountGetNotifySettingsRequest{
		Peer: peer,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
