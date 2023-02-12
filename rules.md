# Rules

Cada check un codigo, asi se puede hacer una descripcion del caso y tambien se puede ignorar.

Puede que el nivel de severidad este entre Warning y Error, ver si hay mejor forma.

Interesante sería a futuro que se pueda conectarse al mysql y validar:

- Sino hay rows.
- Que no tiene operaciones sobre ella.
- Cardinalidad de indices.
- Buscar queries relacionadas y revisar tiempos.

## Object Name

- Warn: Que no tenga puntos.
- Warn: Que no empieze por \_.
- Warn: Que no supere cierta longitud.
- Warn: Todo en minuscula.
- Warn: Si la palabra es muy larga y no hay _ sugerir ponerlos.

## Table name

- Warn: Que no termine con \_tmp / \_temp.

## Table definition

- Error: Engine == InnoDB.
- Warn: Encode y Collate con UTF8 minimo.
- Warn: Se sugiere que tenga algún comentario.

## Fields

- warn: si el varchar es mas grande que x poner text.
- warn: Sugerir que haya el updated_at y el deleted_at.
- warn: Aconsejar de no usar datetime sino timestamp
- warn: que hay null
- warn: que no hay valor por defecto.

## Primary Key

- Error: Debe tener una primary key.
- Error: No debe ser la misma que un unique key o indice.
- Error: Que sea; not null, big int, autoincrement.
- Warn: Puede haber un primary key como uuid/ulid y validar el tipo de dato correcto.

## Foreign Key

- Warn: Sugerir de no tenerlas.
- Error: Que no haya cascade.
- Sugerir que sean \_id.
- Sugerir nombres valido.

## Indices

- Error: No duplicados.
- Error: No confundir con unique key.
- Error: Revisar indice correcto para texto.
- Sugerir nombres valido.
- Indicar de cambiar el tipo de dato de un indice si es datime.
