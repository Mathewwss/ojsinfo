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

// Get journal groups
func (j *Journal) GetSections () error {
	// Sql query
	query := "SELECT DISTINCT"
	query = query + " " + "t1.locale, t1.section_id,"
	query = query + " " + "t1.setting_value"
	query = query + " " + "FROM"
	query = query + " " + "section_settings AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "sections AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.section_id = t2.section_id"
	query = query + " " + "WHERE"
	query = query + " " + "t1.setting_name = 'title'"
	query = query + " " + "AND t2.journal_id = '" + fmt.Sprint(j.ID)
	query = query + "'"
	query = query + " " + "ORDER BY"
	query = query + " " + "t1.locale ASC, t1.section_id ASC"
	query = query + ";"

	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	// Connect db
	db, err := sql.Open(driver, con)

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
	section_name := ""
	section_id := -1
	locale := ""

	// Start map
	j.Sections = map[int]map[string]string{}

	// View results
	for res.Next() {
		// Check errors
		err := res.Scan(&locale, &section_id, &section_name)
		if err != nil {
			// Stop
			return err

		}

		// Check map
		if len(j.Sections[section_id]) == 0 {
			// start map
			j.Sections[section_id] = map[string]string{}

		}

		// Update map
		j.Sections[section_id][locale] = section_name

	}

	// Finish
	return nil

}

// ------------------------------------------------------------------ //
