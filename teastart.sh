#!/usr/bin/env bash

echo "\nHi, $USER!"

echo ""
echo "ÛÛÛÛÛÛÛÛÛÛÛ ÛÛÛÛÛÛÛÛÛÛ   ÛÛÛÛÛÛÛÛÛ   ÛÛÛÛÛ   ÛÛÛÛÛ ÛÛÛÛÛÛÛÛÛÛÛ   ÛÛÛÛÛÛ   ÛÛÛÛÛÛ"
echo "°Û°°°ÛÛÛ°°°Û°°ÛÛÛ°°°°°Û  ÛÛÛ°°°°°ÛÛÛ °°ÛÛÛ   °°ÛÛÛ °°ÛÛÛ°°°°°ÛÛÛ °°ÛÛÛÛÛÛ ÛÛÛÛÛÛ" 
echo "°   °ÛÛÛ  °  °ÛÛÛ  Û °  °ÛÛÛ    °ÛÛÛ  °ÛÛÛ    °ÛÛÛ  °ÛÛÛ    °ÛÛÛ  °ÛÛÛ°ÛÛÛÛÛ°ÛÛÛ" 
echo "    °ÛÛÛ     °ÛÛÛÛÛÛ    °ÛÛÛÛÛÛÛÛÛÛÛ  °ÛÛÛÛÛÛÛÛÛÛÛ  °ÛÛÛÛÛÛÛÛÛÛ   °ÛÛÛ°°ÛÛÛ °ÛÛÛ" 
echo "    °ÛÛÛ     °ÛÛÛ°°Û    °ÛÛÛ°°°°°ÛÛÛ  °ÛÛÛ°°°°°ÛÛÛ  °ÛÛÛ°°°°°ÛÛÛ  °ÛÛÛ °°°  °ÛÛÛ" 
echo "    °ÛÛÛ     °ÛÛÛ °   Û °ÛÛÛ    °ÛÛÛ  °ÛÛÛ    °ÛÛÛ  °ÛÛÛ    °ÛÛÛ  °ÛÛÛ      °ÛÛÛ" 
echo "    ÛÛÛÛÛ    ÛÛÛÛÛÛÛÛÛÛ ÛÛÛÛÛ   ÛÛÛÛÛ ÛÛÛÛÛ   ÛÛÛÛÛ ÛÛÛÛÛ   ÛÛÛÛÛ ÛÛÛÛÛ     ÛÛÛÛÛ"
echo "   °°°°°    °°°°°°°°°° °°°°°   °°°°° °°°°°   °°°°° °°°°°   °°°°° °°°°°     °°°°°" 


echo "GO ENVIRONTMENT VARIABLES FOR TEAHRM"

gvm use go1.7
lsb_release -a | grep Description

export GOPATH=$HOME/go-learn2
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export GOSRC=$GOPATH/src
export GOPKG=$GOPATH/pkg

export TEAHRM_DB_SERVER=postgres
export TEAHRM_DB_USERNAME=postgres
export TEAHRM_DB_PASSWORD=postgres
export TEAHRM_DB_NAME=teahrm
export TEAHRM_DB_TEST_NAME=teahrm_test
export TEAHRM_DB_MIGRATION=$GOSRC/gitlab.com/mhyusufibrahim/teahrm/database

PATH=$GOBIN:$PATH
cd $GOPATH/src/gitlab.com/mhyusufibrahim/teahrm

echo "\nGO paths:"
echo "- GOROOT : " $GOROOT
echo "- GOBIN  : " $GOBIN
echo "- GOSRC  : " $GOSRC
echo "- GOPKG  : " $GOPKG
echo "- GOPATH : " $GOPATH
echo "- GVM_ROOT" $GVM_ROOT

echo "\nTEAHRM Paths:"
echo "- TEAHRM_DB_USERNAME   : " $TEAHRM_DB_USERNAME 
echo "- TEAHRM_DB_PASSWORD   : " $TEAHRM_DB_PASSWORD
echo "- TEAHRM_DB_NAME       : " $TEAHRM_DB_NAME
echo "- TEAHRM_DB_TEST_NAME  : " $TEAHRM_DB_TEST_NAME
echo "- TEAHRM_DB_MIGRATION  : " $TEAHRM_DB_MIGRATION


echo "\nDatabase status:"
service postgresql status
psql -l | grep $TEAHRM_DB_NAME 
