# Instructions

## Prerequisite
Download Terraform [binary](https://www.terraform.io/downloads.html)  and install it.

## Initialization
- Download binaries and genesis file, then place them in `microtick/files`
- Fill in your AWS api keys and change protocol-specific parameters in `microtick/vars.tf`
- Tweak security group rules in `microtick/security_group.tf`. It's currently commented out in the code as I planned on using an existing security group at the time of writing
- Tweak whatever tf you need

## Deployment
```
$ terraform init # under microtick/
$ ssh-keygen -f terra # save a copy of the private key

### alternatively you can import an existing key ####
$ terraform import aws_key_pair.main_key terra

$ terraform validate # check syntax
# terraform plan # dryrun
$ terraform apply # deploy
```

## Monitoring
```
$ sudo systemctl status mtd
$ tail -f /var/log/mtd.log
```
**NOTE: No tendermint node keys will be added during the deployment. You need to do that on your own, first thing when you log on to server.**
```
$ mtcli keys add <key_name>
```
## Termination
```
$ terraform destroy
```