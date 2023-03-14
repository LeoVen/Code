#!/bin/bash
yum update -y
yum install httpd -y
service httpd start
chkconfig httpd on
echo "<html><h1>Hello, World!</h1><p>Welcome</p></html>" > /var/www/html/index.html