package api

import (
	"io/ioutil"
	"strconv"

	errs "github.com/pkg/errors"

	"deus-task/core"
	jsn "deus-task/serializer/js"
	"encoding/json"
	"log"
	"net/http"
)

type VisitsHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	srv core.VisitsService
}

func NewVisitsHandler(vs core.VisitsService) VisitsHandler {
	return &handler{srv: vs}
}
func setupResponse(w http.ResponseWriter, contentType string, statusCode int, body []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serialize(contentType string) core.VisitsSerializer {
	if contentType == "application/json" {
		return &jsn.VisitsSerializer{}
	}
	return &jsn.VisitsSerializer{}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		e := &ErrStr{http.StatusText(http.StatusMethodNotAllowed)}
		JSONError(w, e, http.StatusMethodNotAllowed)
	}
	url := r.URL.Query().Get("url")
	numVisits, err := h.srv.GetUniqueVisitsNumber(url)
	if err != nil {
		if errs.Cause(err) == core.ErrUrlNotFound {
			e := &ErrStr{http.StatusText(http.StatusNotFound)}
			JSONError(w, e, http.StatusNotFound)
			return
		}
		e := &ErrStr{http.StatusText(http.StatusInternalServerError)}
		JSONError(w, e, http.StatusInternalServerError)
		return
	}
	setupResponse(w, "application/json", http.StatusOK, []byte("{visitas: "+strconv.Itoa(numVisits)+"}"))

}
func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		e := &ErrStr{http.StatusText(http.StatusMethodNotAllowed)}
		JSONError(w, e, http.StatusMethodNotAllowed)
	}
	contentType := r.Header.Get("Content-Type")
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e := &ErrStr{"ioutil.ReadAll(r.Body)" + http.StatusText(http.StatusInternalServerError)}
		JSONError(w, e, http.StatusInternalServerError)
		return
	}

	v, err := h.serialize(contentType).Decode(rBody)
	if err != nil {
		e := &ErrStr{"h.serialize(contentType).Decode(rBody)" + http.StatusText(http.StatusInternalServerError)}
		JSONError(w, e, http.StatusInternalServerError)
		return
	}

	err = h.srv.SaveVisit(v)
	if err != nil {
		if errs.Cause(err) == core.ErrUrlInvalid {
			e := &ErrStr{http.StatusText(http.StatusBadRequest)}
			JSONError(w, e, http.StatusBadRequest)
			return
		}
		e := &ErrStr{http.StatusText(http.StatusInternalServerError)}
		JSONError(w, e, http.StatusInternalServerError)
		return
	}
	rBody, err = h.serialize(contentType).Encode(v)
	if err != nil {
		e := &ErrStr{http.StatusText(http.StatusInternalServerError)}
		JSONError(w, e, http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, http.StatusCreated, rBody)
}

type ErrStr struct {
	Error string `json:"error"`
}

func JSONError(w http.ResponseWriter, err *ErrStr, code int) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	e := json.NewEncoder(w).Encode(err)
	if e != nil {
		panic(e)
	}
}
