#!/usr/bin/env bash
go test -ldflags="-r ./php-lib/libs/" $*
