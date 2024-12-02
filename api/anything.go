package api

import (
	"encoding/base64"
	"endpointlab/utils"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

type HttpAnything struct {}

type httpAnythingResponse struct {
	Args    map[string]string `json:"args"`
	Data    string             `json:"data"`
	Files   map[string]string  `json:"files"`
	Form    map[string]string  `json:"form"`
	Headers map[string]string `json:"headers"`
	JSON    interface{}        `json:"json"`
	Origin  string             `json:"origin"`
	URL     string             `json:"url"`
	Method  string             `json:"method"`
}

func NewHttpAnything() *HttpAnything {
	return &HttpAnything{}
}


//	@Summary		Returns anything that is passed to request
//	@Description	Return anything that is passed to the request
//	@Tags			Anything
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object
//	@Router			/anything [get]
//	@Router			/anything [post]
//	@Router			/anything [put]
//	@Router			/anything [delete]
//	@Router			/anything [patch]
//	@Router			/anything [options]
//	@Router			/anything [head]
func (h *HttpAnything) HandleAnything(c *gin.Context) {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	fullURL := scheme + "://" + c.Request.Host + c.Request.URL.String()
	
	response := httpAnythingResponse{
		Args:    utils.ConvertQuery(c.Request.URL.Query()),
		Data:    getData(c.Request.Body, c),
		Files:   getFiles(c.Request.Body, c),
		Form:    getForm(c.Request.Body, c),
		Headers: utils.ConvertHeaders(c.Request.Header),
		Method:  c.Request.Method,
		Origin:  c.Request.RemoteAddr,
		URL:     fullURL,
	}

	c.JSON(200, response)
}

func getData(body io.ReadCloser , c *gin.Context) string {
	if body == nil {
		return ""
	}
	
	bytes, err := io.ReadAll(body)
	if err != nil {
		return ""
	}

	contentType := c.GetHeader("Content-Type")
	if strings.Contains(contentType, "multipart/form-data") {
		return ""
	}

	if strings.Contains(contentType, "application/x-www-form-urlencoded") {
		return ""
	}

	// Check if the content is binary by looking for null bytes or non-printable characters
	isBinary := false
	for _, b := range bytes {
		if b == 0 || (b < 32 && b != '\n' && b != '\r' && b != '\t') {
			isBinary = true
			break
		}
	}

	if isBinary {
		return "data:" + contentType + ";base64," + base64.StdEncoding.EncodeToString(bytes)
	}
	
	return string(bytes)
}

func getFiles(body io.ReadCloser, c *gin.Context) map[string]string {
	contentType := c.GetHeader("Content-Type")
	if !strings.Contains(contentType, "multipart/form-data") {
		return nil
	}

	// Parse multipart form with a reasonable max memory
	err := c.Request.ParseMultipartForm(32 << 20) // 32MB max memory
	if err != nil {
		return nil
	}

	files := make(map[string]string)
	if c.Request.MultipartForm == nil || c.Request.MultipartForm.File == nil {
		return nil
	}

	// Process each uploaded file
	for fieldName, fileHeaders := range c.Request.MultipartForm.File {
		// Only process the first file for each field name
		if len(fileHeaders) > 0 {
			fileHeader := fileHeaders[0]
			file, err := fileHeader.Open()
			if err != nil {
				continue
			}
			defer file.Close()

			// Read file content
			content, err := io.ReadAll(file)
			if err != nil {
				continue
			}

			// Get the content type from the file header
			contentType := fileHeader.Header.Get("Content-Type")
			if contentType == "" {
				contentType = "application/octet-stream"
			}

			// Store base64 encoded content with content-type prefix
			files[fieldName] = "data:" + contentType + ";base64," + base64.StdEncoding.EncodeToString(content)
		}
	}

	return files
}

func getForm(body io.ReadCloser, c *gin.Context) map[string]string {
	contentType := c.GetHeader("Content-Type")
	
	// Handle application/x-www-form-urlencoded
	if strings.Contains(contentType, "application/x-www-form-urlencoded") {
		err := c.Request.ParseForm()
		if err != nil {
			return nil
		}
	
		utils.Log.Printf("PostForm ne: %+v", c.Request)

		return utils.ConvertQuery(c.Request.PostForm)
	}

	// Handle multipart/form-data
	if strings.Contains(contentType, "multipart/form-data") {
		if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
			return nil
		}
		if c.Request.MultipartForm == nil {
			return nil
		}
		formData := make(map[string]string)
		for key, values := range c.Request.MultipartForm.Value {
			if len(values) > 0 {
				formData[key] = values[0]
			}
		}
		return formData
	}

	return nil
}