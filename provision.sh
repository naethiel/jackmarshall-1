#! /usr/bin/env sh

# Update the system globally
pacman -Sy --noconfirm

# Utility tools setup
pacman -S --noconfirm gcc
pacman -S --noconfirm git
pacman -S --noconfirm make

# Node setup
pacman -S --noconfirm npm


# Golang setup
pacman -S --noconfirm go
echo "export GOPATH=/home/vagrant" >> /home/vagrant/.bashrc
echo "export PATH=\$GOPATH/bin:\$PATH" >> /home/vagrant/.bashrc
chown vagrant:vagrant -R /home/vagrant/src
