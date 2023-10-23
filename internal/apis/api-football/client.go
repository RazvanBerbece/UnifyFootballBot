package apiFootball

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mozillazg/go-unidecode"

	apiFootballModels "github.com/RazvanBerbece/UnifyFootballBot/internal/apis/api-football/models"
	"github.com/RazvanBerbece/UnifyFootballBot/internal/globals"
)

type Team struct {
	Id          int
	Name        string
	DisplayName string
	LogoUrl     string
	LogoBase64  string
}

type League struct {
	Id          int
	Name        string
	CountryName string
	Teams       []Team
}

func GetLeaguesForCountry(countryName string, leagues int) []League {

	var retLeagues []League

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/leagues?country=%s", countryName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("An error occured while creating request to retrieve leagues for country %s: %v", countryName, err)
	}

	req.Header.Add("X-RapidAPI-Key", globals.RapidApiFootballKey)
	req.Header.Add("X-RapidAPI-Host", globals.RapidApiHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("An error occured while retrieving leagues for country %s: %v", countryName, err)
	}
	defer res.Body.Close()

	var leaguesResponse apiFootballModels.ApiFootballLeaguesResponse
	err = json.NewDecoder(res.Body).Decode(&leaguesResponse)
	if err != nil {
		fmt.Printf("An error occured while decoding leagues for country response: %v", err)
	}

	// Construct first `league` Leagues for given country
	var i int
	for i = 0; i < leagues; i++ {
		retLeagues = append(retLeagues, League{
			Id:          leaguesResponse.Response[i].League.ID,
			Name:        leaguesResponse.Response[i].League.Name,
			CountryName: countryName,
		})
	}

	return retLeagues

}

func GetTeamsForLeague(leagueId int, season int, countryName string) []Team {

	var retTeams []Team

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/teams?league=%d&season=%d&country=%s", leagueId, season, countryName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("An error occured while creating request to retrieve teams for league %d: %v", leagueId, err)
	}

	req.Header.Add("X-RapidAPI-Key", globals.RapidApiFootballKey)
	req.Header.Add("X-RapidAPI-Host", globals.RapidApiHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("An error occured while retrieving teams for league %d: %v", leagueId, err)
	}
	defer res.Body.Close()

	var teamsResponse apiFootballModels.ApiFootballTeamsForLeagueResponse
	err = json.NewDecoder(res.Body).Decode(&teamsResponse)
	if err != nil {
		fmt.Printf("An error occured while decoding teams for league response: %v", err)
	}

	// Construct the array of teams
	var i int
	for i = 0; i < len(teamsResponse.Response); i++ {
		// we transliterate to closest English representations because Discord does not allow non-English alphanumerical characters
		// as guild reaction names
		teamDisplayName := unidecode.Unidecode(teamsResponse.Response[i].Team.Name)
		// it seems that the Discord API needs emoji names to not have blank spaces, so we do this
		teamName := strings.Replace(teamDisplayName, " ", "_", -1)
		teamToAdd := Team{
			Id:          teamsResponse.Response[i].Team.ID,
			Name:        teamName,
			DisplayName: teamDisplayName,
			LogoUrl:     teamsResponse.Response[i].Team.Logo,
		}
		// Retrieve the image data for all team logos
		encodedString, err := GetImageAsBase64FromUrl(teamToAdd.LogoUrl)
		if err != nil {
			fmt.Printf("An error occured while downloading team logo image data: %v", err)
		}
		teamToAdd.LogoBase64 = encodedString
		retTeams = append(retTeams, teamToAdd)
	}

	return retTeams

}

func GetImageAsBase64FromUrl(url string) (string, error) {

	imageURL := url

	response, err := http.Get(imageURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return "", err
	}
	defer response.Body.Close()

	// Read the response body into a byte slice
	imageBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	// Encode the image data as a base64 string in the Discord format
	encodedImage := base64.StdEncoding.EncodeToString(imageBytes)
	encodedImage += strings.Repeat("=", (4-len(encodedImage)%4)%4)

	return encodedImage, nil

}
