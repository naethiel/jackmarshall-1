#!/usr/bin/env bash

#update package list
pacman -Syu

# Utility tools setup
pacman -S --noconfirm gcc
pacman -S --noconfirm git
pacman -S --noconfirm make

# Golang setup
pacman -S --noconfirm go
echo "export GOPATH=/home/vagrant" >> /home/vagrant/.bashrc
echo "export PATH=\$GOPATH/bin:\$PATH" >> /home/vagrant/.bashrc
chown vagrant:vagrant -R /home/vagrant/src

# Docker setup
pacman -S --noconfirm docker
systemctl enable docker
systemctl start docker
gpasswd -a vagrant docker

# Mongo setup
pacman -S --noconfirm mongodb
systemctl enable mongodb
systemctl start mongodb
