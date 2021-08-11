#!/bin/bash

echo "Downloading aura..."
wget "https://github.com/arghyagod-coder/GoLang/releases/download/1.0.0/aura-linux-amd64" 

echo "Adding aura into PATH..."

mkdir -p ~/.aura;
mv ./aura-linux-amd64 ~/.aura
echo "PATH=$PATH:~/.aura" >> /etc/environment  
echo "aura installation is completed!"
echo "You need to restart the shell to use aura."