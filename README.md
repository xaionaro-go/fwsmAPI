# About

MSWF uses [https://github.com/xaionaro-go/fwsmConfig](https://github.com/xaionaro-go/fwsmConfig) to work with FWSM-compatible configuration files and setup a Linux machine to do the same job using [https://github.com/xaionaro-go/networkControl](https://github.com/xaionaro-go/networkControl). There're two interfaces: Web-interface for dummies and SSH-interfaces for Cisco-lovers.

![interfaces](https://raw.githubusercontent.com/xaionaro-go/mswfAPI/master/doc/interfaces.png)
![mswfAPI](https://raw.githubusercontent.com/xaionaro-go/mswfAPI/master/doc/mswfAPI.png)
![files](https://raw.githubusercontent.com/xaionaro-go/mswfAPI/master/doc/files.png)

# Install

```sh
cd /root
git clone git@devel.mephi.ru:noc/fwsm-config.git

git clone https://github.com/xaionaro-go/mswfAPI
cd mswfAPI
make
```

# Post-install

Installing additional runtime-dependencies
```sh
apt-get install -y bwm-ng
```
(it's used by URI /fwsm/status)

Define passwords in conf/app.conf:
```
user0.login = someLoginHere
user0.password_sha1 = sha1hashhere
user1.login = anotherLogin
user1.password_sha1 = anothersha1hash
user2.login = mswfShell
user2.password = someClearPassword
```

Define also JSON Web Tokens (JWT) secret in conf/app.conf:
```
jwt_secret = someSecretHere
```

Install Web-shell: [https://github.com/xaionaro/mswfWebControl](https://github.com/xaionaro/mswfWebControl)

Install SSH-shell: [https://github.com/xaionaro-go/mswfShell](https://github.com/xaionaro-go/mswfShell)

Create file `/root/fwsm-config/dynamic` and initialize git repository in `/root/fwsm-config` (`(cd /root/fwsm-config && git init && git add dynamic && git commit -a -m 'initial commit')`). `/root/fwsm-config/dynamic` is a FWSM-compatible configuration file.

# Run

```sh
revel run github.com/xaionaro-go/mswfAPI prod
```

(localhost:9000)

