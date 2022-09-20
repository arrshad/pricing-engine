package models

// import "strings"

// get returns field value with index
func (t *Ticket) Get(i int) string {
	switch i {
	case 2:
		return t.AirlineID
	case 3:
		return t.AgencyID
	case 4:
		return t.SupplierID
	default:
		return ""
	}
}

/*
// IsMatch returns ruleID if the ticket matches the given rule
func (t *Ticket) IsMatch(rule string) (ruleID string) {
	// The reason I wrote such a function myself was
	// because it's up to 4x faster than using split
	// and contains functions

	// Remove the 'rule:' from beginning
	rule = rule[5:]
	// Store id of the current rule
	o := strings.IndexByte(rule, ':')
	ruleID = rule[:o]
	// Index of rule string
	i := o + 1
	// Position of current field
	f := 1

	for i < len(rule) {
		// Store rest of rule string
		rule = rule[i:]
		// Check if this is an empty field
		if rule[0] == ':' {
			// Field is null and it means 'all',
			// so we go to the next field
			i += 1
			f += 1
			continue
		}
		// Get next field index
		n := strings.IndexByte(rule, ':')
		// Routes field is a little different,
		// so I seperated checking method of it
		if f == 1 {
			// Find end of current route (routes separator)
			e := strings.IndexByte(rule, ',')
			// Find origin and destination separator
			m := strings.IndexByte(rule[:e], '-')
			// Check if current route match the ticket
			if o := rule[:m]; o == "0" || o == t.OriginID {
				if d := rule[m+1 : e]; d == "0" ||
					d == t.DestinationID {
					// Route found
					// Going to next field
					i = n + 1
					f += 1
					continue
				}
			}
			// Reaching cursor to ',:' means no route
			// match the ticket, so it returns zero.
			if rule[e+1] == ':' {
				return ""
			}
			// Change position of cursor to next route
			i = e + 1

		} else {
			// Get field value with index
			v := t.Get(f)
			// Find end of current value
			e := strings.IndexByte(rule[:n], ',')
			// Check if current value match the ticket
			if rule[:e] == v {
				// Value found
				// Going to next field
				i = n + 1
				f += 1
				continue
			}
			if rule[e+1] == ':' {
				// No value in this field match the ticket
				return ""
			}
			// Change position of cursor to next value
			i = e + 1
		}
	}
	return ruleID
}
*/
