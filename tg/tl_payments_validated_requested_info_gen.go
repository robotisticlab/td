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

// PaymentsValidatedRequestedInfo represents TL type `payments.validatedRequestedInfo#d1451883`.
//
// See https://core.telegram.org/constructor/payments.validatedRequestedInfo for reference.
type PaymentsValidatedRequestedInfo struct {
	// Flags field of PaymentsValidatedRequestedInfo.
	Flags bin.Fields
	// ID field of PaymentsValidatedRequestedInfo.
	//
	// Use SetID and GetID helpers.
	ID string
	// ShippingOptions field of PaymentsValidatedRequestedInfo.
	//
	// Use SetShippingOptions and GetShippingOptions helpers.
	ShippingOptions []ShippingOption
}

// PaymentsValidatedRequestedInfoTypeID is TL type id of PaymentsValidatedRequestedInfo.
const PaymentsValidatedRequestedInfoTypeID = 0xd1451883

func (v *PaymentsValidatedRequestedInfo) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Flags.Zero()) {
		return false
	}
	if !(v.ID == "") {
		return false
	}
	if !(v.ShippingOptions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *PaymentsValidatedRequestedInfo) String() string {
	if v == nil {
		return "PaymentsValidatedRequestedInfo(nil)"
	}
	type Alias PaymentsValidatedRequestedInfo
	return fmt.Sprintf("PaymentsValidatedRequestedInfo%+v", Alias(*v))
}

// FillFrom fills PaymentsValidatedRequestedInfo from given interface.
func (v *PaymentsValidatedRequestedInfo) FillFrom(from interface {
	GetID() (value string, ok bool)
	GetShippingOptions() (value []ShippingOption, ok bool)
}) {
	if val, ok := from.GetID(); ok {
		v.ID = val
	}
	if val, ok := from.GetShippingOptions(); ok {
		v.ShippingOptions = val
	}
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (v *PaymentsValidatedRequestedInfo) TypeID() uint32 {
	return PaymentsValidatedRequestedInfoTypeID
}

// Encode implements bin.Encoder.
func (v *PaymentsValidatedRequestedInfo) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode payments.validatedRequestedInfo#d1451883 as nil")
	}
	b.PutID(PaymentsValidatedRequestedInfoTypeID)
	if !(v.ID == "") {
		v.Flags.Set(0)
	}
	if !(v.ShippingOptions == nil) {
		v.Flags.Set(1)
	}
	if err := v.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.validatedRequestedInfo#d1451883: field flags: %w", err)
	}
	if v.Flags.Has(0) {
		b.PutString(v.ID)
	}
	if v.Flags.Has(1) {
		b.PutVectorHeader(len(v.ShippingOptions))
		for idx, v := range v.ShippingOptions {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode payments.validatedRequestedInfo#d1451883: field shipping_options element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// SetID sets value of ID conditional field.
func (v *PaymentsValidatedRequestedInfo) SetID(value string) {
	v.Flags.Set(0)
	v.ID = value
}

// GetID returns value of ID conditional field and
// boolean which is true if field was set.
func (v *PaymentsValidatedRequestedInfo) GetID() (value string, ok bool) {
	if !v.Flags.Has(0) {
		return value, false
	}
	return v.ID, true
}

// SetShippingOptions sets value of ShippingOptions conditional field.
func (v *PaymentsValidatedRequestedInfo) SetShippingOptions(value []ShippingOption) {
	v.Flags.Set(1)
	v.ShippingOptions = value
}

// GetShippingOptions returns value of ShippingOptions conditional field and
// boolean which is true if field was set.
func (v *PaymentsValidatedRequestedInfo) GetShippingOptions() (value []ShippingOption, ok bool) {
	if !v.Flags.Has(1) {
		return value, false
	}
	return v.ShippingOptions, true
}

// Decode implements bin.Decoder.
func (v *PaymentsValidatedRequestedInfo) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode payments.validatedRequestedInfo#d1451883 to nil")
	}
	if err := b.ConsumeID(PaymentsValidatedRequestedInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.validatedRequestedInfo#d1451883: %w", err)
	}
	{
		if err := v.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.validatedRequestedInfo#d1451883: field flags: %w", err)
		}
	}
	if v.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.validatedRequestedInfo#d1451883: field id: %w", err)
		}
		v.ID = value
	}
	if v.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode payments.validatedRequestedInfo#d1451883: field shipping_options: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ShippingOption
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode payments.validatedRequestedInfo#d1451883: field shipping_options: %w", err)
			}
			v.ShippingOptions = append(v.ShippingOptions, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsValidatedRequestedInfo.
var (
	_ bin.Encoder = &PaymentsValidatedRequestedInfo{}
	_ bin.Decoder = &PaymentsValidatedRequestedInfo{}
)
