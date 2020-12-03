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

// UpdatesState represents TL type `updates.state#a56c2a3e`.
//
// See https://core.telegram.org/constructor/updates.state for reference.
type UpdatesState struct {
	// Pts field of UpdatesState.
	Pts int
	// Qts field of UpdatesState.
	Qts int
	// Date field of UpdatesState.
	Date int
	// Seq field of UpdatesState.
	Seq int
	// UnreadCount field of UpdatesState.
	UnreadCount int
}

// UpdatesStateTypeID is TL type id of UpdatesState.
const UpdatesStateTypeID = 0xa56c2a3e

// Encode implements bin.Encoder.
func (s *UpdatesState) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode updates.state#a56c2a3e as nil")
	}
	b.PutID(UpdatesStateTypeID)
	b.PutInt(s.Pts)
	b.PutInt(s.Qts)
	b.PutInt(s.Date)
	b.PutInt(s.Seq)
	b.PutInt(s.UnreadCount)
	return nil
}

// Decode implements bin.Decoder.
func (s *UpdatesState) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode updates.state#a56c2a3e to nil")
	}
	if err := b.ConsumeID(UpdatesStateTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.state#a56c2a3e: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.state#a56c2a3e: field pts: %w", err)
		}
		s.Pts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.state#a56c2a3e: field qts: %w", err)
		}
		s.Qts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.state#a56c2a3e: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.state#a56c2a3e: field seq: %w", err)
		}
		s.Seq = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.state#a56c2a3e: field unread_count: %w", err)
		}
		s.UnreadCount = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UpdatesState.
var (
	_ bin.Encoder = &UpdatesState{}
	_ bin.Decoder = &UpdatesState{}
)
