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

// SecureValueErrorData represents TL type `secureValueErrorData#e8a40bd9`.
//
// See https://core.telegram.org/constructor/secureValueErrorData for reference.
type SecureValueErrorData struct {
	// Type field of SecureValueErrorData.
	Type SecureValueTypeClass
	// DataHash field of SecureValueErrorData.
	DataHash []byte
	// Field field of SecureValueErrorData.
	Field string
	// Text field of SecureValueErrorData.
	Text string
}

// SecureValueErrorDataTypeID is TL type id of SecureValueErrorData.
const SecureValueErrorDataTypeID = 0xe8a40bd9

// Encode implements bin.Encoder.
func (s *SecureValueErrorData) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueErrorData#e8a40bd9 as nil")
	}
	b.PutID(SecureValueErrorDataTypeID)
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureValueErrorData#e8a40bd9: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValueErrorData#e8a40bd9: field type: %w", err)
	}
	b.PutBytes(s.DataHash)
	b.PutString(s.Field)
	b.PutString(s.Text)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueErrorData) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueErrorData#e8a40bd9 to nil")
	}
	if err := b.ConsumeID(SecureValueErrorDataTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueErrorData#e8a40bd9: %w", err)
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorData#e8a40bd9: field type: %w", err)
		}
		s.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorData#e8a40bd9: field data_hash: %w", err)
		}
		s.DataHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorData#e8a40bd9: field field: %w", err)
		}
		s.Field = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorData#e8a40bd9: field text: %w", err)
		}
		s.Text = value
	}
	return nil
}

// construct implements constructor of SecureValueErrorClass.
func (s SecureValueErrorData) construct() SecureValueErrorClass { return &s }

// Ensuring interfaces in compile-time for SecureValueErrorData.
var (
	_ bin.Encoder = &SecureValueErrorData{}
	_ bin.Decoder = &SecureValueErrorData{}

	_ SecureValueErrorClass = &SecureValueErrorData{}
)

// SecureValueErrorFrontSide represents TL type `secureValueErrorFrontSide#be3dfa`.
//
// See https://core.telegram.org/constructor/secureValueErrorFrontSide for reference.
type SecureValueErrorFrontSide struct {
	// Type field of SecureValueErrorFrontSide.
	Type SecureValueTypeClass
	// FileHash field of SecureValueErrorFrontSide.
	FileHash []byte
	// Text field of SecureValueErrorFrontSide.
	Text string
}

// SecureValueErrorFrontSideTypeID is TL type id of SecureValueErrorFrontSide.
const SecureValueErrorFrontSideTypeID = 0xbe3dfa

// Encode implements bin.Encoder.
func (s *SecureValueErrorFrontSide) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueErrorFrontSide#be3dfa as nil")
	}
	b.PutID(SecureValueErrorFrontSideTypeID)
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureValueErrorFrontSide#be3dfa: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValueErrorFrontSide#be3dfa: field type: %w", err)
	}
	b.PutBytes(s.FileHash)
	b.PutString(s.Text)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueErrorFrontSide) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueErrorFrontSide#be3dfa to nil")
	}
	if err := b.ConsumeID(SecureValueErrorFrontSideTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueErrorFrontSide#be3dfa: %w", err)
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorFrontSide#be3dfa: field type: %w", err)
		}
		s.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorFrontSide#be3dfa: field file_hash: %w", err)
		}
		s.FileHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorFrontSide#be3dfa: field text: %w", err)
		}
		s.Text = value
	}
	return nil
}

// construct implements constructor of SecureValueErrorClass.
func (s SecureValueErrorFrontSide) construct() SecureValueErrorClass { return &s }

// Ensuring interfaces in compile-time for SecureValueErrorFrontSide.
var (
	_ bin.Encoder = &SecureValueErrorFrontSide{}
	_ bin.Decoder = &SecureValueErrorFrontSide{}

	_ SecureValueErrorClass = &SecureValueErrorFrontSide{}
)

// SecureValueErrorReverseSide represents TL type `secureValueErrorReverseSide#868a2aa5`.
//
// See https://core.telegram.org/constructor/secureValueErrorReverseSide for reference.
type SecureValueErrorReverseSide struct {
	// Type field of SecureValueErrorReverseSide.
	Type SecureValueTypeClass
	// FileHash field of SecureValueErrorReverseSide.
	FileHash []byte
	// Text field of SecureValueErrorReverseSide.
	Text string
}

// SecureValueErrorReverseSideTypeID is TL type id of SecureValueErrorReverseSide.
const SecureValueErrorReverseSideTypeID = 0x868a2aa5

// Encode implements bin.Encoder.
func (s *SecureValueErrorReverseSide) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueErrorReverseSide#868a2aa5 as nil")
	}
	b.PutID(SecureValueErrorReverseSideTypeID)
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureValueErrorReverseSide#868a2aa5: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValueErrorReverseSide#868a2aa5: field type: %w", err)
	}
	b.PutBytes(s.FileHash)
	b.PutString(s.Text)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueErrorReverseSide) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueErrorReverseSide#868a2aa5 to nil")
	}
	if err := b.ConsumeID(SecureValueErrorReverseSideTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueErrorReverseSide#868a2aa5: %w", err)
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorReverseSide#868a2aa5: field type: %w", err)
		}
		s.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorReverseSide#868a2aa5: field file_hash: %w", err)
		}
		s.FileHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorReverseSide#868a2aa5: field text: %w", err)
		}
		s.Text = value
	}
	return nil
}

// construct implements constructor of SecureValueErrorClass.
func (s SecureValueErrorReverseSide) construct() SecureValueErrorClass { return &s }

// Ensuring interfaces in compile-time for SecureValueErrorReverseSide.
var (
	_ bin.Encoder = &SecureValueErrorReverseSide{}
	_ bin.Decoder = &SecureValueErrorReverseSide{}

	_ SecureValueErrorClass = &SecureValueErrorReverseSide{}
)

// SecureValueErrorSelfie represents TL type `secureValueErrorSelfie#e537ced6`.
//
// See https://core.telegram.org/constructor/secureValueErrorSelfie for reference.
type SecureValueErrorSelfie struct {
	// Type field of SecureValueErrorSelfie.
	Type SecureValueTypeClass
	// FileHash field of SecureValueErrorSelfie.
	FileHash []byte
	// Text field of SecureValueErrorSelfie.
	Text string
}

// SecureValueErrorSelfieTypeID is TL type id of SecureValueErrorSelfie.
const SecureValueErrorSelfieTypeID = 0xe537ced6

// Encode implements bin.Encoder.
func (s *SecureValueErrorSelfie) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueErrorSelfie#e537ced6 as nil")
	}
	b.PutID(SecureValueErrorSelfieTypeID)
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureValueErrorSelfie#e537ced6: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValueErrorSelfie#e537ced6: field type: %w", err)
	}
	b.PutBytes(s.FileHash)
	b.PutString(s.Text)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueErrorSelfie) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueErrorSelfie#e537ced6 to nil")
	}
	if err := b.ConsumeID(SecureValueErrorSelfieTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueErrorSelfie#e537ced6: %w", err)
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorSelfie#e537ced6: field type: %w", err)
		}
		s.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorSelfie#e537ced6: field file_hash: %w", err)
		}
		s.FileHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorSelfie#e537ced6: field text: %w", err)
		}
		s.Text = value
	}
	return nil
}

// construct implements constructor of SecureValueErrorClass.
func (s SecureValueErrorSelfie) construct() SecureValueErrorClass { return &s }

// Ensuring interfaces in compile-time for SecureValueErrorSelfie.
var (
	_ bin.Encoder = &SecureValueErrorSelfie{}
	_ bin.Decoder = &SecureValueErrorSelfie{}

	_ SecureValueErrorClass = &SecureValueErrorSelfie{}
)

// SecureValueErrorFile represents TL type `secureValueErrorFile#7a700873`.
//
// See https://core.telegram.org/constructor/secureValueErrorFile for reference.
type SecureValueErrorFile struct {
	// Type field of SecureValueErrorFile.
	Type SecureValueTypeClass
	// FileHash field of SecureValueErrorFile.
	FileHash []byte
	// Text field of SecureValueErrorFile.
	Text string
}

// SecureValueErrorFileTypeID is TL type id of SecureValueErrorFile.
const SecureValueErrorFileTypeID = 0x7a700873

// Encode implements bin.Encoder.
func (s *SecureValueErrorFile) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueErrorFile#7a700873 as nil")
	}
	b.PutID(SecureValueErrorFileTypeID)
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureValueErrorFile#7a700873: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValueErrorFile#7a700873: field type: %w", err)
	}
	b.PutBytes(s.FileHash)
	b.PutString(s.Text)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueErrorFile) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueErrorFile#7a700873 to nil")
	}
	if err := b.ConsumeID(SecureValueErrorFileTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueErrorFile#7a700873: %w", err)
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorFile#7a700873: field type: %w", err)
		}
		s.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorFile#7a700873: field file_hash: %w", err)
		}
		s.FileHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorFile#7a700873: field text: %w", err)
		}
		s.Text = value
	}
	return nil
}

// construct implements constructor of SecureValueErrorClass.
func (s SecureValueErrorFile) construct() SecureValueErrorClass { return &s }

// Ensuring interfaces in compile-time for SecureValueErrorFile.
var (
	_ bin.Encoder = &SecureValueErrorFile{}
	_ bin.Decoder = &SecureValueErrorFile{}

	_ SecureValueErrorClass = &SecureValueErrorFile{}
)

// SecureValueErrorFiles represents TL type `secureValueErrorFiles#666220e9`.
//
// See https://core.telegram.org/constructor/secureValueErrorFiles for reference.
type SecureValueErrorFiles struct {
	// Type field of SecureValueErrorFiles.
	Type SecureValueTypeClass
	// FileHash field of SecureValueErrorFiles.
	FileHash [][]byte
	// Text field of SecureValueErrorFiles.
	Text string
}

// SecureValueErrorFilesTypeID is TL type id of SecureValueErrorFiles.
const SecureValueErrorFilesTypeID = 0x666220e9

// Encode implements bin.Encoder.
func (s *SecureValueErrorFiles) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueErrorFiles#666220e9 as nil")
	}
	b.PutID(SecureValueErrorFilesTypeID)
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureValueErrorFiles#666220e9: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValueErrorFiles#666220e9: field type: %w", err)
	}
	b.PutVectorHeader(len(s.FileHash))
	for _, v := range s.FileHash {
		b.PutBytes(v)
	}
	b.PutString(s.Text)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueErrorFiles) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueErrorFiles#666220e9 to nil")
	}
	if err := b.ConsumeID(SecureValueErrorFilesTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueErrorFiles#666220e9: %w", err)
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorFiles#666220e9: field type: %w", err)
		}
		s.Type = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorFiles#666220e9: field file_hash: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode secureValueErrorFiles#666220e9: field file_hash: %w", err)
			}
			s.FileHash = append(s.FileHash, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorFiles#666220e9: field text: %w", err)
		}
		s.Text = value
	}
	return nil
}

// construct implements constructor of SecureValueErrorClass.
func (s SecureValueErrorFiles) construct() SecureValueErrorClass { return &s }

// Ensuring interfaces in compile-time for SecureValueErrorFiles.
var (
	_ bin.Encoder = &SecureValueErrorFiles{}
	_ bin.Decoder = &SecureValueErrorFiles{}

	_ SecureValueErrorClass = &SecureValueErrorFiles{}
)

// SecureValueError represents TL type `secureValueError#869d758f`.
//
// See https://core.telegram.org/constructor/secureValueError for reference.
type SecureValueError struct {
	// Type field of SecureValueError.
	Type SecureValueTypeClass
	// Hash field of SecureValueError.
	Hash []byte
	// Text field of SecureValueError.
	Text string
}

// SecureValueErrorTypeID is TL type id of SecureValueError.
const SecureValueErrorTypeID = 0x869d758f

// Encode implements bin.Encoder.
func (s *SecureValueError) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueError#869d758f as nil")
	}
	b.PutID(SecureValueErrorTypeID)
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureValueError#869d758f: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValueError#869d758f: field type: %w", err)
	}
	b.PutBytes(s.Hash)
	b.PutString(s.Text)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueError) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueError#869d758f to nil")
	}
	if err := b.ConsumeID(SecureValueErrorTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueError#869d758f: %w", err)
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValueError#869d758f: field type: %w", err)
		}
		s.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueError#869d758f: field hash: %w", err)
		}
		s.Hash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueError#869d758f: field text: %w", err)
		}
		s.Text = value
	}
	return nil
}

// construct implements constructor of SecureValueErrorClass.
func (s SecureValueError) construct() SecureValueErrorClass { return &s }

// Ensuring interfaces in compile-time for SecureValueError.
var (
	_ bin.Encoder = &SecureValueError{}
	_ bin.Decoder = &SecureValueError{}

	_ SecureValueErrorClass = &SecureValueError{}
)

// SecureValueErrorTranslationFile represents TL type `secureValueErrorTranslationFile#a1144770`.
//
// See https://core.telegram.org/constructor/secureValueErrorTranslationFile for reference.
type SecureValueErrorTranslationFile struct {
	// Type field of SecureValueErrorTranslationFile.
	Type SecureValueTypeClass
	// FileHash field of SecureValueErrorTranslationFile.
	FileHash []byte
	// Text field of SecureValueErrorTranslationFile.
	Text string
}

// SecureValueErrorTranslationFileTypeID is TL type id of SecureValueErrorTranslationFile.
const SecureValueErrorTranslationFileTypeID = 0xa1144770

// Encode implements bin.Encoder.
func (s *SecureValueErrorTranslationFile) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueErrorTranslationFile#a1144770 as nil")
	}
	b.PutID(SecureValueErrorTranslationFileTypeID)
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureValueErrorTranslationFile#a1144770: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValueErrorTranslationFile#a1144770: field type: %w", err)
	}
	b.PutBytes(s.FileHash)
	b.PutString(s.Text)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueErrorTranslationFile) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueErrorTranslationFile#a1144770 to nil")
	}
	if err := b.ConsumeID(SecureValueErrorTranslationFileTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueErrorTranslationFile#a1144770: %w", err)
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorTranslationFile#a1144770: field type: %w", err)
		}
		s.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorTranslationFile#a1144770: field file_hash: %w", err)
		}
		s.FileHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorTranslationFile#a1144770: field text: %w", err)
		}
		s.Text = value
	}
	return nil
}

// construct implements constructor of SecureValueErrorClass.
func (s SecureValueErrorTranslationFile) construct() SecureValueErrorClass { return &s }

// Ensuring interfaces in compile-time for SecureValueErrorTranslationFile.
var (
	_ bin.Encoder = &SecureValueErrorTranslationFile{}
	_ bin.Decoder = &SecureValueErrorTranslationFile{}

	_ SecureValueErrorClass = &SecureValueErrorTranslationFile{}
)

// SecureValueErrorTranslationFiles represents TL type `secureValueErrorTranslationFiles#34636dd8`.
//
// See https://core.telegram.org/constructor/secureValueErrorTranslationFiles for reference.
type SecureValueErrorTranslationFiles struct {
	// Type field of SecureValueErrorTranslationFiles.
	Type SecureValueTypeClass
	// FileHash field of SecureValueErrorTranslationFiles.
	FileHash [][]byte
	// Text field of SecureValueErrorTranslationFiles.
	Text string
}

// SecureValueErrorTranslationFilesTypeID is TL type id of SecureValueErrorTranslationFiles.
const SecureValueErrorTranslationFilesTypeID = 0x34636dd8

// Encode implements bin.Encoder.
func (s *SecureValueErrorTranslationFiles) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueErrorTranslationFiles#34636dd8 as nil")
	}
	b.PutID(SecureValueErrorTranslationFilesTypeID)
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureValueErrorTranslationFiles#34636dd8: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValueErrorTranslationFiles#34636dd8: field type: %w", err)
	}
	b.PutVectorHeader(len(s.FileHash))
	for _, v := range s.FileHash {
		b.PutBytes(v)
	}
	b.PutString(s.Text)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueErrorTranslationFiles) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueErrorTranslationFiles#34636dd8 to nil")
	}
	if err := b.ConsumeID(SecureValueErrorTranslationFilesTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueErrorTranslationFiles#34636dd8: %w", err)
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorTranslationFiles#34636dd8: field type: %w", err)
		}
		s.Type = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorTranslationFiles#34636dd8: field file_hash: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode secureValueErrorTranslationFiles#34636dd8: field file_hash: %w", err)
			}
			s.FileHash = append(s.FileHash, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueErrorTranslationFiles#34636dd8: field text: %w", err)
		}
		s.Text = value
	}
	return nil
}

// construct implements constructor of SecureValueErrorClass.
func (s SecureValueErrorTranslationFiles) construct() SecureValueErrorClass { return &s }

// Ensuring interfaces in compile-time for SecureValueErrorTranslationFiles.
var (
	_ bin.Encoder = &SecureValueErrorTranslationFiles{}
	_ bin.Decoder = &SecureValueErrorTranslationFiles{}

	_ SecureValueErrorClass = &SecureValueErrorTranslationFiles{}
)

// SecureValueErrorClass represents SecureValueError generic type.
//
// See https://core.telegram.org/type/SecureValueError for reference.
//
// Example:
//  g, err := DecodeSecureValueError(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *SecureValueErrorData: // secureValueErrorData#e8a40bd9
//  case *SecureValueErrorFrontSide: // secureValueErrorFrontSide#be3dfa
//  case *SecureValueErrorReverseSide: // secureValueErrorReverseSide#868a2aa5
//  case *SecureValueErrorSelfie: // secureValueErrorSelfie#e537ced6
//  case *SecureValueErrorFile: // secureValueErrorFile#7a700873
//  case *SecureValueErrorFiles: // secureValueErrorFiles#666220e9
//  case *SecureValueError: // secureValueError#869d758f
//  case *SecureValueErrorTranslationFile: // secureValueErrorTranslationFile#a1144770
//  case *SecureValueErrorTranslationFiles: // secureValueErrorTranslationFiles#34636dd8
//  default: panic(v)
//  }
type SecureValueErrorClass interface {
	bin.Encoder
	bin.Decoder
	construct() SecureValueErrorClass
}

// DecodeSecureValueError implements binary de-serialization for SecureValueErrorClass.
func DecodeSecureValueError(buf *bin.Buffer) (SecureValueErrorClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SecureValueErrorDataTypeID:
		// Decoding secureValueErrorData#e8a40bd9.
		v := SecureValueErrorData{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueErrorClass: %w", err)
		}
		return &v, nil
	case SecureValueErrorFrontSideTypeID:
		// Decoding secureValueErrorFrontSide#be3dfa.
		v := SecureValueErrorFrontSide{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueErrorClass: %w", err)
		}
		return &v, nil
	case SecureValueErrorReverseSideTypeID:
		// Decoding secureValueErrorReverseSide#868a2aa5.
		v := SecureValueErrorReverseSide{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueErrorClass: %w", err)
		}
		return &v, nil
	case SecureValueErrorSelfieTypeID:
		// Decoding secureValueErrorSelfie#e537ced6.
		v := SecureValueErrorSelfie{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueErrorClass: %w", err)
		}
		return &v, nil
	case SecureValueErrorFileTypeID:
		// Decoding secureValueErrorFile#7a700873.
		v := SecureValueErrorFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueErrorClass: %w", err)
		}
		return &v, nil
	case SecureValueErrorFilesTypeID:
		// Decoding secureValueErrorFiles#666220e9.
		v := SecureValueErrorFiles{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueErrorClass: %w", err)
		}
		return &v, nil
	case SecureValueErrorTypeID:
		// Decoding secureValueError#869d758f.
		v := SecureValueError{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueErrorClass: %w", err)
		}
		return &v, nil
	case SecureValueErrorTranslationFileTypeID:
		// Decoding secureValueErrorTranslationFile#a1144770.
		v := SecureValueErrorTranslationFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueErrorClass: %w", err)
		}
		return &v, nil
	case SecureValueErrorTranslationFilesTypeID:
		// Decoding secureValueErrorTranslationFiles#34636dd8.
		v := SecureValueErrorTranslationFiles{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueErrorClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SecureValueErrorClass: %w", bin.NewUnexpectedID(id))
	}
}

// SecureValueError boxes the SecureValueErrorClass providing a helper.
type SecureValueErrorBox struct {
	SecureValueError SecureValueErrorClass
}

// Decode implements bin.Decoder for SecureValueErrorBox.
func (b *SecureValueErrorBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SecureValueErrorBox to nil")
	}
	v, err := DecodeSecureValueError(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SecureValueError = v
	return nil
}

// Encode implements bin.Encode for SecureValueErrorBox.
func (b *SecureValueErrorBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SecureValueError == nil {
		return fmt.Errorf("unable to encode SecureValueErrorClass as nil")
	}
	return b.SecureValueError.Encode(buf)
}
