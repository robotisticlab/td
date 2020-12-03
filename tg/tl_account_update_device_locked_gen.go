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

// AccountUpdateDeviceLockedRequest represents TL type `account.updateDeviceLocked#38df3532`.
//
// See https://core.telegram.org/method/account.updateDeviceLocked for reference.
type AccountUpdateDeviceLockedRequest struct {
	// Period field of AccountUpdateDeviceLockedRequest.
	Period int
}

// AccountUpdateDeviceLockedRequestTypeID is TL type id of AccountUpdateDeviceLockedRequest.
const AccountUpdateDeviceLockedRequestTypeID = 0x38df3532

// Encode implements bin.Encoder.
func (u *AccountUpdateDeviceLockedRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateDeviceLocked#38df3532 as nil")
	}
	b.PutID(AccountUpdateDeviceLockedRequestTypeID)
	b.PutInt(u.Period)
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateDeviceLockedRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateDeviceLocked#38df3532 to nil")
	}
	if err := b.ConsumeID(AccountUpdateDeviceLockedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateDeviceLocked#38df3532: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateDeviceLocked#38df3532: field period: %w", err)
		}
		u.Period = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUpdateDeviceLockedRequest.
var (
	_ bin.Encoder = &AccountUpdateDeviceLockedRequest{}
	_ bin.Decoder = &AccountUpdateDeviceLockedRequest{}
)

// AccountUpdateDeviceLocked invokes method account.updateDeviceLocked#38df3532 returning error if any.
//
// See https://core.telegram.org/method/account.updateDeviceLocked for reference.
func (c *Client) AccountUpdateDeviceLocked(ctx context.Context, request *AccountUpdateDeviceLockedRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
