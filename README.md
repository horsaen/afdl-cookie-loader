# afdl-cookie-loader

sister project to [afreeca-downloader](https://github.com/horsaen/afreeca-downloader)

### PandaTV Cookies

PandaTV requires getting the cookies from the browser itself. In order to do this, you need to manually edit the cookie `sessKey`'s `Expires/Max-Age` value to anything other than `Session` to ensure the cookie is written to disk, rather than kept in memory.

Setting the value of anything works, more testing is needed, but setting the value to anything such as `fasdfjhksfkjlawe` works fine.

Note: please do this on one browser only, the tool will attempt to find all instances of Panda's `sessKey`.