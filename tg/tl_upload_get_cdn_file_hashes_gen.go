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

// UploadGetCdnFileHashesRequest represents TL type `upload.getCdnFileHashes#4da54231`.
//
// See https://core.telegram.org/constructor/upload.getCdnFileHashes for reference.
type UploadGetCdnFileHashesRequest struct {
	// FileToken field of UploadGetCdnFileHashesRequest.
	FileToken []byte
	// Offset field of UploadGetCdnFileHashesRequest.
	Offset int
}

// UploadGetCdnFileHashesRequestTypeID is TL type id of UploadGetCdnFileHashesRequest.
const UploadGetCdnFileHashesRequestTypeID = 0x4da54231

// Encode implements bin.Encoder.
func (g *UploadGetCdnFileHashesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getCdnFileHashes#4da54231 as nil")
	}
	b.PutID(UploadGetCdnFileHashesRequestTypeID)
	b.PutBytes(g.FileToken)
	b.PutInt(g.Offset)
	return nil
}

// Decode implements bin.Decoder.
func (g *UploadGetCdnFileHashesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getCdnFileHashes#4da54231 to nil")
	}
	if err := b.ConsumeID(UploadGetCdnFileHashesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.getCdnFileHashes#4da54231: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getCdnFileHashes#4da54231: field file_token: %w", err)
		}
		g.FileToken = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getCdnFileHashes#4da54231: field offset: %w", err)
		}
		g.Offset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UploadGetCdnFileHashesRequest.
var (
	_ bin.Encoder = &UploadGetCdnFileHashesRequest{}
	_ bin.Decoder = &UploadGetCdnFileHashesRequest{}
)
