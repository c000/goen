// Code generated by https://github.com/kamichidu/goen; DO NOT EDIT THIS FILE.
// Use of this source code is governed by a MIT license that can be found in
// the file located on https://github.com/kamichidu/goen repository.

package example

import (
	"database/sql"
	"time"

	"github.com/kamichidu/goen"
	"github.com/satori/go.uuid"
	"gopkg.in/Masterminds/squirrel.v1"
)

var metaSchema = new(goen.MetaSchema)

func init() {
	metaSchema.Register(Blog{})
}

type BlogSqlizer interface {
	squirrel.Sqlizer

	BlogToSql() (string, []interface{}, error)
}

type _BlogSqlizer struct {
	squirrel.Sqlizer
}

func (sqlizer *_BlogSqlizer) BlogToSql() (string, []interface{}, error) {
	return sqlizer.ToSql()
}

type BlogColumnExpr interface {
	BlogColumnExpr() string

	String() string
}

type BlogOrderExpr interface {
	BlogOrderExpr() string
}

type BlogQueryBuilder struct {
	dbc *goen.DBContext

	includeLoaders goen.IncludeLoaderList

	builder squirrel.SelectBuilder
}

func newBlogQueryBuilder(dbc *goen.DBContext) BlogQueryBuilder {
	stmtBuilder := squirrel.StatementBuilder.PlaceholderFormat(dbc.Dialect().PlaceholderFormat())
	metaT := metaSchema.LoadOf(&Blog{})
	return BlogQueryBuilder{
		dbc: dbc,
		// columns provided later
		builder: stmtBuilder.Select().From(metaT.TableName),
	}
}

func (qb BlogQueryBuilder) Include(loaders ...goen.IncludeLoader) BlogQueryBuilder {
	qb.includeLoaders.Append(loaders...)
	return qb
}

func (qb BlogQueryBuilder) Where(conds ...BlogSqlizer) BlogQueryBuilder {
	for _, cond := range conds {
		qb.builder = qb.builder.Where(cond)
	}
	return qb
}

func (qb BlogQueryBuilder) WhereRaw(conds ...squirrel.Sqlizer) BlogQueryBuilder {
	for _, cond := range conds {
		qb.builder = qb.builder.Where(cond)
	}
	return qb
}

func (qb BlogQueryBuilder) Offset(offset uint64) BlogQueryBuilder {
	qb.builder = qb.builder.Offset(offset)
	return qb
}

func (qb BlogQueryBuilder) Limit(limit uint64) BlogQueryBuilder {
	qb.builder = qb.builder.Limit(limit)
	return qb
}

func (qb BlogQueryBuilder) OrderBy(orderBys ...BlogOrderExpr) BlogQueryBuilder {
	exprs := make([]string, len(orderBys))
	for i := range orderBys {
		exprs[i] = orderBys[i].BlogOrderExpr()
	}
	qb.builder = qb.builder.OrderBy(exprs...)
	return qb
}

func (qb BlogQueryBuilder) Count() (int64, error) {
	query, args, err := qb.builder.Columns("count(*)").ToSql()
	if err != nil {
		return 0, err
	}

	var count int64
	row := qb.dbc.QueryRow(query, args...)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (qb BlogQueryBuilder) Query() ([]*Blog, error) {
	return qb.query()
}

func (qb BlogQueryBuilder) QueryRow() (*Blog, error) {
	qb.builder = qb.builder.Limit(1)
	if records, err := qb.query(); err != nil {
		return nil, err
	} else if len(records) == 0 {
		return nil, sql.ErrNoRows
	} else {
		return records[0], nil
	}
}

func (qb BlogQueryBuilder) query() ([]*Blog, error) {
	// for caching reason, wont support filtering columns
	metaT := metaSchema.LoadOf(&Blog{})
	cols := make([]string, len(metaT.Columns))
	for i := range metaT.Columns {
		cols[i] = metaT.Columns[i].ColumnName
	}

	query, args, err := qb.builder.Columns(cols...).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := qb.dbc.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var records []*Blog
	if err := qb.dbc.Scan(rows, &records); err != nil {
		rows.Close()
		return nil, err
	}
	rows.Close()

	sc := goen.NewScopeCache(metaSchema)
	for _, record := range records {
		sc.AddObject(record)
	}
	if err := qb.dbc.Include(records, sc, qb.includeLoaders); err != nil {
		return nil, err
	}

	return records, nil
}

type _Blog_BlogID_OrderExpr string

func (s _Blog_BlogID_OrderExpr) BlogOrderExpr() string {
	return string(s)
}

type _Blog_BlogID string

func (c _Blog_BlogID) BlogColumnExpr() string {
	return string(c)
}

func (c _Blog_BlogID) String() string {
	return string(c)
}

func (c _Blog_BlogID) Eq(v uuid.UUID) BlogSqlizer {
	return &_BlogSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Blog_BlogID) NotEq(v uuid.UUID) BlogSqlizer {
	return &_BlogSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Blog_BlogID) In(v []uuid.UUID) BlogSqlizer {
	return &_BlogSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Blog_BlogID) NotIn(v []uuid.UUID) BlogSqlizer {
	return &_BlogSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Blog_BlogID) Asc() BlogOrderExpr {
	return _Blog_BlogID_OrderExpr(string(c))
}

func (c _Blog_BlogID) Desc() BlogOrderExpr {
	return _Blog_BlogID_OrderExpr(string(c) + " DESC")
}

type _Blog_Name_OrderExpr string

func (s _Blog_Name_OrderExpr) BlogOrderExpr() string {
	return string(s)
}

type _Blog_Name string

func (c _Blog_Name) BlogColumnExpr() string {
	return string(c)
}

func (c _Blog_Name) String() string {
	return string(c)
}

func (c _Blog_Name) Eq(v string) BlogSqlizer {
	return &_BlogSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Blog_Name) NotEq(v string) BlogSqlizer {
	return &_BlogSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Blog_Name) In(v []string) BlogSqlizer {
	return &_BlogSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Blog_Name) NotIn(v []string) BlogSqlizer {
	return &_BlogSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Blog_Name) Like(v string) BlogSqlizer {
	return &_BlogSqlizer{squirrel.Expr(string(c)+" LIKE ?", v)}
}

func (c _Blog_Name) NotLike(v string) BlogSqlizer {
	return &_BlogSqlizer{squirrel.Expr(string(c)+" NOT LIKE ?", v)}
}

func (c _Blog_Name) Asc() BlogOrderExpr {
	return _Blog_Name_OrderExpr(string(c))
}

func (c _Blog_Name) Desc() BlogOrderExpr {
	return _Blog_Name_OrderExpr(string(c) + " DESC")
}

type _Blog_Author_OrderExpr string

func (s _Blog_Author_OrderExpr) BlogOrderExpr() string {
	return string(s)
}

type _Blog_Author string

func (c _Blog_Author) BlogColumnExpr() string {
	return string(c)
}

func (c _Blog_Author) String() string {
	return string(c)
}

func (c _Blog_Author) Eq(v string) BlogSqlizer {
	return &_BlogSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Blog_Author) NotEq(v string) BlogSqlizer {
	return &_BlogSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Blog_Author) In(v []string) BlogSqlizer {
	return &_BlogSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Blog_Author) NotIn(v []string) BlogSqlizer {
	return &_BlogSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Blog_Author) Like(v string) BlogSqlizer {
	return &_BlogSqlizer{squirrel.Expr(string(c)+" LIKE ?", v)}
}

func (c _Blog_Author) NotLike(v string) BlogSqlizer {
	return &_BlogSqlizer{squirrel.Expr(string(c)+" NOT LIKE ?", v)}
}

func (c _Blog_Author) Asc() BlogOrderExpr {
	return _Blog_Author_OrderExpr(string(c))
}

func (c _Blog_Author) Desc() BlogOrderExpr {
	return _Blog_Author_OrderExpr(string(c) + " DESC")
}

type BlogDBSet struct {
	dbc *goen.DBContext

	BlogID _Blog_BlogID

	Name _Blog_Name

	Author _Blog_Author

	IncludePosts goen.IncludeLoader
}

func newBlogDBSet(dbc *goen.DBContext) *BlogDBSet {
	dbset := &BlogDBSet{
		dbc: dbc,
	}
	dbset.BlogID = "blog_id"
	dbset.Name = "name"
	dbset.Author = "author"

	dbset.IncludePosts = goen.IncludeLoaderFunc(dbset.includePosts)

	return dbset
}

func (dbset *BlogDBSet) String() string {
	return "blogs"
}

func (dbset *BlogDBSet) Insert(v *Blog) {
	dbset.dbc.Patch(metaSchema.InsertPatchOf(v))
}

func (dbset *BlogDBSet) Select() BlogQueryBuilder {
	return newBlogQueryBuilder(dbset.dbc)
}

func (dbset *BlogDBSet) Update(v *Blog) {
	dbset.dbc.Patch(metaSchema.UpdatePatchOf(v))
}

func (dbset *BlogDBSet) Delete(v *Blog) {
	dbset.dbc.Patch(metaSchema.DeletePatchOf(v))
}

func (dbset *BlogDBSet) includePosts(later *goen.IncludeBuffer, sc *goen.ScopeCache, records interface{}) error {
	entities, ok := records.([]*Blog)
	if !ok {
		return nil
	}

	childRowKeyOf := func(v *Blog) goen.RowKey {
		return &goen.MapRowKey{
			Table: "posts",
			Key: map[string]interface{}{
				"blog_id": v.BlogID,
			},
		}
	}

	// filter cached entity
	cachedChildRowKeys := make([]goen.RowKey, 0, len(entities))
	noCachedChildRowKeys := make([]goen.RowKey, 0, len(entities))
	for _, entity := range entities {
		key := childRowKeyOf(entity)
		if sc.HasObject(key) {
			cachedChildRowKeys = append(cachedChildRowKeys, key)
		} else {
			noCachedChildRowKeys = append(noCachedChildRowKeys, key)
		}
	}
	if len(noCachedChildRowKeys) > 0 {
		cond := squirrel.Or{}
		for _, rowKey := range noCachedChildRowKeys {
			cond = append(cond, rowKey)
		}
		stmtBuilder := squirrel.StatementBuilder.PlaceholderFormat(dbset.dbc.Dialect().PlaceholderFormat())
		query, args, err := stmtBuilder.Select(
			"created_at",
			"updated_at",
			"blog_id",
			"post_id",
			"title",
			"content",
		).From("posts").Where(cond).ToSql()
		if err != nil {
			return err
		}
		rows, err := dbset.dbc.Query(query, args...)
		if err != nil {
			return err
		}

		var noCachedEntities []*Post
		if err := dbset.dbc.Scan(rows, &noCachedEntities); err != nil {
			rows.Close()
			return err
		}
		rows.Close()

		for _, entity := range noCachedEntities {
			sc.AddObject(entity)
		}

		// for newly loaded entity, to be filled by includeLoader
		later.AddRecords(noCachedEntities)
	}

	for _, entity := range entities {
		childRowKey := childRowKeyOf(entity)
		raw := sc.GetObject(childRowKey)
		if refes, ok := raw.([]interface{}); ok {
			for _, refe := range refes {
				entity.Posts = append(entity.Posts, refe.(*Post))
			}
		} else if raw != nil {
			entity.Posts = []*Post{raw.(*Post)}
		}
	}

	return nil
}

func init() {
	metaSchema.Register(Post{})
}

type PostSqlizer interface {
	squirrel.Sqlizer

	PostToSql() (string, []interface{}, error)
}

type _PostSqlizer struct {
	squirrel.Sqlizer
}

func (sqlizer *_PostSqlizer) PostToSql() (string, []interface{}, error) {
	return sqlizer.ToSql()
}

type PostColumnExpr interface {
	PostColumnExpr() string

	String() string
}

type PostOrderExpr interface {
	PostOrderExpr() string
}

type PostQueryBuilder struct {
	dbc *goen.DBContext

	includeLoaders goen.IncludeLoaderList

	builder squirrel.SelectBuilder
}

func newPostQueryBuilder(dbc *goen.DBContext) PostQueryBuilder {
	stmtBuilder := squirrel.StatementBuilder.PlaceholderFormat(dbc.Dialect().PlaceholderFormat())
	metaT := metaSchema.LoadOf(&Post{})
	return PostQueryBuilder{
		dbc: dbc,
		// columns provided later
		builder: stmtBuilder.Select().From(metaT.TableName),
	}
}

func (qb PostQueryBuilder) Include(loaders ...goen.IncludeLoader) PostQueryBuilder {
	qb.includeLoaders.Append(loaders...)
	return qb
}

func (qb PostQueryBuilder) Where(conds ...PostSqlizer) PostQueryBuilder {
	for _, cond := range conds {
		qb.builder = qb.builder.Where(cond)
	}
	return qb
}

func (qb PostQueryBuilder) WhereRaw(conds ...squirrel.Sqlizer) PostQueryBuilder {
	for _, cond := range conds {
		qb.builder = qb.builder.Where(cond)
	}
	return qb
}

func (qb PostQueryBuilder) Offset(offset uint64) PostQueryBuilder {
	qb.builder = qb.builder.Offset(offset)
	return qb
}

func (qb PostQueryBuilder) Limit(limit uint64) PostQueryBuilder {
	qb.builder = qb.builder.Limit(limit)
	return qb
}

func (qb PostQueryBuilder) OrderBy(orderBys ...PostOrderExpr) PostQueryBuilder {
	exprs := make([]string, len(orderBys))
	for i := range orderBys {
		exprs[i] = orderBys[i].PostOrderExpr()
	}
	qb.builder = qb.builder.OrderBy(exprs...)
	return qb
}

func (qb PostQueryBuilder) Count() (int64, error) {
	query, args, err := qb.builder.Columns("count(*)").ToSql()
	if err != nil {
		return 0, err
	}

	var count int64
	row := qb.dbc.QueryRow(query, args...)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (qb PostQueryBuilder) Query() ([]*Post, error) {
	return qb.query()
}

func (qb PostQueryBuilder) QueryRow() (*Post, error) {
	qb.builder = qb.builder.Limit(1)
	if records, err := qb.query(); err != nil {
		return nil, err
	} else if len(records) == 0 {
		return nil, sql.ErrNoRows
	} else {
		return records[0], nil
	}
}

func (qb PostQueryBuilder) query() ([]*Post, error) {
	// for caching reason, wont support filtering columns
	metaT := metaSchema.LoadOf(&Post{})
	cols := make([]string, len(metaT.Columns))
	for i := range metaT.Columns {
		cols[i] = metaT.Columns[i].ColumnName
	}

	query, args, err := qb.builder.Columns(cols...).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := qb.dbc.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var records []*Post
	if err := qb.dbc.Scan(rows, &records); err != nil {
		rows.Close()
		return nil, err
	}
	rows.Close()

	sc := goen.NewScopeCache(metaSchema)
	for _, record := range records {
		sc.AddObject(record)
	}
	if err := qb.dbc.Include(records, sc, qb.includeLoaders); err != nil {
		return nil, err
	}

	return records, nil
}

type _Post_CreatedAt_OrderExpr string

func (s _Post_CreatedAt_OrderExpr) PostOrderExpr() string {
	return string(s)
}

type _Post_CreatedAt string

func (c _Post_CreatedAt) PostColumnExpr() string {
	return string(c)
}

func (c _Post_CreatedAt) String() string {
	return string(c)
}

func (c _Post_CreatedAt) Eq(v time.Time) PostSqlizer {
	return &_PostSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Post_CreatedAt) NotEq(v time.Time) PostSqlizer {
	return &_PostSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Post_CreatedAt) In(v []time.Time) PostSqlizer {
	return &_PostSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Post_CreatedAt) NotIn(v []time.Time) PostSqlizer {
	return &_PostSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Post_CreatedAt) Asc() PostOrderExpr {
	return _Post_CreatedAt_OrderExpr(string(c))
}

func (c _Post_CreatedAt) Desc() PostOrderExpr {
	return _Post_CreatedAt_OrderExpr(string(c) + " DESC")
}

type _Post_UpdatedAt_OrderExpr string

func (s _Post_UpdatedAt_OrderExpr) PostOrderExpr() string {
	return string(s)
}

type _Post_UpdatedAt string

func (c _Post_UpdatedAt) PostColumnExpr() string {
	return string(c)
}

func (c _Post_UpdatedAt) String() string {
	return string(c)
}

func (c _Post_UpdatedAt) Eq(v time.Time) PostSqlizer {
	return &_PostSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Post_UpdatedAt) NotEq(v time.Time) PostSqlizer {
	return &_PostSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Post_UpdatedAt) In(v []time.Time) PostSqlizer {
	return &_PostSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Post_UpdatedAt) NotIn(v []time.Time) PostSqlizer {
	return &_PostSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Post_UpdatedAt) Asc() PostOrderExpr {
	return _Post_UpdatedAt_OrderExpr(string(c))
}

func (c _Post_UpdatedAt) Desc() PostOrderExpr {
	return _Post_UpdatedAt_OrderExpr(string(c) + " DESC")
}

type _Post_BlogID_OrderExpr string

func (s _Post_BlogID_OrderExpr) PostOrderExpr() string {
	return string(s)
}

type _Post_BlogID string

func (c _Post_BlogID) PostColumnExpr() string {
	return string(c)
}

func (c _Post_BlogID) String() string {
	return string(c)
}

func (c _Post_BlogID) Eq(v uuid.UUID) PostSqlizer {
	return &_PostSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Post_BlogID) NotEq(v uuid.UUID) PostSqlizer {
	return &_PostSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Post_BlogID) In(v []uuid.UUID) PostSqlizer {
	return &_PostSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Post_BlogID) NotIn(v []uuid.UUID) PostSqlizer {
	return &_PostSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Post_BlogID) Asc() PostOrderExpr {
	return _Post_BlogID_OrderExpr(string(c))
}

func (c _Post_BlogID) Desc() PostOrderExpr {
	return _Post_BlogID_OrderExpr(string(c) + " DESC")
}

type _Post_PostID_OrderExpr string

func (s _Post_PostID_OrderExpr) PostOrderExpr() string {
	return string(s)
}

type _Post_PostID string

func (c _Post_PostID) PostColumnExpr() string {
	return string(c)
}

func (c _Post_PostID) String() string {
	return string(c)
}

func (c _Post_PostID) Eq(v int) PostSqlizer {
	return &_PostSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Post_PostID) NotEq(v int) PostSqlizer {
	return &_PostSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Post_PostID) In(v []int) PostSqlizer {
	return &_PostSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Post_PostID) NotIn(v []int) PostSqlizer {
	return &_PostSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Post_PostID) Asc() PostOrderExpr {
	return _Post_PostID_OrderExpr(string(c))
}

func (c _Post_PostID) Desc() PostOrderExpr {
	return _Post_PostID_OrderExpr(string(c) + " DESC")
}

type _Post_Title_OrderExpr string

func (s _Post_Title_OrderExpr) PostOrderExpr() string {
	return string(s)
}

type _Post_Title string

func (c _Post_Title) PostColumnExpr() string {
	return string(c)
}

func (c _Post_Title) String() string {
	return string(c)
}

func (c _Post_Title) Eq(v string) PostSqlizer {
	return &_PostSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Post_Title) NotEq(v string) PostSqlizer {
	return &_PostSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Post_Title) In(v []string) PostSqlizer {
	return &_PostSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Post_Title) NotIn(v []string) PostSqlizer {
	return &_PostSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Post_Title) Like(v string) PostSqlizer {
	return &_PostSqlizer{squirrel.Expr(string(c)+" LIKE ?", v)}
}

func (c _Post_Title) NotLike(v string) PostSqlizer {
	return &_PostSqlizer{squirrel.Expr(string(c)+" NOT LIKE ?", v)}
}

func (c _Post_Title) Asc() PostOrderExpr {
	return _Post_Title_OrderExpr(string(c))
}

func (c _Post_Title) Desc() PostOrderExpr {
	return _Post_Title_OrderExpr(string(c) + " DESC")
}

type _Post_Content_OrderExpr string

func (s _Post_Content_OrderExpr) PostOrderExpr() string {
	return string(s)
}

type _Post_Content string

func (c _Post_Content) PostColumnExpr() string {
	return string(c)
}

func (c _Post_Content) String() string {
	return string(c)
}

func (c _Post_Content) Eq(v string) PostSqlizer {
	return &_PostSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Post_Content) NotEq(v string) PostSqlizer {
	return &_PostSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Post_Content) In(v []string) PostSqlizer {
	return &_PostSqlizer{squirrel.Eq{string(c): v}}
}

func (c _Post_Content) NotIn(v []string) PostSqlizer {
	return &_PostSqlizer{squirrel.NotEq{string(c): v}}
}

func (c _Post_Content) Like(v string) PostSqlizer {
	return &_PostSqlizer{squirrel.Expr(string(c)+" LIKE ?", v)}
}

func (c _Post_Content) NotLike(v string) PostSqlizer {
	return &_PostSqlizer{squirrel.Expr(string(c)+" NOT LIKE ?", v)}
}

func (c _Post_Content) Asc() PostOrderExpr {
	return _Post_Content_OrderExpr(string(c))
}

func (c _Post_Content) Desc() PostOrderExpr {
	return _Post_Content_OrderExpr(string(c) + " DESC")
}

type PostDBSet struct {
	dbc *goen.DBContext

	CreatedAt _Post_CreatedAt

	UpdatedAt _Post_UpdatedAt

	BlogID _Post_BlogID

	PostID _Post_PostID

	Title _Post_Title

	Content _Post_Content

	IncludeBlog goen.IncludeLoader
}

func newPostDBSet(dbc *goen.DBContext) *PostDBSet {
	dbset := &PostDBSet{
		dbc: dbc,
	}
	dbset.CreatedAt = "created_at"
	dbset.UpdatedAt = "updated_at"
	dbset.BlogID = "blog_id"
	dbset.PostID = "post_id"
	dbset.Title = "title"
	dbset.Content = "content"

	dbset.IncludeBlog = goen.IncludeLoaderFunc(dbset.includeBlog)

	return dbset
}

func (dbset *PostDBSet) String() string {
	return "posts"
}

func (dbset *PostDBSet) Insert(v *Post) {
	dbset.dbc.Patch(metaSchema.InsertPatchOf(v))
}

func (dbset *PostDBSet) Select() PostQueryBuilder {
	return newPostQueryBuilder(dbset.dbc)
}

func (dbset *PostDBSet) Update(v *Post) {
	dbset.dbc.Patch(metaSchema.UpdatePatchOf(v))
}

func (dbset *PostDBSet) Delete(v *Post) {
	dbset.dbc.Patch(metaSchema.DeletePatchOf(v))
}

func (dbset *PostDBSet) includeBlog(later *goen.IncludeBuffer, sc *goen.ScopeCache, records interface{}) error {
	entities, ok := records.([]*Post)
	if !ok {
		return nil
	}

	childRowKeyOf := func(v *Post) goen.RowKey {
		return &goen.MapRowKey{
			Table: "blogs",
			Key: map[string]interface{}{
				"blog_id": v.BlogID,
			},
		}
	}

	// filter cached entity
	cachedChildRowKeys := make([]goen.RowKey, 0, len(entities))
	noCachedChildRowKeys := make([]goen.RowKey, 0, len(entities))
	for _, entity := range entities {
		key := childRowKeyOf(entity)
		if sc.HasObject(key) {
			cachedChildRowKeys = append(cachedChildRowKeys, key)
		} else {
			noCachedChildRowKeys = append(noCachedChildRowKeys, key)
		}
	}
	if len(noCachedChildRowKeys) > 0 {
		cond := squirrel.Or{}
		for _, rowKey := range noCachedChildRowKeys {
			cond = append(cond, rowKey)
		}
		stmtBuilder := squirrel.StatementBuilder.PlaceholderFormat(dbset.dbc.Dialect().PlaceholderFormat())
		query, args, err := stmtBuilder.Select(
			"blog_id",
			"name",
			"author",
		).From("blogs").Where(cond).ToSql()
		if err != nil {
			return err
		}
		rows, err := dbset.dbc.Query(query, args...)
		if err != nil {
			return err
		}

		var noCachedEntities []*Blog
		if err := dbset.dbc.Scan(rows, &noCachedEntities); err != nil {
			rows.Close()
			return err
		}
		rows.Close()

		for _, entity := range noCachedEntities {
			sc.AddObject(entity)
		}

		// for newly loaded entity, to be filled by includeLoader
		later.AddRecords(noCachedEntities)
	}

	for _, entity := range entities {
		childRowKey := childRowKeyOf(entity)
		raw := sc.GetObject(childRowKey)
		entity.Blog = raw.(*Blog)
	}

	return nil
}

type DBContext struct {
	*goen.DBContext

	Blog *BlogDBSet

	Post *PostDBSet
}

func NewDBContext(dialectName string, db *sql.DB) *DBContext {
	dbc := goen.NewDBContext(dialectName, db)
	return &DBContext{
		DBContext: dbc,
		Blog:      newBlogDBSet(dbc),
		Post:      newPostDBSet(dbc),
	}
}

func (dbc *DBContext) UseTx(tx *sql.Tx) *DBContext {
	clone := dbc.DBContext.UseTx(tx)
	return &DBContext{
		DBContext: clone,
		Blog:      newBlogDBSet(clone),
		Post:      newPostDBSet(clone),
	}
}
