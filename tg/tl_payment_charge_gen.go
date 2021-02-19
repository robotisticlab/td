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

// PaymentCharge represents TL type `paymentCharge#ea02c27e`.
// Payment identifier
//
// See https://core.telegram.org/constructor/paymentCharge for reference.
type PaymentCharge struct {
	// Telegram payment identifier
	ID string
	// Provider payment identifier
	ProviderChargeID string
}

// PaymentChargeTypeID is TL type id of PaymentCharge.
const PaymentChargeTypeID = 0xea02c27e

func (p *PaymentCharge) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ID == "") {
		return false
	}
	if !(p.ProviderChargeID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentCharge) String() string {
	if p == nil {
		return "PaymentCharge(nil)"
	}
	type Alias PaymentCharge
	return fmt.Sprintf("PaymentCharge%+v", Alias(*p))
}

// FillFrom fills PaymentCharge from given interface.
func (p *PaymentCharge) FillFrom(from interface {
	GetID() (value string)
	GetProviderChargeID() (value string)
}) {
	p.ID = from.GetID()
	p.ProviderChargeID = from.GetProviderChargeID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PaymentCharge) TypeID() uint32 {
	return PaymentChargeTypeID
}

// Encode implements bin.Encoder.
func (p *PaymentCharge) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentCharge#ea02c27e as nil")
	}
	b.PutID(PaymentChargeTypeID)
	b.PutString(p.ID)
	b.PutString(p.ProviderChargeID)
	return nil
}

// GetID returns value of ID field.
func (p *PaymentCharge) GetID() (value string) {
	return p.ID
}

// GetProviderChargeID returns value of ProviderChargeID field.
func (p *PaymentCharge) GetProviderChargeID() (value string) {
	return p.ProviderChargeID
}

// Decode implements bin.Decoder.
func (p *PaymentCharge) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentCharge#ea02c27e to nil")
	}
	if err := b.ConsumeID(PaymentChargeTypeID); err != nil {
		return fmt.Errorf("unable to decode paymentCharge#ea02c27e: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentCharge#ea02c27e: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentCharge#ea02c27e: field provider_charge_id: %w", err)
		}
		p.ProviderChargeID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentCharge.
var (
	_ bin.Encoder = &PaymentCharge{}
	_ bin.Decoder = &PaymentCharge{}
)
