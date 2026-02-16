#!/usr/bin/env sh
go build -C ../gopls/gopls -o ~/.local/bin/gopls-devel -tags pipetransport .
