# Go Bird API

LINTU REST API GO + FIBER + MONGODB

## Teknologiat

- **Go** - Ohjelmointikieli
- **Fiber** - Web framework
- **MongoDB** - Tietokanta
- **MongoDB Go Driver** - Tietokantayhteys

## API Endpoints

### Linnut (Birds)

- `GET /birds` - Hae kaikki linnut
- `GET /birds/:id` - Hae lintu ID:llä
- `POST /birds` - Luo uusi lintu
- `PATCH /birds/:id` - Päivitä lintu - testattu patchia päivittämään vain osia linnusta
- `DELETE /birds/:id` - Poista lintu

### Uhanalaisuus

- `GET /uhanalaisuus` - Hae uhanalaisuusluokat

## Tietokanta

- **Tietokanta**: BIRDDEX
- **Collections**:
  - `birds` - Lintujen tiedot
  - `uhanalaisuus` - Uhanalaisuusluokat

### Lintuobjekti (Bird)

```json
{
  "nimi": "Varis",
  "tieteellinenNimi": "Corvus corone",
  "siipivali_cm": 100,
  "varitys": "Musta, metallinhohtoinen",
  "elinymparisto": ["metsät", "pellot", "asutusalueet"],
  "ravinto": ["hyönteiset", "pienet nisäkkäät", "siemenet", "jätteet"],
  "uhanalaisuus_id": 1
}
```

### Uhanalaisuus

```json
{
  "uhanalaisuus_id": 1,
  "nimi": "Elinvoimainen",
  "lyhenne": "LC"
}
```

## TEHDYT KOHDAT

Step 1 & 2
