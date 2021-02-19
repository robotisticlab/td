// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// DocumentAttributeImageSize represents TL type `documentAttributeImageSize#6c37c15c`.
// Defines the width and height of an image uploaded as document
//
// See https://core.telegram.org/constructor/documentAttributeImageSize for reference.
type DocumentAttributeImageSize struct {
	// Width of image
	W int
	// Height of image
	H int
}

// DocumentAttributeImageSizeTypeID is TL type id of DocumentAttributeImageSize.
const DocumentAttributeImageSizeTypeID = 0x6c37c15c

func (d *DocumentAttributeImageSize) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.W == 0) {
		return false
	}
	if !(d.H == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DocumentAttributeImageSize) String() string {
	if d == nil {
		return "DocumentAttributeImageSize(nil)"
	}
	type Alias DocumentAttributeImageSize
	return fmt.Sprintf("DocumentAttributeImageSize%+v", Alias(*d))
}

// FillFrom fills DocumentAttributeImageSize from given interface.
func (d *DocumentAttributeImageSize) FillFrom(from interface {
	GetW() (value int)
	GetH() (value int)
}) {
	d.W = from.GetW()
	d.H = from.GetH()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DocumentAttributeImageSize) TypeID() uint32 {
	return DocumentAttributeImageSizeTypeID
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeImageSize) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeImageSize#6c37c15c as nil")
	}
	b.PutID(DocumentAttributeImageSizeTypeID)
	b.PutInt(d.W)
	b.PutInt(d.H)
	return nil
}

// GetW returns value of W field.
func (d *DocumentAttributeImageSize) GetW() (value int) {
	return d.W
}

// GetH returns value of H field.
func (d *DocumentAttributeImageSize) GetH() (value int) {
	return d.H
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeImageSize) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeImageSize#6c37c15c to nil")
	}
	if err := b.ConsumeID(DocumentAttributeImageSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeImageSize#6c37c15c: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeImageSize#6c37c15c: field w: %w", err)
		}
		d.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeImageSize#6c37c15c: field h: %w", err)
		}
		d.H = value
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeImageSize) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeImageSize.
var (
	_ bin.Encoder = &DocumentAttributeImageSize{}
	_ bin.Decoder = &DocumentAttributeImageSize{}

	_ DocumentAttributeClass = &DocumentAttributeImageSize{}
)

// DocumentAttributeAnimated represents TL type `documentAttributeAnimated#11b58939`.
// Defines an animated GIF
//
// See https://core.telegram.org/constructor/documentAttributeAnimated for reference.
type DocumentAttributeAnimated struct {
}

// DocumentAttributeAnimatedTypeID is TL type id of DocumentAttributeAnimated.
const DocumentAttributeAnimatedTypeID = 0x11b58939

func (d *DocumentAttributeAnimated) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *DocumentAttributeAnimated) String() string {
	if d == nil {
		return "DocumentAttributeAnimated(nil)"
	}
	type Alias DocumentAttributeAnimated
	return fmt.Sprintf("DocumentAttributeAnimated%+v", Alias(*d))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DocumentAttributeAnimated) TypeID() uint32 {
	return DocumentAttributeAnimatedTypeID
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeAnimated) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeAnimated#11b58939 as nil")
	}
	b.PutID(DocumentAttributeAnimatedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeAnimated) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeAnimated#11b58939 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeAnimatedTypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeAnimated#11b58939: %w", err)
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeAnimated) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeAnimated.
var (
	_ bin.Encoder = &DocumentAttributeAnimated{}
	_ bin.Decoder = &DocumentAttributeAnimated{}

	_ DocumentAttributeClass = &DocumentAttributeAnimated{}
)

// DocumentAttributeSticker23 represents TL type `documentAttributeSticker23#fb0a5727`.
//
// See https://core.telegram.org/constructor/documentAttributeSticker23 for reference.
type DocumentAttributeSticker23 struct {
}

// DocumentAttributeSticker23TypeID is TL type id of DocumentAttributeSticker23.
const DocumentAttributeSticker23TypeID = 0xfb0a5727

func (d *DocumentAttributeSticker23) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *DocumentAttributeSticker23) String() string {
	if d == nil {
		return "DocumentAttributeSticker23(nil)"
	}
	type Alias DocumentAttributeSticker23
	return fmt.Sprintf("DocumentAttributeSticker23%+v", Alias(*d))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DocumentAttributeSticker23) TypeID() uint32 {
	return DocumentAttributeSticker23TypeID
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeSticker23) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeSticker23#fb0a5727 as nil")
	}
	b.PutID(DocumentAttributeSticker23TypeID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeSticker23) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeSticker23#fb0a5727 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeSticker23TypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeSticker23#fb0a5727: %w", err)
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeSticker23) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeSticker23.
var (
	_ bin.Encoder = &DocumentAttributeSticker23{}
	_ bin.Decoder = &DocumentAttributeSticker23{}

	_ DocumentAttributeClass = &DocumentAttributeSticker23{}
)

// DocumentAttributeVideo represents TL type `documentAttributeVideo#5910cccb`.
// Defines a video
//
// See https://core.telegram.org/constructor/documentAttributeVideo for reference.
type DocumentAttributeVideo struct {
	// Duration in seconds
	Duration int
	// Video width
	W int
	// Video height
	H int
}

// DocumentAttributeVideoTypeID is TL type id of DocumentAttributeVideo.
const DocumentAttributeVideoTypeID = 0x5910cccb

func (d *DocumentAttributeVideo) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Duration == 0) {
		return false
	}
	if !(d.W == 0) {
		return false
	}
	if !(d.H == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DocumentAttributeVideo) String() string {
	if d == nil {
		return "DocumentAttributeVideo(nil)"
	}
	type Alias DocumentAttributeVideo
	return fmt.Sprintf("DocumentAttributeVideo%+v", Alias(*d))
}

// FillFrom fills DocumentAttributeVideo from given interface.
func (d *DocumentAttributeVideo) FillFrom(from interface {
	GetDuration() (value int)
	GetW() (value int)
	GetH() (value int)
}) {
	d.Duration = from.GetDuration()
	d.W = from.GetW()
	d.H = from.GetH()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DocumentAttributeVideo) TypeID() uint32 {
	return DocumentAttributeVideoTypeID
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeVideo) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeVideo#5910cccb as nil")
	}
	b.PutID(DocumentAttributeVideoTypeID)
	b.PutInt(d.Duration)
	b.PutInt(d.W)
	b.PutInt(d.H)
	return nil
}

// GetDuration returns value of Duration field.
func (d *DocumentAttributeVideo) GetDuration() (value int) {
	return d.Duration
}

// GetW returns value of W field.
func (d *DocumentAttributeVideo) GetW() (value int) {
	return d.W
}

// GetH returns value of H field.
func (d *DocumentAttributeVideo) GetH() (value int) {
	return d.H
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeVideo) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeVideo#5910cccb to nil")
	}
	if err := b.ConsumeID(DocumentAttributeVideoTypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeVideo#5910cccb: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeVideo#5910cccb: field duration: %w", err)
		}
		d.Duration = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeVideo#5910cccb: field w: %w", err)
		}
		d.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeVideo#5910cccb: field h: %w", err)
		}
		d.H = value
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeVideo) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeVideo.
var (
	_ bin.Encoder = &DocumentAttributeVideo{}
	_ bin.Decoder = &DocumentAttributeVideo{}

	_ DocumentAttributeClass = &DocumentAttributeVideo{}
)

// DocumentAttributeAudio23 represents TL type `documentAttributeAudio23#51448e5`.
//
// See https://core.telegram.org/constructor/documentAttributeAudio23 for reference.
type DocumentAttributeAudio23 struct {
	// Duration field of DocumentAttributeAudio23.
	Duration int
}

// DocumentAttributeAudio23TypeID is TL type id of DocumentAttributeAudio23.
const DocumentAttributeAudio23TypeID = 0x51448e5

func (d *DocumentAttributeAudio23) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Duration == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DocumentAttributeAudio23) String() string {
	if d == nil {
		return "DocumentAttributeAudio23(nil)"
	}
	type Alias DocumentAttributeAudio23
	return fmt.Sprintf("DocumentAttributeAudio23%+v", Alias(*d))
}

// FillFrom fills DocumentAttributeAudio23 from given interface.
func (d *DocumentAttributeAudio23) FillFrom(from interface {
	GetDuration() (value int)
}) {
	d.Duration = from.GetDuration()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DocumentAttributeAudio23) TypeID() uint32 {
	return DocumentAttributeAudio23TypeID
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeAudio23) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeAudio23#51448e5 as nil")
	}
	b.PutID(DocumentAttributeAudio23TypeID)
	b.PutInt(d.Duration)
	return nil
}

// GetDuration returns value of Duration field.
func (d *DocumentAttributeAudio23) GetDuration() (value int) {
	return d.Duration
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeAudio23) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeAudio23#51448e5 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeAudio23TypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeAudio23#51448e5: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio23#51448e5: field duration: %w", err)
		}
		d.Duration = value
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeAudio23) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeAudio23.
var (
	_ bin.Encoder = &DocumentAttributeAudio23{}
	_ bin.Decoder = &DocumentAttributeAudio23{}

	_ DocumentAttributeClass = &DocumentAttributeAudio23{}
)

// DocumentAttributeFilename represents TL type `documentAttributeFilename#15590068`.
// A simple document with a file name
//
// See https://core.telegram.org/constructor/documentAttributeFilename for reference.
type DocumentAttributeFilename struct {
	// The file name
	FileName string
}

// DocumentAttributeFilenameTypeID is TL type id of DocumentAttributeFilename.
const DocumentAttributeFilenameTypeID = 0x15590068

func (d *DocumentAttributeFilename) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.FileName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DocumentAttributeFilename) String() string {
	if d == nil {
		return "DocumentAttributeFilename(nil)"
	}
	type Alias DocumentAttributeFilename
	return fmt.Sprintf("DocumentAttributeFilename%+v", Alias(*d))
}

// FillFrom fills DocumentAttributeFilename from given interface.
func (d *DocumentAttributeFilename) FillFrom(from interface {
	GetFileName() (value string)
}) {
	d.FileName = from.GetFileName()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DocumentAttributeFilename) TypeID() uint32 {
	return DocumentAttributeFilenameTypeID
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeFilename) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeFilename#15590068 as nil")
	}
	b.PutID(DocumentAttributeFilenameTypeID)
	b.PutString(d.FileName)
	return nil
}

// GetFileName returns value of FileName field.
func (d *DocumentAttributeFilename) GetFileName() (value string) {
	return d.FileName
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeFilename) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeFilename#15590068 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeFilenameTypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeFilename#15590068: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeFilename#15590068: field file_name: %w", err)
		}
		d.FileName = value
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeFilename) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeFilename.
var (
	_ bin.Encoder = &DocumentAttributeFilename{}
	_ bin.Decoder = &DocumentAttributeFilename{}

	_ DocumentAttributeClass = &DocumentAttributeFilename{}
)

// DocumentAttributeAudio45 represents TL type `documentAttributeAudio45#ded218e0`.
//
// See https://core.telegram.org/constructor/documentAttributeAudio45 for reference.
type DocumentAttributeAudio45 struct {
	// Duration field of DocumentAttributeAudio45.
	Duration int
	// Title field of DocumentAttributeAudio45.
	Title string
	// Performer field of DocumentAttributeAudio45.
	Performer string
}

// DocumentAttributeAudio45TypeID is TL type id of DocumentAttributeAudio45.
const DocumentAttributeAudio45TypeID = 0xded218e0

func (d *DocumentAttributeAudio45) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Duration == 0) {
		return false
	}
	if !(d.Title == "") {
		return false
	}
	if !(d.Performer == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DocumentAttributeAudio45) String() string {
	if d == nil {
		return "DocumentAttributeAudio45(nil)"
	}
	type Alias DocumentAttributeAudio45
	return fmt.Sprintf("DocumentAttributeAudio45%+v", Alias(*d))
}

// FillFrom fills DocumentAttributeAudio45 from given interface.
func (d *DocumentAttributeAudio45) FillFrom(from interface {
	GetDuration() (value int)
	GetTitle() (value string)
	GetPerformer() (value string)
}) {
	d.Duration = from.GetDuration()
	d.Title = from.GetTitle()
	d.Performer = from.GetPerformer()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DocumentAttributeAudio45) TypeID() uint32 {
	return DocumentAttributeAudio45TypeID
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeAudio45) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeAudio45#ded218e0 as nil")
	}
	b.PutID(DocumentAttributeAudio45TypeID)
	b.PutInt(d.Duration)
	b.PutString(d.Title)
	b.PutString(d.Performer)
	return nil
}

// GetDuration returns value of Duration field.
func (d *DocumentAttributeAudio45) GetDuration() (value int) {
	return d.Duration
}

// GetTitle returns value of Title field.
func (d *DocumentAttributeAudio45) GetTitle() (value string) {
	return d.Title
}

// GetPerformer returns value of Performer field.
func (d *DocumentAttributeAudio45) GetPerformer() (value string) {
	return d.Performer
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeAudio45) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeAudio45#ded218e0 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeAudio45TypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeAudio45#ded218e0: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio45#ded218e0: field duration: %w", err)
		}
		d.Duration = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio45#ded218e0: field title: %w", err)
		}
		d.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio45#ded218e0: field performer: %w", err)
		}
		d.Performer = value
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeAudio45) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeAudio45.
var (
	_ bin.Encoder = &DocumentAttributeAudio45{}
	_ bin.Decoder = &DocumentAttributeAudio45{}

	_ DocumentAttributeClass = &DocumentAttributeAudio45{}
)

// DocumentAttributeSticker represents TL type `documentAttributeSticker#3a556302`.
// Defines a sticker
//
// See https://core.telegram.org/constructor/documentAttributeSticker for reference.
type DocumentAttributeSticker struct {
	// Alternative emoji representation of sticker
	Alt string
	// Associated stickerset
	Stickerset InputStickerSetClass
}

// DocumentAttributeStickerTypeID is TL type id of DocumentAttributeSticker.
const DocumentAttributeStickerTypeID = 0x3a556302

func (d *DocumentAttributeSticker) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Alt == "") {
		return false
	}
	if !(d.Stickerset == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DocumentAttributeSticker) String() string {
	if d == nil {
		return "DocumentAttributeSticker(nil)"
	}
	type Alias DocumentAttributeSticker
	return fmt.Sprintf("DocumentAttributeSticker%+v", Alias(*d))
}

// FillFrom fills DocumentAttributeSticker from given interface.
func (d *DocumentAttributeSticker) FillFrom(from interface {
	GetAlt() (value string)
	GetStickerset() (value InputStickerSetClass)
}) {
	d.Alt = from.GetAlt()
	d.Stickerset = from.GetStickerset()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DocumentAttributeSticker) TypeID() uint32 {
	return DocumentAttributeStickerTypeID
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeSticker) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeSticker#3a556302 as nil")
	}
	b.PutID(DocumentAttributeStickerTypeID)
	b.PutString(d.Alt)
	if d.Stickerset == nil {
		return fmt.Errorf("unable to encode documentAttributeSticker#3a556302: field stickerset is nil")
	}
	if err := d.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode documentAttributeSticker#3a556302: field stickerset: %w", err)
	}
	return nil
}

// GetAlt returns value of Alt field.
func (d *DocumentAttributeSticker) GetAlt() (value string) {
	return d.Alt
}

// GetStickerset returns value of Stickerset field.
func (d *DocumentAttributeSticker) GetStickerset() (value InputStickerSetClass) {
	return d.Stickerset
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeSticker) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeSticker#3a556302 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeStickerTypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeSticker#3a556302: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeSticker#3a556302: field alt: %w", err)
		}
		d.Alt = value
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeSticker#3a556302: field stickerset: %w", err)
		}
		d.Stickerset = value
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeSticker) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeSticker.
var (
	_ bin.Encoder = &DocumentAttributeSticker{}
	_ bin.Decoder = &DocumentAttributeSticker{}

	_ DocumentAttributeClass = &DocumentAttributeSticker{}
)

// DocumentAttributeAudio represents TL type `documentAttributeAudio#9852f9c6`.
// Represents an audio file
//
// See https://core.telegram.org/constructor/documentAttributeAudio for reference.
type DocumentAttributeAudio struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is a voice message
	Voice bool
	// Duration in seconds
	Duration int
	// Name of song
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// Performer
	//
	// Use SetPerformer and GetPerformer helpers.
	Performer string
	// Waveform
	//
	// Use SetWaveform and GetWaveform helpers.
	Waveform []byte
}

// DocumentAttributeAudioTypeID is TL type id of DocumentAttributeAudio.
const DocumentAttributeAudioTypeID = 0x9852f9c6

func (d *DocumentAttributeAudio) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.Voice == false) {
		return false
	}
	if !(d.Duration == 0) {
		return false
	}
	if !(d.Title == "") {
		return false
	}
	if !(d.Performer == "") {
		return false
	}
	if !(d.Waveform == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DocumentAttributeAudio) String() string {
	if d == nil {
		return "DocumentAttributeAudio(nil)"
	}
	type Alias DocumentAttributeAudio
	return fmt.Sprintf("DocumentAttributeAudio%+v", Alias(*d))
}

// FillFrom fills DocumentAttributeAudio from given interface.
func (d *DocumentAttributeAudio) FillFrom(from interface {
	GetVoice() (value bool)
	GetDuration() (value int)
	GetTitle() (value string, ok bool)
	GetPerformer() (value string, ok bool)
	GetWaveform() (value []byte, ok bool)
}) {
	d.Duration = from.GetDuration()
	if val, ok := from.GetTitle(); ok {
		d.Title = val
	}
	if val, ok := from.GetPerformer(); ok {
		d.Performer = val
	}
	if val, ok := from.GetWaveform(); ok {
		d.Waveform = val
	}
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DocumentAttributeAudio) TypeID() uint32 {
	return DocumentAttributeAudioTypeID
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeAudio) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeAudio#9852f9c6 as nil")
	}
	b.PutID(DocumentAttributeAudioTypeID)
	if !(d.Voice == false) {
		d.Flags.Set(10)
	}
	if !(d.Title == "") {
		d.Flags.Set(0)
	}
	if !(d.Performer == "") {
		d.Flags.Set(1)
	}
	if !(d.Waveform == nil) {
		d.Flags.Set(2)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode documentAttributeAudio#9852f9c6: field flags: %w", err)
	}
	b.PutInt(d.Duration)
	if d.Flags.Has(0) {
		b.PutString(d.Title)
	}
	if d.Flags.Has(1) {
		b.PutString(d.Performer)
	}
	if d.Flags.Has(2) {
		b.PutBytes(d.Waveform)
	}
	return nil
}

// SetVoice sets value of Voice conditional field.
func (d *DocumentAttributeAudio) SetVoice(value bool) {
	if value {
		d.Flags.Set(10)
		d.Voice = true
	} else {
		d.Flags.Unset(10)
		d.Voice = false
	}
}

// GetVoice returns value of Voice conditional field.
func (d *DocumentAttributeAudio) GetVoice() (value bool) {
	return d.Flags.Has(10)
}

// GetDuration returns value of Duration field.
func (d *DocumentAttributeAudio) GetDuration() (value int) {
	return d.Duration
}

// SetTitle sets value of Title conditional field.
func (d *DocumentAttributeAudio) SetTitle(value string) {
	d.Flags.Set(0)
	d.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (d *DocumentAttributeAudio) GetTitle() (value string, ok bool) {
	if !d.Flags.Has(0) {
		return value, false
	}
	return d.Title, true
}

// SetPerformer sets value of Performer conditional field.
func (d *DocumentAttributeAudio) SetPerformer(value string) {
	d.Flags.Set(1)
	d.Performer = value
}

// GetPerformer returns value of Performer conditional field and
// boolean which is true if field was set.
func (d *DocumentAttributeAudio) GetPerformer() (value string, ok bool) {
	if !d.Flags.Has(1) {
		return value, false
	}
	return d.Performer, true
}

// SetWaveform sets value of Waveform conditional field.
func (d *DocumentAttributeAudio) SetWaveform(value []byte) {
	d.Flags.Set(2)
	d.Waveform = value
}

// GetWaveform returns value of Waveform conditional field and
// boolean which is true if field was set.
func (d *DocumentAttributeAudio) GetWaveform() (value []byte, ok bool) {
	if !d.Flags.Has(2) {
		return value, false
	}
	return d.Waveform, true
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeAudio) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeAudio#9852f9c6 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeAudioTypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeAudio#9852f9c6: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio#9852f9c6: field flags: %w", err)
		}
	}
	d.Voice = d.Flags.Has(10)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio#9852f9c6: field duration: %w", err)
		}
		d.Duration = value
	}
	if d.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio#9852f9c6: field title: %w", err)
		}
		d.Title = value
	}
	if d.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio#9852f9c6: field performer: %w", err)
		}
		d.Performer = value
	}
	if d.Flags.Has(2) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio#9852f9c6: field waveform: %w", err)
		}
		d.Waveform = value
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeAudio) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeAudio.
var (
	_ bin.Encoder = &DocumentAttributeAudio{}
	_ bin.Decoder = &DocumentAttributeAudio{}

	_ DocumentAttributeClass = &DocumentAttributeAudio{}
)

// DocumentAttributeVideo66 represents TL type `documentAttributeVideo66#ef02ce6`.
//
// See https://core.telegram.org/constructor/documentAttributeVideo66 for reference.
type DocumentAttributeVideo66 struct {
	// Flags field of DocumentAttributeVideo66.
	Flags bin.Fields
	// RoundMessage field of DocumentAttributeVideo66.
	RoundMessage bool
	// Duration field of DocumentAttributeVideo66.
	Duration int
	// W field of DocumentAttributeVideo66.
	W int
	// H field of DocumentAttributeVideo66.
	H int
}

// DocumentAttributeVideo66TypeID is TL type id of DocumentAttributeVideo66.
const DocumentAttributeVideo66TypeID = 0xef02ce6

func (d *DocumentAttributeVideo66) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.RoundMessage == false) {
		return false
	}
	if !(d.Duration == 0) {
		return false
	}
	if !(d.W == 0) {
		return false
	}
	if !(d.H == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DocumentAttributeVideo66) String() string {
	if d == nil {
		return "DocumentAttributeVideo66(nil)"
	}
	type Alias DocumentAttributeVideo66
	return fmt.Sprintf("DocumentAttributeVideo66%+v", Alias(*d))
}

// FillFrom fills DocumentAttributeVideo66 from given interface.
func (d *DocumentAttributeVideo66) FillFrom(from interface {
	GetRoundMessage() (value bool)
	GetDuration() (value int)
	GetW() (value int)
	GetH() (value int)
}) {
	d.Duration = from.GetDuration()
	d.W = from.GetW()
	d.H = from.GetH()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DocumentAttributeVideo66) TypeID() uint32 {
	return DocumentAttributeVideo66TypeID
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeVideo66) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeVideo66#ef02ce6 as nil")
	}
	b.PutID(DocumentAttributeVideo66TypeID)
	if !(d.RoundMessage == false) {
		d.Flags.Set(0)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode documentAttributeVideo66#ef02ce6: field flags: %w", err)
	}
	b.PutInt(d.Duration)
	b.PutInt(d.W)
	b.PutInt(d.H)
	return nil
}

// SetRoundMessage sets value of RoundMessage conditional field.
func (d *DocumentAttributeVideo66) SetRoundMessage(value bool) {
	if value {
		d.Flags.Set(0)
		d.RoundMessage = true
	} else {
		d.Flags.Unset(0)
		d.RoundMessage = false
	}
}

// GetRoundMessage returns value of RoundMessage conditional field.
func (d *DocumentAttributeVideo66) GetRoundMessage() (value bool) {
	return d.Flags.Has(0)
}

// GetDuration returns value of Duration field.
func (d *DocumentAttributeVideo66) GetDuration() (value int) {
	return d.Duration
}

// GetW returns value of W field.
func (d *DocumentAttributeVideo66) GetW() (value int) {
	return d.W
}

// GetH returns value of H field.
func (d *DocumentAttributeVideo66) GetH() (value int) {
	return d.H
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeVideo66) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeVideo66#ef02ce6 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeVideo66TypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeVideo66#ef02ce6: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode documentAttributeVideo66#ef02ce6: field flags: %w", err)
		}
	}
	d.RoundMessage = d.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeVideo66#ef02ce6: field duration: %w", err)
		}
		d.Duration = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeVideo66#ef02ce6: field w: %w", err)
		}
		d.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeVideo66#ef02ce6: field h: %w", err)
		}
		d.H = value
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeVideo66) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeVideo66.
var (
	_ bin.Encoder = &DocumentAttributeVideo66{}
	_ bin.Decoder = &DocumentAttributeVideo66{}

	_ DocumentAttributeClass = &DocumentAttributeVideo66{}
)

// DocumentAttributeClass represents DocumentAttribute generic type.
//
// See https://core.telegram.org/type/DocumentAttribute for reference.
//
// Example:
//  g, err := e2e.DecodeDocumentAttribute(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *e2e.DocumentAttributeImageSize: // documentAttributeImageSize#6c37c15c
//  case *e2e.DocumentAttributeAnimated: // documentAttributeAnimated#11b58939
//  case *e2e.DocumentAttributeSticker23: // documentAttributeSticker23#fb0a5727
//  case *e2e.DocumentAttributeVideo: // documentAttributeVideo#5910cccb
//  case *e2e.DocumentAttributeAudio23: // documentAttributeAudio23#51448e5
//  case *e2e.DocumentAttributeFilename: // documentAttributeFilename#15590068
//  case *e2e.DocumentAttributeAudio45: // documentAttributeAudio45#ded218e0
//  case *e2e.DocumentAttributeSticker: // documentAttributeSticker#3a556302
//  case *e2e.DocumentAttributeAudio: // documentAttributeAudio#9852f9c6
//  case *e2e.DocumentAttributeVideo66: // documentAttributeVideo66#ef02ce6
//  default: panic(v)
//  }
type DocumentAttributeClass interface {
	bin.Encoder
	bin.Decoder
	construct() DocumentAttributeClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeDocumentAttribute implements binary de-serialization for DocumentAttributeClass.
func DecodeDocumentAttribute(buf *bin.Buffer) (DocumentAttributeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DocumentAttributeImageSizeTypeID:
		// Decoding documentAttributeImageSize#6c37c15c.
		v := DocumentAttributeImageSize{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeAnimatedTypeID:
		// Decoding documentAttributeAnimated#11b58939.
		v := DocumentAttributeAnimated{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeSticker23TypeID:
		// Decoding documentAttributeSticker23#fb0a5727.
		v := DocumentAttributeSticker23{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeVideoTypeID:
		// Decoding documentAttributeVideo#5910cccb.
		v := DocumentAttributeVideo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeAudio23TypeID:
		// Decoding documentAttributeAudio23#51448e5.
		v := DocumentAttributeAudio23{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeFilenameTypeID:
		// Decoding documentAttributeFilename#15590068.
		v := DocumentAttributeFilename{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeAudio45TypeID:
		// Decoding documentAttributeAudio45#ded218e0.
		v := DocumentAttributeAudio45{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeStickerTypeID:
		// Decoding documentAttributeSticker#3a556302.
		v := DocumentAttributeSticker{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeAudioTypeID:
		// Decoding documentAttributeAudio#9852f9c6.
		v := DocumentAttributeAudio{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeVideo66TypeID:
		// Decoding documentAttributeVideo66#ef02ce6.
		v := DocumentAttributeVideo66{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DocumentAttribute boxes the DocumentAttributeClass providing a helper.
type DocumentAttributeBox struct {
	DocumentAttribute DocumentAttributeClass
}

// Decode implements bin.Decoder for DocumentAttributeBox.
func (b *DocumentAttributeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode DocumentAttributeBox to nil")
	}
	v, err := DecodeDocumentAttribute(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DocumentAttribute = v
	return nil
}

// Encode implements bin.Encode for DocumentAttributeBox.
func (b *DocumentAttributeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DocumentAttribute == nil {
		return fmt.Errorf("unable to encode DocumentAttributeClass as nil")
	}
	return b.DocumentAttribute.Encode(buf)
}

// DocumentAttributeClassSlice is adapter for slice of DocumentAttributeClass.
type DocumentAttributeClassSlice []DocumentAttributeClass

// First returns first element of slice (if exists).
func (s DocumentAttributeClassSlice) First() (v DocumentAttributeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DocumentAttributeClassSlice) Last() (v DocumentAttributeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DocumentAttributeClassSlice) PopFirst() (v DocumentAttributeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	a[len(a)-1] = nil
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DocumentAttributeClassSlice) Pop() (v DocumentAttributeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
