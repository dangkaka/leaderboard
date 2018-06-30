# Leaderboard

## Requirements
* Docker 17.06+
* Go 1.9.2+

## Quickstart

1. Setup Go environment
    * Set GOPATH
        ```
        export GOPATH=$HOME/go
        ```
        More information [here](https://github.com/golang/go/wiki/SettingGOPATH)
    * Once GOPATH is set, you can clone this repository by doing 
        ```
        go get github.com/dangkaka/leaderboard
        ```
        It is located in `GOPATH/src/github.com/dangkaka/leaderboard`
1. Start the application
    ```
    docker-compose up -d
    ```
1. `localhost:3000` is ready to use

## Explore the API
`http://localhost:3000/swagger`

```
GET    /health                   
GET    /api/v1/scores           
POST   /api/v1/scores       
DELETE /api/v1/scores/:id
```

