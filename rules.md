# Rules

Cada check un codigo, asi se puede hacer una descripcion del caso y tambien se puede ignorar.

Puede que el nivel de severidad este entre Warning y Error, ver si hay mejor forma. --ignore-code=101,103

Interesante sería a futuro que se pueda conectarse al mysql y validar:

- Sino hay rows.
- Que no tiene operaciones sobre ella.
- Cardinalidad de indices.
- Buscar queries relacionadas y revisar tiempos.

## 100 File

- **101:** Right permissions.
- **102:** Should to be have .sql extension.
- **103:** Imposible to parser.
- **104:** Tiene \n EOF?
- **105:** Is in UTF8 format.

## 200 Object Name

- Warn: Que no tenga puntos.
- Warn: Que no empieze por \_.
- Warn: Que no supere cierta longitud.
- Warn: Todo en minuscula.
- Warn: Si la palabra es muy larga y no hay _ sugerir ponerlos.
- Warn: Que no termine con \_tmp / \_temp.
- Is a inglish word.?

## 300 Table definition

- Error: Engine == InnoDB.
- Warn: Encode y Collate con UTF8 minimo.
- Warn: Se sugiere que tenga algún comentario.
- La sentencia no termina con ;

## 400 Fields

- warn: si el varchar es mas grande que x poner text.
- warn: Sugerir que haya el updated_at y el deleted_at.
- warn: Aconsejar de no usar datetime sino timestamp
- warn: que hay null
- warn: que no hay valor por defecto.
- y si es menor a 50 que sea char.

## 500 Primary Key

- Error: Debe tener una primary key.
- Error: No debe ser la misma que un unique key o indice.
- Error: Que sea; not null, big int, autoincrement.
- Warn: Puede haber un primary key como uuid/ulid y validar el tipo de dato correcto.

## 600 Foreign Key

- Warn: Sugerir de no tenerlas.
- Error: Que no haya cascade.
- Sugerir que sean \_id.
- Sugerir nombres valido.
- Que el nombre termine en \_uq

## 700 Indices

- Error: No duplicados.
- Error: No confundir con unique key.
- Error: Revisar indice correcto para texto.
- Sugerir nombres valido.
- Que el nombre termine en \_id
