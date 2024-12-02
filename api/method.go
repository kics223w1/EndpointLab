package api

import (
	"endpointlab/utils"
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

//	@Summary		Get request returns a simple GET response
//	@Description	Returns a simple GET response
//	@Tags			http-methods
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


//	@Summary		Post request returns a POST response
//	@Description	Returns a POST response
//	@Tags			http-methods
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object
//	@Router			/post [post]
func (h *HttpMethod) HandlePost(c *gin.Context) {
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

//	@Summary		Put request returns a PUT response
//	@Description	Returns a PUT response
//	@Tags			http-methods
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

//	@Summary		Delete request returns a DELETE response
//	@Description	Returns a DELETE response
//	@Tags			http-methods
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

//	@Summary		Patch request returns a PATCH response
//	@Description	Returns a PATCH response
//	@Tags			http-methods
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

