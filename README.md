<h1 align="center">
	ojsinfo
	<img
		src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png"
		width="50px"
	>
</h1> 

Códigos simples para obter informações do banco de dados do
[Open Journal Systems](https://pkp.sfu.ca/software/ojs/) 3.3x.

# Exemplo simples

```go
	// Importando bibliotecas
	import "github.com/Mathewwss/ojsinfo/DbCfg"
	import "github.com/Mathewwss/ojsinfo/Journals"

	// Configurando SGBD, exemplo mysql
	DbCfg.Db_conf.Driver = "mysql"

	// Login no banco
	DbCfg.Db_conf.Settings = "ojs:ojs@tcp(127.0.0.1:3306)/ojs"

	// Exemplo usando revista
	// www.ojs.com/revista_1 => final-url = revista_1
	// www.ojs.com/revistaazul => final-url = revistaazul
	revista, err := Journals.New("final-url")

	// Obtendo grupos dentro do revista
	err := revista.GetGroups()

	// Mostrando grupos
	fmt.Println(err)
	fmt.Println(revista.Groups)

```

# Tipos

## DbCon

```go
	type DbCon struct {
		Driver string
		Settings string
	}
```

É uma struct composta por dois campos do tipo string.

O campo Driver está relacionado ao SGDB utilizado.

O campo Settings está relacionado as configurações de login no banco,
porta, ip, usuário, senha...

## Journal

```go
	type Journal struct {
		ID int
		Names map[string]string
		Path string
		Groups map[int]map[string]string
		Sections map[int]map[string]string
	}
```

É uma struct que agrupa informações de uma revista.

## Submission

```go
	type Submission struct {
		ID int
		Locale string
		Start string
		// language: title
		Section map[string]string
		// language: title
		Titles map[string]string
		// language: journal name
		JournalNames map[string]string
		// language: [keywords]
		Keywords map[string][]string
		// language: abstract
		Abstract map[string]string
	}
```

É uma struct que agrupa informações de determinada submissão.

## User

```go
	type User struct {
		UID int
		Email string
		Username string
		RealNames map[string]string
		Groups map[string]map[int]string
	}
```

É uma struct que agrupa informações de um usuário.

# Variáveis

## Db_conf

```go
	var Db_conf = DbCon{
		"",
		"",
	}
```

É uma variável do tipo DbCon.

Por padrão os valores ficam vazios.

É de extrema importancia configurar essa variável, senão será
impossível se conectar ao banco de dados.

# Funções

## func New (identity string) (Journal, error)

Função do pacote github.com/Mathewwss/ojsinfo/Journals.

É necessário passar o final da url como parâmetro, se a url é
'www.novarevista/azul' o final da url é 'azul'.

Retorna a variável do tipo Journal e um erro.

Se a final da url estiver correta retorna a variável do tipo
Journal com os campos ID e Path, e o erro vazio.

Senão retorna a variável do tipo Journal vazia, e o erro.

```go
	import "github.com/Mathewwss/ojsinfo/DbCfg"
	import "github.com/Mathewwss/ojsinfo/Journals"

	DbCfg.Db_conf.Driver = "mysql"
	DbCfg.Db_conf.Settings = "ojs:ojs@tcp(127.0.0.1:3306)/ojs"

	revista, err := Journals.New("final-url")

	fmt.Println(revista, err)
```

## func (j *Journal) GetGroups () error

Função do pacote github.com/Mathewwss/ojsinfo/Journals.

Mostra o ID e o nome de cada grupo, em todos os idiomas, de uma revista.

```go
	import "github.com/Mathewwss/ojsinfo/DbCfg"
	import "github.com/Mathewwss/ojsinfo/Journals"

	DbCfg.Db_conf.Driver = "mysql"
	DbCfg.Db_conf.Settings = "ojs:ojs@tcp(127.0.0.1:3306)/ojs"

	revista, err := Journals.New("final-url")
	err := revista.GetGroups()

	fmt.Println(revista.Groups, err)
```

## func (j *Journal) GetNames () error

Função do pacote github.com/Mathewwss/ojsinfo/Journals.

Mostra o nome oficial da revista, em todos idiomas.

```go
	import "github.com/Mathewwss/ojsinfo/DbCfg"
	import "github.com/Mathewwss/ojsinfo/Journals"

	DbCfg.Db_conf.Driver = "mysql"
	DbCfg.Db_conf.Settings = "ojs:ojs@tcp(127.0.0.1:3306)/ojs"

	revista, err := Journals.New("final-url")
	err := revista.GetNames()

	fmt.Println(revista.Names, err)
```


