# Instructions

## Prerequisite
Download Terraform [binary](https://www.terraform.io/downloads.html)  and install it.

## Initialization
- Download the [binaries](https://github.com/certikfoundation/chain/releases/) and the genesis file, and place them in `certik/files`
- Fill in your AWS api keys and change protocol-specific parameters in `certik/vars.tf`
- Tweak the security group rules in `certik/security_group.tf`
- Tweak whatever, however you like

## Deployment
```
$ terraform init # under certik/
$ ssh-keygen -f ctk # name it however you want. save a copy of the private key

### alternatively you can import an existing key ####
$ terraform import aws_key_pair.main_key ctk

$ terraform validate # check syntax
# terraform plan # dryrun
$ terraform apply # deploy
```

## Monitoring
```
$ sudo systemctl status certikd
$ tail -f /var/log/certikd.log
```
**NOTE: No tendermint node keys will be added during the deployment. You need to do that on your own, first thing after logging in.**
```
$ certikcli keys add <key_name>
```
## Termination
```
$ terraform destroy
```