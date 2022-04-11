Var Db *sql.DB 
func init() {
	var err error
	Db, err = sql.Open("mingjingyu", "dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}