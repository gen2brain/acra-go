## acra-go

Backend for Application Crash Reports for Android - [ACRA](https://github.com/ACRA/acra).

### Installation

    go install -v github.com/gen2brain/acra-go/acra-go

This will install server in `$GOPATH/bin/acra-go`.

### Download

Binary for linux/amd64 is available:

 - [acra-go-1.0](https://github.com/gen2brain/acra-go/releases/download/1.0/acra-go-1.0.tar.gz)

### Setup

#### Server

Server by default listens on port 55000, you can bind it to other port like this:

    acra-go -bind-addr :80

Usually you will want to proxy requests through your existing server:

##### Nginx

```
location /acra {
    proxy_pass         http://127.0.0.1:55000;
}
```

##### Apache

```
ProxyPass /acra http://127.0.0.1:55000/
```

#### Client

ACRA should send reports to `http://example.com/path/send`. Example annotation of your Android `Application` class:

```java
@ReportsCrashes(formUri = "http://example.com/acra/send",
                formUriBasicAuthLogin = "yourusername", // optional
                formUriBasicAuthPassword = "y0uRpa$$w0rd", // optional
                reportType = org.acra.sender.HttpSender.Type.JSON, // recommended
                mode = ReportingInteractionMode.TOAST,
                resToastText = R.string.crash_toast_text)
public class MyApplication extends Application {
...
}
```

### Usage

```
Usage of acra-go:
  -bind-addr string
        Bind address (default ":55000")
  -database-dir string
        Path to database directory (default ".")
  -htpasswd-backend string
        Path to htpasswd file, if empty backend auth is disabled
  -htpasswd-frontend string
        Path to htpasswd file, if empty frontend auth is disabled
  -read-timeout int
        Read timeout (seconds) (default 5)
  -write-timeout int
        Write timeout (seconds) (default 15)
```

### Pages

  * `/`: dashboard of the crashes
  * `/send`: accepts POST and PUT requests from you app and stores report in db
  * `/view`: view all reports, returns json
  * `/view?id=9823648d-20f6-4663-b8b5-f66f9fc97f81`: view a single report, identified by report id

### Screenshot

![screenshot](https://goo.gl/E38Gyw)
