| Donnée               | Comment c'est représenté en JSON ?                      | Comment c'est représenté en XML ?                                       |
|:---------------------|:--------------------------------------------------------|:------------------------------------------------------------------------|
| Pays                 | stations[] country                                      | station "country,attr"                                                  |
| Coordonnées          | stations[] location {latitude, longitude}               | station coordinate "lat, attr", "lon, attr"                             |
| Altitude             | stations[] altitude_m                                   | station coordinate "altitude, attr"                                     |
| Modèle de capteur    | stations[] device.type                                  | station hardware "model, attr"                                          |
| Température          | stations[] observations[].temperature_celsius           | station observations observation measure "type=temperature," "chardata" |
| Conditions ciel      | stations[] observations.conditions                      | not exists                                                              |
| Vent                 | stations[] observations.wind {speed_kmh, direction_deg} | station observations observation wind "speed, attr", "direction, attr"  |
| Notes (optionnelles) | stations[] observations.notes                           | station observations observation wind "note, chardata"                  |
