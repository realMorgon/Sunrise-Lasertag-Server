class Player:

    def __init__ (self, hit_points):
        self.hit_points = hit_points

        print(f"Created new player with {hit_points} hit points. Phaser ID not set yet.")

    def get_hit_points(self):
        return self.hit_points

    def deal_damage(self, damage):
        old_hp = self.hit_points
        self.hit_points = self.hit_points - damage
        if self.hit_points < 0:
           	self.hit_points = 0
        print(f"Reduced Hit Points from {old_hp} to {self.hit_points}")

    def set_phaser_id(self, phaser_id):
        self.phaser_id = phaser_id
        print(f"Set phaser id to {phaser_id}")

    def get_phaser_id(self):
        return self.phaser_id

    def set_nickname(self, nickname):
        self.nickname = nickname
        print(f"Set nickname to {nickname}")

    def get_nickname(self):
        return self.nickname
