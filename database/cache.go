package database

import (
	"pricing/models"
)

// initCache Copies everything from sql database to redis as cache.
func (db *Database) initCache() (err error) {

	// Add sql database names and ids into redis.

	tables := []string{"cities", "airlines", "agencies", "suppliers"}

	// For loop that iterates over tables list and
	// fetch data from sql database and then add
	// them to redis database as redis sorted sets.
	for _, table := range tables {
		var data []models.Field
		err = db.gorm.Table(table).Find(&data).Error
		if err != nil {
			return
		}
		// Create an empty slice where redis commands will be placed
		a := make([]interface{}, 2+2*len(data))
		a[0], a[1] = "zadd", table
		// Add each id and name as score and member
		// of redis sorted sets, in commands slice.
		for i := range data {
			a[2+2*i] = data[i].ID
			a[2+2*i+1] = data[i].Name
		}
		// Execute redis commands
		err = db.redis.Do(a...).Err()
		if err != nil {
			return
		}
	}

	// Add rules from sql database into redis.

	var rules []models.Rule
	err = db.gorm.Preload(
		"RuleAirlines").Preload(
		"RuleAgencies").Preload(
		"RuleSuppliers").Preload(
		"RuleRoutes").Find(&rules).Error
	if err != nil {
		return
	}

	if len(rules) < 1 {
		return nil
	}

	a := make([]interface{}, 2+2*len(rules))
	a[0], a[1] = "hset", "rules"
	// Add rules encoded string in commands slice.
	for i := range rules {
		a[2+2*i], a[2+2*i+1] = rules[i].Encode()
	}
	// Execute redis commands
	err = db.redis.Do(a...).Err()
	if err != nil {
		return
	}
	return nil
}
