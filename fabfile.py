from __future__ import with_statement
from fabric.api import *
from fabric.contrib.console import confirm
from fabric.contrib.files import exists
from fabric.context_managers import cd
import json 

env.hosts = []
env.shell = "/bin/bash -l -c -i"
project_idx  = 2
lib_needed = ['github.com/speps/go-hashids','github.com/go-sql-driver/mysql','github.com/riunixnix/shorturl']

def host_setting(host,user):    
    env.hosts.append(host)
    env.user = user
    print env.hosts

def install_mysql():
	print "Installing Mysql"			
	# run("sudo apt-get install -y mysql-server")
	# run("mysql_secure_installation")

	db_host 	= raw_input("MySql Host ?")
	db_user 	= raw_input("MySql Username ?")
	db_password = raw_input("MySql Password ?")
	db_name     = "shorturl"

	conf = {
		"Host" : db_host,
		"User" : db_user,
		"Pass" : db_password,
		"Db" : db_name
	}
	conf = json.dumps(conf)
	with cd(run("echo $GOPATH")):
		if not exists('conf'):
			run("mkdir conf")
		run("echo '%s' > conf/db.json" % conf)
		sql_file = "src/"+lib_needed[project_idx]+"/db.sql"
		run("mysql -u %s -p%s -e \"DROP DATABASE %s;CREATE DATABASE %s\"" % (db_user,db_password,db_name,db_name))
		run("mysql -u %s -p%s %s < %s" % (db_user,db_password,db_name,sql_file))

def install_golang():
	with cd(tmp_folder):	
		print "Install Go Lang"	
		if not exists('golang-tools-install-script'):
			run("git clone https://github.com/canha/golang-tools-install-script.git")	
		
		run("cd golang-tools-install-script;sudo bash goinstall.sh --remove") #remove any existing golang
		run("cd golang-tools-install-script;sudo bash goinstall.sh --64")			
		print ""

def update():
	run('sudo add-apt-repository "deb http://archive.ubuntu.com/ubuntu $(lsb_release -sc) main universe"')
	run("sudo apt-get update")
	run("sudo apt-get -y install build-essentials")
	run("sudo apt-get -y install clang zip unzip")
	run("sudo apt-get -y install git")

	tmp_folder = "~/tmp_folder"
	if not exists(tmp_folder): 
		print "Creating temporary folder"
		run("mkdir -p %s;" % tmp_folder)
		print ""

def install_golang_lib():
	run("sudo chown -R %s:%s $GOPATH" % (env.user,env.user))
	with cd(run("echo $GOPATH")):
		for lib in lib_needed:
			print "Installing %s ..." % lib
			run('go get -u %s' % lib)
			run('go install %s' % lib)


def deploy():
	#run("source ~/.bashrc")	

	with settings(warn_only=True):
		#update()

		#install_golang()
		
		install_golang_lib()
		
		install_mysql()
			
				
		

		

