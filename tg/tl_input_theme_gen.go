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

// InputTheme represents TL type `inputTheme#3c5693e9`.
//
// See https://core.telegram.org/constructor/inputTheme for reference.
type InputTheme struct {
	// ID field of InputTheme.
	ID int64
	// AccessHash field of InputTheme.
	AccessHash int64
}

// InputThemeTypeID is TL type id of InputTheme.
const InputThemeTypeID = 0x3c5693e9

// Encode implements bin.Encoder.
func (i *InputTheme) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputTheme#3c5693e9 as nil")
	}
	b.PutID(InputThemeTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputTheme) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputTheme#3c5693e9 to nil")
	}
	if err := b.ConsumeID(InputThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode inputTheme#3c5693e9: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputTheme#3c5693e9: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputTheme#3c5693e9: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputThemeClass.
func (i InputTheme) construct() InputThemeClass { return &i }

// Ensuring interfaces in compile-time for InputTheme.
var (
	_ bin.Encoder = &InputTheme{}
	_ bin.Decoder = &InputTheme{}

	_ InputThemeClass = &InputTheme{}
)

// InputThemeSlug represents TL type `inputThemeSlug#f5890df1`.
//
// See https://core.telegram.org/constructor/inputThemeSlug for reference.
type InputThemeSlug struct {
	// Slug field of InputThemeSlug.
	Slug string
}

// InputThemeSlugTypeID is TL type id of InputThemeSlug.
const InputThemeSlugTypeID = 0xf5890df1

// Encode implements bin.Encoder.
func (i *InputThemeSlug) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputThemeSlug#f5890df1 as nil")
	}
	b.PutID(InputThemeSlugTypeID)
	b.PutString(i.Slug)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputThemeSlug) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputThemeSlug#f5890df1 to nil")
	}
	if err := b.ConsumeID(InputThemeSlugTypeID); err != nil {
		return fmt.Errorf("unable to decode inputThemeSlug#f5890df1: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputThemeSlug#f5890df1: field slug: %w", err)
		}
		i.Slug = value
	}
	return nil
}

// construct implements constructor of InputThemeClass.
func (i InputThemeSlug) construct() InputThemeClass { return &i }

// Ensuring interfaces in compile-time for InputThemeSlug.
var (
	_ bin.Encoder = &InputThemeSlug{}
	_ bin.Decoder = &InputThemeSlug{}

	_ InputThemeClass = &InputThemeSlug{}
)

// InputThemeClass represents InputTheme generic type.
//
// See https://core.telegram.org/type/InputTheme for reference.
//
// Example:
//  g, err := DecodeInputTheme(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputTheme: // inputTheme#3c5693e9
//  case *InputThemeSlug: // inputThemeSlug#f5890df1
//  default: panic(v)
//  }
type InputThemeClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputThemeClass
}

// DecodeInputTheme implements binary de-serialization for InputThemeClass.
func DecodeInputTheme(buf *bin.Buffer) (InputThemeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputThemeTypeID:
		// Decoding inputTheme#3c5693e9.
		v := InputTheme{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputThemeClass: %w", err)
		}
		return &v, nil
	case InputThemeSlugTypeID:
		// Decoding inputThemeSlug#f5890df1.
		v := InputThemeSlug{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputThemeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputThemeClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputTheme boxes the InputThemeClass providing a helper.
type InputThemeBox struct {
	InputTheme InputThemeClass
}

// Decode implements bin.Decoder for InputThemeBox.
func (b *InputThemeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputThemeBox to nil")
	}
	v, err := DecodeInputTheme(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputTheme = v
	return nil
}

// Encode implements bin.Encode for InputThemeBox.
func (b *InputThemeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputTheme == nil {
		return fmt.Errorf("unable to encode InputThemeClass as nil")
	}
	return b.InputTheme.Encode(buf)
}
