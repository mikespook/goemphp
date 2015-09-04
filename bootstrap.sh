#!/usr/bin/env bash
set -e

help() {
	printf "Usage: $0 [PHP version]\n\n"
	printf " * The PHP version can be one of 5.4, 5.5, 5.6 or ng.\n"
}

[ -z $1 ] && help && exit

PHP_VER=$1

PHP_5_6_D=php-5.6.12
PHP_5_5_D=php-5.5.28
PHP_5_4_D=php-5.4.44
PHP_NG_D=php-7.0.0RC1

PHP_5_6=http://php.net/get/${PHP_5_6_D}.tar.bz2/from/this/mirror
PHP_5_5=http://php.net/get/${PHP_5_5_D}.tar.bz2/from/this/mirror
PHP_5_4=http://php.net/get/${PHP_5_4_D}.tar.bz2/from/this/mirror
PHP_NG=http://downloads.php.net/~ab/${PHP_NG_D}.tar.bz2

case $PHP_VER in
5.4)
	TAR=$PHP_5_4
	DIR=$PHP_5_4_D
  ;;
5.5)
	TAR=$PHP_5_5
	DIR=$PHP_5_5_D 
  ;;
5.6)
	TAR=$PHP_5_6
	DIR=$PHP_5_6_D
  ;;
*)
  	TAR=$PHP_NG
	DIR=$PHP_NG_D 
  ;;
esac

WORK_DIR=php-srcs
mkdir -p $WORK_DIR
[ -d $WORK_DIR/$DIR ] && echo "FATAL: $DIR already existed. Please remove $DIR manually and run $0 again." && exit

pushd .
cd $WORK_DIR

wget -O php.tar.bz2 $TAR
tar jxvf php.tar.bz2
rm php.tar.bz2
cd $DIR 
./configure --build=x86_64-linux-gnu --host=x86_64-linux-gnu --sysconfdir=/etc --localstatedir=/var --mandir=/usr/share/man --disable-debug --disable-rpath --disable-static --with-pic --with-layout=GNU --with-pear=/usr/share/php --enable-calendar --enable-sysvsem --enable-sysvshm --enable-sysvmsg --enable-bcmath --with-bz2 --enable-ctype --with-db4 --without-gdbm --with-iconv --enable-exif --enable-ftp --with-gettext --enable-mbstring --with-pcre-regex=/usr --enable-shmop --enable-sockets --enable-wddx --with-libxml-dir=/usr --with-zlib --with-kerberos=/usr --with-openssl=/usr --enable-soap --enable-zip --with-mhash=yes --with-mysql-sock=/var/run/mysqld/mysqld.sock --enable-dtrace --without-mm --with-curl=shared,/usr --with-enchant=shared,/usr --with-zlib-dir=/usr --with-gd=shared,/usr --enable-gd-native-ttf --with-gmp=shared,/usr --with-jpeg-dir=shared,/usr --with-xpm-dir=shared,/usr/X11R6 --with-png-dir=shared,/usr --with-freetype-dir=shared,/usr --enable-intl=shared --with-ldap=shared,/usr --with-ldap-sasl=/usr --with-mysqli=shared,/usr/bin/mysql_config --with-pspell=shared,/usr --with-unixODBC=shared,/usr --with-recode=shared,/usr --with-xsl=shared,/usr --with-snmp=shared,/usr --with-tidy=shared,/usr --with-xmlrpc=shared --with-pgsql=shared,/usr --enable-embed --with-libdir=/lib/x86_64-linux-gnu
make
popd
echo $PHP_VER > php-version
unlink php-lib
ln -s $WORK_DIR/$DIR php-lib

echo "Congratulations!!!"
