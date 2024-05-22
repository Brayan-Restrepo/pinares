package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"pinares/config/variables"
)

func GetClient(clientKey string) variables.ClientConfig {
	clients := map[string]variables.ClientConfig{}

	return clients[clientKey]
}

func GetBasePath(client variables.ClientConfig) string {
	return client.Host + client.Prefix
}

func GetErrorResponse(err error) interface{} {
	return map[string]string{"error": err.Error()}
}

func GetStatusResponse(message string) interface{} {
	return map[string]string{"status": message}
}

func GetBody(c echo.Context) (*bytes.Buffer, error) {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error getting body: ", err)
		return nil, err
	}

	return bytes.NewBuffer(body), nil
}

func GetBodyValue(body *bytes.Buffer, key string) (interface{}, error) {
	var bodyMap map[string]interface{}
	err := json.Unmarshal(body.Bytes(), &bodyMap)
	if err != nil {
		log.Println("Error getting", key, "from body:", err)
		return "nil", err
	}

	value, ok := bodyMap[key]
	if !ok {
		return "nil", fmt.Errorf("Body field %s is required", key)
	}

	return value, nil
}

func StringifyQueryParams(c echo.Context) string {
	return "?" + c.QueryParams().Encode()
}

func GetQueryParam(c echo.Context, key string, required bool) (string, error) {
	value := c.QueryParam(key)
	if value == "" && required {
		return value, fmt.Errorf("Query param %s is required", key)
	}
	return value, nil
}

func CopyHeaders(c echo.Context, req *http.Request) {
	req.Header = c.Request().Header.Clone()
}

func CopyHeader(c echo.Context, req *http.Request, key string) {
	req.Header.Add(key, c.Request().Header.Get(key))
}

func GetHeader(c echo.Context, key string, required bool) (string, error) {
	value := c.Request().Header.Get(key)
	if value == "" && required {
		return value, fmt.Errorf("Header %s is required", key)
	}
	return value, nil
}

func CopyRequest(c echo.Context, body *bytes.Buffer, method, url string) (*http.Request, error) {
	queryParams := StringifyQueryParams(c)
	req, err := http.NewRequest(method, url+queryParams, body)
	if err != nil {
		log.Println("Error creating request to", url, ":", err)
		return nil, err
	}
	CopyHeaders(c, req)

	return req, nil
}
