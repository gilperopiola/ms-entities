# Entities Microservice

## Microservice that handles simple 1 table entities used for different things, kind of a wildcard

### Endpoints

* Create Entity -> **POST** -> /Entities
* Search Entities -> **GET** -> /Entities
* Get Entity -> **GET** -> /Entities/*id_entity*
* Update Entity -> **PUT** -> /Entities/*id_entity*
* Update Entity Importance -> **PUT** -> /Entities/*id_entity*/Importance/*importance*
* Disable Entity -> **DELETE** -> /Entities/*id_entity*