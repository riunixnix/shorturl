# Shorturl
Simple Restful API for convert website url become shorter.
The Aplication use port 9090 to access it.

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

fab host_setting:HOSTNAME,server_username install
#example :  fab host_setting:127.0.0.1:2222,vagrant install

fab host_setting:127.0.0.1:2222,vagrant reload_script #for pull latest update
```

Then ssh to the server, then go to GOPATH
```
#!bash

cd $GOPATH

nohup env conf_path="$GOPATH/conf" ./bin/shorturl > ~/shorturl.log& #for run the webserver

```

For testing from inside server 
```
#!bash
curl -sX POST -H 'Content-Type: application/json' 'http://localhost:9090/shorten/' -d '{"url":"http://a.very.long.urlff"}'

```

For better performance , can combine with nginx and setting proxy to the local.

Todo Later:
- Will Add auto install for nginx and configuration
- make the app run as service
