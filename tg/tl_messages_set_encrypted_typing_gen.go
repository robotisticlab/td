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

// MessagesSetEncryptedTypingRequest represents TL type `messages.setEncryptedTyping#791451ed`.
//
// See https://core.telegram.org/method/messages.setEncryptedTyping for reference.
type MessagesSetEncryptedTypingRequest struct {
	// Peer field of MessagesSetEncryptedTypingRequest.
	Peer InputEncryptedChat
	// Typing field of MessagesSetEncryptedTypingRequest.
	Typing bool
}

// MessagesSetEncryptedTypingRequestTypeID is TL type id of MessagesSetEncryptedTypingRequest.
const MessagesSetEncryptedTypingRequestTypeID = 0x791451ed

// Encode implements bin.Encoder.
func (s *MessagesSetEncryptedTypingRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setEncryptedTyping#791451ed as nil")
	}
	b.PutID(MessagesSetEncryptedTypingRequestTypeID)
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setEncryptedTyping#791451ed: field peer: %w", err)
	}
	b.PutBool(s.Typing)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSetEncryptedTypingRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setEncryptedTyping#791451ed to nil")
	}
	if err := b.ConsumeID(MessagesSetEncryptedTypingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setEncryptedTyping#791451ed: %w", err)
	}
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setEncryptedTyping#791451ed: field peer: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setEncryptedTyping#791451ed: field typing: %w", err)
		}
		s.Typing = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetEncryptedTypingRequest.
var (
	_ bin.Encoder = &MessagesSetEncryptedTypingRequest{}
	_ bin.Decoder = &MessagesSetEncryptedTypingRequest{}
)

// MessagesSetEncryptedTyping invokes method messages.setEncryptedTyping#791451ed returning error if any.
//
// See https://core.telegram.org/method/messages.setEncryptedTyping for reference.
func (c *Client) MessagesSetEncryptedTyping(ctx context.Context, request *MessagesSetEncryptedTypingRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
