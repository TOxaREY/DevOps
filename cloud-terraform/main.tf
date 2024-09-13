terraform {
  backend "s3" {
    endpoint    = "https://storage.yandexcloud.net"
    bucket      = "terraform-state-toxarey"
    region      = "ru-central1"
    key         = "terraform.tfstate"

    skip_region_validation      = true
    skip_credentials_validation = true
    skip_requesting_account_id  = true
    skip_s3_checksum            = true
  }
}

locals {
  folder_id   = var.yc_folder_id
}

resource "yandex_kubernetes_cluster" "k8s-toxarey" {
  name = "k8s-toxarey"
  network_id = yandex_vpc_network.toxareynet.id
  master {
    master_location {
      zone      = yandex_vpc_subnet.toxareysubnet.zone
      subnet_id = yandex_vpc_subnet.toxareysubnet.id      
    }
    public_ip = true
  }
  service_account_id      = yandex_iam_service_account.toxareyaccount.id
  node_service_account_id = yandex_iam_service_account.toxareyaccount.id
  depends_on = [
    yandex_resourcemanager_folder_iam_member.k8s-clusters-agent,
    yandex_resourcemanager_folder_iam_member.vpc-public-admin,
    yandex_resourcemanager_folder_iam_member.images-puller,
    yandex_resourcemanager_folder_iam_member.editor,
    yandex_resourcemanager_folder_iam_member.load-balancer-admin
  ]
}

resource "yandex_kubernetes_node_group" "k8s-toxarey-ng" {
  name        = "k8s-toxarey-ng"
  cluster_id  = yandex_kubernetes_cluster.k8s-toxarey.id
  version     = "1.27"
  instance_template {
    name = "toxarey-{instance.short_id}-{instance_group.id}"
    platform_id = "standard-v3"
    resources {
      cores         = 2
      core_fraction = 50
      memory        = 4
    }
    boot_disk {
      size = 30
      type = "network-ssd"
    }
    network_acceleration_type = "standard"
    network_interface {
      nat                = true
      subnet_ids         = ["${yandex_vpc_subnet.toxareysubnet.id}"]
    }
    scheduling_policy {
      preemptible = true
    }
  }
  scale_policy {
    fixed_scale {
      size = 2
    }
  }
  deploy_policy {
    max_expansion   = 3
    max_unavailable = 1
  }
  maintenance_policy {
    auto_upgrade = true
    auto_repair  = true
    maintenance_window {
      start_time = "22:00"
      duration   = "10h"
    }
  }
  node_labels = {
    node-label1 = "node-value1"
  }
  labels = {
    "template-label1" = "template-value1"
  }
  allowed_unsafe_sysctls = ["kernel.msg*", "net.core.somaxconn"]
}

resource "yandex_vpc_network" "toxareynet" {
  name = "toxareynet"
}

resource "yandex_vpc_subnet" "toxareysubnet" {
  name = "toxareysubnet"
  v4_cidr_blocks = ["10.1.0.0/16"]
  zone           = var.yc_zone
  network_id     = yandex_vpc_network.toxareynet.id
}

resource "yandex_iam_service_account" "toxareyaccount" {
  name        = "zonal-k8s-account"
  description = "K8S zonal service account"
}

resource "yandex_resourcemanager_folder_iam_member" "k8s-clusters-agent" {
  folder_id = local.folder_id
  role      = "k8s.clusters.agent"
  member    = "serviceAccount:${yandex_iam_service_account.toxareyaccount.id}"
}

resource "yandex_resourcemanager_folder_iam_member" "vpc-public-admin" {
  folder_id = local.folder_id
  role      = "vpc.publicAdmin"
  member    = "serviceAccount:${yandex_iam_service_account.toxareyaccount.id}"
}

resource "yandex_resourcemanager_folder_iam_member" "images-puller" {
  folder_id = local.folder_id
  role      = "container-registry.images.puller"
  member    = "serviceAccount:${yandex_iam_service_account.toxareyaccount.id}"
}

resource "yandex_resourcemanager_folder_iam_member" "editor" {
  folder_id = local.folder_id
  role      = "editor"
  member    = "serviceAccount:${yandex_iam_service_account.toxareyaccount.id}"
}

resource "yandex_resourcemanager_folder_iam_member" "load-balancer-admin" {
  folder_id = local.folder_id
  role      = "load-balancer.admin"
  member    = "serviceAccount:${yandex_iam_service_account.toxareyaccount.id}"
}