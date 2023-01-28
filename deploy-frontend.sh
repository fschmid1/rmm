#!/bin/bash

make build.frontend

cd projects/frontend/dist

tar -czf files.tar.gz ./*
ssh rmm 'cd /var/www/rmm.festech.de && rm -rf ./*'
scp files.tar.gz rmm:/var/www/rmm.festech.de
ssh rmm 'cd /var/www/rmm.festech.de && tar -xzf files.tar.gz'
rm files.tar.gz

