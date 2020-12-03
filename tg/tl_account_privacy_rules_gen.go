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

// AccountPrivacyRules represents TL type `account.privacyRules#50a04e45`.
//
// See https://core.telegram.org/constructor/account.privacyRules for reference.
type AccountPrivacyRules struct {
	// Rules field of AccountPrivacyRules.
	Rules []PrivacyRuleClass
	// Chats field of AccountPrivacyRules.
	Chats []ChatClass
	// Users field of AccountPrivacyRules.
	Users []UserClass
}

// AccountPrivacyRulesTypeID is TL type id of AccountPrivacyRules.
const AccountPrivacyRulesTypeID = 0x50a04e45

// Encode implements bin.Encoder.
func (p *AccountPrivacyRules) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode account.privacyRules#50a04e45 as nil")
	}
	b.PutID(AccountPrivacyRulesTypeID)
	b.PutVectorHeader(len(p.Rules))
	for idx, v := range p.Rules {
		if v == nil {
			return fmt.Errorf("unable to encode account.privacyRules#50a04e45: field rules element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.privacyRules#50a04e45: field rules element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Chats))
	for idx, v := range p.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode account.privacyRules#50a04e45: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.privacyRules#50a04e45: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode account.privacyRules#50a04e45: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.privacyRules#50a04e45: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *AccountPrivacyRules) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode account.privacyRules#50a04e45 to nil")
	}
	if err := b.ConsumeID(AccountPrivacyRulesTypeID); err != nil {
		return fmt.Errorf("unable to decode account.privacyRules#50a04e45: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.privacyRules#50a04e45: field rules: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePrivacyRule(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.privacyRules#50a04e45: field rules: %w", err)
			}
			p.Rules = append(p.Rules, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.privacyRules#50a04e45: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.privacyRules#50a04e45: field chats: %w", err)
			}
			p.Chats = append(p.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.privacyRules#50a04e45: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.privacyRules#50a04e45: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountPrivacyRules.
var (
	_ bin.Encoder = &AccountPrivacyRules{}
	_ bin.Decoder = &AccountPrivacyRules{}
)
