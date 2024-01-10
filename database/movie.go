package database

import (
	"github/chinmay/gocrudmux/models"
)

var DbMovies = models.IMovies{
	{Id: 1, Name: "Avengers", Cast: []string{"RDJ", "chris hemsworth"}, Price: 3.3},
	{Id: 2, Name: "Twilight", Cast: []string{"Kristen", "robert"}, Price: 3.3},
	{Id: 3, Name: "Batman", Cast: []string{"christan bale", "anna hathway"}, Price: 3.3},
	{Id: 4, Name: "Jack Reacher", Cast: []string{"tom curise", "Lee child"}, Price: 3.3},
	{Id: 5, Name: "Superman", Cast: []string{"Henry Cavill", "russell"}, Price: 3.3},
}
