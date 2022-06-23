#!/bin/bash


# This script is runnable 

# NOTE:
# There is no support for environment variables. By right, the values should be taken from the configuration file.
# Functional on EC2, Ubuntu 22

gh_repo_url=https://github.com/noellimx/project-6-backend.git
gh_go_repo_url=github.com/noellimx/project-6-backend

app_name=gomoon
repo_dst="$HOME/repo/${app_name}"


go_binary_dir="$HOME/go/bin"

dbname=gomoontest

listeningPort=8080

commit_hash=$1



# Script for deployment, should be running locally on the server machine

# df -H
echo "Running script"

echo "[INSTALL] Checking git installation..."

echo "[INSTALL] Update git"
sudo DEBIAN_FRONTEND=noninteractive apt-get install --assume-yes git
sudo DEBIAN_FRONTEND=noninteractive apt-get upgrade --assume-yes -y git


if [ "$(which git)" == "" ]; then
	echo "Fatal: git not found" && exit 1;
fi

echo "[INSTALL] Checking go installation..."

echo "[INSTALL] Update go"

sudo DEBIAN_FRONTEND=noninteractive apt-get install --assume-yes golang-go
sudo DEBIAN_FRONTEND=noninteractive apt-get upgrade --assume-yes -y golang-go


if [ "$(which go)" == "" ]; then
	echo "Fatal: go not found" && exit 1;
fi

echo "[INSTALL] Checking psql installation..."

echo "[INSTALL] Update psql"

sudo DEBIAN_FRONTEND=noninteractive apt-get install --assume-yes postgresql
sudo DEBIAN_FRONTEND=noninteractive apt-get upgrade --assume-yes -y postgresql

if [ "$(which psql)" == "" ]; then
	echo "Fatal: psql not found" && exit 1;
fi

echo "[SERVICE] WARNING: Deleting DB"

sudo -u postgres dropdb --if-exists ${dbname}

echo "[SERVICE] Creating DB"
sudo -u postgres createdb ${dbname}

sudo -u postgres psql -c "ALTER USER postgres PASSWORD 'postgres';"

echo "[SERVICE]  ------"


echo "[REPO] Cleaning... ${repo_dst}"
rm -rf ${repo_dst}

echo "[REPO] Cloning... ${repo_dst}"
git clone ${gh_repo_url} ${repo_dst}

echo "[REPO] Checking out branch... ${repo_dst}"

git checkout ${commit_hash}

echo "[BINARY] Changing pwd to repo..."
ls ${repo_dst}
cd ${repo_dst}


echo "commit hash ${commit_hash}"
git checkout ${commit_hash}
exit 1;


echo "[BINARY] Cleaning binary... ${go_binary_dir}"
rm -rf ${go_binary_dir}/${app_name}

echo "[BINARY] Installing..."

go install

ls ${go_binary_dir}/${app_name}


echo "[Run] Killing existing process on port ${listeningPort}"

fuser -k ${listeningPort}/tcp

${go_binary_dir}/${app_name}

exit 0;

