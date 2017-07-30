#Install PostgresSQL

#Install Nginx
Setup Request forwader from Nginx 80 to Our app at 9090
 
#Install GO on server
- https://www.digitalocean.com/community/tutorials/how-to-install-go-1-6-on-ubuntu-16-04

#Golang with Ngix
- https://evanbyrne.com/blog/go-production-server-ubuntu-nginx

#Pongo2 Template Engine
- https://github.com/flosch/pongo2
- https://github.com/flosch/pongo/tree/master/template_examples

#Go-Watcher watching .go file changes, and restarting the app in case of an update/delete/add
- https://github.com/canthefason/go-watcher

#Database Migration
- https://github.com/pressly/goose

#OS Variables
gvm use go1.7

export GOROOT=/usr/local/go
export GOPATH=$HOME/go-learn2
export GOBIN=$GOPATH/bin
export GOSRC=$GOPATH/src
export GOPKG=$GOPATH/pkg

export TEAHRM_DB_SERVER=postgres
export TEAHRM_DB_USERNAME=postgres
export TEAHRM_DB_PASSWORD=postgres
export TEAHRM_DB_NAME=teahrm
export TEAHRM_DB_TEST_NAME=teahrm_test
