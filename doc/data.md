## Critical

```json
const critical = [
    {clr: 'muted',  text: 'Нет',  value: 1},
    {clr: 'danger', text: 'Да',   value: 2}
]
```

## Priorities
```json
const priorities = [
    {icon: 'arrow-down ar-icon-trivial',       value: 1, text: 'trivial'},
    {icon: 'angle-double-down ar-icon-minor',  value: 2, text: 'minor'},
    {icon: 'angle-double-up ar-icon-major',    value: 3, text: 'major'},
    {icon: 'arrow-circle-up ar-icon-critical', value: 4, text: 'critical'},
    {icon: 'ban ar-icon-blocker',              value: 5, text: 'blocker'}, 
  ]
```  

## Resolutions
```json
const resolutions = [
  {text: 'Fixed',            value: 1},
  {text: 'Won\'t fix',       value: 2},
  {text: 'Duplicate',        value: 3},
  {text: 'Incomplete',       value: 4},
  {text: 'Can\'t reproduce', value: 5},
  {text: 'Won\'t do',        value: 6},
]
```

### description
* Fixed Исправлено исправление этой проблемы.
* Won't Fix Эта проблема не будет исправлена, например, она больше не может быть релевантной.
* Duplicate Эта проблема является дубликатом существующей проблемы. Примечание: рекомендуется создать ссылку на дублируемую проблему.
* Incomplete Существует не достаточно информации , чтобы работать над этим вопросом.
* Cannot Reproduce Эта проблема не может быть воспроизведена в настоящее время или недостаточно информации для воспроизведения проблемы. Если появится больше информации, повторите попытку.
* Won't Do Не будет - эта проблема не будет предпринята. (Это разрешение такое же, как Will not Fix, и доступно только для программных проектов по умолчанию)

## Types
```json
export default [
    {text: 'bug',         value: 1},
    {text: 'proposal',    value: 2},
    {text: 'task',        value: 3},
    {text: 'enhancment',  value: 4},
]
```


## statuses
```json
const statuses = [
    { clr:'primary', text: 'Open' ,          value: 1},
    { clr:'warning', text: 'In progress',    value: 2},
    { clr:'success', text: 'Resolved',       value: 3},
    { clr:'info',    text: 'Reopened',       value: 4},
    { clr:'danger',  text: 'Cancelled',      value: 5},
    { clr:'light',   text: 'Closed',         value: 6},
    // { clr:'success', text: 'В прогрессе',    value: 7},
    // { clr:'dark',    text: '[Can be named]', value: 8}
  ]
```
















