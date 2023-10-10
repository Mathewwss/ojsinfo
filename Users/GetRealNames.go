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

// Get Real name all languages
func (u *User) GetRealNames () error {
	query := fmt.Sprint("SELECT")
	query = query + " " + "s1.locale, s1.setting_value,"
	query = query + " " + "s2.setting_value"
	query = query + " " + "FROM"
	query = query + " " + "("
	query = query + " " + "SELECT DISTINCT"
	query = query + " " + "locale, setting_value"
	query = query + " " + "FROM"
	query = query + " " + "user_settings"
	query = query + " " + "WHERE"
	query = query + " " + "setting_name = 'givenName'"
	query = query + " " + "AND user_id = '" + fmt.Sprint(u.UID) + "'"
	query = query + " " + "ORDER BY"
	query = query + " " + "locale"
	query = query + " " + ") AS s1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "("
	query = query + " " + "SELECT DISTINCT"
	query = query + " " + "locale, setting_value"
	query = query + " " + "FROM"
	query = query + " " + "user_settings"
	query = query + " " + "WHERE"
	query = query + " " + "setting_name = 'familyName'"
	query = query + " " + "AND user_id = '" + fmt.Sprint(u.UID) + "'"
	query = query + " " + "ORDER BY"
	query = query + " " + "locale"
	query = query + " " + ") AS s2"
	query = query + " " + "ON"
	query = query + " " + "s1.locale = s2.locale"
	query = query + " " + "ORDER BY"
	query = query + " " + "s1.locale"
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
	first := ""
	middles := ""

	u.RealNames = map[string]string{}

	for res.Next() {

		err = res.Scan(&locale, &first, &middles)

		if err != nil {

			return err

		}

		if first + middles == "" {

			continue

		}

		u.RealNames[locale] = first + " " + middles

	}

	return nil

}

// ------------------------------------------------------------------ //
