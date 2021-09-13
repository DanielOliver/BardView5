package models

type RPG struct {
	Id   string
	Name string
}

type RPGList []RPG

var AllRPGs = RPGList{
	{
		"dnd5e",
		"Dungeons & Dragons 5e",
	},
}
