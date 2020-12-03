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

// MessagesAffectedHistory represents TL type `messages.affectedHistory#b45c69d1`.
//
// See https://core.telegram.org/constructor/messages.affectedHistory for reference.
type MessagesAffectedHistory struct {
	// Pts field of MessagesAffectedHistory.
	Pts int
	// PtsCount field of MessagesAffectedHistory.
	PtsCount int
	// Offset field of MessagesAffectedHistory.
	Offset int
}

// MessagesAffectedHistoryTypeID is TL type id of MessagesAffectedHistory.
const MessagesAffectedHistoryTypeID = 0xb45c69d1

// Encode implements bin.Encoder.
func (a *MessagesAffectedHistory) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.affectedHistory#b45c69d1 as nil")
	}
	b.PutID(MessagesAffectedHistoryTypeID)
	b.PutInt(a.Pts)
	b.PutInt(a.PtsCount)
	b.PutInt(a.Offset)
	return nil
}

// Decode implements bin.Decoder.
func (a *MessagesAffectedHistory) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.affectedHistory#b45c69d1 to nil")
	}
	if err := b.ConsumeID(MessagesAffectedHistoryTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.affectedHistory#b45c69d1: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedHistory#b45c69d1: field pts: %w", err)
		}
		a.Pts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedHistory#b45c69d1: field pts_count: %w", err)
		}
		a.PtsCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedHistory#b45c69d1: field offset: %w", err)
		}
		a.Offset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesAffectedHistory.
var (
	_ bin.Encoder = &MessagesAffectedHistory{}
	_ bin.Decoder = &MessagesAffectedHistory{}
)
