#!/bin/sh

HOST=http://101.201.210.248/

curl --head $HOST
#1
./chaos_os create nginx crash
./chaos_os destroy nginx crash

#2
./chaos_os create nginx restart

#3
#mode test
./chaos_os create nginx config 
./chaos_os create nginx config --mode fff

#file
./chaos_os create nginx config --mode file
./chaos_os create nginx config --mode file --file ok.conf
./chaos_os destroy nginx config

#cmd
#list
./chaos_os create nginx config --list true

./chaos_os create nginx config --mode cmd
./chaos_os create nginx config --mode cmd --block-id 100

./chaos_os create nginx config --mode cmd --block-id 3 --set-config='listen=8899'
./chaos_os destroy nginx config

./chaos_os create nginx config --mode cmd --block-id 4 --set-config='proxy_pass=https://www.tmall.com'
./chaos_os destroy nginx config

#4
./chaos_os create nginx response --path / --body 'ok'
curl -v "${HOST}/test"
./chaos_os destroy nginx response

./chaos_os create nginx response --path / --code 500
curl -v "${HOST}/test"
./chaos_os destroy nginx response

./chaos_os create nginx response --path / --code 200 --body '{"a":1}'
curl -v "${HOST}/test"
./chaos_os destroy nginx response

./chaos_os create nginx response --regex /t.* --code 200 --body '{"a":1}' --header 'Server=mock;' --server 0
curl -v "${HOST}/test"
./chaos_os destroy nginx response
