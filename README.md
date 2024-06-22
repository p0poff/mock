# Simple mock service
## Instaling
clone code with GIT and deploy from Docker compose
```bash
git clone ... OR unpack release source code zip (tar.gz)
docker-compose up
```
### Run-time Customization

The container can be customized in runtime by setting environment from docker's command line or as a part of `docker-compose.yml` 
- `FILE_DB` - path to sqlite database, default /srv/data/mock.db
- `FILE_LOG` - path to sqlite database, default /srv/data/mock.log
  
from https://github.com/umputun/baseimage/blob/master/README.md
- `TIME_ZONE` - set container's TZ, default "America/Chicago". For scratch-based `TZ` should be used instead
- `APP_UID` - UID of internal `app` user, default 1001

### Use
```bash
http://localhost:8080/<route>
```
8080 default port, edit `docker-compose.yml` for change it
#### Administration routes
```bash
http://localhost:8080/admin
```
