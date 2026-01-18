<div align="center">
  <a href="https://smurfsatwork.org/papa" target="_blank"><img src="https://smurfsatwork.org/assets/web-app-manifest-512x512.png" width="150" /></a>

  <h1>Lil Papa Smurf</h1>
  <p>
    <strong>Zero-conf papa smurf.</strong>
  </p>
  <p>
    <a href="https://goreportcard.com/report/github.com/SmurfsAtWork/lilpapa"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/SmurfsAtWork/lilpapa"/></a>
  </p>
</div>

## About

**Lil Papa Smurf** is the minified version of [Papa Smurf](https://github.com/SmurfsAtWork/papa) designed to run without the fuss, no configuration, no urls, no nothing, just compile and run and connect smurfs to it and roll.

Lil Papa is supposed to be ran on local networks, where using Papa is a bit of an overkill as everything is on the same network, as Papa's use-case is for distributed smurfs across different newtworks.

## Key differences from Papa

| Diff     | Papa                              | Lil Papa                     |
| -------- | --------------------------------- | ---------------------------- |
| Users    | Multiple users                    | Single admin user            |
| Network  | Anywhere in the world             | Local network                |
| Binaries | Multiple binaries for each server | Single binary for everything |
| Database | MariaDB                           | SQLite3                      |
| Cache    | Redis                             | Memory Cache                 |
| Events   | Standalone events server          | Async task with http server  |
| CDN      | Standalone CDN server             | Async task with http server  |

## Contributing

IDK, it would be really nice of you to contribute, check the poorly written [CONTRIBUTING.md](/CONTRIBUTING.md) for more info.

## Roadmap

- [x] Add admin user
- [ ] Register Smurf
- [ ] Add Smurfs' command(s)
- [ ] Update Smurfs' configuration
- [ ] Add downloadable programs
- [ ] Download programs at Smurfs
- [ ] Push/Pull Smurfs' command(s) logs
- [ ] Push/Pull Smurfs' system status
- [ ] Reboot Smurfs

## Run locally

1. Clone the repo.

```bash
git clone https://github.com/SmurfsAtWork/lilpapa
```

2. Create the docker environment file

```bash
cp .env.example .env.docker
```

4. Run it with docker compose.

```bash
docker compose up -f docker-compose.yml
```

5. Run [Azrael](https://github.com/SmurfsAtWork/azrael) with API URL as [localhost:26171](http://localhost:26171) and create a couple of smurfs placeholders.

6. Login your [Smurf](https://github.com/SmurfsAtWork/smurf)s and set them to work.

---

Made with 🧉 by [Baraa Al-Masri](https://mbaraa.com)
