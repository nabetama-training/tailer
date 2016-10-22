# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "centos6.6"
  config.vm.network "forwarded_port", guest: 80, host: 8080
end
