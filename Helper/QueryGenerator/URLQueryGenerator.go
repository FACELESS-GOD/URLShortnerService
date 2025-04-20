package QueryGenerator

import "github.com/FACELESS-GOD/URLShortnerService.git/Helper/StructStore"

func IsprocessedQueryGenerator(ReqData StructStore.URLRequest) string {
	query := "SELECT COUNT(Original_URL),Original_URL FROM URLDATA WHERE EXISTS (SELECT Original_URL FROM URLDATA WHERE Original_URL ='" + ReqData.URL + "' );"
	return query
}

func InsertQueryGenerator(ReqData StructStore.URLRequest, URLNew string) string {
	query := "INSERT INTO URLDATA (Original_URL, New_URL) VALUES ('" + ReqData.URL + "'+'" + URLNew + "')"
	return query
}

func GetQueryGenerator(ReqData StructStore.URLRequest) string {
	query := "SELECT FIRST(New_URL) FROM URLDATA  WHERE Original_URL = '" + ReqData.URL + "'"
	return query
}
