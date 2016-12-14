define({ "api": [
  {
    "type": "get",
    "url": "/task/kasisemua/:project_id",
    "title": "Get All Task By Project Id",
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
            "type": "Integer",
            "optional": false,
            "field": "Project_Id",
            "description": "<p>Project ID</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Input",
          "content": "http://localhost:9090/task/kasisemua/1",
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
            "field": "content",
            "description": "<p>Data</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "content.title_task",
            "description": "<p>Judul Task</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "content.difficulty",
            "description": "<p>Tingkat Kesukaran Task</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "content.urgency",
            "description": "<p>Tingkat Kepentingan Task</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "content.desc_task",
            "description": "<p>Deskripsi Task</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "content.pemberi",
            "description": "<p>Pemeberi Task</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "content.penerima",
            "description": "<p>Penerima Task</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status_code",
            "description": "<p>Code Response</p>"
          },
          {
            "group": "Success 200",
            "type": "Boolean",
            "optional": false,
            "field": "succes",
            "description": "<p>Determine the Response</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success",
          "content": "{\n  \"content\": [\n    {\n      \"title_task\": \"Membuat Layout Login\",\n      \"difficulty\": \"4\",\n      \"urgency\": \"5\",\n      \"desc_task\": \"Create Login Layout\",\n      \"pemberi\": \"Iman\",\n      \"penerima\": \"Eminarty \"\n    }\n  ],\n  \"status_code\": 200,\n  \"success\": true\n}",
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
    "filename": "./routers/users/list-task-by-project.go",
    "groupTitle": "Users",
    "name": "GetTaskKasisemuaProject_id"
  },
  {
    "type": "post",
    "url": "/authuser/login",
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
    "filename": "./routers/user_auth/auth_r.go",
    "groupTitle": "Users",
    "name": "PostAuthuserLogin"
  },
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./public/apidoc/main.js",
    "group": "_home_ww3_gocode_src_simple_api_public_apidoc_main_js",
    "groupTitle": "_home_ww3_gocode_src_simple_api_public_apidoc_main_js",
    "name": ""
  }
] });
