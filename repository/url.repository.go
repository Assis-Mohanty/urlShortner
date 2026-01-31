package repository

import (
	"database/sql"
	"fmt"
	"urlshortner/models"
)

type UrlRepository interface {
	CreateShortUrl(longUrl string ,shortUrl string)(int,error)
	GetLongUrl(shortUrl string)(string,error )
	GetIdByLongUrl(longUrl string)(int,error)
	CreateLongUrl(longUrl string )(int ,error)
}

type UrlRepositoryImpl struct {
	db *sql.DB
}

func NewUrlRepository(db *sql.DB)UrlRepository{
	return&UrlRepositoryImpl{
		db:db,
	}
}

func (ur *UrlRepositoryImpl)CreateShortUrl(longUrl string ,shortUrl string)(int,error){
	query:= "UPDATE urls SET short_url = ? WHERE long_url = ?;"
	row,err:=ur.db.Exec(query , shortUrl,longUrl)
	fmt.Println(row)
	if err != nil {
		return 0,err
	}
	id,err:=row.LastInsertId()
	if err != nil {
		return 0,err 
	}
	fmt.Println(id)
	return int(id),nil
}

func (ur *UrlRepositoryImpl) GetLongUrl(shortUrl string)(string ,error ){
	query:="SELECT id, long_url  FROM urls WHERE short_url =?;"
	row:=ur.db.QueryRow(query,shortUrl)
	url:=&models.Url{}
	err:=row.Scan(&url.Id,&url.LongUrl)
	if err != nil{
		fmt.Println("long url not found with the short url",shortUrl)
		return "", err
	}
	return url.LongUrl,nil
}

func (ur *UrlRepositoryImpl) GetIdByLongUrl(longUrl string)(int,error){
	query:="SELECT id FROM urls WHERE long_url = ?;"
	row:=ur.db.QueryRow(query,longUrl)
	url:=&models.Url{}
	err:=row.Scan(&url.Id)
	if err != nil{
		fmt.Println("id not found with long url", longUrl)
		return 0, err
	}
	return int(url.Id),nil
}

func (ur *UrlRepositoryImpl)CreateLongUrl(longUrl string )(int ,error){
	query:="INSERT INTO urls (long_url) VALUES(?)"
	row,err:=ur.db.Exec(query,longUrl)
	if err != nil {
		return -1,err
	}
	id,err:=row.LastInsertId()
	return int(id),err
}
