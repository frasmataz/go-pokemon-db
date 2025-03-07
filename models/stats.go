package models

type Stats struct {
	HP              int `yaml:"hp"`
	Attack          int `yaml:"attack"`
	Defence         int `yaml:"defence"`
	Special_attack  int `yaml:"spatk"`
	Special_defence int `yaml:"spdef"`
	Speed           int `yaml:"speed"`
}
