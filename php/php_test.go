// Copyright 2011 Xing Xing <mikespook@gmail.com> All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package php

import (
    "os"
    "io/ioutil"
    "testing"
)

const(
    FileName = "test.php"
    TestFileName = "test.txt"
)

var (
    php *PHP
)

func init() {
    php = NewPHP()
    php.Startup()
}

func TestExec(t *testing.T) {
    if err := ioutil.WriteFile(FileName, []byte("<?php echo 'TestExec\n';file_put_contents('test.txt', 'abcdef');"), 0644); err != nil {
        t.Errorf("ioutil.WriteFile: %s", err)
    }
    defer func() {
        if err := os.Remove(FileName); err != nil {
            t.Errorf("os.Remove: %s", err)
        }
        if err := os.Remove(TestFileName); err != nil {
            t.Errorf("os.Remove: %s", err)
        }
    }()
    if err := php.Exec(FileName); err != nil {
        t.Errorf("php.Exec: %s", err)
    }

    if _, err := os.Stat(TestFileName); err != nil {
        t.Errorf("php::file_put_contents: %s", err)
    }
}

func TestExecErr(t *testing.T) {
    if err := ioutil.WriteFile(FileName, []byte("<?php echo TestExec"), 0644); err != nil {
        t.Errorf("ioutil.WriteFile: %s", err)
    }
    defer func() {
        if err := os.Remove(FileName); err != nil {
            t.Errorf("os.Remove: %s", err)
        }
    }()
    if err := php.Exec(FileName); err == nil {
        t.Errorf("php.Exec should have a panic.")
    }
}

func TestExecNotFound(t *testing.T) {
    if err := php.Exec("not-found.php"); err == nil {
        t.Errorf("php.Exec should have a panic.")
    }
}

func TestEval(t *testing.T) {
    if err := php.Eval("echo 'TestEval\n';"); err != nil {
        t.Errorf("php.Eval: %s", err)
    }
}

func TestEvalThrowExp(t *testing.T) {
    if err := php.Eval("throw new Exception('test exception');"); err == nil {
        t.Errorf("php.Eval should have a panic.")
    }
}

func TestEvalTriggerError(t *testing.T) {
    if err := php.Eval("trigger_error('test error');"); err == nil {
        t.Errorf("php.Eval should have a panic.")
    }
}

func TestEvalErr(t *testing.T) {
    if err := php.Eval("echo ;"); err == nil {
        t.Errorf("php.Eval should have a panic.")
    }
}

func TestVar(t *testing.T) {
    if err := php.Var("v", "test"); err != nil {
        t.Errorf("TestArgs: %s", err)
    }

    if err := php.Eval("var_dump($v);"); err != nil {
        t.Errorf("TestArgs: %s", err)
    }

    if err := php.Var("v", true); err != nil {
        t.Errorf("TestArgs: %s", err)
    }

    if err := php.Eval("var_dump($v);"); err != nil {
        t.Errorf("TestArgs: %s", err)
    }

    if err := php.Var("v", 123); err != nil {
        t.Errorf("TestArgs: %s", err)
    }

    if err := php.Eval("var_dump($v);"); err != nil {
        t.Errorf("TestArgs: %s", err)
    }

    if err := php.Var("v", 123.456); err != nil {
        t.Errorf("TestArgs: %s", err)
    }

    if err := php.Eval("var_dump($v);"); err != nil {
        t.Errorf("TestArgs: %s", err)
    }
}

func TestArray(t *testing.T) {
    php.Array("v", map[string]string{"a":"z", "b":"y", "c":"x"})

    if err := php.Eval("var_dump($v);"); err != nil {
        t.Errorf("TestArgs: %s", err)
    }
}

func TestUnset(t *testing.T) {
    php.Unset("v")

    if err := php.Eval("var_dump($v);"); err == nil {
        t.Error("php.Eval should have a panic.")
    }
}
