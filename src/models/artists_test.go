// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testArtists(t *testing.T) {
	t.Parallel()

	query := Artists()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testArtistsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Artists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArtistsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Artists().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Artists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArtistsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ArtistSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Artists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArtistsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ArtistExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Artist exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ArtistExists to return true, but got false.")
	}
}

func testArtistsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	artistFound, err := FindArtist(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if artistFound == nil {
		t.Error("want a record, got nil")
	}
}

func testArtistsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Artists().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testArtistsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Artists().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testArtistsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	artistOne := &Artist{}
	artistTwo := &Artist{}
	if err = randomize.Struct(seed, artistOne, artistDBTypes, false, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}
	if err = randomize.Struct(seed, artistTwo, artistDBTypes, false, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = artistOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = artistTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Artists().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testArtistsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	artistOne := &Artist{}
	artistTwo := &Artist{}
	if err = randomize.Struct(seed, artistOne, artistDBTypes, false, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}
	if err = randomize.Struct(seed, artistTwo, artistDBTypes, false, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = artistOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = artistTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Artists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func artistBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Artist) error {
	*o = Artist{}
	return nil
}

func artistAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Artist) error {
	*o = Artist{}
	return nil
}

func artistAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Artist) error {
	*o = Artist{}
	return nil
}

func artistBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Artist) error {
	*o = Artist{}
	return nil
}

func artistAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Artist) error {
	*o = Artist{}
	return nil
}

func artistBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Artist) error {
	*o = Artist{}
	return nil
}

func artistAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Artist) error {
	*o = Artist{}
	return nil
}

func artistBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Artist) error {
	*o = Artist{}
	return nil
}

func artistAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Artist) error {
	*o = Artist{}
	return nil
}

func testArtistsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Artist{}
	o := &Artist{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, artistDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Artist object: %s", err)
	}

	AddArtistHook(boil.BeforeInsertHook, artistBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	artistBeforeInsertHooks = []ArtistHook{}

	AddArtistHook(boil.AfterInsertHook, artistAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	artistAfterInsertHooks = []ArtistHook{}

	AddArtistHook(boil.AfterSelectHook, artistAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	artistAfterSelectHooks = []ArtistHook{}

	AddArtistHook(boil.BeforeUpdateHook, artistBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	artistBeforeUpdateHooks = []ArtistHook{}

	AddArtistHook(boil.AfterUpdateHook, artistAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	artistAfterUpdateHooks = []ArtistHook{}

	AddArtistHook(boil.BeforeDeleteHook, artistBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	artistBeforeDeleteHooks = []ArtistHook{}

	AddArtistHook(boil.AfterDeleteHook, artistAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	artistAfterDeleteHooks = []ArtistHook{}

	AddArtistHook(boil.BeforeUpsertHook, artistBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	artistBeforeUpsertHooks = []ArtistHook{}

	AddArtistHook(boil.AfterUpsertHook, artistAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	artistAfterUpsertHooks = []ArtistHook{}
}

func testArtistsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Artists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testArtistsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(artistColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Artists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testArtistsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testArtistsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ArtistSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testArtistsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Artists().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	artistDBTypes = map[string]string{`ID`: `text`, `Name`: `text`, `SpotifyID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_             = bytes.MinRead
)

func testArtistsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(artistPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(artistAllColumns) == len(artistPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Artists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, artistDBTypes, true, artistPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testArtistsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(artistAllColumns) == len(artistPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Artist{}
	if err = randomize.Struct(seed, o, artistDBTypes, true, artistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Artists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, artistDBTypes, true, artistPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(artistAllColumns, artistPrimaryKeyColumns) {
		fields = artistAllColumns
	} else {
		fields = strmangle.SetComplement(
			artistAllColumns,
			artistPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ArtistSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testArtistsUpsert(t *testing.T) {
	t.Parallel()

	if len(artistAllColumns) == len(artistPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Artist{}
	if err = randomize.Struct(seed, &o, artistDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Artist: %s", err)
	}

	count, err := Artists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, artistDBTypes, false, artistPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Artist struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Artist: %s", err)
	}

	count, err = Artists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
