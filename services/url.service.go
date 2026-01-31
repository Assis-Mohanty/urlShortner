package services

import (
	"fmt"
	"urlshortner/repository"

)

type UrlService interface {
	CreateShortUrl(longUrl string ) (int, error)
	GetLongUrl(shortUrl string) (string, error)
	GetIdByLongUrl(longUrl string)(int,error)
	CreateLongUrl(longUrl string )(int ,error)
}

type UrlServiceImpl struct {
	urlRepository repository.UrlRepository
}

func NewUrlServiceImpl(urlRepository repository.UrlRepository)UrlService{
	return &UrlServiceImpl{
		urlRepository: urlRepository,
	}
}

func (us *UrlServiceImpl)CreateShortUrl(longUrl string)(int ,error){
	shortUrl,err:=us.GenerateShortUrlFromLongUrl(longUrl)
	if err != nil {
		return -1,err
	}
	fmt.Println("qqq",shortUrl)
	fmt.Println(longUrl)
	return us.urlRepository.CreateShortUrl(longUrl,shortUrl)
}

func (us *UrlServiceImpl)GetLongUrl(shortUrl string)(string,error){
	return us.urlRepository.GetLongUrl(shortUrl)
}

func (us *UrlServiceImpl)GetIdByLongUrl(longUrl string)(int,error){
	return us.urlRepository.GetIdByLongUrl(longUrl)
}

func (us *UrlServiceImpl) GenerateShortUrlFromLongUrl(longUrl string)(string,error){
	id,err:=us.GetIdByLongUrl(longUrl)
	if err != nil {
		return err.Error(),err
	}
	idToint62:=toBase62(id)
	fmt.Println(idToint62)
	return idToint62,nil
}

func(us *UrlServiceImpl)CreateLongUrl(longUrl string )(int ,error){
	id,err:=us.urlRepository.CreateLongUrl(longUrl)
	if err != nil {
		return -1,err
	}
	return id,err
}
func toBase62(n int) string {
	const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if n == 0 {
		return "0"
	}

	result := ""
	for n > 0 {
		rem := n % 62
		result = string(base62[rem]) + result
		n /= 62
	}
	return result
}
