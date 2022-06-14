## 概要

Vagrant & Ansibleを使って、ローカルマシンで、Nginx - WebAPIアプリ構築の練習サンプル

## 準備
* Virtual Boxのインストール
https://www.oracle.com/jp/virtualization/technologies/vm/downloads/virtualbox-downloads.html

* Vagrantのインストール
https://www.vagrantup.com/docs/installation

* Ansibleのインストール
https://docs.ansible.com/ansible/2.9_ja/installation_guide/intro_installation.html

* Goのインストール
https://go.dev/doc/install

## VirtualBox内のLinuxでプロビジョニング

```bash
cd vagrant-ansible-sample

# main.goのビルド
make build

# Linuxを起動してプロビジョニング
vagrant up
vagrant provision
```

## Nginx -> GoのWebアプリの疎通確認

```bash
curl http://192.168.56.10/api/hello
```
Hello Go Appと出力されればOK

## トラブルシュート

### サービスが起動しているか確認

ログイン
```
vagrant ssh
```
サービスステータスの確認
```
systemctl status
```

### ログの確認
```
sudo journalctl -u goapp
sudo journalctl -u nginx
```
## ansible-playbookからプロビジョニングする場合

```bash
ansible-playbook -i .vagrant/provisioners/ansible/inventory/vagrant_ansible_inventory provisioning/playbook.yml
```

## 後処理

```bash
vagrant destroy
```

## See
https://docs.ansible.com/ansible/2.9_ja/scenario_guides/guide_vagrant.html
