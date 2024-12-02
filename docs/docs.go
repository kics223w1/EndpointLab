// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/kics223w1/EndpointLab"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/anything": {
            "get": {
                "description": "Return anything that is passed to the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Anything"
                ],
                "summary": "Returns anything that is passed to request",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "put": {
                "description": "Return anything that is passed to the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Anything"
                ],
                "summary": "Returns anything that is passed to request",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "description": "Return anything that is passed to the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Anything"
                ],
                "summary": "Returns anything that is passed to request",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Return anything that is passed to the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Anything"
                ],
                "summary": "Returns anything that is passed to request",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "options": {
                "description": "Return anything that is passed to the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Anything"
                ],
                "summary": "Returns anything that is passed to request",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "head": {
                "description": "Return anything that is passed to the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Anything"
                ],
                "summary": "Returns anything that is passed to request",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "patch": {
                "description": "Return anything that is passed to the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Anything"
                ],
                "summary": "Returns anything that is passed to request",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/cache": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response Inspection"
                ],
                "summary": "Returns a 304 if an If-Modified-Since header or If-None-Match is present. Returns the same as a GET otherwise.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Optional header to check if the resource has been modified since the specified date",
                        "name": "If-Modified-Since",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Optional header to check if the resource matches the given ETag",
                        "name": "If-None-Match",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Cached response"
                    },
                    "304": {
                        "description": "Modified"
                    }
                }
            }
        },
        "/cache/{value}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response Inspection"
                ],
                "summary": "Sets a Cache-Control header for n seconds.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Number of seconds for the Cache-Control max-age directive",
                        "name": "value",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Cache control set"
                    }
                }
            }
        },
        "/cookies": {
            "get": {
                "description": "Returns the cookies sent by the client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cookies"
                ],
                "summary": "Get cookies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/cookies/delete": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cookies"
                ],
                "summary": "Deletes cookie(s) as provided by the query string and redirects to cookie list.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Freeform cookie value",
                        "name": "freeform",
                        "in": "query"
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Redirects to cookie list",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cookies/set": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cookies"
                ],
                "summary": "Sets cookie(s) as provided by the query string and redirects to cookie list.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Freeform cookie value",
                        "name": "freeform",
                        "in": "query"
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Redirects to cookie list",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cookies/set/{name}/{value}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cookies"
                ],
                "summary": "Set a cookie with specified name and value",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cookie name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Cookie value",
                        "name": "value",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Redirects to cookie list",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/delete": {
            "delete": {
                "description": "Returns the DELETE parameters of the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP Methods"
                ],
                "summary": "The request's DELETE parameters.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/etag/{etag}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response Inspection"
                ],
                "summary": "Assumes the resource has the given etag and responds to If-None-Match and If-Match headers appropriately.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Optional header to check if the resource does not match the given ETag",
                        "name": "If-None-Match",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Optional header to check if the resource matches the given ETag",
                        "name": "If-Match",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Normal response"
                    },
                    "412": {
                        "description": "match"
                    }
                }
            }
        },
        "/get": {
            "get": {
                "description": "Returns the query parameters of the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP Methods"
                ],
                "summary": "The request's query parameters.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/headers": {
            "get": {
                "description": "Returns all headers of the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Request Inspection"
                ],
                "summary": "Returns the request headers.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/image": {
            "get": {
                "description": "Returns a simple image of the type suggested by the Accept header",
                "produces": [
                    "image/webp",
                    " image/png",
                    " image/jpeg",
                    " image/svg+xml"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "Returns an image based on the Accept header",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/image/jpeg": {
            "get": {
                "description": "Returns a simple JPEG image",
                "produces": [
                    "image/jpeg"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "Returns a JPEG image",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/image/png": {
            "get": {
                "description": "Returns a simple PNG image",
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "Returns a PNG image",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/image/svg": {
            "get": {
                "description": "Returns a simple SVG image",
                "produces": [
                    "image/svg+xml"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "Returns an SVG image",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/image/webp": {
            "get": {
                "description": "Returns a simple WEBP image",
                "produces": [
                    "image/webp"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "Returns a WEBP image",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/ip": {
            "get": {
                "description": "Returns the IP address of the client making the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Request Inspection"
                ],
                "summary": "Returns the client's IP address.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/patch": {
            "patch": {
                "description": "Returns the PATCH parameters of the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP Methods"
                ],
                "summary": "The request's PATCH parameters.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/post": {
            "post": {
                "description": "Returns the POST parameters of the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP Methods"
                ],
                "summary": "The request's POST parameters.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/put": {
            "put": {
                "description": "Returns the PUT parameters of the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP Methods"
                ],
                "summary": "The request's PUT parameters.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/response-headers": {
            "get": {
                "description": "Returns all response headers including freeform values.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response Inspection"
                ],
                "summary": "Returns response headers.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Freeform query parameter",
                        "name": "freeform",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response headers"
                    }
                }
            },
            "post": {
                "description": "Returns all response headers including freeform values.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response Inspection"
                ],
                "summary": "Returns response headers.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Freeform query parameter",
                        "name": "freeform",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response headers"
                    }
                }
            }
        },
        "/status/{code}": {
            "get": {
                "description": "Returns a status code based on the path",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Return status code or random status code if more than one are given",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "HTTP Status Code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "put": {
                "description": "Returns a status code based on the path",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Return status code or random status code if more than one are given",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "HTTP Status Code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "description": "Returns a status code based on the path",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Return status code or random status code if more than one are given",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "HTTP Status Code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Returns a status code based on the path",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Return status code or random status code if more than one are given",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "HTTP Status Code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "options": {
                "description": "Returns a status code based on the path",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Return status code or random status code if more than one are given",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "HTTP Status Code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "head": {
                "description": "Returns a status code based on the path",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Return status code or random status code if more than one are given",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "HTTP Status Code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "patch": {
                "description": "Returns a status code based on the path",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Return status code or random status code if more than one are given",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "HTTP Status Code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user-agent": {
            "get": {
                "description": "Returns the User-Agent string of the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Request Inspection"
                ],
                "summary": "Returns the User-Agent string.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "EndpointLab",
	Description:      "An alternative to httpbin.org in Gin.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
