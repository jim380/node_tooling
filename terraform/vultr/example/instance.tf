# api doc https://www.vultr.com/api/v2/
resource "vultr_instance" "monitor" {
  region   = var.region
  plan     = var.plan
  hostname = var.hostname
  os_id    = var.os_id
  # ssh_key_ids = [""] # main
  label = var.label
  # tag = "node"
  # snapshot_id = " "
  enable_private_network = true
  enable_ipv6            = true
  # backups = "enabled"
  # ddos_protection = false
  # activation_email = true

  # install grafana
  provisioner "remote-exec" {
    inline = [
      "sudo apt-get install -y apt-transport-https",
      "sudo apt-get install -y software-properties-common wget",
      "wget -q -O - https://packages.grafana.com/gpg.key | sudo apt-key add -",
      "echo \"deb https://packages.grafana.com/oss/deb stable main\" | sudo tee -a /etc/apt/sources.list.d/grafana.list",
      "sudo apt-get update",
      "sudo apt-get install grafana",
      "sudo systemctl daemon-reload",
      "sudo systemctl start grafana-server",
      "sudo systemctl enable grafana-server.service",
    ]

    connection {
      type = "ssh"
      host = self.main_ip
      user     = "root"
      password = self.default_password
    }
  }

  # install prometheus
  provisioner "file" {
    content = <<-EOF
            [Unit]
            Description=Prometheus
            After=network.target

            [Service]
            User=jim380
            Restart=always
            ExecStart=/usr/local/bin/prometheus --config.file=/etc/prometheus/prometheus.yml --storage.tsdb.path=/var/lib/prometheus --web.console.templates=/etc/prometheus/consoles --web.console.libraries=/etc/prometheus/console_libraries --web.listen-address=0.0.0.0:9090
            ExecReload=/bin/kill -HUP $MAINPID
            SyslogIdentifier=prometheus
            Restart=always
            RestartSec=3
            LimitNOFILE=4096
            TimeoutStartSec=0

            [Install]
            WantedBy=multi-user.target
        EOF

    destination = "/tmp/prometheus.service"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo mkdir -p /etc/prometheus /var/lib/prometheus /home/jim380/Documents /home/jim380/Downloads",
      "cd /home/jim380/Downloads",
      "curl -Lk https://github.com/prometheus/prometheus/releases/download/${data.github_release.prometheus.release_tag}/prometheus-${trimprefix(data.github_release.prometheus.release_tag, "v")}.linux-amd64.tar.gz --retry 5 | sudo tar xvf --directory ~/Downloads",
      "cd prometheus-${trimprefix(data.github_release.prometheus.release_tag, "v")}.linux-amd64",
      "sudo cp {prometheus,promtool} /usr/local/bin/",
      "sudo chown jim380 /usr/local/bin/{prometheus,promtool}",
      "sudo chown -R jim380 /etc/prometheus",
      # /var/lib/prometheus -> data
      "sudo chown jim380 /var/lib/prometheus",
      # /etc/prometheus -> config
      "sudo cp -r {consoles,console_libraries} /etc/prometheus/",
      "sudo cp prometheus.yml /etc/prometheus/",
      "sudo mv /tmp/prometheus.service /etc/systemd/system",
      "sudo systemctl daemon-reload",
      "sudo systemctl enable prometheus",
      "sudo systemctl start prometheus",
    ]

    connection {
      type = "ssh"
      host = self.main_ip
      user     = "root"
      password = self.default_password
    }
  }

  # install node exporter
  provisioner "file" {
    content = <<-EOF
            [Unit]
            Description=Node Exporter
            After=network.target

            [Service]
            User=jim380
            Restart=always
            ExecStart=/usr/local/bin/node_exporter
            Restart=always
            RestartSec=3
            LimitNOFILE=4096
            TimeoutStartSec=0

            [Install]
            WantedBy=multi-user.target
        EOF

    destination = "/tmp/node_exporter.service"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo mkdir -p /home/jim380/Downloads",
      "curl -Lk https://github.com/prometheus/node_exporter/releases/download/${data.github_release.node_exporter.release_tag}/node_exporter-${trimprefix(data.github_release.node_exporter.release_tag, "v")}.linux-amd64.tar.gz --retry 5 | sudo tar xvf --directory ~/Downloads",
      "sudo mv /home/jim380/Downloads/node_exporter-${trimprefix(data.github_release.node_exporter.release_tag, "v")}.linux-amd64 /usr/local/bin",
      "sudo mv /tmp/node_exporter.service /etc/systemd/system",
      "sudo systemctl daemon-reload",
      "sudo systemctl enable node_exporter",
      "sudo systemctl start node_exporter",
    ]

    connection {
      type = "ssh"
      host = self.main_ip
      user     = "root"
      password = self.default_password
    }
  }
}

resource "vultr_startup_script" "script" {
  name   = "add a new system user"
  script = "c3VkbyB1c2VyYWRkIC1tIC1yIGppbTM4MA=="
}