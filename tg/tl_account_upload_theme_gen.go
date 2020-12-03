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

// AccountUploadThemeRequest represents TL type `account.uploadTheme#1c3db333`.
//
// See https://core.telegram.org/method/account.uploadTheme for reference.
type AccountUploadThemeRequest struct {
	// Flags field of AccountUploadThemeRequest.
	Flags bin.Fields
	// File field of AccountUploadThemeRequest.
	File InputFileClass
	// Thumb field of AccountUploadThemeRequest.
	//
	// Use SetThumb and GetThumb helpers.
	Thumb InputFileClass
	// FileName field of AccountUploadThemeRequest.
	FileName string
	// MimeType field of AccountUploadThemeRequest.
	MimeType string
}

// AccountUploadThemeRequestTypeID is TL type id of AccountUploadThemeRequest.
const AccountUploadThemeRequestTypeID = 0x1c3db333

// Encode implements bin.Encoder.
func (u *AccountUploadThemeRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.uploadTheme#1c3db333 as nil")
	}
	b.PutID(AccountUploadThemeRequestTypeID)
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.uploadTheme#1c3db333: field flags: %w", err)
	}
	if u.File == nil {
		return fmt.Errorf("unable to encode account.uploadTheme#1c3db333: field file is nil")
	}
	if err := u.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.uploadTheme#1c3db333: field file: %w", err)
	}
	if u.Flags.Has(0) {
		if u.Thumb == nil {
			return fmt.Errorf("unable to encode account.uploadTheme#1c3db333: field thumb is nil")
		}
		if err := u.Thumb.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.uploadTheme#1c3db333: field thumb: %w", err)
		}
	}
	b.PutString(u.FileName)
	b.PutString(u.MimeType)
	return nil
}

// SetThumb sets value of Thumb conditional field.
func (u *AccountUploadThemeRequest) SetThumb(value InputFileClass) {
	u.Flags.Set(0)
	u.Thumb = value
}

// GetThumb returns value of Thumb conditional field and
// boolean which is true if field was set.
func (u *AccountUploadThemeRequest) GetThumb() (value InputFileClass, ok bool) {
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.Thumb, true
}

// Decode implements bin.Decoder.
func (u *AccountUploadThemeRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.uploadTheme#1c3db333 to nil")
	}
	if err := b.ConsumeID(AccountUploadThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.uploadTheme#1c3db333: %w", err)
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.uploadTheme#1c3db333: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadTheme#1c3db333: field file: %w", err)
		}
		u.File = value
	}
	if u.Flags.Has(0) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadTheme#1c3db333: field thumb: %w", err)
		}
		u.Thumb = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadTheme#1c3db333: field file_name: %w", err)
		}
		u.FileName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadTheme#1c3db333: field mime_type: %w", err)
		}
		u.MimeType = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUploadThemeRequest.
var (
	_ bin.Encoder = &AccountUploadThemeRequest{}
	_ bin.Decoder = &AccountUploadThemeRequest{}
)

// AccountUploadTheme invokes method account.uploadTheme#1c3db333 returning error if any.
//
// See https://core.telegram.org/method/account.uploadTheme for reference.
func (c *Client) AccountUploadTheme(ctx context.Context, request *AccountUploadThemeRequest) (DocumentClass, error) {
	var result DocumentBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Document, nil
}
