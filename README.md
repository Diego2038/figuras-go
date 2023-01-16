### Module figuras prueba consumo módulo de terceros
Para consumirlo realizar el siguiente procedimiento:<br/>
En el caso no haya creado el módulo dentro del código de Go

```
go mod init ${inserte_nombre}
```

<br>
Una vez tenga el archivo go.mod, añadir este módulo de forma externa con el siguiente comando en la terminal:<br>

```
go get github.com/Diego2038/figuras-go
```
<br>
Después, en el archivo .go donde se vaya a utilizar insertar en el import: <br>

```
import (
  "github.com/Diego2038/figuras-go"
)
```
<br>
De esta manera será posible utilizar el paquete <b>figuras</b> con sus respectivas funciones, clases y métodos.