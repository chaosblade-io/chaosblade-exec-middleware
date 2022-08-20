#!/bin/sh

#
# Copyright 1999-2020 Alibaba Group Holding Ltd.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

HOST=http://101.201.210.248/

curl --head $HOST
#1
./chaos_middleware create nginx crash
./chaos_middleware destroy nginx crash

#2
./chaos_middleware create nginx restart

#3
#mode test
./chaos_middleware create nginx config
./chaos_middleware create nginx config --mode fff

#file
./chaos_middleware create nginx config --mode file
./chaos_middleware create nginx config --mode file --file ok.conf
./chaos_middleware create nginx config --mode file --file wrong.conf
./chaos_middleware destroy nginx config

#cmd
#list
./chaos_middleware create nginx config --mode cmd

./chaos_middleware create nginx config --mode cmd --block 'http.server[0]' --set-config='listen=8899'
./chaos_middleware destroy nginx config

./chaos_middleware create nginx config --mode cmd --block 'http.server[0].location[0]' --set-config='proxy_pass=https://www.tmall.com'
./chaos_middleware destroy nginx config

#4
./chaos_middleware create nginx response --path / --body 'ok'
curl -v ${HOST}
./chaos_middleware destroy nginx response

./chaos_middleware create nginx response --path / --code 500
curl -v ${HOST}
./chaos_middleware destroy nginx response

./chaos_middleware create nginx response --path / --code 200 --body '{"a":1}'
curl -v ${HOST}
./chaos_middleware destroy nginx response

./chaos_middleware create nginx response --regex /t.* --code 200 --body '{"a":1}' --header 'Server=mock;' --server 0
curl -v ${HOST}
./chaos_middleware destroy nginx response
