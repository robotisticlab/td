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

// HelpCountriesListNotModified represents TL type `help.countriesListNotModified#93cc1f32`.
//
// See https://core.telegram.org/constructor/help.countriesListNotModified for reference.
type HelpCountriesListNotModified struct {
}

// HelpCountriesListNotModifiedTypeID is TL type id of HelpCountriesListNotModified.
const HelpCountriesListNotModifiedTypeID = 0x93cc1f32

// Encode implements bin.Encoder.
func (c *HelpCountriesListNotModified) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode help.countriesListNotModified#93cc1f32 as nil")
	}
	b.PutID(HelpCountriesListNotModifiedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *HelpCountriesListNotModified) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode help.countriesListNotModified#93cc1f32 to nil")
	}
	if err := b.ConsumeID(HelpCountriesListNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode help.countriesListNotModified#93cc1f32: %w", err)
	}
	return nil
}

// construct implements constructor of HelpCountriesListClass.
func (c HelpCountriesListNotModified) construct() HelpCountriesListClass { return &c }

// Ensuring interfaces in compile-time for HelpCountriesListNotModified.
var (
	_ bin.Encoder = &HelpCountriesListNotModified{}
	_ bin.Decoder = &HelpCountriesListNotModified{}

	_ HelpCountriesListClass = &HelpCountriesListNotModified{}
)

// HelpCountriesList represents TL type `help.countriesList#87d0759e`.
//
// See https://core.telegram.org/constructor/help.countriesList for reference.
type HelpCountriesList struct {
	// Countries field of HelpCountriesList.
	Countries []HelpCountry
	// Hash field of HelpCountriesList.
	Hash int
}

// HelpCountriesListTypeID is TL type id of HelpCountriesList.
const HelpCountriesListTypeID = 0x87d0759e

// Encode implements bin.Encoder.
func (c *HelpCountriesList) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode help.countriesList#87d0759e as nil")
	}
	b.PutID(HelpCountriesListTypeID)
	b.PutVectorHeader(len(c.Countries))
	for idx, v := range c.Countries {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.countriesList#87d0759e: field countries element with index %d: %w", idx, err)
		}
	}
	b.PutInt(c.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (c *HelpCountriesList) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode help.countriesList#87d0759e to nil")
	}
	if err := b.ConsumeID(HelpCountriesListTypeID); err != nil {
		return fmt.Errorf("unable to decode help.countriesList#87d0759e: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.countriesList#87d0759e: field countries: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value HelpCountry
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode help.countriesList#87d0759e: field countries: %w", err)
			}
			c.Countries = append(c.Countries, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.countriesList#87d0759e: field hash: %w", err)
		}
		c.Hash = value
	}
	return nil
}

// construct implements constructor of HelpCountriesListClass.
func (c HelpCountriesList) construct() HelpCountriesListClass { return &c }

// Ensuring interfaces in compile-time for HelpCountriesList.
var (
	_ bin.Encoder = &HelpCountriesList{}
	_ bin.Decoder = &HelpCountriesList{}

	_ HelpCountriesListClass = &HelpCountriesList{}
)

// HelpCountriesListClass represents help.CountriesList generic type.
//
// See https://core.telegram.org/type/help.CountriesList for reference.
//
// Example:
//  g, err := DecodeHelpCountriesList(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *HelpCountriesListNotModified: // help.countriesListNotModified#93cc1f32
//  case *HelpCountriesList: // help.countriesList#87d0759e
//  default: panic(v)
//  }
type HelpCountriesListClass interface {
	bin.Encoder
	bin.Decoder
	construct() HelpCountriesListClass
}

// DecodeHelpCountriesList implements binary de-serialization for HelpCountriesListClass.
func DecodeHelpCountriesList(buf *bin.Buffer) (HelpCountriesListClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpCountriesListNotModifiedTypeID:
		// Decoding help.countriesListNotModified#93cc1f32.
		v := HelpCountriesListNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpCountriesListClass: %w", err)
		}
		return &v, nil
	case HelpCountriesListTypeID:
		// Decoding help.countriesList#87d0759e.
		v := HelpCountriesList{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpCountriesListClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpCountriesListClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpCountriesList boxes the HelpCountriesListClass providing a helper.
type HelpCountriesListBox struct {
	CountriesList HelpCountriesListClass
}

// Decode implements bin.Decoder for HelpCountriesListBox.
func (b *HelpCountriesListBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpCountriesListBox to nil")
	}
	v, err := DecodeHelpCountriesList(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CountriesList = v
	return nil
}

// Encode implements bin.Encode for HelpCountriesListBox.
func (b *HelpCountriesListBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.CountriesList == nil {
		return fmt.Errorf("unable to encode HelpCountriesListClass as nil")
	}
	return b.CountriesList.Encode(buf)
}
