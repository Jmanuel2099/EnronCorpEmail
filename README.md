# Enron e-mail indexer
**`EnronCorpEmail`** is a application that allows you to search information from the ENRON Corp. e-mail database. This database is indexed in the ZincSearch tool.

## Technologies

- Backend: [Go](https://go.dev/)
- API Router: [Chi](https://github.com/go-chi/chi)
- Search Engine: [Zinc](https://github.com/zinclabs/zinc)
- Frontend: [Vue](https://vuejs.org/) + [JavaScript](https://developer.mozilla.org/es/docs/Web/JavaScript)

## Setup

- Download [Enron Corp's email database](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz )

### ZincSearch

- Download and install the [ZincSearch](https://github.com/zinclabs/zinc/releases) tool

- Configure [ZincSearch](https://docs.zincsearch.com/installation/) locally. 
  At this time it is configured for windows

```bash
    set ZINC_FIRST_ADMIN_USER=admin
    set ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123
    mkdir data
    zinc.exe
```

### Running the project

- Backend project

```bash
    cd EnronCorpEmail\\Backend\\cmd
    go run main.go
```

- Frontend project

```bash
    cd EnronCorpEmail\\frontend
    npm run serve
```


