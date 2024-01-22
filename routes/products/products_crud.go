package product_routes

import (
	"encoding/json"
	"muxapi/database"
	"muxapi/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
)

func getProducts(w http.ResponseWriter, _ *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  data := models.Products{}
  database.Database.Order(
    clause.OrderByColumn{
      Column: clause.Column{Name: "id"}, 
      Desc: true,
    }).Preload("UserCategory").Find(&data)

  json.NewEncoder(w).Encode(data)
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
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
  var data models.Product
  if err := database.Database.Preload("UserCategory").First(&data, id); err.Error != nil {
    w.WriteHeader(http.StatusInternalServerError)
    encoder.Encode(map[string]string{
      "status":"Error",
      "message":"Unexpected Error",
    })
    return
  }

  encoder.Encode(data)
}

func postProduct(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  productDto := models.ProductDTO{}
  encoder := json.NewEncoder(w)

  if err := json.NewDecoder(r.Body).Decode(&productDto); err != nil {
    w.WriteHeader(http.StatusNoContent)
    encoder.Encode(map[string]string{
      "status":"Error",
      "message":"Unexpected Error",
    })
    return
  }

  data := models.Product{
    Name: productDto.Name,
    Description: productDto.Description,
    Price: productDto.Price,
    Stock: productDto.Stock,
    Category_id: productDto.Category_id,
    Date: time.Now(),
  }
  database.Database.Save(&data)

  encoder.Encode(map[string]string{
    "status": "ok",
    "message": "Stored Successfully",
  })
}

func putProduct(w http.ResponseWriter, r *http.Request) {
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

  productDto := models.ProductDTO{}
  if err:=json.NewDecoder(r.Body).Decode(&productDto); err != nil{
    w.WriteHeader(http.StatusBadRequest)
    encoder.Encode(map[string]string{
      "status": "Error",
      "message": "Unexpected Error",
    })
    return
  }

  data := models.Product{}
  
  if err := database.Database.First(&data, id); err.Error != nil{
    w.WriteHeader(http.StatusBadRequest)
        encoder.Encode(map[string]string{
      "status": "Error",
      "message": "Unexpected Error",
    })
    return
  }

  data.Name = productDto.Name
  data.Description = productDto.Description
  data.Price = productDto.Price
  data.Stock = productDto.Stock
  data.Category_id = productDto.Category_id

  database.Database.Save(&data)
  encoder.Encode(map[string]string{
    "status": "Ok",
    "message": "Stored Successfully",
  })
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
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

  data := models.Product{}
  
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
