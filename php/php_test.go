package php

import (
    "os"
    "testing"
)

const(
    FileName = "test.php"
)

func TestExec(t *testing.T) {
    f, err := os.OpenFile(FileName, os.O_CREATE, 0644)
    if err != nil {
        t.Errorf("os.OpenFile: %s", err)
    }
    defer func() {
        f.Close()
        if err := os.Remove(FileName); err != nil {
            t.Errorf("os.Remove: %s", err)
        }
    }()

    if ret, err := f.WriteString("<?php phpinfo();"); err != nil {
        t.Errorf("File.WriteString: %s", err)
    } else {
        t.Logf("File.WriteString: %d", ret)
    }

    php := NewPHP()
    php.Exec(FileName)
    php.Close()
}

func TestEval(t *testing.T) {
    php := NewPHP()
    php.Eval("phpinfo();")
    php.Close()
}
