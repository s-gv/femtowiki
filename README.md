femtowiki
=========

[Femtowiki](http://www.goodoldweb.com) is a tiny wiki written in golang. It allows users to create pages using markdown.
The wiki few dependencies and uses very little javascript.
Try the latest version hosted [here](https://wiki.goodoldweb.com). Please contact [info@goodoldweb.com](mailto:info@goodoldweb.com)
if you have any questions or want support.

How to use
----------

By default, sqlite is used, so it's easy to get started.
[Download](https://github.com/s-gv/femtowiki/releases) the binary and migrate the database with:

```
./femtowiki -migrate
```

Create a superadmin:

```
./femtowiki -createsuperuser
```

Finally, start the server:

```
./femtowiki
```

Notes
-----

Femtowiki allows users to create wiki pages using markdown.
Admin users control who can perform various actions with the following level of granularity.

- Users belonging to the `PageMaster` group can create, modify, and delete all pages.
- Users belonging to the `FileMaster` group can upload and delete files.
- For users outside the `PageMaster` and `FileMaster` groups, read and edit rights can be restricted at the level
of individual pages and files.
- Groups can be created and modified only by admin users.
- The special group `everybody` represents all users and cannot be modified.
- New user registration can be disabled (but admin users can still signup users).


Dependencies
------------

- Go 1.8 (only for compiling)
- Postgres 9.5 (or use embedded sqlite3)

Options
-------

- `-addr <port>`: Use `./femtowiki -addr :8086` to listen on port 8086.
- `-dbdriver <db>` and `-dsn <data_source_name>`: PostgreSQL and SQLite are supported. SQLite is the default driver.

To use postgres, run `./femtowiki -dbdriver postgres -dsn postgres://pguser:pgpasswd@localhost/dbname`

To save an sqlite db at a different location, run `./femtowiki -dsn path/to/mywiki.db`.

Commands
--------

- `-help`: Show a list of all commands and options.
- `-migrate`: Migrate the database. Run this once after updating the femtowiki binary (or when starting afresh).
- `-createsuperuser`: Create a super admin.
- `-createuser`: Create a new user with no special privileges.
- `-changepasswd`: Change password of a user.
- `-deletesessions`: Drop all sessions and log out all users.