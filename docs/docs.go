// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Get The Server Status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Root"
                ],
                "summary": "Show The Status of The Server",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/auth": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get Authentication Token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Root"
                ],
                "summary": "Generate Authentication Token",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/group": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get Joined Groups Information from WhatsApp",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Group"
                ],
                "summary": "Get Joined Groups Information",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/group/join": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Joining to Group From Invitation Link from WhatsApp",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Group"
                ],
                "summary": "Join Group From Invitation Link",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group Invitation Link",
                        "name": "link",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/group/leave": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Leaving Group By Group ID from WhatsApp",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Group"
                ],
                "summary": "Leave Group By Group ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group ID",
                        "name": "groupid",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/login": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get QR Code for WhatsApp Multi-Device Login",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json",
                    "text/html"
                ],
                "tags": [
                    "WhatsApp Authentication"
                ],
                "summary": "Generate QR Code for WhatsApp Multi-Device Login",
                "parameters": [
                    {
                        "enum": [
                            "html",
                            "json"
                        ],
                        "type": "string",
                        "default": "html",
                        "description": "Change Output Format in HTML or JSON",
                        "name": "output",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/login/pair": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get Pairing Code for WhatsApp Multi-Device Login",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Authentication"
                ],
                "summary": "Pair Phone for WhatsApp Multi-Device Login",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Make Device Logout from WhatsApp Multi-Device",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Authentication"
                ],
                "summary": "Logout Device from WhatsApp Multi-Device",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/message/delete": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete Message to Spesific WhatsApp Personal ID or Group ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Message"
                ],
                "summary": "Delete Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination WhatsApp Personal ID or Group ID",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Message ID",
                        "name": "messageid",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/message/edit": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update Message to Spesific WhatsApp Personal ID or Group ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Message"
                ],
                "summary": "Update Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination WhatsApp Personal ID or Group ID",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Message ID",
                        "name": "messageid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Text Message",
                        "name": "message",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/message/react": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "React Message to Spesific WhatsApp Personal ID or Group ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Message"
                ],
                "summary": "React Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination WhatsApp Personal ID or Group ID",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Message ID",
                        "name": "messageid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Reaction Emoji",
                        "name": "emoji",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/registered": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Check WhatsApp Personal ID is Registered",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Information"
                ],
                "summary": "Check If WhatsApp Personal ID is Registered",
                "parameters": [
                    {
                        "type": "string",
                        "description": "WhatsApp Personal ID to Check",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/send/audio": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Audio Message to Spesific WhatsApp Personal ID or Group ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Send Message"
                ],
                "summary": "Send Audio Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination WhatsApp Personal ID or Group ID",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Audio File",
                        "name": "audio",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/send/contact": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Contact Message to Spesific WhatsApp Personal ID or Group ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Send Message"
                ],
                "summary": "Send Contact Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination WhatsApp Personal ID or Group ID",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Contact Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Contact Phone",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/send/document": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Document Message to Spesific WhatsApp Personal ID or Group ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Send Message"
                ],
                "summary": "Send Document Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination WhatsApp Personal ID or Group ID",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Document File",
                        "name": "document",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/send/image": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Image Message to Spesific WhatsApp Personal ID or Group ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Send Message"
                ],
                "summary": "Send Image Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination WhatsApp Personal ID or Group ID",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Caption Image Message",
                        "name": "caption",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Image File",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "default": false,
                        "description": "Is View Once",
                        "name": "viewonce",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/send/link": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Link Message to Spesific WhatsApp Personal ID or Group ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Send Message"
                ],
                "summary": "Send Link Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination WhatsApp Personal ID or Group ID",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Link Caption",
                        "name": "caption",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Link URL",
                        "name": "url",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/send/location": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Location Message to Spesific WhatsApp Personal ID or Group ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Send Message"
                ],
                "summary": "Send Location Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination WhatsApp Personal ID or Group ID",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Location Latitude",
                        "name": "latitude",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Location Longitude",
                        "name": "longitude",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/send/poll": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Poll to Spesific WhatsApp Personal ID or Group ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Send Message"
                ],
                "summary": "Send Poll",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination WhatsApp Personal ID or Group ID",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Poll Question",
                        "name": "question",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Poll Options (Comma Seperated for New Options)",
                        "name": "options",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "default": false,
                        "description": "Is Multiple Answer",
                        "name": "multianswer",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/send/sticker": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Sticker Message to Spesific WhatsApp Personal ID or Group ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Send Message"
                ],
                "summary": "Send Sticker Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination WhatsApp Personal ID or Group ID",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Sticker File",
                        "name": "sticker",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/send/text": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Text Message to Spesific WhatsApp Personal ID or Group ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Send Message"
                ],
                "summary": "Send Text Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination WhatsApp Personal ID or Group ID",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Text Message",
                        "name": "message",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/send/video": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send Video Message to Spesific WhatsApp Personal ID or Group ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Send Message"
                ],
                "summary": "Send Video Message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Destination WhatsApp Personal ID or Group ID",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Caption Video Message",
                        "name": "caption",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Video File",
                        "name": "video",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "default": false,
                        "description": "Is View Once",
                        "name": "viewonce",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        },
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.x",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Go WhatsApp Multi-Device REST API",
	Description:      "This is WhatsApp Multi-Device Implementation in Go REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
