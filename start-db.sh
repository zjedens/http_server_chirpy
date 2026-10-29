#!/usr/bin/bash

sudo -u postgres pg_ctl -D /var/lib/pgsql/data -l /var/lib/pgsql/db.log start
