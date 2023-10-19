// ----------------------------- Package ---------------------------- //

package Journals

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "database/sql"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func (j *Journal) GetGroups () error {
	// Sql query
	query := "SELECT DISTINCT"
	query = query + " " + "t1.locale, t1.user_group_id,"
	query = query + " " + "t1.setting_value"
	query = query + " " + "FROM"
	query = query + " " + "user_group_settings AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "user_groups AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.user_group_id = t2.user_group_id"
	query = query + " " + "INNER JOIN"
	query = query + " " + "journals AS t3"
	query = query + " " + "ON"
	query = query + " " + "t2.context_id = t3.journal_id"
	query = query + " " + "WHERE"
	query = query + " " + "t1.setting_name = 'name'"
	query = query + " " + "AND t3.journal_id = '" + fmt.Sprint(j.ID)
	query = query + "'"
	query = query + " " + "ORDER BY"
	query = query + " " + "t1.locale ASC, t1.user_group_id ASC"
	query = query + ";"

	// Database connection settings
	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	// Connect db
	db, err := sql.Open(driver, con)

	// Finish Connection
	defer db.Close()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Check connection
	err = db.Ping()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Run query
	res, err := db.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Start variable
	group := ""
	gid := -1
	locale := ""

	// Start map
	j.Groups = map[int]map[string]string{}

	// View results
	for res.Next() {
		// Check errors
		err := res.Scan(&locale, &gid, &group)
		if err != nil {
			// Stop
			return err

		}

		// Check map
		if len(j.Groups[gid]) == 0 {
			// start map
			j.Groups[gid] = map[string]string{}

		}

		// Update map
		j.Groups[gid][locale] = group

	}

	// Finish
	return nil

}

// ------------------------------------------------------------------ //
