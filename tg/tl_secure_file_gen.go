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

// SecureFileEmpty represents TL type `secureFileEmpty#64199744`.
//
// See https://core.telegram.org/constructor/secureFileEmpty for reference.
type SecureFileEmpty struct {
}

// SecureFileEmptyTypeID is TL type id of SecureFileEmpty.
const SecureFileEmptyTypeID = 0x64199744

// Encode implements bin.Encoder.
func (s *SecureFileEmpty) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureFileEmpty#64199744 as nil")
	}
	b.PutID(SecureFileEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureFileEmpty) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureFileEmpty#64199744 to nil")
	}
	if err := b.ConsumeID(SecureFileEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode secureFileEmpty#64199744: %w", err)
	}
	return nil
}

// construct implements constructor of SecureFileClass.
func (s SecureFileEmpty) construct() SecureFileClass { return &s }

// Ensuring interfaces in compile-time for SecureFileEmpty.
var (
	_ bin.Encoder = &SecureFileEmpty{}
	_ bin.Decoder = &SecureFileEmpty{}

	_ SecureFileClass = &SecureFileEmpty{}
)

// SecureFile represents TL type `secureFile#e0277a62`.
//
// See https://core.telegram.org/constructor/secureFile for reference.
type SecureFile struct {
	// ID field of SecureFile.
	ID int64
	// AccessHash field of SecureFile.
	AccessHash int64
	// Size field of SecureFile.
	Size int
	// DCID field of SecureFile.
	DCID int
	// Date field of SecureFile.
	Date int
	// FileHash field of SecureFile.
	FileHash []byte
	// Secret field of SecureFile.
	Secret []byte
}

// SecureFileTypeID is TL type id of SecureFile.
const SecureFileTypeID = 0xe0277a62

// Encode implements bin.Encoder.
func (s *SecureFile) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureFile#e0277a62 as nil")
	}
	b.PutID(SecureFileTypeID)
	b.PutLong(s.ID)
	b.PutLong(s.AccessHash)
	b.PutInt(s.Size)
	b.PutInt(s.DCID)
	b.PutInt(s.Date)
	b.PutBytes(s.FileHash)
	b.PutBytes(s.Secret)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureFile) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureFile#e0277a62 to nil")
	}
	if err := b.ConsumeID(SecureFileTypeID); err != nil {
		return fmt.Errorf("unable to decode secureFile#e0277a62: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field access_hash: %w", err)
		}
		s.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field size: %w", err)
		}
		s.Size = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field dc_id: %w", err)
		}
		s.DCID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field file_hash: %w", err)
		}
		s.FileHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#e0277a62: field secret: %w", err)
		}
		s.Secret = value
	}
	return nil
}

// construct implements constructor of SecureFileClass.
func (s SecureFile) construct() SecureFileClass { return &s }

// Ensuring interfaces in compile-time for SecureFile.
var (
	_ bin.Encoder = &SecureFile{}
	_ bin.Decoder = &SecureFile{}

	_ SecureFileClass = &SecureFile{}
)

// SecureFileClass represents SecureFile generic type.
//
// See https://core.telegram.org/type/SecureFile for reference.
//
// Example:
//  g, err := DecodeSecureFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *SecureFileEmpty: // secureFileEmpty#64199744
//  case *SecureFile: // secureFile#e0277a62
//  default: panic(v)
//  }
type SecureFileClass interface {
	bin.Encoder
	bin.Decoder
	construct() SecureFileClass
}

// DecodeSecureFile implements binary de-serialization for SecureFileClass.
func DecodeSecureFile(buf *bin.Buffer) (SecureFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SecureFileEmptyTypeID:
		// Decoding secureFileEmpty#64199744.
		v := SecureFileEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureFileClass: %w", err)
		}
		return &v, nil
	case SecureFileTypeID:
		// Decoding secureFile#e0277a62.
		v := SecureFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SecureFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// SecureFile boxes the SecureFileClass providing a helper.
type SecureFileBox struct {
	SecureFile SecureFileClass
}

// Decode implements bin.Decoder for SecureFileBox.
func (b *SecureFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SecureFileBox to nil")
	}
	v, err := DecodeSecureFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SecureFile = v
	return nil
}

// Encode implements bin.Encode for SecureFileBox.
func (b *SecureFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SecureFile == nil {
		return fmt.Errorf("unable to encode SecureFileClass as nil")
	}
	return b.SecureFile.Encode(buf)
}
