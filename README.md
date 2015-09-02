# GoEmPHP

This package is built for Embedding PHP into Golang.

It is easy to use:

    script = php.New()
    script.Startup()
    defer script.Close()
    if err := script.Eval("phpinfo();"); err != nil {
        log.Fatal(err)
    }
    if err := script.Exec("foobar.php"); err != nil {
        log.Fatal(err)
    }

For more examples, please read the souce code: `php_test.go`.

# INSTALL

> $ go get github.com/mikespook/goemphp/php
	
# Contacts

 * Xing Xing <mikespook@gmail.com>

 * [Blog](http://mikespook.com)

 * [@Twitter](http://twitter.com/mikespook)
