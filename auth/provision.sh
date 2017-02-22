#! /usr/bin/env sh

# Import the jackmarshall deploy key
sudo -u vagrant echo -e "-----BEGIN RSA PRIVATE KEY-----\n\
MIIEowIBAAKCAQEAucbOPA2mZhJrCTbtbOgDmgpv/OgI8FfwskbROvu6mjd2nuvB\n\
26IRtbXuzP0GkHjTOHiRohMXOssbldLmpmkXLZ96ZcAG0SHCOxtWZbb+ut+Nsj2+\n\
xM3cCRvMWuC12XLeN/kYFnT7Ak1Kfwrui1OdxOFFhIGNgNUWdsReHQJ+zjyfyjTM\n\
xDn1ReZx1HPvfFiWHA2BSsO4DUrgUejRyNl2jaTbu0/WesBb2zrsVoFcXpRCfR77\n\
K9a8X0q67W7UIq8OemMWSiWsT5XXxbppoIOz+Zi48YJdlm1fG8XemmGabmm7IXmZ\n\
OM7Nt1fqFZ0hlH7SRTnCy+EJcJbhTzI3A9tp3wIDAQABAoIBAQCaZYrTSDjqDhad\n\
EuRiJbWQmWoXU7TSIxQs5kRP9BRSxRO14pQ7+EclsO2lughxm6lX/oRyodElkNX9\n\
P8lntmGIDknINL61oovtWbwFTwAHyXHXGA/rOnfLrim5wZYBAcGD3WbSiyht8lSe\n\
nzQ/4R93GA4RoSY8U1yXGn5pN8Cxny+IbaxKrkiIphf6yfA5Kzc1A1cxixb8Uc1S\n\
sJ8tjozGsuU4hycoT9dV+viW2UfIKnu0FbB9S4Jj5HbzmUFV8YGTNbCGTtVAc7Z8\n\
8Ir1s9f86IRsohop+TcipO7LOTwCANDNx2PGJIKWXX7wP0py2k4PprtcXYspr+ke\n\
H8UUy4vxAoGBAOmadWcKqM6vRx0JJs8APCCcZLeB6VZVFs8P0fsfxI9vrfEI9XuO\n\
F7o31tC2wNxA3PB7StGX2Rd+d1qnCB5gThFA3PAe9UnxYziHSdp0ReFBbQTvBi4w\n\
uWOtOl221LNpQ85B2NmQsENl00ue5/UK9SUqnG4D9Tkf6Xktvm383Pr5AoGBAMuW\n\
faAlNxnDGxdU4f9i4QZHfs5W6rBFMW6t+ulDOingQi842VzSM69xeYADoSqqkTeg\n\
pwi74WhJlQGuPUnTTf0rbVyvYwjW2DiBlcef3pKODjVhCzbzMWFQ//SkXLmyLGv2\n\
Frkdy1EoDuFWfMw7Tk9JHRMw0Ecxiyms+GvCvamXAoGAc5hZLCocKlqVyAp1R82H\n\
82Mwe5YvjucAKGpdr/SiYZAT9MhIIWHI3/8tr44flV4LdnysI1cOuA9+pD8vR0jE\n\
mMkOC0xmeLaYMKrRouYShkrH1b8pwcWel3JrBWCplDYstW3kDpUIg/FpQ2xYv+sy\n\
L7gMPGx89k2STMKUDv0npJECgYAp/9HKx+KmHeno1koJNYzWjslKKAQieSPJgXrR\n\
xFyrZuhAdanhzDt9Zht6nUUSWD29DBdpZdkKeqOSxS8/m6frIppdFJGJsUjJ33HW\n\
zNRNx6CGwo5dt63oNZ2YCpBWfrtLSmlFpy/afjkTo5P7TgDJunuVCeztwvrCjjxx\n\
fqtB4QKBgHBFlaF5loMG/8Mkjsalll2H8Ghv62PbbjygxSbA5OmmSMjECl5NdI6e\n\
0QuynZgQ00+ziPTH64dOuYPDDxdxk3SyUbv8DvNB1ypvGFIXioCkdaOEhl4L/OfB\n\
hRYXAtkiZ/avJ9hd5SDLwxxIV44Jjp35KJxuMYqEUfanpce4PFnW\n\
-----END RSA PRIVATE KEY-----" > /home/vagrant/.ssh/id_rsa
sudo -u vagrant echo -e "Host git.elwinar.com\n    Port 2222" >> /home/vagrant/.ssh/config
chown -R vagrant:vagrant /home/vagrant/.ssh
chmod 600 /home/vagrant/.ssh/id_rsa

# Update the system globally
pacman -Sy --noconfirm

# Install utilities
pacman -S --noconfirm gcc
pacman -S --noconfirm git
pacman -S --noconfirm make

# Golang tools
pacman -S --noconfirm go
echo "export GOPATH=/home/vagrant" >> /home/vagrant/.bashrc
echo "export PATH=\$GOPATH/bin:\$PATH" >> /home/vagrant/.bashrc
chown vagrant:vagrant -R /home/vagrant/src

# Docker tools
pacman -S --noconfirm docker
systemctl enable docker
systemctl start docker
gpasswd -a vagrant docker

# Redis
pacman -S --noconfirm redis
systemctl enable redis
systemctl start redis

export APP_PASSWORD_COST=12
export APP_PORT=8080
export APP_REDIS_PORT_6379_TCP_ADDR=127.0.0.1
export APP_REDIS_PORT_6379_TCP_PORT=6379
export APP_SECRET=secret
