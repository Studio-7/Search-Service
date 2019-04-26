package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sort"
	"github.com/cvhariharan/Utils/utils"
	ct "github.com/cvhariharan/Utils/customtype"
	"github.com/fatih/set"
	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
	"github.com/joho/godotenv"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

var (
	postSearcher = riot.Engine{}
	userSearcher = riot.Engine{}
)

var session *r.Session

var db string

var tcsJSON []byte

var profileJSON []byte

// Initializes the riot search engine index with posts that already exist in the database
func initPostIndex(rows *r.Cursor) {

	var response map[string]interface{}

	for rows.Next(&response) {
		message := response["body"].(map[string]interface{})["message"]
		postSearcher.Index(response["id"].(string), types.DocData{Content: message.(string)})
		postSearcher.Flush()
	}
}

// Initializes the riot search engine index with users that already exist in the database
func initUserIndex(rows *r.Cursor) {

	var response map[string]interface{}

	for rows.Next(&response) {
		name := response["username"]
		userSearcher.Index(response["username"].(string), types.DocData{Content: name.(string)})
		userSearcher.Flush()
	}
}

// Listens to RethinkDB Changefeeds and adds new posts to the Index
func listenToPostChangefeeds(res *r.Cursor) {

	var value map[string]interface{}

	for res.Next(&value) {
		fmt.Println("Value: ", value)
		if value["old_val"] == nil {
			n := value["new_val"]
			addToPostIndex(n)
		}
	}
}

// Adds relevant fields of the post to the Index
func addToPostIndex(val interface{}) {

	p := val.(map[string]interface{})
	body := p["body"].(map[string]interface{})

	postSearcher.Index(p["id"].(string), types.DocData{Content: body["message"].(string)})
	postSearcher.Flush()
}

func listenToUserChangefeeds(res *r.Cursor) {

	var value map[string]interface{}

	for res.Next(&value) {
		fmt.Println("Value: ", value)
		if value["old_val"] == nil {
			n := value["new_val"]
			addToUserIndex(n)
		}
	}
}

// Adds relevant fields of the user to the Index
func addToUserIndex(val interface{}) {

	u := val.(map[string]interface{})
	FName := u["FName"]

	userSearcher.Index(u["UName"].(string), types.DocData{Content: FName.(string)})
	userSearcher.Flush()
}

// Writes a JSON response to the search query
func searchHandler(w http.ResponseWriter, r *http.Request) {
	var jsonString string
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	searchterm := r.FormValue("query")
	username := r.FormValue("username")
	
	tcsJSON := getRelevantPosts(searchterm, session)
	profileJSON := getRelevantUsers(searchterm, session)

	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	result := `{ "profile": ` + string(profileJSON) + `, "post": ` + string(tcsJSON) + "}"

	jwt := utils.GenerateJWT(username, session)
	jsonString = `{ "result": ` + string(result) + `, "token": "` + jwt + "\"}"
	
	fmt.Println(jsonString)

	w.Write([]byte(jsonString))
}

// Runs the search engine to get relevant results
func getRelevantPosts(term string, session *r.Session) []byte {

	req := types.SearchReq{Text: term}
	search := postSearcher.Search(req)
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
	var tcs []ct.TravelCapsule
	var fields ct.TravelCapsule
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

	sort.Sort(ct.TravelCapsules(tcs))
	tcsJSON, _ = json.Marshal(tcs)

	return tcsJSON
}

// Runs the search engine to get relevant results
func getRelevantUsers(term string, session *r.Session) []byte {

	req := types.SearchReq{Text: term}
	search := userSearcher.Search(req)
	docs := (search.Docs).(types.ScoredDocs)

	var ids []string
	var profilesList []ct.Profile
	tol := 5

	// Retrieving the most relevant users wrt the query term
	for i := 0; i < tol; i++ {
		if i < len(docs) {
			ids = append(ids, docs[i].DocId)
		}
	}

	fmt.Println(ids)

	for _, uname := range ids {

		profiles := utils.GetProfile(uname, session)

		// profileRows, _ := r.DB(db).Table(os.Getenv("TCTABLE")).Get(tcID).Run(session)
		// tcRows.One(&fields)

		profilesList = append(profilesList, profiles)
	}

	profileJSON, _ = json.Marshal(profilesList)

	//profileJSON:= []byte("wait")
	return profileJSON
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

	postRes, err := r.DB(db).Table(os.Getenv("POSTTABLE")).Changes().Run(session)
	userRes, err := r.DB(db).Table(os.Getenv("USERTABLE")).Changes().Run(session)

	if err != nil {
		fmt.Println("Could not run Changefeeds")
		log.Fatalln(err)
	}

	// Initializing the riot searcher
	postSearcher.Init(types.EngineOpts{
		NotUseGse: true,
	})

	userSearcher.Init(types.EngineOpts{
		NotUseGse: true,
	})

	// Retrieving already-existing posts to initialize the index with
	postRows, err := r.DB(db).Table(os.Getenv("POSTTABLE")).Run(session)

	if err != nil {
		fmt.Println("Could not retrieve posts from database")
		fmt.Println(err)
		return
	}

	userRows, err := r.DB(db).Table(os.Getenv("USERTABLE")).Run(session)

	if err != nil {
		fmt.Println("Could not retrieve users from database")
		fmt.Println(err)
		return
	}

	initPostIndex(postRows)
	initUserIndex(userRows)

	// Running another goroutine that listens to Changefeeds
	go listenToPostChangefeeds(postRes)
	go listenToUserChangefeeds(userRes)

	http.HandleFunc("/search/find", utils.AuthMiddleware(searchHandler, session))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}