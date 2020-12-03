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

// MessagesSetTypingRequest represents TL type `messages.setTyping#58943ee2`.
//
// See https://core.telegram.org/method/messages.setTyping for reference.
type MessagesSetTypingRequest struct {
	// Flags field of MessagesSetTypingRequest.
	Flags bin.Fields
	// Peer field of MessagesSetTypingRequest.
	Peer InputPeerClass
	// TopMsgID field of MessagesSetTypingRequest.
	//
	// Use SetTopMsgID and GetTopMsgID helpers.
	TopMsgID int
	// Action field of MessagesSetTypingRequest.
	Action SendMessageActionClass
}

// MessagesSetTypingRequestTypeID is TL type id of MessagesSetTypingRequest.
const MessagesSetTypingRequestTypeID = 0x58943ee2

// Encode implements bin.Encoder.
func (s *MessagesSetTypingRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setTyping#58943ee2 as nil")
	}
	b.PutID(MessagesSetTypingRequestTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.TopMsgID)
	}
	if s.Action == nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field action is nil")
	}
	if err := s.Action.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field action: %w", err)
	}
	return nil
}

// SetTopMsgID sets value of TopMsgID conditional field.
func (s *MessagesSetTypingRequest) SetTopMsgID(value int) {
	s.Flags.Set(0)
	s.TopMsgID = value
}

// GetTopMsgID returns value of TopMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSetTypingRequest) GetTopMsgID() (value int, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.TopMsgID, true
}

// Decode implements bin.Decoder.
func (s *MessagesSetTypingRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setTyping#58943ee2 to nil")
	}
	if err := b.ConsumeID(MessagesSetTypingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setTyping#58943ee2: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setTyping#58943ee2: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setTyping#58943ee2: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setTyping#58943ee2: field top_msg_id: %w", err)
		}
		s.TopMsgID = value
	}
	{
		value, err := DecodeSendMessageAction(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setTyping#58943ee2: field action: %w", err)
		}
		s.Action = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetTypingRequest.
var (
	_ bin.Encoder = &MessagesSetTypingRequest{}
	_ bin.Decoder = &MessagesSetTypingRequest{}
)

// MessagesSetTyping invokes method messages.setTyping#58943ee2 returning error if any.
//
// See https://core.telegram.org/method/messages.setTyping for reference.
func (c *Client) MessagesSetTyping(ctx context.Context, request *MessagesSetTypingRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
