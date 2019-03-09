# Apparat API

Work in progress 

## Endepunkter

### `/employees`

Returnerer en full liste av info om alle ansatte

```
GET /employees
```

### `/employees/:name`

Returnerer en info om en ansatt

```
GET /employees/:name
```

### `/employeenames`

Returnerer en liste av ansattenavn. Navnene i denne listen kan brukes som parameter i kall til `/employees/:name`.

```
GET /employeenames
```

### `/public/profilepictures/:picture-file-name`

Returnerer profilbilde for en ansatt
