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

// ContactsAddContactRequest represents TL type `contacts.addContact#e8f463d0`.
//
// See https://core.telegram.org/method/contacts.addContact for reference.
type ContactsAddContactRequest struct {
	// Flags field of ContactsAddContactRequest.
	Flags bin.Fields
	// AddPhonePrivacyException field of ContactsAddContactRequest.
	AddPhonePrivacyException bool
	// ID field of ContactsAddContactRequest.
	ID InputUserClass
	// FirstName field of ContactsAddContactRequest.
	FirstName string
	// LastName field of ContactsAddContactRequest.
	LastName string
	// Phone field of ContactsAddContactRequest.
	Phone string
}

// ContactsAddContactRequestTypeID is TL type id of ContactsAddContactRequest.
const ContactsAddContactRequestTypeID = 0xe8f463d0

// Encode implements bin.Encoder.
func (a *ContactsAddContactRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode contacts.addContact#e8f463d0 as nil")
	}
	b.PutID(ContactsAddContactRequestTypeID)
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.addContact#e8f463d0: field flags: %w", err)
	}
	if a.ID == nil {
		return fmt.Errorf("unable to encode contacts.addContact#e8f463d0: field id is nil")
	}
	if err := a.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.addContact#e8f463d0: field id: %w", err)
	}
	b.PutString(a.FirstName)
	b.PutString(a.LastName)
	b.PutString(a.Phone)
	return nil
}

// SetAddPhonePrivacyException sets value of AddPhonePrivacyException conditional field.
func (a *ContactsAddContactRequest) SetAddPhonePrivacyException(value bool) {
	if value {
		a.Flags.Set(0)
	} else {
		a.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (a *ContactsAddContactRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode contacts.addContact#e8f463d0 to nil")
	}
	if err := b.ConsumeID(ContactsAddContactRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: %w", err)
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field flags: %w", err)
		}
	}
	a.AddPhonePrivacyException = a.Flags.Has(0)
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field id: %w", err)
		}
		a.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field first_name: %w", err)
		}
		a.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field last_name: %w", err)
		}
		a.LastName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field phone: %w", err)
		}
		a.Phone = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsAddContactRequest.
var (
	_ bin.Encoder = &ContactsAddContactRequest{}
	_ bin.Decoder = &ContactsAddContactRequest{}
)

// ContactsAddContact invokes method contacts.addContact#e8f463d0 returning error if any.
//
// See https://core.telegram.org/method/contacts.addContact for reference.
func (c *Client) ContactsAddContact(ctx context.Context, request *ContactsAddContactRequest) (UpdatesClass, error) {
	var result UpdatesBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
