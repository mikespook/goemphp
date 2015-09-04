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

The first step is downloading the package. Please note that you must specify
the param `-d`.

> $ go get -d github.com/mikespook/goemphp

Then, following step is running `bootstrap.sh` to prepar the embeded PHP 
library. It has one paramater with 4 values: [5.4 | 5.5 | 5.6 | ng]

Eg.,

> `./bootstrap.sh 5.6`

will `wget`, `configure` and `make` the libphp5.so file which can be used for embedding PHP.

It will be a long time waiting. When you see `Congratulations!!!`, it means you have already got the proper .so file placed in `./php-lib/libs/`. You could check it manually.

The third step is calling `go generate` to prepare source files which will be built.

If generate process do not report any issue, then you could call `go build` to build GoEmPHP.

After that, please use `./test.sh` for testing the package. And of course, you could run `go test -ldflags="-r ./php-lib/libs/"` manually, or put the .so into one of system library directories and run `go test`. The same library mechanism should be used when you use this library in your application.

# Contacts

 * Xing Xing <mikespook@gmail.com>

 * [Blog](http://mikespook.com)

 * [@Twitter](http://twitter.com/mikespook)

# Open Source

See LICENSE for more information.

