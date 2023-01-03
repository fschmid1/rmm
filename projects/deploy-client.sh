#!/bin/bash

myArray=("hoku01" "hoku02" "hoku04" "loadbalancer" "dns")
arm=("hoku03" "hoku05")

make build.client.linux
pwd=$(bazel cquery //projects/rmm/go/client:RMM_CLIENT_LINUX --output=files)

for str in ${myArray[@]}; do
  scp ../../$pwd $str:/home/felix/RMM_CLIENT
done

make build.client.arm
pwd=$(bazel cquery //projects/rmm/go/client:RMM_CLIENT_ARM --output=files)

for str in ${arm[@]}; do
  scp ../../$pwd $str:/home/felix/RMM_CLIENT
done

for str in ${myArray[@]}; do
	ssh $str "sudo systemctl stop rmm && sudo mv ./RMM_CLIENT /opt/rmm/RMM_CLIENT && sudo rm -rf /etc/fes-rmm/device && sudo systemctl start rmm && sleep 4 && sudo systemctl restart rmm"
done

for str in ${arm[@]}; do
	ssh $str "sudo systemctl stop rmm && sudo mv ./RMM_CLIENT /opt/rmm/RMM_CLIENT && sudo rm -rf /etc/fes-rmm/device && sudo systemctl start rmm && sleep 4 && sudo systemctl restart rmm"
done