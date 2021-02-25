package user

import (
	"config"
	"encoding/json"
	"models"
	"net/http"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		responWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		users, err2 := userModel.FindAll()
		if err2 != nil {
			responWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responWithJson(response, http.StatusOK, users)
		}
	}
}

func Search(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := config.GetDB()
	if err != nil {
		responWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		users, err2 := userModel.Search()
		if err2 != nil {
			responWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responWithJson(response, http.StatusOK, users)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	err := json.NewDecoder((request.Body).Decode(&user)
	db, err := config.GetDB()
	if err != nil {
		responWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		err2 := userModel.Create(&user)
		if err2 != nil {
			responWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responWithJson(response, http.StatusOK, users)
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	err := json.NewDecoder((request.Body).Decode(&user)
	db, err := config.GetDB()
	if err != nil {
		responWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		_, err2 := userModel.Update(&user)
		if err2 != nil {
			responWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responWithJson(response, http.StatusOK, users)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)
	err := json.NewDecoder((request.Body).Decode(&user)
	db, err := config.GetDB()
	if err != nil {
		responWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		RowsAffected,err2 := userModel.Delete(id)
		if err2 != nil {
			responWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responWithJson(response, http.StatusOK, map[string] int {
				"RowsAffected" : RowsAffected,
			})
			
		}
	}
}

func responWithError(w http.ResponseWriter, code int, msg string) {
	responWithJson(w, code, map[string]string{"error": msg})
}

func responWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
