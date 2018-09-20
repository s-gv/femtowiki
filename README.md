femtowiki
=========

[Femtowiki](www.goodoldweb.com/femtowiki/) is a tiny wiki written in golang.
It allows users to create pages using markdown.
The wiki has few dependencies and uses very little javascript. It doesn't need
a database system and can also be used as a static site generator.
Try the latest version hosted [here](https://wiki.goodoldweb.com/).
Please contact [info@goodoldweb.com](mailto:info@goodoldweb.com) if you have any questions or want support.

How to use
----------

[Download](https://github.com/s-gv/femtowiki/releases) the binary.

- To use this as a static site generator, do the following:

Create a few markdown files (with extension .md) in a directory, say `/home/user/wikiroot/`.

Then, run:

```
./femtowiki -wikiroot <path/to/wikiroot>
```

A HTML file is created for every markdown file using a default template (see `templates/` in this repo).
To specify custom templates, use:

```
./femtowiki -wikiroot <path/to/wikiroot> -templateroot <path/to/templates>
```

- If you'd like to be able to edit the wiki in a browser, then create a new user with:

TODO: Current version does not yet support this!

```
./femtowiki -createuser -users <path/to/wikiusers.json>
```

Start the server:

```
./femtowiki -wikiroot <path/to/wikiroot> -users <path/to/wikiusers.json> -serve
```

Dependencies
------------

- Go 1.8 (only for compiling)

Options
-------

- `-addr <port>`: Use `./femtowiki -addr :8086` to listen on port 8086.
- `-wikiroot <path/to/wikiroot>`: Specify directory containing markdown files.
- `-templateroot <path/to/templateroot>`: Specify directory containing template files.
- `-users <path/to/users.json>`: Specify file containing list of users.

Commands
--------

- `-help`: Show a list of all commands and options.
- `-createuser`: Create a new user.
- `-changepasswd`: Change password of a user.
- `-serve`: Run the online editor.

