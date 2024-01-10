package models

import (
	"strings"
)

type IMovie struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Cast  []string `json:"cast"`
	Price float64  `json:"price"`
}

type IMovies []IMovie

func (m *IMovies) AddMovie(p IMovie) IMovies {
	*m = append(*m, p)
	return *m
}

func (m *IMovies) FindByName(s string) IMovies {
	var arr IMovies
	for _, v := range *m {

		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(s)) {
			arr = append(arr, v)
		}
	}
	return arr
}
