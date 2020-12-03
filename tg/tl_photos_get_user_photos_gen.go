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

// PhotosGetUserPhotosRequest represents TL type `photos.getUserPhotos#91cd32a8`.
//
// See https://core.telegram.org/method/photos.getUserPhotos for reference.
type PhotosGetUserPhotosRequest struct {
	// UserID field of PhotosGetUserPhotosRequest.
	UserID InputUserClass
	// Offset field of PhotosGetUserPhotosRequest.
	Offset int
	// MaxID field of PhotosGetUserPhotosRequest.
	MaxID int64
	// Limit field of PhotosGetUserPhotosRequest.
	Limit int
}

// PhotosGetUserPhotosRequestTypeID is TL type id of PhotosGetUserPhotosRequest.
const PhotosGetUserPhotosRequestTypeID = 0x91cd32a8

// Encode implements bin.Encoder.
func (g *PhotosGetUserPhotosRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode photos.getUserPhotos#91cd32a8 as nil")
	}
	b.PutID(PhotosGetUserPhotosRequestTypeID)
	if g.UserID == nil {
		return fmt.Errorf("unable to encode photos.getUserPhotos#91cd32a8: field user_id is nil")
	}
	if err := g.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photos.getUserPhotos#91cd32a8: field user_id: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutLong(g.MaxID)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *PhotosGetUserPhotosRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode photos.getUserPhotos#91cd32a8 to nil")
	}
	if err := b.ConsumeID(PhotosGetUserPhotosRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: %w", err)
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: field user_id: %w", err)
		}
		g.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photos.getUserPhotos#91cd32a8: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhotosGetUserPhotosRequest.
var (
	_ bin.Encoder = &PhotosGetUserPhotosRequest{}
	_ bin.Decoder = &PhotosGetUserPhotosRequest{}
)

// PhotosGetUserPhotos invokes method photos.getUserPhotos#91cd32a8 returning error if any.
//
// See https://core.telegram.org/method/photos.getUserPhotos for reference.
func (c *Client) PhotosGetUserPhotos(ctx context.Context, request *PhotosGetUserPhotosRequest) (PhotosPhotosClass, error) {
	var result PhotosPhotosBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Photos, nil
}
