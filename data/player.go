package data

import (
	"fmt"
	"strconv"
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
	p.Hit_points = max(p.Hit_points-damage, 0)
	fmt.Println("Reduced Hit Points from " + strconv.Itoa(old_hp) + " to " + strconv.Itoa(p.Hit_points))
}

func (p *player) GetHealth() int {
	return p.Hit_points
}

func (p *player) GetVestID() int {
	return p.Vest_ID
}
