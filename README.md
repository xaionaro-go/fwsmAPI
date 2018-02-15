# Install

```sh
cd /root
git clone git@devel.mephi.ru:noc/fwsm-config.git

git clone https://github.com/xaionaro-go/fwsmAPI
cd fwsmAPI
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
```

Define also JSON Web Tokens (JWT) secret in conf/app.conf:
```
jwt_secret = someSecretHere
```

Install: https://github.com/xaionaro/fwsmWebControl

# Run

```sh
revel run github.com/xaionaro-go/fwsmAPI prod
```

(localhost:9000)

