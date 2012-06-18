# GoEmPHP

This package is built for Embedding PHP into Golang.

It is easy to use:

    php = NewPHP()
    php.Startup()
    defer php.Close()
    if err := php.Eval("phpinfo();"); err != nil {
        log.Fatal(err)
    }
    if err := php.Exec("foobar.php"); err != nil {
        log.Fatal(err)
    }

For more examples, please read the php\_test.php.

# INSTALL

> $ go get bitbucket.org/mikespook/goemphp/php
	
# Contacts

Xing Xing <mikespook@gmail.com>

[Blog](http://mikespook.com)

[@Twitter](http://twitter.com/mikespook)
