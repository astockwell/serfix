serfix
======

Fix PHP serialized objects in MySQL dumps after find/replace.


About
-----

__`serfix` corrects the character counts in PHP serialized string objects, within a MySQL dump file (.sql), when you need to do a mass change of the strings themselves.__

`serfix` was conceived out of a need to automate complex Wordpress development and deployment workflows for a large number of projects. Specifically, when changing environments between development/staging/production, a find/replace tool (such as `sed`) can be used on the database dumps to update the site's root URL, however this will generally have unforseen and undesireable consequences when it breaks the PHP serialized objects in the dump file.

`serfix` was originally a Python script, but for a variety of reasons it was desirable to write it in a compiled binary form.


Build
-----

`go get github.com/astockwell/serfix`

-or-

1. Clone & `cd` into the repo
2. `go build serfix.go` (requires go compiler installed)
3. Copy the resulting `serfix` binary into your path


Usage
-----

`serfix` can be used as a stdin line filter (e.g. chained together with pipes and other unix commands) or as a standalone program.

`serfix [flags] filename [outfilename]`

#### Flags

- `-f`, `--force`: Force overwrite of destination file if it exists (only used with `[outfilename]` specified)
- `-h`, `--help`: Print `serfix` help

#### Line Filter Examples

```
cat filename.sql | serfix > fixed_filename.sql

ssh -C user@host mysqldump --single-transaction --opt \
--net_buffer_length=75000 -u'username' -p'password' db_name \
| sed 's/development.com/production.com/g' | serfix | gzip > \
db_name_$(date +"%Y.%m.%d_%H.%M").sql.gz
```

#### Standalone Examples

```
serfix myfile.sql   # saves over source file

serfix myfile.sql new_filename.sql   # saves to new file

serfix -f myfile.sql existing_filename.sql   # saves over destination file
```


Workflow
--------

#### 1. Dump MySQL database (say, from your dev environment)

Gives you a .sql file full of `s:20:"http://mydevsite.com";` PHP serialized string objects.

#### 2. Find/replace 'mydevsite.com' with 'myproductionsite.com' in your .sql file

Gives you a .sql file full of `s:20:"http://myproductionsite.com";` objects that PHP will reject because the character count doesn't match the string.

#### 3. Run `serfix`

Gives you a .sql file full of `s:27:"http://myproductionsite.com";` objects with the character count correctly matching the string.

#### 4. Import MySQL dump (say, to your prod environment)

High fives all around.


Considerations
--------------

`serfix` uses a 2MB buffer of RAM, so line lengths exceeding 2MB in your .sql file or stdin will cause an error. This can be mitigated by using the `--net_buffer_length=75000` flag in your `mysqldump` commands, which is good practice anyway. `--net_buffer_length` of <=100000 have been tested without issue.


Benchmarks
----------

`serfix` is fast but not wildly so. It should run on an average Wordpress database in <1s. In testing, it took ~38s to fix a 250MB .sql file that contained ~940k regexp match/fixes (dumped with `--net_buffer_length=75000`) on a 2 x 2.4GHz Quad-Core Xeon Mac Pro with 16GB RAM. YMMV.


Roadmap
-------

- Expand test suite
- Remove the second regexp search that is called on every match to find the submatches (this seems to be necessary with the Go standard regexp package, but may not be in the future)
- Better yet, rewrite all regexp operations using a proper lexer


Background
----------

An example of a PHP serialized string is `s:20:"http://mywebsite.com";`. The number is the character count (with some wrinkles) of the subsequent string, so if the string is changed, PHP will choke when this entry is parsed back into an object.

The need to modify a value (the char count) based on another value (the string), plus the "wrinkes", makes this very difficult (if not impossible) to do with basic unix tools (`awk`, `sed`).
