# Crawler
A simple web crawler visits the eBay web page and collects data on items.

A crawler, written in Golang, visits the following page to extract data:

`https://www.ebay.com/sch/garlandcomputer/m.html`

From there, it extracts the title, price, product URL, and condition (new or pre-owned) information for each listed item and stores the results in a folder named 'data.'
The results are stored as individual files, each containing a JSON file with the data defined above. Each product URL follows this format:
`https://www.ebay.com/itm/234365295029?hash=item36b1a09854:g:czcAAOSwHoRlFfxf`

In this format, the ITEM ID is: 234365295029. The Item ID is used as the filename for the result file. For example, the file named 234365295029.json should contain JSON formatted in this way:
```json
{
    "title": "Zócalo procesador de CPU Intel Core i7-3820 3,6 GHz 4 núcleos LGA2011 __ SR0LD",
    "condition": "De segunda mano",
    "price": "MXN $768.99",
    "product_url": "https://www.ebay.com/itm/234365295029?hash=item36914299b5:g:oo8AAOSwTXxkHNLV"
}
```
The crawler utilizes Colly, the Golang Scraping Framework, which facilitates bypassing eBay's anti-scraping protection
The crawler writes the result files asynchronously.
Pagination is implemented by following the 'next page' links found on the visited pages.

Condition filtering is supported by providing an optional command-line parameter to crawl only items in a specific condition (New, Pre-Owned, etc.).
```shell
% ./crawler -help            
Usage of ./crawler:
  -condition value
        Specifies the condition of the product. Possible values: any, new, used, unknown. Default: any
```

## Build the project
Execute the following command to build the project:
```
make build
```
## Run the project
Execute the following command to run the project:
```
./crawler
```


