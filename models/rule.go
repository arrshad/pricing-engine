package models

import (
	"strconv"
	"strings"
)

// Encode returns an encoded string from Rule fields.
func (r *Rule) Encode() (string, string) {

	var key strings.Builder

	key.WriteString("rule:")
	key.WriteString(strconv.Itoa(r.ID))
	key.WriteByte(':')

	for i := range r.RuleRoutes {
		writeInt(&key, r.RuleRoutes[i].OriginID)
		key.WriteByte('-')
		writeInt(&key, r.RuleRoutes[i].DestinationID)
		key.WriteByte(',')
	}
	key.WriteByte(':')

	for i := range r.RuleAirlines {
		writeInt(&key, r.RuleAirlines[i].AirlineID)
		key.WriteByte(',')
	}
	key.WriteByte(':')

	for i := range r.RuleAgencies {
		writeInt(&key, r.RuleAgencies[i].AgencyID)
		key.WriteByte(',')
	}
	key.WriteByte(':')

	for i := range r.RuleSuppliers {
		writeInt(&key, r.RuleSuppliers[i].SupplierID)
		key.WriteByte(',')
	}
	key.WriteByte(':')

	var val strings.Builder

	if r.IsPercent {
		val.WriteByte('1')
	} else {
		val.WriteByte('0')
	}
	val.WriteByte(':')
	val.WriteString(strconv.Itoa(r.Amount))

	// Output has the following structure:
	// rule:[id]:[routes]:[airelines]:[agencies]:[suppliers]
	// [IsPercent]:[amountValue]
	return key.String(), val.String()

}

func writeInt(s *strings.Builder, i int) {
	s.WriteString(strconv.Itoa(i))
}
