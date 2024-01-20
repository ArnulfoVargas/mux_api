package example_routes

import (
  "encoding/json"
  "io"
  "muxapi/dto"
  "muxapi/models"
  "net/http"
  "os"
  "strings"
  "time"

  "github.com/gorilla/mux"
)

func HandleExample(mux *mux.Router){
  prefix := "/api/example"

  mux.HandleFunc(prefix + "/", exampleGet).Methods(http.MethodGet)
  mux.HandleFunc(prefix + "/", examplePost).Methods(http.MethodPost)
  mux.HandleFunc(prefix + "/{id:[0-9]+}", examplePut).Methods(http.MethodPut)
  mux.HandleFunc(prefix + "/{id:[0-9]+}", exampleDelete).Methods(http.MethodDelete)
  mux.HandleFunc(prefix + "/{id:[0-9]+}", exampleParamsGet).Methods(http.MethodGet)
  mux.HandleFunc(prefix + "/query", exampleQueryGet).Methods(http.MethodGet)
  mux.HandleFunc(prefix + "/file", exampleFileUpload).Methods(http.MethodPost)
  mux.HandleFunc(prefix + "/file", exampleFilePreview).Methods(http.MethodGet)
}

func exampleGet(w http.ResponseWriter, _ *http.Request){
  response := models.GenericResponse{ Status: 200, Message: "Method Get"}
  jsonResponse, err := json.Marshal(response)

  if err != nil {
    response.Status = 505
    response.Message = "Couldnt Parse"
  }
  w.Header().Set("Content-Type", "application/json")
  w.Header().Add("Arnulfo", "https://arnulfo-vue-portfolio.netlify.app/")
  w.Write(jsonResponse)
}

func examplePost(w http.ResponseWriter, r *http.Request){
  response := models.GenericResponse{ Status: 200, Message: "Method Post"}

  var userDto dto.ExampleUserDto
  err := json.NewDecoder(r.Body).Decode(&userDto)

  auth := r.Header.Get("Authorization")
  if err != nil {
    response.Status = http.StatusNoContent
    response.Message = "No data in body"
  }

  output := map[string]interface{} {
    "status" : response.Status,
    "message": response.Message,
    "name"   : userDto.Name,
    "auth"   : auth,
  }

  jsonResponse, err := json.Marshal(output)

  if err != nil {
    response.Status = 505
    response.Message = response.Message + "Couldnt Parse"
  }

  w.Header().Set("Content-Type", "application/json")
  w.Header().Add("Arnulfo", "https://arnulfo-vue-portfolio.netlify.app/")
  w.WriteHeader(http.StatusCreated)

  w.Write(jsonResponse)
}

func examplePut(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  response := models.GenericResponse{ Status: 200, Message: "Method Put | Id: " + vars["id"]}
  jsonResponse, err := json.Marshal(response)

  if err != nil {
    response.Status = 505
    response.Message = "Couldnt Parse"
  }
  w.Header().Set("Content-Type", "application/json")
  w.Header().Add("Arnulfo", "https://arnulfo-vue-portfolio.netlify.app/")
  w.Write(jsonResponse)
}

func exampleDelete(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  response := models.GenericResponse{ Status: 200, Message: "Method Delete | Id: " + vars["id"]}
  jsonResponse, err := json.Marshal(response)

  if err != nil {
    response.Status = 505
    response.Message = "Couldnt Parse"
  }
  w.Header().Set("Content-Type", "application/json")
  w.Header().Add("Arnulfo", "https://arnulfo-vue-portfolio.netlify.app/")
  w.Write(jsonResponse)
}
func exampleParamsGet(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  response := models.GenericResponse{ Status: 200, Message: "Method Get Params | Id: " + vars["id"]}
  jsonResponse, err := json.Marshal(response)

  if err != nil {
    response.Status = 505
    response.Message = "Couldnt Parse"
  }
  w.Header().Set("Content-Type", "application/json")
  w.Header().Add("Arnulfo", "https://arnulfo-vue-portfolio.netlify.app/")
  w.Write(jsonResponse)
}
func exampleQueryGet(w http.ResponseWriter, r *http.Request){
  query := r.URL.Query()

  if query.Has("id"){
    response := models.GenericResponse{ Status: 200, Message: "Method Get Query | Id: " + query.Get("id")}
    jsonResponse, err := json.Marshal(response)

    if err != nil {
      response.Status = 505
      response.Message = "Couldnt Parse"
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Add("Arnulfo", "https://arnulfo-vue-portfolio.netlify.app/")
    w.Write(jsonResponse)

  } else {
    w.Write([]byte("Query does not contain id"))
  }
}

func exampleFileUpload(w http.ResponseWriter, r *http.Request){
  file, handler, err := r.FormFile("image")
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  nameArr := strings.Split(handler.Filename, ".")
  extension := nameArr[len(nameArr) - 1]
  currentTime := strings.Split(time.Now().String(), " ")
  fileName := string(currentTime[4][6:14]) + "." + extension

  fileDir := "public/uploads/images/" + fileName

  receiver, err := os.OpenFile(fileDir, os.O_WRONLY | os.O_CREATE, 0777)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  _, err = io.Copy(receiver, file)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  response := models.GenericResponse{
    Status: http.StatusOK,
    Message: "File uploaded",
  }

  jsonResponse, err := json.Marshal(response)
  if err != nil {
    w.WriteHeader(http.StatusTeapot)
    w.Write([]byte("Error sending a request, File succesfully uploaded"))
    return
  }

  w.Write(jsonResponse)
}

func exampleFilePreview(w http.ResponseWriter, r *http.Request){
  query := r.URL.Query()
  
  if !query.Has("file") {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  fileName := query.Get("file")

  file, err := os.Open("public/uploads/images/" + fileName)

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  
  _, err = io.Copy(w, file)
  
  if err != nil {
    http.Error(w, "Unexpected error", http.StatusNotFound)
  }
}
