## APIDOC
**Endpoint** - '/search/find'  
Required
- *query* string: Search Query Term
```json
Success
[
    {
        "created_by": "johnwick",
        "created_on": "2019-04-10T12:30:16.093Z",
        "hashtags": [],
        "id": "383b5b9a-e9ea-4f00-bdc8-9b67e009a497",
        "likes": 2,
        "posts": [
            "1331ceb9-e0ed-4c28-b710-d4b1ec584d4b"
        ],
        "title": "Criterion"
    },
    {
        "created_by": "johnwick",
        "created_on": "2019-04-11T08:46:50.213Z",
        "hashtags": [],
        "id": "e6409a37-a403-4fd6-8738-fb0ec5ba9d06",
        "likes": 0,
        "posts": [
            "32b4ed9b-bf79-4968-aa48-07073355ec76",
            "1f3cb937-752b-4703-b62d-900a0d878a56",
            "4be5a719-208e-4fea-b8ce-ad91f32ef150",
            "eeef6da6-cfc1-4f51-95c7-57b9e851bd02"
        ],
        "title": "Hogmanay"
    }
]

Error
{ "error" : <ERROR>,
  "token": <NEWTOKEN>
}
The error could be "Could not run Changefeeds".