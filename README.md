# mylinter, is a MySQL Migration Lint
This tool help the data definition language review process for developers, the idea is to identifying some common and uncommon mistakes that are made during coding to optimize the data model.

For example, analize this create table:

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

If you lint this simple model with the follow command:

```bash
mylinter --path=assets/examples/case01.sql
```

Have the follow output:

```
> File: assets/examples/case01.sql
- [302] Table engine is not InnoDB.
- [303] Table charset is not set to use UTF8.
- [304] Table collate is not set to use UTF8.
- [311] Table name has capital letter: userExternal.
- [404] Field name contains dot's, please remove it: user.name
- [407] Field with char type should by have length less than 50 chars: status CHAR(255)
- [408] Field varchar type with length great than 255 should by text type: user.name VARCHAR(1024)
- [408] Field varchar type with length great than 255 should by text type: description VARCHAR(2000)
- [409] Field datetime type is defined, should by timestamp: update_at
- [503] Primary Key field should by NOT NULL: id
- [504] Primary key field must be BIGINT: id INT
- [505] Primary Key field should by unsigned: id
- [506] Primary Key field should by auto increment: id
```

## Install

```bash
bash < <(curl -s https://debeando.com/mylinter.sh)
```

**NOTE:** Now, the install script support only Linux amd64.

## Configure

This step is optional, but maybe need change default limits or ignore specific error code, this is the way: Is very easy to configure, only need to create file witch this name `.mylinter.yaml` on directory you have \*.sql, add folow lines and adjust by your preferences:

```yaml
---
ignore: 302, 303, 304, 305, 406
fields-max: 20
char-length-max: 51
varchar-length-max: 256
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
