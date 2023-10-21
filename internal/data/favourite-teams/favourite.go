package favouriteteams

import (
	"fmt"

	databaseconn "github.com/RazvanBerbece/UnifyFootballBot/internal/database-conn"
)

type FavouriteTeamsRepositoryInterface interface {
	InsertFavouriteTeam(string, string) error
	RemoveFavouriteTeam(string) error
}

type FavouriteTeamsRepository struct {
	conn databaseconn.Database
}

func NewFavouriteTeamsRepository() *FavouriteTeamsRepository {
	repo := new(FavouriteTeamsRepository)
	repo.conn.ConnectDatabaseHandle()
	return repo
}

func (r FavouriteTeamsRepository) InsertFavouriteTeam(userId string, teamName string) (int64, error) {
	result, err := r.conn.Db.Exec("INSERT INTO FavouriteTeams (userId, favouriteTeam) VALUES (?, ?)", userId, teamName)
	if err != nil {
		return 0, fmt.Errorf("InsertFavouriteTeam: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("InsertFavouriteTeam: %v", err)
	}
	return id, nil
}
