package movie

type Character struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Concept struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Image struct {
	IconURL        string `json:"icon_url"`
	MediumURL      string `json:"medium_url"`
	ScreenURL      string `json:"screen_url"`
	ScreenLargeURL string `json:"screen_large_url"`
	SmallURL       string `json:"small_url"`
	SuperURL       string `json:"super_url"`
	ThumbURL       string `json:"thumb_url"`
	TinyURL        string `json:"tiny_url"`
	OriginalURL    string `json:"original_url"`
	ImageTags      string `json:"image_tags"`
}

type Location struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Object struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Producer struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Studio struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Team struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Writer struct {
	APIDetailURL  string `json:"api_detail_url"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailURL string `json:"site_detail_url"`
}

type Results struct {
	APIDetailURL     string      `json:"api_detail_url"`
	BoxOfficeRevenue string      `json:"box_office_revenue"`
	Budget           string      `json:"budget"`
	Characters       []Character `json:"characters"`
	Concepts         []Concept   `json:"concepts"`
	DateAdded        string      `json:"date_added"`
	DateLastUpdated  string      `json:"date_last_updated"`
	Deck             string      `json:"deck"`
	Description      string      `json:"description"`
	Distributor      any         `json:"distributor"`
	HasStaffReview   any         `json:"has_staff_review"`
	ID               int         `json:"id"`
	Image            Image       `json:"image"`
	Locations        []Location  `json:"locations"`
	Name             string      `json:"name"`
	Objects          []Object    `json:"objects"`
	Producers        []Producer  `json:"producers"`
	Rating           string      `json:"rating"`
	ReleaseDate      string      `json:"release_date"`
	Runtime          string      `json:"runtime"`
	SiteDetailURL    string      `json:"site_detail_url"`
	Studios          []Studio    `json:"studios"`
	Teams            []Team      `json:"teams"`
	TotalRevenue     string      `json:"total_revenue"`
	Writers          []Writer    `json:"writers"`
}

type MainType struct {
	Error                string  `json:"error"`
	Limit                int     `json:"limit"`
	Offset               int     `json:"offset"`
	NumberOfPageResults  int     `json:"number_of_page_results"`
	NumberOfTotalResults int     `json:"number_of_total_results"`
	StatusCode           int     `json:"status_code"`
	Results              Results `json:"results"`
	Version              string  `json:"version"`
}
