package pictures_routes

import (
	"encoding/json"
	"io"
	"muxapi/database"
	"muxapi/models"
	"muxapi/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func getPictureByProd(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  vars := mux.Vars(r)
  idStr := vars["id"]
  var id uint

  if val, err := strconv.Atoi(idStr); err != nil {
    utils.EncodeGenericResponse(w, "Error", "Unexpected Error")
    return
  } else {
    id = uint(val)
  }

  product := models.Product{}

  if err := database.Database.First(&product, id); err.Error  != nil{
    utils.EncodeGenericResponse(w, "Error", "Unexpected Error")
    return
  }

  data := models.Pictures{}

  if err := database.Database.Preload("Product.UserCategory").Where("product_id=?", id).Find(&data); err.Error != nil {
    utils.EncodeGenericResponse(w, "Error", "Unexpected Error")
    return
  }

  json.NewEncoder(w).Encode(data)
}

func postPicture(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  vars := mux.Vars(r)
  idStr := vars["id"]
  var id uint 

  if v, err := strconv.Atoi(idStr); err != nil {
    utils.EncodeGenericResponse(w, "Error", "Unexpected Error")
    return
  } else {
    id = uint(v)
  }

  file, handler, err := r.FormFile("image")
  if err != nil {
    utils.EncodeGenericResponse(w, "Error", "Unexpected Error")
    return
  }
  nameArr := strings.Split(handler.Filename, ".")
  extension := nameArr[len(nameArr) - 1]
  currentTime := strings.Split(time.Now().String(), " ")
  fileName := string(currentTime[4][6:14]) + "." + extension

  fileDir := "public/uploads/products/" + fileName

  receiver, err := os.OpenFile(fileDir, os.O_WRONLY | os.O_CREATE, 0777)

  if err != nil {
    utils.EncodeGenericResponse(w, "Error", "Unexpected Error")
    return
  }

  _, err = io.Copy(receiver, file)

  if err != nil {
    utils.EncodeGenericResponse(w, "Error", "Unexpected Error")
    return
  }

  data := models.Picture{
    Name: fileName,
    ProductId: id,
  }
  database.Database.Save(&data)
  utils.EncodeGenericResponse(w, "Ok", "Stored Successfully")
}

func deletePicture(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  idStr := vars["id"]
  id, err := strconv.Atoi(idStr)

  if err != nil {
    utils.EncodeGenericResponse(w, "Error", "Unexpected Error")
    return
  }

  data := models.Picture{}

  if err := database.Database.Find(&data, id); err.Error != nil {
    utils.EncodeGenericResponse(w, "Error", "Unexpected Error")
    return
  }

  err = os.Remove("public/uploads/products/" + data.Name)

  if err != nil {
    utils.EncodeGenericResponse(w, "Error", "Error Deleting")
    return
  }

  database.Database.Delete(&data)
  utils.EncodeGenericResponse(w, "Ok", "Deleted Successfully")
}
