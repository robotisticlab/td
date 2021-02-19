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

// DocumentEmpty represents TL type `documentEmpty#36f8c871`.
// Empty constructor, document doesn't exist.
//
// See https://core.telegram.org/constructor/documentEmpty for reference.
type DocumentEmpty struct {
	// Document ID or 0
	ID int64
}

// DocumentEmptyTypeID is TL type id of DocumentEmpty.
const DocumentEmptyTypeID = 0x36f8c871

func (d *DocumentEmpty) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DocumentEmpty) String() string {
	if d == nil {
		return "DocumentEmpty(nil)"
	}
	type Alias DocumentEmpty
	return fmt.Sprintf("DocumentEmpty%+v", Alias(*d))
}

// FillFrom fills DocumentEmpty from given interface.
func (d *DocumentEmpty) FillFrom(from interface {
	GetID() (value int64)
}) {
	d.ID = from.GetID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DocumentEmpty) TypeID() uint32 {
	return DocumentEmptyTypeID
}

// Encode implements bin.Encoder.
func (d *DocumentEmpty) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentEmpty#36f8c871 as nil")
	}
	b.PutID(DocumentEmptyTypeID)
	b.PutLong(d.ID)
	return nil
}

// GetID returns value of ID field.
func (d *DocumentEmpty) GetID() (value int64) {
	return d.ID
}

// Decode implements bin.Decoder.
func (d *DocumentEmpty) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentEmpty#36f8c871 to nil")
	}
	if err := b.ConsumeID(DocumentEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode documentEmpty#36f8c871: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode documentEmpty#36f8c871: field id: %w", err)
		}
		d.ID = value
	}
	return nil
}

// construct implements constructor of DocumentClass.
func (d DocumentEmpty) construct() DocumentClass { return &d }

// Ensuring interfaces in compile-time for DocumentEmpty.
var (
	_ bin.Encoder = &DocumentEmpty{}
	_ bin.Decoder = &DocumentEmpty{}

	_ DocumentClass = &DocumentEmpty{}
)

// Document represents TL type `document#1e87342b`.
// Document
//
// See https://core.telegram.org/constructor/document for reference.
type Document struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Document ID
	ID int64
	// Check sum, dependant on document ID
	AccessHash int64
	// File reference¹
	//
	// Links:
	//  1) https://core.telegram.org/api/file_reference
	FileReference []byte
	// Creation date
	Date int
	// MIME type
	MimeType string
	// Size
	Size int
	// Thumbnails
	//
	// Use SetThumbs and GetThumbs helpers.
	Thumbs []PhotoSizeClass
	// Video thumbnails
	//
	// Use SetVideoThumbs and GetVideoThumbs helpers.
	VideoThumbs []VideoSize
	// DC ID
	DCID int
	// Attributes
	Attributes []DocumentAttributeClass
}

// DocumentTypeID is TL type id of Document.
const DocumentTypeID = 0x1e87342b

func (d *Document) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.ID == 0) {
		return false
	}
	if !(d.AccessHash == 0) {
		return false
	}
	if !(d.FileReference == nil) {
		return false
	}
	if !(d.Date == 0) {
		return false
	}
	if !(d.MimeType == "") {
		return false
	}
	if !(d.Size == 0) {
		return false
	}
	if !(d.Thumbs == nil) {
		return false
	}
	if !(d.VideoThumbs == nil) {
		return false
	}
	if !(d.DCID == 0) {
		return false
	}
	if !(d.Attributes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *Document) String() string {
	if d == nil {
		return "Document(nil)"
	}
	type Alias Document
	return fmt.Sprintf("Document%+v", Alias(*d))
}

// FillFrom fills Document from given interface.
func (d *Document) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetFileReference() (value []byte)
	GetDate() (value int)
	GetMimeType() (value string)
	GetSize() (value int)
	GetThumbs() (value []PhotoSizeClass, ok bool)
	GetVideoThumbs() (value []VideoSize, ok bool)
	GetDCID() (value int)
	GetAttributes() (value []DocumentAttributeClass)
}) {
	d.ID = from.GetID()
	d.AccessHash = from.GetAccessHash()
	d.FileReference = from.GetFileReference()
	d.Date = from.GetDate()
	d.MimeType = from.GetMimeType()
	d.Size = from.GetSize()
	if val, ok := from.GetThumbs(); ok {
		d.Thumbs = val
	}
	if val, ok := from.GetVideoThumbs(); ok {
		d.VideoThumbs = val
	}
	d.DCID = from.GetDCID()
	d.Attributes = from.GetAttributes()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *Document) TypeID() uint32 {
	return DocumentTypeID
}

// Encode implements bin.Encoder.
func (d *Document) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode document#1e87342b as nil")
	}
	b.PutID(DocumentTypeID)
	if !(d.Thumbs == nil) {
		d.Flags.Set(0)
	}
	if !(d.VideoThumbs == nil) {
		d.Flags.Set(1)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode document#1e87342b: field flags: %w", err)
	}
	b.PutLong(d.ID)
	b.PutLong(d.AccessHash)
	b.PutBytes(d.FileReference)
	b.PutInt(d.Date)
	b.PutString(d.MimeType)
	b.PutInt(d.Size)
	if d.Flags.Has(0) {
		b.PutVectorHeader(len(d.Thumbs))
		for idx, v := range d.Thumbs {
			if v == nil {
				return fmt.Errorf("unable to encode document#1e87342b: field thumbs element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode document#1e87342b: field thumbs element with index %d: %w", idx, err)
			}
		}
	}
	if d.Flags.Has(1) {
		b.PutVectorHeader(len(d.VideoThumbs))
		for idx, v := range d.VideoThumbs {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode document#1e87342b: field video_thumbs element with index %d: %w", idx, err)
			}
		}
	}
	b.PutInt(d.DCID)
	b.PutVectorHeader(len(d.Attributes))
	for idx, v := range d.Attributes {
		if v == nil {
			return fmt.Errorf("unable to encode document#1e87342b: field attributes element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode document#1e87342b: field attributes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (d *Document) GetID() (value int64) {
	return d.ID
}

// GetAccessHash returns value of AccessHash field.
func (d *Document) GetAccessHash() (value int64) {
	return d.AccessHash
}

// GetFileReference returns value of FileReference field.
func (d *Document) GetFileReference() (value []byte) {
	return d.FileReference
}

// GetDate returns value of Date field.
func (d *Document) GetDate() (value int) {
	return d.Date
}

// GetMimeType returns value of MimeType field.
func (d *Document) GetMimeType() (value string) {
	return d.MimeType
}

// GetSize returns value of Size field.
func (d *Document) GetSize() (value int) {
	return d.Size
}

// SetThumbs sets value of Thumbs conditional field.
func (d *Document) SetThumbs(value []PhotoSizeClass) {
	d.Flags.Set(0)
	d.Thumbs = value
}

// GetThumbs returns value of Thumbs conditional field and
// boolean which is true if field was set.
func (d *Document) GetThumbs() (value []PhotoSizeClass, ok bool) {
	if !d.Flags.Has(0) {
		return value, false
	}
	return d.Thumbs, true
}

// MapThumbs returns field Thumbs wrapped in PhotoSizeClassSlice helper.
func (d *Document) MapThumbs() (value PhotoSizeClassSlice, ok bool) {
	if !d.Flags.Has(0) {
		return value, false
	}
	return PhotoSizeClassSlice(d.Thumbs), true
}

// SetVideoThumbs sets value of VideoThumbs conditional field.
func (d *Document) SetVideoThumbs(value []VideoSize) {
	d.Flags.Set(1)
	d.VideoThumbs = value
}

// GetVideoThumbs returns value of VideoThumbs conditional field and
// boolean which is true if field was set.
func (d *Document) GetVideoThumbs() (value []VideoSize, ok bool) {
	if !d.Flags.Has(1) {
		return value, false
	}
	return d.VideoThumbs, true
}

// GetDCID returns value of DCID field.
func (d *Document) GetDCID() (value int) {
	return d.DCID
}

// GetAttributes returns value of Attributes field.
func (d *Document) GetAttributes() (value []DocumentAttributeClass) {
	return d.Attributes
}

// MapAttributes returns field Attributes wrapped in DocumentAttributeClassSlice helper.
func (d *Document) MapAttributes() (value DocumentAttributeClassSlice) {
	return DocumentAttributeClassSlice(d.Attributes)
}

// Decode implements bin.Decoder.
func (d *Document) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode document#1e87342b to nil")
	}
	if err := b.ConsumeID(DocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode document#1e87342b: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field flags: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field id: %w", err)
		}
		d.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field access_hash: %w", err)
		}
		d.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field file_reference: %w", err)
		}
		d.FileReference = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field date: %w", err)
		}
		d.Date = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field mime_type: %w", err)
		}
		d.MimeType = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field size: %w", err)
		}
		d.Size = value
	}
	if d.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field thumbs: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePhotoSize(b)
			if err != nil {
				return fmt.Errorf("unable to decode document#1e87342b: field thumbs: %w", err)
			}
			d.Thumbs = append(d.Thumbs, value)
		}
	}
	if d.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field video_thumbs: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value VideoSize
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode document#1e87342b: field video_thumbs: %w", err)
			}
			d.VideoThumbs = append(d.VideoThumbs, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field dc_id: %w", err)
		}
		d.DCID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field attributes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocumentAttribute(b)
			if err != nil {
				return fmt.Errorf("unable to decode document#1e87342b: field attributes: %w", err)
			}
			d.Attributes = append(d.Attributes, value)
		}
	}
	return nil
}

// construct implements constructor of DocumentClass.
func (d Document) construct() DocumentClass { return &d }

// Ensuring interfaces in compile-time for Document.
var (
	_ bin.Encoder = &Document{}
	_ bin.Decoder = &Document{}

	_ DocumentClass = &Document{}
)

// DocumentClass represents Document generic type.
//
// See https://core.telegram.org/type/Document for reference.
//
// Example:
//  g, err := tg.DecodeDocument(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.DocumentEmpty: // documentEmpty#36f8c871
//  case *tg.Document: // document#1e87342b
//  default: panic(v)
//  }
type DocumentClass interface {
	bin.Encoder
	bin.Decoder
	construct() DocumentClass

	// Document ID or 0
	GetID() (value int64)

	// AsNotEmpty tries to map DocumentClass to Document.
	AsNotEmpty() (*Document, bool)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// AsNotEmpty tries to map DocumentClass to Document.
func (d *DocumentEmpty) AsNotEmpty() (*Document, bool) {
	return nil, false
}

// AsNotEmpty tries to map DocumentClass to Document.
func (d *Document) AsNotEmpty() (*Document, bool) {
	return d, true
}

// DecodeDocument implements binary de-serialization for DocumentClass.
func DecodeDocument(buf *bin.Buffer) (DocumentClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DocumentEmptyTypeID:
		// Decoding documentEmpty#36f8c871.
		v := DocumentEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentClass: %w", err)
		}
		return &v, nil
	case DocumentTypeID:
		// Decoding document#1e87342b.
		v := Document{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DocumentClass: %w", bin.NewUnexpectedID(id))
	}
}

// Document boxes the DocumentClass providing a helper.
type DocumentBox struct {
	Document DocumentClass
}

// Decode implements bin.Decoder for DocumentBox.
func (b *DocumentBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode DocumentBox to nil")
	}
	v, err := DecodeDocument(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Document = v
	return nil
}

// Encode implements bin.Encode for DocumentBox.
func (b *DocumentBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Document == nil {
		return fmt.Errorf("unable to encode DocumentClass as nil")
	}
	return b.Document.Encode(buf)
}

// DocumentClassSlice is adapter for slice of DocumentClass.
type DocumentClassSlice []DocumentClass

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s DocumentClassSlice) AppendOnlyNotEmpty(to []*Document) []*Document {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s DocumentClassSlice) AsNotEmpty() (to []*Document) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s DocumentClassSlice) FirstAsNotEmpty() (v *Document, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s DocumentClassSlice) LastAsNotEmpty() (v *Document, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// First returns first element of slice (if exists).
func (s DocumentClassSlice) First() (v DocumentClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DocumentClassSlice) Last() (v DocumentClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DocumentClassSlice) PopFirst() (v DocumentClass, ok bool) {
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
func (s *DocumentClassSlice) Pop() (v DocumentClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
