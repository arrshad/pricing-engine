package database

import (
	"errors"
	"log"
	"pricing/models"
	"pricing/utils"
)

// AddRule inserts given rawRule into database and cache
func (db *Database) AddRule(rawRule models.RawRule) (int, error) {

	var err error

	var rule models.Rule
	// Add routes to rule model
	for i := range rawRule.Routes {

		var oID, dID = 1, 1

		// Get origin id and validate it
		if rawRule.Routes[i].Origin != "" {
			oID, err = db.getValidID(CITY, rawRule.Routes[i].Origin)
			if err != nil {
				return 0, err
			}
		}

		// Get destination id and validate it
		if rawRule.Routes[i].Destination != "" {
			dID, err = db.getValidID(CITY, rawRule.Routes[i].Destination)
			if err != nil {
				return 0, err
			}
		}

		rule.RuleRoutes = append(rule.RuleRoutes, models.RuleRoute{
			OriginID:      oID,
			DestinationID: dID,
		})
	}

	// Add airlines to rule model
	for i := range rawRule.Airlines {
		id, err := db.getValidID(AIRLINE, rawRule.Airlines[i])
		if err != nil {
			return 0, err
		}
		rule.RuleAirlines = append(
			rule.RuleAirlines,
			models.RuleAirline{AirlineID: id},
		)
	}

	// Add agencies to rule model
	for i := range rawRule.Agencies {
		id, err := db.getValidID(AGENCY, rawRule.Agencies[i])
		if err != nil {
			return 0, err
		}
		rule.RuleAgencies = append(
			rule.RuleAgencies,
			models.RuleAgency{AgencyID: id},
		)
	}

	// Add suppliers to rule model
	for i := range rawRule.Suppliers {
		id, err := db.getValidID(SUPPLIER, rawRule.Suppliers[i])
		if err != nil {
			return 0, err
		}
		rule.RuleSuppliers = append(
			rule.RuleSuppliers,
			models.RuleSupplier{SupplierID: id},
		)
	}

	rule.IsPercent = rawRule.AmountType == "PERCENTAGE"
	rule.Amount = rawRule.AmountValue

	// Insert rule model into sql database
	err = db.gorm.Create(&rule).Error
	if err != nil {
		return 0, err
	}
	// Add rule model into redis
	key, val := rule.Encode()
	err = db.redis.HSet("rules", key, val).Err()
	if err != nil {
		return 0, err
	}

	return rule.ID, nil
}

// FindRules returns a slice of matching rules
// as an array with given ticket for each ticket
func (db *Database) FindRules(t []models.RawTicket) ([][][2]string, error) {

	n := len(t)

	tickets := make([]models.Ticket, n)

	var err error

	// Get ids and store them as ticket struct
	// It returns if id not found
	for i := range tickets {
		tickets[i].OriginID = db.getID(CITY, t[i].Origin)
		tickets[i].DestinationID = db.getID(CITY, t[i].Destination)
		tickets[i].AirlineID = db.getID(AIRLINE, t[i].Airline)
		tickets[i].AgencyID = db.getID(AGENCY, t[i].Agency)
		tickets[i].SupplierID = db.getID(SUPPLIER, t[i].Supplier)
	}

	// Use redis HSCAN to get all rules for checking
	iter := db.redis.HScan("rules", 0, "", 0).Iterator()
	isKey := true
	lastRule := ""
	isMatch := make([]bool, n)

	rules := make([][][2]string, n)

	for iter.Next() {
		val := iter.Val()
		// The HSCAN command return key in first line
		// and value in the next row, so here we should
		// know which row is checking
		if isKey {
			// Check if the ticket matches given rule
			isMatch, lastRule = utils.Match(tickets, val)
			// Next row is a value
			isKey = false
		} else {
			for i := range isMatch {
				if isMatch[i] {
					rules[i] = append(rules[i], [2]string{lastRule, val})
				}
			}
			// Next row is a key
			isKey = true
		}
	}
	if err = iter.Err(); err != nil {
		log.Println(err)
		return rules, errors.New("something went wrong")
	}
	return rules, nil
}
