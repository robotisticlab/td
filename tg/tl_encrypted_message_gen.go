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

// EncryptedMessage represents TL type `encryptedMessage#ed18c118`.
//
// See https://core.telegram.org/constructor/encryptedMessage for reference.
type EncryptedMessage struct {
	// RandomID field of EncryptedMessage.
	RandomID int64
	// ChatID field of EncryptedMessage.
	ChatID int
	// Date field of EncryptedMessage.
	Date int
	// Bytes field of EncryptedMessage.
	Bytes []byte
	// File field of EncryptedMessage.
	File EncryptedFileClass
}

// EncryptedMessageTypeID is TL type id of EncryptedMessage.
const EncryptedMessageTypeID = 0xed18c118

// Encode implements bin.Encoder.
func (e *EncryptedMessage) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedMessage#ed18c118 as nil")
	}
	b.PutID(EncryptedMessageTypeID)
	b.PutLong(e.RandomID)
	b.PutInt(e.ChatID)
	b.PutInt(e.Date)
	b.PutBytes(e.Bytes)
	if e.File == nil {
		return fmt.Errorf("unable to encode encryptedMessage#ed18c118: field file is nil")
	}
	if err := e.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode encryptedMessage#ed18c118: field file: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedMessage) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedMessage#ed18c118 to nil")
	}
	if err := b.ConsumeID(EncryptedMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedMessage#ed18c118: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessage#ed18c118: field random_id: %w", err)
		}
		e.RandomID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessage#ed18c118: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessage#ed18c118: field date: %w", err)
		}
		e.Date = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessage#ed18c118: field bytes: %w", err)
		}
		e.Bytes = value
	}
	{
		value, err := DecodeEncryptedFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessage#ed18c118: field file: %w", err)
		}
		e.File = value
	}
	return nil
}

// construct implements constructor of EncryptedMessageClass.
func (e EncryptedMessage) construct() EncryptedMessageClass { return &e }

// Ensuring interfaces in compile-time for EncryptedMessage.
var (
	_ bin.Encoder = &EncryptedMessage{}
	_ bin.Decoder = &EncryptedMessage{}

	_ EncryptedMessageClass = &EncryptedMessage{}
)

// EncryptedMessageService represents TL type `encryptedMessageService#23734b06`.
//
// See https://core.telegram.org/constructor/encryptedMessageService for reference.
type EncryptedMessageService struct {
	// RandomID field of EncryptedMessageService.
	RandomID int64
	// ChatID field of EncryptedMessageService.
	ChatID int
	// Date field of EncryptedMessageService.
	Date int
	// Bytes field of EncryptedMessageService.
	Bytes []byte
}

// EncryptedMessageServiceTypeID is TL type id of EncryptedMessageService.
const EncryptedMessageServiceTypeID = 0x23734b06

// Encode implements bin.Encoder.
func (e *EncryptedMessageService) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedMessageService#23734b06 as nil")
	}
	b.PutID(EncryptedMessageServiceTypeID)
	b.PutLong(e.RandomID)
	b.PutInt(e.ChatID)
	b.PutInt(e.Date)
	b.PutBytes(e.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedMessageService) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedMessageService#23734b06 to nil")
	}
	if err := b.ConsumeID(EncryptedMessageServiceTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedMessageService#23734b06: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessageService#23734b06: field random_id: %w", err)
		}
		e.RandomID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessageService#23734b06: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessageService#23734b06: field date: %w", err)
		}
		e.Date = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessageService#23734b06: field bytes: %w", err)
		}
		e.Bytes = value
	}
	return nil
}

// construct implements constructor of EncryptedMessageClass.
func (e EncryptedMessageService) construct() EncryptedMessageClass { return &e }

// Ensuring interfaces in compile-time for EncryptedMessageService.
var (
	_ bin.Encoder = &EncryptedMessageService{}
	_ bin.Decoder = &EncryptedMessageService{}

	_ EncryptedMessageClass = &EncryptedMessageService{}
)

// EncryptedMessageClass represents EncryptedMessage generic type.
//
// See https://core.telegram.org/type/EncryptedMessage for reference.
//
// Example:
//  g, err := DecodeEncryptedMessage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *EncryptedMessage: // encryptedMessage#ed18c118
//  case *EncryptedMessageService: // encryptedMessageService#23734b06
//  default: panic(v)
//  }
type EncryptedMessageClass interface {
	bin.Encoder
	bin.Decoder
	construct() EncryptedMessageClass
}

// DecodeEncryptedMessage implements binary de-serialization for EncryptedMessageClass.
func DecodeEncryptedMessage(buf *bin.Buffer) (EncryptedMessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case EncryptedMessageTypeID:
		// Decoding encryptedMessage#ed18c118.
		v := EncryptedMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedMessageClass: %w", err)
		}
		return &v, nil
	case EncryptedMessageServiceTypeID:
		// Decoding encryptedMessageService#23734b06.
		v := EncryptedMessageService{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedMessageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EncryptedMessageClass: %w", bin.NewUnexpectedID(id))
	}
}

// EncryptedMessage boxes the EncryptedMessageClass providing a helper.
type EncryptedMessageBox struct {
	EncryptedMessage EncryptedMessageClass
}

// Decode implements bin.Decoder for EncryptedMessageBox.
func (b *EncryptedMessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode EncryptedMessageBox to nil")
	}
	v, err := DecodeEncryptedMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EncryptedMessage = v
	return nil
}

// Encode implements bin.Encode for EncryptedMessageBox.
func (b *EncryptedMessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EncryptedMessage == nil {
		return fmt.Errorf("unable to encode EncryptedMessageClass as nil")
	}
	return b.EncryptedMessage.Encode(buf)
}
