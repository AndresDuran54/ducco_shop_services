package network

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type (
	HttpClientRequest struct {
		Method       string
		Url          string
		InsecureSkip bool
		Headers      map[string]string
		Params       map[string]string
		Data         interface{}
	}

	HttpClientResponse struct {
		Status    int
		StatusMsg string
		Success   bool
		Url       string
		Headers   map[string]string
		Request   HttpClientRequest
		Data      map[string]interface{}
		Error     error
	}
)

type HttpClient struct {
}

func (obj HttpClient) Call(data HttpClientRequest) HttpClientResponse {
	//+ Cuerpo de la respuesta
	var body []byte

	//+ Request
	var request *http.Request

	//+ Variable para almacenar el error
	var err error

	//+ Clonamos el transporte HTTP predeterminado
	transportHttp := http.DefaultTransport.(*http.Transport).Clone()

	//+ Deshabilitamos keep-alive
	transportHttp.DisableKeepAlives = true

	//+ Deshabilitamos o no la verificación SSL en las solicitudes
	transportHttp.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: data.InsecureSkip,
	}

	//+ Establecemos el transporte HTTP en el cliente
	clientHttp := &http.Client{
		Transport: transportHttp,
	}

	//+ Agregamos el body en la request
	if data.Data != nil {

		body, err = json.Marshal(data.Data)

		if err != nil {
			return HttpClientResponse{Error: err, Request: data}
		}

		request, err = http.NewRequest(data.Method, data.Url, bytes.NewBuffer(body))

		if request.Body != nil {
			defer request.Body.Close()
		}
	} else {
		request, err = http.NewRequest(data.Method, data.Url, nil)
	}

	//+ En caso de error retornamos
	if err != nil {
		return HttpClientResponse{Error: err, Request: data}
	}

	//+ Agregamos los headers en la request
	for key, value := range data.Headers {
		request.Header.Add(key, value)
	}

	//+ Agregamos el cabecero Content-Type
	request.Header.Add("Content-Type", "application/json")

	//+ Agregamos los query params
	query := request.URL.Query()
	for key, value := range data.Params {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()

	//+ Url final de la request
	finalUrl := request.URL.String()

	//+ Ejecutamos la request
	resultHttp, err := clientHttp.Do(request)

	//+ En caso de error retornamos
	if err != nil {
		return HttpClientResponse{Error: err, Request: data, Url: finalUrl}
	}

	//+ Finalizamos el buffer al finalizar la ejecución de la función
	defer resultHttp.Body.Close()

	//+ Obtenemos el body de la response
	body, err = ioutil.ReadAll(resultHttp.Body)

	//+ En caso de error retornamos
	if err != nil {
		return HttpClientResponse{Error: err, Request: data, Url: finalUrl}
	}

	//+ Construimos la respuesta
	response := HttpClientResponse{
		Error:   nil,
		Request: data,
		Headers: map[string]string{},
		Url:     finalUrl,
	}

	//+ Parseamos el body de la respuesta
	err = json.Unmarshal(body, &response.Data)

	//+ En caso de error retornamos
	if err != nil {
		response.Error = err
		return response
	}

	//+ Establecemos el éxito de la respuesta
	response.Success = resultHttp.StatusCode >= 200 && resultHttp.StatusCode < 300

	//+ Código de estado
	response.Status = resultHttp.StatusCode

	//+ Mensaje de la respuesta
	response.StatusMsg = resultHttp.Status

	//+ Obtenemos los headers de la response
	for key := range resultHttp.Header {
		response.Headers[key] = resultHttp.Header.Get(key)
	}

	return response
}
