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

// MessagesGetPollResultsRequest represents TL type `messages.getPollResults#73bb643b`.
// Get poll results
//
// See https://core.telegram.org/method/messages.getPollResults for reference.
type MessagesGetPollResultsRequest struct {
	// Peer where the poll was found
	Peer InputPeerClass
	// Message ID of poll message
	MsgID int
}

// MessagesGetPollResultsRequestTypeID is TL type id of MessagesGetPollResultsRequest.
const MessagesGetPollResultsRequestTypeID = 0x73bb643b

func (g *MessagesGetPollResultsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetPollResultsRequest) String() string {
	if g == nil {
		return "MessagesGetPollResultsRequest(nil)"
	}
	type Alias MessagesGetPollResultsRequest
	return fmt.Sprintf("MessagesGetPollResultsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetPollResultsRequest from given interface.
func (g *MessagesGetPollResultsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
}) {
	g.Peer = from.GetPeer()
	g.MsgID = from.GetMsgID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetPollResultsRequest) TypeID() uint32 {
	return MessagesGetPollResultsRequestTypeID
}

// Encode implements bin.Encoder.
func (g *MessagesGetPollResultsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPollResults#73bb643b as nil")
	}
	b.PutID(MessagesGetPollResultsRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getPollResults#73bb643b: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getPollResults#73bb643b: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetPollResultsRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetMsgID returns value of MsgID field.
func (g *MessagesGetPollResultsRequest) GetMsgID() (value int) {
	return g.MsgID
}

// Decode implements bin.Decoder.
func (g *MessagesGetPollResultsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPollResults#73bb643b to nil")
	}
	if err := b.ConsumeID(MessagesGetPollResultsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getPollResults#73bb643b: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPollResults#73bb643b: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPollResults#73bb643b: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetPollResultsRequest.
var (
	_ bin.Encoder = &MessagesGetPollResultsRequest{}
	_ bin.Decoder = &MessagesGetPollResultsRequest{}
)

// MessagesGetPollResults invokes method messages.getPollResults#73bb643b returning error if any.
// Get poll results
//
// Possible errors:
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//
// See https://core.telegram.org/method/messages.getPollResults for reference.
func (c *Client) MessagesGetPollResults(ctx context.Context, request *MessagesGetPollResultsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
