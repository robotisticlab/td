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

// StatsMegagroupStats represents TL type `stats.megagroupStats#ef7ff916`.
//
// See https://core.telegram.org/constructor/stats.megagroupStats for reference.
type StatsMegagroupStats struct {
	// Period field of StatsMegagroupStats.
	Period StatsDateRangeDays
	// Members field of StatsMegagroupStats.
	Members StatsAbsValueAndPrev
	// Messages field of StatsMegagroupStats.
	Messages StatsAbsValueAndPrev
	// Viewers field of StatsMegagroupStats.
	Viewers StatsAbsValueAndPrev
	// Posters field of StatsMegagroupStats.
	Posters StatsAbsValueAndPrev
	// GrowthGraph field of StatsMegagroupStats.
	GrowthGraph StatsGraphClass
	// MembersGraph field of StatsMegagroupStats.
	MembersGraph StatsGraphClass
	// NewMembersBySourceGraph field of StatsMegagroupStats.
	NewMembersBySourceGraph StatsGraphClass
	// LanguagesGraph field of StatsMegagroupStats.
	LanguagesGraph StatsGraphClass
	// MessagesGraph field of StatsMegagroupStats.
	MessagesGraph StatsGraphClass
	// ActionsGraph field of StatsMegagroupStats.
	ActionsGraph StatsGraphClass
	// TopHoursGraph field of StatsMegagroupStats.
	TopHoursGraph StatsGraphClass
	// WeekdaysGraph field of StatsMegagroupStats.
	WeekdaysGraph StatsGraphClass
	// TopPosters field of StatsMegagroupStats.
	TopPosters []StatsGroupTopPoster
	// TopAdmins field of StatsMegagroupStats.
	TopAdmins []StatsGroupTopAdmin
	// TopInviters field of StatsMegagroupStats.
	TopInviters []StatsGroupTopInviter
	// Users field of StatsMegagroupStats.
	Users []UserClass
}

// StatsMegagroupStatsTypeID is TL type id of StatsMegagroupStats.
const StatsMegagroupStatsTypeID = 0xef7ff916

// Encode implements bin.Encoder.
func (m *StatsMegagroupStats) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode stats.megagroupStats#ef7ff916 as nil")
	}
	b.PutID(StatsMegagroupStatsTypeID)
	if err := m.Period.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field period: %w", err)
	}
	if err := m.Members.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field members: %w", err)
	}
	if err := m.Messages.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field messages: %w", err)
	}
	if err := m.Viewers.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field viewers: %w", err)
	}
	if err := m.Posters.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field posters: %w", err)
	}
	if m.GrowthGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field growth_graph is nil")
	}
	if err := m.GrowthGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field growth_graph: %w", err)
	}
	if m.MembersGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field members_graph is nil")
	}
	if err := m.MembersGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field members_graph: %w", err)
	}
	if m.NewMembersBySourceGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field new_members_by_source_graph is nil")
	}
	if err := m.NewMembersBySourceGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field new_members_by_source_graph: %w", err)
	}
	if m.LanguagesGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field languages_graph is nil")
	}
	if err := m.LanguagesGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field languages_graph: %w", err)
	}
	if m.MessagesGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field messages_graph is nil")
	}
	if err := m.MessagesGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field messages_graph: %w", err)
	}
	if m.ActionsGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field actions_graph is nil")
	}
	if err := m.ActionsGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field actions_graph: %w", err)
	}
	if m.TopHoursGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_hours_graph is nil")
	}
	if err := m.TopHoursGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_hours_graph: %w", err)
	}
	if m.WeekdaysGraph == nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field weekdays_graph is nil")
	}
	if err := m.WeekdaysGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field weekdays_graph: %w", err)
	}
	b.PutVectorHeader(len(m.TopPosters))
	for idx, v := range m.TopPosters {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_posters element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.TopAdmins))
	for idx, v := range m.TopAdmins {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_admins element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.TopInviters))
	for idx, v := range m.TopInviters {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field top_inviters element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Users))
	for idx, v := range m.Users {
		if v == nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.megagroupStats#ef7ff916: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *StatsMegagroupStats) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode stats.megagroupStats#ef7ff916 to nil")
	}
	if err := b.ConsumeID(StatsMegagroupStatsTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: %w", err)
	}
	{
		if err := m.Period.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field period: %w", err)
		}
	}
	{
		if err := m.Members.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field members: %w", err)
		}
	}
	{
		if err := m.Messages.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field messages: %w", err)
		}
	}
	{
		if err := m.Viewers.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field viewers: %w", err)
		}
	}
	{
		if err := m.Posters.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field posters: %w", err)
		}
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field growth_graph: %w", err)
		}
		m.GrowthGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field members_graph: %w", err)
		}
		m.MembersGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field new_members_by_source_graph: %w", err)
		}
		m.NewMembersBySourceGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field languages_graph: %w", err)
		}
		m.LanguagesGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field messages_graph: %w", err)
		}
		m.MessagesGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field actions_graph: %w", err)
		}
		m.ActionsGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_hours_graph: %w", err)
		}
		m.TopHoursGraph = value
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field weekdays_graph: %w", err)
		}
		m.WeekdaysGraph = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_posters: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StatsGroupTopPoster
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_posters: %w", err)
			}
			m.TopPosters = append(m.TopPosters, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_admins: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StatsGroupTopAdmin
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_admins: %w", err)
			}
			m.TopAdmins = append(m.TopAdmins, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_inviters: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StatsGroupTopInviter
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field top_inviters: %w", err)
			}
			m.TopInviters = append(m.TopInviters, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode stats.megagroupStats#ef7ff916: field users: %w", err)
			}
			m.Users = append(m.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsMegagroupStats.
var (
	_ bin.Encoder = &StatsMegagroupStats{}
	_ bin.Decoder = &StatsMegagroupStats{}
)
