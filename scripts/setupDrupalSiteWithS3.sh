# Move the website files to the top level
mv /var/www/html/drupal-7.8/* /var/www/html
mv /var/www/html/drupal-7.8/.htaccess /var/www/html
rm -Rf /var/www/html/drupal-7.8

# Mount the S3 bucket
mv /var/www/html/sites/default/files /var/www/html/sites/default/files_original
mkdir -p /var/www/html/sites/default/files
sh ~ec2-user/mountS3Bucket.sh

# Make changes to Apache Web Server configuration
sed -i 's/AllowOverride None/AllowOverride All/g'  /etc/httpd/conf/httpd.conf
service httpd restart

# Only execute the site install if we are the first host up - otherwise we'll end up losing all the data
read first < /var/www/html/sites/default/files/hosts
if [ `hostname` = $first ]
then
  
  sh ~ec2-user/installDrupal.sh

  # use the S3 bucket for shared file storage
  cp -R sites/default/files_original/* sites/default/files
  cp -R sites/default/files_original/.htaccess sites/default/files
fi
# Copy settings.php file since everything else is configured
cp /home/ec2-user/settings.php /var/www/html/sites/default
rm /home/ec2-user/settings.php
