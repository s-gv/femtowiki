femtowiki
=========

[Femtowiki](www.goodoldweb.com/femtowiki/) is a tiny wiki written in golang.
It allows users to create pages using markdown.
The wiki few dependencies and uses very little javascript.
Try the latest version hosted [here](https://wiki.goodoldweb.com/).
Please contact [info@goodoldweb.com](mailto:info@goodoldweb.com) if you have any questions or want support.

How to use
----------

[Download](https://github.com/s-gv/femtowiki/releases) the binary.
Create a new file `config.json` with the contents below:

```
{
    "wikiroot": "./wiki", // Replace with the path where you want the wiki files
    "users": "./users.json", // Replace with where you want a file having a list of users
}
```

Create a new user:

```
./femtowiki -config ./config.json -createuser
```

Start the server:

```
./femtowiki -config ./config.json
```

Dependencies
------------

- Go 1.8 (only for compiling)

Options
-------

- `-addr <port>`: Use `./femtowiki -addr :8086` to listen on port 8086.
- `-config <path/to/config.json>`: Use `./femtowiki -config /home/user/config.json` to use `/home/user/config.json` as the the config file.

Commands
--------

- `-help`: Show a list of all commands and options.
- `-createuser`: Create a new user.
- `-changepasswd`: Change password of a user.

