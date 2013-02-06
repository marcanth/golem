golem
=====

golem - file-copy and backup tool

An animated anthropomorphic being, created entirely from inanimate matter.


Installing golem
----------------

	git clone git://github.com/joaosoda/golem.git
	go build -o golem golem/main.go

Usage
-----

	golem <source> <dest>

config.json
-----------

Put a **config.json** inside your project to tell commands something to *golem*.

###Ignored
Ignore specified files and/or folders.

Exemple

	{
	    "Ignored": ["file1", "file2", "folder/file3"]
	}

With this configuration, golem not will copy and backup *file1*, *file2* and *file3* in *folder*
