# Sale Notifier
## Get notifications when games go on sale in the eShop.
Developed in [Go](https://go.dev/).

To build the ios library, run:
```console
cd appactions/
gomobile bind -target ios -o ../out/Appactions.xcframework
```
This will output the library to `out/` for use in [sale-notifier-app](https://github.com/AntonioFigueiredo/sale-notifier-app).
