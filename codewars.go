package main

// User - user information.
type User struct {
	Username            string         `json:"username"`
	Name                string         `json:"name"`
	Honor               int            `json:"honor"`
	Clan                string         `json:"clan"`
	LeaderboardPosition int            `json:"leaderboardPosition"`
	Skills              []string       `json:"skills"`
	Rank                Ranks          `json:"ranks"`
	CodeChallenges      CodeChallenges `json:"codeChallenges"`
}

// Ranks - user ranking information
type Ranks struct {
	Overall   Overall   `json:"overall"`
	Languages Languages `json:"languages"`
}

// Overall - overall user information
type Overall struct {
	Rank  int    `json:"rank"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Score int    `json:"score"`
}

// CodeChallenges - code challenges created and completed
type CodeChallenges struct {
	TotalAuthored  int `json:"totalAuthored"`
	TotalCompleted int `json:"totalCompleted"`
}

// Language - language ranking information
type Language struct {
	Rank  int    `json:"rank"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Score int    `json:"score"`
}

// Languages - set of languages supported by codewars
type Languages struct {
	C           Language `json:"c, omitempty"`
	Closure     Language `json:"closure, omitempty"`
	Coffescript Language `json:"coffescript, omitempty"`
	Cplusplus   Language `json:"c++, omitempty"`
	Crystal     Language `json:"crystal, omitempty"`
	CSharp      Language `json:"c#, omitempty"`
	Dart        Language `json:"dart, omitempty"`
	Elixit      Language `json:"elixir, omitempty"`
	FSharp      Language `json:"f#, omitempty"`
	Haskell     Language `json:"haskell, omitempty"`
	Java        Language `json:"java, omitempty"`
	Javascript  Language `json:"javascript, omitempty"`
	ObjectiveC  Language `json:"objective-c, omitempty"`
	OCaml       Language `json:"ocaml, omitempty"`
	PHP         Language `json:"php, omitempty"`
	Python      Language `json:"python, omitempty"`
	Ruby        Language `json:"ruby, omitempty"`
	Rust        Language `json:"rust, omitempty"`
	Shell       Language `json:"shell, omitempty"`
	SQL         Language `json:"sql, omitempty"`
	Swift       Language `json:"swift, omitempty"`
	Typescript  Language `json:"typescript, omitempty"`
}
