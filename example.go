// ----------------------------- Package ---------------------------- //

package main

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "github.com/Mathewwss/ojsinfo/Users"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func main () {
	// Configure database
	DbCfg.Db_conf.Driver = "mysql"
	DbCfg.Db_conf.Settings = "ojs:ojs@tcp(127.0.0.1:3306)/ojs"

	// Get user
	user, err := Users.New("rxvt")

	// Get info
	err = user.GetEmail()
	err = user.GetUsername()
	err = user.GetRealNames()
	err = user.GetGroups()

	// Show error
	fmt.Println(err)

	// Show user
	fmt.Println(user)
}

// ------------------------------------------------------------------ //
