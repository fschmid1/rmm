#!/bin/bash

prod="export const apiBase = 'https:\/\/rmm.festech.de\/api';"
oldFile=$(cat projects/frontend/src/vars.ts)

echo "$(sed "1s/.*/$prod/1" projects/frontend/src/vars.ts)" > projects/frontend/src/vars.ts

npm run build
cd projects/frontend/dist

tar -czf files.tar.gz ./*
ssh rmm 'cd /var/www/rmm.festech.de && rm -rf ./*'
scp files.tar.gz rmm:/var/www/rmm.festech.de
ssh rmm 'cd /var/www/rmm.festech.de && tar -xzf files.tar.gz'
rm files.tar.gz

cd ..
echo "$oldFile" > src/vars.ts
cd ../../
