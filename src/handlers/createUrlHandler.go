package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/faizkhan-06/gozap/config"
	"github.com/faizkhan-06/gozap/src/models"
	"github.com/faizkhan-06/gozap/types"
	"github.com/faizkhan-06/gozap/utils"
)

func CreateUrlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	
	var urlData types.UrlData
	var existedUrl models.Urls
	
	json.NewDecoder(r.Body).Decode(&urlData)

	if urlData.LongUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(types.Response{
			Message: "URL is required",
			Status: http.StatusBadRequest,
		})
		return
	}

	matched, err := regexp.Match(`^(https?:\/\/)?([a-zA-Z0-9.-]+)\.([a-zA-Z]{2,6})(\/[^\s]*)?$`, []byte(urlData.LongUrl))
	if err != nil {
		log.Fatal("Match Url failed")
	}

	if !matched {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(types.Response{
			Message: "Please enter valid url",
			Status: http.StatusBadRequest,
		})
		return
	}

	result := config.DB.Where("long_url = ?",urlData.LongUrl).First(&existedUrl)

	if result.RowsAffected > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(types.Response{
			Message: "ok",
			Status: http.StatusOK,
			Data: existedUrl,
		})
		return
	}


	shortId := utils.GenerateShortId()

	data := models.Urls{
		ShortUrl: shortId,
		LongUrl:  urlData.LongUrl,
	}

	res := config.DB.Model(&models.Urls{}).Create(&data)

	if res.Error != nil{
		 w.WriteHeader(http.StatusInternalServerError)
		 json.NewEncoder(w).Encode(types.Response{
			Message: res.Error.Error(),
			Status: http.StatusInternalServerError,
		 })
		return
	}
	
	data.ShortUrl = os.Getenv("DOMAIN") + shortId

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(types.Response{
		Message: "OK",
		Status: http.StatusOK,
		Data: data,
	})

}