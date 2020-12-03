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

// MessagesSetBotCallbackAnswerRequest represents TL type `messages.setBotCallbackAnswer#d58f130a`.
//
// See https://core.telegram.org/method/messages.setBotCallbackAnswer for reference.
type MessagesSetBotCallbackAnswerRequest struct {
	// Flags field of MessagesSetBotCallbackAnswerRequest.
	Flags bin.Fields
	// Alert field of MessagesSetBotCallbackAnswerRequest.
	Alert bool
	// QueryID field of MessagesSetBotCallbackAnswerRequest.
	QueryID int64
	// Message field of MessagesSetBotCallbackAnswerRequest.
	//
	// Use SetMessage and GetMessage helpers.
	Message string
	// URL field of MessagesSetBotCallbackAnswerRequest.
	//
	// Use SetURL and GetURL helpers.
	URL string
	// CacheTime field of MessagesSetBotCallbackAnswerRequest.
	CacheTime int
}

// MessagesSetBotCallbackAnswerRequestTypeID is TL type id of MessagesSetBotCallbackAnswerRequest.
const MessagesSetBotCallbackAnswerRequestTypeID = 0xd58f130a

// Encode implements bin.Encoder.
func (s *MessagesSetBotCallbackAnswerRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setBotCallbackAnswer#d58f130a as nil")
	}
	b.PutID(MessagesSetBotCallbackAnswerRequestTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setBotCallbackAnswer#d58f130a: field flags: %w", err)
	}
	b.PutLong(s.QueryID)
	if s.Flags.Has(0) {
		b.PutString(s.Message)
	}
	if s.Flags.Has(2) {
		b.PutString(s.URL)
	}
	b.PutInt(s.CacheTime)
	return nil
}

// SetAlert sets value of Alert conditional field.
func (s *MessagesSetBotCallbackAnswerRequest) SetAlert(value bool) {
	if value {
		s.Flags.Set(1)
	} else {
		s.Flags.Unset(1)
	}
}

// SetMessage sets value of Message conditional field.
func (s *MessagesSetBotCallbackAnswerRequest) SetMessage(value string) {
	s.Flags.Set(0)
	s.Message = value
}

// GetMessage returns value of Message conditional field and
// boolean which is true if field was set.
func (s *MessagesSetBotCallbackAnswerRequest) GetMessage() (value string, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Message, true
}

// SetURL sets value of URL conditional field.
func (s *MessagesSetBotCallbackAnswerRequest) SetURL(value string) {
	s.Flags.Set(2)
	s.URL = value
}

// GetURL returns value of URL conditional field and
// boolean which is true if field was set.
func (s *MessagesSetBotCallbackAnswerRequest) GetURL() (value string, ok bool) {
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.URL, true
}

// Decode implements bin.Decoder.
func (s *MessagesSetBotCallbackAnswerRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setBotCallbackAnswer#d58f130a to nil")
	}
	if err := b.ConsumeID(MessagesSetBotCallbackAnswerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field flags: %w", err)
		}
	}
	s.Alert = s.Flags.Has(1)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field query_id: %w", err)
		}
		s.QueryID = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field message: %w", err)
		}
		s.Message = value
	}
	if s.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field url: %w", err)
		}
		s.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field cache_time: %w", err)
		}
		s.CacheTime = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetBotCallbackAnswerRequest.
var (
	_ bin.Encoder = &MessagesSetBotCallbackAnswerRequest{}
	_ bin.Decoder = &MessagesSetBotCallbackAnswerRequest{}
)

// MessagesSetBotCallbackAnswer invokes method messages.setBotCallbackAnswer#d58f130a returning error if any.
//
// See https://core.telegram.org/method/messages.setBotCallbackAnswer for reference.
func (c *Client) MessagesSetBotCallbackAnswer(ctx context.Context, request *MessagesSetBotCallbackAnswerRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
