#!/bin/bash


# This file is the driver code for remote deployment.
# Assumes control of the remote server and executes the deployment script.

# pem_name="go-moon-ec2.pem"
# pem_path="${HOME}/keystore"

echo "Deployment Driver"

pem_full_path=$1
echo $pem_full_path
exit 1;

ec2_ip=54.255.120.79
ec2=ubuntu@"${ec2_ip}"
ec2_config_dir=/home/ubuntu/customkeystore/production
commit_hash=5f7a8e9e4ff4a03b5540598641f5e302cc93eec8


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
