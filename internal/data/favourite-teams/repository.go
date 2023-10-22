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

func (r FavouriteTeamsRepository) GetFavouriteTeams(userId string) ([]favouriteTeamModel.FavouriteTeam, error) {

	var favTeams []favouriteTeamModel.FavouriteTeam

	rows, err := r.conn.Db.Query("SELECT * FROM FavouriteTeams WHERE userId = ?", userId)
	if err != nil {
		return nil, fmt.Errorf("GetFavouriteTeams %s: %v", userId, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var favTeam favouriteTeamModel.FavouriteTeam
		if err := rows.Scan(&favTeam.Id, &favTeam.UserId, &favTeam.TeamName); err != nil {
			return nil, fmt.Errorf("GetFavouriteTeams %s: %v", userId, err)
		}
		favTeams = append(favTeams, favTeam)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetFavouriteTeams %s: %v", userId, err)
	}

	return favTeams, nil
}

func (r FavouriteTeamsRepository) DeleteFavouriteTeam(userId string, teamName string) (int64, error) {
	var favTeamResult favouriteTeamModel.FavouriteTeam
	row := r.conn.Db.QueryRow("DELETE FROM FavouriteTeams WHERE userId = ? AND favouriteTeam = ?", userId, teamName)
	if err := row.Scan(&favTeamResult.Id, &favTeamResult.UserId, &favTeamResult.TeamName); err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("DeleteFavouriteTeam %s: user does not have a favourite team to be deleted", userId)
		}
		return 0, fmt.Errorf("DeleteFavouriteTeam %s: %v", userId, err)
	}
	return 1, nil
}
