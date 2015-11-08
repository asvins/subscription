# subscription
A fresh new service for asvins

## Usage

`go install`
`subscripton`

## Endpoints

`GET /subscription/show?email="john.doe@example.com" // scopes: subscription OR 'user AND owner'`
`GET /subscription/list?page=1 // 10 subscriptions per page, scopes: subscriptions`
`POST /subscription/new // scopes: user OR subscription`
`POST /subscription/pay // scopes: user`
`POST /subscriptions/delayed //scopes: subscription`
