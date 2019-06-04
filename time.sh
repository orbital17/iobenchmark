#!/bin/sh
time ( ./go.out < inputlarge ) 2>&1 1>/dev/null
time ( ./cpp.out < inputlarge ) 2>&1 1>/dev/null