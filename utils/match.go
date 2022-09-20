package utils

import (
	"pricing/models"
	"strings"
)

// Match returns true if for each ticket matches the given rule
func Match(tickets []models.Ticket, rule string) ([]bool, string) {
	// The reason I wrote such a function myself was
	// because it's up to 7x faster than using golang
	// split and contains functions

	// Initializing search variables
	ok := make([]bool, len(tickets))
	for i := range ok {
		ok[i] = true
	}
	found := make([]bool, len(tickets))
	for i := range found {
		found[i] = false
	}

	// Remove the 'rule:' from beginning
	rule = rule[5:]
	// Store id of the current rule
	n := strings.IndexByte(rule, ':')
	ruleID := rule[:n]
	// Index of rule string
	i := n + 1
	// Position of current field
	f := 1

	for i < len(rule) {
		if isAll(ok, false) {
			return ok, ""
		}
		// Store rest of rule string
		rule = rule[i:]
		// Check if this is an empty field
		if rule[0] == ':' {
			// Field is null and it means 'all',
			// so we go to the next field
			i = 1
			f += 1
			continue
		}
		// Get next field index
		n = strings.IndexByte(rule, ':')
		// Find end of current value
		e := strings.IndexByte(rule[:n], ',')
		// Store the current value
		c := rule[:e]

		// Routes field is a little different,
		// so I seperated checking method of it
		if f == 1 {
			// Find origin and destination separator
			m := strings.IndexByte(c, '-')
			// Check if current route match the ticket
			o := c[:m]
			d := c[m+1:]
			for t := range tickets {
				if !found[t] {
					// City id 1 means 'all' in database
					if (o == "1" || o == tickets[t].OriginID) &&
						(d == "1" || d == tickets[t].DestinationID) {
						found[t] = true
					}
				}
			}
		} else {
			// Check if current value match the ticket
			for t := range tickets {
				if ok[t] && !found[t] {
					if tickets[t].Get(f) == c {
						// Value found
						found[t] = true
					}
				}
			}
		}
		if isAll(found, true) {
			for j := range found {
				// Reset values of found
				found[j] = false
			}
			// Going to next field
			f += 1
			i = n + 1
			continue
		}
		if e+1 == n {
			if isAll(found, false) {
				return found, ""
			}
			for j := range ok {
				// Update ok status of ticket
				if ok[j] && !found[j] {
					ok[j] = false
				}
				found[j] = false
			}
			// Going to next field
			f += 1
			i = n + 1
			continue
		}
		// Change position of cursor to next value
		i = e + 1
	}
	return ok, ruleID
}

// isAll returns true if all members of
// gievn slice have a true value
func isAll(s []bool, b bool) bool {
	for i := range s {
		if s[i] != b {
			return false
		}
	}
	return true
}
