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

// MessagesGetEmojiKeywordsLanguagesRequest represents TL type `messages.getEmojiKeywordsLanguages#4e9963b2`.
//
// See https://core.telegram.org/constructor/messages.getEmojiKeywordsLanguages for reference.
type MessagesGetEmojiKeywordsLanguagesRequest struct {
	// LangCodes field of MessagesGetEmojiKeywordsLanguagesRequest.
	LangCodes []string
}

// MessagesGetEmojiKeywordsLanguagesRequestTypeID is TL type id of MessagesGetEmojiKeywordsLanguagesRequest.
const MessagesGetEmojiKeywordsLanguagesRequestTypeID = 0x4e9963b2

// Encode implements bin.Encoder.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiKeywordsLanguages#4e9963b2 as nil")
	}
	b.PutID(MessagesGetEmojiKeywordsLanguagesRequestTypeID)
	b.PutVectorHeader(len(g.LangCodes))
	for _, v := range g.LangCodes {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiKeywordsLanguages#4e9963b2 to nil")
	}
	if err := b.ConsumeID(MessagesGetEmojiKeywordsLanguagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getEmojiKeywordsLanguages#4e9963b2: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiKeywordsLanguages#4e9963b2: field lang_codes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode messages.getEmojiKeywordsLanguages#4e9963b2: field lang_codes: %w", err)
			}
			g.LangCodes = append(g.LangCodes, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetEmojiKeywordsLanguagesRequest.
var (
	_ bin.Encoder = &MessagesGetEmojiKeywordsLanguagesRequest{}
	_ bin.Decoder = &MessagesGetEmojiKeywordsLanguagesRequest{}
)
