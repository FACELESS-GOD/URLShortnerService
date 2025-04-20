package Model

import (
	"errors"

	"github.com/FACELESS-GOD/URLShortnerService.git/Helper/QueryGenerator"
	"github.com/FACELESS-GOD/URLShortnerService.git/Helper/StructStore"
	"github.com/FACELESS-GOD/URLShortnerService.git/Package/Utility"
)

func IsURLProcessed(ReqData StructStore.URLRequest) (bool, error) {
	numericResponse := StructStore.DbNumericResponse{}

	query := QueryGenerator.IsprocessedQueryGenerator(ReqData)

	dbResponseInstance, err := Utility.DbInstance.Query(query)
	if err != nil {
		return false, err
	}
	for dbResponseInstance.Next() {
		err := dbResponseInstance.Scan(&numericResponse.Count)
		if err != nil {
			return false, err
		}
	}

	if numericResponse.Count > 0 {
		return true, nil
	} else {
		return false, nil
	}

}

func AddURLTODb(ReqData StructStore.URLRequest) (string, error) {
	isProcessed, err := IsURLProcessed(ReqData)
	if err != nil {
		return "", err
	}

	if isProcessed != true {

		newQuery := "http://localhost:9000/ID/{" + Utility.RandomURLString() + "}/"

		dbQuery := QueryGenerator.InsertQueryGenerator(ReqData, newQuery)

		dbResponseInstance, err := Utility.DbInstance.Query(dbQuery)
		if err != nil {
			return "", err
		}

		for dbResponseInstance.Next() {
		}

		errRedis := Utility.RedisInstance.Set(Utility.Context, newQuery, ReqData.URL, 0).Err()
		if errRedis != nil {
			return "", err
		}

		return newQuery, nil

	} else {

		numericResponse := StructStore.DbNumericResponse_Adv{}

		dbQuery := QueryGenerator.GetQueryGenerator(ReqData)

		dbResponseInstance, err := Utility.DbInstance.Query(dbQuery)
		if err != nil {
			return "", err
		}

		for dbResponseInstance.Next() {
			err := dbResponseInstance.Scan(&numericResponse.ShortenedURL)
			if err != nil {
				return "", err
			}
		}

		if numericResponse.ShortenedURL != "" {
			return numericResponse.ShortenedURL, nil
		} else {
			err := errors.New("Error Occured")
			return "", err
		}

	}

}
