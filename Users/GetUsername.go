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
func (u *User) GetUsername () error {
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "username"
	query = query + " " + "FROM"
	query = query + " " + "users"
	query = query + " " + "WHERE"
	query = query + " " + "user_id = '" + fmt.Sprint(u.UID) + "'"
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

	for res.Next() {

		err = res.Scan(&u.Username)

		if err != nil {

			return err

		}

	}

	return nil

}

// ------------------------------------------------------------------ //
