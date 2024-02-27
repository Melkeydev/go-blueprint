## Air - Live Reloading Utility

[Air](https://github.com/cosmtrek/air) is a live-reloading utility designed to enhance the development experience by automatically rebuilding and restarting your Go application whenever changes are detected in the source code.

The Makefile provided in the project repository includes a command make watch, which triggers Air to monitor file changes and initiate rebuilds and restarts as necessary. Additionally, if Air is not installed on your machine, the Makefile provides an option to install it automatically.

Air's `.air.toml` configuration file allows customization of various aspects of its behavior.

## Live Preview

```bash
make watch

  __    _   ___  
 / /\  | | | |_) 
/_/--\ |_| |_| \_ v1.51.0, built with Go go1.22.0

mkdir /home/ujstor/code/blueprint-version-test/ws-test4/tmp
watching .
watching cmd
watching cmd/api
watching cmd/web
watching cmd/web/js
watching internal
watching internal/database
watching internal/server
watching tests
!exclude tmp
building...
make[1]: Entering directory '/home/ujstor/code/blueprint-version-test/ws-test4'
Building...
Processing path: /home/ujstor/code/blueprint-version-test/ws-test4
Generated code for "/home/ujstor/code/blueprint-version-test/ws-test4/cmd/web/base.templ" in 914.556µs
Generated code for "/home/ujstor/code/blueprint-version-test/ws-test4/cmd/web/hello.templ" in 963.157µs
Generated code for 2 templates with 0 errors in 1.274392ms
make[1]: Leaving directory '/home/ujstor/code/blueprint-version-test/ws-test4'
running...
internal/server/routes.go has changed
building...
make[1]: Entering directory '/home/ujstor/code/blueprint-version-test/ws-test4'
Building...
Processing path: /home/ujstor/code/blueprint-version-test/ws-test4
Generated code for "/home/ujstor/code/blueprint-version-test/ws-test4/cmd/web/base.templ" in 907.426µs
Generated code for "/home/ujstor/code/blueprint-version-test/ws-test4/cmd/web/hello.templ" in 1.16142ms
Generated code for 2 templates with 0 errors in 1.527556ms
make[1]: Leaving directory '/home/ujstor/code/blueprint-version-test/ws-test4'
running...
```

Integrating Air into your development workflow alongside the provided Makefile enables a smooth and efficient process for building, testing, and running your Go applications. With automatic live-reloading, you can focus more on coding and less on manual build and restart steps.