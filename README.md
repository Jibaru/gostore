<h3 align="center">
  <img src="https://raw.githubusercontent.com/Jibaru/gostore/main/assets/images/logo.png" width="300" alt="Gostore Logo"/><br/>
</h3>

<p align="center"><i>Simplistic and minimalist storage.</i></p>

## Features

| Feature                                 | Method | Endpoint                     |
|-----------------------------------------|--------|------------------------------|
| List Buckets                            | GET    | /buckets                     |
| Create a Bucket                         | POST   | /buckets                     |
| Upload an Object to a Bucket            | POST   | /buckets/{bucketID}/objects  |
| List Buckets in a first level of Bucket | GET    | /buckets/{bucketID}/buckets  |
| List Objects in a first level of Bucket | GET    | /buckets/{bucketID}/objects  |
| Download an Object                      | GET    | /objects/{objectID}/download |

# Database Implementations

Store the information about buckets and objects as a metadata to process some requests simple as possible.

| Implementation | Description                |
|----------------|----------------------------|
| In Memory      | Uses server's RAM          |
| File           | Uses a JSON representation |
