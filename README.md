# mylinter, is a MySQL Migration Lint
This tool help the data definition language review process for developers, the idea is to identifying some common and uncommon mistakes that are made during coding to optimize the data model.

For example, if you lint this simple model:

```sql
CREATE TABLE `userExternal` (
    `id` INT(20),
    `user.name` VARCHAR(1024),
    `status` CHAR(255),
    `description` VARCHAR(2000) DEFAULT NULL,
    `update_at` datetime,
    PRIMARY KEY (`id`)
)ENGINE=MyISAM;
```

Have the follow output:

```
> File: assets/examples/case01.sql
- [302] Table engine is not InnoDB.
- [303] Table charset is not set to use UTF8.
- [304] Table collate is not set to use UTF8.
- [305] Table no have description.
- [311] Table name has capital letter: userExternal.
- [404] Field name contains dot's, please remove it: user.name
- [406] Field should by have comment: id
- [406] Field should by have comment: user.name
- [406] Field should by have comment: status
- [406] Field should by have comment: description
- [406] Field should by have comment: update_at
- [407] Field with char type should by have length less than 50 chars: CHAR(255)
- [408] Field varchar type with length great than 255 should by text type: VARCHAR(1024)
- [408] Field varchar type with length great than 255 should by text type: VARCHAR(2000)
- [409] Field datetime type is defined, should by timestamp: update_at
- [503] Primary Key field should by NOT NULL: id
- [504] Primary key field must be BIGINT: id INT
- [505] Primary Key field should by unsigned: id
- [506] Primary Key field should by auto increment: id
```

## Configure

This step is optional, but maybe need ignore specific error code, this is the way: Is very easy to configure, only need to create file witch this name `.mylinter.yaml` on directory you have \*.sql, add folow lines and put the number of errors code to ignore:

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
