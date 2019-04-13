## APIDOC
**Endpoint** - '/search/find'  
Required
- *query* string: Search Query Term
- *username* string
- *token* string: A valid JWT
```json
Success
{
    "result": [
        {
            "created_by": "johnwick",
            "created_on": "2019-04-07T17:02:36.362Z",
            "hashtags": [],
            "id": "c5ae9d75-2d03-4714-a198-2609fca51a00",
            "likes": 0,
            "posts": [
                "ca7e18fc-b488-4435-be21-8c762fb3944e"
            ],
            "title": "Live Capsule"
        },
        {
            "created_by": "johnwick",
            "created_on": "2019-04-10T12:30:16.093Z",
            "hashtags": [],
            "id": "e6409a37-a403-4fd6-8738-fb0ec5ba9d06",
            "likes": 0,
            "posts": [
                "32b4ed9b-bf79-4968-aa48-07073355ec76",
                "1f3cb937-752b-4703-b62d-900a0d878a56",
                "4be5a719-208e-4fea-b8ce-ad91f32ef150",
                "eeef6da6-cfc1-4f51-95c7-57b9e851bd02",
                "96b3434c-a02a-4e48-bcd3-e863d91b5333",
                "5e5db671-f1e8-4bb5-aae4-364f5f0ce36e",
                "068b96ab-b2ae-4a31-ba4f-6523709e2c35",
                "5b1472df-a5fc-44d5-af66-55f3dab3edfa",
                "0f0aec15-b2f3-4379-8fd0-9237031f1fab",
                "0eef73c3-544b-49a0-8760-705b398d92a4",
                "a45cbd45-5535-416d-93e2-4cba22997e1f",
                "73b8d650-fb53-4eb7-aff1-59a24a945d74",
                "90ca85a4-7ef1-4dcd-accc-3b3e1dfdc86a",
                "50c3449c-2a32-46f3-8e3d-5dd18d8db725"
            ],
            "title": "Search Test"
        }
    ],
    "token": "dGhvcg==.WlhpYWtnWHNTTXRhSmp3elltc3BWSFV1Q1llREJ1TWo=.Z0JxZVAzOEUyVjFEVmVDZXV6K3ZKbVY1TnU5MDR3NDdqMmtydmtOdjdDUT0="
}

Error
{ "error" : <ERROR>,
  "token": <NEWTOKEN>
}
The error could be "Could not run Changefeeds".