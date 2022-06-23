#!/bin/bash


# This file is the driver code for remote deployment.
# Assumes control of the remote server and executes the deployment script.



echo "Deployment Driver"

pem_full_path=$1 # file path to private key to access the instance

ec2_ip=$2

ec2_username=ubuntu
ec2=$ec2_username@"${ec2_ip}"


# ec2_config_dir=/home/ubuntu/customkeystore/production
commit_hash=$3
echo $commit_hash

exit 1;

deploy_script=some.sh

deploy_script_path="./${deploy_script}"

echo "Supply configuration file"

ssh -t -i "${pem_full_path}" "${ec2}" "sudo rm -rf ${ec2_config_dir}"
scp -r -i "${pem_full_path}" ${ec2_config_dir} "${ec2}":"${ec2_config_dir}"

echo "Copying deployment script to ec2 instance"
scp -i ${pem_full_path} ./some.sh "${ec2}":./ 

echo "Execute script on server environment"
ssh -t -i "${pem_full_path}" "${ec2}" "chmod +x ${deploy_script_path} && ${deploy_script_path} ${commit_hash}"

echo "End of Deployment"
