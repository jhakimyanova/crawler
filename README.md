# Crawler
A simple web crawler visiting EBAY web page and collecting item's data 

A crawler written in Golang visits the following page:

https://www.ebay.com/sch/garlandcomputer/m.html

From there it extracts the title, price, product URL and condition (new/pre-owned) information of each
listed item and stores the results in a folder called “data”.
Results are individual files containing a JSON with the data defined above.
Each product URL follows this format:
https://www.ebay.com/itm/234365295029?hash=item36b1a09854:g:czcAAOSwHoRlFfxf

In which the ITEM ID is: 234365295029
Item ID is used as the filename for the result file.
For example: the file 234365295029.json should contain a json formatted in this way:
```json
{
    "title": "Zócalo procesador de CPU Intel Core i7-3820 3,6 GHz 4 núcleos LGA2011 __ SR0LD",
    "condition": "De segunda mano",
    "price": "MXN $768.99",
    "product_url": "https://www.ebay.com/itm/234365295029?hash=item36914299b5:g:oo8AAOSwTXxkHNLV"
}
```
The crawler uses Colly: the Golang Scraping Framework which facilitates breaking through EBAY anti-scraping protection.

The crawler writes the result files asynchronously. 
Pagination is implemented by following next page links found on the visited page.

Condition filtering is supported by providing an optional comman-line parameter to only crawl items
in a specific condition (New, Pre-Owned, etc).
```shell
% ./crawler -help
Usage of ./crawler:
  -condition value
        Condition: new, used, unknown
```

## Building the project
Execute the following command to build the project:
```
make build
```
Execute the following command to run the project:
```
./crawler
```


