// Copyright 2011 Xing Xing <mikespook@gmail.com> All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package php

// #cgo CFLAGS: -I/usr/include/php5 -I/usr/include/php5/main -I/usr/include/php5/TSRM -I/usr/include/php5/Zend -I/usr/include/php5/ext -I/usr/include/php5/
// #cgo LDFLAGS: -lphp5
// void php_set_ini(char *ini);
// void php_startup();
// char * php_exec_file(char *filename);
// char * php_eval_script(char *script);  
// void php_shutdown(void);
import "C"

import (
    "os"
    "errors"
    "syscall"
)

const (
    Success = 0
    Failure = -1
)

type PHP struct {
    stdout, stderr *os.File
    inifile string
}

func NewPHP() (php *PHP) {
    php = &PHP{
        stdout: os.Stdout,
        stderr: os.Stderr,
    }
    return
}

func (php *PHP) Stdout(f *os.File) {
    php.stdout = f
}

func (php *PHP) Stderr(f *os.File) {
    php.stderr = f
}

func (php *PHP) IniFile(ini string) {
    php.inifile = ini
}

func (php *PHP) Startup() {
    syscall.Dup2(int(php.stdout.Fd()), 1)
    syscall.Dup2(int(php.stderr.Fd()), 2)
    C.php_set_ini(C.CString(php.inifile))
    C.php_startup()
}

func (php * PHP) Exec(filepath string) (err error) {
    if _, err = os.Stat(filepath); err != nil {
        return
    }
    if err := C.php_exec_file(C.CString(filepath)); err != nil {
        return errors.New(C.GoString(err))
    }
    return
//    return php.Eval("require('" + filepath + "');")
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
