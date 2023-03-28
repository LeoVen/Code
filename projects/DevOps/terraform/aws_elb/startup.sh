#!/bin/bash
yum update -y
yum install -y httpd
systemctl start httpd
systemctl enable httpd
INTERFACE=$(curl -s http://169.254.169.254/latest/meta-data/network/interfaces/macs/)
SUBNETID=$(curl -s http://169.254.169.254/latest/meta-data/network/interfaces/macs/${INTERFACE}/subnet-id)
INSTANCEID=$(curl -s http://169.254.169.254/latest/meta-data/instance-id)
echo '<html><center><h1>This instance (INSTANCEID) is in the subnet wih ID: SUBNETID </h1></center></html>' > /var/www/html/index.txt
sed "s/SUBNETID/$SUBNETID/;s/INSTANCEID/$INSTANCEID/" /var/www/html/index.txt > /var/www/html/index.html