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

// AccountSendChangePhoneCodeRequest represents TL type `account.sendChangePhoneCode#82574ae5`.
//
// See https://core.telegram.org/method/account.sendChangePhoneCode for reference.
type AccountSendChangePhoneCodeRequest struct {
	// PhoneNumber field of AccountSendChangePhoneCodeRequest.
	PhoneNumber string
	// Settings field of AccountSendChangePhoneCodeRequest.
	Settings CodeSettings
}

// AccountSendChangePhoneCodeRequestTypeID is TL type id of AccountSendChangePhoneCodeRequest.
const AccountSendChangePhoneCodeRequestTypeID = 0x82574ae5

// Encode implements bin.Encoder.
func (s *AccountSendChangePhoneCodeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sendChangePhoneCode#82574ae5 as nil")
	}
	b.PutID(AccountSendChangePhoneCodeRequestTypeID)
	b.PutString(s.PhoneNumber)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.sendChangePhoneCode#82574ae5: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSendChangePhoneCodeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sendChangePhoneCode#82574ae5 to nil")
	}
	if err := b.ConsumeID(AccountSendChangePhoneCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.sendChangePhoneCode#82574ae5: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.sendChangePhoneCode#82574ae5: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.sendChangePhoneCode#82574ae5: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSendChangePhoneCodeRequest.
var (
	_ bin.Encoder = &AccountSendChangePhoneCodeRequest{}
	_ bin.Decoder = &AccountSendChangePhoneCodeRequest{}
)

// AccountSendChangePhoneCode invokes method account.sendChangePhoneCode#82574ae5 returning error if any.
//
// See https://core.telegram.org/method/account.sendChangePhoneCode for reference.
func (c *Client) AccountSendChangePhoneCode(ctx context.Context, request *AccountSendChangePhoneCodeRequest) (*AuthSentCode, error) {
	var result AuthSentCode
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
