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

// EncryptedFileEmpty represents TL type `encryptedFileEmpty#c21f497e`.
//
// See https://core.telegram.org/constructor/encryptedFileEmpty for reference.
type EncryptedFileEmpty struct {
}

// EncryptedFileEmptyTypeID is TL type id of EncryptedFileEmpty.
const EncryptedFileEmptyTypeID = 0xc21f497e

// Encode implements bin.Encoder.
func (e *EncryptedFileEmpty) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedFileEmpty#c21f497e as nil")
	}
	b.PutID(EncryptedFileEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedFileEmpty) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedFileEmpty#c21f497e to nil")
	}
	if err := b.ConsumeID(EncryptedFileEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedFileEmpty#c21f497e: %w", err)
	}
	return nil
}

// construct implements constructor of EncryptedFileClass.
func (e EncryptedFileEmpty) construct() EncryptedFileClass { return &e }

// Ensuring interfaces in compile-time for EncryptedFileEmpty.
var (
	_ bin.Encoder = &EncryptedFileEmpty{}
	_ bin.Decoder = &EncryptedFileEmpty{}

	_ EncryptedFileClass = &EncryptedFileEmpty{}
)

// EncryptedFile represents TL type `encryptedFile#4a70994c`.
//
// See https://core.telegram.org/constructor/encryptedFile for reference.
type EncryptedFile struct {
	// ID field of EncryptedFile.
	ID int64
	// AccessHash field of EncryptedFile.
	AccessHash int64
	// Size field of EncryptedFile.
	Size int
	// DCID field of EncryptedFile.
	DCID int
	// KeyFingerprint field of EncryptedFile.
	KeyFingerprint int
}

// EncryptedFileTypeID is TL type id of EncryptedFile.
const EncryptedFileTypeID = 0x4a70994c

// Encode implements bin.Encoder.
func (e *EncryptedFile) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedFile#4a70994c as nil")
	}
	b.PutID(EncryptedFileTypeID)
	b.PutLong(e.ID)
	b.PutLong(e.AccessHash)
	b.PutInt(e.Size)
	b.PutInt(e.DCID)
	b.PutInt(e.KeyFingerprint)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedFile) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedFile#4a70994c to nil")
	}
	if err := b.ConsumeID(EncryptedFileTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedFile#4a70994c: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#4a70994c: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#4a70994c: field access_hash: %w", err)
		}
		e.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#4a70994c: field size: %w", err)
		}
		e.Size = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#4a70994c: field dc_id: %w", err)
		}
		e.DCID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#4a70994c: field key_fingerprint: %w", err)
		}
		e.KeyFingerprint = value
	}
	return nil
}

// construct implements constructor of EncryptedFileClass.
func (e EncryptedFile) construct() EncryptedFileClass { return &e }

// Ensuring interfaces in compile-time for EncryptedFile.
var (
	_ bin.Encoder = &EncryptedFile{}
	_ bin.Decoder = &EncryptedFile{}

	_ EncryptedFileClass = &EncryptedFile{}
)

// EncryptedFileClass represents EncryptedFile generic type.
//
// See https://core.telegram.org/type/EncryptedFile for reference.
//
// Example:
//  g, err := DecodeEncryptedFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *EncryptedFileEmpty: // encryptedFileEmpty#c21f497e
//  case *EncryptedFile: // encryptedFile#4a70994c
//  default: panic(v)
//  }
type EncryptedFileClass interface {
	bin.Encoder
	bin.Decoder
	construct() EncryptedFileClass
}

// DecodeEncryptedFile implements binary de-serialization for EncryptedFileClass.
func DecodeEncryptedFile(buf *bin.Buffer) (EncryptedFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case EncryptedFileEmptyTypeID:
		// Decoding encryptedFileEmpty#c21f497e.
		v := EncryptedFileEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedFileClass: %w", err)
		}
		return &v, nil
	case EncryptedFileTypeID:
		// Decoding encryptedFile#4a70994c.
		v := EncryptedFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EncryptedFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// EncryptedFile boxes the EncryptedFileClass providing a helper.
type EncryptedFileBox struct {
	EncryptedFile EncryptedFileClass
}

// Decode implements bin.Decoder for EncryptedFileBox.
func (b *EncryptedFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode EncryptedFileBox to nil")
	}
	v, err := DecodeEncryptedFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EncryptedFile = v
	return nil
}

// Encode implements bin.Encode for EncryptedFileBox.
func (b *EncryptedFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EncryptedFile == nil {
		return fmt.Errorf("unable to encode EncryptedFileClass as nil")
	}
	return b.EncryptedFile.Encode(buf)
}
