#!/bin/bash

mysqldump --column-statistics=0 -h 0.0.0.0 -P 3306  -uroot -p123qwe ecode > ecode-`date "+%Y-%m-%d-%H-%M-%S"`.sql