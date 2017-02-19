from __future__ import with_statement
from fabric.api import *
from fabric.contrib.console import confirm
from fabric.contrib.files import exists
from fabric.context_managers import cd

env.hosts = []

def host_setting(host,user):    
    env.hosts.append(host)
    env.user = user
    print env.hosts

def deploy():
	with settings(warn_only=True):
		# run('sudo add-apt-repository "deb http://archive.ubuntu.com/ubuntu $(lsb_release -sc) main universe"')
		# run("sudo apt-get update")
		# run("sudo apt-get -y install build-essentials")
		# run("sudo apt-get -y install clang zip unzip")
		# run("sudo apt-get -y install git")

		tmp_folder = "~/tmp_folder"
		if not exists(tmp_folder): 
			print "Creating temporary folder"
			run("mkdir -p %s;" % tmp_folder)

		with cd(tmp_folder):		
			if not exists('golang-tools-install-script'):
				run("git clone https://github.com/canha/golang-tools-install-script.git")	
			
			run("cd golang-tools-install-script;sudo bash goinstall.sh --remove") #remove any existing golang
			run("cd golang-tools-install-script;sudo bash goinstall.sh --64")
		run("source ~/.bashrc")