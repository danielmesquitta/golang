//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package rcvFunc

import "fmt"

type Player struct {
	Name      string
	Health    uint
	MaxHealth uint
	Energy    uint
	MaxEnergy uint
}

func NewPlayer(name string, health uint, energy uint) Player {
	return Player{
		Name:      name,
		Health:    health,
		MaxHealth: health,
		Energy:    energy,
		MaxEnergy: energy,
	}
}

func (p *Player) Heal(healthIncrease uint) {
	p.Health += healthIncrease

	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
}

func (p *Player) Hurt(healthDecrease uint) {
	if p.Health < healthDecrease {
		p.Health = 0
	} else {
		p.Health -= healthDecrease
	}
}

func (p *Player) Rest(energyIncrease uint) {
	p.Energy += energyIncrease

	if p.Energy > p.MaxEnergy {
		p.Energy = p.MaxEnergy
	}
}

func (p *Player) Run(energyDecrease uint) {
	if p.Energy < energyDecrease {
		p.Energy = 0
	} else {
		p.Energy -= energyDecrease
	}

}

func (p *Player) PrintStatus() {
	fmt.Println("Player:", p.Name)
	fmt.Println("Health:", p.Health, "/", p.MaxHealth)
	fmt.Println("Energy:", p.Energy, "/", p.MaxEnergy)
	fmt.Println()
}

func RcvFunc() {
	player := NewPlayer("Mario", 100, 100)

	player.PrintStatus()

	player.Hurt(25)
	fmt.Println(player.Name, "was hurt!")

	player.PrintStatus()

	player.Heal(10)
	fmt.Println(player.Name, "healed!")

	player.PrintStatus()

	player.Run(150)
	fmt.Println(player.Name, "ran!")

	player.PrintStatus()

	player.Rest(10)
	fmt.Println(player.Name, "rested!")

	player.PrintStatus()
}
