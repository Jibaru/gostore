<h3 align="center">
  <img src="https://raw.githubusercontent.com/Jibaru/gostore/main/assets/images/logo.png" width="300" alt="Gostore Logo"/><br/>
</h3>

<p align="center"><i>Simplistic and minimalist storage.</i></p>

## Get started

Download the latest version of gostore for your os.
Make a directory at the same level of the binary.

```bash
$ mkdir storage
$ ls
  gostore_1.0.0_windows_amd64.exe*  storage/
```

Open a terminal an execute the application.
Then, you can use the features' endpoints to interact with the storage api.

```bash
$ curl --location 'localhost:80/api/v1/buckets' \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "my bucket"
  }'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   109  100    82  100    27    821    270 --:--:-- --:--:-- --:--:--  1101{"id":"6893ee7a-d370-4120-8735-9be2329c788a","name":"my bucket","parent_id":null}

$ ls storage/
  17a0dba2-4e19-4acb-a890-d7283c659e81/  buckets.json
```

## Flags

You can customize the execution using the following flags:

```
  -f string
        storage folder name (alias) (default "storage")
  -folder string
        storage folder name (default "storage")
  -h string
        host (alias) (default "localhost")
  -help
        shows help
  -host string
        host (default "localhost")
  -p int
        port (alias) (default 80)
  -port int
        port (default 80)
```

## Features

| Feature                                 | Method | Endpoint (/api/v1)           |
|-----------------------------------------|--------|------------------------------|
| List Buckets                            | GET    | /buckets                     |
| Create a Bucket                         | POST   | /buckets                     |
| Upload an Object to a Bucket            | POST   | /buckets/{bucketID}/objects  |
| List Buckets in a first level of Bucket | GET    | /buckets/{bucketID}/buckets  |
| List Objects in a first level of Bucket | GET    | /buckets/{bucketID}/objects  |
| Download an Object                      | GET    | /objects/{objectID}/download |
| Delete an Object                        | DELETE | /objects/{objectID}          |
| Delete a Bucket                         | DELETE | /buckets/{bucketID}          |

## Repositories Implementations

Store the information about buckets and objects as a metadata to process some requests simple as possible.

| Implementation | Description                |
|----------------|----------------------------|
| In Memory      | Uses server's RAM          |
| File           | Uses a JSON representation |

## How it works?

**Gostore** save information into a local storage folder, keeping the structure with parents, but using uuids as a names.
Buckets are stored as folders, objects are stored as files.

When a request has started, the program give the control to the router, that send the request to some controller.
Then, a service handle the action, interact with the filesystem (if necessary), and then save or update the metadata about the buckets or objects.
This metadata helps the program to get information fast.

**Gostore** was made following hexagonal architecture with some principles of DDD.

```mermaid
sequenceDiagram
    main-->>+CreateBucketCtrl: Handle(request)
    CreateBucketCtrl-->>+CreateBucketServ: Do(name)
    CreateBucketServ-->>+Filesystem: MakeDirectory(path)
    Filesystem-->>-CreateBucketServ: error
    CreateBucketServ-->>+Bucket: new(id, name, parentID)
    Bucket-->>-CreateBucketServ: pointer
    CreateBucketServ-->>+BucketRepository: Save(bucket)
    BucketRepository-->>-CreateBucketServ: error
    CreateBucketServ-->>-CreateBucketCtrl: bucket, error
    CreateBucketCtrl-->-main: response
```
