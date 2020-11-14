## GREEDY-GAME-TEST-API

## Command to run the API - docker-compose-up

### API Documentation

#### 1. Insert API

    - Endpoint: /v1/insert
    - Method: POST
    - Request body (Test cases)
      1.  {
                "dim": [{
                        "key": "device",
                        "val": "mobile"
                },
                {
                    "key": "country",
                    "val": "IN"
                }],
                "metrics": [{
                    "key": "webreq",
                    "val": 70
                },
                {
                    "key": "timespent",
                    "val": 30
                }]
          }


      2.  {
                "dim": [{
                        "key": "device",
                        "val": "web"
                },
                {
                    "key": "country",
                    "val": "IN"
                }],
                "metrics": [{
                    "key": "webreq",
                    "val": 50
                },
                {
                    "key": "timespent",
                    "val": 30
                }]
          }
    - Response : 200 OK
    - Response Body
      {
            "message": "success"
      }

#### 2. Query API

    - Endpoint: /v1/query
    - Method: GET
    - Request Body (Test case)
       {
            "dim" : [{
                "key" : "country",
                "val" : "IN"
            }]
       }
    - Response : 200 OK
    - Response Body
      {
            "dim": [
                {
                    "key": "country",
                    "val": "IN"
                }
            ],
            "metrics": [
                {
                    "key": "webreq",
                    "val": 120
                },
                {
                    "key": "timespent",
                    "val": 60
                }
            ]
     }
