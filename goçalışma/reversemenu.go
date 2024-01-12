package main

func ReverseMenuIndex(menu []string) []string {

	cikti := []string{}

	for i, _ := range menu {

		cikti = append(cikti, menu[len(menu)-1-i])
	}
}
