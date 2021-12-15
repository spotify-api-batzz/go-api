// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Thumbnail is an object representing the database table.
type Thumbnail struct {
	ID        string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Entity    null.String `boil:"entity" json:"entity,omitempty" toml:"entity" yaml:"entity,omitempty"`
	EntityID  null.String `boil:"entity_id" json:"entity_id,omitempty" toml:"entity_id" yaml:"entity_id,omitempty"`
	Width     null.Int16  `boil:"width" json:"width,omitempty" toml:"width" yaml:"width,omitempty"`
	Height    null.Int16  `boil:"height" json:"height,omitempty" toml:"height" yaml:"height,omitempty"`
	URL       null.String `boil:"url" json:"url,omitempty" toml:"url" yaml:"url,omitempty"`
	CreatedAt null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *thumbnailR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L thumbnailL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ThumbnailColumns = struct {
	ID        string
	Entity    string
	EntityID  string
	Width     string
	Height    string
	URL       string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Entity:    "entity",
	EntityID:  "entity_id",
	Width:     "width",
	Height:    "height",
	URL:       "url",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

type whereHelpernull_Int16 struct{ field string }

func (w whereHelpernull_Int16) EQ(x null.Int16) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int16) NEQ(x null.Int16) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int16) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int16) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int16) LT(x null.Int16) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int16) LTE(x null.Int16) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int16) GT(x null.Int16) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int16) GTE(x null.Int16) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ThumbnailWhere = struct {
	ID        whereHelperstring
	Entity    whereHelpernull_String
	EntityID  whereHelpernull_String
	Width     whereHelpernull_Int16
	Height    whereHelpernull_Int16
	URL       whereHelpernull_String
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperstring{field: "\"thumbnails\".\"id\""},
	Entity:    whereHelpernull_String{field: "\"thumbnails\".\"entity\""},
	EntityID:  whereHelpernull_String{field: "\"thumbnails\".\"entity_id\""},
	Width:     whereHelpernull_Int16{field: "\"thumbnails\".\"width\""},
	Height:    whereHelpernull_Int16{field: "\"thumbnails\".\"height\""},
	URL:       whereHelpernull_String{field: "\"thumbnails\".\"url\""},
	CreatedAt: whereHelpernull_Time{field: "\"thumbnails\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"thumbnails\".\"updated_at\""},
}

// ThumbnailRels is where relationship names are stored.
var ThumbnailRels = struct {
}{}

// thumbnailR is where relationships are stored.
type thumbnailR struct {
}

// NewStruct creates a new relationship struct
func (*thumbnailR) NewStruct() *thumbnailR {
	return &thumbnailR{}
}

// thumbnailL is where Load methods for each relationship are stored.
type thumbnailL struct{}

var (
	thumbnailAllColumns            = []string{"id", "entity", "entity_id", "width", "height", "url", "created_at", "updated_at"}
	thumbnailColumnsWithoutDefault = []string{"id", "entity", "entity_id", "width", "height", "url", "created_at", "updated_at"}
	thumbnailColumnsWithDefault    = []string{}
	thumbnailPrimaryKeyColumns     = []string{"id"}
)

type (
	// ThumbnailSlice is an alias for a slice of pointers to Thumbnail.
	// This should generally be used opposed to []Thumbnail.
	ThumbnailSlice []*Thumbnail
	// ThumbnailHook is the signature for custom Thumbnail hook methods
	ThumbnailHook func(context.Context, boil.ContextExecutor, *Thumbnail) error

	thumbnailQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	thumbnailType                 = reflect.TypeOf(&Thumbnail{})
	thumbnailMapping              = queries.MakeStructMapping(thumbnailType)
	thumbnailPrimaryKeyMapping, _ = queries.BindMapping(thumbnailType, thumbnailMapping, thumbnailPrimaryKeyColumns)
	thumbnailInsertCacheMut       sync.RWMutex
	thumbnailInsertCache          = make(map[string]insertCache)
	thumbnailUpdateCacheMut       sync.RWMutex
	thumbnailUpdateCache          = make(map[string]updateCache)
	thumbnailUpsertCacheMut       sync.RWMutex
	thumbnailUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var thumbnailBeforeInsertHooks []ThumbnailHook
var thumbnailBeforeUpdateHooks []ThumbnailHook
var thumbnailBeforeDeleteHooks []ThumbnailHook
var thumbnailBeforeUpsertHooks []ThumbnailHook

var thumbnailAfterInsertHooks []ThumbnailHook
var thumbnailAfterSelectHooks []ThumbnailHook
var thumbnailAfterUpdateHooks []ThumbnailHook
var thumbnailAfterDeleteHooks []ThumbnailHook
var thumbnailAfterUpsertHooks []ThumbnailHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Thumbnail) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thumbnailBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Thumbnail) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thumbnailBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Thumbnail) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thumbnailBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Thumbnail) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thumbnailBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Thumbnail) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thumbnailAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Thumbnail) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thumbnailAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Thumbnail) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thumbnailAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Thumbnail) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thumbnailAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Thumbnail) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thumbnailAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddThumbnailHook registers your hook function for all future operations.
func AddThumbnailHook(hookPoint boil.HookPoint, thumbnailHook ThumbnailHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		thumbnailBeforeInsertHooks = append(thumbnailBeforeInsertHooks, thumbnailHook)
	case boil.BeforeUpdateHook:
		thumbnailBeforeUpdateHooks = append(thumbnailBeforeUpdateHooks, thumbnailHook)
	case boil.BeforeDeleteHook:
		thumbnailBeforeDeleteHooks = append(thumbnailBeforeDeleteHooks, thumbnailHook)
	case boil.BeforeUpsertHook:
		thumbnailBeforeUpsertHooks = append(thumbnailBeforeUpsertHooks, thumbnailHook)
	case boil.AfterInsertHook:
		thumbnailAfterInsertHooks = append(thumbnailAfterInsertHooks, thumbnailHook)
	case boil.AfterSelectHook:
		thumbnailAfterSelectHooks = append(thumbnailAfterSelectHooks, thumbnailHook)
	case boil.AfterUpdateHook:
		thumbnailAfterUpdateHooks = append(thumbnailAfterUpdateHooks, thumbnailHook)
	case boil.AfterDeleteHook:
		thumbnailAfterDeleteHooks = append(thumbnailAfterDeleteHooks, thumbnailHook)
	case boil.AfterUpsertHook:
		thumbnailAfterUpsertHooks = append(thumbnailAfterUpsertHooks, thumbnailHook)
	}
}

// One returns a single thumbnail record from the query.
func (q thumbnailQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Thumbnail, error) {
	o := &Thumbnail{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for thumbnails")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Thumbnail records from the query.
func (q thumbnailQuery) All(ctx context.Context, exec boil.ContextExecutor) (ThumbnailSlice, error) {
	var o []*Thumbnail

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Thumbnail slice")
	}

	if len(thumbnailAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Thumbnail records in the query.
func (q thumbnailQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count thumbnails rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q thumbnailQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if thumbnails exists")
	}

	return count > 0, nil
}

// Thumbnails retrieves all the records using an executor.
func Thumbnails(mods ...qm.QueryMod) thumbnailQuery {
	mods = append(mods, qm.From("\"thumbnails\""))
	return thumbnailQuery{NewQuery(mods...)}
}

// FindThumbnail retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindThumbnail(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Thumbnail, error) {
	thumbnailObj := &Thumbnail{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"thumbnails\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, thumbnailObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from thumbnails")
	}

	return thumbnailObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Thumbnail) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no thumbnails provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(thumbnailColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	thumbnailInsertCacheMut.RLock()
	cache, cached := thumbnailInsertCache[key]
	thumbnailInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			thumbnailAllColumns,
			thumbnailColumnsWithDefault,
			thumbnailColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(thumbnailType, thumbnailMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(thumbnailType, thumbnailMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"thumbnails\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"thumbnails\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into thumbnails")
	}

	if !cached {
		thumbnailInsertCacheMut.Lock()
		thumbnailInsertCache[key] = cache
		thumbnailInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Thumbnail.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Thumbnail) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	thumbnailUpdateCacheMut.RLock()
	cache, cached := thumbnailUpdateCache[key]
	thumbnailUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			thumbnailAllColumns,
			thumbnailPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update thumbnails, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"thumbnails\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, thumbnailPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(thumbnailType, thumbnailMapping, append(wl, thumbnailPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update thumbnails row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for thumbnails")
	}

	if !cached {
		thumbnailUpdateCacheMut.Lock()
		thumbnailUpdateCache[key] = cache
		thumbnailUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q thumbnailQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for thumbnails")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for thumbnails")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ThumbnailSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thumbnailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"thumbnails\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, thumbnailPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in thumbnail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all thumbnail")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Thumbnail) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no thumbnails provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(thumbnailColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	thumbnailUpsertCacheMut.RLock()
	cache, cached := thumbnailUpsertCache[key]
	thumbnailUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			thumbnailAllColumns,
			thumbnailColumnsWithDefault,
			thumbnailColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			thumbnailAllColumns,
			thumbnailPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert thumbnails, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(thumbnailPrimaryKeyColumns))
			copy(conflict, thumbnailPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"thumbnails\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(thumbnailType, thumbnailMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(thumbnailType, thumbnailMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert thumbnails")
	}

	if !cached {
		thumbnailUpsertCacheMut.Lock()
		thumbnailUpsertCache[key] = cache
		thumbnailUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Thumbnail record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Thumbnail) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Thumbnail provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), thumbnailPrimaryKeyMapping)
	sql := "DELETE FROM \"thumbnails\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from thumbnails")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for thumbnails")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q thumbnailQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no thumbnailQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from thumbnails")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for thumbnails")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ThumbnailSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(thumbnailBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thumbnailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"thumbnails\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thumbnailPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from thumbnail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for thumbnails")
	}

	if len(thumbnailAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Thumbnail) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindThumbnail(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ThumbnailSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ThumbnailSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thumbnailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"thumbnails\".* FROM \"thumbnails\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thumbnailPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ThumbnailSlice")
	}

	*o = slice

	return nil
}

// ThumbnailExists checks if the Thumbnail row exists.
func ThumbnailExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"thumbnails\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if thumbnails exists")
	}

	return exists, nil
}