// ----------------------------- Package ---------------------------- //

package main

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "github.com/Mathewwss/ojsinfo/Users"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

type funcs func () error

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func main () {
	// Configure database
	DbCfg.Db_conf.Driver = "mysql"
	DbCfg.Db_conf.Settings = "ojs:ojs@tcp(127.0.0.1:3306)/ojs"

	// Try connect database
	DbCfg.Db_conf.OpenCon()

	// View errors
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Show error
		fmt.Println(err)

		// Stop
		return

	}

	// Example using users
	u, err := Users.New("user_test")

	// CHeck errors
	if err != nil {
		// Show erros
		fmt.Println(err)

		// Stop
		return

	}

	// All functions
	list := []funcs{
		u.GetEmail,
		u.GetRealNames,
		u.GetUsername,
		u.GetGroups,
	}

	// View functions
	for a := 0; a < len(list); a++ {
		// Run function
		err := list[a]()

		// Check errors
		if err != nil {
			// Show error
			fmt.Println(err)

			// Stop
			return

		}
	}

	// Show user
	fmt.Println(u)

	// Close database connection
	DbCfg.Db_conf.Con.Close()
}

// ------------------------------------------------------------------ //
