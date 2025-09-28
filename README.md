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

## Lintuobjekti (Bird)

```json
{
  "nimi": "Varis",
  "tieteellinenNimi": "Corvus corone",
  "koko": {
    "pituus_cm": 50,
    "siipivali_cm": 100,
    "paino_g": 500
  },
  "varitys": "Musta",
  "elinymparisto": ["metsä", "kaupunki"],
  "ravinto": ["hyönteiset", "siemenet"],
  "uhanalaisuus": "Elinvoimainen (LC)"
}
```

## TEHDYN KOHDAT

Step 1 & 2
