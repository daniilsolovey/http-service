package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/reconquest/pkg/log"

	"github.com/go-chi/chi"
)

func (handler *Handler) PrintUsers(w http.ResponseWriter, r *http.Request) {
	result, err := json.Marshal(handler.users)
	if err != nil {
		log.Errorf(
			err,
			"unable to marshal users data to json",
		)

		fmt.Fprintln(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintln(w, string(result))
	if err != nil {
		log.Errorf(
			err,
			"unable to print users",
		)

		fmt.Fprintln(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *Handler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Errorf(
			err,
			"unable to parse id",
		)

		fmt.Fprintln(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var newFields User
	err = json.NewDecoder(r.Body).Decode(&newFields)
	if err != nil {
		log.Errorf(
			err,
			"unable to handle http request",
		)

		fmt.Fprintln(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	handler.updateUser(id, newFields)
}

func (handler *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Errorf(
			err,
			"unable to parse id",
		)

		fmt.Fprintln(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := handler.findUser(id)
	if err != nil {
		log.Errorf(
			err,
			"unable to find user with id: %d",
			id,
		)

		fmt.Fprintln(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(&user)
	if err != nil {
		log.Errorf(
			err,
			"unable to marshal users data to json",
		)

		fmt.Fprintln(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintln(w, string(result))
	if err != nil {
		log.Errorf(
			err,
			"unable to print users",
		)

		fmt.Fprintln(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
