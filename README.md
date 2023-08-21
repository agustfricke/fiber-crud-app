## CRUD WEB APP CON FIBER Y DOCKER

### Instalar y usar

- para hacerlo mas practico y no tener errores con VCS, vamos a eliminar el folder .git
- les dejo los archivo .env y .air.toml ya listos pero es de buena practica no subirlos a GitHub

```bash
git clone https://github.com/agustfricke/fiber-crud-app.git
cd fiber-crud-app
sudo rm -r .git
```

```bash
sudo docker compose build
sudo docker compose up
```


