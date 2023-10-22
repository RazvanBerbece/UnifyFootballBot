package favouriteteams

import (
	"database/sql"
	"fmt"

	favouriteTeamModel "github.com/RazvanBerbece/UnifyFootballBot/internal/data/favourite-teams/models"
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

func (r FavouriteTeamsRepository) GetFavouriteTeam(userId string) (string, error) {
	var favTeamResult favouriteTeamModel.FavouriteTeam
	row := r.conn.Db.QueryRow("SELECT * FROM FavouriteTeams WHERE userId = ?", userId)
	if err := row.Scan(&favTeamResult.Id, &favTeamResult.UserId, &favTeamResult.TeamName); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("GetFavouriteTeam %s: user does not have a favourite team", userId)
		}
		return "", fmt.Errorf("GetFavouriteTeam %s: %v", userId, err)
	}
	return favTeamResult.TeamName, nil
}

func (r FavouriteTeamsRepository) DeleteFavouriteTeam(userId string, teamName string) (int64, error) {
	var favTeamResult favouriteTeamModel.FavouriteTeam
	row := r.conn.Db.QueryRow("DELETE FROM FavouriteTeams WHERE userId = ? AND favouriteTeam = ?", userId, teamName)
	if err := row.Scan(&favTeamResult.Id, &favTeamResult.UserId, &favTeamResult.TeamName); err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("DeleteFavouriteTeam %s: user does not have a favourite team %s to be deleted", teamName, userId)
		}
		return 0, fmt.Errorf("DeleteFavouriteTeam %s, %s: %v", userId, teamName, err)
	}
	return 1, nil
}
