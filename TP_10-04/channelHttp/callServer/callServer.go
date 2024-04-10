package callServer

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	RespText string
	Err error
}

func CallServer (url string, response chan(Response)) {
	res, err := http.Get(url)
	if err!= nil {
		log.Fatal(err)
		response <- Response{Err: err}
		return
	}
	if res.StatusCode != 200 {
		response <- Response{ Err: errors.New("Le code retourné par le serveur indique une erreur: " + strconv.Itoa(res.StatusCode)) }
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err!= nil {
        log.Fatal(err)
        response <- Response{Err: err}
        return
    }
	response <- Response{RespText: string(body)}
}