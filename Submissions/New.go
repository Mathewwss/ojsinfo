// ----------------------------- Package ---------------------------- //

package Submissions

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

// Create submission
func New (identity int) (Submission, error) {
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "submission_id"
	query = query + " " + "FROM"
	query = query + " " + "submissions"
	query = query + " " + "WHERE"
	query = query + " " + "submission_id = '" + fmt.Sprint(identity)
	query = query + "'"
	query = query + ";"

	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	db, err := sql.Open(driver, con)

	if err != nil {

		return Submission{}, err

	}

	err = db.Ping()

	if err != nil {

		return Submission{}, err

	}

	res, err := db.Query(query)

	if err != nil {

		return Submission{}, err

	}

	s := Submission{}
	s.ID = -1

	for res.Next() {

		err = res.Scan(&s.ID)

		if err != nil {

			return Submission{}, err

		}

	}

	if s.ID == -1 {

		msg := "[ERROR] -> Not found submission!"
		err = errors.New(msg)

		return Submission{}, err

	}

	return s, nil

}

// ------------------------------------------------------------------ //
