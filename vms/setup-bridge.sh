#!/bin/bash

# Nome da bridge
BRIDGE=br-lan

# Criação das interfaces TAP
for i in 1 2 3; do
    sudo ip tuntap add dev tap$i mode tap user $(whoami)
    sudo ip link set tap$i up
    sudo brctl addif $BRIDGE tap$i
done

# Ativa a bridge, se necessário
sudo ip link set $BRIDGE up

