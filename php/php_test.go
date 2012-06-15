package php

import (
    "os"
    "io/ioutil"
    "testing"
)

const(
    FileName = "test.php"
)

func init() {
//    os.Stderr.Close()
//    os.Stdout.Close()
}

func TestExec(t *testing.T) {
    if err := ioutil.WriteFile(FileName, []byte("<?php echo 'TestExec\n';"), 0644); err != nil {
        t.Errorf("ioutil.WriteFile: %s", err)
    }
    defer func() {
        if err := os.Remove(FileName); err != nil {
            t.Errorf("os.Remove: %s", err)
        }
    }()
    php := NewPHP()
    php.Startup()
    defer php.Close()
    if err := php.Exec(FileName); err != nil {
        t.Errorf("php.Exec: %s", err)
    }
}

func TestExecErr(t *testing.T) {
    if err := ioutil.WriteFile(FileName, []byte("<?php echo 'TestExec\n'"), 0644); err != nil {
        t.Errorf("ioutil.WriteFile: %s", err)
    }
    defer func() {
        if err := os.Remove(FileName); err != nil {
            t.Errorf("os.Remove: %s", err)
        }
    }()
    php := NewPHP()
    php.Startup()
    defer php.Close()
    if err := php.Exec(FileName); err == nil {
        t.Errorf("php.Exec should have a panic.")
    } else {
        t.Logf("php.Exec: %s", err)
    }
}

func TestExecNotFound(t *testing.T) {
    php := NewPHP()
    php.Startup()
    defer php.Close()
    if err := php.Exec("not-found.php"); err == nil {
        t.Errorf("php.Exec should have a panic.")
    } else {
        t.Logf("php.Exec: %s", err)
    }
}

func TestEval(t *testing.T) {
    php := NewPHP()
    php.Startup()
    defer php.Close()
    if err := php.Eval("echo 'TestEval';"); err != nil {
        t.Errorf("php.Eval: %s", err)
    }
}

func TestEvalErr(t *testing.T) {
    php := NewPHP()
    php.Startup()
    defer php.Close()
    if err := php.Eval("echo 'TestEval\n'"); err == nil {
        t.Errorf("php.Eval should have a panic.")
    } else {
        t.Logf("php.Eval: %s", err)
    }
}
