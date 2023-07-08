//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//

package if_else

import "fmt"

// Days of the week
const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func IfElse() {
	// The day and role. Change these to check your work.
	today, role := Tuesday, Guest

	//* Access at any time: Admin, Manager
	if role == Admin || role == Manager {
		accessGranted()
		//* Access weekends: Contractor
	} else if role == Contractor && (today == Saturday || today == Sunday) {
		accessGranted()
		//* Access weekdays: Member
	} else if role == Member && (today >= Monday && today <= Friday) {
		accessGranted()
		//* Access Mondays, Wednesdays, and Fridays: Guest
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		accessGranted()
	} else {
		accessDenied()
	}

}
