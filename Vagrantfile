Vagrant.require_version ">= 1.8.0"

Vagrant.configure(2) do |config|
  config.vm.box = "almalinux/8"
  config.vm.network "private_network", ip: "192.168.56.10"

  config.vm.provision "ansible" do |ansible|
    ansible.verbose = "v"
    ansible.playbook = "provisioning/playbook.yml"
  end
end