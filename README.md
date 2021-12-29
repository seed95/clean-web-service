# clean-web-service
A simple web service implemented with clean architecture.
This implementation is with the help of KianooshAz tutorials
and its videos inside
[YouTube](https://www.youtube.com/watch?v=iQNxYVt5ZYY&list=PLVdhomlRFDQzduTUFnI9oD7VzvU89ADMX).

## Table of Contents
**[Log](#Log)**  
**[Translate](#Translate)**  
**[CLI](#Command-Line-Interface)**  
**[Server](#Server)**  

### Log
Used [logrus](https://github.com/sirupsen/logrus) package for log and
[file-rotate](https://github.com/lestrrat-go/file-rotatelogs) package
for writing with rotating and have lifetime and maximum size for file.  

### Translate
Used [i18n](https://github.com/nicksnyder/go-i18n) package for the 
translator.

### Command Line Interface
Used [cobra](https://github.com/spf13/cobra) package for creating modern CLI applications.

### Server
Used below packages for the web service and restApi.
* [echo](https://echo.labstack.com/)
* [gin](https://github.com/gin-gonic/gin)
