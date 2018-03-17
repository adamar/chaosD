# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.ssh.forward_agent = true

  config.vm.define "chaosD" do |b|
    b.vm.box = "bento/ubuntu-16.04"
    b.vm.hostname = 'chaosD'
    b.vm.network :private_network, ip: "192.168.201.232"

    b.vm.provider :virtualbox do |v|
      v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
      v.customize ["modifyvm", :id, "--memory", 512]
      v.customize ["modifyvm", :id, "--name", "chaosD"]
    end

    config.vm.synced_folder "~", "/home/vagrant/shared"

  end

end
