# den
docker-compose down pull and down in a simple way

#### how to get it ?
```apple js
go get -u github.com/fadeAce/den
```
#### how to use it ?
by set env 'DEN' in ~/.bash_profile
or in cmd line underlying
```apple js
# for example an env should be like 
docker-composefilepath=/Users/fadeAce/docker-compose.yml
export DEN=$docker-composefilepath
```
#### what a usage for ?
```
den
```
that's all , the tool will monitor a docker-compose down docker-compose pull and docker-compose up -d by a sequence automatically by a readable progress tip