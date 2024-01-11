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

func (m *IMovies) DeleteMovie(s string) int {

	temp := *m
	index := -1

	for idx, v := range *m {
		if strings.EqualFold(strings.ToLower(v.Name), strings.ToLower(s)) {
			index = idx
		}
	}

	if index == -1 {
		return -1
	} else if index == 0 {
		*m = temp[index+1:]
		return 1
	} else if index == len(temp)-1 {
		*m = temp[:index]
		return 1
	} else {
		temp = append(temp[:index], temp[index+1:]...)
		*m = temp
		return 1
	}
}
