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

// AuthImportLoginTokenRequest represents TL type `auth.importLoginToken#95ac5ce4`.
//
// See https://core.telegram.org/method/auth.importLoginToken for reference.
type AuthImportLoginTokenRequest struct {
	// Token field of AuthImportLoginTokenRequest.
	Token []byte
}

// AuthImportLoginTokenRequestTypeID is TL type id of AuthImportLoginTokenRequest.
const AuthImportLoginTokenRequestTypeID = 0x95ac5ce4

// Encode implements bin.Encoder.
func (i *AuthImportLoginTokenRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode auth.importLoginToken#95ac5ce4 as nil")
	}
	b.PutID(AuthImportLoginTokenRequestTypeID)
	b.PutBytes(i.Token)
	return nil
}

// Decode implements bin.Decoder.
func (i *AuthImportLoginTokenRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode auth.importLoginToken#95ac5ce4 to nil")
	}
	if err := b.ConsumeID(AuthImportLoginTokenRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.importLoginToken#95ac5ce4: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importLoginToken#95ac5ce4: field token: %w", err)
		}
		i.Token = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthImportLoginTokenRequest.
var (
	_ bin.Encoder = &AuthImportLoginTokenRequest{}
	_ bin.Decoder = &AuthImportLoginTokenRequest{}
)

// AuthImportLoginToken invokes method auth.importLoginToken#95ac5ce4 returning error if any.
//
// See https://core.telegram.org/method/auth.importLoginToken for reference.
func (c *Client) AuthImportLoginToken(ctx context.Context, request *AuthImportLoginTokenRequest) (AuthLoginTokenClass, error) {
	var result AuthLoginTokenBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.LoginToken, nil
}
