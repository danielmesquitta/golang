// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package tests

import (
	"testing"

	rcvFunc "github.com/danielmesquitta/golang/11_rcv_func"
)

func TestHeal(t *testing.T) {
	player := rcvFunc.NewPlayer("Mario", 100, 100)

	player.Health = 50

	player.Heal(10)
	if player.Health != 60 {
		t.Errorf("Player health is wrong, got %d, want %d", player.Health, 60)
	}

	player.Heal(100)
	if player.Health != 100 {
		t.Errorf("Player health should not go above maximum")
	}
}

func TestHurt(t *testing.T) {
	player := rcvFunc.NewPlayer("Mario", 100, 100)

	player.Hurt(50)

	if player.Health != 50 {
		t.Errorf("Player health is wrong, got %d, want %d", player.Health, 50)
	}

	player.Hurt(100)
	if player.Health != 0 {
		t.Errorf("Player health should not go bellow 0")
	}
}

func TestRest(t *testing.T) {
	player := rcvFunc.NewPlayer("Mario", 100, 100)

	player.Energy = 50

	player.Rest(10)
	if player.Energy != 60 {
		t.Errorf("Player energy is wrong, got %d, want %d", player.Energy, 60)
	}

	player.Rest(100)
	if player.Energy != 100 {
		t.Errorf("Player energy should not go above maximum")
	}
}

func TestRun(t *testing.T) {
	player := rcvFunc.NewPlayer("Mario", 100, 100)

	player.Run(50)

	if player.Energy != 50 {
		t.Errorf("Player energy is wrong, got %d, want %d", player.Energy, 50)
	}

	player.Run(100)
	if player.Energy != 0 {
		t.Errorf("Player energy should not go bellow 0")
	}
}
