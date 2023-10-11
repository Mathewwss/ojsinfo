// ----------------------------- Package ---------------------------- //

package Submissions

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

// Get journal name by language
func (s *Submission) GetJournalNames () error {
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "t2.locale, t2.setting_value"
	query = query + " " + "FROM"
	query = query + " " + "submissions AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "journal_settings AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.context_id = t2.journal_id"
	query = query + " " + "WHERE"
	query = query + " " + "t2.setting_name = 'name'"
	query = query + " " + "AND t1.submission_id = '" + fmt.Sprint(s.ID)
	query = query + "'"
	query = query + " " + "ORDER BY"
	query = query + " " + "t2.locale"
	query = query + ";"

	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	db, err := sql.Open(driver, con)

	if err != nil {

		return err

	}

	err = db.Ping()

	if err != nil {

		return err

	}

	res, err := db.Query(query)

	if err != nil {

		return err

	}

	locale := ""
	name := ""

	s.JournalNames = map[string]string{}

	for res.Next() {

		err = res.Scan(&locale, &name)

		if err != nil {

			return err

		}

		s.JournalNames[locale] = name

	}

	return nil

}

// ------------------------------------------------------------------ //
