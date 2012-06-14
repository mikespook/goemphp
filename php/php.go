package php

// #cgo CFLAGS: -I. -I/usr/include/php5 -I/usr/include/php5/main -I/usr/include/php5/TSRM -I/usr/include/php5/Zend -I/usr/include/php5/ext -I/usr/include/php5/
// #cgo LDFLAGS: -L. -lphp5
// void php_startup(void);
// void php_execute(char *filename);
// void php_shutdown(void);
import "C"

func ExePHP(filepath string) {
    //var fileHandle C.zend_file_handle
    C.php_startup()
    C.php_shutdown()
}
