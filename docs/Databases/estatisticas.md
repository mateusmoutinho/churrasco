# `Estatisticas`

`sandbox/internal/databases/estatisticas/`, keys under `estatisticas`. Build one with
`estatisticas.New(sandbox)` — it touches no key, so building one is free.

## `visita`

| Field | Type | Required | Target |
| --- | --- | --- | --- |
| `criadoem` | `int` | yes |  |

## `calculo`

| Field | Type | Required | Target |
| --- | --- | --- | --- |
| `adultos` | `int` | yes |  |
| `criancas` | `int` |  |  |
| `carneg` | `int` |  |  |
| `cervejas` | `int` |  |  |
| `criadoem` | `int` | yes |  |

## Methods

| Method | What it does |
| --- | --- |
| `AddVisita(props VisitaNew) (VisitaItem, error)` | inserts one visita record |
| `FindVisitaById(id int64) (VisitaItem, bool)` | reads one visita record by its permanent id |
| `ListVisita(filtrage VisitaFiltrage) ([]VisitaItem, error)` | reads every visita record the filtrage keeps |
| `PageVisita(position int, chunk int) ([]VisitaItem, error)` | reads one page of visita records, counted from 1 |
| `CountVisita() (int, error)` | is how many visita records are live |
| `UpdateVisitaCriadoem(id int64, value int64) error` | writes a new criadoem on one visita record |
| `RemoveVisita(id int64) error` | deletes one visita record and everything nested under it |
| `AddCalculo(props CalculoNew) (CalculoItem, error)` | inserts one calculo record |
| `FindCalculoById(id int64) (CalculoItem, bool)` | reads one calculo record by its permanent id |
| `ListCalculo(filtrage CalculoFiltrage) ([]CalculoItem, error)` | reads every calculo record the filtrage keeps |
| `PageCalculo(position int, chunk int) ([]CalculoItem, error)` | reads one page of calculo records, counted from 1 |
| `CountCalculo() (int, error)` | is how many calculo records are live |
| `UpdateCalculoAdultos(id int64, value int64) error` | writes a new adultos on one calculo record |
| `UpdateCalculoCriancas(id int64, value int64) error` | writes a new criancas on one calculo record |
| `UpdateCalculoCarneg(id int64, value int64) error` | writes a new carneg on one calculo record |
| `UpdateCalculoCervejas(id int64, value int64) error` | writes a new cervejas on one calculo record |
| `UpdateCalculoCriadoem(id int64, value int64) error` | writes a new criadoem on one calculo record |
| `RemoveCalculo(id int64) error` | deletes one calculo record and everything nested under it |

A query this page does not list goes in `methods_custom.go`, hand-written beside these and
rewritten by no build.

[every database](doc.md)
