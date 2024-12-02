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
