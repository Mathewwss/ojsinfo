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

func (j *Journal) GetGroups () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

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

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

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
