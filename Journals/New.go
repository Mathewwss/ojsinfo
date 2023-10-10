// ----------------------------- Package ---------------------------- //

package Journals

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "database/sql"
import "errors"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

// Create journal with exists on database
func New (identity string) (Journal, error) {
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "journal_id, path"
	query = query + " " + "FROM"
	query = query + " " + "journals"
	query = query + " " + "WHERE"
	query = query + " " + "path = '" + identity + "'"
	query = query + ";"

	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	db, err := sql.Open(driver, con)

	if err != nil {

		return Journal{}, err

	}

	err = db.Ping()

	if err != nil {

		return Journal{}, err

	}

	res, err := db.Query(query)

	if err != nil {

		return Journal{}, err

	}

	j := Journal{}

	j.ID = -1

	for res.Next() {

		err = res.Scan(&j.ID, &j.Path)

		if err != nil {

			return Journal{}, err

		}

	}

	if j.ID == -1 {

		msg := "[ERROR] -> Not found journal by path '" + identity
		msg = msg + "'!"
		err = errors.New(msg)

		return Journal{}, err

	}

	return j, nil

}

// ------------------------------------------------------------------ //
