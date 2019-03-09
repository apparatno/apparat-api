package app

type Employee struct {
	Candidate Candidate        `json:"candidate"`
	Work      []WorkExperience `json:"work"`
	Education []Education      `json:"education"`
}

type Candidate struct {
	Name      string          `json:"name"`
	Title     string          `json:"title"`
	EMail     string          `json:"mail"`
	Phone     string          `json:"phone"`
	Picture   string          `json:"picture"`
	Location  string          `json:"location"`
	Website   string          `json:"website"`
	Linkedin  string          `json:"linkedin"`
	GitHub    string          `json:"github"`
	Skills    []string        `json:"skills"`
	Languages []LanguageSkill `json:"languages"`
	Hobbies   []string        `json:"hobbies"`
}

type Education struct {
	Title    string `json:"title"`
	Location string `json:"location"`
	Period   string `json:"period"`
}

type WorkExperience struct {
	Company     string   `json:"company"`
	Description string   `json:"description"`
	Period      string   `json:"period"`
	Location    string   `json:"location"`
	Website     string   `json:"website"`
	Skills      string   `json:"skills"`
	Experience  []string `json:"experience"`
	References  []string `json:"references"`
}

type LanguageSkill struct {
	Lang string `json:"lang"`
	Rank int    `json:"rank"`
}
