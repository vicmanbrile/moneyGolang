# Money program in golang

## Archivo estado de cuenta

```json
{
    "wallets": {
        "average" : 240,
        "cash":  700.5,
        "banking":  0
    },
    "debts" : [
        {
            "name": "Gabriela León",
            "amount" :  1168.66,
            "days" : 120
        },
        {
            "name": "Victor Briseño",
            "amount" : 2926.15,
            "days" : 60
        }
    ],
    "suscriptions": [
        {
            "name": "Spotify",
            "type": "monthly",
            "pricing": 60
        },
        {
            "name": "YouTube",
            "type": "monthly",
            "pricing": 180
        },
        {
            "name": "Dentista",
            "type": "monthly",
            "pricing": 700
        },
        {
            "name": "Platzi",
            "type": "monthly",
            "pricing": 580
        },
        {
            "name": "Barbero",
            "type": "monthly",
            "pricing": 300
        },
        {
            "name": "Telcel",
            "type": "monthly",
            "pricing": 300
        }
    ],
    "percentile":[
        {
            "name":"Gabriela León",
            "percentage": 0.2
        }
    ]
}
```

## CLI

```zsh
 go run . 
            -file <Filename.json>
            "gastado.csv" "guardados.csv" "extras.csv"

```
