//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//

package if_else

import "fmt"

type DayOfWeek byte

const (
	Sunday DayOfWeek = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d DayOfWeek) String() string {
	return []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[d]
}

type Role byte

const (
	Admin Role = iota
	Manager
	Contractor
	Member
	Guest
)

func (d Role) String() string {
	return []string{"Admin", "Manager", "Contractor", "Member", "Guest"}[d]
}

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func IfElse() {
	// The day and role. Change these to check your work.
	today, role := Tuesday, Guest

	fmt.Println(today, role)

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
