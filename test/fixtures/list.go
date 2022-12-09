package fixtures

var FailBindListRequest = `{
    "filter" : {
        "startDate": "",
        "endDate" : "",
        "isDone" : null
    },
    "pageSpec" : {
        "pageNumber" : "1",
        "itemPerPage" : 2
    }
}`

var SuccessListRequest = `{
    "filter" : {
        "startDate": "",
        "endDate" : "",
        "isDone" : null
    },
    "pageSpec" : {
        "pageNumber" : 1,
        "itemPerPage" : 1
    }
}`
