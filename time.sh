#!/bin/sh
time ( ./go.out < inputml ) 2>&1 1>/dev/null
time ( ./cpp.out < inputml ) 2>&1 1>/dev/null