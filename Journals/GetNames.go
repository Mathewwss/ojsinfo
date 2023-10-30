// ----------------------------- Package ---------------------------- //

package Journals

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func (j *Journal) GetNames () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Sql query
	query := "SELECT DISTINCT"
	query = query + " " + "locale, setting_value"
	query = query + " " + "FROM"
	query = query + " " + "journal_settings"
	query = query + " " + "WHERE"
	query = query + " " + "setting_name = 'name'"
	query = query + " " + "AND journal_id = '" + fmt.Sprint(j.ID) + "'"
	query = query + " " + "ORDER BY"
	query = query + " " + "locale"
	query = query + ";"

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Start variable
	name := ""
	locale := ""
	j.Names = map[string]string{}

	// View results
	for res.Next() {
		// Check errors
		err := res.Scan(&locale, &name)
		if err != nil {
			// Stop
			return err

		}

		// Update map
		j.Names[locale] = name

	}

	// Finish
	return nil

}

// ------------------------------------------------------------------ //
