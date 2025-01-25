package main

type locAreaJSON struct {
	Name string
	Url  string
}

type mapJSON struct {
	Next     string
	Previous string
	Results  []locAreaJSON
}
