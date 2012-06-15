package php

import (
    "os"
    "io/ioutil"
    "testing"
)

const(
    FileName = "test.php"
)

func TestExec(t *testing.T) {
    if err := ioutil.WriteFile(FileName, []byte("<?php echo 'TestExec\n';"), 0644); err != nil {
        t.Errorf("ioutil.WriteFile: %s", err)
    }
    defer func() {
        if err := os.Remove(FileName); err != nil {
            t.Errorf("os.Remove: %s", err)
        }
    }()
    php := NewPHP("/etc/php5/cli/")
    defer php.Close()
    if err := php.Exec(FileName); err != nil {
        t.Errorf("php.Exec: %s", err)
    }
}

func TestEval(t *testing.T) {
    php := NewPHP("/etc/php5/cli/")
    defer php.Close()
    if err := php.Eval("echo 'TestEval\n';"); err != nil {
        t.Errorf("php.Eval: %s", err)
    }
}

func TestEvalErr(t *testing.T) {
    php := NewPHP("/etc/php5/cli/")
    defer php.Close()
    if err := php.Eval("echo 'TestEval\n'"); err == nil {
        t.Errorf("php.Eval should have panic.")
    } else {
        t.Logf("php.Eval: ", err)
    }
}

func _TestInfo(t *testing.T) {
    php := NewPHP("/etc/php5/cli/")
    defer php.Close()
    if err := php.Info(); err != nil {
        t.Errorf("php.Info: %s", err)
    }
}
