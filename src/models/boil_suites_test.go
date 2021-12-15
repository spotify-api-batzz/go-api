// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Albums", testAlbums)
	t.Run("Artists", testArtists)
	t.Run("RecentListenData", testRecentListenData)
	t.Run("RecentListens", testRecentListens)
	t.Run("Songs", testSongs)
	t.Run("Thumbnails", testThumbnails)
	t.Run("TopArtistData", testTopArtistData)
	t.Run("TopArtists", testTopArtists)
	t.Run("TopSongData", testTopSongData)
	t.Run("TopSongs", testTopSongs)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Albums", testAlbumsDelete)
	t.Run("Artists", testArtistsDelete)
	t.Run("RecentListenData", testRecentListenDataDelete)
	t.Run("RecentListens", testRecentListensDelete)
	t.Run("Songs", testSongsDelete)
	t.Run("Thumbnails", testThumbnailsDelete)
	t.Run("TopArtistData", testTopArtistDataDelete)
	t.Run("TopArtists", testTopArtistsDelete)
	t.Run("TopSongData", testTopSongDataDelete)
	t.Run("TopSongs", testTopSongsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Albums", testAlbumsQueryDeleteAll)
	t.Run("Artists", testArtistsQueryDeleteAll)
	t.Run("RecentListenData", testRecentListenDataQueryDeleteAll)
	t.Run("RecentListens", testRecentListensQueryDeleteAll)
	t.Run("Songs", testSongsQueryDeleteAll)
	t.Run("Thumbnails", testThumbnailsQueryDeleteAll)
	t.Run("TopArtistData", testTopArtistDataQueryDeleteAll)
	t.Run("TopArtists", testTopArtistsQueryDeleteAll)
	t.Run("TopSongData", testTopSongDataQueryDeleteAll)
	t.Run("TopSongs", testTopSongsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Albums", testAlbumsSliceDeleteAll)
	t.Run("Artists", testArtistsSliceDeleteAll)
	t.Run("RecentListenData", testRecentListenDataSliceDeleteAll)
	t.Run("RecentListens", testRecentListensSliceDeleteAll)
	t.Run("Songs", testSongsSliceDeleteAll)
	t.Run("Thumbnails", testThumbnailsSliceDeleteAll)
	t.Run("TopArtistData", testTopArtistDataSliceDeleteAll)
	t.Run("TopArtists", testTopArtistsSliceDeleteAll)
	t.Run("TopSongData", testTopSongDataSliceDeleteAll)
	t.Run("TopSongs", testTopSongsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Albums", testAlbumsExists)
	t.Run("Artists", testArtistsExists)
	t.Run("RecentListenData", testRecentListenDataExists)
	t.Run("RecentListens", testRecentListensExists)
	t.Run("Songs", testSongsExists)
	t.Run("Thumbnails", testThumbnailsExists)
	t.Run("TopArtistData", testTopArtistDataExists)
	t.Run("TopArtists", testTopArtistsExists)
	t.Run("TopSongData", testTopSongDataExists)
	t.Run("TopSongs", testTopSongsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Albums", testAlbumsFind)
	t.Run("Artists", testArtistsFind)
	t.Run("RecentListenData", testRecentListenDataFind)
	t.Run("RecentListens", testRecentListensFind)
	t.Run("Songs", testSongsFind)
	t.Run("Thumbnails", testThumbnailsFind)
	t.Run("TopArtistData", testTopArtistDataFind)
	t.Run("TopArtists", testTopArtistsFind)
	t.Run("TopSongData", testTopSongDataFind)
	t.Run("TopSongs", testTopSongsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Albums", testAlbumsBind)
	t.Run("Artists", testArtistsBind)
	t.Run("RecentListenData", testRecentListenDataBind)
	t.Run("RecentListens", testRecentListensBind)
	t.Run("Songs", testSongsBind)
	t.Run("Thumbnails", testThumbnailsBind)
	t.Run("TopArtistData", testTopArtistDataBind)
	t.Run("TopArtists", testTopArtistsBind)
	t.Run("TopSongData", testTopSongDataBind)
	t.Run("TopSongs", testTopSongsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Albums", testAlbumsOne)
	t.Run("Artists", testArtistsOne)
	t.Run("RecentListenData", testRecentListenDataOne)
	t.Run("RecentListens", testRecentListensOne)
	t.Run("Songs", testSongsOne)
	t.Run("Thumbnails", testThumbnailsOne)
	t.Run("TopArtistData", testTopArtistDataOne)
	t.Run("TopArtists", testTopArtistsOne)
	t.Run("TopSongData", testTopSongDataOne)
	t.Run("TopSongs", testTopSongsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Albums", testAlbumsAll)
	t.Run("Artists", testArtistsAll)
	t.Run("RecentListenData", testRecentListenDataAll)
	t.Run("RecentListens", testRecentListensAll)
	t.Run("Songs", testSongsAll)
	t.Run("Thumbnails", testThumbnailsAll)
	t.Run("TopArtistData", testTopArtistDataAll)
	t.Run("TopArtists", testTopArtistsAll)
	t.Run("TopSongData", testTopSongDataAll)
	t.Run("TopSongs", testTopSongsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Albums", testAlbumsCount)
	t.Run("Artists", testArtistsCount)
	t.Run("RecentListenData", testRecentListenDataCount)
	t.Run("RecentListens", testRecentListensCount)
	t.Run("Songs", testSongsCount)
	t.Run("Thumbnails", testThumbnailsCount)
	t.Run("TopArtistData", testTopArtistDataCount)
	t.Run("TopArtists", testTopArtistsCount)
	t.Run("TopSongData", testTopSongDataCount)
	t.Run("TopSongs", testTopSongsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Albums", testAlbumsHooks)
	t.Run("Artists", testArtistsHooks)
	t.Run("RecentListenData", testRecentListenDataHooks)
	t.Run("RecentListens", testRecentListensHooks)
	t.Run("Songs", testSongsHooks)
	t.Run("Thumbnails", testThumbnailsHooks)
	t.Run("TopArtistData", testTopArtistDataHooks)
	t.Run("TopArtists", testTopArtistsHooks)
	t.Run("TopSongData", testTopSongDataHooks)
	t.Run("TopSongs", testTopSongsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Albums", testAlbumsInsert)
	t.Run("Albums", testAlbumsInsertWhitelist)
	t.Run("Artists", testArtistsInsert)
	t.Run("Artists", testArtistsInsertWhitelist)
	t.Run("RecentListenData", testRecentListenDataInsert)
	t.Run("RecentListenData", testRecentListenDataInsertWhitelist)
	t.Run("RecentListens", testRecentListensInsert)
	t.Run("RecentListens", testRecentListensInsertWhitelist)
	t.Run("Songs", testSongsInsert)
	t.Run("Songs", testSongsInsertWhitelist)
	t.Run("Thumbnails", testThumbnailsInsert)
	t.Run("Thumbnails", testThumbnailsInsertWhitelist)
	t.Run("TopArtistData", testTopArtistDataInsert)
	t.Run("TopArtistData", testTopArtistDataInsertWhitelist)
	t.Run("TopArtists", testTopArtistsInsert)
	t.Run("TopArtists", testTopArtistsInsertWhitelist)
	t.Run("TopSongData", testTopSongDataInsert)
	t.Run("TopSongData", testTopSongDataInsertWhitelist)
	t.Run("TopSongs", testTopSongsInsert)
	t.Run("TopSongs", testTopSongsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Albums", testAlbumsReload)
	t.Run("Artists", testArtistsReload)
	t.Run("RecentListenData", testRecentListenDataReload)
	t.Run("RecentListens", testRecentListensReload)
	t.Run("Songs", testSongsReload)
	t.Run("Thumbnails", testThumbnailsReload)
	t.Run("TopArtistData", testTopArtistDataReload)
	t.Run("TopArtists", testTopArtistsReload)
	t.Run("TopSongData", testTopSongDataReload)
	t.Run("TopSongs", testTopSongsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Albums", testAlbumsReloadAll)
	t.Run("Artists", testArtistsReloadAll)
	t.Run("RecentListenData", testRecentListenDataReloadAll)
	t.Run("RecentListens", testRecentListensReloadAll)
	t.Run("Songs", testSongsReloadAll)
	t.Run("Thumbnails", testThumbnailsReloadAll)
	t.Run("TopArtistData", testTopArtistDataReloadAll)
	t.Run("TopArtists", testTopArtistsReloadAll)
	t.Run("TopSongData", testTopSongDataReloadAll)
	t.Run("TopSongs", testTopSongsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Albums", testAlbumsSelect)
	t.Run("Artists", testArtistsSelect)
	t.Run("RecentListenData", testRecentListenDataSelect)
	t.Run("RecentListens", testRecentListensSelect)
	t.Run("Songs", testSongsSelect)
	t.Run("Thumbnails", testThumbnailsSelect)
	t.Run("TopArtistData", testTopArtistDataSelect)
	t.Run("TopArtists", testTopArtistsSelect)
	t.Run("TopSongData", testTopSongDataSelect)
	t.Run("TopSongs", testTopSongsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Albums", testAlbumsUpdate)
	t.Run("Artists", testArtistsUpdate)
	t.Run("RecentListenData", testRecentListenDataUpdate)
	t.Run("RecentListens", testRecentListensUpdate)
	t.Run("Songs", testSongsUpdate)
	t.Run("Thumbnails", testThumbnailsUpdate)
	t.Run("TopArtistData", testTopArtistDataUpdate)
	t.Run("TopArtists", testTopArtistsUpdate)
	t.Run("TopSongData", testTopSongDataUpdate)
	t.Run("TopSongs", testTopSongsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Albums", testAlbumsSliceUpdateAll)
	t.Run("Artists", testArtistsSliceUpdateAll)
	t.Run("RecentListenData", testRecentListenDataSliceUpdateAll)
	t.Run("RecentListens", testRecentListensSliceUpdateAll)
	t.Run("Songs", testSongsSliceUpdateAll)
	t.Run("Thumbnails", testThumbnailsSliceUpdateAll)
	t.Run("TopArtistData", testTopArtistDataSliceUpdateAll)
	t.Run("TopArtists", testTopArtistsSliceUpdateAll)
	t.Run("TopSongData", testTopSongDataSliceUpdateAll)
	t.Run("TopSongs", testTopSongsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}