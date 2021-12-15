package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"spotify-server/models"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Database *Database
}

type Response struct {
	Message string        `json:"message"`
	Data    interface{}   `json:"data"`
	Status  bool          `json:"status"`
	Opts    *QueryOptions `json:"meta,omitempty"`
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func (server *Server) Serve() {
	fmt.Println("starting server")
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	// r.HandleFunc("/albums/{id}", server.HandleAlbum)
	r.HandleFunc("/album", server.HandleAlbums)
	r.HandleFunc("/songs", server.HandleSongs)

	r.HandleFunc("/album/{id}", server.HandleAlbum)
	r.HandleFunc("/user/{userid}/songs", server.HandleUserTopSongs)
	r.HandleFunc("/user/{userid}/recents", server.HandleUserRecents)
	// r.HandleFunc("/user/{userid}/artists", server.HandleSongs)
	r.HandleFunc("/topsongs/{topsongid}", server.HandleTopSongs)
	r.Use(loggingMiddleware)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func (server *Server) HandleAlbums(response http.ResponseWriter, req *http.Request) {
	opts, err := parseQueryOptions(req.URL.Query())
	if err != nil {
		handleErr(response, err)
		return
	}

	albums, err := server.Database.FetchAlbums(opts)
	if err != nil {
		handleErr(response, err)
		return
	}

	if len(albums) == 0 {
		json.NewEncoder(response).Encode(&Response{Message: "Failed to fetch albums", Status: false, Opts: opts})
		return
	}

	artistIDs := []interface{}{}
	for _, album := range albums {
		artistIDs = append(artistIDs, album.ArtistID)
	}

	artists, err := server.Database.FetchArtistsByID(artistIDs)
	if err != nil {
		handleErr(response, err)
		return
	}

	albums = models.AttachArtistsToAlbums(albums, artists)

	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(200)

	json.NewEncoder(response).Encode(&Response{Message: "Succesfully fetched albums", Data: albums, Status: true, Opts: opts})
}

func (server *Server) HandleAlbum(response http.ResponseWriter, req *http.Request) {
	opts, err := parseQueryOptions(req.URL.Query())
	if err != nil {
		handleErr(response, err)
		return
	}

	vars := mux.Vars(req)
	var albumID string
	if id, ok := vars["id"]; ok {
		albumID = id
	}

	album, err := server.Database.FetchAlbumByID(albumID)
	if err != nil {
		handleErr(response, err)
		return
	}

	artist, err := server.Database.FetchArtistByID(album.ArtistID)
	if err != nil {
		handleErr(response, err)
		return
	}

	album.Artist = &artist

	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(200)

	json.NewEncoder(response).Encode(&Response{Message: fmt.Sprintf("Succesfully fetched album with id %s", albumID), Data: album, Status: true, Opts: opts})
}

func (server *Server) HandleSongs(response http.ResponseWriter, req *http.Request) {
	start := time.Now()
	opts, err := parseQueryOptions(req.URL.Query())
	if err != nil {
		handleErr(response, err)
		return
	}

	songs, err := server.Database.FetchSongs(opts)
	if err != nil {
		handleErr(response, err)
		return
	}

	if len(songs) == 0 {
		json.NewEncoder(response).Encode(&Response{Message: "Failed to fetch songs", Status: false, Opts: opts})
		return
	}

	artistIDs := []interface{}{}
	albumIDs := []interface{}{}
	for _, song := range songs {
		artistIDs = append(artistIDs, song.ArtistID)
		albumIDs = append(albumIDs, song.AlbumID)
	}

	artists, err := server.Database.FetchArtistsByID(artistIDs)
	if err != nil {
		handleErr(response, err)
		return
	}

	albums, err := server.Database.FetchAlbumsByID(albumIDs)
	if err != nil {
		handleErr(response, err)
		return
	}

	songs = models.AttachArtistsToSongs(songs, artists)
	songs = models.AttachAlbumsToSongs(songs, albums)

	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(200)

	elapsed := time.Since(start)
	log.Printf("Fetched songs in %s", elapsed)

	json.NewEncoder(response).Encode(&Response{Message: "Succesfully fetched songs", Data: songs, Status: true, Opts: opts})
}

func (server *Server) HandleUserTopSongs(response http.ResponseWriter, req *http.Request) {
	opts, err := parseQueryOptions(req.URL.Query())
	if err != nil {
		handleErr(response, err)
		return
	}

	vars := mux.Vars(req)
	topSongs, err := server.Database.FetchTopSongsByUserID(vars["userid"], opts)
	if err != nil {
		handleErr(response, err)
		return
	}

	if len(topSongs) == 0 {
		json.NewEncoder(response).Encode(&Response{Message: "No top song records for user", Status: false, Opts: opts})
		return
	}

	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(200)

	json.NewEncoder(response).Encode(&Response{Message: "Succesfully fetched top songs", Data: topSongs, Status: true, Opts: opts})
}

func (server *Server) HandleUserRecents(response http.ResponseWriter, req *http.Request) {
	opts, err := parseQueryOptions(req.URL.Query())
	if err != nil {
		handleErr(response, err)
		return
	}

	vars := mux.Vars(req)
	recentListens, err := server.Database.FetchRecentListensByUser(vars["userid"])
	if err != nil {
		handleErr(response, err)
		return
	}

	if len(recentListens) == 0 {
		json.NewEncoder(response).Encode(&Response{Message: "No recently listened data for user", Status: false, Opts: opts})
		return
	}

	recentListenIDs := []interface{}{}
	for _, recentListen := range recentListens {
		recentListenIDs = append(recentListenIDs, recentListen.ID.String())
	}

	recentListenData, err := server.Database.FetchRecentListenDataByRecentListenID(recentListenIDs, opts)
	if err != nil {
		handleErr(response, err)
		return
	}

	songIDs := []interface{}{}
	for _, recentListenData := range recentListenData {
		songIDs = append(songIDs, recentListenData.SongID)
	}

	songs, err := server.Database.FetchSongsByID(songIDs)
	if err != nil {
		handleErr(response, err)
		return
	}

	artistIDs := []interface{}{}
	for _, song := range songs {
		artistIDs = append(artistIDs, song.ArtistID)
	}

	artists, err := server.Database.FetchArtistsByID(artistIDs)
	if err != nil {
		handleErr(response, err)
		return
	}

	songs = models.AttachArtistsToSongs(songs, artists)
	recentListenData = models.AttachSongsToRecentListenData(songs, recentListenData)

	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(200)

	json.NewEncoder(response).Encode(&Response{Message: "Succesfully fetched recently listened data", Data: recentListenData, Status: true, Opts: opts})

}

func (server *Server) HandleTopSongs(response http.ResponseWriter, req *http.Request) {
	opts, err := parseQueryOptions(req.URL.Query())
	if err != nil {
		handleErr(response, err)
		return
	}

	vars := mux.Vars(req)
	topSongs, err := server.Database.FetchTopSongsByID(vars["topsongid"], opts)
	if err != nil {
		handleErr(response, err)
		return
	}

	if len(topSongs) == 0 {
		json.NewEncoder(response).Encode(&Response{Message: "No top song records for user", Status: false, Opts: opts})
		return
	}

	topSongIDs := []interface{}{}
	for _, topSong := range topSongs {
		topSongIDs = append(topSongIDs, topSong.ID.String())
	}

	topSongData, err := server.Database.FetchTopSongDataByTopSongIDs(topSongIDs, opts)
	if err != nil {
		handleErr(response, err)
		return
	}

	songIDs := []interface{}{}
	for _, songData := range topSongData {
		songIDs = append(songIDs, songData.SongID)
	}

	songs, err := server.Database.FetchSongsByIDFiltered(songIDs, opts)
	if err != nil {
		handleErr(response, err)
		return
	}

	topSongData = models.AttachSongsToTopSongData(topSongData, songs)
	topSongs = models.AttachTopSongDataToTopSongs(topSongs, topSongData)

	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(200)

	json.NewEncoder(response).Encode(&Response{Message: "Succesfully fetched top song data", Data: topSongs, Status: true, Opts: opts})
}

func handleErr(response http.ResponseWriter, err error) {
	log.Print(err.Error())
	response.WriteHeader(500)
	json.NewEncoder(response).Encode(&Response{Message: "Server error"})
}
