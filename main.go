package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cvhariharan/Utils/utils"
	"github.com/fatih/set"
	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
	"github.com/joho/godotenv"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

var (
	searcher = riot.Engine{}
)

var session *r.Session

var db string

var tcsJSON []byte

// Initializes the riot search engine index with posts that already exist in the database
func initIndex(rows *r.Cursor) {

	var response map[string]interface{}

	for rows.Next(&response) {
		message := response["body"].(map[string]interface{})["message"]
		searcher.Index(response["id"].(string), types.DocData{Content: message.(string)})
		searcher.Flush()
	}
}

// Listens to RethinkDB Changefeeds and adds new posts to the Index
func listenToChangefeeds(res *r.Cursor) {

	var value map[string]interface{}

	for res.Next(&value) {
		fmt.Println("Value: ", value)
		if value["old_val"] == nil {
			n := value["new_val"]
			addToIndex(n)
		}
	}
}

// Adds relevant fields of the post to the Index
func addToIndex(val interface{}) {

	p := val.(map[string]interface{})
	body := p["body"].(map[string]interface{})

	searcher.Index(p["id"].(string), types.DocData{Content: body["message"].(string)})
	searcher.Flush()
}

// Runs the search engine to get relevant results
func getResults(term string, session *r.Session) []byte {

	req := types.SearchReq{Text: term}
	search := searcher.Search(req)
	docs := (search.Docs).(types.ScoredDocs)

	var ids []string
	tol := 15

	// Retrieving the most relevant posts wrt the query term
	for i := 0; i < tol; i++ {
		if i < len(docs) {
			ids = append(ids, docs[i].DocId)
		}
	}

	// Utilizing set variable to deal with duplicate results
	var tcs []map[string]interface{}
	var fields map[string]interface{}
	s := set.New(set.ThreadSafe)

	for _, id := range ids {
		tcIDs := utils.GetTC(id, session)
		for _, t := range tcIDs {
			s.Add(t)
		}
	}

	for _, tcID := range s.List() {

		tcRows, _ := r.DB(db).Table(os.Getenv("TCTABLE")).Get(tcID).Run(session)
		tcRows.One(&fields)
		tcs = append(tcs, fields)

	}

	tcsJSON, _ = json.Marshal(tcs)

	return tcsJSON
}

// Writes a JSON response to the search query
func searchHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	searchterm := r.FormValue("query")
	tcsJSON := getResults(searchterm, session)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tcsJSON)
}

func main() {

	e := godotenv.Load()
	if e != nil {
		log.Fatal(e)
	}
	endpoints := os.Getenv("DBURL")
	url := strings.Split(endpoints, ",")
	dbpass := os.Getenv("DBPASS")
	s, err := r.Connect(r.ConnectOpts{
		Addresses: url,
		Password:  dbpass,
	})

	db = os.Getenv("DB")

	session = s

	res, err := r.DB(db).Table(os.Getenv("POSTTABLE")).Changes().Run(session)

	if err != nil {
		fmt.Println("Could not run Changefeeds")
		log.Fatalln(err)
	}

	// Initializing the riot searcher
	searcher.Init(types.EngineOpts{
		NotUseGse: true,
	})

	// Retrieving already-existing posts to initialize the index with
	rows, err := r.DB(db).Table(os.Getenv("POSTTABLE")).Run(session)

	if err != nil {
		fmt.Println("Could not retrive posts from database")
		fmt.Println(err)
		return
	}

	initIndex(rows)

	// Running another goroutine that listens to Changefeeds
	go listenToChangefeeds(res)

	http.HandleFunc("/search/find", searchHandler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}