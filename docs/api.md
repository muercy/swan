| Method        | API                                         | Desc                         |
| ------------- |:--------------------------------------------|------------------------------|
| GET           | /v1/apps                                    | List all apps                |
| POST          | /v1/apps                                    | Create a app                 |
| GET           | /v1/apps/{app_id}                           | Inspect a app                |
| DELETE        | /v1/apps/{app_id}                           | Delete a app                 |
| POST          | /v1/apps/{app_id}/scale                     | Scale up/down                |
| PUT           | /v1/apps/{app_id}                           | Rolling update a app         |
| POST          | /v1/apps/{app_id}/rollback                  | Roll back a app              |
| GET           | /v1/apps/{app_id}/tasks                     | List all tasks for a app     |
| GET           | /v1/apps/{app_id}/tasks/{task_id}           | Inspect a task               |
| GET           | /v1/apps/{app_id}/versions                  | List all versions for a app  |
| GET           | /v1/apps/{app_id}/versions/{version_id}     | Inspect a version            |
| POST          | /v1/apps/{app_id}/versions                  | Create a version for a app   |
| PATCH         | /v1/apps/{app_id}/weights                   | Update tasks's weights       |
| PATCH         | /v1/apps/{app_id}/tasks/{task_id}/weight    | Update task's weight         |
| POST          | /v1/compose                                 | Create a compose             |
| GET           | /v1/compose                                 | List all compose             |
| GET           | /v1/compose/{compose_id}                    | Inspect a compose            |
| DELETE        | /v1/compose/{compose_id}                    | Delete a compose             |
| GET           | /ping                                       | Health check                 |
| GET           | /version                                    | Version information          |
| GET           | /v1/leader                                  | Inspect leader info          |
| GET           | /v1/events                                  | Event Subscription           |

#### List all apps
```
GET /v1/apps 
```
Example request:
```
GET /v1/apps
```
Example response:
```
[
  {
    "id": "nginx0r1.default.xcm.dataman",
    "name": "nginx0r1",
    "alias": "",
    "runAs": "xcm",
    "priority": 0,
    "cluster": "dataman",
    "operationStatus": "noop",
    "tasks": [
      {
        "id": "e6404f0324d2.0.nginx0r1.default.xcm.dataman",
        "name": "0.nginx0r1.default.xcm.dataman",
        "ip": "192.168.1.102",
        "port": 31008,
        "healthy": "unset",
        "weight": 100,
        "agentId": "7a40294e-b16b-4ac3-bbe4-1865df4a4705-S6",
        "version": "1498029948754163146",
        "status": "TASK_RUNNING",
        "errmsg": ""
        "created": "2017-06-21T15:25:48.78944685+08:00",
        "updated": "2017-06-21T15:25:48.78944688+08:00"
      }
    ],
    "currentVersion": [
      "1498029948754163146"
    ],
    "versions": [
      {
        "id": "1498029948754163146",
        "name": "nginx0r1",
        "cmd": "",
        "cpus": 0.01,
        "mem": 32,
        "disk": 0,
        "instances": 1,
        "runAs": "xcm",
        "priority": 100,
        "container": {
          "type": "DOCKER",
          "docker": {
            "image": "nginx",
            "network": "bridge",
            "parameters": [
              {
                "key": "label",
                "value": "APP_ID=wordpress"
              }
            ],
            "portMappings": [
              {
                "containerPort": 80,
                "hostPort": 80,
                "name": "web",
                "protocol": "tcp"
              }
            ],
            "privileged": true
          },
          "volumes": [
            {
              "containerPath": "/data",
              "hostPath": "/home",
              "mode": "RW"
            }
          ]
        },
        "labels": null,
        "healthCheck": null,
        "env": {},
        "killPolicy": null,
        "updatPolicy": null,
        "constraints": [],
        "uris": [],
        "ips": null,
        "proxy": {
          "enabled": false,
          "alias": ""
        }
      }
    ],
    "status": "available",
    "health": {
      "healthy": 0,
      "unhealthy": 0,
      "unset": 1
    },
    "created": "2017-06-21T15:25:48.754164732+08:00",
    "updated": "2017-06-21T15:25:48.754164753+08:00"
  }
]
```
Query parameters:
```
```

#### Create a app
```
POST /v1/apps
```
Example request:
```
 POST /v1/apps HTTP/1.1
 Content-Type: application/json
 
 {
  "name": "nginx002",
  "cmd": null,
  "args": null,
  "cpus": 0.01,
  "mem": 32,
  "disk": 0,
  "runAs": "xcm",
  "priority": 100,
  "instances": 10,
  "constraints": [],
  "container": {
    "docker": {
      "image": "nginx",
      "network": "bridge",
      "forcePullImage": false,
      "privileged": true,
      "parameters": [
        {
            "key": "label",
            "value": "APP_ID=wordpress"
        }
      ],
      "portMappings": [
              {
                      "name": "web",
                      "protocol": "tcp",
                      "containerPort": 80,
                      "hostPort": 80
              }
      ]
    },
    "type": "DOCKER",
    "volumes": [
      {
        "hostPath": "/home",
        "containerPath": "/data",
        "mode": "RW"
      }
  ]
  },
  "env": {
    "WORDPRESS_DB_HOST": "192.168.1.210",
    "WORDPRESS_DB_PASSWORD": "root"
  },
  "uris": [
  ],
  "label": {
    "USER_ID": "1"
  },
  "healthCheck":
    {
      "protocol": "http",
      "path": "/",
      "delaySeconds": 2,
      "gracePeriodSeconds": 5,
      "intervalSeconds": 1,
      "portName": "web",
      "timeoutSeconds": 1,
      "consecutiveFailures": 5
    },
  "proxy": {
            "enabled": false,
            "alias": ""
  }
}
```
Example response:
```
  HTTP/1.1 201 Created
  Content-Type: application/json

  {
       "Id":"nginx0r1.default.xcm.dataman"
  }
```
#### Inspect a app
```
GET /v1/apps/{app_id}
```
Example request:
``
GET /v1/apps/nginx0r2.default.xcm.dataman
```
Example response:
```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "id":"nginx0r2.default.xcm.dataman",
    "name":"nginx0r2",
    "alias":"",
    "runAs":"xcm",
    "priority":0,
    "cluster":"dataman",
    "operationStatus":"noop",
    "tasks":[
        {
            "id":"731ea4512976.0.nginx0r2.default.xcm.dataman",
            "name":"0.nginx0r2.default.xcm.dataman",
            "ip":"192.168.1.102",
            "port":31010,
            "healthy":"unset",
            "weight":100,
            "agentId":"7a40294e-b16b-4ac3-bbe4-1865df4a4705-S6",
            "version":"1498030396211326306",
            "status":"TASK_RUNNING",
            "errmsg":"",
            "created":"2017-06-21T15:33:16.238348516+08:00",
            "updated":"2017-06-21T15:33:16.238348626+08:00"
        }
    ],
    "currentVersion":[
        "1498030396211326306"
    ],
    "versions":[
        {
            "id":"1498030396211326306",
            "name":"nginx0r2",
            "cmd":"",
            "cpus":0.01,
            "mem":32,
            "disk":0,
            "instances":1,
            "runAs":"xcm",
            "priority":100,
            "container":{
                "type":"DOCKER",
                "docker":{
                    "image":"nginx",
                    "network":"bridge",
                    "parameters":[
                        {
                            "key":"label",
                            "value":"APP_ID=wordpress"
                        }
                    ],
                    "portMappings":[
                        {
                            "containerPort":80,
                            "hostPort":80,
                            "name":"web",
                            "protocol":"tcp"
                        }
                    ],
                    "privileged":true
                },
                "volumes":[
                    {
                        "containerPath":"/data",
                        "hostPath":"/home",
                        "mode":"RW"
                    }
                ]
            },
            "labels":null,
            "healthCheck":null,
            "env":{

            },
            "killPolicy":null,
            "updatPolicy":null,
            "constraints":[

            ],
            "uris":[

            ],
            "ips":null,
            "proxy":{
                "enabled":false,
                "alias":""
            }
        }
    ],
    "status":"available",
    "health":{
        "healthy":0,
        "unhealthy":0,
        "unset":1
    },
    "created":"2017-06-21T15:33:16.211327705+08:00",
    "updated":"2017-06-21T15:33:16.211327722+08:00"
}
```

#### Delete a app
```
DELETE /v1/apps/{app_id}
```
Example request:
```
DELETE /v1/apps/nginx0r2.default.xcm.dataman
```
Example response:
```
HTTP/1.1 204 No Content
```
##### Scale up/down
```
POST /v1/apps/{app_id}/scale
```
Example request:
```
POST /v1/apps/nginx0r2.default.xcm.dataman/scale HTTP/1.1
Content-Type: application/json
{
    "instances": 5, 
    "ips": [],
}
```
Json parameters:
```
instances     : the goal to scale up/down
ips(optional) : ip list for static ip(brige or host or scale down ignore)
```
Example response:
```
HTTP/1.1 202 Accepted
```

#### Update

First, create a new version for the app you plan to update.

```
POST /v1/apps/{app_id}/versions
```
Example request:
```
 POST /v1/apps/{nginx004.default.testuser.dataman}/versions HTTP/1.1
 Content-Type: application/json
 
 {
  "name": "nginx002",
  "cmd": null,
  "args": null,
  "cpus": 0.01,
  "mem": 32,
  "disk": 0,
  "runAs": "xcm",
  "priority": 100,
  "instances": 10,
  "constraints": [],
  "container": {
    "docker": {
      "image": "nginx",
      "network": "bridge",
      "forcePullImage": false,
      "privileged": true,
      "parameters": [
        {
            "key": "label",
            "value": "APP_ID=wordpress"
        }
      ],
      "portMappings": [
              {
                      "name": "web",
                      "protocol": "tcp",
                      "containerPort": 80,
                      "hostPort": 80
              }
      ]
    },
    "type": "DOCKER",
    "volumes": [
      {
        "hostPath": "/home",
        "containerPath": "/data",
        "mode": "RW"
      }
  ]
  },
  "env": {
    "WORDPRESS_DB_HOST": "192.168.1.210",
    "WORDPRESS_DB_PASSWORD": "root"
  },
  "uris": [
  ],
  "label": {
    "USER_ID": "1"
  },
  "healthCheck":
    {
      "protocol": "http",
      "path": "/",
      "delaySeconds": 2,
      "gracePeriodSeconds": 5,
      "intervalSeconds": 1,
      "portName": "web",
      "timeoutSeconds": 1,
      "consecutiveFailures": 5
    },
  "proxy": {
            "enabled": false,
            "alias": ""
  }
}
```
Example response:
```
  HTTP/1.1 201 Created
  Content-Type: application/json

  {
      "Id":"1498791358276219465"
  }
```
Second, process updating.

```
PUT /v1/apps/{app_id}
```
Example request:
```
PUT /v1/apps/nginx0r2.default.xcm.dataman HTTP/1.1
Content-Type: application/json
{
    "instances": 5, 
    "canary": {
        "enabled": false,
        "value": 0.1,
    }
}
```
Json parameters:
```
instances: the number of tasks to be updated.
canary:
    enabled: Disable or Enabled canary publish(gray publish). If Disabled, value is ignore. Default is Disabled.
    value:   Percentage of traffic for new version. 0.1 means ten percent. value is between (0, 1].
```
Example response:
```
HTTP/1.1 202 Accepted
```


```
During the update process, swan will be automatic find the new version you created previous step, and updated the old tasks to the new version
according you strategy.

If canary is enabled, the traffic control will be enabled. you can control how much traffic the new version is divided.
```


#### Roll back
```
POST /v1/apps/{app_id}/rollback
```
Example request:
```
POST /v1/apps/nginx0r2.default.xcm.dataman/rollback
```
```
Rollback will update app to the previous version.
```
Example response:
```
HTTP/1.1 202 Accepted
```

#### List all tasks for a app

```
GET /v1/apps/{app_id}/tasks
```
Example request:
```
GET /v1/apps/nginx0r1.default.xcm.dataman/tasks
```
Example response:
```json
HTTP/1.1 200 OK
Content-Type: application/json
[
  {
    "id": "e6404f0324d2.0.nginx0r1.default.xcm.dataman",
    "name": "0.nginx0r1.default.xcm.dataman",
    "ip": "192.168.1.102",
    "port": 31008,
    "healthy": "unset",
    "weight": 100,
    "agentId": "7a40294e-b16b-4ac3-bbe4-1865df4a4705-S6",
    "version": "1498029948754163146",
    "status": "TASK_RUNNING",
    "errmsg": "",
    "created": "2017-06-21T15:25:48.78944685+08:00",
    "updated": "2017-06-21T15:25:48.78944688+08:00"
  }
]
```

#### List all versions for a app
```
GET /v1/apps/{app_id}/versions
```
Example request:
```
GET /v1/apps/nginx0r1.default.xcm.dataman/versions
```
Example response:
```json
HTTP/1.1 200 OK
Content-Type: application/json
[
  {
    "id": "1498029948754163146",
    "name": "nginx0r1",
    "cmd": "",
    "cpus": 0.01,
    "mem": 32,
    "disk": 0,
    "instances": 1,
    "runAs": "xcm",
    "priority": 100,
    "container": {
      "type": "DOCKER",
      "docker": {
        "image": "nginx",
        "network": "bridge",
        "parameters": [
          {
            "key": "label",
            "value": "APP_ID=wordpress"
          }
        ],
        "portMappings": [
          {
            "containerPort": 80,
            "hostPort": 80,
            "name": "web",
            "protocol": "tcp"
          }
        ],
        "privileged": true
      },
      "volumes": [
        {
          "containerPath": "/data",
          "hostPath": "/home",
          "mode": "RW"
        }
      ]
    },
    "labels": null,
    "healthCheck": null,
    "env": {},
    "killPolicy": null,
    "updatPolicy": null,
    "constraints": [],
    "uris": [],
    "ips": null,
    "proxy": {
      "enabled": false,
      "alias": ""
    }
  }
]
```

#### Inspect a task

```
GET /v1/apps/{app_id}/tasks/{task_id}
```
Example request:
```
GET /v1/apps/nginx0r1.default.xcm.dataman/tasks/e6404f0324d2.0.nginx0r1.default.xcm.dataman
```
Example response:
```json
HTTP/1.1 200 OK
Content-Type: application/json
{
  "id": "e6404f0324d2.0.nginx0r1.default.xcm.dataman",
  "name": "0.nginx0r1.default.xcm.dataman",
  "ip": "192.168.1.102",
  "port": 31008,
  "healthy": "unset",
  "weight": 100,
  "agentId": "7a40294e-b16b-4ac3-bbe4-1865df4a4705-S6",
  "version": "1498029948754163146",
  "status": "TASK_RUNNING",
  "errmsg": "",
  "created": "2017-06-21T15:25:48.78944685+08:00",
  "updated": "2017-06-21T15:25:48.78944688+08:00"
}
```

#### Inspect a version 
```
GET /v1/apps/{app_id}/versions/{version_id}
```
Example request:
```
GET /v1/apps/nginx0r1.default.xcm.dataman/versions/1498029948754163146
```
Example response:
```json
HTTP/1.1 200 OK
Content-Type: application/json
{
  "id": "1498029948754163146",
  "name": "nginx0r1",
  "cmd": "",
  "cpus": 0.01,
  "mem": 32,
  "disk": 0,
  "instances": 1,
  "runAs": "xcm",
  "priority": 100,
  "container": {
    "type": "DOCKER",
    "docker": {
      "image": "nginx",
      "network": "bridge",
      "parameters": [
        {
          "key": "label",
          "value": "APP_ID=wordpress"
        }
      ],
      "portMappings": [
        {
          "containerPort": 80,
          "hostPort": 80,
          "name": "web",
          "protocol": "tcp"
        }
      ],
      "privileged": true
    },
    "volumes": [
      {
        "containerPath": "/data",
        "hostPath": "/home",
        "mode": "RW"
      }
    ]
  },
  "labels": null,
  "healthCheck": null,
  "env": {},
  "killPolicy": null,
  "updatPolicy": null,
  "constraints": [],
  "uris": [],
  "ips": null,
  "proxy": {
    "enabled": false,
    "alias": ""
  }
}
```
#### Create a version 
```
POST /v1/apps/{app_id}/versions
```
Example request:
```
 POST /v1/apps/{nginx004.default.testuser.dataman}/versions HTTP/1.1
 Content-Type: application/json
 
 {
  "name": "nginx002",
  "cmd": null,
  "args": null,
  "cpus": 0.01,
  "mem": 32,
  "disk": 0,
  "runAs": "xcm",
  "priority": 100,
  "instances": 10,
  "constraints": [],
  "container": {
    "docker": {
      "image": "nginx",
      "network": "bridge",
      "forcePullImage": false,
      "privileged": true,
      "parameters": [
        {
            "key": "label",
            "value": "APP_ID=wordpress"
        }
      ],
      "portMappings": [
              {
                      "name": "web",
                      "protocol": "tcp",
                      "containerPort": 80,
                      "hostPort": 80
              }
      ]
    },
    "type": "DOCKER",
    "volumes": [
      {
        "hostPath": "/home",
        "containerPath": "/data",
        "mode": "RW"
      }
  ]
  },
  "env": {
    "WORDPRESS_DB_HOST": "192.168.1.210",
    "WORDPRESS_DB_PASSWORD": "root"
  },
  "uris": [
  ],
  "label": {
    "USER_ID": "1"
  },
  "healthCheck":
    {
      "protocol": "http",
      "path": "/",
      "delaySeconds": 2,
      "gracePeriodSeconds": 5,
      "intervalSeconds": 1,
      "portName": "web",
      "timeoutSeconds": 1,
      "consecutiveFailures": 5
    },
  "proxy": {
            "enabled": false,
            "alias": ""
  }
}
```
Example response:
```
  HTTP/1.1 201 Created
  Content-Type: application/json

  {
      "Id":"1498791358276219465"
  }
```
#### Update weights 
```
PATCH /v1/apps/{app_id}/weights
```
Example request:
```
 PATCH /v1/apps/nginx004.default.testuser.dataman/weights HTTP/1.1
 Content-Type: application/json
 
 {
     Weights: {
         "0": 10,
         "1": 20,
         "3": 30,
     }
 }
```
Json parameters:
```
Weights:
    Key(string)  : task index, eg. 0, 1, 2...
    Value(float) : weight of task. value is between (0, 100].
```
Example response:
```
HTTP/1.1 202 Accepted
```

#### Update weight
```
PATCH /v1/apps/{app_id}/tasks/{task_id}/weight
```
Example request:
```
PATCH /v1/apps/nginx004.default.testuser.dataman}/tasks/0.nginx004.default.testuser.dataman/weight HTTP/1.1
Content-Type: application/json
{
    Weight: 50,
}
```
Json parameters:
```
Weight: the task weight
```
Example response:
```
HTTP/1.1 202 Accepted
```

#### Ping
```
GET /ping
```
Example request:
```
GET /ping
```
Example response:
```
"pong"
```

#### Version
```
GET /version
```

Example response:
```
{
  "version": "v0.1.11",
  "commit": "v0.1.11-153-g397f477-dirty",
  "build_time": "2017-05-27:11-27-19",
  "go_version": "go1.8.1",
  "os": "linux",
  "arch": "amd64"
}
```

#### Leader
```
GET /v1/leader
```

Example response:
```
{
    "leader": "192.168.1.92:5016"
}
```

