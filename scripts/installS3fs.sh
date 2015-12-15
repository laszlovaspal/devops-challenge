cd /home/ec2-user
wget https://s3fs.googlecode.com/files/s3fs-1.74.tar.gz
tar xzf s3fs-1.74.tar.gz
cd s3fs-1.74
./configure --prefix=/usr
make
make install