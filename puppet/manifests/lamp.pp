
class httpd {

  exec { 'yum-update':
    command => '/usr/bin/yum -y update'
  }

  package { "httpd":
    ensure => present
  }

  package { "httpd-devel":
    ensure  => present
  }

  service { 'httpd':
    name      => 'httpd',
    require   => Package["httpd"],
    ensure    => running,
    enable    => true
  }

  file { "/etc/httpd/conf.d/vhost.conf":
    owner   => "root",
    group   => "root",
    mode    => 644,
    replace => true,
    ensure  => present,
    source  => "/home/ec2-user/vhost.conf",
    require => Package["httpd"],
    notify  => Service["httpd"]
  }

}

class php {

  package { "php":
    ensure  => present,
  }

  package { "php-cli":
    ensure  => present,
  }

  package { "php-common":
    ensure  => present,
  }

  package { "php-devel":
    ensure  => present,
  }

  package { "php-gd":
    ensure  => present,
  }

  package { "php-mcrypt":
    ensure  => present,
  }

  package { "php-intl":
    ensure  => present,
  }

  package { "php-ldap":
    ensure  => present,
  }

  package { "php-mbstring":
    ensure  => present,
  }

  package { "php-mysql":
    ensure  => present,
  }

  package { "php-pdo":
    ensure  => present,
  }

  package { "php-pear":
    ensure  => present,
  }

  package { "php-pecl-apc":
    ensure  => present,
  }

  package { "php-soap":
    ensure  => present,
  }

  package { "php-xml":
    ensure  => present,
  }

  package { "uuid-php":
    ensure  => present,
  }

}

class mysql {

  package { "mysql":
    ensure  => present,
  }

}

class devel {

  package { "gcc":
    ensure  => present,
  }

  package { "make":
    ensure  => present,
  }

  package { "libstdc++-devel":
    ensure  => present,
  }

  package { "gcc-c++":
    ensure  => present,
  }

  package { "fuse":
    ensure  => present,
  }

  package { "fuse-devel":
    ensure  => present,
  }

  package { "libcurl-devel":
    ensure  => present,
  }

  package { "libxml2-devel":
    ensure  => present,
  }

  package { "openssl-devel":
    ensure  => present,
  }

}

class s3fs {

  exec { 'install-s3fs':
    command => '/bin/bash /home/ec2-user/installS3fs.sh',
    require => Class['devel'],
    logoutput => on_failure,
  }

}

class drupal {

  exec { 'install-drupal':
    command => '/bin/bash /home/ec2-user/setupDrupalSiteWithS3.sh',
    require => Class['s3fs'],
    logoutput => on_failure,
  }

}

include httpd
include php
include mysql
include devel
include s3fs
include drupal
