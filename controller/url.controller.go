package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"urlshortner/services"
)

type UrlController struct {
	urlService services.UrlService
}

func NewUrlController(urlService services.UrlService)*UrlController{
	return&UrlController{
		urlService: urlService,
	}
}

func (uc *UrlController)CreateShortUrlHandler(w http.ResponseWriter,r *http.Request){
	var requestBody map[string]string
	if err:=json.NewDecoder(r.Body).Decode(&requestBody);err!=nil{
		http.Error(w,"Cannot decode request body",http.StatusBadRequest)
		return
	}
	longUrl:=requestBody["longUrl"]
	if longUrl==""{
		http.Error(w,"long url required",http.StatusBadRequest)
		return
	}
	url,err:=uc.urlService.CreateShortUrl(longUrl)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	response:=map[string]any{
		"message":"Successfully created a short url",
		"success":true,
		"error":nil,
		"data":url,
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	if err:=json.NewEncoder(w).Encode(response);err!=nil{
		http.Error(w,"failed to encode response",http.StatusInternalServerError)
		return
	}
}


func (uc *UrlController)GetLongUrlHandler(w http.ResponseWriter,r *http.Request){
	shortUrl:=r.PathValue("shorturl")
	if shortUrl==""{
		http.Error(w,"short url required",http.StatusBadRequest)
		return
	}
	longUrl,err:=uc.urlService.GetLongUrl(shortUrl)
	if err != nil {
		http.Error(w,err.Error(),http.StatusNotFound)
		return
	}
	
	w.Header().Set("Location",longUrl)
	http.Redirect(w,r,longUrl,http.StatusFound)
}

func (uc *UrlController)GetIdByLongUrlHandler(w http.ResponseWriter,r *http.Request){
	fmt.Println("GetIdByLongUrlHandler reached")
	longUrl:=r.URL.Query().Get("longurl")
	if longUrl == "" {
		http.Error(w, "longurl is required", http.StatusInternalServerError)
		return
	}
	fmt.Println(longUrl)
	id,err:=uc.urlService.GetIdByLongUrl(longUrl)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	if err:=json.NewEncoder(w).Encode(id);err!=nil{
		http.Error(w,"Cannot encode the response",http.StatusInternalServerError)
		return
	}
}


func (uc *UrlController)CreateLongUrl(w http.ResponseWriter,r *http.Request){
	var requestBody map[string]string
	if err:=json.NewDecoder(r.Body).Decode(&requestBody);err!=nil{
		http.Error(w,"Cannot decode request body",http.StatusBadRequest)
		return
	}
	longUrl:=requestBody["longUrl"]
	if longUrl==""{
		http.Error(w,"long url required",http.StatusBadRequest)
		return
	}
	id,err:=uc.urlService.CreateLongUrl(longUrl)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	response:=map[string]any{
		"message":"Successfully created a short url",
		"success":true,
		"error":nil,
		"data":id,
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	if err:=json.NewEncoder(w).Encode(response);err!=nil{
		http.Error(w,"failed to encode response",http.StatusInternalServerError)
		return
	}
	
}