#!/bin/bash


# This file is the driver code for remote deployment.
# Assumes control of the remote server and executes the deployment script.



echo "Deployment Driver"

pem_full_path=$1 # file path to private key to access the instance

ec2_ip=$2

ec2_username=ubuntu
ec2="$ec2_username@${ec2_ip}"


# ec2_config_dir=/home/ubuntu/customkeystore/production
commit_hash=$3
deploy_script_path=$4


echo "Copying deployment script to ec2 instance ${ec2}"


scp -tt "${deploy_script_path}" localhost:"$HOME/"

scp -i "${pem_full_path}" -tt "${deploy_script_path}" ${ec2}:"$HOME/"

echo $deploy_script_path
exit 1;


echo "Execute script on server environment"
ssh -t -i "${pem_full_path}" "${ec2}" "chmod +x ${deploy_script_path} && ${deploy_script_path} ${commit_hash}"

echo "End of Deployment"
