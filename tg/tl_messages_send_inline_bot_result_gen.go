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

// MessagesSendInlineBotResultRequest represents TL type `messages.sendInlineBotResult#220815b0`.
// Send a result obtained using messages.getInlineBotResults¹.
//
// Links:
//  1) https://core.telegram.org/method/messages.getInlineBotResults
//
// See https://core.telegram.org/method/messages.sendInlineBotResult for reference.
type MessagesSendInlineBotResultRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to send the message silently (no notification will be triggered on the other client)
	Silent bool
	// Whether to send the message in background
	Background bool
	// Whether to clear the draft¹
	//
	// Links:
	//  1) https://core.telegram.org/api/drafts
	ClearDraft bool
	// Whether to hide the via @botname in the resulting message (only for bot usernames encountered in the config¹)
	//
	// Links:
	//  1) https://core.telegram.org/constructor/config
	HideVia bool
	// Destination
	Peer InputPeerClass
	// ID of the message this message should reply to
	//
	// Use SetReplyToMsgID and GetReplyToMsgID helpers.
	ReplyToMsgID int
	// Random ID to avoid resending the same query
	RandomID int64
	// Query ID from messages.getInlineBotResults¹
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.getInlineBotResults
	QueryID int64
	// Result ID from messages.getInlineBotResults¹
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.getInlineBotResults
	ID string
	// Scheduled message date for scheduled messages
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
}

// MessagesSendInlineBotResultRequestTypeID is TL type id of MessagesSendInlineBotResultRequest.
const MessagesSendInlineBotResultRequestTypeID = 0x220815b0

func (s *MessagesSendInlineBotResultRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Silent == false) {
		return false
	}
	if !(s.Background == false) {
		return false
	}
	if !(s.ClearDraft == false) {
		return false
	}
	if !(s.HideVia == false) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.ReplyToMsgID == 0) {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}
	if !(s.QueryID == 0) {
		return false
	}
	if !(s.ID == "") {
		return false
	}
	if !(s.ScheduleDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendInlineBotResultRequest) String() string {
	if s == nil {
		return "MessagesSendInlineBotResultRequest(nil)"
	}
	type Alias MessagesSendInlineBotResultRequest
	return fmt.Sprintf("MessagesSendInlineBotResultRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendInlineBotResultRequest from given interface.
func (s *MessagesSendInlineBotResultRequest) FillFrom(from interface {
	GetSilent() (value bool)
	GetBackground() (value bool)
	GetClearDraft() (value bool)
	GetHideVia() (value bool)
	GetPeer() (value InputPeerClass)
	GetReplyToMsgID() (value int, ok bool)
	GetRandomID() (value int64)
	GetQueryID() (value int64)
	GetID() (value string)
	GetScheduleDate() (value int, ok bool)
}) {
	s.Peer = from.GetPeer()
	if val, ok := from.GetReplyToMsgID(); ok {
		s.ReplyToMsgID = val
	}
	s.RandomID = from.GetRandomID()
	s.QueryID = from.GetQueryID()
	s.ID = from.GetID()
	if val, ok := from.GetScheduleDate(); ok {
		s.ScheduleDate = val
	}
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *MessagesSendInlineBotResultRequest) TypeID() uint32 {
	return MessagesSendInlineBotResultRequestTypeID
}

// Encode implements bin.Encoder.
func (s *MessagesSendInlineBotResultRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendInlineBotResult#220815b0 as nil")
	}
	b.PutID(MessagesSendInlineBotResultRequestTypeID)
	if !(s.Silent == false) {
		s.Flags.Set(5)
	}
	if !(s.Background == false) {
		s.Flags.Set(6)
	}
	if !(s.ClearDraft == false) {
		s.Flags.Set(7)
	}
	if !(s.HideVia == false) {
		s.Flags.Set(11)
	}
	if !(s.ReplyToMsgID == 0) {
		s.Flags.Set(0)
	}
	if !(s.ScheduleDate == 0) {
		s.Flags.Set(10)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendInlineBotResult#220815b0: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendInlineBotResult#220815b0: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendInlineBotResult#220815b0: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.ReplyToMsgID)
	}
	b.PutLong(s.RandomID)
	b.PutLong(s.QueryID)
	b.PutString(s.ID)
	if s.Flags.Has(10) {
		b.PutInt(s.ScheduleDate)
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendInlineBotResultRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(5)
		s.Silent = true
	} else {
		s.Flags.Unset(5)
		s.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (s *MessagesSendInlineBotResultRequest) GetSilent() (value bool) {
	return s.Flags.Has(5)
}

// SetBackground sets value of Background conditional field.
func (s *MessagesSendInlineBotResultRequest) SetBackground(value bool) {
	if value {
		s.Flags.Set(6)
		s.Background = true
	} else {
		s.Flags.Unset(6)
		s.Background = false
	}
}

// GetBackground returns value of Background conditional field.
func (s *MessagesSendInlineBotResultRequest) GetBackground() (value bool) {
	return s.Flags.Has(6)
}

// SetClearDraft sets value of ClearDraft conditional field.
func (s *MessagesSendInlineBotResultRequest) SetClearDraft(value bool) {
	if value {
		s.Flags.Set(7)
		s.ClearDraft = true
	} else {
		s.Flags.Unset(7)
		s.ClearDraft = false
	}
}

// GetClearDraft returns value of ClearDraft conditional field.
func (s *MessagesSendInlineBotResultRequest) GetClearDraft() (value bool) {
	return s.Flags.Has(7)
}

// SetHideVia sets value of HideVia conditional field.
func (s *MessagesSendInlineBotResultRequest) SetHideVia(value bool) {
	if value {
		s.Flags.Set(11)
		s.HideVia = true
	} else {
		s.Flags.Unset(11)
		s.HideVia = false
	}
}

// GetHideVia returns value of HideVia conditional field.
func (s *MessagesSendInlineBotResultRequest) GetHideVia() (value bool) {
	return s.Flags.Has(11)
}

// GetPeer returns value of Peer field.
func (s *MessagesSendInlineBotResultRequest) GetPeer() (value InputPeerClass) {
	return s.Peer
}

// SetReplyToMsgID sets value of ReplyToMsgID conditional field.
func (s *MessagesSendInlineBotResultRequest) SetReplyToMsgID(value int) {
	s.Flags.Set(0)
	s.ReplyToMsgID = value
}

// GetReplyToMsgID returns value of ReplyToMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSendInlineBotResultRequest) GetReplyToMsgID() (value int, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ReplyToMsgID, true
}

// GetRandomID returns value of RandomID field.
func (s *MessagesSendInlineBotResultRequest) GetRandomID() (value int64) {
	return s.RandomID
}

// GetQueryID returns value of QueryID field.
func (s *MessagesSendInlineBotResultRequest) GetQueryID() (value int64) {
	return s.QueryID
}

// GetID returns value of ID field.
func (s *MessagesSendInlineBotResultRequest) GetID() (value string) {
	return s.ID
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (s *MessagesSendInlineBotResultRequest) SetScheduleDate(value int) {
	s.Flags.Set(10)
	s.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (s *MessagesSendInlineBotResultRequest) GetScheduleDate() (value int, ok bool) {
	if !s.Flags.Has(10) {
		return value, false
	}
	return s.ScheduleDate, true
}

// Decode implements bin.Decoder.
func (s *MessagesSendInlineBotResultRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendInlineBotResult#220815b0 to nil")
	}
	if err := b.ConsumeID(MessagesSendInlineBotResultRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(5)
	s.Background = s.Flags.Has(6)
	s.ClearDraft = s.Flags.Has(7)
	s.HideVia = s.Flags.Has(11)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field reply_to_msg_id: %w", err)
		}
		s.ReplyToMsgID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field query_id: %w", err)
		}
		s.QueryID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field id: %w", err)
		}
		s.ID = value
	}
	if s.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field schedule_date: %w", err)
		}
		s.ScheduleDate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSendInlineBotResultRequest.
var (
	_ bin.Encoder = &MessagesSendInlineBotResultRequest{}
	_ bin.Decoder = &MessagesSendInlineBotResultRequest{}
)

// MessagesSendInlineBotResult invokes method messages.sendInlineBotResult#220815b0 returning error if any.
// Send a result obtained using messages.getInlineBotResults¹.
//
// Links:
//  1) https://core.telegram.org/method/messages.getInlineBotResults
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_RESTRICTED: You can't send messages in this chat, you were restricted
//  403 CHAT_SEND_GIFS_FORBIDDEN: You can't send gifs in this chat
//  403 CHAT_SEND_INLINE_FORBIDDEN: You can't send inline messages in this group
//  403 CHAT_SEND_MEDIA_FORBIDDEN: You can't send media in this chat
//  403 CHAT_SEND_STICKERS_FORBIDDEN: You can't send stickers in this chat.
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 INLINE_RESULT_EXPIRED: The inline query expired
//  400 INPUT_USER_DEACTIVATED: The specified user was deleted
//  400 MEDIA_EMPTY: The provided media object is invalid
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 QUERY_ID_EMPTY: The query ID is empty
//  400 RESULT_ID_EMPTY: Result ID empty
//  420 SLOWMODE_WAIT_X: Slowmode is enabled in this chat: you must wait for the specified number of seconds before sending another message to the chat.
//  400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels
//  400 WEBPAGE_CURL_FAILED: Failure while fetching the webpage with cURL
//  400 WEBPAGE_MEDIA_EMPTY: Webpage media empty
//  400 YOU_BLOCKED_USER: You blocked this user
//
// See https://core.telegram.org/method/messages.sendInlineBotResult for reference.
func (c *Client) MessagesSendInlineBotResult(ctx context.Context, request *MessagesSendInlineBotResultRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
