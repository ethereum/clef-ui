## Clef UI

UI Implementation for [Clef](https://github.com/ethereum/go-ethereum/tree/master/cmd/clef), the geth-based signer.

It starts the `signer` binary as a subprocess, and communicates with the `signer` over standard input/output. The `signer` opens an external API to receive RPC request. Once a request comes in, the `signer` will then inform the UI via Standard I/O.

## Getting Started
This UI uses a QT binding for Go to create the UI. To start building, you will first need to install 
[thereceipe/qt](https://github.com/therecipe/qt/wiki/Installation), which comes with a CLI for compiling QT codes, and a WYSIWIG editor for QML. The installation process takes about 25 to 35 minutes.

## Environment Setup
After installing *thereceipe/qt*, you will need to install the dependencies by running:
```cgo
make deps
```

Since this is the first time you are building the UI, you will then need to compile all QT-related code by running the following. In the future, you will need to run this again if you change any QT-related code.  
```cgo
make deploy
```

You can then start the UI by running:
```cgo
make run
```

## License
This is licensed under [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html)




 
