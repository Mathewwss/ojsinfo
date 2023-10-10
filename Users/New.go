// ----------------------------- Package ---------------------------- //

package Users

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

// Create user with exists on database
func New (identity string) (User, error) {
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "user_id"
	query = query + " " + "FROM"
	query = query + " " + "users"
	query = query + " " + "WHERE"
	query = query + " " + "email = '" + identity + "'"
	query = query + " " + "OR username = '" + identity + "'"
	query = query + ";"

	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	db, err := sql.Open(driver, con)

	if err != nil {

		return User{}, err

	}

	err = db.Ping()

	if err != nil {

		return User{}, err

	}

	res, err := db.Query(query)

	if err != nil {

		return User{}, err

	}

	u := User{}

	for res.Next() {

		err = res.Scan(&u.UID)

		if err != nil {

			return User{}, err

		}

	}

	if u.UID == 0 {

		msg := "[ERROR] -> Not found user!"
		err = errors.New(msg)

		return User{}, err

	}

	return u, nil

}

// ------------------------------------------------------------------ //
