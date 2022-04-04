```bash
# Instalando dependências do go.mod
$ go get

# Inicializando OCR do golang
$ docker run -it --rm otiai10/gosseract

# Sair do terminal do docker
$ exit

# Corrigindo erro
$ apt-cache search tesseract | grep dev
$ sudo apt-get install libtesseract-dev

# Ferramenta OCR C++
$ sudo apt install tesseract-ocr
$ sudo apt install libtesseract-dev

# Instalando dependências da ferramenta de OCR
$ sudo apt-get install g++ # or clang++ (presumably)
$ sudo apt-get install autoconf automake libtool
$ sudo apt-get install pkg-config
$ sudo apt-get install libpng-dev
$ sudo apt-get install libjpeg8-dev
$ sudo apt-get install libtiff5-dev
$ sudo apt-get install zlib1g-dev
```
