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

// TextEmpty represents TL type `textEmpty#dc3d824f`.
//
// See https://core.telegram.org/constructor/textEmpty for reference.
type TextEmpty struct {
}

// TextEmptyTypeID is TL type id of TextEmpty.
const TextEmptyTypeID = 0xdc3d824f

// Encode implements bin.Encoder.
func (t *TextEmpty) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textEmpty#dc3d824f as nil")
	}
	b.PutID(TextEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TextEmpty) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textEmpty#dc3d824f to nil")
	}
	if err := b.ConsumeID(TextEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode textEmpty#dc3d824f: %w", err)
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextEmpty) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextEmpty.
var (
	_ bin.Encoder = &TextEmpty{}
	_ bin.Decoder = &TextEmpty{}

	_ RichTextClass = &TextEmpty{}
)

// TextPlain represents TL type `textPlain#744694e0`.
//
// See https://core.telegram.org/constructor/textPlain for reference.
type TextPlain struct {
	// Text field of TextPlain.
	Text string
}

// TextPlainTypeID is TL type id of TextPlain.
const TextPlainTypeID = 0x744694e0

// Encode implements bin.Encoder.
func (t *TextPlain) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textPlain#744694e0 as nil")
	}
	b.PutID(TextPlainTypeID)
	b.PutString(t.Text)
	return nil
}

// Decode implements bin.Decoder.
func (t *TextPlain) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textPlain#744694e0 to nil")
	}
	if err := b.ConsumeID(TextPlainTypeID); err != nil {
		return fmt.Errorf("unable to decode textPlain#744694e0: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode textPlain#744694e0: field text: %w", err)
		}
		t.Text = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextPlain) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextPlain.
var (
	_ bin.Encoder = &TextPlain{}
	_ bin.Decoder = &TextPlain{}

	_ RichTextClass = &TextPlain{}
)

// TextBold represents TL type `textBold#6724abc4`.
//
// See https://core.telegram.org/constructor/textBold for reference.
type TextBold struct {
	// Text field of TextBold.
	Text RichTextClass
}

// TextBoldTypeID is TL type id of TextBold.
const TextBoldTypeID = 0x6724abc4

// Encode implements bin.Encoder.
func (t *TextBold) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textBold#6724abc4 as nil")
	}
	b.PutID(TextBoldTypeID)
	if t.Text == nil {
		return fmt.Errorf("unable to encode textBold#6724abc4: field text is nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textBold#6724abc4: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TextBold) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textBold#6724abc4 to nil")
	}
	if err := b.ConsumeID(TextBoldTypeID); err != nil {
		return fmt.Errorf("unable to decode textBold#6724abc4: %w", err)
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode textBold#6724abc4: field text: %w", err)
		}
		t.Text = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextBold) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextBold.
var (
	_ bin.Encoder = &TextBold{}
	_ bin.Decoder = &TextBold{}

	_ RichTextClass = &TextBold{}
)

// TextItalic represents TL type `textItalic#d912a59c`.
//
// See https://core.telegram.org/constructor/textItalic for reference.
type TextItalic struct {
	// Text field of TextItalic.
	Text RichTextClass
}

// TextItalicTypeID is TL type id of TextItalic.
const TextItalicTypeID = 0xd912a59c

// Encode implements bin.Encoder.
func (t *TextItalic) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textItalic#d912a59c as nil")
	}
	b.PutID(TextItalicTypeID)
	if t.Text == nil {
		return fmt.Errorf("unable to encode textItalic#d912a59c: field text is nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textItalic#d912a59c: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TextItalic) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textItalic#d912a59c to nil")
	}
	if err := b.ConsumeID(TextItalicTypeID); err != nil {
		return fmt.Errorf("unable to decode textItalic#d912a59c: %w", err)
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode textItalic#d912a59c: field text: %w", err)
		}
		t.Text = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextItalic) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextItalic.
var (
	_ bin.Encoder = &TextItalic{}
	_ bin.Decoder = &TextItalic{}

	_ RichTextClass = &TextItalic{}
)

// TextUnderline represents TL type `textUnderline#c12622c4`.
//
// See https://core.telegram.org/constructor/textUnderline for reference.
type TextUnderline struct {
	// Text field of TextUnderline.
	Text RichTextClass
}

// TextUnderlineTypeID is TL type id of TextUnderline.
const TextUnderlineTypeID = 0xc12622c4

// Encode implements bin.Encoder.
func (t *TextUnderline) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textUnderline#c12622c4 as nil")
	}
	b.PutID(TextUnderlineTypeID)
	if t.Text == nil {
		return fmt.Errorf("unable to encode textUnderline#c12622c4: field text is nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textUnderline#c12622c4: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TextUnderline) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textUnderline#c12622c4 to nil")
	}
	if err := b.ConsumeID(TextUnderlineTypeID); err != nil {
		return fmt.Errorf("unable to decode textUnderline#c12622c4: %w", err)
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode textUnderline#c12622c4: field text: %w", err)
		}
		t.Text = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextUnderline) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextUnderline.
var (
	_ bin.Encoder = &TextUnderline{}
	_ bin.Decoder = &TextUnderline{}

	_ RichTextClass = &TextUnderline{}
)

// TextStrike represents TL type `textStrike#9bf8bb95`.
//
// See https://core.telegram.org/constructor/textStrike for reference.
type TextStrike struct {
	// Text field of TextStrike.
	Text RichTextClass
}

// TextStrikeTypeID is TL type id of TextStrike.
const TextStrikeTypeID = 0x9bf8bb95

// Encode implements bin.Encoder.
func (t *TextStrike) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textStrike#9bf8bb95 as nil")
	}
	b.PutID(TextStrikeTypeID)
	if t.Text == nil {
		return fmt.Errorf("unable to encode textStrike#9bf8bb95: field text is nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textStrike#9bf8bb95: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TextStrike) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textStrike#9bf8bb95 to nil")
	}
	if err := b.ConsumeID(TextStrikeTypeID); err != nil {
		return fmt.Errorf("unable to decode textStrike#9bf8bb95: %w", err)
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode textStrike#9bf8bb95: field text: %w", err)
		}
		t.Text = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextStrike) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextStrike.
var (
	_ bin.Encoder = &TextStrike{}
	_ bin.Decoder = &TextStrike{}

	_ RichTextClass = &TextStrike{}
)

// TextFixed represents TL type `textFixed#6c3f19b9`.
//
// See https://core.telegram.org/constructor/textFixed for reference.
type TextFixed struct {
	// Text field of TextFixed.
	Text RichTextClass
}

// TextFixedTypeID is TL type id of TextFixed.
const TextFixedTypeID = 0x6c3f19b9

// Encode implements bin.Encoder.
func (t *TextFixed) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textFixed#6c3f19b9 as nil")
	}
	b.PutID(TextFixedTypeID)
	if t.Text == nil {
		return fmt.Errorf("unable to encode textFixed#6c3f19b9: field text is nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textFixed#6c3f19b9: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TextFixed) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textFixed#6c3f19b9 to nil")
	}
	if err := b.ConsumeID(TextFixedTypeID); err != nil {
		return fmt.Errorf("unable to decode textFixed#6c3f19b9: %w", err)
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode textFixed#6c3f19b9: field text: %w", err)
		}
		t.Text = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextFixed) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextFixed.
var (
	_ bin.Encoder = &TextFixed{}
	_ bin.Decoder = &TextFixed{}

	_ RichTextClass = &TextFixed{}
)

// TextUrl represents TL type `textUrl#3c2884c1`.
//
// See https://core.telegram.org/constructor/textUrl for reference.
type TextUrl struct {
	// Text field of TextUrl.
	Text RichTextClass
	// URL field of TextUrl.
	URL string
	// WebpageID field of TextUrl.
	WebpageID int64
}

// TextUrlTypeID is TL type id of TextUrl.
const TextUrlTypeID = 0x3c2884c1

// Encode implements bin.Encoder.
func (t *TextUrl) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textUrl#3c2884c1 as nil")
	}
	b.PutID(TextUrlTypeID)
	if t.Text == nil {
		return fmt.Errorf("unable to encode textUrl#3c2884c1: field text is nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textUrl#3c2884c1: field text: %w", err)
	}
	b.PutString(t.URL)
	b.PutLong(t.WebpageID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TextUrl) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textUrl#3c2884c1 to nil")
	}
	if err := b.ConsumeID(TextUrlTypeID); err != nil {
		return fmt.Errorf("unable to decode textUrl#3c2884c1: %w", err)
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode textUrl#3c2884c1: field text: %w", err)
		}
		t.Text = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode textUrl#3c2884c1: field url: %w", err)
		}
		t.URL = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode textUrl#3c2884c1: field webpage_id: %w", err)
		}
		t.WebpageID = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextUrl) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextUrl.
var (
	_ bin.Encoder = &TextUrl{}
	_ bin.Decoder = &TextUrl{}

	_ RichTextClass = &TextUrl{}
)

// TextEmail represents TL type `textEmail#de5a0dd6`.
//
// See https://core.telegram.org/constructor/textEmail for reference.
type TextEmail struct {
	// Text field of TextEmail.
	Text RichTextClass
	// Email field of TextEmail.
	Email string
}

// TextEmailTypeID is TL type id of TextEmail.
const TextEmailTypeID = 0xde5a0dd6

// Encode implements bin.Encoder.
func (t *TextEmail) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textEmail#de5a0dd6 as nil")
	}
	b.PutID(TextEmailTypeID)
	if t.Text == nil {
		return fmt.Errorf("unable to encode textEmail#de5a0dd6: field text is nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textEmail#de5a0dd6: field text: %w", err)
	}
	b.PutString(t.Email)
	return nil
}

// Decode implements bin.Decoder.
func (t *TextEmail) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textEmail#de5a0dd6 to nil")
	}
	if err := b.ConsumeID(TextEmailTypeID); err != nil {
		return fmt.Errorf("unable to decode textEmail#de5a0dd6: %w", err)
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode textEmail#de5a0dd6: field text: %w", err)
		}
		t.Text = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode textEmail#de5a0dd6: field email: %w", err)
		}
		t.Email = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextEmail) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextEmail.
var (
	_ bin.Encoder = &TextEmail{}
	_ bin.Decoder = &TextEmail{}

	_ RichTextClass = &TextEmail{}
)

// TextConcat represents TL type `textConcat#7e6260d7`.
//
// See https://core.telegram.org/constructor/textConcat for reference.
type TextConcat struct {
	// Texts field of TextConcat.
	Texts []RichTextClass
}

// TextConcatTypeID is TL type id of TextConcat.
const TextConcatTypeID = 0x7e6260d7

// Encode implements bin.Encoder.
func (t *TextConcat) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textConcat#7e6260d7 as nil")
	}
	b.PutID(TextConcatTypeID)
	b.PutVectorHeader(len(t.Texts))
	for idx, v := range t.Texts {
		if v == nil {
			return fmt.Errorf("unable to encode textConcat#7e6260d7: field texts element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode textConcat#7e6260d7: field texts element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TextConcat) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textConcat#7e6260d7 to nil")
	}
	if err := b.ConsumeID(TextConcatTypeID); err != nil {
		return fmt.Errorf("unable to decode textConcat#7e6260d7: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode textConcat#7e6260d7: field texts: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeRichText(b)
			if err != nil {
				return fmt.Errorf("unable to decode textConcat#7e6260d7: field texts: %w", err)
			}
			t.Texts = append(t.Texts, value)
		}
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextConcat) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextConcat.
var (
	_ bin.Encoder = &TextConcat{}
	_ bin.Decoder = &TextConcat{}

	_ RichTextClass = &TextConcat{}
)

// TextSubscript represents TL type `textSubscript#ed6a8504`.
//
// See https://core.telegram.org/constructor/textSubscript for reference.
type TextSubscript struct {
	// Text field of TextSubscript.
	Text RichTextClass
}

// TextSubscriptTypeID is TL type id of TextSubscript.
const TextSubscriptTypeID = 0xed6a8504

// Encode implements bin.Encoder.
func (t *TextSubscript) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textSubscript#ed6a8504 as nil")
	}
	b.PutID(TextSubscriptTypeID)
	if t.Text == nil {
		return fmt.Errorf("unable to encode textSubscript#ed6a8504: field text is nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textSubscript#ed6a8504: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TextSubscript) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textSubscript#ed6a8504 to nil")
	}
	if err := b.ConsumeID(TextSubscriptTypeID); err != nil {
		return fmt.Errorf("unable to decode textSubscript#ed6a8504: %w", err)
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode textSubscript#ed6a8504: field text: %w", err)
		}
		t.Text = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextSubscript) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextSubscript.
var (
	_ bin.Encoder = &TextSubscript{}
	_ bin.Decoder = &TextSubscript{}

	_ RichTextClass = &TextSubscript{}
)

// TextSuperscript represents TL type `textSuperscript#c7fb5e01`.
//
// See https://core.telegram.org/constructor/textSuperscript for reference.
type TextSuperscript struct {
	// Text field of TextSuperscript.
	Text RichTextClass
}

// TextSuperscriptTypeID is TL type id of TextSuperscript.
const TextSuperscriptTypeID = 0xc7fb5e01

// Encode implements bin.Encoder.
func (t *TextSuperscript) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textSuperscript#c7fb5e01 as nil")
	}
	b.PutID(TextSuperscriptTypeID)
	if t.Text == nil {
		return fmt.Errorf("unable to encode textSuperscript#c7fb5e01: field text is nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textSuperscript#c7fb5e01: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TextSuperscript) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textSuperscript#c7fb5e01 to nil")
	}
	if err := b.ConsumeID(TextSuperscriptTypeID); err != nil {
		return fmt.Errorf("unable to decode textSuperscript#c7fb5e01: %w", err)
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode textSuperscript#c7fb5e01: field text: %w", err)
		}
		t.Text = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextSuperscript) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextSuperscript.
var (
	_ bin.Encoder = &TextSuperscript{}
	_ bin.Decoder = &TextSuperscript{}

	_ RichTextClass = &TextSuperscript{}
)

// TextMarked represents TL type `textMarked#34b8621`.
//
// See https://core.telegram.org/constructor/textMarked for reference.
type TextMarked struct {
	// Text field of TextMarked.
	Text RichTextClass
}

// TextMarkedTypeID is TL type id of TextMarked.
const TextMarkedTypeID = 0x34b8621

// Encode implements bin.Encoder.
func (t *TextMarked) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textMarked#34b8621 as nil")
	}
	b.PutID(TextMarkedTypeID)
	if t.Text == nil {
		return fmt.Errorf("unable to encode textMarked#34b8621: field text is nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textMarked#34b8621: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TextMarked) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textMarked#34b8621 to nil")
	}
	if err := b.ConsumeID(TextMarkedTypeID); err != nil {
		return fmt.Errorf("unable to decode textMarked#34b8621: %w", err)
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode textMarked#34b8621: field text: %w", err)
		}
		t.Text = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextMarked) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextMarked.
var (
	_ bin.Encoder = &TextMarked{}
	_ bin.Decoder = &TextMarked{}

	_ RichTextClass = &TextMarked{}
)

// TextPhone represents TL type `textPhone#1ccb966a`.
//
// See https://core.telegram.org/constructor/textPhone for reference.
type TextPhone struct {
	// Text field of TextPhone.
	Text RichTextClass
	// Phone field of TextPhone.
	Phone string
}

// TextPhoneTypeID is TL type id of TextPhone.
const TextPhoneTypeID = 0x1ccb966a

// Encode implements bin.Encoder.
func (t *TextPhone) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textPhone#1ccb966a as nil")
	}
	b.PutID(TextPhoneTypeID)
	if t.Text == nil {
		return fmt.Errorf("unable to encode textPhone#1ccb966a: field text is nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textPhone#1ccb966a: field text: %w", err)
	}
	b.PutString(t.Phone)
	return nil
}

// Decode implements bin.Decoder.
func (t *TextPhone) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textPhone#1ccb966a to nil")
	}
	if err := b.ConsumeID(TextPhoneTypeID); err != nil {
		return fmt.Errorf("unable to decode textPhone#1ccb966a: %w", err)
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode textPhone#1ccb966a: field text: %w", err)
		}
		t.Text = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode textPhone#1ccb966a: field phone: %w", err)
		}
		t.Phone = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextPhone) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextPhone.
var (
	_ bin.Encoder = &TextPhone{}
	_ bin.Decoder = &TextPhone{}

	_ RichTextClass = &TextPhone{}
)

// TextImage represents TL type `textImage#81ccf4f`.
//
// See https://core.telegram.org/constructor/textImage for reference.
type TextImage struct {
	// DocumentID field of TextImage.
	DocumentID int64
	// W field of TextImage.
	W int
	// H field of TextImage.
	H int
}

// TextImageTypeID is TL type id of TextImage.
const TextImageTypeID = 0x81ccf4f

// Encode implements bin.Encoder.
func (t *TextImage) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textImage#81ccf4f as nil")
	}
	b.PutID(TextImageTypeID)
	b.PutLong(t.DocumentID)
	b.PutInt(t.W)
	b.PutInt(t.H)
	return nil
}

// Decode implements bin.Decoder.
func (t *TextImage) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textImage#81ccf4f to nil")
	}
	if err := b.ConsumeID(TextImageTypeID); err != nil {
		return fmt.Errorf("unable to decode textImage#81ccf4f: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode textImage#81ccf4f: field document_id: %w", err)
		}
		t.DocumentID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode textImage#81ccf4f: field w: %w", err)
		}
		t.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode textImage#81ccf4f: field h: %w", err)
		}
		t.H = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextImage) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextImage.
var (
	_ bin.Encoder = &TextImage{}
	_ bin.Decoder = &TextImage{}

	_ RichTextClass = &TextImage{}
)

// TextAnchor represents TL type `textAnchor#35553762`.
//
// See https://core.telegram.org/constructor/textAnchor for reference.
type TextAnchor struct {
	// Text field of TextAnchor.
	Text RichTextClass
	// Name field of TextAnchor.
	Name string
}

// TextAnchorTypeID is TL type id of TextAnchor.
const TextAnchorTypeID = 0x35553762

// Encode implements bin.Encoder.
func (t *TextAnchor) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textAnchor#35553762 as nil")
	}
	b.PutID(TextAnchorTypeID)
	if t.Text == nil {
		return fmt.Errorf("unable to encode textAnchor#35553762: field text is nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textAnchor#35553762: field text: %w", err)
	}
	b.PutString(t.Name)
	return nil
}

// Decode implements bin.Decoder.
func (t *TextAnchor) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textAnchor#35553762 to nil")
	}
	if err := b.ConsumeID(TextAnchorTypeID); err != nil {
		return fmt.Errorf("unable to decode textAnchor#35553762: %w", err)
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode textAnchor#35553762: field text: %w", err)
		}
		t.Text = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode textAnchor#35553762: field name: %w", err)
		}
		t.Name = value
	}
	return nil
}

// construct implements constructor of RichTextClass.
func (t TextAnchor) construct() RichTextClass { return &t }

// Ensuring interfaces in compile-time for TextAnchor.
var (
	_ bin.Encoder = &TextAnchor{}
	_ bin.Decoder = &TextAnchor{}

	_ RichTextClass = &TextAnchor{}
)

// RichTextClass represents RichText generic type.
//
// See https://core.telegram.org/type/RichText for reference.
//
// Example:
//  g, err := DecodeRichText(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *TextEmpty: // textEmpty#dc3d824f
//  case *TextPlain: // textPlain#744694e0
//  case *TextBold: // textBold#6724abc4
//  case *TextItalic: // textItalic#d912a59c
//  case *TextUnderline: // textUnderline#c12622c4
//  case *TextStrike: // textStrike#9bf8bb95
//  case *TextFixed: // textFixed#6c3f19b9
//  case *TextUrl: // textUrl#3c2884c1
//  case *TextEmail: // textEmail#de5a0dd6
//  case *TextConcat: // textConcat#7e6260d7
//  case *TextSubscript: // textSubscript#ed6a8504
//  case *TextSuperscript: // textSuperscript#c7fb5e01
//  case *TextMarked: // textMarked#34b8621
//  case *TextPhone: // textPhone#1ccb966a
//  case *TextImage: // textImage#81ccf4f
//  case *TextAnchor: // textAnchor#35553762
//  default: panic(v)
//  }
type RichTextClass interface {
	bin.Encoder
	bin.Decoder
	construct() RichTextClass
}

// DecodeRichText implements binary de-serialization for RichTextClass.
func DecodeRichText(buf *bin.Buffer) (RichTextClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case TextEmptyTypeID:
		// Decoding textEmpty#dc3d824f.
		v := TextEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextPlainTypeID:
		// Decoding textPlain#744694e0.
		v := TextPlain{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextBoldTypeID:
		// Decoding textBold#6724abc4.
		v := TextBold{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextItalicTypeID:
		// Decoding textItalic#d912a59c.
		v := TextItalic{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextUnderlineTypeID:
		// Decoding textUnderline#c12622c4.
		v := TextUnderline{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextStrikeTypeID:
		// Decoding textStrike#9bf8bb95.
		v := TextStrike{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextFixedTypeID:
		// Decoding textFixed#6c3f19b9.
		v := TextFixed{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextUrlTypeID:
		// Decoding textUrl#3c2884c1.
		v := TextUrl{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextEmailTypeID:
		// Decoding textEmail#de5a0dd6.
		v := TextEmail{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextConcatTypeID:
		// Decoding textConcat#7e6260d7.
		v := TextConcat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextSubscriptTypeID:
		// Decoding textSubscript#ed6a8504.
		v := TextSubscript{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextSuperscriptTypeID:
		// Decoding textSuperscript#c7fb5e01.
		v := TextSuperscript{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextMarkedTypeID:
		// Decoding textMarked#34b8621.
		v := TextMarked{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextPhoneTypeID:
		// Decoding textPhone#1ccb966a.
		v := TextPhone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextImageTypeID:
		// Decoding textImage#81ccf4f.
		v := TextImage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	case TextAnchorTypeID:
		// Decoding textAnchor#35553762.
		v := TextAnchor{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RichTextClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode RichTextClass: %w", bin.NewUnexpectedID(id))
	}
}

// RichText boxes the RichTextClass providing a helper.
type RichTextBox struct {
	RichText RichTextClass
}

// Decode implements bin.Decoder for RichTextBox.
func (b *RichTextBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode RichTextBox to nil")
	}
	v, err := DecodeRichText(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.RichText = v
	return nil
}

// Encode implements bin.Encode for RichTextBox.
func (b *RichTextBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.RichText == nil {
		return fmt.Errorf("unable to encode RichTextClass as nil")
	}
	return b.RichText.Encode(buf)
}
