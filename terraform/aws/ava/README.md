# Instructions

## Prerequisite
Download Terraform [binary](https://www.terraform.io/downloads.html)  and install it.

## Initialization
- Fill in your AWS api keys in `ava/vars.tf`
- Tweak security group rules in `ava/security_group.tf`
- Tweak whatever tf you need

## Deployment
```
$ terraform init # under ava/
$ ssh-keygen -f terra # save a copy of the private key
$ terraform validate # check syntax
# terraform plan # dryrun
$ terraform apply # deploy
```

## Monitoring
```
$ sudo systemctl status ava
$ tail -f /var/log/ava.log
```

## Termination
```
$ terraform destroy
```