# mylinter, is a MySQL Migration Lint
This tool help the data definition language review process for developers, the idea is to identifying some common and uncommon mistakes that are made during coding to optimize the data model.

## Configure

This setp is optional, but maybe need ignore specific error code, this is the way: Is very easy to configure, only need to create file witch this name `.mylinter.yaml` on directory you have \*.sql, add folow lines and put the number of errors code to ignore:

```yaml
---
ignore: 302, 303, 304, 305, 406
tables:
  - name: actors
    ignore: 405
```

You can ignore by all tables, or by table. To ignore all tables use first line `ignore: 302, 303, 304, 305, 406`, to ignore specific errors code on table `actors` create a yaml list like this:

```yaml
tables:
  - name: actors
    ignore: 405
  - name: users
    ignore: 504, 405
```

## TODO:

- All messages in array and CLI arg to print.
- Change setting about min/max length data type.
- Verify syntax and spell check for all messages.
- Indexes lint.
- Foreign Key lint.
- Verify cardinality with MySQL and data.
- List all codes and checks.
