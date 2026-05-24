package data

import (
	"fmt"
)

type Player interface {
	GetHealth() int
	DealDamage(damage int)
	GetVestID() int
}

type player struct {
	Hit_points int
	Vest_ID    int
}

func NewPlayer(vest_id int) Player {
	p := &player{
		Hit_points: 1,
		Vest_ID:    vest_id,
	}
	fmt.Println("Created Player")
	return p
}

func (p *player) DealDamage(damage int) {
	var old_hp = p.Hit_points
	p.Hit_points = p.Hit_points - damage
	if p.Hit_points < 0 {
		p.Hit_points = 0
	}
	fmt.Println("Reduced Hit Points from %i to %i", old_hp, p.Hit_points)
}

func (p *player) GetHealth() int {
	return p.Hit_points
}

func (p *player) GetVestID() int {
	return p.Vest_ID
}
