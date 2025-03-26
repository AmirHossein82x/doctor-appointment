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
        "/otp/generate": {
            "post": {
                "description": "Generates an OTP for phone number verification",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OTP"
                ],
                "summary": "Generate OTP",
                "parameters": [
                    {
                        "description": "Phone Number",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.OTPRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/otp/verify": {
            "post": {
                "description": "Verifies the OTP entered by the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OTP"
                ],
                "summary": "Verify OTP",
                "parameters": [
                    {
                        "description": "Phone Number and OTP",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.VerifyOTPRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/users/register": {
            "post": {
                "description": "creating users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "registering user",
                "parameters": [
                    {
                        "description": "Phone Number",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserRegisterRequest"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "dto.OTPRequest": {
            "type": "object",
            "properties": {
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "dto.UserRegisterRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "verified_token": {
                    "type": "string"
                }
            }
        },
        "dto.VerifyOTPRequest": {
            "type": "object",
            "properties": {
                "otp_code": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
