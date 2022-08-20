#!/bin/sh

HOST=http://101.201.210.248/

curl --head $HOST
#1
./blade create nginx crash
./blade destroy nginx crash

#2
./blade create nginx restart

#3
#mode test
./blade create nginx config 
./blade create nginx config --mode fff

#file
./blade create nginx config --mode file
./blade create nginx config --mode file --file ok.conf
./blade create nginx config --mode file --file wrong.conf
./blade destroy nginx config

#cmd
#list
./blade create nginx config --list

./blade create nginx config --mode cmd
./blade create nginx config --mode cmd --block-id 100

./blade create nginx config --mode cmd --block-id 3 --set-config='listen=8899'
./blade destroy nginx config

./blade create nginx config --mode cmd --block-id 4 --set-config='proxy_pass=https://www.tmall.com'
./blade destroy nginx config

#4
./blade create nginx response --path / --body 'ok'
curl -v "${HOST}/test"
./blade destroy nginx response

./blade create nginx response --path / --code 500
curl -v "${HOST}/test"
./blade destroy nginx response

./blade create nginx response --path / --code 200 --body '{"a":1}'
curl -v "${HOST}/test"
./blade destroy nginx response

./blade create nginx response --regex /t.* --code 200 --body '{"a":1}' --header 'Server=mock;' --server 0
curl -v "${HOST}/test"
./blade destroy nginx response
