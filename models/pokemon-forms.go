package models

type Stats struct {
	HP              int `yaml:"hp"`
	Attack          int `yaml:"attack"`
	Defence         int `yaml:"defence"`
	Special_attack  int `yaml:"spatk"`
	Special_defence int `yaml:"spdef"`
	Speed           int `yaml:"speed"`
}

type PokemonForms struct {
	Name           string `yaml:"name"`
	Release        string `yaml:"release"`
	Type1          string `yaml:"type1"`
	Type2          string `yaml:"type2"`
	Stats          Stats  `yaml:"stats"`
	Species        string `yaml:"species"`
	Height         int    `yaml:"height"`
	Weight         int    `yaml:"weight"`
	Gender         string `yaml:"gender"`
	CatchRate      int    `yaml:"catch-rate"`
	BaseExperience int    `yaml:"base-exp"`
	EggCycles      int    `yaml:"egg-cycles"`
	Friendship     string `yaml:"friendship"`
	GrowthRate     string `yaml:"growth-rate"`
	EVYield        Stats  `yaml:"ev-yield"`
}
