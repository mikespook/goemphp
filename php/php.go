package php

// #cgo CFLAGS: -I. -I/usr/include/php5 -I/usr/include/php5/main -I/usr/include/php5/TSRM -I/usr/include/php5/Zend -I/usr/include/php5/ext -I/usr/include/php5/
// #cgo LDFLAGS: -L. -lphp5
// void php_startup(char *ini);
// char * php_exec_file(char *filename);
// char * php_eval_script(char *script);  
// void php_shutdown(void);
// int php_info(void);
import "C"

import (
    "errors"
)

var (
    ErrInternal = errors.New("Internal Error.")
)

type PHP struct {}

func NewPHP(ini string) (php *PHP) {
    php = &PHP{}
    C.php_startup(C.CString(ini))
    return
}

func (php * PHP) Exec(filepath string) (err error) {
    if err := C.php_exec_file(C.CString(filepath)); err != nil {
        return errors.New(C.GoString(err))
    }
    return
}

func (php * PHP) Eval(script string) (err error) {
    if err := C.php_eval_script(C.CString(script)); err != nil {
        return errors.New(C.GoString(err))
    }
    return
}

func (php * PHP) Close() {
    C.php_shutdown()
}

func (php *PHP) Info() error {
    if 0 != C.php_info() {
        return ErrInternal
    }
    return nil
}
