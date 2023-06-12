# performance_go_python_api
Performance test : python HTTP vs Go HTTP


### MySQL Server 
- [#issue](https://github.com/robert-min/performance_go_python_api/issues/2)


### Go HTTP Server
- [#issue](https://github.com/robert-min/performance_go_python_api/issues/3)


### Python Server
- [#issue](https://github.com/robert-min/performance_go_python_api/issues/4)

### nGrinder Test
- [#issue](https://github.com/robert-min/performance_go_python_api/issues/7)


## Project Tree

```sh
├── README.md                           ##  README 
├── data                                ##  data folder 
│   ├── csv_covert.py                   ##  conver data script
├── go_api                              ##  go_api folder 
│   ├── Dockerfile                      ##  go_api Dockerfile 
│   └── server                          ##  go_api server folder 
│       ├── go.mod                      ##  go_api mod 
│       ├── handlers                    ##  go_api handlers folder
│       │   ├── handler.go              ##  go_api web handler
│       │   └── register.go             ##  go_api handler register
│       └── server.go                   ##  go_api server
├── go_api_db                           ##  go_api_db folder 
│   ├── Dockerfile                      ##  go_api_db Dockerfile 
│   └── server                          ##  go_api_db server folder 
│       ├── conf                        ##  go_api_db conf folder 
│       │   └── conf.json               ##  go_api_db conf file 
│       ├── go.mod                      ##  go_api_db mod file
│       ├── go.sum                      ##  go_api_db sum file
│       ├── handlers                    ##  go_api_db handlers folder
│       │   ├── handler.go              ##  go_api_db web handler
│       │   └── register.go             ##  go_api_db handler register
│       ├── lib                         ##  go_api_db lib folder
│       │   └── db_connect.go           ##  go_api_db database connect
│       └── server.go                   ##  go_api_db server
├── python_api                          ##  python_api folder
│   ├── Dockerfile                      ##  python_api Dockerfile
│   ├── app.py                          ##  python_api run server file
│   ├── backend                         ##  python_api backend folder
│   │   └── src                         ##  python_api src folder
│   │       ├── __init__.py             ##  python_api src init file
│   │       └── movies.py               ##  python_api get movies file
│   └── requirements.txt                ##  python_api requirements.txt
└── python_api_db                       ##  python_api_db folder 
    ├── Dockerfile                      ##  python_api_db Dockerfile 
    ├── app.py                          ##  python_api_db run server file 
    ├── backend                         ##  python_api_db backend 
    │   ├── conf                        ##  python_api_db conf folder 
    │   │   └── conf.json               ##  python_api_db conf file 
    │   ├── lib                         ##  python_api_db lib folder 
    │   │   ├── __init__.py             ##  python_api_db lib init file 
    │   │   ├── db_connect.py           ##  python_api_db database connect 
    │   │   └── model.py                ##  python_api_db ORM model file 
    │   └── src                         ##  python_api_db src folder 
    │       ├── __init__.py             ##  python_api_db src init
    │       └── search.py               ##  python_api_db get search file 
    └── requirements.txt                ##  python_api_db requirements.txt 

```
