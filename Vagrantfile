Vagrant.configure("2") do |config|
  config.vm.box = "kaorimatz/archlinux-x86_64"
  config.vm.hostname = "jackmarshall"
  config.ssh.insert_key = false
  config.vm.synced_folder ".", "/home/vagrant/src/github.com/chibimi/jackmarshall"
  config.vm.provision :shell, path: "provision.sh"
  config.vm.network :forwarded_port, guest: 8080, host: 8081, auto_correct: true
end
