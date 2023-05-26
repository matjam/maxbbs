# What is MaxBBS?

MaxBBS is a toy BBS server that will at some point implement most common BBS
functions such as message areas, file areas, chat, doors, etc.

MaxBBS is designed as a homage to the very fine Maximus BBS; specifically it
is intended to support a version of MECCA very close to the original.

MECCA is a templating language for MaxBBS that allows you to build your BBS
screens using UTF-8 text with tokens, and these templates are converted into
whatever codepage you want such as [Code page 437](https://en.wikipedia.org/wiki/Code_page_437). which was commonly used
by BBSs of the era to display ANSI art.

The tokens provide a way to substitute in data from MaxBBS such as the number
of messages in the current board, who is logged on, etc. These templates
are not really a programming language, but there is conditionals and the
ability to "goto" another part of the template so there is support for
displaying different data depending on who is logged on, etc. This feature was
one of the defining features of MaximusBBS and I wanted to bring it along if
I could.

# Current status

Complete

- You can telnet to it?

In Progress

- MECCA Tokenizer implementation
- Data model
- Telnet interface

Planned

- FTP Server access to fileareas
- IRC access to chat
- SMTP support for message delivery
- Message areas
- File areas
