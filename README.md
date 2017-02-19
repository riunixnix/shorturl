# Shorturl
Simple Restful API for convert website url become shorter.

# Language 
Go ( Go Language ), MySql, Fabric

# Install
The installation tools use Fabric, so we need to install it first:

```
#!bash

sudo apt-get install fabric
#or can follow instruction from here http://www.fabfile.org/installing.html

git clone git@github.com:riunixnix/shorturl.git

cd shorturl

fab host_setting:HOSTNAME,server_username deploy
#example :  fab host_setting:127.0.0.1:2222,vagrant deploy
```


**Hashids**

```
#!bash

go get github.com/speps/go-hashids
```
**Go Sql Driver**

```
#!bash

go get github.com/go-sql-driver/mysql
```