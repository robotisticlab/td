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

// AccountUpdateStatusRequest represents TL type `account.updateStatus#6628562c`.
//
// See https://core.telegram.org/method/account.updateStatus for reference.
type AccountUpdateStatusRequest struct {
	// Offline field of AccountUpdateStatusRequest.
	Offline bool
}

// AccountUpdateStatusRequestTypeID is TL type id of AccountUpdateStatusRequest.
const AccountUpdateStatusRequestTypeID = 0x6628562c

// Encode implements bin.Encoder.
func (u *AccountUpdateStatusRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateStatus#6628562c as nil")
	}
	b.PutID(AccountUpdateStatusRequestTypeID)
	b.PutBool(u.Offline)
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateStatusRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateStatus#6628562c to nil")
	}
	if err := b.ConsumeID(AccountUpdateStatusRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateStatus#6628562c: %w", err)
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateStatus#6628562c: field offline: %w", err)
		}
		u.Offline = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUpdateStatusRequest.
var (
	_ bin.Encoder = &AccountUpdateStatusRequest{}
	_ bin.Decoder = &AccountUpdateStatusRequest{}
)

// AccountUpdateStatus invokes method account.updateStatus#6628562c returning error if any.
//
// See https://core.telegram.org/method/account.updateStatus for reference.
func (c *Client) AccountUpdateStatus(ctx context.Context, request *AccountUpdateStatusRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
