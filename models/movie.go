package models

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
