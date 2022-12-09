package fixtures

var FailBindGetByIDsRequest = `{
    "ids" : "1",
    "filter" : {
        "startDate" : "",
        "endDate" : "",
        "level" : 0,
        "isDone" : null,
        "showTreeData" : true
    }
}`

var SuccessGetByIDsRequest = `{
    "ids" : [1,5],
    "filter" : {
        "startDate" : "",
        "endDate" : "",
        "level" : 0,
        "isDone" : null,
        "showTreeData" : true
    }
}`
