# Fullcycle
## Features

- server
- client

> Nos meus testes aqui, essa api externa de cotação (https://economia.awesomeapi.com.br/json/last/USD-BRL)
> está com response time médio >= 350 ms, então esse timeout no context de 200ms parece muito alto, 
> sugestão é subir ela para uns 500ms.


## How to use

Usei args para separar o server do client:

Para rodar o server:
```sh
go run maing. server
```

Para rodar o client:
```sh
go run maing. client
```