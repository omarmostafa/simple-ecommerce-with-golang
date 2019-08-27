package controllers

import (
	"gopkg.in/matryer/respond.v1"
	"net/http"
)

func Respond(res http.ResponseWriter, req *http.Request,status_code int,status bool,message string, data interface{})  {
	response :=map[string]interface{} {"status" : status, "message" : message,"data" : data}
	respond.With(res, req, status_code, response)
}

func RespondAccepted(res http.ResponseWriter, req *http.Request,message string,data interface{})  {
	Respond(res,req,http.StatusAccepted,true,message,data)
}

func RespondCreated(res http.ResponseWriter, req *http.Request,message string,data interface{})  {
	Respond(res,req,http.StatusCreated,true,message,data)
}

func RespondNotFound(res http.ResponseWriter, req *http.Request,message string,data interface{})  {
	Respond(res,req,http.StatusNotFound,false,message,data)
}

func RespondBadRequest(res http.ResponseWriter, req *http.Request,message string,data interface{})  {
	Respond(res,req,http.StatusBadRequest,false,message,data)
}

func RespondNotAuthorized(res http.ResponseWriter, req *http.Request,message string,data interface{})  {
	Respond(res,req,http.StatusUnauthorized,false,message,data)
}

func RespondForbidden(res http.ResponseWriter, req *http.Request,message string,data interface{})  {
	Respond(res,req,http.StatusForbidden,false,message,data)
}



