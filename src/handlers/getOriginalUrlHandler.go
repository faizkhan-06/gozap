package handlers

import (
	"net/http"
	"path"
	"strings"

	"github.com/faizkhan-06/gozap/config"
	"github.com/faizkhan-06/gozap/src/models"
)

func GetOriginalUrlHandler(w http.ResponseWriter, r *http.Request) {
	var longUrl models.Urls
	id := path.Base(r.URL.Path)

	if id == "" {
		http.Error(w, "404 - Not Found", http.StatusNotFound)
		return
	}

	if res:= config.DB.Where("short_url = ?", id).First(&longUrl).Select("long_url"); res.Error != nil || res.RowsAffected == 0 {
		http.Error(w, "404 - Not Found", http.StatusNotFound)
		return
	}

	if !strings.HasPrefix(longUrl.LongUrl, "http://") && !strings.HasPrefix(longUrl.LongUrl, "https://") {
		longUrl.LongUrl = "https://" + longUrl.LongUrl
	}

	http.Redirect(w, r, longUrl.LongUrl, http.StatusFound)
	
}