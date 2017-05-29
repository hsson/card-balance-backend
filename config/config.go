package backend

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config is used to conigure the backend
type Config struct {
	CardSiteURL string `yaml:"card_site_url"`
	Restaurants []struct {
		Name         string  `yaml:"name"`
		MenuURL      string  `yaml:"menu_url"`
		ImageURL     string  `yaml:"image"`
		WebsiteURL   string  `yaml:"website"`
		Rating       float32 `yaml:"rating"`
		AveragePrice int     `yaml:"avg_price"`
		Campus       string  `yaml:"campus"`
		OpenHours    []struct {
			DayOfWeek int `yaml:"day_of_week"`
			StartTime int `yaml:"start_hour"`
			EndTime   int `yaml:"end_hour"`
		} `yaml:"open_hours"`
	} `yaml:"restaurants"`
}

// LoadConfig loads config from given file path
func LoadConfig(configFile string) Config {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("ERROR: Could not load config file")
		panic(err)
	}
	conf := Config{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println("ERROR: Could not parse config file")
		panic(err)
	}
	return conf
}
