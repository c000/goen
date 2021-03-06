package example

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/kamichidu/goen"
	uuid "github.com/satori/go.uuid"
)

func unifyQuery(s string) string {
	s = strings.Replace(s, "$1", "?", -1)
	s = strings.Replace(s, `"`, "`", -1)
	return s
}

func Example() {
	dbc := NewDBContext(prepareDB())
	dbc.DebugMode(true)

	src := []*Blog{
		&Blog{
			BlogID: uuid.Must(uuid.FromString("d03bc237-eef4-4b6f-afe1-ea901357d828")),
			Name:   "testing1",
			Author: "kamichidu",
		},
		&Blog{
			BlogID: uuid.Must(uuid.FromString("b95e5d4d-7eb9-4612-882d-224daa4a59ee")),
			Name:   "testing2",
			Author: "kamichidu",
		},
		&Blog{
			BlogID: uuid.Must(uuid.FromString("22f931c8-ac87-4520-88e8-83fc0604b8f5")),
			Name:   "testing3",
			Author: "kamichidu",
		},
		&Blog{
			BlogID: uuid.Must(uuid.FromString("065c6554-9aff-4b42-ab3b-141ed5ef5624")),
			Name:   "testing4",
			Author: "kamichidu",
		},
	}
	for _, blog := range src {
		dbc.Blog.Insert(blog)
	}
	func(blog *Blog) {
		now, err := time.Parse(time.RFC3339, "2018-06-01T12:00:00Z")
		if err != nil {
			panic(err)
		}
		dbc.Post.Insert(&Post{
			BlogID:  blog.BlogID,
			Title:   "titleA",
			Content: "contentA",
			Timestamp: Timestamp{
				CreatedAt: now,
				UpdatedAt: now,
			},
		})
		dbc.Post.Insert(&Post{
			BlogID:  blog.BlogID,
			Title:   "titleB",
			Content: "contentB",
			Timestamp: Timestamp{
				CreatedAt: now,
				UpdatedAt: now,
			},
		})
	}(src[0])
	src[1].Author = "unknown"
	dbc.Blog.Update(src[1])
	dbc.Blog.Delete(src[2])
	src[3].Name = "updating"
	dbc.Blog.Update(src[3])
	if err := dbc.SaveChanges(); err != nil {
		panic(err)
	}

	// counting all records
	numBlogs, err := dbc.Blog.Select().Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("all blogs = %d\n", numBlogs)

	// querying with conditions
	blogs, err := dbc.Blog.Select().
		Include(dbc.Blog.IncludePosts, dbc.Post.IncludeBlog).
		Where(dbc.Blog.Name.Like(`%testing%`)).
		OrderBy(dbc.Blog.Name.Asc()).
		Query()
	if err != nil {
		panic(err)
	}
	fmt.Printf("found blogs = %d\n", len(blogs))
	spew.Config.SortKeys = true
	spew.Config.MaxDepth = 1
	for _, blog := range blogs {
		spew.Printf("%#v\n", blog)

		for _, post := range blog.Posts {
			spew.Printf("- %#v\n", post)
			spew.Printf("  CreatedAt:%q\n", post.Timestamp.CreatedAt.Format(time.RFC3339))
			spew.Printf("  UpdatedAt:%q\n", post.Timestamp.UpdatedAt.Format(time.RFC3339))
			if post.DeletedAt != nil {
				spew.Printf("  DeletedAt:%q\n", post.DeletedAt.Format(time.RFC3339))
			} else {
				spew.Print("  DeletedAt:nil\n")
			}
		}
	}
	// Output:
	// all blogs = 3
	// found blogs = 2
	// (*example.Blog){BlogID:(uuid.UUID)d03bc237-eef4-4b6f-afe1-ea901357d828 Name:(string)testing1 Author:(string)kamichidu Posts:([]*example.Post)[<max>]}
	// - (*example.Post){Timestamp:(example.Timestamp){<max>} BlogID:(uuid.UUID)d03bc237-eef4-4b6f-afe1-ea901357d828 PostID:(int)1 Title:(string)titleA Content:(string)contentA Order:(int)0 Blog:(*example.Blog){<max>}}
	//   CreatedAt:"2018-06-01T12:00:00Z"
	//   UpdatedAt:"2018-06-01T12:00:00Z"
	//   DeletedAt:nil
	// - (*example.Post){Timestamp:(example.Timestamp){<max>} BlogID:(uuid.UUID)d03bc237-eef4-4b6f-afe1-ea901357d828 PostID:(int)2 Title:(string)titleB Content:(string)contentB Order:(int)0 Blog:(*example.Blog){<max>}}
	//   CreatedAt:"2018-06-01T12:00:00Z"
	//   UpdatedAt:"2018-06-01T12:00:00Z"
	//   DeletedAt:nil
	// (*example.Blog){BlogID:(uuid.UUID)b95e5d4d-7eb9-4612-882d-224daa4a59ee Name:(string)testing2 Author:(string)unknown Posts:([]*example.Post)<nil>}
}

func Example_queryRow() {
	dbc := NewDBContext(prepareDB())
	dbc.DebugMode(true)

	src := []*Blog{
		&Blog{
			BlogID: uuid.Must(uuid.FromString("d03bc237-eef4-4b6f-afe1-ea901357d828")),
			Name:   "testing1",
			Author: "kamichidu",
		},
		&Blog{
			BlogID: uuid.Must(uuid.FromString("b95e5d4d-7eb9-4612-882d-224daa4a59ee")),
			Name:   "testing2",
			Author: "kamichidu",
		},
		&Blog{
			BlogID: uuid.Must(uuid.FromString("22f931c8-ac87-4520-88e8-83fc0604b8f5")),
			Name:   "testing3",
			Author: "kamichidu",
		},
		&Blog{
			BlogID: uuid.Must(uuid.FromString("065c6554-9aff-4b42-ab3b-141ed5ef5624")),
			Name:   "testing4",
			Author: "kamichidu",
		},
	}
	for _, blog := range src {
		dbc.Blog.Insert(blog)
	}
	func(blog *Blog) {
		now, err := time.Parse(time.RFC3339, "2018-06-01T12:00:00Z")
		if err != nil {
			panic(err)
		}
		dbc.Post.Insert(&Post{
			BlogID:  blog.BlogID,
			Title:   "titleA",
			Content: "contentA",
			Timestamp: Timestamp{
				CreatedAt: now,
				UpdatedAt: now,
			},
		})
		dbc.Post.Insert(&Post{
			BlogID:  blog.BlogID,
			Title:   "titleB",
			Content: "contentB",
			Timestamp: Timestamp{
				CreatedAt: now,
				UpdatedAt: now,
			},
		})
	}(src[0])
	if err := dbc.SaveChanges(); err != nil {
		panic(err)
	}

	// when a record was not found, will get sql.ErrNoRows
	_, err := dbc.Blog.Select().
		Where(dbc.Blog.Author.Eq("non-exists-author")).
		QueryRow()
	if err == sql.ErrNoRows {
		fmt.Print("QueryRow returns sql.ErrNoRows when a record was not found.\n")
	}

	// querying a record with conditions
	blog, err := dbc.Blog.Select().
		Include(dbc.Blog.IncludePosts, dbc.Post.IncludeBlog).
		Where(dbc.Blog.Name.Eq(`testing1`)).
		QueryRow()
	if err != nil {
		panic(err)
	}
	spew.Config.SortKeys = true
	spew.Config.MaxDepth = 1
	spew.Printf("%#v\n", blog)
	for _, post := range blog.Posts {
		spew.Printf("- %#v\n", post)
		spew.Printf("  CreatedAt:%q\n", post.Timestamp.CreatedAt.Format(time.RFC3339))
		spew.Printf("  UpdatedAt:%q\n", post.Timestamp.UpdatedAt.Format(time.RFC3339))
	}
	// Output:
	// QueryRow returns sql.ErrNoRows when a record was not found.
	// (*example.Blog){BlogID:(uuid.UUID)d03bc237-eef4-4b6f-afe1-ea901357d828 Name:(string)testing1 Author:(string)kamichidu Posts:([]*example.Post)[<max>]}
	// - (*example.Post){Timestamp:(example.Timestamp){<max>} BlogID:(uuid.UUID)d03bc237-eef4-4b6f-afe1-ea901357d828 PostID:(int)1 Title:(string)titleA Content:(string)contentA Order:(int)0 Blog:(*example.Blog){<max>}}
	//   CreatedAt:"2018-06-01T12:00:00Z"
	//   UpdatedAt:"2018-06-01T12:00:00Z"
	// - (*example.Post){Timestamp:(example.Timestamp){<max>} BlogID:(uuid.UUID)d03bc237-eef4-4b6f-afe1-ea901357d828 PostID:(int)2 Title:(string)titleB Content:(string)contentB Order:(int)0 Blog:(*example.Blog){<max>}}
	//   CreatedAt:"2018-06-01T12:00:00Z"
	//   UpdatedAt:"2018-06-01T12:00:00Z"
}

func Example_count() {
	dbc := NewDBContext(prepareDB())
	dbc.DebugMode(true)

	for i := 0; i < 11; i++ {
		dbc.Blog.Insert(&Blog{
			BlogID: uuid.Must(uuid.NewV4()),
			Name:   fmt.Sprintf("name-%d", i),
			Author: "kamichidu",
		})
	}

	dbc.Compiler = goen.BulkCompiler
	if err := dbc.SaveChanges(); err != nil {
		panic(err)
	}

	// counting a record with conditions
	count, err := dbc.Blog.Select().
		Where(dbc.Blog.Name.In(
			`name-3`,
			`name-4`,
			`name-5`,
			`name-6`,
			`name-7`,
		)).
		Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("count = %d\n", count)
	// Output:
	// count = 5
}

func Example_generatedSchemaFields() {
	dbc := NewDBContext(dialectName, nil)

	// can get quoted names
	fmt.Printf("dbc.Blog = %s / %s\n", dbc.Blog, unifyQuery(dbc.Blog.QuotedString()))
	fmt.Printf("dbc.Blog.BlogID = %s / %s\n", dbc.Blog.BlogID, unifyQuery(dbc.Blog.BlogID.QuotedString()))
	fmt.Printf("dbc.Blog.Name = %s / %s\n", dbc.Blog.Name, unifyQuery(dbc.Blog.Name.QuotedString()))
	fmt.Printf("dbc.Blog.Author = %s / %s\n", dbc.Blog.Author, unifyQuery(dbc.Blog.Author.QuotedString()))
	// Output:
	// dbc.Blog = blogs / `blogs`
	// dbc.Blog.BlogID = blog_id / `blog_id`
	// dbc.Blog.Name = name / `name`
	// dbc.Blog.Author = author / `author`
}

func Example_customDelete() {
	dbc := NewDBContext(prepareDB())

	// expect `delete from posts where blog_id = ?`
	dbc.Patch(goen.DeletePatch(
		dbc.Post.String(),
		&goen.MapRowKey{
			Table: dbc.Post.String(),
			Key: map[string]interface{}{
				dbc.Post.BlogID.String(): uuid.Must(uuid.FromString("d03bc237-eef4-4b6f-afe1-ea901357d828")),
			},
		}))
	if err := dbc.SaveChanges(); err != nil {
		panic(err)
	}
	// Output:
}

func Example_transaction() {
	dbc := NewDBContext(prepareDB())
	dbc.DebugMode(true)

	tx, err := dbc.DB.Begin()
	if err != nil {
		panic(err)
	}

	txc := dbc.UseTx(tx)
	blogID := uuid.Must(uuid.FromString("d03bc237-eef4-4b6f-afe1-ea901357d828"))
	txc.Blog.Insert(&Blog{
		BlogID: blogID,
		Name:   "tx",
		Author: "kamichidu",
	})
	if err := txc.SaveChanges(); err != nil {
		panic(err)
	}

	n, err := dbc.Blog.Select().
		Where(dbc.Blog.BlogID.Eq(blogID)).
		Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dbc founds %d blogs when not committed yet\n", n)

	n, err = txc.Blog.Select().
		Where(dbc.Blog.BlogID.Eq(blogID)).
		Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("txc founds %d blogs when not committed yet since it's in same transaction\n", n)

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	n, err = dbc.Blog.Select().
		Where(dbc.Blog.BlogID.Eq(blogID)).
		Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dbc founds %d blogs after committed\n", n)

	// Output:
	// dbc founds 0 blogs when not committed yet
	// txc founds 1 blogs when not committed yet since it's in same transaction
	// dbc founds 1 blogs after committed
}

func Example_nullableColumn() {
	dbc := NewDBContext(prepareDB())
	dbc.DebugMode(true)

	blogID := uuid.Must(uuid.FromString("d03bc237-eef4-4b6f-afe1-ea901357d828"))
	dbc.Blog.Insert(&Blog{
		BlogID: blogID,
		Name:   "nullable",
		Author: "kamichidu",
	})
	now, err := time.Parse(time.RFC3339, "2018-08-09T13:48:29Z")
	if err != nil {
		panic(err)
	}
	dbc.Post.Insert(&Post{
		BlogID: blogID,
		Title:  "p1",
		Timestamp: Timestamp{
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: &now,
		},
	})
	dbc.Post.Insert(&Post{
		BlogID: blogID,
		Title:  "p2",
		Timestamp: Timestamp{
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
	})
	if err := dbc.SaveChanges(); err != nil {
		panic(err)
	}

	fmt.Println("print all rows")
	posts, err := dbc.Post.Select().
		OrderBy(dbc.Post.Title.Asc()).
		Query()
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		if post.DeletedAt != nil {
			fmt.Printf("%q > DeletedAt = %q\n", post.Title, post.DeletedAt.Format(time.RFC3339))
		} else {
			fmt.Printf("%q > DeletedAt = nil\n", post.Title)
		}
	}

	fmt.Println("print filtered rows that deleted at is null")
	posts, err = dbc.Post.Select().
		Where(dbc.Post.DeletedAt.Eq(nil)).
		OrderBy(dbc.Post.Title.Asc()).
		Query()
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		if post.DeletedAt != nil {
			fmt.Printf("%q > DeletedAt = %q\n", post.Title, post.DeletedAt.Format(time.RFC3339))
		} else {
			fmt.Printf("%q > DeletedAt = nil\n", post.Title)
		}
	}

	// Output:
	// print all rows
	// "p1" > DeletedAt = "2018-08-09T13:48:29Z"
	// "p2" > DeletedAt = nil
	// print filtered rows that deleted at is null
	// "p2" > DeletedAt = nil
}

func Example_queryBuilderAsSqlizer() {
	dbc := NewDBContext(prepareDB())

	query, args, err := dbc.Blog.Select().
		Where(dbc.Blog.Author.Eq("kamichidu")).
		ToSqlizer(dbc.Blog.BlogID.QuotedString()).
		ToSql()
	if err != nil {
		panic(err)
	}
	// for stable output
	query = unifyQuery(query)
	fmt.Printf("%q\n", query)
	fmt.Printf("%q\n", args)

	// Output:
	// "SELECT `blog_id` FROM `blogs` WHERE `author` = ?"
	// ["kamichidu"]
}

func Example_brokenManyToOneRelationWithNonForeignKey() {
	dbc := NewDBContext(prepareDB())
	dropForeignKeys(dbc.DB)

	validBlogID := uuid.Must(uuid.FromString("d03bc237-eef4-4b6f-afe1-ea901357d828"))
	invalidBlogID := uuid.Must(uuid.FromString("0c816298-1d35-4948-870a-6f7b5bba14d3"))
	dbc.Blog.Insert(&Blog{
		BlogID: validBlogID,
		Name:   "the blog",
	})
	dbc.Post.Insert(&Post{
		BlogID: validBlogID,
		Title:  "0 the valid post",
		Timestamp: Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
	dbc.Post.Insert(&Post{
		BlogID: invalidBlogID,
		Title:  "1 the invalid post",
		Timestamp: Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
	if err := dbc.SaveChanges(); err != nil {
		panic(err)
	}

	posts, err := dbc.Post.Select().
		Include(dbc.Post.IncludeBlog).
		OrderBy(dbc.Post.Title.Asc()).
		Query()
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		var blogName string
		if post.Blog != nil {
			blogName = post.Blog.Name
		} else {
			blogName = "<nil>"
		}
		fmt.Printf("%q with related blog %q\n", post.Title, blogName)
	}
	// Output:
	// "0 the valid post" with related blog "the blog"
	// "1 the invalid post" with related blog "<nil>"
}
