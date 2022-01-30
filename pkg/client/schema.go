package client

type SearchPhotosRequest struct {
	// The search query. Ocean, Tigers, Pears, etc.
	Query string `json:"query" url:"query"`
	// Desired photo orientation. The current supported orientations are: landscape, portrait or square.
	Orientation string `json:"orientation" url:"orientation"`
	// Minimum photo size. The current supported sizes are: large(24MP), medium(12MP) or small(4MP).
	Size string `json:"size" url:"size"`
	// Desired photo color. Supported colors: red, orange, yellow, green, turquoise, blue, violet, pink, brown, black, gray, white or any hexidecimal color code (eg. #ffffff).
	Color string `json:"color" url:"color"`
	// The locale of the search you are performing. The current supported locales are: 'en-US' 'pt-BR' 'es-ES' 'ca-ES' 'de-DE' 'it-IT' 'fr-FR' 'sv-SE' 'id-ID' 'pl-PL' 'ja-JP' 'zh-TW' 'zh-CN' 'ko-KR' 'th-TH' 'nl-NL' 'hu-HU' 'vi-VN' 'cs-CZ' 'da-DK' 'fi-FI' 'uk-UA' 'el-GR' 'ro-RO' 'nb-NO' 'sk-SK' 'tr-TR' 'ru-RU'.
	Local string `json:"local" url:"local"`
	// The page number you are requesting. Default: 1
	Page int `json:"page" url:"page"`
	// The number of results you are requesting per page. Default: 15 Max: 80
	PerPage int `json:"per_page" url:"per_page"`
}
type SearchPhotosResponse struct {
	TotalResults int `json:"total_results"`
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
	Photos       []struct {
		ID              int    `json:"id"`
		Width           int    `json:"width"`
		Height          int    `json:"height"`
		URL             string `json:"url"`
		Photographer    string `json:"photographer"`
		PhotographerURL string `json:"photographer_url"`
		PhotographerID  int    `json:"photographer_id"`
		AvgColor        string `json:"avg_color"`
		Src             struct {
			Original  string `json:"original"`
			Large2X   string `json:"large2x"`
			Large     string `json:"large"`
			Medium    string `json:"medium"`
			Small     string `json:"small"`
			Portrait  string `json:"portrait"`
			Landscape string `json:"landscape"`
			Tiny      string `json:"tiny"`
		} `json:"src"`
		Liked bool   `json:"liked"`
		Alt   string `json:"alt"`
	} `json:"photos"`
	NextPage string `json:"next_page"`
}
