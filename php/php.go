package php

// #cgo CFLAGS: -I. -I/usr/include/php5 -I/usr/include/php5/main -I/usr/include/php5/TSRM -I/usr/include/php5/Zend -I/usr/include/php5/ext -I/usr/include/php5/
// #cgo LDFLAGS: -L. -lphp5
// void php_startup(void);
// void php_exec_file(char *filename);
// void php_eval_script(char *script);  
// void php_shutdown(void);
import "C"

type PHP struct {}

func NewPHP() (php *PHP) {
    php = &PHP{}
    C.php_startup()
    return
}

func (php * PHP) Exec(filepath string) {
    C.php_exec_file(C.CString(filepath))
}

func (php * PHP) Eval(script string) {
    C.php_eval_script(C.CString(script))
}

func (php * PHP) Close() {
    C.php_shutdown()
}
