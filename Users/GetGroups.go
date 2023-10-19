// ----------------------------- Package ---------------------------- //

package Users

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

// Get login name
func (u *User) GetGroups () error {
	// Sql query
	query := "SELECT DISTINCT"
	query = query + " " + "t1.user_group_id, t3.setting_value,"
	query = query + " " + "t4.setting_value"
	query = query + " " + "FROM"
	query = query + " " + "user_user_groups AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "user_groups AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.user_group_id = t2.user_group_id"
	query = query + " " + "INNER JOIN"
	query = query + " " + "user_group_settings AS t3"
	query = query + " " + "ON"
	query = query + " " + "t1.user_group_id = t3.user_group_id"
	query = query + " " + "INNER JOIN"
	query = query + " " + "journal_settings AS t4"
	query = query + " " + "ON"
	query = query + " " + "t2.context_id = t4.journal_id"
	query = query + " " + "WHERE"
	query = query + " " + "t3.locale = 'pt_BR'"
	query = query + " " + "AND t3.setting_name = 'name'"
	query = query + " " + "AND t4.locale = 'pt_BR'"
	query = query + " " + "AND t4.setting_name = 'name'"
	query = query + " " + "AND t1.user_id = '" + fmt.Sprint(u.UID) + "'"
	query = query + " " + "ORDER BY"
	query = query + " " + "t4.setting_value ASC,"
	query = query + " " + "t1.user_group_id ASC"
	query = query + ";"

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
	journal := ""
	group := ""
	gid := -1

	// Start map
	u.Groups = map[string]map[int]string{}

	// View results
	for res.Next() {
		// Check errors
		err := res.Scan(&gid, &group, &journal)
		if err != nil {
			// Stop
			return err

		}

		// Check map
		if len(u.Groups[journal]) == 0 {
			// start map
			u.Groups[journal] = map[int]string{}

		}

		// Update map
		u.Groups[journal][gid] = group

	}

	// Show user
	return nil
}

// ------------------------------------------------------------------ //
