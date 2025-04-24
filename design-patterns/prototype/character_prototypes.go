package prototype

// EnemyPrototype is an abstract enemy NPC prototype
var EnemyPrototype = Character{
	behavior: Hostile,
	Speed:    5.0,
	Shape: &Shape{
		Head: ShapePart{
			Color: "#000",
			model: Pyramid,
			Size:  1,
			Count: 1,
		},
		Body: ShapePart{
			Color: "#111",
			model: Cone,
			Size:  1.5,
			Count: 1,
		},
		Legs: ShapePart{
			Color: "#000",
			model: Cylinder,
			Size:  2,
			Count: 2,
		},
	},
}

// FriendPrototype is an abstract friend NPC prototype
var FriendPrototype = Character{
	behavior: Friendly,
	Speed:    8.5,
	Shape: &Shape{
		Head: ShapePart{
			Color: "#fff",
			model: Sphere,
			Size:  1,
			Count: 1,
		},
		Body: ShapePart{
			Color: "#eee",
			model: Cylinder,
			Size:  1.5,
			Count: 1,
		},
		Legs: ShapePart{
			Color: "#ddd",
			model: Cylinder,
			Size:  2,
			Count: 2,
		},
	},
}
