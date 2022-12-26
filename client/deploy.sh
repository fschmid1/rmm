#!/bin/bash

myArray=("hoku01" "hoku02" "hoku04" "loadbalancer" "dns")

make build

for str in ${myArray[@]}; do
  scp build/RMM_CLIENT $str:/home/felix/RMM_CLIENT
done

for str in ${myArray[@]}; do
	ssh $str "sudo systemctl stop rmm && sudo mv ./RMM_CLIENT /opt/rmm/RMM_CLIENT && sudo rm -rf /etc/fes-rmm/device && sudo systemctl start rmm && sleep 4 && sudo systemctl restart rmm"
done