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

// PaymentCharge represents TL type `paymentCharge#ea02c27e`.
//
// See https://core.telegram.org/constructor/paymentCharge for reference.
type PaymentCharge struct {
	// ID field of PaymentCharge.
	ID string
	// ProviderChargeID field of PaymentCharge.
	ProviderChargeID string
}

// PaymentChargeTypeID is TL type id of PaymentCharge.
const PaymentChargeTypeID = 0xea02c27e

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
