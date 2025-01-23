# Coding test SuperIndo:

Buatlah sebuah API dengan endpoint /product untuk menambahkan dan mengambil data product super
indo, dengan spesifikasi sebagai berikut :

- [x] Dapat melakukan penambahkan data product
- [x] Dapat menampilkan list data product
- [x] Dapat melakukan pencarian bedasarkan nama dan id product
- [x] Dapat melakukan filter produk berdasarkan tipe produk Sayuran, Protein, Buah dan Snack
- [x] Dapat melakukan sorting berdasarkan tanggal, harga dan nama product

## Tech Stack :

- [x] Language : Golang
- [x] Database : SQL / NoSQL + Seeder + migration
- [x] Cache : Redis
- [ ] Dependency Injection : wire (Optional)
- [x] Unittest (Optional)
- [x] Docker (Optional)

## Documentation

1.  `/product` API

    - method **POST** is used for creating new product. the payload must followed this format

      ```json
      {
        "name": "kopi luwak",
        "type": "Snack",
        "price": 10000
      }
      ```

      if the product is created will get status **201**

      ```json
      {
        "data": {
          "id": 168
        }
      }
      ```

      if the product is exist in db with status **409**

      ```json
      {
        "error": "product already exist"
      }
      ```

    - method **GET** is used for retrieving list of product. This method can combined with query params with 5 possible values

      1. search
         is used for searching by id or name

         example

         ```
         /product?search=semangka
         ```

      2. sorts
         is used for sorting the data, can be used only for name, price, and type. The for mat of sorts is `key` + ":" + `asc` or `key` + ":" + `desc`

         example

         ```
         /product?sorts=updated_at:asc&sorts=name:desc&sorts=price:asc
         ```

      3. types
         is used for filtering the data, can be used with the value of product types

         example

         ```
         /product?types=buah&types=snack
         ```

      4. page
         is used for splitting the data with page

         example

         ```
         /product?page=1
         ```

      5. limit
         is used for limiting the data each page
         example

         ```
         /product?page=1
         ```

      and the response will be

      ```json
      {
        "data": {
          "total_data": 2,
          "total_page": 2,
          "products": [
            {
              "id": 168,
              "name": "kopi luwak",
              "price": 10000,
              "type": "snack",
              "updated_at": "2025-01-23T10:51:05.445274Z"
            },
            {
              "id": 167,
              "name": "kopi aDasaassa",
              "price": 10000,
              "type": "snack",
              "updated_at": "2025-01-23T10:39:33.187086Z"
            }
          ]
        }
      }
      ```

    - another method will be rejected with status **405**

      ```json
      {
        "error": "invalid method"
      }
      ```
