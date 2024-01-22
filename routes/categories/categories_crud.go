package category_routes

import (
	"encoding/json"
	"muxapi/database"
	"muxapi/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
)


func getCategory(w http.ResponseWriter, _ *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  data := models.UserCategories{}
  database.Database.Order(
    clause.OrderByColumn{
      Column: clause.Column{Name: "id"}, 
      Desc: true,
    }).Find(&data)

  json.NewEncoder(w).Encode(data)
}

func getCategoryByID(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  w.Header().Set("Content-Type", "application/json")
  encoder := json.NewEncoder(w)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    encoder.Encode(map[string]string{
      "status":"Error",
      "message":"Unexpected Error",
    })
    return
  }

  w.Header().Set("Content-Type", "application/json")
  var data models.UserCategory
  if err := database.Database.First(&data, id); err.Error != nil {
    w.WriteHeader(http.StatusInternalServerError)
    encoder.Encode(map[string]string{
      "status":"Error",
      "message":"Unexpected Error",
    })
    return
  }

  encoder.Encode(data)
}

func postCategory(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  userDto := models.UserCategoryDTO{}
  encoder := json.NewEncoder(w)

  if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
    w.WriteHeader(http.StatusNoContent)
    encoder.Encode(map[string]string{
      "status":"Error",
      "message":"Unexpected Error",
    })
    return
  }

  data := models.UserCategory{
    Name: userDto.Name,
    Message: userDto.Message,
  }
  database.Database.Save(&data)

  encoder.Encode(map[string]string{
    "status": "ok",
    "message": "Stored Successfully",
  })
}

func putCategory(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  vars := mux.Vars(r)
  encoder := json.NewEncoder(w)
  id, err := strconv.Atoi(vars["id"])
  
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    encoder.Encode(map[string]string{
      "status": "Error",
      "message": "Unexpected Error",
    })
    return
  }

  userDto := models.UserCategoryDTO{}
  if err:=json.NewDecoder(r.Body).Decode(&userDto); err != nil{
    w.WriteHeader(http.StatusBadRequest)
    encoder.Encode(map[string]string{
      "status": "Error",
      "message": "Unexpected Error",
    })
    return
  }

  data := models.UserCategory{}
  
  if err := database.Database.First(&data, id); err.Error != nil{
    w.WriteHeader(http.StatusBadRequest)
        encoder.Encode(map[string]string{
      "status": "Error",
      "message": "Unexpected Error",
    })
    return
  }

  data.Message = userDto.Message
  data.Name = userDto.Name

  database.Database.Save(&data)
  encoder.Encode(map[string]string{
    "status": "Ok",
    "message": "Stored Successfully",
  })
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  vars := mux.Vars(r)
  encoder := json.NewEncoder(w)
  id, err := strconv.Atoi(vars["id"])
  
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    encoder.Encode(map[string]string{
      "status": "Error",
      "message": "Unexpected Error",
    })
    return
  }

  data := models.UserCategory{}
  
  if err := database.Database.First(&data, id); err.Error != nil{
    w.WriteHeader(http.StatusBadRequest)
        encoder.Encode(map[string]string{
      "status": "Error",
      "message": "Unexpected Error",
    })
    return
  }

  database.Database.Delete(&data)
  encoder.Encode(map[string]string{
    "status": "Ok",
    "message": "Deleted Successfully",
  })
}
