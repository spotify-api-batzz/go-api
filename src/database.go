package main

import (
	"fmt"
	"spotify-server/models"
	"spotify-server/utils"
	"strings"

	"github.com/batzz-00/goutils/logger"

	"github.com/batzz-00/goutils"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	DB *sqlx.DB
}

// Connect opens up a conection to the database
func (database *Database) Connect() error {
	dbUser := utils.MustGetEnv("DB_USER")
	dbIP := utils.MustGetEnv("DB_IP")
	dbPass := utils.MustGetEnv("DB_PASS")
	dbPort := utils.MustGetEnv("DB_PORT")
	dbTable := utils.MustGetEnv("DB_TABLE")
	url := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", dbIP, dbPort, dbTable, dbUser, dbPass)

	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		logger.Log(err, logger.Error)
		// do something here
		return err
	}

	database.DB = db
	return nil
}

func (d *Database) FetchUsersByNames(names []interface{}) ([]models.User, error) {
	user := []models.User{}
	sql := fmt.Sprintf("SELECT * FROM users WHERE username IN (%s)", PrepareInStringPG(1, len(names), 1))
	err := d.DB.Select(&user, sql, names...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *Database) FetchUserByName(name string) (models.User, error) {
	user := models.User{}
	err := d.DB.Get(&user, "SELECT * FROM users WHERE username = $1", name)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (d *Database) FetchArtistBySpotifyID(spotifyID string) (models.Artist, error) {
	artist := models.Artist{}
	err := d.DB.Get(&artist, "SELECT * FROM artists WHERE spotify_id = $1", spotifyID)
	if err != nil {
		return models.Artist{}, err
	}
	return artist, nil
}

func (d *Database) FetchSongsBySpotifyID(spotifyIDs []interface{}) ([]models.Song, error) {
	songs := []models.Song{}
	sql := fmt.Sprintf("SELECT * FROM songs WHERE spotify_id IN (%s)", PrepareBatchValuesPG(1, len(spotifyIDs)))
	err := d.DB.Select(&songs, sql, spotifyIDs...)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (d *Database) FetchSongs(queryOptions *QueryOptions) ([]models.Song, error) {
	songs := []models.Song{}
	sql := Transform("SELECT * FROM songs", queryOptions)
	err := d.DB.Select(&songs, sql)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (d *Database) FetchSongsByIDFiltered(songIDs []interface{}, queryOptions *QueryOptions) ([]models.Song, error) {
	songs := []models.Song{}
	columnNames := goutils.ColumnNamesExclusive(&models.Song{})
	tableName := (&models.Song{}).TableName()
	sql := Transform(fmt.Sprintf("SELECT %s FROM %s WHERE id IN (%s)", columnNames, tableName, PrepareBatchValuesPG(1, len(songIDs))), queryOptions)
	err := d.DB.Select(&songs, sql, songIDs...)
	if err != nil {
		return nil, err
	}
	return songs, nil
}
func (d *Database) FetchSongsByID(songIDs []interface{}) ([]models.Song, error) {
	songs := []models.Song{}
	columnNames := goutils.ColumnNamesExclusive(&models.Song{})
	tableName := (&models.Song{}).TableName()
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE id IN (%s)", columnNames, tableName, PrepareBatchValuesPG(1, len(songIDs)))
	err := d.DB.Select(&songs, sql, songIDs...)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (d *Database) FetchTopSongsByUserID(userID string, queryOptions *QueryOptions) ([]models.TopSong, error) {
	topSongs := []models.TopSong{}
	columnNames := goutils.ColumnNamesExclusive(&models.TopSong{})
	tableName := (&models.TopSong{}).TableName()
	sql := Transform(fmt.Sprintf("SELECT %s FROM %s WHERE user_id = $1", columnNames, tableName), queryOptions)
	err := d.DB.Select(&topSongs, sql, userID)
	if err != nil {
		return nil, err
	}
	return topSongs, nil
}

func (d *Database) FetchTopSongsByID(topSongID string, queryOptions *QueryOptions) ([]models.TopSong, error) {
	topSongs := []models.TopSong{}
	columnNames := goutils.ColumnNamesExclusive(&models.TopSong{})
	tableName := (&models.TopSong{}).TableName()
	sql := Transform(fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", columnNames, tableName), queryOptions)
	err := d.DB.Select(&topSongs, sql, topSongID)
	if err != nil {
		return nil, err
	}
	return topSongs, nil
}

func (d *Database) FetchTopSongDataByTopSongIDs(topSongIDs []interface{}, queryOptions *QueryOptions) ([]models.TopSongData, error) {
	topSongData := []models.TopSongData{}
	sql := Transform(fmt.Sprintf("SELECT * FROM top_song_data WHERE top_song_id IN (%s)", PrepareBatchValuesPG(1, len(topSongIDs))), queryOptions)
	err := d.DB.Select(&topSongData, sql, topSongIDs...)
	if err != nil {
		return nil, err
	}
	return topSongData, nil
}

func (d *Database) FetchAlbumsBySpotifyID(spotifyIDs []interface{}, queryOptions *QueryOptions) ([]models.Album, error) {
	albums := []models.Album{}
	sql := Transform(fmt.Sprintf("SELECT * FROM albums WHERE spotify_id IN (%s)", PrepareBatchValuesPG(1, len(spotifyIDs))), queryOptions)
	err := d.DB.Select(&albums, sql, spotifyIDs...)
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (d *Database) FetchAlbums(queryOptions *QueryOptions) ([]models.Album, error) {
	albums := []models.Album{}
	sql := Transform("SELECT * FROM albums", queryOptions)
	err := d.DB.Select(&albums, sql)
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (d *Database) FetchAlbumByID(id string) (models.Album, error) {
	album := models.Album{}
	err := d.DB.Get(&album, "SELECT * FROM albums WHERE id = $1", id)
	if err != nil {
		return models.Album{}, err
	}
	return album, nil
}

func (d *Database) FetchAlbumsByID(ids []interface{}) ([]models.Album, error) {
	album := []models.Album{}
	sql := fmt.Sprintf("SELECT * FROM albums WHERE id IN (%s)", PrepareBatchValuesPG(1, len(ids)))
	err := d.DB.Select(&album, sql, ids...)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (d *Database) FetchArtistsBySpotifyID(spotifyIDs []interface{}) ([]models.Artist, error) {
	artists := []models.Artist{}
	sql := fmt.Sprintf("SELECT * FROM artists WHERE spotify_id IN (%s)", PrepareBatchValuesPG(1, len(spotifyIDs)))
	err := d.DB.Select(&artists, sql, spotifyIDs...)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func (d *Database) FetchArtistByID(id string) (models.Artist, error) {
	artist := models.Artist{}
	err := d.DB.Get(&artist, "SELECT * FROM artists WHERE id = $1", id)
	if err != nil {
		return models.Artist{}, err
	}
	return artist, nil
}

func (d *Database) FetchArtistsByID(ids []interface{}) ([]models.Artist, error) {
	artist := []models.Artist{}
	sql := fmt.Sprintf("SELECT * FROM artists WHERE id IN (%s)", PrepareBatchValuesPG(1, len(ids)))
	err := d.DB.Select(&artist, sql, ids...)
	if err != nil {
		return nil, err
	}
	return artist, nil
}

func (d *Database) FetchRecentListensByTime(playedAts []interface{}, userID string) ([]models.RecentListen, error) {
	recentListens := []models.RecentListen{}
	sql := fmt.Sprintf("SELECT * FROM recent_listens WHERE user_id = $1 AND played_at IN (%s)", PrepareInStringPG(1, len(playedAts), 2))
	vars := []interface{}{}
	vars = append(vars, userID)
	vars = append(vars, playedAts...)
	err := d.DB.Select(&recentListens, sql, vars...)
	if err != nil {
		return nil, err
	}
	return recentListens, nil
}

func (d *Database) FetchRecentListensByUser(userID string) ([]models.RecentListen, error) {
	recentListens := []models.RecentListen{}
	sql := "SELECT * FROM recent_listens WHERE user_id = $1"
	err := d.DB.Select(&recentListens, sql, userID)
	if err != nil {
		return nil, err
	}
	return recentListens, nil
}

func (d *Database) FetchRecentListenDataByRecentListenID(recentListenIDs []interface{}, opts *QueryOptions) ([]models.RecentListenData, error) {
	recentListenData := []models.RecentListenData{}
	sql := Transform(fmt.Sprintf("SELECT * FROM recent_listen_data WHERE recent_listen_id IN(%s)", PrepareInStringPG(1, len(recentListenIDs), 1)), opts)
	err := d.DB.Select(&recentListenData, sql, recentListenIDs...)
	if err != nil {
		return nil, err
	}
	return recentListenData, nil
}

func (d *Database) FetchThumbnailsByEntityID(entityIDs []interface{}) ([]models.Thumbnail, error) {
	thumbnails := []models.Thumbnail{}
	sql := fmt.Sprintf("SELECT * FROM thumbnails WHERE entity_id IN (%s)", PrepareInStringPG(1, len(entityIDs), 1))
	err := d.DB.Select(&thumbnails, sql, entityIDs...)
	if err != nil {
		return nil, err
	}
	return thumbnails, nil
}

func PrepareBatchValuesPG(paramLength int, valueLength int) string {
	counter := 1
	var values string
	for i := 0; i < valueLength; i++ {
		values = fmt.Sprintf("%s, %s", values, genValString(paramLength, &counter))
	}
	return strings.TrimPrefix(values, ", ")
}

func PrepareInStringPG(paramLength int, valueLength int, counter int) string {
	if counter == 0 {
		counter = 1
	}
	var values string
	for i := 0; i < valueLength; i++ {
		values = fmt.Sprintf("%s, %s", values, genValString(paramLength, &counter))
	}
	return strings.TrimPrefix(values, ", ")
}

func genValString(paramLength int, counter *int) string {
	var valString string
	for i := 0; i < paramLength; i++ {
		valString = valString + fmt.Sprintf("$%d,", *counter)
		*counter++
	}
	valString = fmt.Sprintf("(%s)", strings.TrimSuffix(valString, ","))
	return valString
}
