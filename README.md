# Tax-Calculator
A simple web application to calculate tax using Go language.

## Requirement
- [Git](https://git-scm.com/downloads)
- [Docker](https://docs.docker.com/install/) & [Docker Compose](https://docs.docker.com/compose/install/)

## Installation
Run the following command
```
git clone https://github.com/rl404/tax-calculator
cd tax-calculator
docker-compose up
```
Web can be access from **[http://localhost:9001](http://localhost:9001)**

## API Documentation
- ### Create a New Bill
    Send a list of tax name, tax code, and price and return a bill with detailed tax.
    - **URL**: `/api/createbill`
    - **Method**: `POST`
    - **Required data**: List of tax model
         <details>
         <summary>Example</summary>
         <pre>
         [{
            "name": "Lucky Stretch",
            "taxcode": 2,
            "price": 1000
         }, 
         {
            "name": "Big Mac",
            "taxcode": 1,
            "price": 1000
         }, 
         {
            "name": "Movie",
            "taxcode": 3,
            "price": 150
         }] 
         </pre>
         </details>
    - **Response**: Http response json with Bill model
      <details>
         <summary>Example</summary>
         <pre>
         {
             "data": {
                 "billid": 3,
                 "detail": [
                     {
                         "name": "Lucky Stretch",
                         "taxcode": 2,
                         "price": 1000,
                         "type": "Tobacco",
                         "refundable": "no",
                         "tax": 30,
                         "amount": 1030
                     },
                     {
                         "name": "Big Mac",
                         "taxcode": 1,
                         "price": 1000,
                         "type": "Food & Beverage",
                         "refundable": "yes",
                         "tax": 100,
                         "amount": 1100
                     },
                     {
                         "name": "Movie",
                         "taxcode": 3,
                         "price": 150,
                         "type": "Entertainment",
                         "refundable": "no",
                         "tax": 0.5,
                         "amount": 150.5
                     }
                 ],
                 "pricetotal": 2150,
                 "taxtotal": 130.5,
                 "grandtotal": 2280.5,
                 "createddate": 1558276347
             },
             "message": "Success",
             "status": 200
         }
         </pre>
         </details>
- ### Get Bill Detail
    Get bill detail with list of tax
    - **URL**: `/api/getbill`
    - **Method**: `GET`
    - **Required param**: `bill`
        <details>
         <summary>Example</summary>
         <code>/api/getbill?bill=1</code>
        </details>
    - **Response**: Http response json with Bill model
         <details>
         <summary>Example</summary>
         <pre>
         {
             "data": {
                 "billid": 3,
                 "detail": [
                     {
                         "name": "Lucky Stretch",
                         "taxcode": 2,
                         "price": 1000,
                         "type": "Tobacco",
                         "refundable": "no",
                         "tax": 30,
                         "amount": 1030
                     },
                     {
                         "name": "Big Mac",
                         "taxcode": 1,
                         "price": 1000,
                         "type": "Food & Beverage",
                         "refundable": "yes",
                         "tax": 100,
                         "amount": 1100
                     },
                     {
                         "name": "Movie",
                         "taxcode": 3,
                         "price": 150,
                         "type": "Entertainment",
                         "refundable": "no",
                         "tax": 0.5,
                         "amount": 150.5
                     }
                 ],
                 "pricetotal": 2150,
                 "taxtotal": 130.5,
                 "grandtotal": 2280.5,
                 "createddate": 1558276347
             },
             "message": "Success",
             "status": 200
         }
         </pre>
         </details>
