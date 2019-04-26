## APIDOC
**Endpoint** - '/search/find'  
Required
- *query* string: Search Query Term
- *username* string
- *token* string: A valid JWT
```json
Success
{
    "result": {
        "profile": [
            {
                "FName": "thor",
                "LName": "odinson",
                "UName": "thor",
                "Email": "test@test.com",
                "ProfilePic": {
                    "Link": "https://ipfs.io/ipfs/QmUuW4aAFXdAoLNXBFqVMDjpZNk1912x18LTggbsydh9jP",
                    "CreatedOn": "",
                    "CreatedBy": "thor",
                    "UploadedOn": "2019-04-26T14:37:13.023+05:30",
                    "Manufacturer": "",
                    "Model": ""
                },
                "Followers": [
                    "thor"
                ],
                "Following": [
                    "ac491",
                    "divya21raj"
                ],
                "FollowersCount": 1,
                "FollowingCount": 2,
                "TravelCapsules": [
                    {
                        "Id": "309a95c4-e132-4e7c-b922-a138dfa352ae",
                        "Title": "Ragnarok",
                        "Posts": [
                            "15b64971-91aa-4bc1-9ad1-24745a97d68e",
                            "ff1e53f5-4ca4-4f4a-9f0c-9c164ba2578c",
                            "8fc19aa6-9c24-4b15-8c08-6d5ce7bb4157"
                        ],
                        "CreatedOn": "2019-04-11T11:51:42.915Z",
                        "CreatedBy": "thor",
                        "UpdatedOn": "2019-04-26T11:38:41.354+05:30",
                        "Hashtags": [],
                        "Likes": 0,
                        "ProfilePic": ""
                    },
                    {
                        "Id": "5fdda452-57d5-42c1-9bf4-b22e05b115af",
                        "Title": "Respawn",
                        "Posts": [],
                        "CreatedOn": "2019-04-25T17:17:22.371Z",
                        "CreatedBy": "thor",
                        "UpdatedOn": "2019-04-25T17:17:22.371Z",
                        "Hashtags": [],
                        "Likes": 0,
                        "ProfilePic": "https://i.pinimg.com/originals/bd/93/3a/bd933a1b384c808326201c104c24ee6d.png"
                    },
                    {
                        "Id": "6118e1b5-e091-48a3-acc0-66a35ab9d515",
                        "Title": "Capcom",
                        "Posts": [
                            "5bc8f523-f009-4257-b67f-cb4e69aac371",
                            "dc290f58-12e2-457b-9ef9-1eb1113286f8",
                            "c1e5122a-5267-49d6-a086-3a7627686c8a",
                            "a69ba25f-022c-4490-abd9-8452c4184dc1",
                            "964fe57a-0608-46db-8b81-64761c4b926b",
                            "08f6e3c2-073c-4208-bc72-083cf70ab434",
                            "543105af-eb85-4a8c-afba-cad53bce0ae6",
                            "b8b79838-ca02-47d5-be91-89c2c54c8267",
                            "74900418-bafe-4fce-831c-a5f3bf36be50",
                            "f0afc238-ff98-4936-8001-fb7222e1376c",
                            "ad275790-43bd-44ab-9942-8a7c2edfa507",
                            "e3d01660-2980-46fc-ae23-5c2c6546161f",
                            "48700a85-cde9-494c-ba96-7d2ef3a1733b",
                            "1aab8b4e-f2ee-46ed-8a25-e160615c81d4"
                        ],
                        "CreatedOn": "2019-04-23T17:16:55.177Z",
                        "CreatedBy": "thor",
                        "UpdatedOn": "2019-04-25T17:10:53.18Z",
                        "Hashtags": [],
                        "Likes": 0,
                        "ProfilePic": ""
                    },
                    {
                        "Id": "0732162d-a4c4-4b1c-aa65-93ed1042a087",
                        "Title": "Bethesda",
                        "Posts": [
                            "241ac94e-e909-4104-b238-5a1a9df42bc2",
                            "67bfbb1e-a73e-4e48-a6c4-fe8eea1a196e"
                        ],
                        "CreatedOn": "2019-04-23T22:48:51.418+05:30",
                        "CreatedBy": "thor",
                        "UpdatedOn": "2019-04-25T13:47:54.534Z",
                        "Hashtags": [],
                        "Likes": 0,
                        "ProfilePic": ""
                    },
                    {
                        "Id": "a6614c96-b86d-40cb-a8bd-ed0593c1cb58",
                        "Title": "Rockstar",
                        "Posts": [],
                        "CreatedOn": "2019-04-24T09:07:12.511Z",
                        "CreatedBy": "thor",
                        "UpdatedOn": "0001-01-01T00:00:00Z",
                        "Hashtags": [],
                        "Likes": 0,
                        "ProfilePic": ""
                    }
                ],
                "Images": [
                    "https://ipfs.io/ipfs/QmcteUnUzoPQ52NaCfjEpEQk5CSBsHdW7bJJNi8Yt1KY8y"
                ]
            }
        ],
        "post": [
            {
                "Id": "383b5b9a-e9ea-4f00-bdc8-9b67e009a497",
                "Title": "Search Test 5",
                "Posts": [
                    "1331ceb9-e0ed-4c28-b710-d4b1ec584d4b",
                    "44f51848-4df2-47cc-9005-34b9e801f38e",
                    "8d67167a-f8c1-4240-9b16-1f8dc8303c8b",
                    "898aa3a3-1312-4bfe-8a17-fd8d2570f725"
                ],
                "CreatedOn": "2019-04-11T08:46:50.213Z",
                "CreatedBy": "johnwick",
                "UpdatedOn": "2019-04-26T17:11:50.783Z",
                "Hashtags": [],
                "Likes": 0,
                "ProfilePic": ""
            }
        ]
    },
    "token": ".QnpnYmFpQ01SQWpXd2hUSGN0Y3VBeGh4S1FGRGFGcEw=.ZkQ5U1NkSTNwL2FBcGk4Qk50T3JXTzA5ZzczQnBMNlVNYXE5QTV2aHBwST0="
}

Error
{ "error" : <ERROR>,
  "token": <NEWTOKEN>
}
The error could be "Could not run Changefeeds".