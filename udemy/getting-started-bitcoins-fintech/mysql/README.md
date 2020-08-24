## start mysql container

```
$ docker build -t golang-mysql .

$ docker run -it --name golang-mysql -d -p 3306:3306 golang-mysql
```
