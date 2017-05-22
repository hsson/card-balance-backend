package backend

// Config is used to conigure the backend
type Config struct {
	CardSiteURL string `yaml:"card_site_url"`
	Restaurants []struct {
		Name       string  `yaml:"name"`
		MenuURL    string  `yaml:"menu_url"`
		ImageURL   string  `yaml:"image"`
		WebsiteURL string  `yaml:"website"`
		Rating     float32 `yaml:"rating"`
	} `yaml:"restaurants"`
}
