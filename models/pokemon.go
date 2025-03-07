package models

type Pokemon struct {
	NationalDexID int    `yaml:"national"`
	Name          string `yaml:"name"`
	Generation    int    `yaml:"gen"`
}
