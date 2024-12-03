package api

import (
	"encoding/json"
	"endpointlab/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpMethod struct {

}

type HTTPMethodResponse struct {
	Args    map[string]string `json:"args"`
	Data    string             `json:"data,omitempty"`
	Files   map[string]string  `json:"files,omitempty"`
	Form    map[string]string  `json:"form,omitempty"`
	Headers map[string]string `json:"headers"`
	JSON    interface{}        `json:"json,omitempty"`
	Origin  string             `json:"origin"`
	URL     string             `json:"url"`
	Method  string             `json:"method"`
}

func NewHttpMethod() *HttpMethod {
	return &HttpMethod{}
}	

func parseJSONIfApplicable(headers http.Header, rawBody []byte) interface{} {
	// Convert the raw data to a string
	bodyString := string(rawBody)
	fmt.Printf("Request Body as String: %s\n", bodyString)

	// Check if the Content-Type is application/json
	if headers.Get("Content-Type") == "application/json" {
		var jsonData interface{}
		// Unmarshal the JSON data from the string
		if err := json.Unmarshal([]byte(bodyString), &jsonData); err == nil {
			return jsonData
		} else {
			fmt.Println("Failed to unmarshal JSON:", err)
		}
	}
	return nil
}

//	@Summary		The request's query parameters.
//	@Description	Returns the query parameters of the request
//	@Tags			HTTP Methods
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object
//	@Router			/get [get]
func (h *HttpMethod) HandleGet(c *gin.Context) {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	fullURL := scheme + "://" + c.Request.Host + c.Request.URL.String()

	response := HTTPMethodResponse{
		Args:    utils.ConvertQuery(c.Request.URL.Query()),
		Headers: utils.ConvertHeaders(c.Request.Header),
		JSON:    nil,
		Origin:  c.ClientIP(),
		URL:     fullURL,
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
}

//	@Summary		The request's POST parameters.
//	@Description	Returns the POST parameters of the request
//	@Tags			HTTP Methods
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object
//	@Router			/post [post]
func (h *HttpMethod) HandlePost(c *gin.Context) {
	// Read body data
	bodyData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	// Parse JSON if applicable
	jsonData := parseJSONIfApplicable(c.Request.Header, bodyData)

	fmt.Printf("jsonData: %v\n", jsonData)

	response := HTTPMethodResponse{
		Args:    utils.ConvertQuery(c.Request.URL.Query()),
		Data:    string(bodyData),  // Convert bodyData to string
		Files:   make(map[string]string),
		Form:    make(map[string]string),
		Headers: utils.ConvertHeaders(c.Request.Header),
		JSON:    jsonData,  // Use the successfully bound JSON data, if any
		Origin:  c.ClientIP(),
		URL:     c.Request.URL.String(),
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
}

//	@Summary		The request's PUT parameters.
//	@Description	Returns the PUT parameters of the request
//	@Tags			HTTP Methods
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object
//	@Router			/put [put]
func (h *HttpMethod) HandlePut(c *gin.Context) {
	response := HTTPMethodResponse{
		Args:    utils.ConvertQuery(c.Request.URL.Query()),
		Data:    "",
		Files:   make(map[string]string),
		Form:    make(map[string]string),
		Headers: utils.ConvertHeaders(c.Request.Header),
		JSON:    nil,
		Origin:  c.ClientIP(),
		URL:     c.Request.URL.String(),
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
}

//	@Summary		The request's DELETE parameters.
//	@Description	Returns the DELETE parameters of the request
//	@Tags			HTTP Methods
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object
//	@Router			/delete [delete]
func (h *HttpMethod) HandleDelete(c *gin.Context) {
	response := HTTPMethodResponse{
		Args:   utils.ConvertQuery(c.Request.URL.Query()),
		Data:    "",
		Files:   make(map[string]string),
		Form:    make(map[string]string),
		Headers: utils.ConvertHeaders(c.Request.Header),
		JSON:    nil,
		Origin:  c.ClientIP(),
		URL:     c.Request.URL.String(),
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
}

//	@Summary		The request's PATCH parameters.
//	@Description	Returns the PATCH parameters of the request
//	@Tags			HTTP Methods
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object
//	@Router			/patch [patch]
func (h *HttpMethod) HandlePatch(c *gin.Context) {
	response := HTTPMethodResponse{
		Args:   utils.ConvertQuery(c.Request.URL.Query()),
		Data:    "",
		Files:   make(map[string]string),
		Form:    make(map[string]string),
		Headers: utils.ConvertHeaders(c.Request.Header),
		JSON:    nil,
		Origin:  c.ClientIP(),
		URL:     c.Request.URL.String(),
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
}

