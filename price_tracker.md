# Names?
- mpt
- metr
- meprtr


# what is going to do?

im going to give some meli link and start tracking it.


# what im going to use?

Im thinking in go with go-colly for an easy way to get the info of the page and cobra for the cli commands.

# where to persist de data?

For now I only think in a JSON file that keeps the thing i want to track, an id, the link and the price.

```json
[
    {
        "id": 12332,
        "name": "bordeadora"
        "links": [
            {
                "link": "meli.com.ar/track-1",
                "price": 150000.65
            }
            {
                "link": "meli.com.ar/track-2",
                "price": 150000.65
            }
            {
                "link": "meli.com.ar/track-3",
                "price": 110000.65
            }
        ]
    }
]

```


# what commands im going to use?

## mpt new

receive a parameter with the new item i want to track. Optionaly can accept the link(s) to track.
mpt new bordeadora --links link1 link2 link3


## mpt add 

add a new link(s) to track to the existing item
mpt add 1 --links

## mpt list

print the list of items tracked and the link with the lower price. --all shows all the information.
mpt list: list the items and the links with the lower price.
| Id | Name | Link | Price |
| --------------- | --------------- | --------------- | --------------- |
| 1 | Bordeadora | https://mercadolibre.com.ar/... | 549.12 |
| 2 | Cortadora de Pelo | https://mercadolibre.com.ar/... | 120.12 |
| 3 | Pileta | https://mercadolibre.com.ar/... | 12312.12 |

mpt list --all: list the items and all the links.

| Id | Name | Link | Price |
| --------------- | --------------- | --------------- | --------------- |
| 1 | Bordeadora | https://mercadolibre.com.ar/... | 549.12 |
| 1 | Bordeadora | https://mercadolibre.com.ar/... | 549.12 |
| 1 | Bordeadora | https://mercadolibre.com.ar/... | 549.12 |
| 2 | Cortadora de Pelo | https://mercadolibre.com.ar/... | 120.12 |
| 2 | Cortadora de Pelo | https://mercadolibre.com.ar/... | 120.12 |
| 3 | Pileta | https://mercadolibre.com.ar/... | 12312.12 |


mpt list [id]

| Id | Name | Link | Price |
| --------------- | --------------- | --------------- | --------------- |
| 1 | Bordeadora | https://mercadolibre.com.ar/... | 549.12 |
mpt list id --all

| Id | Name | Link | Price |
| --------------- | --------------- | --------------- | --------------- |
| 1 | Bordeadora | https://mercadolibre.com.ar/... | 549.12 |
| 1 | Bordeadora | https://mercadolibre.com.ar/... | 549.12 |
| 1 | Bordeadora | https://mercadolibre.com.ar/... | 549.12 |
## mpt track

starts tracking the elements of the list. can pass a parameter of the id of an item and only track that one.
