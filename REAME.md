## 概要

Vagrant & Ansibleを使って、ローカルマシンで、nginx - WebAPIアプリ構築の練習サンプル

## 準備
* Virtual Boxのインストール
https://www.oracle.com/jp/virtualization/technologies/vm/downloads/virtualbox-downloads.html

* Vagrantのインストール
https://www.vagrantup.com/docs/installation

* Ansibleのインストール
https://docs.ansible.com/ansible/2.9_ja/installation_guide/intro_installation.html

* Goのインストール
https://go.dev/doc/install

## 開発環境のprivate_networkの変更

```
vi Vagrantfile
```
private_networkの ipを自身のPC環境に合わせて書き換える。
既存設定は、192.168.56.10

## VirtualBox内のLinuxでプロビジョニング
```bash
cd vagrant-ansible-sample

# main.goのビルド
make build

# 起動してプロビジョニング
vagrant up
vagrant provision
```

## Nginx -> GoのWebアプリの疎通確認

```
curl http://192.168.56.10/api/hello
```
private_networkの ipを自身のPC環境に合わせて書き換える

## ansible-playbookからプロビジョニングする場合
```
ansible-playbook -i .vagrant/provisioners/ansible/inventory/vagrant_ansible_inventory provisioning/playbook.yml
```

## 後処理
```
vagrant destroy
```

## See
https://docs.ansible.com/ansible/2.9_ja/scenario_guides/guide_vagrant.html
