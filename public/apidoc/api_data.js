define({ "api": [
  {
    "type": "post",
    "url": "/v1/auth/login",
    "title": "Login",
    "group": "Users",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "application/json",
            "optional": false,
            "field": "Content-Type",
            "description": "<p>Accept application/json</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>User username</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>User Password</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Input",
          "content": "   {\n     \"username\": \"your username\",\n\t\t\"password\"\t : \"your password\"\t\t\n   }",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "authenticate",
            "description": "<p>Response</p>"
          },
          {
            "group": "Success 200",
            "type": "Boolean",
            "optional": false,
            "field": "authenticate.success",
            "description": "<p>Status</p>"
          },
          {
            "group": "Success 200",
            "type": "Integer",
            "optional": false,
            "field": "authenticate.statuscode",
            "description": "<p>Status Code</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "authenticate.message",
            "description": "<p>Authenticate Message</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "authenticate.token",
            "description": "<p>Your JSON Token</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success",
          "content": "   {\n\t\t\"authenticate\": { \t\n\t \t \t \"statuscode\": 200,\n    \t\t \"success\": true,\n     \t \"message\": \"Login Successfully\",\n\t \t     \"token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ\"\n   \t\t}\n\t  }",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "List error",
          "content": "HTTP/1.1 500 Internal Server Error",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./routers/v1/auth_r.go",
    "groupTitle": "Users",
    "name": "PostV1AuthLogin"
  }
] });
