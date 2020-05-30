# Create a server. api doc https://www.vultr.com/api/
resource "vultr_server" "dev_server" {
    region_id = "39" # miami https://api.vultr.com/v1/regions/list
    plan_id = "204" # https://api.vultr.com/v1/plans/list?type=ssd
    hostname = "node"
    os_id = "270" # https://api.vultr.com/v1/os/list
    # ssh_key_ids = [""] # main
    # tags          = {
    #     Name        = "Test Server"
    #     Environment = "development"
    # }
    label = "node"
    # snapshot_id = " "
    enable_private_network= true
    enable_ipv6 = true
    auto_backup = false
    ddos_protection = false
    notify_activate = false
}

